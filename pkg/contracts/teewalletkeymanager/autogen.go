// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teewalletkeymanager

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

// ITeeWalletKeyManagerKeyConfigConstants is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletKeyManagerKeyConfigConstants struct {
	AdminsPublicKeys   []PublicKey
	AdminsThreshold    uint64
	Cosigners          []common.Address
	CosignersThreshold uint64
}

// ITeeWalletKeyManagerKeyExistence is an auto generated low-level Go binding around an user-defined struct.
type ITeeWalletKeyManagerKeyExistence struct {
	TeeId           common.Address
	WalletId        [32]byte
	KeyId           uint64
	KeyType         [32]byte
	SigningAlgo     [32]byte
	PublicKey       []byte
	Nonce           *big.Int
	Restored        bool
	ConfigConstants ITeeWalletKeyManagerKeyConfigConstants
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

// TeeWalletKeyManagerMetaData contains all meta data concerning the TeeWalletKeyManager contract.
var TeeWalletKeyManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSettings\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTeeSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyNotGeneratedOnTeeMachine\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyNotRestoredOnTeeMachine\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeIdAlreadyAdded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"name\":\"WalletKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"WalletKeyConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"name\":\"WalletKeyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"keyIds\",\"type\":\"uint64[]\"}],\"name\":\"WalletKeysNotAvailable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"multisigThreshold\",\"type\":\"uint64\"}],\"name\":\"WalletMultisigThresholdSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"KEY_DELETE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KEY_GENERATE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WALLET_OP_TYPE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"addKey\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"cleanUpTeeIds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"signingAlgo\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"restored\",\"type\":\"bool\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"internalType\":\"structPublicKey[]\",\"name\":\"adminsPublicKeys\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64\",\"name\":\"adminsThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"}],\"internalType\":\"structITeeWalletKeyManager.KeyConfigConstants\",\"name\":\"configConstants\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"settingsVersion\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"settings\",\"type\":\"bytes\"}],\"internalType\":\"structITeeWalletKeyManager.KeyExistence\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature\",\"name\":\"_teeSignature\",\"type\":\"tuple\"}],\"name\":\"confirmKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"deleteKey\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"getWalletKeyPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"getWalletKeyTeeIds\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletKeysInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_multisigThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64[]\",\"name\":\"_keyIds\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64\",\"name\":\"_counter\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_teeId\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_keyId\",\"type\":\"uint64\"}],\"name\":\"increaseKeyNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"receivingTeesAndKeys\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"keyId\",\"type\":\"uint64\"}],\"internalType\":\"structTeeIdKeyIdPair[]\",\"name\":\"_teeIdKeyIdPairs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_multisigThreshold\",\"type\":\"uint64\"}],\"name\":\"setMultisigThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeExtensionRegistry\",\"outputs\":[{\"internalType\":\"contractITeeExtensionRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeMachineRegistry\",\"outputs\":[{\"internalType\":\"contractITeeMachineRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletBackupManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletBackupManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teeWalletProjectManager\",\"outputs\":[{\"internalType\":\"contractITeeWalletProjectManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TeeWalletKeyManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeWalletKeyManagerMetaData.ABI instead.
var TeeWalletKeyManagerABI = TeeWalletKeyManagerMetaData.ABI

// TeeWalletKeyManager is an auto generated Go binding around an Ethereum contract.
type TeeWalletKeyManager struct {
	TeeWalletKeyManagerCaller     // Read-only binding to the contract
	TeeWalletKeyManagerTransactor // Write-only binding to the contract
	TeeWalletKeyManagerFilterer   // Log filterer for contract events
}

// TeeWalletKeyManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeWalletKeyManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletKeyManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeWalletKeyManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletKeyManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeWalletKeyManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeWalletKeyManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeWalletKeyManagerSession struct {
	Contract     *TeeWalletKeyManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TeeWalletKeyManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeWalletKeyManagerCallerSession struct {
	Contract *TeeWalletKeyManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// TeeWalletKeyManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeWalletKeyManagerTransactorSession struct {
	Contract     *TeeWalletKeyManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// TeeWalletKeyManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeWalletKeyManagerRaw struct {
	Contract *TeeWalletKeyManager // Generic contract binding to access the raw methods on
}

// TeeWalletKeyManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeWalletKeyManagerCallerRaw struct {
	Contract *TeeWalletKeyManagerCaller // Generic read-only contract binding to access the raw methods on
}

// TeeWalletKeyManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeWalletKeyManagerTransactorRaw struct {
	Contract *TeeWalletKeyManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeWalletKeyManager creates a new instance of TeeWalletKeyManager, bound to a specific deployed contract.
func NewTeeWalletKeyManager(address common.Address, backend bind.ContractBackend) (*TeeWalletKeyManager, error) {
	contract, err := bindTeeWalletKeyManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManager{TeeWalletKeyManagerCaller: TeeWalletKeyManagerCaller{contract: contract}, TeeWalletKeyManagerTransactor: TeeWalletKeyManagerTransactor{contract: contract}, TeeWalletKeyManagerFilterer: TeeWalletKeyManagerFilterer{contract: contract}}, nil
}

