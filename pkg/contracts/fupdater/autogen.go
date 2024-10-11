// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fupdater

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

// Bn256G1Point is an auto generated low-level Go binding around an user-defined struct.
type Bn256G1Point struct {
	X *big.Int
	Y *big.Int
}

// IFastUpdaterFastUpdates is an auto generated low-level Go binding around an user-defined struct.
type IFastUpdaterFastUpdates struct {
	SortitionBlock      *big.Int
	SortitionCredential SortitionCredential
	Deltas              []byte
	Signature           IFastUpdaterSignature
}

// IFastUpdaterSignature is an auto generated low-level Go binding around an user-defined struct.
type IFastUpdaterSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// SortitionCredential is an auto generated low-level Go binding around an user-defined struct.
type SortitionCredential struct {
	Replicate *big.Int
	Gamma     Bn256G1Point
	C         *big.Int
	S         *big.Int
}

// FUpdaterMetaData contains all meta data concerning the FUpdater contract.
var FUpdaterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_flareDaemon\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_firstVotingRoundStartTs\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_votingEpochDurationSeconds\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_submissionWindow\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"FastUpdateFeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"votingRoundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes21\",\"name\":\"id\",\"type\":\"bytes21\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int8\",\"name\":\"decimals\",\"type\":\"int8\"}],\"name\":\"FastUpdateFeedReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"votingEpochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"feeds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"int8[]\",\"name\":\"decimals\",\"type\":\"int8[]\"}],\"name\":\"FastUpdateFeeds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"votingRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"FastUpdateFeedsSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BLOCKS_HISTORY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEED_AGE_IN_VOTING_EPOCHS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"}],\"name\":\"blockScoreCutoff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_cutoff\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewardEpochId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentScoreCutoff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_cutoff\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"currentSortitionWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daemonize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fastUpdateIncentiveManager\",\"outputs\":[{\"internalType\":\"contractIIFastUpdateIncentiveManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fastUpdatesConfiguration\",\"outputs\":[{\"internalType\":\"contractIFastUpdatesConfiguration\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchAllCurrentFeeds\",\"outputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_feedIds\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_feeds\",\"type\":\"uint256[]\"},{\"internalType\":\"int8[]\",\"name\":\"_decimals\",\"type\":\"int8[]\"},{\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indices\",\"type\":\"uint256[]\"}],\"name\":\"fetchCurrentFeeds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_feeds\",\"type\":\"uint256[]\"},{\"internalType\":\"int8[]\",\"name\":\"_decimals\",\"type\":\"int8[]\"},{\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstVotingRoundStartTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareDaemon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftsoFeedPublisher\",\"outputs\":[{\"internalType\":\"contractIFtsoFeedPublisher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_historySize\",\"type\":\"uint256\"}],\"name\":\"numberOfUpdates\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_noOfUpdates\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberOfUpdatesInBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indices\",\"type\":\"uint256[]\"}],\"name\":\"removeFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indices\",\"type\":\"uint256[]\"}],\"name\":\"resetFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_submissionWindow\",\"type\":\"uint8\"}],\"name\":\"setSubmissionWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submissionWindow\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sortitionBlock\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"replicate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structBn256.G1Point\",\"name\":\"gamma\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structSortitionCredential\",\"name\":\"sortitionCredential\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"deltas\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFastUpdater.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structIFastUpdater.FastUpdates\",\"name\":\"_updates\",\"type\":\"tuple\"}],\"name\":\"submitUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToFallbackMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_part1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_part2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_verificationData\",\"type\":\"bytes\"}],\"name\":\"verifyPublicKey\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistry\",\"outputs\":[{\"internalType\":\"contractIIVoterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FUpdaterABI is the input ABI used to generate the binding from.
// Deprecated: Use FUpdaterMetaData.ABI instead.
var FUpdaterABI = FUpdaterMetaData.ABI

// FUpdater is an auto generated Go binding around an Ethereum contract.
type FUpdater struct {
	FUpdaterCaller     // Read-only binding to the contract
	FUpdaterTransactor // Write-only binding to the contract
	FUpdaterFilterer   // Log filterer for contract events
}

