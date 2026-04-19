// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teeinstructions

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

// IInstructionsFacetTeeInstructionParams is an auto generated low-level Go binding around an user-defined struct.
type IInstructionsFacetTeeInstructionParams struct {
	OpType             [32]byte
	OpCommand          [32]byte
	Message            []byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	ClaimBackAddress   common.Address
}

// IMachineManagerFacetTeeMachine is an auto generated low-level Go binding around an user-defined struct.
type IMachineManagerFacetTeeMachine struct {
	TeeId      common.Address
	TeeProxyId common.Address
	Url        string
}

// TeeInstructionsMetaData contains all meta data concerning the TeeInstructions contract.
var TeeInstructionsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AvailabilityCheckTimestampInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"DuplicatedCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExtensionIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAvailabilityCheckStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGovernanceHash\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPublicKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidResponseData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningAlgo\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidThreshold\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWalletStatus\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"}],\"name\":\"KeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MessageEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoTeeMachinesSpecified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExtensionOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyOwnerOrBackupManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProposedOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemInstructionsSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationCommandEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationTypeEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OwnerNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"instructionsSender\",\"type\":\"address\"}],\"name\":\"SystemInstructionsSenderAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"instructionsSender\",\"type\":\"address\"}],\"name\":\"SystemInstructionsSenderNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"}],\"name\":\"SystemOpTypeNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeMachineNotAvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TeeNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionNotSupported\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"instructionsSenders\",\"type\":\"address[]\"}],\"name\":\"SystemInstructionsSendersRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"instructionsSenders\",\"type\":\"address[]\"}],\"name\":\"SystemInstructionsSendersUnregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"extensionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"rewardEpochId\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structIMachineManagerFacet.TeeMachine[]\",\"name\":\"teeMachines\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TeeInstructionsSent\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getSystemInstructionsSenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_instructionsSenders\",\"type\":\"address[]\"}],\"name\":\"registerSystemInstructionsSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"}],\"internalType\":\"structIInstructionsFacet.TeeInstructionParams\",\"name\":\"_instructionParams\",\"type\":\"tuple\"}],\"name\":\"sendInstructions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"}],\"internalType\":\"structIInstructionsFacet.TeeInstructionParams\",\"name\":\"_instructionParams\",\"type\":\"tuple\"}],\"name\":\"sendSystemInstructions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_instructionId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"teeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"teeProxyId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"internalType\":\"structIMachineManagerFacet.TeeMachine[]\",\"name\":\"_teeMachines\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"}],\"internalType\":\"structIInstructionsFacet.TeeInstructionParams\",\"name\":\"_instructionParams\",\"type\":\"tuple\"}],\"name\":\"sendSystemInstructions\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_instructionsSenders\",\"type\":\"address[]\"}],\"name\":\"unregisterSystemInstructionsSenders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeeInstructionsABI is the input ABI used to generate the binding from.
// Deprecated: Use TeeInstructionsMetaData.ABI instead.
var TeeInstructionsABI = TeeInstructionsMetaData.ABI

// TeeInstructions is an auto generated Go binding around an Ethereum contract.
type TeeInstructions struct {
	TeeInstructionsCaller     // Read-only binding to the contract
	TeeInstructionsTransactor // Write-only binding to the contract
	TeeInstructionsFilterer   // Log filterer for contract events
}

// TeeInstructionsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeeInstructionsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeInstructionsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeeInstructionsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeInstructionsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeeInstructionsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeeInstructionsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeeInstructionsSession struct {
	Contract     *TeeInstructions  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeeInstructionsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeeInstructionsCallerSession struct {
	Contract *TeeInstructionsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TeeInstructionsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeeInstructionsTransactorSession struct {
	Contract     *TeeInstructionsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TeeInstructionsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeeInstructionsRaw struct {
	Contract *TeeInstructions // Generic contract binding to access the raw methods on
}

// TeeInstructionsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeeInstructionsCallerRaw struct {
	Contract *TeeInstructionsCaller // Generic read-only contract binding to access the raw methods on
}

// TeeInstructionsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeeInstructionsTransactorRaw struct {
	Contract *TeeInstructionsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeeInstructions creates a new instance of TeeInstructions, bound to a specific deployed contract.
func NewTeeInstructions(address common.Address, backend bind.ContractBackend) (*TeeInstructions, error) {
	contract, err := bindTeeInstructions(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeeInstructions{TeeInstructionsCaller: TeeInstructionsCaller{contract: contract}, TeeInstructionsTransactor: TeeInstructionsTransactor{contract: contract}, TeeInstructionsFilterer: TeeInstructionsFilterer{contract: contract}}, nil
}

// NewTeeInstructionsCaller creates a new read-only instance of TeeInstructions, bound to a specific deployed contract.
func NewTeeInstructionsCaller(address common.Address, caller bind.ContractCaller) (*TeeInstructionsCaller, error) {
	contract, err := bindTeeInstructions(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsCaller{contract: contract}, nil
}

// NewTeeInstructionsTransactor creates a new write-only instance of TeeInstructions, bound to a specific deployed contract.
func NewTeeInstructionsTransactor(address common.Address, transactor bind.ContractTransactor) (*TeeInstructionsTransactor, error) {
	contract, err := bindTeeInstructions(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsTransactor{contract: contract}, nil
}

// NewTeeInstructionsFilterer creates a new log filterer instance of TeeInstructions, bound to a specific deployed contract.
func NewTeeInstructionsFilterer(address common.Address, filterer bind.ContractFilterer) (*TeeInstructionsFilterer, error) {
	contract, err := bindTeeInstructions(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsFilterer{contract: contract}, nil
}

// bindTeeInstructions binds a generic wrapper to an already deployed contract.
func bindTeeInstructions(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeeInstructionsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeInstructions *TeeInstructionsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeInstructions.Contract.TeeInstructionsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeInstructions *TeeInstructionsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeInstructions.Contract.TeeInstructionsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeInstructions *TeeInstructionsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeInstructions.Contract.TeeInstructionsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeeInstructions *TeeInstructionsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeeInstructions.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeeInstructions *TeeInstructionsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeeInstructions.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeeInstructions *TeeInstructionsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeeInstructions.Contract.contract.Transact(opts, method, params...)
}

// GetSystemInstructionsSenders is a free data retrieval call binding the contract method 0xa37768e5.
//
// Solidity: function getSystemInstructionsSenders() view returns(address[])
func (_TeeInstructions *TeeInstructionsCaller) GetSystemInstructionsSenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _TeeInstructions.contract.Call(opts, &out, "getSystemInstructionsSenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSystemInstructionsSenders is a free data retrieval call binding the contract method 0xa37768e5.
//
// Solidity: function getSystemInstructionsSenders() view returns(address[])
func (_TeeInstructions *TeeInstructionsSession) GetSystemInstructionsSenders() ([]common.Address, error) {
	return _TeeInstructions.Contract.GetSystemInstructionsSenders(&_TeeInstructions.CallOpts)
}

// GetSystemInstructionsSenders is a free data retrieval call binding the contract method 0xa37768e5.
//
// Solidity: function getSystemInstructionsSenders() view returns(address[])
func (_TeeInstructions *TeeInstructionsCallerSession) GetSystemInstructionsSenders() ([]common.Address, error) {
	return _TeeInstructions.Contract.GetSystemInstructionsSenders(&_TeeInstructions.CallOpts)
}

// RegisterSystemInstructionsSenders is a paid mutator transaction binding the contract method 0x4d9003cf.
//
// Solidity: function registerSystemInstructionsSenders(address[] _instructionsSenders) returns()
func (_TeeInstructions *TeeInstructionsTransactor) RegisterSystemInstructionsSenders(opts *bind.TransactOpts, _instructionsSenders []common.Address) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "registerSystemInstructionsSenders", _instructionsSenders)
}

// RegisterSystemInstructionsSenders is a paid mutator transaction binding the contract method 0x4d9003cf.
//
// Solidity: function registerSystemInstructionsSenders(address[] _instructionsSenders) returns()
func (_TeeInstructions *TeeInstructionsSession) RegisterSystemInstructionsSenders(_instructionsSenders []common.Address) (*types.Transaction, error) {
	return _TeeInstructions.Contract.RegisterSystemInstructionsSenders(&_TeeInstructions.TransactOpts, _instructionsSenders)
}

// RegisterSystemInstructionsSenders is a paid mutator transaction binding the contract method 0x4d9003cf.
//
// Solidity: function registerSystemInstructionsSenders(address[] _instructionsSenders) returns()
func (_TeeInstructions *TeeInstructionsTransactorSession) RegisterSystemInstructionsSenders(_instructionsSenders []common.Address) (*types.Transaction, error) {
	return _TeeInstructions.Contract.RegisterSystemInstructionsSenders(&_TeeInstructions.TransactOpts, _instructionsSenders)
}

// SendInstructions is a paid mutator transaction binding the contract method 0xf731df53.
//
// Solidity: function sendInstructions(address[] _teeIds, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (_TeeInstructions *TeeInstructionsTransactor) SendInstructions(opts *bind.TransactOpts, _teeIds []common.Address, _instructionParams IInstructionsFacetTeeInstructionParams) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "sendInstructions", _teeIds, _instructionParams)
}

// SendInstructions is a paid mutator transaction binding the contract method 0xf731df53.
//
// Solidity: function sendInstructions(address[] _teeIds, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (_TeeInstructions *TeeInstructionsSession) SendInstructions(_teeIds []common.Address, _instructionParams IInstructionsFacetTeeInstructionParams) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SendInstructions(&_TeeInstructions.TransactOpts, _teeIds, _instructionParams)
}

// SendInstructions is a paid mutator transaction binding the contract method 0xf731df53.
//
// Solidity: function sendInstructions(address[] _teeIds, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (_TeeInstructions *TeeInstructionsTransactorSession) SendInstructions(_teeIds []common.Address, _instructionParams IInstructionsFacetTeeInstructionParams) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SendInstructions(&_TeeInstructions.TransactOpts, _teeIds, _instructionParams)
}

// SendSystemInstructions is a paid mutator transaction binding the contract method 0x1456290e.
//
// Solidity: function sendSystemInstructions(bytes32 _instructionId, address[] _teeIds, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (_TeeInstructions *TeeInstructionsTransactor) SendSystemInstructions(opts *bind.TransactOpts, _instructionId [32]byte, _teeIds []common.Address, _instructionParams IInstructionsFacetTeeInstructionParams) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "sendSystemInstructions", _instructionId, _teeIds, _instructionParams)
}