// NewTeeWalletKeyManagerCaller creates a new read-only instance of TeeWalletKeyManager, bound to a specific deployed contract.
func NewTeeWalletKeyManagerCaller(address common.Address, caller bind.ContractCaller) (*TeeWalletKeyManagerCaller, error) {
	contract, err := bindTeeWalletKeyManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerCaller{contract: contract}, nil
}

// NewTeeWalletKeyManagerTransactor creates a new write-only instance of TeeWalletKeyManager, bound to a specific deployed contract.
func NewTeeWalletKeyManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeWalletKeyManagerTransactor, error) {
	contract, err := bindTeeWalletKeyManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerTransactor{contract: contract}, nil
}

// NewTeeWalletKeyManagerFilterer creates a new log filterer instance of TeeWalletKeyManager, bound to a specific deployed contract.
func NewTeeWalletKeyManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeWalletKeyManagerFilterer, error) {
	contract, err := bindTeeWalletKeyManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerFilterer{contract: contract}, nil
}

// bindTeeWalletKeyManager binds a generic wrapper to an already deployed contract.
func bindTeeWalletKeyManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeWalletKeyManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletKeyManager *TeeWalletKeyManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletKeyManager.Contract.TeeWalletKeyManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletKeyManager *TeeWalletKeyManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.TeeWalletKeyManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletKeyManager *TeeWalletKeyManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.TeeWalletKeyManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeWalletKeyManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.contract.Transact(opts, method, params...)
}

// KEYDELETE is a free data retrieval call binding the contract method 0x84b17286.
//
// Solidity: function KEY_DELETE() view returns(bytes32)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) KEYDELETE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "KEY_DELETE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KEYDELETE is a free data retrieval call binding the contract method 0x84b17286.
//
// Solidity: function KEY_DELETE() view returns(bytes32)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) KEYDELETE() ([32]byte, error) {
	return _TeeWalletKeyManager.Contract.KEYDELETE(&_TeeWalletKeyManager.CallOpts)
}

// KEYDELETE is a free data retrieval call binding the contract method 0x84b17286.
//
// Solidity: function KEY_DELETE() view returns(bytes32)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) KEYDELETE() ([32]byte, error) {
	return _TeeWalletKeyManager.Contract.KEYDELETE(&_TeeWalletKeyManager.CallOpts)
}

// KEYGENERATE is a free data retrieval call binding the contract method 0xe33529bd.
//
// Solidity: function KEY_GENERATE() view returns(bytes32)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) KEYGENERATE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "KEY_GENERATE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KEYGENERATE is a free data retrieval call binding the contract method 0xe33529bd.
//
// Solidity: function KEY_GENERATE() view returns(bytes32)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) KEYGENERATE() ([32]byte, error) {
	return _TeeWalletKeyManager.Contract.KEYGENERATE(&_TeeWalletKeyManager.CallOpts)
}

// KEYGENERATE is a free data retrieval call binding the contract method 0xe33529bd.
//
// Solidity: function KEY_GENERATE() view returns(bytes32)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) KEYGENERATE() ([32]byte, error) {
	return _TeeWalletKeyManager.Contract.KEYGENERATE(&_TeeWalletKeyManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeWalletKeyManager.Contract.UPGRADEINTERFACEVERSION(&_TeeWalletKeyManager.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeeWalletKeyManager.Contract.UPGRADEINTERFACEVERSION(&_TeeWalletKeyManager.CallOpts)
}

// WALLETOPTYPE is a free data retrieval call binding the contract method 0xc3500300.
//
// Solidity: function WALLET_OP_TYPE() view returns(bytes32)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) WALLETOPTYPE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "WALLET_OP_TYPE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WALLETOPTYPE is a free data retrieval call binding the contract method 0xc3500300.
//
// Solidity: function WALLET_OP_TYPE() view returns(bytes32)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) WALLETOPTYPE() ([32]byte, error) {
	return _TeeWalletKeyManager.Contract.WALLETOPTYPE(&_TeeWalletKeyManager.CallOpts)
}