// FUpdaterCaller is an auto generated read-only Go binding around an Ethereum contract.
type FUpdaterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FUpdaterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FUpdaterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FUpdaterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FUpdaterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FUpdaterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FUpdaterSession struct {
	Contract     *FUpdater         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FUpdaterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FUpdaterCallerSession struct {
	Contract *FUpdaterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// FUpdaterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FUpdaterTransactorSession struct {
	Contract     *FUpdaterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// FUpdaterRaw is an auto generated low-level Go binding around an Ethereum contract.
type FUpdaterRaw struct {
	Contract *FUpdater // Generic contract binding to access the raw methods on
}

// FUpdaterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FUpdaterCallerRaw struct {
	Contract *FUpdaterCaller // Generic read-only contract binding to access the raw methods on
}

// FUpdaterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FUpdaterTransactorRaw struct {
	Contract *FUpdaterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFUpdater creates a new instance of FUpdater, bound to a specific deployed contract.
func NewFUpdater(address common.Address, backend bind.ContractBackend) (*FUpdater, error) {
	contract, err := bindFUpdater(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FUpdater{FUpdaterCaller: FUpdaterCaller{contract: contract}, FUpdaterTransactor: FUpdaterTransactor{contract: contract}, FUpdaterFilterer: FUpdaterFilterer{contract: contract}}, nil
}

// NewFUpdaterCaller creates a new read-only instance of FUpdater, bound to a specific deployed contract.
func NewFUpdaterCaller(address common.Address, caller bind.ContractCaller) (*FUpdaterCaller, error) {
	contract, err := bindFUpdater(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FUpdaterCaller{contract: contract}, nil
}

// NewFUpdaterTransactor creates a new write-only instance of FUpdater, bound to a specific deployed contract.
func NewFUpdaterTransactor(address common.Address, transactor bind.ContractTransactor) (*FUpdaterTransactor, error) {
	contract, err := bindFUpdater(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FUpdaterTransactor{contract: contract}, nil
}

// NewFUpdaterFilterer creates a new log filterer instance of FUpdater, bound to a specific deployed contract.
func NewFUpdaterFilterer(address common.Address, filterer bind.ContractFilterer) (*FUpdaterFilterer, error) {
	contract, err := bindFUpdater(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FUpdaterFilterer{contract: contract}, nil
}

// bindFUpdater binds a generic wrapper to an already deployed contract.
func bindFUpdater(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FUpdaterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FUpdater *FUpdaterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FUpdater.Contract.FUpdaterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FUpdater *FUpdaterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FUpdater.Contract.FUpdaterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FUpdater *FUpdaterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FUpdater.Contract.FUpdaterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FUpdater *FUpdaterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FUpdater.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FUpdater *FUpdaterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FUpdater.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FUpdater *FUpdaterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FUpdater.Contract.contract.Transact(opts, method, params...)
}

// MAXBLOCKSHISTORY is a free data retrieval call binding the contract method 0xc1bff139.
//
// Solidity: function MAX_BLOCKS_HISTORY() view returns(uint256)
func (_FUpdater *FUpdaterCaller) MAXBLOCKSHISTORY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "MAX_BLOCKS_HISTORY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBLOCKSHISTORY is a free data retrieval call binding the contract method 0xc1bff139.
//
// Solidity: function MAX_BLOCKS_HISTORY() view returns(uint256)
func (_FUpdater *FUpdaterSession) MAXBLOCKSHISTORY() (*big.Int, error) {
	return _FUpdater.Contract.MAXBLOCKSHISTORY(&_FUpdater.CallOpts)
}

// MAXBLOCKSHISTORY is a free data retrieval call binding the contract method 0xc1bff139.
//
// Solidity: function MAX_BLOCKS_HISTORY() view returns(uint256)
func (_FUpdater *FUpdaterCallerSession) MAXBLOCKSHISTORY() (*big.Int, error) {
	return _FUpdater.Contract.MAXBLOCKSHISTORY(&_FUpdater.CallOpts)
}

// MAXFEEDAGEINVOTINGEPOCHS is a free data retrieval call binding the contract method 0x7fe3341a.
//
// Solidity: function MAX_FEED_AGE_IN_VOTING_EPOCHS() view returns(uint256)
func (_FUpdater *FUpdaterCaller) MAXFEEDAGEINVOTINGEPOCHS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "MAX_FEED_AGE_IN_VOTING_EPOCHS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEEDAGEINVOTINGEPOCHS is a free data retrieval call binding the contract method 0x7fe3341a.
//
// Solidity: function MAX_FEED_AGE_IN_VOTING_EPOCHS() view returns(uint256)
func (_FUpdater *FUpdaterSession) MAXFEEDAGEINVOTINGEPOCHS() (*big.Int, error) {
	return _FUpdater.Contract.MAXFEEDAGEINVOTINGEPOCHS(&_FUpdater.CallOpts)
}

// MAXFEEDAGEINVOTINGEPOCHS is a free data retrieval call binding the contract method 0x7fe3341a.
//
// Solidity: function MAX_FEED_AGE_IN_VOTING_EPOCHS() view returns(uint256)
func (_FUpdater *FUpdaterCallerSession) MAXFEEDAGEINVOTINGEPOCHS() (*big.Int, error) {
	return _FUpdater.Contract.MAXFEEDAGEINVOTINGEPOCHS(&_FUpdater.CallOpts)
}

// BlockScoreCutoff is a free data retrieval call binding the contract method 0xdcb1476e.
//
// Solidity: function blockScoreCutoff(uint256 _blockNum) view returns(uint256 _cutoff)
func (_FUpdater *FUpdaterCaller) BlockScoreCutoff(opts *bind.CallOpts, _blockNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "blockScoreCutoff", _blockNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockScoreCutoff is a free data retrieval call binding the contract method 0xdcb1476e.
//
// Solidity: function blockScoreCutoff(uint256 _blockNum) view returns(uint256 _cutoff)
func (_FUpdater *FUpdaterSession) BlockScoreCutoff(_blockNum *big.Int) (*big.Int, error) {
	return _FUpdater.Contract.BlockScoreCutoff(&_FUpdater.CallOpts, _blockNum)
}

// BlockScoreCutoff is a free data retrieval call binding the contract method 0xdcb1476e.
//
// Solidity: function blockScoreCutoff(uint256 _blockNum) view returns(uint256 _cutoff)
func (_FUpdater *FUpdaterCallerSession) BlockScoreCutoff(_blockNum *big.Int) (*big.Int, error) {
	return _FUpdater.Contract.BlockScoreCutoff(&_FUpdater.CallOpts, _blockNum)
}

// CurrentRewardEpochId is a free data retrieval call binding the contract method 0x8e0e9f7c.
//
// Solidity: function currentRewardEpochId() view returns(uint24)
func (_FUpdater *FUpdaterCaller) CurrentRewardEpochId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "currentRewardEpochId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentRewardEpochId is a free data retrieval call binding the contract method 0x8e0e9f7c.
//
// Solidity: function currentRewardEpochId() view returns(uint24)
func (_FUpdater *FUpdaterSession) CurrentRewardEpochId() (*big.Int, error) {
	return _FUpdater.Contract.CurrentRewardEpochId(&_FUpdater.CallOpts)
}

// CurrentRewardEpochId is a free data retrieval call binding the contract method 0x8e0e9f7c.
//
// Solidity: function currentRewardEpochId() view returns(uint24)
func (_FUpdater *FUpdaterCallerSession) CurrentRewardEpochId() (*big.Int, error) {
	return _FUpdater.Contract.CurrentRewardEpochId(&_FUpdater.CallOpts)
}

// CurrentScoreCutoff is a free data retrieval call binding the contract method 0x0799fe75.
//
// Solidity: function currentScoreCutoff() view returns(uint256 _cutoff)
func (_FUpdater *FUpdaterCaller) CurrentScoreCutoff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "currentScoreCutoff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentScoreCutoff is a free data retrieval call binding the contract method 0x0799fe75.
//
// Solidity: function currentScoreCutoff() view returns(uint256 _cutoff)
func (_FUpdater *FUpdaterSession) CurrentScoreCutoff() (*big.Int, error) {
	return _FUpdater.Contract.CurrentScoreCutoff(&_FUpdater.CallOpts)
}

// CurrentScoreCutoff is a free data retrieval call binding the contract method 0x0799fe75.
//
// Solidity: function currentScoreCutoff() view returns(uint256 _cutoff)
func (_FUpdater *FUpdaterCallerSession) CurrentScoreCutoff() (*big.Int, error) {
	return _FUpdater.Contract.CurrentScoreCutoff(&_FUpdater.CallOpts)
}

// CurrentSortitionWeight is a free data retrieval call binding the contract method 0xa14634a7.
//
// Solidity: function currentSortitionWeight(address _signingPolicyAddress) view returns(uint256 _weight)
func (_FUpdater *FUpdaterCaller) CurrentSortitionWeight(opts *bind.CallOpts, _signingPolicyAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "currentSortitionWeight", _signingPolicyAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentSortitionWeight is a free data retrieval call binding the contract method 0xa14634a7.
//
// Solidity: function currentSortitionWeight(address _signingPolicyAddress) view returns(uint256 _weight)
func (_FUpdater *FUpdaterSession) CurrentSortitionWeight(_signingPolicyAddress common.Address) (*big.Int, error) {
	return _FUpdater.Contract.CurrentSortitionWeight(&_FUpdater.CallOpts, _signingPolicyAddress)
}

// CurrentSortitionWeight is a free data retrieval call binding the contract method 0xa14634a7.
//
// Solidity: function currentSortitionWeight(address _signingPolicyAddress) view returns(uint256 _weight)
func (_FUpdater *FUpdaterCallerSession) CurrentSortitionWeight(_signingPolicyAddress common.Address) (*big.Int, error) {
	return _FUpdater.Contract.CurrentSortitionWeight(&_FUpdater.CallOpts, _signingPolicyAddress)
}

// FastUpdateIncentiveManager is a free data retrieval call binding the contract method 0x7925eaca.
//
// Solidity: function fastUpdateIncentiveManager() view returns(address)
func (_FUpdater *FUpdaterCaller) FastUpdateIncentiveManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "fastUpdateIncentiveManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FastUpdateIncentiveManager is a free data retrieval call binding the contract method 0x7925eaca.
//
// Solidity: function fastUpdateIncentiveManager() view returns(address)
func (_FUpdater *FUpdaterSession) FastUpdateIncentiveManager() (common.Address, error) {
	return _FUpdater.Contract.FastUpdateIncentiveManager(&_FUpdater.CallOpts)
}

// FastUpdateIncentiveManager is a free data retrieval call binding the contract method 0x7925eaca.
//
// Solidity: function fastUpdateIncentiveManager() view returns(address)
func (_FUpdater *FUpdaterCallerSession) FastUpdateIncentiveManager() (common.Address, error) {
	return _FUpdater.Contract.FastUpdateIncentiveManager(&_FUpdater.CallOpts)
}

// FastUpdatesConfiguration is a free data retrieval call binding the contract method 0xc10f489a.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (_FUpdater *FUpdaterCaller) FastUpdatesConfiguration(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "fastUpdatesConfiguration")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FastUpdatesConfiguration is a free data retrieval call binding the contract method 0xc10f489a.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (_FUpdater *FUpdaterSession) FastUpdatesConfiguration() (common.Address, error) {
	return _FUpdater.Contract.FastUpdatesConfiguration(&_FUpdater.CallOpts)
}

// FastUpdatesConfiguration is a free data retrieval call binding the contract method 0xc10f489a.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (_FUpdater *FUpdaterCallerSession) FastUpdatesConfiguration() (common.Address, error) {
	return _FUpdater.Contract.FastUpdatesConfiguration(&_FUpdater.CallOpts)
}

// FetchAllCurrentFeeds is a free data retrieval call binding the contract method 0x4691377f.
//
// Solidity: function fetchAllCurrentFeeds() view returns(bytes21[] _feedIds, uint256[] _feeds, int8[] _decimals, uint64 _timestamp)
func (_FUpdater *FUpdaterCaller) FetchAllCurrentFeeds(opts *bind.CallOpts) (struct {
	FeedIds   [][21]byte
	Feeds     []*big.Int
	Decimals  []int8
	Timestamp uint64
}, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "fetchAllCurrentFeeds")

	outstruct := new(struct {
		FeedIds   [][21]byte
		Feeds     []*big.Int
		Decimals  []int8
		Timestamp uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeedIds = *abi.ConvertType(out[0], new([][21]byte)).(*[][21]byte)
	outstruct.Feeds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Decimals = *abi.ConvertType(out[2], new([]int8)).(*[]int8)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(uint64)).(*uint64)

	return *outstruct, err

}

// FetchAllCurrentFeeds is a free data retrieval call binding the contract method 0x4691377f.
//
// Solidity: function fetchAllCurrentFeeds() view returns(bytes21[] _feedIds, uint256[] _feeds, int8[] _decimals, uint64 _timestamp)
func (_FUpdater *FUpdaterSession) FetchAllCurrentFeeds() (struct {
	FeedIds   [][21]byte
	Feeds     []*big.Int
	Decimals  []int8
	Timestamp uint64
}, error) {
	return _FUpdater.Contract.FetchAllCurrentFeeds(&_FUpdater.CallOpts)
}

// FetchAllCurrentFeeds is a free data retrieval call binding the contract method 0x4691377f.
//
// Solidity: function fetchAllCurrentFeeds() view returns(bytes21[] _feedIds, uint256[] _feeds, int8[] _decimals, uint64 _timestamp)
func (_FUpdater *FUpdaterCallerSession) FetchAllCurrentFeeds() (struct {
	FeedIds   [][21]byte
	Feeds     []*big.Int
	Decimals  []int8
	Timestamp uint64
}, error) {
	return _FUpdater.Contract.FetchAllCurrentFeeds(&_FUpdater.CallOpts)
}

// FetchCurrentFeeds is a free data retrieval call binding the contract method 0x45a15d3c.
//
// Solidity: function fetchCurrentFeeds(uint256[] _indices) view returns(uint256[] _feeds, int8[] _decimals, uint64 _timestamp)
func (_FUpdater *FUpdaterCaller) FetchCurrentFeeds(opts *bind.CallOpts, _indices []*big.Int) (struct {
	Feeds     []*big.Int
	Decimals  []int8
	Timestamp uint64
}, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "fetchCurrentFeeds", _indices)

	outstruct := new(struct {
		Feeds     []*big.Int
		Decimals  []int8
		Timestamp uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Feeds = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Decimals = *abi.ConvertType(out[1], new([]int8)).(*[]int8)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// FetchCurrentFeeds is a free data retrieval call binding the contract method 0x45a15d3c.
//
// Solidity: function fetchCurrentFeeds(uint256[] _indices) view returns(uint256[] _feeds, int8[] _decimals, uint64 _timestamp)
func (_FUpdater *FUpdaterSession) FetchCurrentFeeds(_indices []*big.Int) (struct {
	Feeds     []*big.Int
	Decimals  []int8
	Timestamp uint64
}, error) {
	return _FUpdater.Contract.FetchCurrentFeeds(&_FUpdater.CallOpts, _indices)
}

// FetchCurrentFeeds is a free data retrieval call binding the contract method 0x45a15d3c.
//
// Solidity: function fetchCurrentFeeds(uint256[] _indices) view returns(uint256[] _feeds, int8[] _decimals, uint64 _timestamp)
func (_FUpdater *FUpdaterCallerSession) FetchCurrentFeeds(_indices []*big.Int) (struct {
	Feeds     []*big.Int
	Decimals  []int8
	Timestamp uint64
}, error) {
	return _FUpdater.Contract.FetchCurrentFeeds(&_FUpdater.CallOpts, _indices)
}

// FirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (_FUpdater *FUpdaterCaller) FirstVotingRoundStartTs(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "firstVotingRoundStartTs")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// FirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (_FUpdater *FUpdaterSession) FirstVotingRoundStartTs() (uint64, error) {
	return _FUpdater.Contract.FirstVotingRoundStartTs(&_FUpdater.CallOpts)
}

// FirstVotingRoundStartTs is a free data retrieval call binding the contract method 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (_FUpdater *FUpdaterCallerSession) FirstVotingRoundStartTs() (uint64, error) {
	return _FUpdater.Contract.FirstVotingRoundStartTs(&_FUpdater.CallOpts)
}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_FUpdater *FUpdaterCaller) FlareDaemon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "flareDaemon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_FUpdater *FUpdaterSession) FlareDaemon() (common.Address, error) {
	return _FUpdater.Contract.FlareDaemon(&_FUpdater.CallOpts)
}

// FlareDaemon is a free data retrieval call binding the contract method 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (_FUpdater *FUpdaterCallerSession) FlareDaemon() (common.Address, error) {
	return _FUpdater.Contract.FlareDaemon(&_FUpdater.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_FUpdater *FUpdaterCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_FUpdater *FUpdaterSession) FlareSystemsManager() (common.Address, error) {
	return _FUpdater.Contract.FlareSystemsManager(&_FUpdater.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_FUpdater *FUpdaterCallerSession) FlareSystemsManager() (common.Address, error) {
	return _FUpdater.Contract.FlareSystemsManager(&_FUpdater.CallOpts)
}

// FtsoFeedPublisher is a free data retrieval call binding the contract method 0x29bfe39d.
//
// Solidity: function ftsoFeedPublisher() view returns(address)
func (_FUpdater *FUpdaterCaller) FtsoFeedPublisher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "ftsoFeedPublisher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FtsoFeedPublisher is a free data retrieval call binding the contract method 0x29bfe39d.
//
// Solidity: function ftsoFeedPublisher() view returns(address)
func (_FUpdater *FUpdaterSession) FtsoFeedPublisher() (common.Address, error) {
	return _FUpdater.Contract.FtsoFeedPublisher(&_FUpdater.CallOpts)
}

// FtsoFeedPublisher is a free data retrieval call binding the contract method 0x29bfe39d.
//
// Solidity: function ftsoFeedPublisher() view returns(address)
func (_FUpdater *FUpdaterCallerSession) FtsoFeedPublisher() (common.Address, error) {
	return _FUpdater.Contract.FtsoFeedPublisher(&_FUpdater.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FUpdater *FUpdaterCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FUpdater *FUpdaterSession) GetAddressUpdater() (common.Address, error) {
	return _FUpdater.Contract.GetAddressUpdater(&_FUpdater.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_FUpdater *FUpdaterCallerSession) GetAddressUpdater() (common.Address, error) {
	return _FUpdater.Contract.GetAddressUpdater(&_FUpdater.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FUpdater *FUpdaterCaller) GetContractName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "getContractName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FUpdater *FUpdaterSession) GetContractName() (string, error) {
	return _FUpdater.Contract.GetContractName(&_FUpdater.CallOpts)
}

// GetContractName is a free data retrieval call binding the contract method 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (_FUpdater *FUpdaterCallerSession) GetContractName() (string, error) {
	return _FUpdater.Contract.GetContractName(&_FUpdater.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FUpdater *FUpdaterCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FUpdater *FUpdaterSession) Governance() (common.Address, error) {
	return _FUpdater.Contract.Governance(&_FUpdater.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_FUpdater *FUpdaterCallerSession) Governance() (common.Address, error) {
	return _FUpdater.Contract.Governance(&_FUpdater.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FUpdater *FUpdaterCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FUpdater *FUpdaterSession) GovernanceSettings() (common.Address, error) {
	return _FUpdater.Contract.GovernanceSettings(&_FUpdater.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_FUpdater *FUpdaterCallerSession) GovernanceSettings() (common.Address, error) {
	return _FUpdater.Contract.GovernanceSettings(&_FUpdater.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FUpdater *FUpdaterCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FUpdater *FUpdaterSession) IsExecutor(_address common.Address) (bool, error) {
	return _FUpdater.Contract.IsExecutor(&_FUpdater.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_FUpdater *FUpdaterCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _FUpdater.Contract.IsExecutor(&_FUpdater.CallOpts, _address)
}

// NumberOfUpdates is a free data retrieval call binding the contract method 0xe36da7b7.
//
// Solidity: function numberOfUpdates(uint256 _historySize) view returns(uint256[] _noOfUpdates)
func (_FUpdater *FUpdaterCaller) NumberOfUpdates(opts *bind.CallOpts, _historySize *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "numberOfUpdates", _historySize)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// NumberOfUpdates is a free data retrieval call binding the contract method 0xe36da7b7.
//
// Solidity: function numberOfUpdates(uint256 _historySize) view returns(uint256[] _noOfUpdates)
func (_FUpdater *FUpdaterSession) NumberOfUpdates(_historySize *big.Int) ([]*big.Int, error) {
	return _FUpdater.Contract.NumberOfUpdates(&_FUpdater.CallOpts, _historySize)
}

// NumberOfUpdates is a free data retrieval call binding the contract method 0xe36da7b7.
//
// Solidity: function numberOfUpdates(uint256 _historySize) view returns(uint256[] _noOfUpdates)
func (_FUpdater *FUpdaterCallerSession) NumberOfUpdates(_historySize *big.Int) ([]*big.Int, error) {
	return _FUpdater.Contract.NumberOfUpdates(&_FUpdater.CallOpts, _historySize)
}

// NumberOfUpdatesInBlock is a free data retrieval call binding the contract method 0xfc79c300.
//
// Solidity: function numberOfUpdatesInBlock(uint256 _blockNumber) view returns(uint256)
func (_FUpdater *FUpdaterCaller) NumberOfUpdatesInBlock(opts *bind.CallOpts, _blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "numberOfUpdatesInBlock", _blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfUpdatesInBlock is a free data retrieval call binding the contract method 0xfc79c300.
//
// Solidity: function numberOfUpdatesInBlock(uint256 _blockNumber) view returns(uint256)
func (_FUpdater *FUpdaterSession) NumberOfUpdatesInBlock(_blockNumber *big.Int) (*big.Int, error) {
	return _FUpdater.Contract.NumberOfUpdatesInBlock(&_FUpdater.CallOpts, _blockNumber)
}

// NumberOfUpdatesInBlock is a free data retrieval call binding the contract method 0xfc79c300.
//
// Solidity: function numberOfUpdatesInBlock(uint256 _blockNumber) view returns(uint256)
func (_FUpdater *FUpdaterCallerSession) NumberOfUpdatesInBlock(_blockNumber *big.Int) (*big.Int, error) {
	return _FUpdater.Contract.NumberOfUpdatesInBlock(&_FUpdater.CallOpts, _blockNumber)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FUpdater *FUpdaterCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FUpdater *FUpdaterSession) ProductionMode() (bool, error) {
	return _FUpdater.Contract.ProductionMode(&_FUpdater.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_FUpdater *FUpdaterCallerSession) ProductionMode() (bool, error) {
	return _FUpdater.Contract.ProductionMode(&_FUpdater.CallOpts)
}

// SubmissionWindow is a free data retrieval call binding the contract method 0xe621dbc7.
//
// Solidity: function submissionWindow() view returns(uint8)
func (_FUpdater *FUpdaterCaller) SubmissionWindow(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "submissionWindow")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SubmissionWindow is a free data retrieval call binding the contract method 0xe621dbc7.
//
// Solidity: function submissionWindow() view returns(uint8)
func (_FUpdater *FUpdaterSession) SubmissionWindow() (uint8, error) {
	return _FUpdater.Contract.SubmissionWindow(&_FUpdater.CallOpts)
}

// SubmissionWindow is a free data retrieval call binding the contract method 0xe621dbc7.
//
// Solidity: function submissionWindow() view returns(uint8)
func (_FUpdater *FUpdaterCallerSession) SubmissionWindow() (uint8, error) {
	return _FUpdater.Contract.SubmissionWindow(&_FUpdater.CallOpts)
}

// SwitchToFallbackMode is a free data retrieval call binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (_FUpdater *FUpdaterCaller) SwitchToFallbackMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "switchToFallbackMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SwitchToFallbackMode is a free data retrieval call binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (_FUpdater *FUpdaterSession) SwitchToFallbackMode() (bool, error) {
	return _FUpdater.Contract.SwitchToFallbackMode(&_FUpdater.CallOpts)
}

// SwitchToFallbackMode is a free data retrieval call binding the contract method 0xe22fdece.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (_FUpdater *FUpdaterCallerSession) SwitchToFallbackMode() (bool, error) {
	return _FUpdater.Contract.SwitchToFallbackMode(&_FUpdater.CallOpts)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FUpdater *FUpdaterCaller) TimelockedCalls(opts *bind.CallOpts, selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "timelockedCalls", selector)

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
func (_FUpdater *FUpdaterSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FUpdater.Contract.TimelockedCalls(&_FUpdater.CallOpts, selector)
}

// TimelockedCalls is a free data retrieval call binding the contract method 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FUpdater *FUpdaterCallerSession) TimelockedCalls(selector [4]byte) (struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}, error) {
	return _FUpdater.Contract.TimelockedCalls(&_FUpdater.CallOpts, selector)
}

// VerifyPublicKey is a free data retrieval call binding the contract method 0x70473f2f.
//
// Solidity: function verifyPublicKey(address _voter, bytes32 _part1, bytes32 _part2, bytes _verificationData) view returns()
func (_FUpdater *FUpdaterCaller) VerifyPublicKey(opts *bind.CallOpts, _voter common.Address, _part1 [32]byte, _part2 [32]byte, _verificationData []byte) error {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "verifyPublicKey", _voter, _part1, _part2, _verificationData)

	if err != nil {
		return err
	}

	return err

}

// VerifyPublicKey is a free data retrieval call binding the contract method 0x70473f2f.
//
// Solidity: function verifyPublicKey(address _voter, bytes32 _part1, bytes32 _part2, bytes _verificationData) view returns()
func (_FUpdater *FUpdaterSession) VerifyPublicKey(_voter common.Address, _part1 [32]byte, _part2 [32]byte, _verificationData []byte) error {
	return _FUpdater.Contract.VerifyPublicKey(&_FUpdater.CallOpts, _voter, _part1, _part2, _verificationData)
}

// VerifyPublicKey is a free data retrieval call binding the contract method 0x70473f2f.
//
// Solidity: function verifyPublicKey(address _voter, bytes32 _part1, bytes32 _part2, bytes _verificationData) view returns()
func (_FUpdater *FUpdaterCallerSession) VerifyPublicKey(_voter common.Address, _part1 [32]byte, _part2 [32]byte, _verificationData []byte) error {
	return _FUpdater.Contract.VerifyPublicKey(&_FUpdater.CallOpts, _voter, _part1, _part2, _verificationData)
}

// VoterRegistry is a free data retrieval call binding the contract method 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (_FUpdater *FUpdaterCaller) VoterRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "voterRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoterRegistry is a free data retrieval call binding the contract method 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (_FUpdater *FUpdaterSession) VoterRegistry() (common.Address, error) {
	return _FUpdater.Contract.VoterRegistry(&_FUpdater.CallOpts)
}

// VoterRegistry is a free data retrieval call binding the contract method 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (_FUpdater *FUpdaterCallerSession) VoterRegistry() (common.Address, error) {
	return _FUpdater.Contract.VoterRegistry(&_FUpdater.CallOpts)
}

// VotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (_FUpdater *FUpdaterCaller) VotingEpochDurationSeconds(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FUpdater.contract.Call(opts, &out, "votingEpochDurationSeconds")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// VotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (_FUpdater *FUpdaterSession) VotingEpochDurationSeconds() (uint64, error) {
	return _FUpdater.Contract.VotingEpochDurationSeconds(&_FUpdater.CallOpts)
}

// VotingEpochDurationSeconds is a free data retrieval call binding the contract method 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (_FUpdater *FUpdaterCallerSession) VotingEpochDurationSeconds() (uint64, error) {
	return _FUpdater.Contract.VotingEpochDurationSeconds(&_FUpdater.CallOpts)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FUpdater *FUpdaterTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FUpdater.contract.Transact(opts, "cancelGovernanceCall", _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FUpdater *FUpdaterSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FUpdater.Contract.CancelGovernanceCall(&_FUpdater.TransactOpts, _selector)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x67fc4029.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (_FUpdater *FUpdaterTransactorSession) CancelGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FUpdater.Contract.CancelGovernanceCall(&_FUpdater.TransactOpts, _selector)
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_FUpdater *FUpdaterTransactor) Daemonize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FUpdater.contract.Transact(opts, "daemonize")
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_FUpdater *FUpdaterSession) Daemonize() (*types.Transaction, error) {
	return _FUpdater.Contract.Daemonize(&_FUpdater.TransactOpts)
}

// Daemonize is a paid mutator transaction binding the contract method 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (_FUpdater *FUpdaterTransactorSession) Daemonize() (*types.Transaction, error) {
	return _FUpdater.Contract.Daemonize(&_FUpdater.TransactOpts)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FUpdater *FUpdaterTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _selector [4]byte) (*types.Transaction, error) {
	return _FUpdater.contract.Transact(opts, "executeGovernanceCall", _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FUpdater *FUpdaterSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FUpdater.Contract.ExecuteGovernanceCall(&_FUpdater.TransactOpts, _selector)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x5ff27079.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (_FUpdater *FUpdaterTransactorSession) ExecuteGovernanceCall(_selector [4]byte) (*types.Transaction, error) {
	return _FUpdater.Contract.ExecuteGovernanceCall(&_FUpdater.TransactOpts, _selector)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FUpdater *FUpdaterTransactor) Initialise(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FUpdater.contract.Transact(opts, "initialise", _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FUpdater *FUpdaterSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FUpdater.Contract.Initialise(&_FUpdater.TransactOpts, _governanceSettings, _initialGovernance)
}

// Initialise is a paid mutator transaction binding the contract method 0xef88bf13.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (_FUpdater *FUpdaterTransactorSession) Initialise(_governanceSettings common.Address, _initialGovernance common.Address) (*types.Transaction, error) {
	return _FUpdater.Contract.Initialise(&_FUpdater.TransactOpts, _governanceSettings, _initialGovernance)
}

// RemoveFeeds is a paid mutator transaction binding the contract method 0xabfaf170.
//
// Solidity: function removeFeeds(uint256[] _indices) returns()
func (_FUpdater *FUpdaterTransactor) RemoveFeeds(opts *bind.TransactOpts, _indices []*big.Int) (*types.Transaction, error) {
	return _FUpdater.contract.Transact(opts, "removeFeeds", _indices)
}

// RemoveFeeds is a paid mutator transaction binding the contract method 0xabfaf170.
//
// Solidity: function removeFeeds(uint256[] _indices) returns()
func (_FUpdater *FUpdaterSession) RemoveFeeds(_indices []*big.Int) (*types.Transaction, error) {
	return _FUpdater.Contract.RemoveFeeds(&_FUpdater.TransactOpts, _indices)
}

// RemoveFeeds is a paid mutator transaction binding the contract method 0xabfaf170.
//
// Solidity: function removeFeeds(uint256[] _indices) returns()
func (_FUpdater *FUpdaterTransactorSession) RemoveFeeds(_indices []*big.Int) (*types.Transaction, error) {
	return _FUpdater.Contract.RemoveFeeds(&_FUpdater.TransactOpts, _indices)
}

// ResetFeeds is a paid mutator transaction binding the contract method 0x63f921db.
//
// Solidity: function resetFeeds(uint256[] _indices) returns()
func (_FUpdater *FUpdaterTransactor) ResetFeeds(opts *bind.TransactOpts, _indices []*big.Int) (*types.Transaction, error) {
	return _FUpdater.contract.Transact(opts, "resetFeeds", _indices)
}

// ResetFeeds is a paid mutator transaction binding the contract method 0x63f921db.
//
// Solidity: function resetFeeds(uint256[] _indices) returns()
func (_FUpdater *FUpdaterSession) ResetFeeds(_indices []*big.Int) (*types.Transaction, error) {
	return _FUpdater.Contract.ResetFeeds(&_FUpdater.TransactOpts, _indices)
}

// ResetFeeds is a paid mutator transaction binding the contract method 0x63f921db.
//
// Solidity: function resetFeeds(uint256[] _indices) returns()
func (_FUpdater *FUpdaterTransactorSession) ResetFeeds(_indices []*big.Int) (*types.Transaction, error) {
	return _FUpdater.Contract.ResetFeeds(&_FUpdater.TransactOpts, _indices)
}

// SetSubmissionWindow is a paid mutator transaction binding the contract method 0x0a166051.
//
// Solidity: function setSubmissionWindow(uint8 _submissionWindow) returns()
func (_FUpdater *FUpdaterTransactor) SetSubmissionWindow(opts *bind.TransactOpts, _submissionWindow uint8) (*types.Transaction, error) {
	return _FUpdater.contract.Transact(opts, "setSubmissionWindow", _submissionWindow)
}

// SetSubmissionWindow is a paid mutator transaction binding the contract method 0x0a166051.
//
// Solidity: function setSubmissionWindow(uint8 _submissionWindow) returns()
func (_FUpdater *FUpdaterSession) SetSubmissionWindow(_submissionWindow uint8) (*types.Transaction, error) {
	return _FUpdater.Contract.SetSubmissionWindow(&_FUpdater.TransactOpts, _submissionWindow)
}

// SetSubmissionWindow is a paid mutator transaction binding the contract method 0x0a166051.
//
// Solidity: function setSubmissionWindow(uint8 _submissionWindow) returns()
func (_FUpdater *FUpdaterTransactorSession) SetSubmissionWindow(_submissionWindow uint8) (*types.Transaction, error) {
	return _FUpdater.Contract.SetSubmissionWindow(&_FUpdater.TransactOpts, _submissionWindow)
}

// SubmitUpdates is a paid mutator transaction binding the contract method 0x470e91df.
//
// Solidity: function submitUpdates((uint256,(uint256,(uint256,uint256),uint256,uint256),bytes,(uint8,bytes32,bytes32)) _updates) returns()
func (_FUpdater *FUpdaterTransactor) SubmitUpdates(opts *bind.TransactOpts, _updates IFastUpdaterFastUpdates) (*types.Transaction, error) {
	return _FUpdater.contract.Transact(opts, "submitUpdates", _updates)
}

// SubmitUpdates is a paid mutator transaction binding the contract method 0x470e91df.
//
// Solidity: function submitUpdates((uint256,(uint256,(uint256,uint256),uint256,uint256),bytes,(uint8,bytes32,bytes32)) _updates) returns()
func (_FUpdater *FUpdaterSession) SubmitUpdates(_updates IFastUpdaterFastUpdates) (*types.Transaction, error) {
	return _FUpdater.Contract.SubmitUpdates(&_FUpdater.TransactOpts, _updates)
}

// SubmitUpdates is a paid mutator transaction binding the contract method 0x470e91df.
//
// Solidity: function submitUpdates((uint256,(uint256,(uint256,uint256),uint256,uint256),bytes,(uint8,bytes32,bytes32)) _updates) returns()
func (_FUpdater *FUpdaterTransactorSession) SubmitUpdates(_updates IFastUpdaterFastUpdates) (*types.Transaction, error) {
	return _FUpdater.Contract.SubmitUpdates(&_FUpdater.TransactOpts, _updates)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FUpdater *FUpdaterTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FUpdater.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FUpdater *FUpdaterSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FUpdater.Contract.SwitchToProductionMode(&_FUpdater.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_FUpdater *FUpdaterTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _FUpdater.Contract.SwitchToProductionMode(&_FUpdater.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FUpdater *FUpdaterTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FUpdater.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FUpdater *FUpdaterSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FUpdater.Contract.UpdateContractAddresses(&_FUpdater.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_FUpdater *FUpdaterTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _FUpdater.Contract.UpdateContractAddresses(&_FUpdater.TransactOpts, _contractNameHashes, _contractAddresses)
}

// FUpdaterFastUpdateFeedRemovedIterator is returned from FilterFastUpdateFeedRemoved and is used to iterate over the raw logs and unpacked data for FastUpdateFeedRemoved events raised by the FUpdater contract.
type FUpdaterFastUpdateFeedRemovedIterator struct {
	Event *FUpdaterFastUpdateFeedRemoved // Event containing the contract specifics and raw log

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
func (it *FUpdaterFastUpdateFeedRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUpdaterFastUpdateFeedRemoved)
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
		it.Event = new(FUpdaterFastUpdateFeedRemoved)
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
func (it *FUpdaterFastUpdateFeedRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUpdaterFastUpdateFeedRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUpdaterFastUpdateFeedRemoved represents a FastUpdateFeedRemoved event raised by the FUpdater contract.
type FUpdaterFastUpdateFeedRemoved struct {
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterFastUpdateFeedRemoved is a free log retrieval operation binding the contract event 0x5a17f564b9fd53c971695a2e033e0cd39ee8ad08d8117242cdecad8b017335c8.
//
// Solidity: event FastUpdateFeedRemoved(uint256 indexed index)
func (_FUpdater *FUpdaterFilterer) FilterFastUpdateFeedRemoved(opts *bind.FilterOpts, index []*big.Int) (*FUpdaterFastUpdateFeedRemovedIterator, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _FUpdater.contract.FilterLogs(opts, "FastUpdateFeedRemoved", indexRule)
	if err != nil {
		return nil, err
	}
	return &FUpdaterFastUpdateFeedRemovedIterator{contract: _FUpdater.contract, event: "FastUpdateFeedRemoved", logs: logs, sub: sub}, nil
}

// WatchFastUpdateFeedRemoved is a free log subscription operation binding the contract event 0x5a17f564b9fd53c971695a2e033e0cd39ee8ad08d8117242cdecad8b017335c8.
//
// Solidity: event FastUpdateFeedRemoved(uint256 indexed index)
func (_FUpdater *FUpdaterFilterer) WatchFastUpdateFeedRemoved(opts *bind.WatchOpts, sink chan<- *FUpdaterFastUpdateFeedRemoved, index []*big.Int) (event.Subscription, error) {

	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}

	logs, sub, err := _FUpdater.contract.WatchLogs(opts, "FastUpdateFeedRemoved", indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUpdaterFastUpdateFeedRemoved)
				if err := _FUpdater.contract.UnpackLog(event, "FastUpdateFeedRemoved", log); err != nil {
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

// ParseFastUpdateFeedRemoved is a log parse operation binding the contract event 0x5a17f564b9fd53c971695a2e033e0cd39ee8ad08d8117242cdecad8b017335c8.
//
// Solidity: event FastUpdateFeedRemoved(uint256 indexed index)
func (_FUpdater *FUpdaterFilterer) ParseFastUpdateFeedRemoved(log types.Log) (*FUpdaterFastUpdateFeedRemoved, error) {
	event := new(FUpdaterFastUpdateFeedRemoved)
	if err := _FUpdater.contract.UnpackLog(event, "FastUpdateFeedRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUpdaterFastUpdateFeedResetIterator is returned from FilterFastUpdateFeedReset and is used to iterate over the raw logs and unpacked data for FastUpdateFeedReset events raised by the FUpdater contract.
type FUpdaterFastUpdateFeedResetIterator struct {
	Event *FUpdaterFastUpdateFeedReset // Event containing the contract specifics and raw log

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
func (it *FUpdaterFastUpdateFeedResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUpdaterFastUpdateFeedReset)
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
		it.Event = new(FUpdaterFastUpdateFeedReset)
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
func (it *FUpdaterFastUpdateFeedResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUpdaterFastUpdateFeedResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUpdaterFastUpdateFeedReset represents a FastUpdateFeedReset event raised by the FUpdater contract.
type FUpdaterFastUpdateFeedReset struct {
	VotingRoundId *big.Int
	Index         *big.Int
	Id            [21]byte
	Value         *big.Int
	Decimals      int8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFastUpdateFeedReset is a free log retrieval operation binding the contract event 0xfa800fadb5e7b72652da40bcd7ca9a96cd4e53c9ea4c68b8afbba027e77a7cf5.
//
// Solidity: event FastUpdateFeedReset(uint256 indexed votingRoundId, uint256 indexed index, bytes21 indexed id, uint256 value, int8 decimals)
func (_FUpdater *FUpdaterFilterer) FilterFastUpdateFeedReset(opts *bind.FilterOpts, votingRoundId []*big.Int, index []*big.Int, id [][21]byte) (*FUpdaterFastUpdateFeedResetIterator, error) {

	var votingRoundIdRule []interface{}
	for _, votingRoundIdItem := range votingRoundId {
		votingRoundIdRule = append(votingRoundIdRule, votingRoundIdItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FUpdater.contract.FilterLogs(opts, "FastUpdateFeedReset", votingRoundIdRule, indexRule, idRule)
	if err != nil {
		return nil, err
	}
	return &FUpdaterFastUpdateFeedResetIterator{contract: _FUpdater.contract, event: "FastUpdateFeedReset", logs: logs, sub: sub}, nil
}

// WatchFastUpdateFeedReset is a free log subscription operation binding the contract event 0xfa800fadb5e7b72652da40bcd7ca9a96cd4e53c9ea4c68b8afbba027e77a7cf5.
//
// Solidity: event FastUpdateFeedReset(uint256 indexed votingRoundId, uint256 indexed index, bytes21 indexed id, uint256 value, int8 decimals)
func (_FUpdater *FUpdaterFilterer) WatchFastUpdateFeedReset(opts *bind.WatchOpts, sink chan<- *FUpdaterFastUpdateFeedReset, votingRoundId []*big.Int, index []*big.Int, id [][21]byte) (event.Subscription, error) {

	var votingRoundIdRule []interface{}
	for _, votingRoundIdItem := range votingRoundId {
		votingRoundIdRule = append(votingRoundIdRule, votingRoundIdItem)
	}
	var indexRule []interface{}
	for _, indexItem := range index {
		indexRule = append(indexRule, indexItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _FUpdater.contract.WatchLogs(opts, "FastUpdateFeedReset", votingRoundIdRule, indexRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUpdaterFastUpdateFeedReset)
				if err := _FUpdater.contract.UnpackLog(event, "FastUpdateFeedReset", log); err != nil {
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

// ParseFastUpdateFeedReset is a log parse operation binding the contract event 0xfa800fadb5e7b72652da40bcd7ca9a96cd4e53c9ea4c68b8afbba027e77a7cf5.
//
// Solidity: event FastUpdateFeedReset(uint256 indexed votingRoundId, uint256 indexed index, bytes21 indexed id, uint256 value, int8 decimals)
func (_FUpdater *FUpdaterFilterer) ParseFastUpdateFeedReset(log types.Log) (*FUpdaterFastUpdateFeedReset, error) {
	event := new(FUpdaterFastUpdateFeedReset)
	if err := _FUpdater.contract.UnpackLog(event, "FastUpdateFeedReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUpdaterFastUpdateFeedsIterator is returned from FilterFastUpdateFeeds and is used to iterate over the raw logs and unpacked data for FastUpdateFeeds events raised by the FUpdater contract.
type FUpdaterFastUpdateFeedsIterator struct {
	Event *FUpdaterFastUpdateFeeds // Event containing the contract specifics and raw log

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
func (it *FUpdaterFastUpdateFeedsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUpdaterFastUpdateFeeds)
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
		it.Event = new(FUpdaterFastUpdateFeeds)
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
func (it *FUpdaterFastUpdateFeedsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUpdaterFastUpdateFeedsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUpdaterFastUpdateFeeds represents a FastUpdateFeeds event raised by the FUpdater contract.
type FUpdaterFastUpdateFeeds struct {
	VotingEpochId *big.Int
	Feeds         []*big.Int
	Decimals      []int8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterFastUpdateFeeds is a free log retrieval operation binding the contract event 0x7195d3bbca575b55b0f4b62da395f7224b75225d9c08741d216e6cf10a83eabe.
//
// Solidity: event FastUpdateFeeds(uint256 indexed votingEpochId, uint256[] feeds, int8[] decimals)
func (_FUpdater *FUpdaterFilterer) FilterFastUpdateFeeds(opts *bind.FilterOpts, votingEpochId []*big.Int) (*FUpdaterFastUpdateFeedsIterator, error) {

	var votingEpochIdRule []interface{}
	for _, votingEpochIdItem := range votingEpochId {
		votingEpochIdRule = append(votingEpochIdRule, votingEpochIdItem)
	}

	logs, sub, err := _FUpdater.contract.FilterLogs(opts, "FastUpdateFeeds", votingEpochIdRule)
	if err != nil {
		return nil, err
	}
	return &FUpdaterFastUpdateFeedsIterator{contract: _FUpdater.contract, event: "FastUpdateFeeds", logs: logs, sub: sub}, nil
}

// WatchFastUpdateFeeds is a free log subscription operation binding the contract event 0x7195d3bbca575b55b0f4b62da395f7224b75225d9c08741d216e6cf10a83eabe.
//
// Solidity: event FastUpdateFeeds(uint256 indexed votingEpochId, uint256[] feeds, int8[] decimals)
func (_FUpdater *FUpdaterFilterer) WatchFastUpdateFeeds(opts *bind.WatchOpts, sink chan<- *FUpdaterFastUpdateFeeds, votingEpochId []*big.Int) (event.Subscription, error) {

	var votingEpochIdRule []interface{}
	for _, votingEpochIdItem := range votingEpochId {
		votingEpochIdRule = append(votingEpochIdRule, votingEpochIdItem)
	}

	logs, sub, err := _FUpdater.contract.WatchLogs(opts, "FastUpdateFeeds", votingEpochIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUpdaterFastUpdateFeeds)
				if err := _FUpdater.contract.UnpackLog(event, "FastUpdateFeeds", log); err != nil {
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

// ParseFastUpdateFeeds is a log parse operation binding the contract event 0x7195d3bbca575b55b0f4b62da395f7224b75225d9c08741d216e6cf10a83eabe.
//
// Solidity: event FastUpdateFeeds(uint256 indexed votingEpochId, uint256[] feeds, int8[] decimals)
func (_FUpdater *FUpdaterFilterer) ParseFastUpdateFeeds(log types.Log) (*FUpdaterFastUpdateFeeds, error) {
	event := new(FUpdaterFastUpdateFeeds)
	if err := _FUpdater.contract.UnpackLog(event, "FastUpdateFeeds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUpdaterFastUpdateFeedsSubmittedIterator is returned from FilterFastUpdateFeedsSubmitted and is used to iterate over the raw logs and unpacked data for FastUpdateFeedsSubmitted events raised by the FUpdater contract.
type FUpdaterFastUpdateFeedsSubmittedIterator struct {
	Event *FUpdaterFastUpdateFeedsSubmitted // Event containing the contract specifics and raw log

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
func (it *FUpdaterFastUpdateFeedsSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUpdaterFastUpdateFeedsSubmitted)
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
		it.Event = new(FUpdaterFastUpdateFeedsSubmitted)
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
func (it *FUpdaterFastUpdateFeedsSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUpdaterFastUpdateFeedsSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUpdaterFastUpdateFeedsSubmitted represents a FastUpdateFeedsSubmitted event raised by the FUpdater contract.
type FUpdaterFastUpdateFeedsSubmitted struct {
	VotingRoundId        uint32
	SigningPolicyAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterFastUpdateFeedsSubmitted is a free log retrieval operation binding the contract event 0x63db91b14b3d088c677f046180aefcea7a236649704d90ce810cde455d38d936.
//
// Solidity: event FastUpdateFeedsSubmitted(uint32 indexed votingRoundId, address indexed signingPolicyAddress)
func (_FUpdater *FUpdaterFilterer) FilterFastUpdateFeedsSubmitted(opts *bind.FilterOpts, votingRoundId []uint32, signingPolicyAddress []common.Address) (*FUpdaterFastUpdateFeedsSubmittedIterator, error) {

	var votingRoundIdRule []interface{}
	for _, votingRoundIdItem := range votingRoundId {
		votingRoundIdRule = append(votingRoundIdRule, votingRoundIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}

	logs, sub, err := _FUpdater.contract.FilterLogs(opts, "FastUpdateFeedsSubmitted", votingRoundIdRule, signingPolicyAddressRule)
	if err != nil {
		return nil, err
	}
	return &FUpdaterFastUpdateFeedsSubmittedIterator{contract: _FUpdater.contract, event: "FastUpdateFeedsSubmitted", logs: logs, sub: sub}, nil
}

// WatchFastUpdateFeedsSubmitted is a free log subscription operation binding the contract event 0x63db91b14b3d088c677f046180aefcea7a236649704d90ce810cde455d38d936.
//
// Solidity: event FastUpdateFeedsSubmitted(uint32 indexed votingRoundId, address indexed signingPolicyAddress)
func (_FUpdater *FUpdaterFilterer) WatchFastUpdateFeedsSubmitted(opts *bind.WatchOpts, sink chan<- *FUpdaterFastUpdateFeedsSubmitted, votingRoundId []uint32, signingPolicyAddress []common.Address) (event.Subscription, error) {

	var votingRoundIdRule []interface{}
	for _, votingRoundIdItem := range votingRoundId {
		votingRoundIdRule = append(votingRoundIdRule, votingRoundIdItem)
	}
	var signingPolicyAddressRule []interface{}
	for _, signingPolicyAddressItem := range signingPolicyAddress {
		signingPolicyAddressRule = append(signingPolicyAddressRule, signingPolicyAddressItem)
	}

	logs, sub, err := _FUpdater.contract.WatchLogs(opts, "FastUpdateFeedsSubmitted", votingRoundIdRule, signingPolicyAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUpdaterFastUpdateFeedsSubmitted)
				if err := _FUpdater.contract.UnpackLog(event, "FastUpdateFeedsSubmitted", log); err != nil {
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

// ParseFastUpdateFeedsSubmitted is a log parse operation binding the contract event 0x63db91b14b3d088c677f046180aefcea7a236649704d90ce810cde455d38d936.
//
// Solidity: event FastUpdateFeedsSubmitted(uint32 indexed votingRoundId, address indexed signingPolicyAddress)
func (_FUpdater *FUpdaterFilterer) ParseFastUpdateFeedsSubmitted(log types.Log) (*FUpdaterFastUpdateFeedsSubmitted, error) {
	event := new(FUpdaterFastUpdateFeedsSubmitted)
	if err := _FUpdater.contract.UnpackLog(event, "FastUpdateFeedsSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUpdaterGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the FUpdater contract.
type FUpdaterGovernanceCallTimelockedIterator struct {
	Event *FUpdaterGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *FUpdaterGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUpdaterGovernanceCallTimelocked)
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
		it.Event = new(FUpdaterGovernanceCallTimelocked)
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
func (it *FUpdaterGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUpdaterGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUpdaterGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FUpdater contract.
type FUpdaterGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FUpdater *FUpdaterFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*FUpdaterGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _FUpdater.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &FUpdaterGovernanceCallTimelockedIterator{contract: _FUpdater.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0xed948300a3694aa01d4a6b258bfd664350193d770c0b51f8387277f6d83ea3b6.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (_FUpdater *FUpdaterFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *FUpdaterGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _FUpdater.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUpdaterGovernanceCallTimelocked)
				if err := _FUpdater.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_FUpdater *FUpdaterFilterer) ParseGovernanceCallTimelocked(log types.Log) (*FUpdaterGovernanceCallTimelocked, error) {
	event := new(FUpdaterGovernanceCallTimelocked)
	if err := _FUpdater.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUpdaterGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the FUpdater contract.
type FUpdaterGovernanceInitialisedIterator struct {
	Event *FUpdaterGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *FUpdaterGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUpdaterGovernanceInitialised)
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
		it.Event = new(FUpdaterGovernanceInitialised)
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
func (it *FUpdaterGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUpdaterGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUpdaterGovernanceInitialised represents a GovernanceInitialised event raised by the FUpdater contract.
type FUpdaterGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FUpdater *FUpdaterFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*FUpdaterGovernanceInitialisedIterator, error) {

	logs, sub, err := _FUpdater.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &FUpdaterGovernanceInitialisedIterator{contract: _FUpdater.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_FUpdater *FUpdaterFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *FUpdaterGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _FUpdater.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUpdaterGovernanceInitialised)
				if err := _FUpdater.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_FUpdater *FUpdaterFilterer) ParseGovernanceInitialised(log types.Log) (*FUpdaterGovernanceInitialised, error) {
	event := new(FUpdaterGovernanceInitialised)
	if err := _FUpdater.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUpdaterGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the FUpdater contract.
type FUpdaterGovernedProductionModeEnteredIterator struct {
	Event *FUpdaterGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *FUpdaterGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUpdaterGovernedProductionModeEntered)
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
		it.Event = new(FUpdaterGovernedProductionModeEntered)
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
func (it *FUpdaterGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUpdaterGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUpdaterGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FUpdater contract.
type FUpdaterGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FUpdater *FUpdaterFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*FUpdaterGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _FUpdater.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &FUpdaterGovernedProductionModeEnteredIterator{contract: _FUpdater.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_FUpdater *FUpdaterFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *FUpdaterGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _FUpdater.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUpdaterGovernedProductionModeEntered)
				if err := _FUpdater.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_FUpdater *FUpdaterFilterer) ParseGovernedProductionModeEntered(log types.Log) (*FUpdaterGovernedProductionModeEntered, error) {
	event := new(FUpdaterGovernedProductionModeEntered)
	if err := _FUpdater.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUpdaterTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the FUpdater contract.
type FUpdaterTimelockedGovernanceCallCanceledIterator struct {
	Event *FUpdaterTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *FUpdaterTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUpdaterTimelockedGovernanceCallCanceled)
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
		it.Event = new(FUpdaterTimelockedGovernanceCallCanceled)
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
func (it *FUpdaterTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUpdaterTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUpdaterTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FUpdater contract.
type FUpdaterTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FUpdater *FUpdaterFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*FUpdaterTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _FUpdater.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &FUpdaterTimelockedGovernanceCallCanceledIterator{contract: _FUpdater.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x7735b2391c38a81419c513e30ca578db7158eadd7101511b23e221c654d19cf8.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (_FUpdater *FUpdaterFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *FUpdaterTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _FUpdater.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUpdaterTimelockedGovernanceCallCanceled)
				if err := _FUpdater.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_FUpdater *FUpdaterFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*FUpdaterTimelockedGovernanceCallCanceled, error) {
	event := new(FUpdaterTimelockedGovernanceCallCanceled)
	if err := _FUpdater.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FUpdaterTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the FUpdater contract.
type FUpdaterTimelockedGovernanceCallExecutedIterator struct {
	Event *FUpdaterTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *FUpdaterTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FUpdaterTimelockedGovernanceCallExecuted)
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
		it.Event = new(FUpdaterTimelockedGovernanceCallExecuted)
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
func (it *FUpdaterTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FUpdaterTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FUpdaterTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FUpdater contract.
type FUpdaterTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FUpdater *FUpdaterFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*FUpdaterTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _FUpdater.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &FUpdaterTimelockedGovernanceCallExecutedIterator{contract: _FUpdater.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xa7326b57fc9cfe267aaea5e7f0b01757154d265620a0585819416ee9ddd2c438.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (_FUpdater *FUpdaterFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *FUpdaterTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _FUpdater.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FUpdaterTimelockedGovernanceCallExecuted)
				if err := _FUpdater.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_FUpdater *FUpdaterFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*FUpdaterTimelockedGovernanceCallExecuted, error) {
	event := new(FUpdaterTimelockedGovernanceCallExecuted)
	if err := _FUpdater.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