// SendSystemInstructions is a paid mutator transaction binding the contract method 0x1456290e.
//
// Solidity: function sendSystemInstructions(bytes32 _instructionId, address[] _teeIds, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (_TeeInstructions *TeeInstructionsSession) SendSystemInstructions(_instructionId [32]byte, _teeIds []common.Address, _instructionParams IInstructionsFacetTeeInstructionParams) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SendSystemInstructions(&_TeeInstructions.TransactOpts, _instructionId, _teeIds, _instructionParams)
}

// SendSystemInstructions is a paid mutator transaction binding the contract method 0x1456290e.
//
// Solidity: function sendSystemInstructions(bytes32 _instructionId, address[] _teeIds, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (_TeeInstructions *TeeInstructionsTransactorSession) SendSystemInstructions(_instructionId [32]byte, _teeIds []common.Address, _instructionParams IInstructionsFacetTeeInstructionParams) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SendSystemInstructions(&_TeeInstructions.TransactOpts, _instructionId, _teeIds, _instructionParams)
}

// SendSystemInstructions0 is a paid mutator transaction binding the contract method 0x32103728.
//
// Solidity: function sendSystemInstructions(bytes32 _instructionId, (address,address,string)[] _teeMachines, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (_TeeInstructions *TeeInstructionsTransactor) SendSystemInstructions0(opts *bind.TransactOpts, _instructionId [32]byte, _teeMachines []IMachineManagerFacetTeeMachine, _instructionParams IInstructionsFacetTeeInstructionParams) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "sendSystemInstructions0", _instructionId, _teeMachines, _instructionParams)
}