// WALLETOPTYPE is a free data retrieval call binding the contract method 0xc3500300.
//
// Solidity: function WALLET_OP_TYPE() view returns(bytes32)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) WALLETOPTYPE() ([32]byte, error) {
	return _TeeWalletKeyManager.Contract.WALLETOPTYPE(&_TeeWalletKeyManager.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) FlareSystemsManager() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.FlareSystemsManager(&_TeeWalletKeyManager.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) FlareSystemsManager() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.FlareSystemsManager(&_TeeWalletKeyManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.GetAddressUpdater(&_TeeWalletKeyManager.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.GetAddressUpdater(&_TeeWalletKeyManager.CallOpts)
}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GetWalletKeyPublicKey(opts *bind.CallOpts, _walletId [32]byte, _keyId uint64) ([]byte, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "getWalletKeyPublicKey", _walletId, _keyId)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GetWalletKeyPublicKey(_walletId [32]byte, _keyId uint64) ([]byte, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyPublicKey(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyPublicKey is a free data retrieval call binding the contract method 0xbb825ac3.
//
// Solidity: function getWalletKeyPublicKey(bytes32 _walletId, uint64 _keyId) view returns(bytes _publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GetWalletKeyPublicKey(_walletId [32]byte, _keyId uint64) ([]byte, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyPublicKey(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GetWalletKeyTeeIds(opts *bind.CallOpts, _walletId [32]byte, _keyId uint64) ([]common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "getWalletKeyTeeIds", _walletId, _keyId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GetWalletKeyTeeIds(_walletId [32]byte, _keyId uint64) ([]common.Address, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyTeeIds(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeyTeeIds is a free data retrieval call binding the contract method 0x08480213.
//
// Solidity: function getWalletKeyTeeIds(bytes32 _walletId, uint64 _keyId) view returns(address[] _teeIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GetWalletKeyTeeIds(_walletId [32]byte, _keyId uint64) ([]common.Address, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeyTeeIds(&_TeeWalletKeyManager.CallOpts, _walletId, _keyId)
}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint64[] _keyIds, uint64 _counter)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GetWalletKeysInfo(opts *bind.CallOpts, _walletId [32]byte) (struct {
	MultisigThreshold uint64
	KeyIds            []uint64
	Counter           uint64
}, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "getWalletKeysInfo", _walletId)

	outstruct := new(struct {
		MultisigThreshold uint64
		KeyIds            []uint64
		Counter           uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MultisigThreshold = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.KeyIds = *abi.ConvertType(out[1], new([]uint64)).(*[]uint64)
	outstruct.Counter = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint64[] _keyIds, uint64 _counter)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GetWalletKeysInfo(_walletId [32]byte) (struct {
	MultisigThreshold uint64
	KeyIds            []uint64
	Counter           uint64
}, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeysInfo(&_TeeWalletKeyManager.CallOpts, _walletId)
}

// GetWalletKeysInfo is a free data retrieval call binding the contract method 0x8338292b.
//
// Solidity: function getWalletKeysInfo(bytes32 _walletId) view returns(uint64 _multisigThreshold, uint64[] _keyIds, uint64 _counter)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GetWalletKeysInfo(_walletId [32]byte) (struct {
	MultisigThreshold uint64
	KeyIds            []uint64
	Counter           uint64
}, error) {
	return _TeeWalletKeyManager.Contract.GetWalletKeysInfo(&_TeeWalletKeyManager.CallOpts, _walletId)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) Governance() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.Governance(&_TeeWalletKeyManager.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) Governance() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.Governance(&_TeeWalletKeyManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) GovernanceSettings() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.GovernanceSettings(&_TeeWalletKeyManager.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.GovernanceSettings(&_TeeWalletKeyManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) Implementation() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.Implementation(&_TeeWalletKeyManager.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) Implementation() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.Implementation(&_TeeWalletKeyManager.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeWalletKeyManager.Contract.IsExecutor(&_TeeWalletKeyManager.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeeWalletKeyManager.Contract.IsExecutor(&_TeeWalletKeyManager.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) ProductionMode() (bool, error) {
	return _TeeWalletKeyManager.Contract.ProductionMode(&_TeeWalletKeyManager.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) ProductionMode() (bool, error) {
	return _TeeWalletKeyManager.Contract.ProductionMode(&_TeeWalletKeyManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeWalletKeyManager.Contract.ProxiableUUID(&_TeeWalletKeyManager.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeeWalletKeyManager.Contract.ProxiableUUID(&_TeeWalletKeyManager.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) TeeExtensionRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "teeExtensionRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.TeeExtensionRegistry(&_TeeWalletKeyManager.CallOpts)
}

// TeeExtensionRegistry is a free data retrieval call binding the contract method 0xa435d58a.
//
// Solidity: function teeExtensionRegistry() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) TeeExtensionRegistry() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.TeeExtensionRegistry(&_TeeWalletKeyManager.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) TeeMachineRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "teeMachineRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) TeeMachineRegistry() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.TeeMachineRegistry(&_TeeWalletKeyManager.CallOpts)
}

// TeeMachineRegistry is a free data retrieval call binding the contract method 0x524967d7.
//
// Solidity: function teeMachineRegistry() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) TeeMachineRegistry() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.TeeMachineRegistry(&_TeeWalletKeyManager.CallOpts)
}

// TeeWalletBackupManager is a free data retrieval call binding the contract method 0xf1362c31.
//
// Solidity: function teeWalletBackupManager() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) TeeWalletBackupManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "teeWalletBackupManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletBackupManager is a free data retrieval call binding the contract method 0xf1362c31.
//
// Solidity: function teeWalletBackupManager() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) TeeWalletBackupManager() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.TeeWalletBackupManager(&_TeeWalletKeyManager.CallOpts)
}

// TeeWalletBackupManager is a free data retrieval call binding the contract method 0xf1362c31.
//
// Solidity: function teeWalletBackupManager() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) TeeWalletBackupManager() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.TeeWalletBackupManager(&_TeeWalletKeyManager.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) TeeWalletManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "teeWalletManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) TeeWalletManager() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.TeeWalletManager(&_TeeWalletKeyManager.CallOpts)
}

// TeeWalletManager is a free data retrieval call binding the contract method 0x9ce03893.
//
// Solidity: function teeWalletManager() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) TeeWalletManager() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.TeeWalletManager(&_TeeWalletKeyManager.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) TeeWalletProjectManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "teeWalletProjectManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.TeeWalletProjectManager(&_TeeWalletKeyManager.CallOpts)
}

// TeeWalletProjectManager is a free data retrieval call binding the contract method 0x1e915a85.
//
// Solidity: function teeWalletProjectManager() view returns(address)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) TeeWalletProjectManager() (common.Address, error) {
	return _TeeWalletKeyManager.Contract.TeeWalletProjectManager(&_TeeWalletKeyManager.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _TeeWalletKeyManager.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeWalletKeyManager.Contract.TimelockedCalls(&_TeeWalletKeyManager.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletKeyManager *TeeWalletKeyManagerCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _TeeWalletKeyManager.Contract.TimelockedCalls(&_TeeWalletKeyManager.CallOpts, selector)
}

// AddKey is a paid mutator transaction binding the contract method 0x7875fa5c.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId) payable returns(uint64 _keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) AddKey(opts *bind.TransactOpts, _teeId common.Address, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "addKey", _teeId, _walletId)
}

// AddKey is a paid mutator transaction binding the contract method 0x7875fa5c.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId) payable returns(uint64 _keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) AddKey(_teeId common.Address, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.AddKey(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId)
}

