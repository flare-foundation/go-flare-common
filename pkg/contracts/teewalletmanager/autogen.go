// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teewalletmanager

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

// ITeeWalletManagerResumeKeyData is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletManagerResumeKeyData struct {
	KeyId uint64
	TeeId common.Address
	Nonce *big.Int
}

// PublicKey is an auto generated low-level Go binding around an user-defined struct.
type PublicKey struct {
	X [32]byte
	Y [32]byte
}

// TeeWalletManagerMetaData contains all meta data concerning the TeeWalletManager contract.
var TeeWalletManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AdminsNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"name\":\"DuplicatedPublicKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdminsThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCosignersThreshold\",\"type\":\"error\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey\",\"name\":\"publicKey\",\"type\":\"tuple\"}],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MultisigThresholdNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"NotAllAdminsConfirmed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"NotAllCosignersConfirmed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughAdmins\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughKeys\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"WalletAdminConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structPublicKey[]\",\"name\":\"adminsPublicKeys\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"adminsThreshold\",\"type\":\"uint64\"}],\"name\":\"WalletAdminsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"WalletCosignerConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"WalletCosignersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"}],\"name\":\"WalletPaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RESUME\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_PAUSING_ADDRESSES\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WALLET_OP_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"closeWalletInitialization\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"confirmAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"confirmCosigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"createWallet\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"enableWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getProjectWalletIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_walletIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAdminsAndThreshold\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_admins\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_adminsThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAdminsPublicKeysAndThreshold\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"_adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"_adminsThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletCosignersAndThreshold\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletProjectId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletStatus\",\"outputs\":[{\"internalType\":\"enumITeeWalletManager.WalletStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"pauseWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structITeeWalletManager.ResumeKeyData[]\",\"name\":\"_keysData\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"_adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"_adminsThreshold\",\"type\":\"uint64\"}],\"name\":\"setAdmins\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"_cosignersThreshold\",\"type\":\"uint64\"}],\"name\":\"setCosigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_pausingAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"setPausingAddresses\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeExtensionRegistry\",\"outputs\":[{\"internalType\":\"contractITeeExtensionRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeMachineRegistry\",\"outputs\":[{\"internalType\":\"contractITeeMachineRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletKeyManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletKeyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletProjectManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletProjectManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"walletCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TeeWalletManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeWalletManagerMetaData.ABI instead.
var TeeWalletManagerABI = TeeWalletManagerMetaData.ABI

// TeeWalletManager is an auto generated Go binding around an Ethereum contract.
type TeeWalletManager struct {
	TeeWalletManagerCaller     // Read-only binding to the contract
	TeeWalletManagerTransactor // Write-only binding to the contract
	TeeWalletManagerFilterer   // Log filterer for contract events
}

// TeeWalletManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeWalletManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeWalletManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeWalletManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeWalletManagerSession struct {
	Contract     *TeeWalletManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeWalletManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeWalletManagerCallerSession struct {
	Contract *TeeWalletManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TeeWalletManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeWalletManagerTransactorSession struct {
	Contract     *TeeWalletManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TeeWalletManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeWalletManagerRaw struct {
	Contract *TeeWalletManager // Generic contract binding to access the raw methods on
}

// TeeWalletManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeWalletManagerCallerRaw struct {
	Contract *TeeWalletManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TeeWalletManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeWalletManagerTransactorRaw struct {
	Contract *TeeWalletManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeWalletManager creates a new instance of TeeWalletManager, bound to a specific deployed contract.
func NewTeeWalletManager(address common.Address, backend bind.ContractBackend) (*TeeWalletManager, error) {
	contract, err := bindTeeWalletManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManager{TeeWalletManagerCaller: TeeWalletManagerCaller{contract: contract}, TeeWalletManagerTransactor: TeeWalletManagerTransactor{contract: contract}, TeeWalletManagerFilterer: TeeWalletManagerFilterer{contract: contract}}, nil
}

// NewTeeWalletManagerCaller creates a new read-only instance of TeeWalletManager, bound to a specific deployed contract.
func NewTeeWalletManagerCaller(address common.Address, caller bind.ContractCaller) (*TeeWalletManagerCaller, error) {
	contract, err := bindTeeWalletManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerCaller{contract: contract}, nil
}

// NewTeeWalletManagerTransactor creates a new write-only instance of TeeWalletManager, bound to a specific deployed contract.
func NewTeeWalletManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeWalletManagerTransactor, error) {
	contract, err := bindTeeWalletManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerTransactor{contract: contract}, nil
}

// NewTeeWalletManagerFilterer creates a new log filterer instance of TeeWalletManager, bound to a specific deployed contract.
func NewTeeWalletManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeWalletManagerFilterer, error) {
	contract, err := bindTeeWalletManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerFilterer{contract: contract}, nil
}

// bindTeeWalletManager binds a generic wrapper to an already deployed contract.
func bindTeeWalletManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeWalletManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletManager *TeeWalletManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletManager.Contract.TeeWalletManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletManager *TeeWalletManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.TeeWalletManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletManager *TeeWalletManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.TeeWalletManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletManager *TeeWalletManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletManager *TeeWalletManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletManager *TeeWalletManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.contract.Transact(opts, method, params...)
}

// RESUME is a free data retrieval call binding the contract method 0x692a1100.
//
// Solidity: function RESUME() view returns(bytes32)
func (_TeeWalletManager *TeeWalletManagerCaller) RESUME(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "RESUME")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUME is a free data retrieval call binding the contract method 0x692a1100.
//
// Solidity: function RESUME() view returns(bytes32)
func (_TeeWalletManager *TeeWalletManagerSession) RESUME() ([32]byte, error) {
	return _TeeWalletManager.Contract.RESUME(&_TeeWalletManager.CallOpts)
}

// RESUME is a free data retrieval call binding the contract method 0x692a1100.
//
// Solidity: function RESUME() view returns(bytes32)
func (_TeeWalletManager *TeeWalletManagerCallerSession) RESUME() ([32]byte, error) {
	return _TeeWalletManager.Contract.RESUME(&_TeeWalletManager.CallOpts)
}

// SETPAUSINGADDRESSES is a free data retrieval call binding the contract method 0xab014c19.
//
// Solidity: function SET_PAUSING_ADDRESSES() view returns(bytes32)
func (_TeeWalletManager *TeeWalletManagerCaller) SETPAUSINGADDRESSES(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "SET_PAUSING_ADDRESSES")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETPAUSINGADDRESSES is a free data retrieval call binding the contract method 0xab014c19.
//
// Solidity: function SET_PAUSING_ADDRESSES() view returns(bytes32)
func (_TeeWalletManager *TeeWalletManagerSession) SETPAUSINGADDRESSES() ([32]byte, error) {
	return _TeeWalletManager.Contract.SETPAUSINGADDRESSES(&_TeeWalletManager.CallOpts)
}