// SendSystemInstructions0 is a paid mutator transaction binding the contract method 0x32103728.
//
// Solidity: function sendSystemInstructions(bytes32 _instructionId, (address,address,string)[] _teeMachines, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (_TeeInstructions *TeeInstructionsSession) SendSystemInstructions0(_instructionId [32]byte, _teeMachines []IMachineManagerFacetTeeMachine, _instructionParams IInstructionsFacetTeeInstructionParams) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SendSystemInstructions0(&_TeeInstructions.TransactOpts, _instructionId, _teeMachines, _instructionParams)
}

// SendSystemInstructions0 is a paid mutator transaction binding the contract method 0x32103728.
//
// Solidity: function sendSystemInstructions(bytes32 _instructionId, (address,address,string)[] _teeMachines, (bytes32,bytes32,bytes,address[],uint64,address) _instructionParams) payable returns(bytes32)
func (_TeeInstructions *TeeInstructionsTransactorSession) SendSystemInstructions0(_instructionId [32]byte, _teeMachines []IMachineManagerFacetTeeMachine, _instructionParams IInstructionsFacetTeeInstructionParams) (*types.Transaction, error) {
	return _TeeInstructions.Contract.SendSystemInstructions0(&_TeeInstructions.TransactOpts, _instructionId, _teeMachines, _instructionParams)
}

// UnregisterSystemInstructionsSenders is a paid mutator transaction binding the contract method 0xf123abf3.
//
// Solidity: function unregisterSystemInstructionsSenders(address[] _instructionsSenders) returns()
func (_TeeInstructions *TeeInstructionsTransactor) UnregisterSystemInstructionsSenders(opts *bind.TransactOpts, _instructionsSenders []common.Address) (*types.Transaction, error) {
	return _TeeInstructions.contract.Transact(opts, "unregisterSystemInstructionsSenders", _instructionsSenders)
}

// UnregisterSystemInstructionsSenders is a paid mutator transaction binding the contract method 0xf123abf3.
//
// Solidity: function unregisterSystemInstructionsSenders(address[] _instructionsSenders) returns()
func (_TeeInstructions *TeeInstructionsSession) UnregisterSystemInstructionsSenders(_instructionsSenders []common.Address) (*types.Transaction, error) {
	return _TeeInstructions.Contract.UnregisterSystemInstructionsSenders(&_TeeInstructions.TransactOpts, _instructionsSenders)
}

// UnregisterSystemInstructionsSenders is a paid mutator transaction binding the contract method 0xf123abf3.
//
// Solidity: function unregisterSystemInstructionsSenders(address[] _instructionsSenders) returns()
func (_TeeInstructions *TeeInstructionsTransactorSession) UnregisterSystemInstructionsSenders(_instructionsSenders []common.Address) (*types.Transaction, error) {
	return _TeeInstructions.Contract.UnregisterSystemInstructionsSenders(&_TeeInstructions.TransactOpts, _instructionsSenders)
}

// TeeInstructionsGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeeInstructions contract.
type TeeInstructionsGovernanceCallTimelockedIterator struct {
	Event *TeeInstructionsGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeeInstructionsGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeInstructionsGovernanceCallTimelocked)
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
		it.Event = new(TeeInstructionsGovernanceCallTimelocked)
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
func (it *TeeInstructionsGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeInstructionsGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeInstructionsGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeeInstructions contract.
type TeeInstructionsGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeeInstructions *TeeInstructionsFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeeInstructionsGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeeInstructions.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsGovernanceCallTimelockedIterator{contract: _TeeInstructions.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeeInstructions *TeeInstructionsFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeeInstructionsGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeeInstructions.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeInstructionsGovernanceCallTimelocked)
				if err := _TeeInstructions.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeeInstructions *TeeInstructionsFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeeInstructionsGovernanceCallTimelocked, error) {
	event := new(TeeInstructionsGovernanceCallTimelocked)
	if err := _TeeInstructions.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeInstructionsGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeeInstructions contract.
type TeeInstructionsGovernanceInitialisedIterator struct {
	Event *TeeInstructionsGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeeInstructionsGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeInstructionsGovernanceInitialised)
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
		it.Event = new(TeeInstructionsGovernanceInitialised)
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
func (it *TeeInstructionsGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeInstructionsGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeInstructionsGovernanceInitialised represents a GovernanceInitialised event raised by the TeeInstructions contract.
type TeeInstructionsGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeInstructions *TeeInstructionsFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeeInstructionsGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeeInstructions.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsGovernanceInitialisedIterator{contract: _TeeInstructions.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeeInstructions *TeeInstructionsFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeeInstructionsGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeeInstructions.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeInstructionsGovernanceInitialised)
				if err := _TeeInstructions.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeeInstructions *TeeInstructionsFilterer) ParseGovernanceInitialised(log types.Log) (*TeeInstructionsGovernanceInitialised, error) {
	event := new(TeeInstructionsGovernanceInitialised)
	if err := _TeeInstructions.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeInstructionsSystemInstructionsSendersRegisteredIterator is returned from FilterSystemInstructionsSendersRegistered and is used to iterate over the raw logs and unpacked data for SystemInstructionsSendersRegistered events raised by the TeeInstructions contract.
type TeeInstructionsSystemInstructionsSendersRegisteredIterator struct {
	Event *TeeInstructionsSystemInstructionsSendersRegistered // Event containing the contract specifics and raw log

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
func (it *TeeInstructionsSystemInstructionsSendersRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeInstructionsSystemInstructionsSendersRegistered)
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
		it.Event = new(TeeInstructionsSystemInstructionsSendersRegistered)
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
func (it *TeeInstructionsSystemInstructionsSendersRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeInstructionsSystemInstructionsSendersRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeInstructionsSystemInstructionsSendersRegistered represents a SystemInstructionsSendersRegistered event raised by the TeeInstructions contract.
type TeeInstructionsSystemInstructionsSendersRegistered struct {
	InstructionsSenders []common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSystemInstructionsSendersRegistered is a free log retrieval operation binding the contract event 0x11d02ff409d2d9cc57a3b2a97a6ac66324adfbb582cabaa3bcbe740724ddae68.
//
// Solidity: event SystemInstructionsSendersRegistered(address[] instructionsSenders)
func (_TeeInstructions *TeeInstructionsFilterer) FilterSystemInstructionsSendersRegistered(opts *bind.FilterOpts) (*TeeInstructionsSystemInstructionsSendersRegisteredIterator, error) {

	logs, sub, err := _TeeInstructions.contract.FilterLogs(opts, "SystemInstructionsSendersRegistered")
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsSystemInstructionsSendersRegisteredIterator{contract: _TeeInstructions.contract, event: "SystemInstructionsSendersRegistered", logs: logs, sub: sub}, nil
}

// WatchSystemInstructionsSendersRegistered is a free log subscription operation binding the contract event 0x11d02ff409d2d9cc57a3b2a97a6ac66324adfbb582cabaa3bcbe740724ddae68.
//
// Solidity: event SystemInstructionsSendersRegistered(address[] instructionsSenders)
func (_TeeInstructions *TeeInstructionsFilterer) WatchSystemInstructionsSendersRegistered(opts *bind.WatchOpts, sink chan<- *TeeInstructionsSystemInstructionsSendersRegistered) (event.Subscription, error) {

	logs, sub, err := _TeeInstructions.contract.WatchLogs(opts, "SystemInstructionsSendersRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeInstructionsSystemInstructionsSendersRegistered)
				if err := _TeeInstructions.contract.UnpackLog(event, "SystemInstructionsSendersRegistered", log); err != nil {
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

// ParseSystemInstructionsSendersRegistered is a log parse operation binding the contract event 0x11d02ff409d2d9cc57a3b2a97a6ac66324adfbb582cabaa3bcbe740724ddae68.
//
// Solidity: event SystemInstructionsSendersRegistered(address[] instructionsSenders)
func (_TeeInstructions *TeeInstructionsFilterer) ParseSystemInstructionsSendersRegistered(log types.Log) (*TeeInstructionsSystemInstructionsSendersRegistered, error) {
	event := new(TeeInstructionsSystemInstructionsSendersRegistered)
	if err := _TeeInstructions.contract.UnpackLog(event, "SystemInstructionsSendersRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeInstructionsSystemInstructionsSendersUnregisteredIterator is returned from FilterSystemInstructionsSendersUnregistered and is used to iterate over the raw logs and unpacked data for SystemInstructionsSendersUnregistered events raised by the TeeInstructions contract.
type TeeInstructionsSystemInstructionsSendersUnregisteredIterator struct {
	Event *TeeInstructionsSystemInstructionsSendersUnregistered // Event containing the contract specifics and raw log

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
func (it *TeeInstructionsSystemInstructionsSendersUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeInstructionsSystemInstructionsSendersUnregistered)
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
		it.Event = new(TeeInstructionsSystemInstructionsSendersUnregistered)
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
func (it *TeeInstructionsSystemInstructionsSendersUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeInstructionsSystemInstructionsSendersUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeInstructionsSystemInstructionsSendersUnregistered represents a SystemInstructionsSendersUnregistered event raised by the TeeInstructions contract.
type TeeInstructionsSystemInstructionsSendersUnregistered struct {
	InstructionsSenders []common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSystemInstructionsSendersUnregistered is a free log retrieval operation binding the contract event 0x63220dc7b394b4e42ed020c0289cb0b17d72c91f6e04221035d71da3d93eba99.
//
// Solidity: event SystemInstructionsSendersUnregistered(address[] instructionsSenders)
func (_TeeInstructions *TeeInstructionsFilterer) FilterSystemInstructionsSendersUnregistered(opts *bind.FilterOpts) (*TeeInstructionsSystemInstructionsSendersUnregisteredIterator, error) {

	logs, sub, err := _TeeInstructions.contract.FilterLogs(opts, "SystemInstructionsSendersUnregistered")
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsSystemInstructionsSendersUnregisteredIterator{contract: _TeeInstructions.contract, event: "SystemInstructionsSendersUnregistered", logs: logs, sub: sub}, nil
}

// WatchSystemInstructionsSendersUnregistered is a free log subscription operation binding the contract event 0x63220dc7b394b4e42ed020c0289cb0b17d72c91f6e04221035d71da3d93eba99.
//
// Solidity: event SystemInstructionsSendersUnregistered(address[] instructionsSenders)
func (_TeeInstructions *TeeInstructionsFilterer) WatchSystemInstructionsSendersUnregistered(opts *bind.WatchOpts, sink chan<- *TeeInstructionsSystemInstructionsSendersUnregistered) (event.Subscription, error) {

	logs, sub, err := _TeeInstructions.contract.WatchLogs(opts, "SystemInstructionsSendersUnregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeInstructionsSystemInstructionsSendersUnregistered)
				if err := _TeeInstructions.contract.UnpackLog(event, "SystemInstructionsSendersUnregistered", log); err != nil {
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

// ParseSystemInstructionsSendersUnregistered is a log parse operation binding the contract event 0x63220dc7b394b4e42ed020c0289cb0b17d72c91f6e04221035d71da3d93eba99.
//
// Solidity: event SystemInstructionsSendersUnregistered(address[] instructionsSenders)
func (_TeeInstructions *TeeInstructionsFilterer) ParseSystemInstructionsSendersUnregistered(log types.Log) (*TeeInstructionsSystemInstructionsSendersUnregistered, error) {
	event := new(TeeInstructionsSystemInstructionsSendersUnregistered)
	if err := _TeeInstructions.contract.UnpackLog(event, "SystemInstructionsSendersUnregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeeInstructionsTeeInstructionsSentIterator is returned from FilterTeeInstructionsSent and is used to iterate over the raw logs and unpacked data for TeeInstructionsSent events raised by the TeeInstructions contract.
type TeeInstructionsTeeInstructionsSentIterator struct {
	Event *TeeInstructionsTeeInstructionsSent // Event containing the contract specifics and raw log

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
func (it *TeeInstructionsTeeInstructionsSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeeInstructionsTeeInstructionsSent)
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
		it.Event = new(TeeInstructionsTeeInstructionsSent)
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
func (it *TeeInstructionsTeeInstructionsSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeeInstructionsTeeInstructionsSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeeInstructionsTeeInstructionsSent represents a TeeInstructionsSent event raised by the TeeInstructions contract.
type TeeInstructionsTeeInstructionsSent struct {
	ExtensionId        *big.Int
	InstructionId      [32]byte
	RewardEpochId      uint32
	TeeMachines        []IMachineManagerFacetTeeMachine
	OpType             [32]byte
	OpCommand          [32]byte
	Message            []byte
	Cosigners          []common.Address
	CosignersThreshold uint64
	ClaimBackAddress   common.Address
	Fee                *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterTeeInstructionsSent is a free log retrieval operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeInstructions *TeeInstructionsFilterer) FilterTeeInstructionsSent(opts *bind.FilterOpts, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (*TeeInstructionsTeeInstructionsSentIterator, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeInstructions.contract.FilterLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &TeeInstructionsTeeInstructionsSentIterator{contract: _TeeInstructions.contract, event: "TeeInstructionsSent", logs: logs, sub: sub}, nil
}

// WatchTeeInstructionsSent is a free log subscription operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeInstructions *TeeInstructionsFilterer) WatchTeeInstructionsSent(opts *bind.WatchOpts, sink chan<- *TeeInstructionsTeeInstructionsSent, extensionId []*big.Int, instructionId [][32]byte, rewardEpochId []uint32) (event.Subscription, error) {

	var extensionIdRule []interface{}
	for _, extensionIdItem := range extensionId {
		extensionIdRule = append(extensionIdRule, extensionIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var rewardEpochIdRule []interface{}
	for _, rewardEpochIdItem := range rewardEpochId {
		rewardEpochIdRule = append(rewardEpochIdRule, rewardEpochIdItem)
	}

	logs, sub, err := _TeeInstructions.contract.WatchLogs(opts, "TeeInstructionsSent", extensionIdRule, instructionIdRule, rewardEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeeInstructionsTeeInstructionsSent)
				if err := _TeeInstructions.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
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

// ParseTeeInstructionsSent is a log parse operation binding the contract event 0xf770e69a9fc05b7180797556ec4cedb6108ce2c56ffa76c84aa087efeb5e6963.
//
// Solidity: event TeeInstructionsSent(uint256 indexed extensionId, bytes32 indexed instructionId, uint32 indexed rewardEpochId, (address,address,string)[] teeMachines, bytes32 opType, bytes32 opCommand, bytes message, address[] cosigners, uint64 cosignersThreshold, address claimBackAddress, uint256 fee)
func (_TeeInstructions *TeeInstructionsFilterer) ParseTeeInstructionsSent(log types.Log) (*TeeInstructionsTeeInstructionsSent, error) {
	event := new(TeeInstructionsTeeInstructionsSent)
	if err := _TeeInstructions.contract.UnpackLog(event, "TeeInstructionsSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