// AddKey is a paid mutator transaction binding the contract method 0x7875fa5c.
//
// Solidity: function addKey(address _teeId, bytes32 _walletId) payable returns(uint64 _keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) AddKey(_teeId common.Address, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.AddKey(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.CancelGovernanceCall(&_TeeWalletKeyManager.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.CancelGovernanceCall(&_TeeWalletKeyManager.TransactOpts, _selector)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0xaec4c979.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) CleanUpTeeIds(opts *bind.TransactOpts, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "cleanUpTeeIds", _walletId, _keyId)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0xaec4c979.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) CleanUpTeeIds(_walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.CleanUpTeeIds(&_TeeWalletKeyManager.TransactOpts, _walletId, _keyId)
}

// CleanUpTeeIds is a paid mutator transaction binding the contract method 0xaec4c979.
//
// Solidity: function cleanUpTeeIds(bytes32 _walletId, uint64 _keyId) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) CleanUpTeeIds(_walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.CleanUpTeeIds(&_TeeWalletKeyManager.TransactOpts, _walletId, _keyId)
}

// ConfirmKey is a paid mutator transaction binding the contract method 0xa02d6fe1.
//
// Solidity: function confirmKey((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) _proof, (uint8,bytes32,bytes32) _teeSignature) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) ConfirmKey(opts *bind.TransactOpts, _proof ITeeWalletKeyManagerKeyExistence, _teeSignature Signature) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "confirmKey", _proof, _teeSignature)
}

// ConfirmKey is a paid mutator transaction binding the contract method 0xa02d6fe1.
//
// Solidity: function confirmKey((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) _proof, (uint8,bytes32,bytes32) _teeSignature) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) ConfirmKey(_proof ITeeWalletKeyManagerKeyExistence, _teeSignature Signature) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ConfirmKey(&_TeeWalletKeyManager.TransactOpts, _proof, _teeSignature)
}

// ConfirmKey is a paid mutator transaction binding the contract method 0xa02d6fe1.
//
// Solidity: function confirmKey((address,bytes32,uint64,bytes32,bytes32,bytes,uint256,bool,((bytes32,bytes32)[],uint64,address[],uint64),bytes32,bytes) _proof, (uint8,bytes32,bytes32) _teeSignature) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) ConfirmKey(_proof ITeeWalletKeyManagerKeyExistence, _teeSignature Signature) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ConfirmKey(&_TeeWalletKeyManager.TransactOpts, _proof, _teeSignature)
}

// DeleteKey is a paid mutator transaction binding the contract method 0x90e31e7e.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) DeleteKey(opts *bind.TransactOpts, _teeId common.Address, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "deleteKey", _teeId, _walletId, _keyId)
}

// DeleteKey is a paid mutator transaction binding the contract method 0x90e31e7e.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) DeleteKey(_teeId common.Address, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.DeleteKey(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId, _keyId)
}