// SETPAUSINGADDRESSES is a free data retrieval call binding the contract method 0xab014c19.
//
// Solidity: function SET_PAUSING_ADDRESSES() view returns(bytes32)
func (_TeeWalletManager *TeeWalletManagerCallerSession) SETPAUSINGADDRESSES() ([32]byte, error) {
	return _TeeWalletManager.Contract.SETPAUSINGADDRESSES(&_TeeWalletManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeWalletManager *TeeWalletManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeWalletManager *TeeWalletManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeWalletManager.Contract.UPGRADEINTERFACEVERSION(&_TeeWalletManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeWalletManager *TeeWalletManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeWalletManager.Contract.UPGRADEINTERFACEVERSION(&_TeeWalletManager.CallOpts)
}

// WALLETOPTYPE is a free data retrieval call binding the contract method 0xc3500300.
//
// Solidity: function WALLET_OP_TYPE() view returns(bytes32)
func (_TeeWalletManager *TeeWalletManagerCaller) WALLETOPTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "WALLET_OP_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WALLETOPTYPE is a free data retrieval call binding the contract method 0xc3500300.
//
// Solidity: function WALLET_OP_TYPE() view returns(bytes32)
func (_TeeWalletManager *TeeWalletManagerSession) WALLETOPTYPE() ([32]byte, error) {
	return _TeeWalletManager.Contract.WALLETOPTYPE(&_TeeWalletManager.CallOpts)
}

// WALLETOPTYPE is a free data retrieval call binding the contract method 0xc3500300.
//
// Solidity: function WALLET_OP_TYPE() view returns(bytes32)
func (_TeeWalletManager *TeeWalletManagerCallerSession) WALLETOPTYPE() ([32]byte, error) {
	return _TeeWalletManager.Contract.WALLETOPTYPE(&_TeeWalletManager.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeWalletManager *TeeWalletManagerSession) FlareSystemsManager() (common.Address, error) {
	return _TeeWalletManager.Contract.FlareSystemsManager(&_TeeWalletManager.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCallerSession) FlareSystemsManager() (common.Address, error) {
	return _TeeWalletManager.Contract.FlareSystemsManager(&_TeeWalletManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeWalletManager *TeeWalletManagerCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeWalletManager *TeeWalletManagerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeWalletManager.Contract.GetAddressUpdater(&_TeeWalletManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeWalletManager.Contract.GetAddressUpdater(&_TeeWalletManager.CallOpts)
}

// GetProjectWalletIds is a free data retrieval call binding the contract method 0x92911c30.
//
// Solidity: function getProjectWalletIds(bytes32 _projectId) view returns(bytes32[] _walletIds)
func (_TeeWalletManager *TeeWalletManagerCaller) GetProjectWalletIds(opts *bind.CallOpts, _projectId [32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getProjectWalletIds", _projectId)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProjectWalletIds is a free data retrieval call binding the contract method 0x92911c30.
//
// Solidity: function getProjectWalletIds(bytes32 _projectId) view returns(bytes32[] _walletIds)
func (_TeeWalletManager *TeeWalletManagerSession) GetProjectWalletIds(_projectId [32]byte) ([][32]byte, error) {
	return _TeeWalletManager.Contract.GetProjectWalletIds(&_TeeWalletManager.CallOpts, _projectId)
}

// GetProjectWalletIds is a free data retrieval call binding the contract method 0x92911c30.
//
// Solidity: function getProjectWalletIds(bytes32 _projectId) view returns(bytes32[] _walletIds)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetProjectWalletIds(_projectId [32]byte) ([][32]byte, error) {
	return _TeeWalletManager.Contract.GetProjectWalletIds(&_TeeWalletManager.CallOpts, _projectId)
}

// GetWalletAdminsAndThreshold is a free data retrieval call binding the contract method 0x3309abd0.
//
// Solidity: function getWalletAdminsAndThreshold(bytes32 _walletId) view returns(address[] _admins, uint64 _adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletAdminsAndThreshold(opts *bind.CallOpts, _walletId [32]byte) (struct {
	Admins          []common.Address
	AdminsThreshold uint64
}, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletAdminsAndThreshold", _walletId)

	outstruct := new(struct {
		Admins          []common.Address
		AdminsThreshold uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Admins = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.AdminsThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetWalletAdminsAndThreshold is a free data retrieval call binding the contract method 0x3309abd0.
//
// Solidity: function getWalletAdminsAndThreshold(bytes32 _walletId) view returns(address[] _admins, uint64 _adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletAdminsAndThreshold(_walletId [32]byte) (struct {
	Admins          []common.Address
	AdminsThreshold uint64
}, error) {
	return _TeeWalletManager.Contract.GetWalletAdminsAndThreshold(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletAdminsAndThreshold is a free data retrieval call binding the contract method 0x3309abd0.
//
// Solidity: function getWalletAdminsAndThreshold(bytes32 _walletId) view returns(address[] _admins, uint64 _adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletAdminsAndThreshold(_walletId [32]byte) (struct {
	Admins          []common.Address
	AdminsThreshold uint64
}, error) {
	return _TeeWalletManager.Contract.GetWalletAdminsAndThreshold(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletAdminsPublicKeysAndThreshold is a free data retrieval call binding the contract method 0xf3f0efb4.
//
// Solidity: function getWalletAdminsPublicKeysAndThreshold(bytes32 _walletId) view returns((bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletAdminsPublicKeysAndThreshold(opts *bind.CallOpts, _walletId [32]byte) (struct {
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
}, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletAdminsPublicKeysAndThreshold", _walletId)

	outstruct := new(struct {
		AdminsPublicKeys []PublicKey
		AdminsThreshold  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AdminsPublicKeys = *abi.ConvertType(out[0], new([]PublicKey)).(*[]PublicKey)
	outstruct.AdminsThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetWalletAdminsPublicKeysAndThreshold is a free data retrieval call binding the contract method 0xf3f0efb4.
//
// Solidity: function getWalletAdminsPublicKeysAndThreshold(bytes32 _walletId) view returns((bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletAdminsPublicKeysAndThreshold(_walletId [32]byte) (struct {
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
}, error) {
	return _TeeWalletManager.Contract.GetWalletAdminsPublicKeysAndThreshold(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletAdminsPublicKeysAndThreshold is a free data retrieval call binding the contract method 0xf3f0efb4.
//
// Solidity: function getWalletAdminsPublicKeysAndThreshold(bytes32 _walletId) view returns((bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletAdminsPublicKeysAndThreshold(_walletId [32]byte) (struct {
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
}, error) {
	return _TeeWalletManager.Contract.GetWalletAdminsPublicKeysAndThreshold(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletCosignersAndThreshold is a free data retrieval call binding the contract method 0x881567e5.
//
// Solidity: function getWalletCosignersAndThreshold(bytes32 _walletId) view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletCosignersAndThreshold(opts *bind.CallOpts, _walletId [32]byte) (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletCosignersAndThreshold", _walletId)

	outstruct := new(struct {
		Cosigners          []common.Address
		CosignersThreshold uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Cosigners = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.CosignersThreshold = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetWalletCosignersAndThreshold is a free data retrieval call binding the contract method 0x881567e5.
//
// Solidity: function getWalletCosignersAndThreshold(bytes32 _walletId) view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletCosignersAndThreshold(_walletId [32]byte) (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	return _TeeWalletManager.Contract.GetWalletCosignersAndThreshold(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletCosignersAndThreshold is a free data retrieval call binding the contract method 0x881567e5.
//
// Solidity: function getWalletCosignersAndThreshold(bytes32 _walletId) view returns(address[] _cosigners, uint64 _cosignersThreshold)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletCosignersAndThreshold(_walletId [32]byte) (struct {
	Cosigners          []common.Address
	CosignersThreshold uint64
}, error) {
	return _TeeWalletManager.Contract.GetWalletCosignersAndThreshold(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletProjectId is a free data retrieval call binding the contract method 0xf8d5d64f.
//
// Solidity: function getWalletProjectId(bytes32 _walletId) view returns(bytes32 _projectId)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletProjectId(opts *bind.CallOpts, _walletId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletProjectId", _walletId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWalletProjectId is a free data retrieval call binding the contract method 0xf8d5d64f.
//
// Solidity: function getWalletProjectId(bytes32 _walletId) view returns(bytes32 _projectId)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletProjectId(_walletId [32]byte) ([32]byte, error) {
	return _TeeWalletManager.Contract.GetWalletProjectId(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletProjectId is a free data retrieval call binding the contract method 0xf8d5d64f.
//
// Solidity: function getWalletProjectId(bytes32 _walletId) view returns(bytes32 _projectId)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletProjectId(_walletId [32]byte) ([32]byte, error) {
	return _TeeWalletManager.Contract.GetWalletProjectId(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 _walletId) view returns(uint8 _status)
func (_TeeWalletManager *TeeWalletManagerCaller) GetWalletStatus(opts *bind.CallOpts, _walletId [32]byte) (uint8, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "getWalletStatus", _walletId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 _walletId) view returns(uint8 _status)
func (_TeeWalletManager *TeeWalletManagerSession) GetWalletStatus(_walletId [32]byte) (uint8, error) {
	return _TeeWalletManager.Contract.GetWalletStatus(&_TeeWalletManager.CallOpts, _walletId)
}

// GetWalletStatus is a free data retrieval call binding the contract method 0xbfb5dcf9.
//
// Solidity: function getWalletStatus(bytes32 _walletId) view returns(uint8 _status)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GetWalletStatus(_walletId [32]byte) (uint8, error) {
	return _TeeWalletManager.Contract.GetWalletStatus(&_TeeWalletManager.CallOpts, _walletId)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeWalletManager *TeeWalletManagerSession) Governance() (common.Address, error) {
	return _TeeWalletManager.Contract.Governance(&_TeeWalletManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCallerSession) Governance() (common.Address, error) {
	return _TeeWalletManager.Contract.Governance(&_TeeWalletManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeWalletManager *TeeWalletManagerSession) GovernanceSettings() (common.Address, error) {
	return _TeeWalletManager.Contract.GovernanceSettings(&_TeeWalletManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeeWalletManager.Contract.GovernanceSettings(&_TeeWalletManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeWalletManager *TeeWalletManagerSession) Implementation() (common.Address, error) {
	return _TeeWalletManager.Contract.Implementation(&_TeeWalletManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCallerSession) Implementation() (common.Address, error) {
	return _TeeWalletManager.Contract.Implementation(&_TeeWalletManager.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeWalletManager *TeeWalletManagerCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeWalletManager *TeeWalletManagerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeWalletManager.Contract.IsExecutor(&_TeeWalletManager.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeWalletManager *TeeWalletManagerCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeWalletManager.Contract.IsExecutor(&_TeeWalletManager.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeWalletManager *TeeWalletManagerCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeWalletManager *TeeWalletManagerSession) ProductionMode() (bool, error) {
	return _TeeWalletManager.Contract.ProductionMode(&_TeeWalletManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeWalletManager *TeeWalletManagerCallerSession) ProductionMode() (bool, error) {
	return _TeeWalletManager.Contract.ProductionMode(&_TeeWalletManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeWalletManager *TeeWalletManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeWalletManager *TeeWalletManagerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeWalletManager.Contract.ProxiableUUID(&_TeeWalletManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeWalletManager *TeeWalletManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeWalletManager.Contract.ProxiableUUID(&_TeeWalletManager.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCaller) TeeExtensionRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "teeExtensionRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeWalletManager *TeeWalletManagerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeWalletManager.Contract.TeeExtensionRegistry(&_TeeWalletManager.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCallerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeWalletManager.Contract.TeeExtensionRegistry(&_TeeWalletManager.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCaller) TeeMachineRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "teeMachineRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeWalletManager *TeeWalletManagerSession) TeeMachineRegistry() (common.Address, error) {
	return _TeeWalletManager.Contract.TeeMachineRegistry(&_TeeWalletManager.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCallerSession) TeeMachineRegistry() (common.Address, error) {
	return _TeeWalletManager.Contract.TeeMachineRegistry(&_TeeWalletManager.CallOpts)
}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCaller) TeeWalletKeyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "teeWalletKeyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeWalletManager *TeeWalletManagerSession) TeeWalletKeyManager() (common.Address, error) {
	return _TeeWalletManager.Contract.TeeWalletKeyManager(&_TeeWalletManager.CallOpts)
}

// TeeWalletKeyManager is a free data retrieval call binding the contract method 0x6b03b74d.
//
// Solidity: function teeWalletKeyManager() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCallerSession) TeeWalletKeyManager() (common.Address, error) {
	return _TeeWalletManager.Contract.TeeWalletKeyManager(&_TeeWalletManager.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCaller) TeeWalletProjectManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "teeWalletProjectManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeWalletManager *TeeWalletManagerSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeeWalletManager.Contract.TeeWalletProjectManager(&_TeeWalletManager.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeWalletManager *TeeWalletManagerCallerSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeeWalletManager.Contract.TeeWalletProjectManager(&_TeeWalletManager.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletManager *TeeWalletManagerCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_TeeWalletManager *TeeWalletManagerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeWalletManager.Contract.TimelockedCalls(&_TeeWalletManager.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletManager *TeeWalletManagerCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeWalletManager.Contract.TimelockedCalls(&_TeeWalletManager.CallOpts, selector)
}

// WalletCounter is a free data retrieval call binding the contract method 0xf64d7060.
//
// Solidity: function walletCounter() view returns(uint256)
func (_TeeWalletManager *TeeWalletManagerCaller) WalletCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeeWalletManager.contract.Call(opts, &out, "walletCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WalletCounter is a free data retrieval call binding the contract method 0xf64d7060.
//
// Solidity: function walletCounter() view returns(uint256)
func (_TeeWalletManager *TeeWalletManagerSession) WalletCounter() (*big.Int, error) {
	return _TeeWalletManager.Contract.WalletCounter(&_TeeWalletManager.CallOpts)
}

// WalletCounter is a free data retrieval call binding the contract method 0xf64d7060.
//
// Solidity: function walletCounter() view returns(uint256)
func (_TeeWalletManager *TeeWalletManagerCallerSession) WalletCounter() (*big.Int, error) {
	return _TeeWalletManager.Contract.WalletCounter(&_TeeWalletManager.CallOpts)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletManager *TeeWalletManagerSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.CancelGovernanceCall(&_TeeWalletManager.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.CancelGovernanceCall(&_TeeWalletManager.TransactOpts, _selector)
}

// CloseWalletInitialization is a paid mutator transaction binding the contract method 0xa3c6d865.
//
// Solidity: function closeWalletInitialization(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) CloseWalletInitialization(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "closeWalletInitialization", _walletId)
}

// CloseWalletInitialization is a paid mutator transaction binding the contract method 0xa3c6d865.
//
// Solidity: function closeWalletInitialization(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerSession) CloseWalletInitialization(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.CloseWalletInitialization(&_TeeWalletManager.TransactOpts, _walletId)
}

// CloseWalletInitialization is a paid mutator transaction binding the contract method 0xa3c6d865.
//
// Solidity: function closeWalletInitialization(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) CloseWalletInitialization(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.CloseWalletInitialization(&_TeeWalletManager.TransactOpts, _walletId)
}

// ConfirmAdmin is a paid mutator transaction binding the contract method 0x9ab3b981.
//
// Solidity: function confirmAdmin(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) ConfirmAdmin(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "confirmAdmin", _walletId)
}

// ConfirmAdmin is a paid mutator transaction binding the contract method 0x9ab3b981.
//
// Solidity: function confirmAdmin(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerSession) ConfirmAdmin(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ConfirmAdmin(&_TeeWalletManager.TransactOpts, _walletId)
}

// ConfirmAdmin is a paid mutator transaction binding the contract method 0x9ab3b981.
//
// Solidity: function confirmAdmin(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) ConfirmAdmin(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ConfirmAdmin(&_TeeWalletManager.TransactOpts, _walletId)
}

// ConfirmCosigner is a paid mutator transaction binding the contract method 0x77f74c1a.
//
// Solidity: function confirmCosigner(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) ConfirmCosigner(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "confirmCosigner", _walletId)
}

// ConfirmCosigner is a paid mutator transaction binding the contract method 0x77f74c1a.
//
// Solidity: function confirmCosigner(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerSession) ConfirmCosigner(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ConfirmCosigner(&_TeeWalletManager.TransactOpts, _walletId)
}

// ConfirmCosigner is a paid mutator transaction binding the contract method 0x77f74c1a.
//
// Solidity: function confirmCosigner(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) ConfirmCosigner(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ConfirmCosigner(&_TeeWalletManager.TransactOpts, _walletId)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x1d647605.
//
// Solidity: function createWallet(bytes32 _projectId) returns(bytes32 _walletId)
func (_TeeWalletManager *TeeWalletManagerTransactor) CreateWallet(opts *bind.TransactOpts, _projectId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "createWallet", _projectId)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x1d647605.
//
// Solidity: function createWallet(bytes32 _projectId) returns(bytes32 _walletId)
func (_TeeWalletManager *TeeWalletManagerSession) CreateWallet(_projectId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.CreateWallet(&_TeeWalletManager.TransactOpts, _projectId)
}

// CreateWallet is a paid mutator transaction binding the contract method 0x1d647605.
//
// Solidity: function createWallet(bytes32 _projectId) returns(bytes32 _walletId)
func (_TeeWalletManager *TeeWalletManagerTransactorSession) CreateWallet(_projectId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.CreateWallet(&_TeeWalletManager.TransactOpts, _projectId)
}

// EnableWallet is a paid mutator transaction binding the contract method 0x676f358c.
//
// Solidity: function enableWallet(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) EnableWallet(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "enableWallet", _walletId)
}

// EnableWallet is a paid mutator transaction binding the contract method 0x676f358c.
//
// Solidity: function enableWallet(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerSession) EnableWallet(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.EnableWallet(&_TeeWalletManager.TransactOpts, _walletId)
}

// EnableWallet is a paid mutator transaction binding the contract method 0x676f358c.
//
// Solidity: function enableWallet(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) EnableWallet(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.EnableWallet(&_TeeWalletManager.TransactOpts, _walletId)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletManager *TeeWalletManagerSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ExecuteGovernanceCall(&_TeeWalletManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.ExecuteGovernanceCall(&_TeeWalletManager.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeWalletManager *TeeWalletManagerSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.Initialise(&_TeeWalletManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.Initialise(&_TeeWalletManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeWalletManager *TeeWalletManagerSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.Initialize(&_TeeWalletManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.Initialize(&_TeeWalletManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// PauseWallet is a paid mutator transaction binding the contract method 0x963a310d.
//
// Solidity: function pauseWallet(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) PauseWallet(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "pauseWallet", _walletId)
}

// PauseWallet is a paid mutator transaction binding the contract method 0x963a310d.
//
// Solidity: function pauseWallet(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerSession) PauseWallet(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.PauseWallet(&_TeeWalletManager.TransactOpts, _walletId)
}

// PauseWallet is a paid mutator transaction binding the contract method 0x963a310d.
//
// Solidity: function pauseWallet(bytes32 _walletId) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) PauseWallet(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.PauseWallet(&_TeeWalletManager.TransactOpts, _walletId)
}

// Resume is a paid mutator transaction binding the contract method 0xa353fe2a.
//
// Solidity: function resume(bytes32 _walletId, (uint64,address,uint256)[] _keysData, address _claimBackAddress) payable returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) Resume(opts *bind.TransactOpts, _walletId [32]byte, _keysData []ITeeWalletManagerResumeKeyData, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "resume", _walletId, _keysData, _claimBackAddress)
}

// Resume is a paid mutator transaction binding the contract method 0xa353fe2a.
//
// Solidity: function resume(bytes32 _walletId, (uint64,address,uint256)[] _keysData, address _claimBackAddress) payable returns()
func (_TeeWalletManager *TeeWalletManagerSession) Resume(_walletId [32]byte, _keysData []ITeeWalletManagerResumeKeyData, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.Resume(&_TeeWalletManager.TransactOpts, _walletId, _keysData, _claimBackAddress)
}

// Resume is a paid mutator transaction binding the contract method 0xa353fe2a.
//
// Solidity: function resume(bytes32 _walletId, (uint64,address,uint256)[] _keysData, address _claimBackAddress) payable returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) Resume(_walletId [32]byte, _keysData []ITeeWalletManagerResumeKeyData, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.Resume(&_TeeWalletManager.TransactOpts, _walletId, _keysData, _claimBackAddress)
}

// SetAdmins is a paid mutator transaction binding the contract method 0x42e1aa9a.
//
// Solidity: function setAdmins(bytes32 _walletId, (bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) SetAdmins(opts *bind.TransactOpts, _walletId [32]byte, _adminsPublicKeys []PublicKey, _adminsThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "setAdmins", _walletId, _adminsPublicKeys, _adminsThreshold)
}

// SetAdmins is a paid mutator transaction binding the contract method 0x42e1aa9a.
//
// Solidity: function setAdmins(bytes32 _walletId, (bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold) returns()
func (_TeeWalletManager *TeeWalletManagerSession) SetAdmins(_walletId [32]byte, _adminsPublicKeys []PublicKey, _adminsThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetAdmins(&_TeeWalletManager.TransactOpts, _walletId, _adminsPublicKeys, _adminsThreshold)
}

// SetAdmins is a paid mutator transaction binding the contract method 0x42e1aa9a.
//
// Solidity: function setAdmins(bytes32 _walletId, (bytes32,bytes32)[] _adminsPublicKeys, uint64 _adminsThreshold) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) SetAdmins(_walletId [32]byte, _adminsPublicKeys []PublicKey, _adminsThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetAdmins(&_TeeWalletManager.TransactOpts, _walletId, _adminsPublicKeys, _adminsThreshold)
}

// SetCosigners is a paid mutator transaction binding the contract method 0x3e00a2d4.
//
// Solidity: function setCosigners(bytes32 _walletId, address[] _cosigners, uint64 _cosignersThreshold) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) SetCosigners(opts *bind.TransactOpts, _walletId [32]byte, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "setCosigners", _walletId, _cosigners, _cosignersThreshold)
}

// SetCosigners is a paid mutator transaction binding the contract method 0x3e00a2d4.
//
// Solidity: function setCosigners(bytes32 _walletId, address[] _cosigners, uint64 _cosignersThreshold) returns()
func (_TeeWalletManager *TeeWalletManagerSession) SetCosigners(_walletId [32]byte, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetCosigners(&_TeeWalletManager.TransactOpts, _walletId, _cosigners, _cosignersThreshold)
}

// SetCosigners is a paid mutator transaction binding the contract method 0x3e00a2d4.
//
// Solidity: function setCosigners(bytes32 _walletId, address[] _cosigners, uint64 _cosignersThreshold) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) SetCosigners(_walletId [32]byte, _cosigners []common.Address, _cosignersThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetCosigners(&_TeeWalletManager.TransactOpts, _walletId, _cosigners, _cosignersThreshold)
}

// SetPausingAddresses is a paid mutator transaction binding the contract method 0xac7355e0.
//
// Solidity: function setPausingAddresses(bytes32 _walletId, address[] _pausingAddresses, address _claimBackAddress) payable returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) SetPausingAddresses(opts *bind.TransactOpts, _walletId [32]byte, _pausingAddresses []common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "setPausingAddresses", _walletId, _pausingAddresses, _claimBackAddress)
}

// SetPausingAddresses is a paid mutator transaction binding the contract method 0xac7355e0.
//
// Solidity: function setPausingAddresses(bytes32 _walletId, address[] _pausingAddresses, address _claimBackAddress) payable returns()
func (_TeeWalletManager *TeeWalletManagerSession) SetPausingAddresses(_walletId [32]byte, _pausingAddresses []common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetPausingAddresses(&_TeeWalletManager.TransactOpts, _walletId, _pausingAddresses, _claimBackAddress)
}

// SetPausingAddresses is a paid mutator transaction binding the contract method 0xac7355e0.
//
// Solidity: function setPausingAddresses(bytes32 _walletId, address[] _pausingAddresses, address _claimBackAddress) payable returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) SetPausingAddresses(_walletId [32]byte, _pausingAddresses []common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SetPausingAddresses(&_TeeWalletManager.TransactOpts, _walletId, _pausingAddresses, _claimBackAddress)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeWalletManager *TeeWalletManagerSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SwitchToProductionMode(&_TeeWalletManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeWalletManager.Contract.SwitchToProductionMode(&_TeeWalletManager.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeWalletManager *TeeWalletManagerSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.UpdateContractAddresses(&_TeeWalletManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.UpdateContractAddresses(&_TeeWalletManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeWalletManager *TeeWalletManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeWalletManager.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeWalletManager *TeeWalletManagerSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.UpgradeToAndCall(&_TeeWalletManager.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeWalletManager *TeeWalletManagerTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeWalletManager.Contract.UpgradeToAndCall(&_TeeWalletManager.TransactOpts, _newImplementation, _data)
}

// TeeWalletManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeWalletManager contract.
type TeeWalletManagerGovernanceCallTimelockedIterator struct {
	Event *TeeWalletManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerGovernanceCallTimelocked)
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
		it.Event = new(TeeWalletManagerGovernanceCallTimelocked)
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
func (it *TeeWalletManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeWalletManager contract.
type TeeWalletManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeWalletManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerGovernanceCallTimelockedIterator{contract: _TeeWalletManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerGovernanceCallTimelocked)
				if err := _TeeWalletManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeWalletManagerGovernanceCallTimelocked, error) {
	event := new(TeeWalletManagerGovernanceCallTimelocked)
	if err := _TeeWalletManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeWalletManager contract.
type TeeWalletManagerGovernanceInitialisedIterator struct {
	Event *TeeWalletManagerGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerGovernanceInitialised)
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
		it.Event = new(TeeWalletManagerGovernanceInitialised)
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
func (it *TeeWalletManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerGovernanceInitialised represents a GovernanceInitialised event raised by the TeeWalletManager contract.
type TeeWalletManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeWalletManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerGovernanceInitialisedIterator{contract: _TeeWalletManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerGovernanceInitialised)
				if err := _TeeWalletManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseGovernanceInitialised(log types.Log) (*TeeWalletManagerGovernanceInitialised, error) {
	event := new(TeeWalletManagerGovernanceInitialised)
	if err := _TeeWalletManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeeWalletManager contract.
type TeeWalletManagerGovernedProductionModeEnteredIterator struct {
	Event *TeeWalletManagerGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerGovernedProductionModeEntered)
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
		it.Event = new(TeeWalletManagerGovernedProductionModeEntered)
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
func (it *TeeWalletManagerGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeWalletManager contract.
type TeeWalletManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeeWalletManagerGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerGovernedProductionModeEnteredIterator{contract: _TeeWalletManager.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerGovernedProductionModeEntered)
				if err := _TeeWalletManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeeWalletManagerGovernedProductionModeEntered, error) {
	event := new(TeeWalletManagerGovernedProductionModeEntered)
	if err := _TeeWalletManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeeWalletManager contract.
type TeeWalletManagerTimelockedGovernanceCallCanceledIterator struct {
	Event *TeeWalletManagerTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeeWalletManagerTimelockedGovernanceCallCanceled)
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
func (it *TeeWalletManagerTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeWalletManager contract.
type TeeWalletManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeeWalletManagerTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerTimelockedGovernanceCallCanceledIterator{contract: _TeeWalletManager.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerTimelockedGovernanceCallCanceled)
				if err := _TeeWalletManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeeWalletManagerTimelockedGovernanceCallCanceled, error) {
	event := new(TeeWalletManagerTimelockedGovernanceCallCanceled)
	if err := _TeeWalletManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeeWalletManager contract.
type TeeWalletManagerTimelockedGovernanceCallExecutedIterator struct {
	Event *TeeWalletManagerTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeeWalletManagerTimelockedGovernanceCallExecuted)
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
func (it *TeeWalletManagerTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeWalletManager contract.
type TeeWalletManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeeWalletManagerTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerTimelockedGovernanceCallExecutedIterator{contract: _TeeWalletManager.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerTimelockedGovernanceCallExecuted)
				if err := _TeeWalletManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeeWalletManagerTimelockedGovernanceCallExecuted, error) {
	event := new(TeeWalletManagerTimelockedGovernanceCallExecuted)
	if err := _TeeWalletManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeeWalletManager contract.
type TeeWalletManagerUpgradedIterator struct {
	Event *TeeWalletManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerUpgraded)
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
		it.Event = new(TeeWalletManagerUpgraded)
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
func (it *TeeWalletManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerUpgraded represents a Upgraded event raised by the TeeWalletManager contract.
type TeeWalletManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeeWalletManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerUpgradedIterator{contract: _TeeWalletManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerUpgraded)
				if err := _TeeWalletManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseUpgraded(log types.Log) (*TeeWalletManagerUpgraded, error) {
	event := new(TeeWalletManagerUpgraded)
	if err := _TeeWalletManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletAdminConfirmedIterator is returned from FilterWalletAdminConfirmed and is used to iterate over the raw logs and unpacked data for WalletAdminConfirmed events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletAdminConfirmedIterator struct {
	Event *TeeWalletManagerWalletAdminConfirmed // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletAdminConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletAdminConfirmed)
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
		it.Event = new(TeeWalletManagerWalletAdminConfirmed)
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
func (it *TeeWalletManagerWalletAdminConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletAdminConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletAdminConfirmed represents a WalletAdminConfirmed event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletAdminConfirmed struct {
	WalletId [32]byte
	Admin    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletAdminConfirmed is a free log retrieval operation binding the contract event 0xefdffb32952334eb7adb154fb07bfa263335e73f8ad33a54aba665efe346cb08.
//
// Solidity: event WalletAdminConfirmed(bytes32 indexed walletId, address indexed admin)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletAdminConfirmed(opts *bind.FilterOpts, walletId [][32]byte, admin []common.Address) (*TeeWalletManagerWalletAdminConfirmedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletAdminConfirmed", walletIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletAdminConfirmedIterator{contract: _TeeWalletManager.contract, event: "WalletAdminConfirmed", logs: logs, sub: sub}, nil
}

// WatchWalletAdminConfirmed is a free log subscription operation binding the contract event 0xefdffb32952334eb7adb154fb07bfa263335e73f8ad33a54aba665efe346cb08.
//
// Solidity: event WalletAdminConfirmed(bytes32 indexed walletId, address indexed admin)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletAdminConfirmed(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletAdminConfirmed, walletId [][32]byte, admin []common.Address) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var adminRule []interface{}
	for _, adminItem := range admin {
		adminRule = append(adminRule, adminItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletAdminConfirmed", walletIdRule, adminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletAdminConfirmed)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletAdminConfirmed", log); err != nil {
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

// ParseWalletAdminConfirmed is a log parse operation binding the contract event 0xefdffb32952334eb7adb154fb07bfa263335e73f8ad33a54aba665efe346cb08.
//
// Solidity: event WalletAdminConfirmed(bytes32 indexed walletId, address indexed admin)
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletAdminConfirmed(log types.Log) (*TeeWalletManagerWalletAdminConfirmed, error) {
	event := new(TeeWalletManagerWalletAdminConfirmed)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletAdminConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletAdminsSetIterator is returned from FilterWalletAdminsSet and is used to iterate over the raw logs and unpacked data for WalletAdminsSet events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletAdminsSetIterator struct {
	Event *TeeWalletManagerWalletAdminsSet // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletAdminsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletAdminsSet)
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
		it.Event = new(TeeWalletManagerWalletAdminsSet)
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
func (it *TeeWalletManagerWalletAdminsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletAdminsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletAdminsSet represents a WalletAdminsSet event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletAdminsSet struct {
	WalletId         [32]byte
	AdminsPublicKeys []PublicKey
	AdminsThreshold  uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterWalletAdminsSet is a free log retrieval operation binding the contract event 0xfac536eb8a80dfc8ca75bf6543d4b1f435e119551bafd3733fcb28850d29d8aa.
//
// Solidity: event WalletAdminsSet(bytes32 indexed walletId, (bytes32,bytes32)[] adminsPublicKeys, uint64 adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletAdminsSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletManagerWalletAdminsSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletAdminsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletAdminsSetIterator{contract: _TeeWalletManager.contract, event: "WalletAdminsSet", logs: logs, sub: sub}, nil
}

// WatchWalletAdminsSet is a free log subscription operation binding the contract event 0xfac536eb8a80dfc8ca75bf6543d4b1f435e119551bafd3733fcb28850d29d8aa.
//
// Solidity: event WalletAdminsSet(bytes32 indexed walletId, (bytes32,bytes32)[] adminsPublicKeys, uint64 adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletAdminsSet(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletAdminsSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletAdminsSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletAdminsSet)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletAdminsSet", log); err != nil {
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

// ParseWalletAdminsSet is a log parse operation binding the contract event 0xfac536eb8a80dfc8ca75bf6543d4b1f435e119551bafd3733fcb28850d29d8aa.
//
// Solidity: event WalletAdminsSet(bytes32 indexed walletId, (bytes32,bytes32)[] adminsPublicKeys, uint64 adminsThreshold)
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletAdminsSet(log types.Log) (*TeeWalletManagerWalletAdminsSet, error) {
	event := new(TeeWalletManagerWalletAdminsSet)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletAdminsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletCosignerConfirmedIterator is returned from FilterWalletCosignerConfirmed and is used to iterate over the raw logs and unpacked data for WalletCosignerConfirmed events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletCosignerConfirmedIterator struct {
	Event *TeeWalletManagerWalletCosignerConfirmed // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletCosignerConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletCosignerConfirmed)
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
		it.Event = new(TeeWalletManagerWalletCosignerConfirmed)
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
func (it *TeeWalletManagerWalletCosignerConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletCosignerConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletCosignerConfirmed represents a WalletCosignerConfirmed event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletCosignerConfirmed struct {
	WalletId [32]byte
	Cosigner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletCosignerConfirmed is a free log retrieval operation binding the contract event 0x10ab8f96b70aef2c884291a57e690b6d9e4d97ca15e4c8b2514626b8d0cd023b.
//
// Solidity: event WalletCosignerConfirmed(bytes32 indexed walletId, address indexed cosigner)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletCosignerConfirmed(opts *bind.FilterOpts, walletId [][32]byte, cosigner []common.Address) (*TeeWalletManagerWalletCosignerConfirmedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var cosignerRule []interface{}
	for _, cosignerItem := range cosigner {
		cosignerRule = append(cosignerRule, cosignerItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletCosignerConfirmed", walletIdRule, cosignerRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletCosignerConfirmedIterator{contract: _TeeWalletManager.contract, event: "WalletCosignerConfirmed", logs: logs, sub: sub}, nil
}

// WatchWalletCosignerConfirmed is a free log subscription operation binding the contract event 0x10ab8f96b70aef2c884291a57e690b6d9e4d97ca15e4c8b2514626b8d0cd023b.
//
// Solidity: event WalletCosignerConfirmed(bytes32 indexed walletId, address indexed cosigner)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletCosignerConfirmed(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletCosignerConfirmed, walletId [][32]byte, cosigner []common.Address) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var cosignerRule []interface{}
	for _, cosignerItem := range cosigner {
		cosignerRule = append(cosignerRule, cosignerItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletCosignerConfirmed", walletIdRule, cosignerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletCosignerConfirmed)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCosignerConfirmed", log); err != nil {
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

// ParseWalletCosignerConfirmed is a log parse operation binding the contract event 0x10ab8f96b70aef2c884291a57e690b6d9e4d97ca15e4c8b2514626b8d0cd023b.
//
// Solidity: event WalletCosignerConfirmed(bytes32 indexed walletId, address indexed cosigner)
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletCosignerConfirmed(log types.Log) (*TeeWalletManagerWalletCosignerConfirmed, error) {
	event := new(TeeWalletManagerWalletCosignerConfirmed)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCosignerConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletCosignersSetIterator is returned from FilterWalletCosignersSet and is used to iterate over the raw logs and unpacked data for WalletCosignersSet events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletCosignersSetIterator struct {
	Event *TeeWalletManagerWalletCosignersSet // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletCosignersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletCosignersSet)
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
		it.Event = new(TeeWalletManagerWalletCosignersSet)
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
func (it *TeeWalletManagerWalletCosignersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletCosignersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletCosignersSet represents a WalletCosignersSet event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletCosignersSet struct {
	WalletId           [32]byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWalletCosignersSet is a free log retrieval operation binding the contract event 0x48419616e657a6824f128cbf23c09b8b4750e1c50999dc05dd869a97ac4a9c31.
//
// Solidity: event WalletCosignersSet(bytes32 indexed walletId, address[] cosigners, uint64 cosignersThreshold)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletCosignersSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletManagerWalletCosignersSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletCosignersSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletCosignersSetIterator{contract: _TeeWalletManager.contract, event: "WalletCosignersSet", logs: logs, sub: sub}, nil
}

// WatchWalletCosignersSet is a free log subscription operation binding the contract event 0x48419616e657a6824f128cbf23c09b8b4750e1c50999dc05dd869a97ac4a9c31.
//
// Solidity: event WalletCosignersSet(bytes32 indexed walletId, address[] cosigners, uint64 cosignersThreshold)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletCosignersSet(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletCosignersSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletCosignersSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletCosignersSet)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCosignersSet", log); err != nil {
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

// ParseWalletCosignersSet is a log parse operation binding the contract event 0x48419616e657a6824f128cbf23c09b8b4750e1c50999dc05dd869a97ac4a9c31.
//
// Solidity: event WalletCosignersSet(bytes32 indexed walletId, address[] cosigners, uint64 cosignersThreshold)
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletCosignersSet(log types.Log) (*TeeWalletManagerWalletCosignersSet, error) {
	event := new(TeeWalletManagerWalletCosignersSet)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCosignersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletCreatedIterator is returned from FilterWalletCreated and is used to iterate over the raw logs and unpacked data for WalletCreated events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletCreatedIterator struct {
	Event *TeeWalletManagerWalletCreated // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletCreated)
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
		it.Event = new(TeeWalletManagerWalletCreated)
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
func (it *TeeWalletManagerWalletCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletCreated represents a WalletCreated event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletCreated struct {
	ProjectId [32]byte
	WalletId  [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletCreated is a free log retrieval operation binding the contract event 0xbe8f27cef1f3d94120c9c547c3614f5b992fdb0c0a497cc920fde06546291ab4.
//
// Solidity: event WalletCreated(bytes32 indexed projectId, bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletCreated(opts *bind.FilterOpts, projectId [][32]byte, walletId [][32]byte) (*TeeWalletManagerWalletCreatedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletCreated", projectIdRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletCreatedIterator{contract: _TeeWalletManager.contract, event: "WalletCreated", logs: logs, sub: sub}, nil
}

// WatchWalletCreated is a free log subscription operation binding the contract event 0xbe8f27cef1f3d94120c9c547c3614f5b992fdb0c0a497cc920fde06546291ab4.
//
// Solidity: event WalletCreated(bytes32 indexed projectId, bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletCreated(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletCreated, projectId [][32]byte, walletId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletCreated", projectIdRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletCreated)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCreated", log); err != nil {
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

// ParseWalletCreated is a log parse operation binding the contract event 0xbe8f27cef1f3d94120c9c547c3614f5b992fdb0c0a497cc920fde06546291ab4.
//
// Solidity: event WalletCreated(bytes32 indexed projectId, bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletCreated(log types.Log) (*TeeWalletManagerWalletCreated, error) {
	event := new(TeeWalletManagerWalletCreated)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletEnabledIterator is returned from FilterWalletEnabled and is used to iterate over the raw logs and unpacked data for WalletEnabled events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletEnabledIterator struct {
	Event *TeeWalletManagerWalletEnabled // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletEnabled)
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
		it.Event = new(TeeWalletManagerWalletEnabled)
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
func (it *TeeWalletManagerWalletEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletEnabled represents a WalletEnabled event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletEnabled struct {
	WalletId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletEnabled is a free log retrieval operation binding the contract event 0x9dd94a650e5c2e19300b404a2e83ff44d26999ca24996bbff99dcd70f81661bf.
//
// Solidity: event WalletEnabled(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletEnabled(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletManagerWalletEnabledIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletEnabled", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletEnabledIterator{contract: _TeeWalletManager.contract, event: "WalletEnabled", logs: logs, sub: sub}, nil
}

// WatchWalletEnabled is a free log subscription operation binding the contract event 0x9dd94a650e5c2e19300b404a2e83ff44d26999ca24996bbff99dcd70f81661bf.
//
// Solidity: event WalletEnabled(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletEnabled(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletEnabled, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletEnabled", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletEnabled)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletEnabled", log); err != nil {
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

// ParseWalletEnabled is a log parse operation binding the contract event 0x9dd94a650e5c2e19300b404a2e83ff44d26999ca24996bbff99dcd70f81661bf.
//
// Solidity: event WalletEnabled(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletEnabled(log types.Log) (*TeeWalletManagerWalletEnabled, error) {
	event := new(TeeWalletManagerWalletEnabled)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletInitializedIterator is returned from FilterWalletInitialized and is used to iterate over the raw logs and unpacked data for WalletInitialized events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletInitializedIterator struct {
	Event *TeeWalletManagerWalletInitialized // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletInitialized)
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
		it.Event = new(TeeWalletManagerWalletInitialized)
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
func (it *TeeWalletManagerWalletInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletInitialized represents a WalletInitialized event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletInitialized struct {
	WalletId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletInitialized is a free log retrieval operation binding the contract event 0x09032cc96c5d9d0239f5528329abee7a1c72f41813ebdd51e4f51f8f962c7223.
//
// Solidity: event WalletInitialized(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletInitialized(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletManagerWalletInitializedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletInitialized", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletInitializedIterator{contract: _TeeWalletManager.contract, event: "WalletInitialized", logs: logs, sub: sub}, nil
}

// WatchWalletInitialized is a free log subscription operation binding the contract event 0x09032cc96c5d9d0239f5528329abee7a1c72f41813ebdd51e4f51f8f962c7223.
//
// Solidity: event WalletInitialized(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletInitialized(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletInitialized, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletInitialized", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletInitialized)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletInitialized", log); err != nil {
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

// ParseWalletInitialized is a log parse operation binding the contract event 0x09032cc96c5d9d0239f5528329abee7a1c72f41813ebdd51e4f51f8f962c7223.
//
// Solidity: event WalletInitialized(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletInitialized(log types.Log) (*TeeWalletManagerWalletInitialized, error) {
	event := new(TeeWalletManagerWalletInitialized)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletManagerWalletPausedIterator is returned from FilterWalletPaused and is used to iterate over the raw logs and unpacked data for WalletPaused events raised by the TeeWalletManager contract.
type TeeWalletManagerWalletPausedIterator struct {
	Event *TeeWalletManagerWalletPaused // Event containing the contract specifics and raw log

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
func (it *TeeWalletManagerWalletPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletManagerWalletPaused)
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
		it.Event = new(TeeWalletManagerWalletPaused)
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
func (it *TeeWalletManagerWalletPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletManagerWalletPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletManagerWalletPaused represents a WalletPaused event raised by the TeeWalletManager contract.
type TeeWalletManagerWalletPaused struct {
	WalletId [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletPaused is a free log retrieval operation binding the contract event 0x5bc3bc472409bb4331cff333acd11b5d5b6d97052d783936292f19955b25eef3.
//
// Solidity: event WalletPaused(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) FilterWalletPaused(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletManagerWalletPausedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.FilterLogs(opts, "WalletPaused", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletManagerWalletPausedIterator{contract: _TeeWalletManager.contract, event: "WalletPaused", logs: logs, sub: sub}, nil
}

// WatchWalletPaused is a free log subscription operation binding the contract event 0x5bc3bc472409bb4331cff333acd11b5d5b6d97052d783936292f19955b25eef3.
//
// Solidity: event WalletPaused(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) WatchWalletPaused(opts *bind.WatchOpts, sink chan<- *TeeWalletManagerWalletPaused, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletManager.contract.WatchLogs(opts, "WalletPaused", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletManagerWalletPaused)
				if err := _TeeWalletManager.contract.UnpackLog(event, "WalletPaused", log); err != nil {
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

// ParseWalletPaused is a log parse operation binding the contract event 0x5bc3bc472409bb4331cff333acd11b5d5b6d97052d783936292f19955b25eef3.
//
// Solidity: event WalletPaused(bytes32 indexed walletId)
func (_TeeWalletManager *TeeWalletManagerFilterer) ParseWalletPaused(log types.Log) (*TeeWalletManagerWalletPaused, error) {
	event := new(TeeWalletManagerWalletPaused)
	if err := _TeeWalletManager.contract.UnpackLog(event, "WalletPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