// DeleteKey is a paid mutator transaction binding the contract method 0x90e31e7e.
//
// Solidity: function deleteKey(address _teeId, bytes32 _walletId, uint64 _keyId) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) DeleteKey(_teeId common.Address, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.DeleteKey(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId, _keyId)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ExecuteGovernanceCall(&_TeeWalletKeyManager.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ExecuteGovernanceCall(&_TeeWalletKeyManager.TransactOpts, _selector)
}

// IncreaseKeyNonce is a paid mutator transaction binding the contract method 0xf3d3161c.
//
// Solidity: function increaseKeyNonce(address _teeId, bytes32 _walletId, uint64 _keyId) returns(uint256 _nonce)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) IncreaseKeyNonce(opts *bind.TransactOpts, _teeId common.Address, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "increaseKeyNonce", _teeId, _walletId, _keyId)
}

// IncreaseKeyNonce is a paid mutator transaction binding the contract method 0xf3d3161c.
//
// Solidity: function increaseKeyNonce(address _teeId, bytes32 _walletId, uint64 _keyId) returns(uint256 _nonce)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) IncreaseKeyNonce(_teeId common.Address, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.IncreaseKeyNonce(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId, _keyId)
}

// IncreaseKeyNonce is a paid mutator transaction binding the contract method 0xf3d3161c.
//
// Solidity: function increaseKeyNonce(address _teeId, bytes32 _walletId, uint64 _keyId) returns(uint256 _nonce)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) IncreaseKeyNonce(_teeId common.Address, _walletId [32]byte, _keyId uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.IncreaseKeyNonce(&_TeeWalletKeyManager.TransactOpts, _teeId, _walletId, _keyId)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.Initialise(&_TeeWalletKeyManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.Initialise(&_TeeWalletKeyManager.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.Initialize(&_TeeWalletKeyManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.Initialize(&_TeeWalletKeyManager.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// ReceivingTeesAndKeys is a paid mutator transaction binding the contract method 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,uint64)[] _teeIdKeyIdPairs)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) ReceivingTeesAndKeys(opts *bind.TransactOpts, _walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "receivingTeesAndKeys", _walletId)
}

// ReceivingTeesAndKeys is a paid mutator transaction binding the contract method 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,uint64)[] _teeIdKeyIdPairs)
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) ReceivingTeesAndKeys(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ReceivingTeesAndKeys(&_TeeWalletKeyManager.TransactOpts, _walletId)
}

// ReceivingTeesAndKeys is a paid mutator transaction binding the contract method 0xfaddcc06.
//
// Solidity: function receivingTeesAndKeys(bytes32 _walletId) returns((address,uint64)[] _teeIdKeyIdPairs)
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) ReceivingTeesAndKeys(_walletId [32]byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.ReceivingTeesAndKeys(&_TeeWalletKeyManager.TransactOpts, _walletId)
}

// SetMultisigThreshold is a paid mutator transaction binding the contract method 0x62101d0a.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint64 _multisigThreshold) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) SetMultisigThreshold(opts *bind.TransactOpts, _walletId [32]byte, _multisigThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "setMultisigThreshold", _walletId, _multisigThreshold)
}

// SetMultisigThreshold is a paid mutator transaction binding the contract method 0x62101d0a.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint64 _multisigThreshold) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) SetMultisigThreshold(_walletId [32]byte, _multisigThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.SetMultisigThreshold(&_TeeWalletKeyManager.TransactOpts, _walletId, _multisigThreshold)
}

// SetMultisigThreshold is a paid mutator transaction binding the contract method 0x62101d0a.
//
// Solidity: function setMultisigThreshold(bytes32 _walletId, uint64 _multisigThreshold) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) SetMultisigThreshold(_walletId [32]byte, _multisigThreshold uint64) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.SetMultisigThreshold(&_TeeWalletKeyManager.TransactOpts, _walletId, _multisigThreshold)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.SwitchToProductionMode(&_TeeWalletKeyManager.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.SwitchToProductionMode(&_TeeWalletKeyManager.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.UpdateContractAddresses(&_TeeWalletKeyManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.UpdateContractAddresses(&_TeeWalletKeyManager.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.UpgradeToAndCall(&_TeeWalletKeyManager.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeeWalletKeyManager *TeeWalletKeyManagerTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeeWalletKeyManager.Contract.UpgradeToAndCall(&_TeeWalletKeyManager.TransactOpts, _newImplementation, _data)
}

// TeeWalletKeyManagerGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerGovernanceCallTimelockedIterator struct {
	Event *TeeWalletKeyManagerGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerGovernanceCallTimelocked)
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
		it.Event = new(TeeWalletKeyManagerGovernanceCallTimelocked)
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
func (it *TeeWalletKeyManagerGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeWalletKeyManagerGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerGovernanceCallTimelockedIterator{contract: _TeeWalletKeyManager.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerGovernanceCallTimelocked)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeWalletKeyManagerGovernanceCallTimelocked, error) {
	event := new(TeeWalletKeyManagerGovernanceCallTimelocked)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerGovernanceInitialisedIterator struct {
	Event *TeeWalletKeyManagerGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerGovernanceInitialised)
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
		it.Event = new(TeeWalletKeyManagerGovernanceInitialised)
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
func (it *TeeWalletKeyManagerGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerGovernanceInitialised represents a GovernanceInitialised event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeWalletKeyManagerGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerGovernanceInitialisedIterator{contract: _TeeWalletKeyManager.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerGovernanceInitialised)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseGovernanceInitialised(log types.Log) (*TeeWalletKeyManagerGovernanceInitialised, error) {
	event := new(TeeWalletKeyManagerGovernanceInitialised)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerGovernedProductionModeEnteredIterator struct {
	Event *TeeWalletKeyManagerGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerGovernedProductionModeEntered)
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
		it.Event = new(TeeWalletKeyManagerGovernedProductionModeEntered)
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
func (it *TeeWalletKeyManagerGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeeWalletKeyManagerGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerGovernedProductionModeEnteredIterator{contract: _TeeWalletKeyManager.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerGovernedProductionModeEntered)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeeWalletKeyManagerGovernedProductionModeEntered, error) {
	event := new(TeeWalletKeyManagerGovernedProductionModeEntered)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerTimelockedGovernanceCallCanceledIterator struct {
	Event *TeeWalletKeyManagerTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeeWalletKeyManagerTimelockedGovernanceCallCanceled)
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
func (it *TeeWalletKeyManagerTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeeWalletKeyManagerTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerTimelockedGovernanceCallCanceledIterator{contract: _TeeWalletKeyManager.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerTimelockedGovernanceCallCanceled)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeeWalletKeyManagerTimelockedGovernanceCallCanceled, error) {
	event := new(TeeWalletKeyManagerTimelockedGovernanceCallCanceled)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerTimelockedGovernanceCallExecutedIterator struct {
	Event *TeeWalletKeyManagerTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeeWalletKeyManagerTimelockedGovernanceCallExecuted)
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
func (it *TeeWalletKeyManagerTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeeWalletKeyManagerTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerTimelockedGovernanceCallExecutedIterator{contract: _TeeWalletKeyManager.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerTimelockedGovernanceCallExecuted)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeeWalletKeyManagerTimelockedGovernanceCallExecuted, error) {
	event := new(TeeWalletKeyManagerTimelockedGovernanceCallExecuted)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerUpgradedIterator struct {
	Event *TeeWalletKeyManagerUpgraded // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerUpgraded)
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
		it.Event = new(TeeWalletKeyManagerUpgraded)
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
func (it *TeeWalletKeyManagerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerUpgraded represents a Upgraded event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeeWalletKeyManagerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerUpgradedIterator{contract: _TeeWalletKeyManager.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerUpgraded)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseUpgraded(log types.Log) (*TeeWalletKeyManagerUpgraded, error) {
	event := new(TeeWalletKeyManagerUpgraded)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerWalletKeyAddedIterator is returned from FilterWalletKeyAdded and is used to iterate over the raw logs and unpacked data for WalletKeyAdded events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeyAddedIterator struct {
	Event *TeeWalletKeyManagerWalletKeyAdded // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerWalletKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerWalletKeyAdded)
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
		it.Event = new(TeeWalletKeyManagerWalletKeyAdded)
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
func (it *TeeWalletKeyManagerWalletKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerWalletKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerWalletKeyAdded represents a WalletKeyAdded event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeyAdded struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeyAdded is a free log retrieval operation binding the contract event 0x8798bb666f83ef50fc32aae926ecc61732c5f6f056210ade84ff0adeda273955.
//
// Solidity: event WalletKeyAdded(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletKeyAdded(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []uint64) (*TeeWalletKeyManagerWalletKeyAddedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "WalletKeyAdded", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerWalletKeyAddedIterator{contract: _TeeWalletKeyManager.contract, event: "WalletKeyAdded", logs: logs, sub: sub}, nil
}

// WatchWalletKeyAdded is a free log subscription operation binding the contract event 0x8798bb666f83ef50fc32aae926ecc61732c5f6f056210ade84ff0adeda273955.
//
// Solidity: event WalletKeyAdded(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletKeyAdded(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletKeyAdded, teeId []common.Address, walletId [][32]byte, keyId []uint64) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "WalletKeyAdded", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerWalletKeyAdded)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeyAdded", log); err != nil {
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

// ParseWalletKeyAdded is a log parse operation binding the contract event 0x8798bb666f83ef50fc32aae926ecc61732c5f6f056210ade84ff0adeda273955.
//
// Solidity: event WalletKeyAdded(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletKeyAdded(log types.Log) (*TeeWalletKeyManagerWalletKeyAdded, error) {
	event := new(TeeWalletKeyManagerWalletKeyAdded)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerWalletKeyConfirmedIterator is returned from FilterWalletKeyConfirmed and is used to iterate over the raw logs and unpacked data for WalletKeyConfirmed events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeyConfirmedIterator struct {
	Event *TeeWalletKeyManagerWalletKeyConfirmed // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerWalletKeyConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerWalletKeyConfirmed)
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
		it.Event = new(TeeWalletKeyManagerWalletKeyConfirmed)
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
func (it *TeeWalletKeyManagerWalletKeyConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerWalletKeyConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerWalletKeyConfirmed represents a WalletKeyConfirmed event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeyConfirmed struct {
	TeeId     common.Address
	WalletId  [32]byte
	KeyId     uint64
	PublicKey []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWalletKeyConfirmed is a free log retrieval operation binding the contract event 0x9eccb667b5a2a1ebae9bbef46a76d93aee881470e069aa53cdbcb7af7a2b3b21.
//
// Solidity: event WalletKeyConfirmed(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, bytes publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletKeyConfirmed(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []uint64) (*TeeWalletKeyManagerWalletKeyConfirmedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "WalletKeyConfirmed", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerWalletKeyConfirmedIterator{contract: _TeeWalletKeyManager.contract, event: "WalletKeyConfirmed", logs: logs, sub: sub}, nil
}

// WatchWalletKeyConfirmed is a free log subscription operation binding the contract event 0x9eccb667b5a2a1ebae9bbef46a76d93aee881470e069aa53cdbcb7af7a2b3b21.
//
// Solidity: event WalletKeyConfirmed(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, bytes publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletKeyConfirmed(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletKeyConfirmed, teeId []common.Address, walletId [][32]byte, keyId []uint64) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "WalletKeyConfirmed", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerWalletKeyConfirmed)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeyConfirmed", log); err != nil {
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

// ParseWalletKeyConfirmed is a log parse operation binding the contract event 0x9eccb667b5a2a1ebae9bbef46a76d93aee881470e069aa53cdbcb7af7a2b3b21.
//
// Solidity: event WalletKeyConfirmed(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId, bytes publicKey)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletKeyConfirmed(log types.Log) (*TeeWalletKeyManagerWalletKeyConfirmed, error) {
	event := new(TeeWalletKeyManagerWalletKeyConfirmed)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeyConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerWalletKeyDeletedIterator is returned from FilterWalletKeyDeleted and is used to iterate over the raw logs and unpacked data for WalletKeyDeleted events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeyDeletedIterator struct {
	Event *TeeWalletKeyManagerWalletKeyDeleted // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerWalletKeyDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerWalletKeyDeleted)
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
		it.Event = new(TeeWalletKeyManagerWalletKeyDeleted)
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
func (it *TeeWalletKeyManagerWalletKeyDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerWalletKeyDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerWalletKeyDeleted represents a WalletKeyDeleted event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeyDeleted struct {
	TeeId    common.Address
	WalletId [32]byte
	KeyId    uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeyDeleted is a free log retrieval operation binding the contract event 0x067c5c5fab799f3b1a41b54608f7f1fd5ec48e8575327ff3da9e4cc149f8226f.
//
// Solidity: event WalletKeyDeleted(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletKeyDeleted(opts *bind.FilterOpts, teeId []common.Address, walletId [][32]byte, keyId []uint64) (*TeeWalletKeyManagerWalletKeyDeletedIterator, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "WalletKeyDeleted", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerWalletKeyDeletedIterator{contract: _TeeWalletKeyManager.contract, event: "WalletKeyDeleted", logs: logs, sub: sub}, nil
}

// WatchWalletKeyDeleted is a free log subscription operation binding the contract event 0x067c5c5fab799f3b1a41b54608f7f1fd5ec48e8575327ff3da9e4cc149f8226f.
//
// Solidity: event WalletKeyDeleted(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletKeyDeleted(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletKeyDeleted, teeId []common.Address, walletId [][32]byte, keyId []uint64) (event.Subscription, error) {

	var teeIdRule []interface{}
	for _, teeIdItem := range teeId {
		teeIdRule = append(teeIdRule, teeIdItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var keyIdRule []interface{}
	for _, keyIdItem := range keyId {
		keyIdRule = append(keyIdRule, keyIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "WalletKeyDeleted", teeIdRule, walletIdRule, keyIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerWalletKeyDeleted)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeyDeleted", log); err != nil {
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

// ParseWalletKeyDeleted is a log parse operation binding the contract event 0x067c5c5fab799f3b1a41b54608f7f1fd5ec48e8575327ff3da9e4cc149f8226f.
//
// Solidity: event WalletKeyDeleted(address indexed teeId, bytes32 indexed walletId, uint64 indexed keyId)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletKeyDeleted(log types.Log) (*TeeWalletKeyManagerWalletKeyDeleted, error) {
	event := new(TeeWalletKeyManagerWalletKeyDeleted)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeyDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerWalletKeysNotAvailableIterator is returned from FilterWalletKeysNotAvailable and is used to iterate over the raw logs and unpacked data for WalletKeysNotAvailable events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeysNotAvailableIterator struct {
	Event *TeeWalletKeyManagerWalletKeysNotAvailable // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerWalletKeysNotAvailableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerWalletKeysNotAvailable)
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
		it.Event = new(TeeWalletKeyManagerWalletKeysNotAvailable)
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
func (it *TeeWalletKeyManagerWalletKeysNotAvailableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerWalletKeysNotAvailableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerWalletKeysNotAvailable represents a WalletKeysNotAvailable event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletKeysNotAvailable struct {
	WalletId [32]byte
	KeyIds   []uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWalletKeysNotAvailable is a free log retrieval operation binding the contract event 0xa82e9445a32f966f6330fc8dc7dbb1a5abf41ce17c2695f248a967d11dc8199c.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint64[] keyIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletKeysNotAvailable(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletKeyManagerWalletKeysNotAvailableIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "WalletKeysNotAvailable", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerWalletKeysNotAvailableIterator{contract: _TeeWalletKeyManager.contract, event: "WalletKeysNotAvailable", logs: logs, sub: sub}, nil
}

// WatchWalletKeysNotAvailable is a free log subscription operation binding the contract event 0xa82e9445a32f966f6330fc8dc7dbb1a5abf41ce17c2695f248a967d11dc8199c.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint64[] keyIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletKeysNotAvailable(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletKeysNotAvailable, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "WalletKeysNotAvailable", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerWalletKeysNotAvailable)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeysNotAvailable", log); err != nil {
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

// ParseWalletKeysNotAvailable is a log parse operation binding the contract event 0xa82e9445a32f966f6330fc8dc7dbb1a5abf41ce17c2695f248a967d11dc8199c.
//
// Solidity: event WalletKeysNotAvailable(bytes32 indexed walletId, uint64[] keyIds)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletKeysNotAvailable(log types.Log) (*TeeWalletKeyManagerWalletKeysNotAvailable, error) {
	event := new(TeeWalletKeyManagerWalletKeysNotAvailable)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletKeysNotAvailable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeWalletKeyManagerWalletMultisigThresholdSetIterator is returned from FilterWalletMultisigThresholdSet and is used to iterate over the raw logs and unpacked data for WalletMultisigThresholdSet events raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletMultisigThresholdSetIterator struct {
	Event *TeeWalletKeyManagerWalletMultisigThresholdSet // Event containing the contract specifics and raw log

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
func (it *TeeWalletKeyManagerWalletMultisigThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeWalletKeyManagerWalletMultisigThresholdSet)
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
		it.Event = new(TeeWalletKeyManagerWalletMultisigThresholdSet)
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
func (it *TeeWalletKeyManagerWalletMultisigThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeWalletKeyManagerWalletMultisigThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeWalletKeyManagerWalletMultisigThresholdSet represents a WalletMultisigThresholdSet event raised by the TeeWalletKeyManager contract.
type TeeWalletKeyManagerWalletMultisigThresholdSet struct {
	WalletId          [32]byte
	MultisigThreshold uint64
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWalletMultisigThresholdSet is a free log retrieval operation binding the contract event 0x3874a09a3d4be5b9ad2530ce7aeac3a302ade2cdc68a57336f7c59bad64cad0c.
//
// Solidity: event WalletMultisigThresholdSet(bytes32 indexed walletId, uint64 multisigThreshold)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) FilterWalletMultisigThresholdSet(opts *bind.FilterOpts, walletId [][32]byte) (*TeeWalletKeyManagerWalletMultisigThresholdSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.FilterLogs(opts, "WalletMultisigThresholdSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeWalletKeyManagerWalletMultisigThresholdSetIterator{contract: _TeeWalletKeyManager.contract, event: "WalletMultisigThresholdSet", logs: logs, sub: sub}, nil
}

// WatchWalletMultisigThresholdSet is a free log subscription operation binding the contract event 0x3874a09a3d4be5b9ad2530ce7aeac3a302ade2cdc68a57336f7c59bad64cad0c.
//
// Solidity: event WalletMultisigThresholdSet(bytes32 indexed walletId, uint64 multisigThreshold)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) WatchWalletMultisigThresholdSet(opts *bind.WatchOpts, sink chan<- *TeeWalletKeyManagerWalletMultisigThresholdSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _TeeWalletKeyManager.contract.WatchLogs(opts, "WalletMultisigThresholdSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeWalletKeyManagerWalletMultisigThresholdSet)
				if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletMultisigThresholdSet", log); err != nil {
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

// ParseWalletMultisigThresholdSet is a log parse operation binding the contract event 0x3874a09a3d4be5b9ad2530ce7aeac3a302ade2cdc68a57336f7c59bad64cad0c.
//
// Solidity: event WalletMultisigThresholdSet(bytes32 indexed walletId, uint64 multisigThreshold)
func (_TeeWalletKeyManager *TeeWalletKeyManagerFilterer) ParseWalletMultisigThresholdSet(log types.Log) (*TeeWalletKeyManagerWalletMultisigThresholdSet, error) {
	event := new(TeeWalletKeyManagerWalletMultisigThresholdSet)
	if err := _TeeWalletKeyManager.contract.UnpackLog(event, "WalletMultisigThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
