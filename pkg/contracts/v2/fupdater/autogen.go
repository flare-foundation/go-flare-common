// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fupdater

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
var FUpdaterMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_flareDaemon\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_firstVotingRoundStartTs\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_votingEpochDurationSeconds\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_submissionWindow\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"FastUpdateFeedRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"votingRoundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes21\",\"name\":\"id\",\"type\":\"bytes21\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int8\",\"name\":\"decimals\",\"type\":\"int8\"}],\"name\":\"FastUpdateFeedReset\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"votingEpochId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"feeds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"int8[]\",\"name\":\"decimals\",\"type\":\"int8[]\"}],\"name\":\"FastUpdateFeeds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"votingRoundId\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"FastUpdateFeedsSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BLOCKS_HISTORY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_FEED_AGE_IN_VOTING_EPOCHS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNum\",\"type\":\"uint256\"}],\"name\":\"blockScoreCutoff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_cutoff\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewardEpochId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentScoreCutoff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_cutoff\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signingPolicyAddress\",\"type\":\"address\"}],\"name\":\"currentSortitionWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_weight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daemonize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fastUpdateIncentiveManager\",\"outputs\":[{\"internalType\":\"contractIIFastUpdateIncentiveManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fastUpdatesConfiguration\",\"outputs\":[{\"internalType\":\"contractIFastUpdatesConfiguration\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchAllCurrentFeeds\",\"outputs\":[{\"internalType\":\"bytes21[]\",\"name\":\"_feedIds\",\"type\":\"bytes21[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_feeds\",\"type\":\"uint256[]\"},{\"internalType\":\"int8[]\",\"name\":\"_decimals\",\"type\":\"int8[]\"},{\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indices\",\"type\":\"uint256[]\"}],\"name\":\"fetchCurrentFeeds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_feeds\",\"type\":\"uint256[]\"},{\"internalType\":\"int8[]\",\"name\":\"_decimals\",\"type\":\"int8[]\"},{\"internalType\":\"uint64\",\"name\":\"_timestamp\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstVotingRoundStartTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareDaemon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ftsoFeedPublisher\",\"outputs\":[{\"internalType\":\"contractIFtsoFeedPublisher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_historySize\",\"type\":\"uint256\"}],\"name\":\"numberOfUpdates\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_noOfUpdates\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberOfUpdatesInBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indices\",\"type\":\"uint256[]\"}],\"name\":\"removeFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_indices\",\"type\":\"uint256[]\"}],\"name\":\"resetFeeds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_submissionWindow\",\"type\":\"uint8\"}],\"name\":\"setSubmissionWindow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submissionWindow\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"sortitionBlock\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"replicate\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structBn256.G1Point\",\"name\":\"gamma\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"}],\"internalType\":\"structSortitionCredential\",\"name\":\"sortitionCredential\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"deltas\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFastUpdater.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structIFastUpdater.FastUpdates\",\"name\":\"_updates\",\"type\":\"tuple\"}],\"name\":\"submitUpdates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToFallbackMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_part1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_part2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_verificationData\",\"type\":\"bytes\"}],\"name\":\"verifyPublicKey\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistry\",\"outputs\":[{\"internalType\":\"contractIIVoterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "FUpdater",
}

// FUpdater is an auto generated Go binding around an Ethereum contract.
type FUpdater struct {
	abi abi.ABI
}

// NewFUpdater creates a new instance of FUpdater.
func NewFUpdater() *FUpdater {
	parsed, err := FUpdaterMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &FUpdater{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *FUpdater) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _governanceSettings, address _initialGovernance, address _addressUpdater, address _flareDaemon, uint32 _firstVotingRoundStartTs, uint8 _votingEpochDurationSeconds, uint8 _submissionWindow) returns()
func (fUpdater *FUpdater) PackConstructor(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _flareDaemon common.Address, _firstVotingRoundStartTs uint32, _votingEpochDurationSeconds uint8, _submissionWindow uint8) []byte {
	enc, err := fUpdater.abi.Pack("", _governanceSettings, _initialGovernance, _addressUpdater, _flareDaemon, _firstVotingRoundStartTs, _votingEpochDurationSeconds, _submissionWindow)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackMAXBLOCKSHISTORY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc1bff139.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MAX_BLOCKS_HISTORY() view returns(uint256)
func (fUpdater *FUpdater) PackMAXBLOCKSHISTORY() []byte {
	enc, err := fUpdater.abi.Pack("MAX_BLOCKS_HISTORY")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMAXBLOCKSHISTORY is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc1bff139.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MAX_BLOCKS_HISTORY() view returns(uint256)
func (fUpdater *FUpdater) TryPackMAXBLOCKSHISTORY() ([]byte, error) {
	return fUpdater.abi.Pack("MAX_BLOCKS_HISTORY")
}

// UnpackMAXBLOCKSHISTORY is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc1bff139.
//
// Solidity: function MAX_BLOCKS_HISTORY() view returns(uint256)
func (fUpdater *FUpdater) UnpackMAXBLOCKSHISTORY(data []byte) (*big.Int, error) {
	out, err := fUpdater.abi.Unpack("MAX_BLOCKS_HISTORY", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackMAXFEEDAGEINVOTINGEPOCHS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7fe3341a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function MAX_FEED_AGE_IN_VOTING_EPOCHS() view returns(uint256)
func (fUpdater *FUpdater) PackMAXFEEDAGEINVOTINGEPOCHS() []byte {
	enc, err := fUpdater.abi.Pack("MAX_FEED_AGE_IN_VOTING_EPOCHS")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackMAXFEEDAGEINVOTINGEPOCHS is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7fe3341a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function MAX_FEED_AGE_IN_VOTING_EPOCHS() view returns(uint256)
func (fUpdater *FUpdater) TryPackMAXFEEDAGEINVOTINGEPOCHS() ([]byte, error) {
	return fUpdater.abi.Pack("MAX_FEED_AGE_IN_VOTING_EPOCHS")
}

// UnpackMAXFEEDAGEINVOTINGEPOCHS is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7fe3341a.
//
// Solidity: function MAX_FEED_AGE_IN_VOTING_EPOCHS() view returns(uint256)
func (fUpdater *FUpdater) UnpackMAXFEEDAGEINVOTINGEPOCHS(data []byte) (*big.Int, error) {
	out, err := fUpdater.abi.Unpack("MAX_FEED_AGE_IN_VOTING_EPOCHS", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackBlockScoreCutoff is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdcb1476e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function blockScoreCutoff(uint256 _blockNum) view returns(uint256 _cutoff)
func (fUpdater *FUpdater) PackBlockScoreCutoff(blockNum *big.Int) []byte {
	enc, err := fUpdater.abi.Pack("blockScoreCutoff", blockNum)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackBlockScoreCutoff is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdcb1476e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function blockScoreCutoff(uint256 _blockNum) view returns(uint256 _cutoff)
func (fUpdater *FUpdater) TryPackBlockScoreCutoff(blockNum *big.Int) ([]byte, error) {
	return fUpdater.abi.Pack("blockScoreCutoff", blockNum)
}

// UnpackBlockScoreCutoff is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdcb1476e.
//
// Solidity: function blockScoreCutoff(uint256 _blockNum) view returns(uint256 _cutoff)
func (fUpdater *FUpdater) UnpackBlockScoreCutoff(data []byte) (*big.Int, error) {
	out, err := fUpdater.abi.Unpack("blockScoreCutoff", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67fc4029.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (fUpdater *FUpdater) PackCancelGovernanceCall(selector [4]byte) []byte {
	enc, err := fUpdater.abi.Pack("cancelGovernanceCall", selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67fc4029.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (fUpdater *FUpdater) TryPackCancelGovernanceCall(selector [4]byte) ([]byte, error) {
	return fUpdater.abi.Pack("cancelGovernanceCall", selector)
}

// PackCurrentRewardEpochId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8e0e9f7c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function currentRewardEpochId() view returns(uint24)
func (fUpdater *FUpdater) PackCurrentRewardEpochId() []byte {
	enc, err := fUpdater.abi.Pack("currentRewardEpochId")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCurrentRewardEpochId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8e0e9f7c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function currentRewardEpochId() view returns(uint24)
func (fUpdater *FUpdater) TryPackCurrentRewardEpochId() ([]byte, error) {
	return fUpdater.abi.Pack("currentRewardEpochId")
}

// UnpackCurrentRewardEpochId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8e0e9f7c.
//
// Solidity: function currentRewardEpochId() view returns(uint24)
func (fUpdater *FUpdater) UnpackCurrentRewardEpochId(data []byte) (*big.Int, error) {
	out, err := fUpdater.abi.Unpack("currentRewardEpochId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCurrentScoreCutoff is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0799fe75.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function currentScoreCutoff() view returns(uint256 _cutoff)
func (fUpdater *FUpdater) PackCurrentScoreCutoff() []byte {
	enc, err := fUpdater.abi.Pack("currentScoreCutoff")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCurrentScoreCutoff is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0799fe75.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function currentScoreCutoff() view returns(uint256 _cutoff)
func (fUpdater *FUpdater) TryPackCurrentScoreCutoff() ([]byte, error) {
	return fUpdater.abi.Pack("currentScoreCutoff")
}

// UnpackCurrentScoreCutoff is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0799fe75.
//
// Solidity: function currentScoreCutoff() view returns(uint256 _cutoff)
func (fUpdater *FUpdater) UnpackCurrentScoreCutoff(data []byte) (*big.Int, error) {
	out, err := fUpdater.abi.Unpack("currentScoreCutoff", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackCurrentSortitionWeight is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa14634a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function currentSortitionWeight(address _signingPolicyAddress) view returns(uint256 _weight)
func (fUpdater *FUpdater) PackCurrentSortitionWeight(signingPolicyAddress common.Address) []byte {
	enc, err := fUpdater.abi.Pack("currentSortitionWeight", signingPolicyAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCurrentSortitionWeight is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa14634a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function currentSortitionWeight(address _signingPolicyAddress) view returns(uint256 _weight)
func (fUpdater *FUpdater) TryPackCurrentSortitionWeight(signingPolicyAddress common.Address) ([]byte, error) {
	return fUpdater.abi.Pack("currentSortitionWeight", signingPolicyAddress)
}

// UnpackCurrentSortitionWeight is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa14634a7.
//
// Solidity: function currentSortitionWeight(address _signingPolicyAddress) view returns(uint256 _weight)
func (fUpdater *FUpdater) UnpackCurrentSortitionWeight(data []byte) (*big.Int, error) {
	out, err := fUpdater.abi.Unpack("currentSortitionWeight", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackDaemonize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d0e8c34.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function daemonize() returns(bool)
func (fUpdater *FUpdater) PackDaemonize() []byte {
	enc, err := fUpdater.abi.Pack("daemonize")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDaemonize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d0e8c34.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function daemonize() returns(bool)
func (fUpdater *FUpdater) TryPackDaemonize() ([]byte, error) {
	return fUpdater.abi.Pack("daemonize")
}

// UnpackDaemonize is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (fUpdater *FUpdater) UnpackDaemonize(data []byte) (bool, error) {
	out, err := fUpdater.abi.Unpack("daemonize", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ff27079.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (fUpdater *FUpdater) PackExecuteGovernanceCall(selector [4]byte) []byte {
	enc, err := fUpdater.abi.Pack("executeGovernanceCall", selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5ff27079.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function executeGovernanceCall(bytes4 _selector) returns()
func (fUpdater *FUpdater) TryPackExecuteGovernanceCall(selector [4]byte) ([]byte, error) {
	return fUpdater.abi.Pack("executeGovernanceCall", selector)
}

// PackFastUpdateIncentiveManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7925eaca.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fastUpdateIncentiveManager() view returns(address)
func (fUpdater *FUpdater) PackFastUpdateIncentiveManager() []byte {
	enc, err := fUpdater.abi.Pack("fastUpdateIncentiveManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFastUpdateIncentiveManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7925eaca.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fastUpdateIncentiveManager() view returns(address)
func (fUpdater *FUpdater) TryPackFastUpdateIncentiveManager() ([]byte, error) {
	return fUpdater.abi.Pack("fastUpdateIncentiveManager")
}

// UnpackFastUpdateIncentiveManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7925eaca.
//
// Solidity: function fastUpdateIncentiveManager() view returns(address)
func (fUpdater *FUpdater) UnpackFastUpdateIncentiveManager(data []byte) (common.Address, error) {
	out, err := fUpdater.abi.Unpack("fastUpdateIncentiveManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFastUpdatesConfiguration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc10f489a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (fUpdater *FUpdater) PackFastUpdatesConfiguration() []byte {
	enc, err := fUpdater.abi.Pack("fastUpdatesConfiguration")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFastUpdatesConfiguration is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc10f489a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (fUpdater *FUpdater) TryPackFastUpdatesConfiguration() ([]byte, error) {
	return fUpdater.abi.Pack("fastUpdatesConfiguration")
}

// UnpackFastUpdatesConfiguration is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc10f489a.
//
// Solidity: function fastUpdatesConfiguration() view returns(address)
func (fUpdater *FUpdater) UnpackFastUpdatesConfiguration(data []byte) (common.Address, error) {
	out, err := fUpdater.abi.Unpack("fastUpdatesConfiguration", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFetchAllCurrentFeeds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4691377f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fetchAllCurrentFeeds() view returns(bytes21[] _feedIds, uint256[] _feeds, int8[] _decimals, uint64 _timestamp)
func (fUpdater *FUpdater) PackFetchAllCurrentFeeds() []byte {
	enc, err := fUpdater.abi.Pack("fetchAllCurrentFeeds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFetchAllCurrentFeeds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4691377f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fetchAllCurrentFeeds() view returns(bytes21[] _feedIds, uint256[] _feeds, int8[] _decimals, uint64 _timestamp)
func (fUpdater *FUpdater) TryPackFetchAllCurrentFeeds() ([]byte, error) {
	return fUpdater.abi.Pack("fetchAllCurrentFeeds")
}

// FetchAllCurrentFeedsOutput serves as a container for the return parameters of contract
// method FetchAllCurrentFeeds.
type FetchAllCurrentFeedsOutput struct {
	FeedIds   [][21]byte
	Feeds     []*big.Int
	Decimals  []int8
	Timestamp uint64
}

// UnpackFetchAllCurrentFeeds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4691377f.
//
// Solidity: function fetchAllCurrentFeeds() view returns(bytes21[] _feedIds, uint256[] _feeds, int8[] _decimals, uint64 _timestamp)
func (fUpdater *FUpdater) UnpackFetchAllCurrentFeeds(data []byte) (FetchAllCurrentFeedsOutput, error) {
	out, err := fUpdater.abi.Unpack("fetchAllCurrentFeeds", data)
	outstruct := new(FetchAllCurrentFeedsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.FeedIds = *abi.ConvertType(out[0], new([][21]byte)).(*[][21]byte)
	outstruct.Feeds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Decimals = *abi.ConvertType(out[2], new([]int8)).(*[]int8)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackFetchCurrentFeeds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x45a15d3c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fetchCurrentFeeds(uint256[] _indices) view returns(uint256[] _feeds, int8[] _decimals, uint64 _timestamp)
func (fUpdater *FUpdater) PackFetchCurrentFeeds(indices []*big.Int) []byte {
	enc, err := fUpdater.abi.Pack("fetchCurrentFeeds", indices)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFetchCurrentFeeds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x45a15d3c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fetchCurrentFeeds(uint256[] _indices) view returns(uint256[] _feeds, int8[] _decimals, uint64 _timestamp)
func (fUpdater *FUpdater) TryPackFetchCurrentFeeds(indices []*big.Int) ([]byte, error) {
	return fUpdater.abi.Pack("fetchCurrentFeeds", indices)
}

// FetchCurrentFeedsOutput serves as a container for the return parameters of contract
// method FetchCurrentFeeds.
type FetchCurrentFeedsOutput struct {
	Feeds     []*big.Int
	Decimals  []int8
	Timestamp uint64
}

// UnpackFetchCurrentFeeds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x45a15d3c.
//
// Solidity: function fetchCurrentFeeds(uint256[] _indices) view returns(uint256[] _feeds, int8[] _decimals, uint64 _timestamp)
func (fUpdater *FUpdater) UnpackFetchCurrentFeeds(data []byte) (FetchCurrentFeedsOutput, error) {
	out, err := fUpdater.abi.Unpack("fetchCurrentFeeds", data)
	outstruct := new(FetchCurrentFeedsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Feeds = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Decimals = *abi.ConvertType(out[1], new([]int8)).(*[]int8)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackFirstVotingRoundStartTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe8d0e70a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (fUpdater *FUpdater) PackFirstVotingRoundStartTs() []byte {
	enc, err := fUpdater.abi.Pack("firstVotingRoundStartTs")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFirstVotingRoundStartTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe8d0e70a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (fUpdater *FUpdater) TryPackFirstVotingRoundStartTs() ([]byte, error) {
	return fUpdater.abi.Pack("firstVotingRoundStartTs")
}

// UnpackFirstVotingRoundStartTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (fUpdater *FUpdater) UnpackFirstVotingRoundStartTs(data []byte) (uint64, error) {
	out, err := fUpdater.abi.Unpack("firstVotingRoundStartTs", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackFlareDaemon is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1077532.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareDaemon() view returns(address)
func (fUpdater *FUpdater) PackFlareDaemon() []byte {
	enc, err := fUpdater.abi.Pack("flareDaemon")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFlareDaemon is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1077532.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function flareDaemon() view returns(address)
func (fUpdater *FUpdater) TryPackFlareDaemon() ([]byte, error) {
	return fUpdater.abi.Pack("flareDaemon")
}

// UnpackFlareDaemon is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (fUpdater *FUpdater) UnpackFlareDaemon(data []byte) (common.Address, error) {
	out, err := fUpdater.abi.Unpack("flareDaemon", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFlareSystemsManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaae7fc9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareSystemsManager() view returns(address)
func (fUpdater *FUpdater) PackFlareSystemsManager() []byte {
	enc, err := fUpdater.abi.Pack("flareSystemsManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFlareSystemsManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaae7fc9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function flareSystemsManager() view returns(address)
func (fUpdater *FUpdater) TryPackFlareSystemsManager() ([]byte, error) {
	return fUpdater.abi.Pack("flareSystemsManager")
}

// UnpackFlareSystemsManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (fUpdater *FUpdater) UnpackFlareSystemsManager(data []byte) (common.Address, error) {
	out, err := fUpdater.abi.Unpack("flareSystemsManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFtsoFeedPublisher is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x29bfe39d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function ftsoFeedPublisher() view returns(address)
func (fUpdater *FUpdater) PackFtsoFeedPublisher() []byte {
	enc, err := fUpdater.abi.Pack("ftsoFeedPublisher")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFtsoFeedPublisher is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x29bfe39d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function ftsoFeedPublisher() view returns(address)
func (fUpdater *FUpdater) TryPackFtsoFeedPublisher() ([]byte, error) {
	return fUpdater.abi.Pack("ftsoFeedPublisher")
}

// UnpackFtsoFeedPublisher is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x29bfe39d.
//
// Solidity: function ftsoFeedPublisher() view returns(address)
func (fUpdater *FUpdater) UnpackFtsoFeedPublisher(data []byte) (common.Address, error) {
	out, err := fUpdater.abi.Unpack("ftsoFeedPublisher", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (fUpdater *FUpdater) PackGetAddressUpdater() []byte {
	enc, err := fUpdater.abi.Pack("getAddressUpdater")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (fUpdater *FUpdater) TryPackGetAddressUpdater() ([]byte, error) {
	return fUpdater.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (fUpdater *FUpdater) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := fUpdater.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetContractName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5f5ba72.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getContractName() pure returns(string)
func (fUpdater *FUpdater) PackGetContractName() []byte {
	enc, err := fUpdater.abi.Pack("getContractName")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetContractName is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5f5ba72.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getContractName() pure returns(string)
func (fUpdater *FUpdater) TryPackGetContractName() ([]byte, error) {
	return fUpdater.abi.Pack("getContractName")
}

// UnpackGetContractName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (fUpdater *FUpdater) UnpackGetContractName(data []byte) (string, error) {
	out, err := fUpdater.abi.Unpack("getContractName", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governance() view returns(address)
func (fUpdater *FUpdater) PackGovernance() []byte {
	enc, err := fUpdater.abi.Pack("governance")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function governance() view returns(address)
func (fUpdater *FUpdater) TryPackGovernance() ([]byte, error) {
	return fUpdater.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (fUpdater *FUpdater) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := fUpdater.abi.Unpack("governance", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGovernanceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62354e03.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governanceSettings() view returns(address)
func (fUpdater *FUpdater) PackGovernanceSettings() []byte {
	enc, err := fUpdater.abi.Pack("governanceSettings")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGovernanceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62354e03.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function governanceSettings() view returns(address)
func (fUpdater *FUpdater) TryPackGovernanceSettings() ([]byte, error) {
	return fUpdater.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (fUpdater *FUpdater) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := fUpdater.abi.Unpack("governanceSettings", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitialise is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef88bf13.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (fUpdater *FUpdater) PackInitialise(governanceSettings common.Address, initialGovernance common.Address) []byte {
	enc, err := fUpdater.abi.Pack("initialise", governanceSettings, initialGovernance)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialise is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef88bf13.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (fUpdater *FUpdater) TryPackInitialise(governanceSettings common.Address, initialGovernance common.Address) ([]byte, error) {
	return fUpdater.abi.Pack("initialise", governanceSettings, initialGovernance)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (fUpdater *FUpdater) PackIsExecutor(address common.Address) []byte {
	enc, err := fUpdater.abi.Pack("isExecutor", address)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (fUpdater *FUpdater) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return fUpdater.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (fUpdater *FUpdater) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := fUpdater.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackNumberOfUpdates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe36da7b7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function numberOfUpdates(uint256 _historySize) view returns(uint256[] _noOfUpdates)
func (fUpdater *FUpdater) PackNumberOfUpdates(historySize *big.Int) []byte {
	enc, err := fUpdater.abi.Pack("numberOfUpdates", historySize)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNumberOfUpdates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe36da7b7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function numberOfUpdates(uint256 _historySize) view returns(uint256[] _noOfUpdates)
func (fUpdater *FUpdater) TryPackNumberOfUpdates(historySize *big.Int) ([]byte, error) {
	return fUpdater.abi.Pack("numberOfUpdates", historySize)
}

// UnpackNumberOfUpdates is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe36da7b7.
//
// Solidity: function numberOfUpdates(uint256 _historySize) view returns(uint256[] _noOfUpdates)
func (fUpdater *FUpdater) UnpackNumberOfUpdates(data []byte) ([]*big.Int, error) {
	out, err := fUpdater.abi.Unpack("numberOfUpdates", data)
	if err != nil {
		return *new([]*big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	return out0, nil
}

// PackNumberOfUpdatesInBlock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc79c300.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function numberOfUpdatesInBlock(uint256 _blockNumber) view returns(uint256)
func (fUpdater *FUpdater) PackNumberOfUpdatesInBlock(blockNumber *big.Int) []byte {
	enc, err := fUpdater.abi.Pack("numberOfUpdatesInBlock", blockNumber)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNumberOfUpdatesInBlock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfc79c300.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function numberOfUpdatesInBlock(uint256 _blockNumber) view returns(uint256)
func (fUpdater *FUpdater) TryPackNumberOfUpdatesInBlock(blockNumber *big.Int) ([]byte, error) {
	return fUpdater.abi.Pack("numberOfUpdatesInBlock", blockNumber)
}

// UnpackNumberOfUpdatesInBlock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfc79c300.
//
// Solidity: function numberOfUpdatesInBlock(uint256 _blockNumber) view returns(uint256)
func (fUpdater *FUpdater) UnpackNumberOfUpdatesInBlock(data []byte) (*big.Int, error) {
	out, err := fUpdater.abi.Unpack("numberOfUpdatesInBlock", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (fUpdater *FUpdater) PackProductionMode() []byte {
	enc, err := fUpdater.abi.Pack("productionMode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function productionMode() view returns(bool)
func (fUpdater *FUpdater) TryPackProductionMode() ([]byte, error) {
	return fUpdater.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (fUpdater *FUpdater) UnpackProductionMode(data []byte) (bool, error) {
	out, err := fUpdater.abi.Unpack("productionMode", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRemoveFeeds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xabfaf170.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeFeeds(uint256[] _indices) returns()
func (fUpdater *FUpdater) PackRemoveFeeds(indices []*big.Int) []byte {
	enc, err := fUpdater.abi.Pack("removeFeeds", indices)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveFeeds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xabfaf170.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeFeeds(uint256[] _indices) returns()
func (fUpdater *FUpdater) TryPackRemoveFeeds(indices []*big.Int) ([]byte, error) {
	return fUpdater.abi.Pack("removeFeeds", indices)
}

// PackResetFeeds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x63f921db.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function resetFeeds(uint256[] _indices) returns()
func (fUpdater *FUpdater) PackResetFeeds(indices []*big.Int) []byte {
	enc, err := fUpdater.abi.Pack("resetFeeds", indices)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackResetFeeds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x63f921db.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function resetFeeds(uint256[] _indices) returns()
func (fUpdater *FUpdater) TryPackResetFeeds(indices []*big.Int) ([]byte, error) {
	return fUpdater.abi.Pack("resetFeeds", indices)
}

// PackSetSubmissionWindow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a166051.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSubmissionWindow(uint8 _submissionWindow) returns()
func (fUpdater *FUpdater) PackSetSubmissionWindow(submissionWindow uint8) []byte {
	enc, err := fUpdater.abi.Pack("setSubmissionWindow", submissionWindow)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSubmissionWindow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0a166051.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSubmissionWindow(uint8 _submissionWindow) returns()
func (fUpdater *FUpdater) TryPackSetSubmissionWindow(submissionWindow uint8) ([]byte, error) {
	return fUpdater.abi.Pack("setSubmissionWindow", submissionWindow)
}

// PackSubmissionWindow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe621dbc7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submissionWindow() view returns(uint8)
func (fUpdater *FUpdater) PackSubmissionWindow() []byte {
	enc, err := fUpdater.abi.Pack("submissionWindow")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmissionWindow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe621dbc7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submissionWindow() view returns(uint8)
func (fUpdater *FUpdater) TryPackSubmissionWindow() ([]byte, error) {
	return fUpdater.abi.Pack("submissionWindow")
}

// UnpackSubmissionWindow is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe621dbc7.
//
// Solidity: function submissionWindow() view returns(uint8)
func (fUpdater *FUpdater) UnpackSubmissionWindow(data []byte) (uint8, error) {
	out, err := fUpdater.abi.Unpack("submissionWindow", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackSubmitUpdates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x470e91df.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submitUpdates((uint256,(uint256,(uint256,uint256),uint256,uint256),bytes,(uint8,bytes32,bytes32)) _updates) returns()
func (fUpdater *FUpdater) PackSubmitUpdates(updates IFastUpdaterFastUpdates) []byte {
	enc, err := fUpdater.abi.Pack("submitUpdates", updates)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmitUpdates is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x470e91df.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submitUpdates((uint256,(uint256,(uint256,uint256),uint256,uint256),bytes,(uint8,bytes32,bytes32)) _updates) returns()
func (fUpdater *FUpdater) TryPackSubmitUpdates(updates IFastUpdaterFastUpdates) ([]byte, error) {
	return fUpdater.abi.Pack("submitUpdates", updates)
}

// PackSwitchToFallbackMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe22fdece.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (fUpdater *FUpdater) PackSwitchToFallbackMode() []byte {
	enc, err := fUpdater.abi.Pack("switchToFallbackMode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwitchToFallbackMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe22fdece.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (fUpdater *FUpdater) TryPackSwitchToFallbackMode() ([]byte, error) {
	return fUpdater.abi.Pack("switchToFallbackMode")
}

// UnpackSwitchToFallbackMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe22fdece.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (fUpdater *FUpdater) UnpackSwitchToFallbackMode(data []byte) (bool, error) {
	out, err := fUpdater.abi.Unpack("switchToFallbackMode", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (fUpdater *FUpdater) PackSwitchToProductionMode() []byte {
	enc, err := fUpdater.abi.Pack("switchToProductionMode")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function switchToProductionMode() returns()
func (fUpdater *FUpdater) TryPackSwitchToProductionMode() ([]byte, error) {
	return fUpdater.abi.Pack("switchToProductionMode")
}

// PackTimelockedCalls is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74e6310e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (fUpdater *FUpdater) PackTimelockedCalls(selector [4]byte) []byte {
	enc, err := fUpdater.abi.Pack("timelockedCalls", selector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTimelockedCalls is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74e6310e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (fUpdater *FUpdater) TryPackTimelockedCalls(selector [4]byte) ([]byte, error) {
	return fUpdater.abi.Pack("timelockedCalls", selector)
}

// TimelockedCallsOutput serves as a container for the return parameters of contract
// method TimelockedCalls.
type TimelockedCallsOutput struct {
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
}

// UnpackTimelockedCalls is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x74e6310e.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (fUpdater *FUpdater) UnpackTimelockedCalls(data []byte) (TimelockedCallsOutput, error) {
	out, err := fUpdater.abi.Unpack("timelockedCalls", data)
	outstruct := new(TimelockedCallsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AllowedAfterTimestamp = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	return *outstruct, nil
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (fUpdater *FUpdater) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := fUpdater.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (fUpdater *FUpdater) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return fUpdater.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackVerifyPublicKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70473f2f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function verifyPublicKey(address _voter, bytes32 _part1, bytes32 _part2, bytes _verificationData) view returns()
func (fUpdater *FUpdater) PackVerifyPublicKey(voter common.Address, part1 [32]byte, part2 [32]byte, verificationData []byte) []byte {
	enc, err := fUpdater.abi.Pack("verifyPublicKey", voter, part1, part2, verificationData)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVerifyPublicKey is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70473f2f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function verifyPublicKey(address _voter, bytes32 _part1, bytes32 _part2, bytes _verificationData) view returns()
func (fUpdater *FUpdater) TryPackVerifyPublicKey(voter common.Address, part1 [32]byte, part2 [32]byte, verificationData []byte) ([]byte, error) {
	return fUpdater.abi.Pack("verifyPublicKey", voter, part1, part2, verificationData)
}

// PackVoterRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbe60040e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function voterRegistry() view returns(address)
func (fUpdater *FUpdater) PackVoterRegistry() []byte {
	enc, err := fUpdater.abi.Pack("voterRegistry")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVoterRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbe60040e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function voterRegistry() view returns(address)
func (fUpdater *FUpdater) TryPackVoterRegistry() ([]byte, error) {
	return fUpdater.abi.Pack("voterRegistry")
}

// UnpackVoterRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (fUpdater *FUpdater) UnpackVoterRegistry(data []byte) (common.Address, error) {
	out, err := fUpdater.abi.Unpack("voterRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackVotingEpochDurationSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5a832088.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (fUpdater *FUpdater) PackVotingEpochDurationSeconds() []byte {
	enc, err := fUpdater.abi.Pack("votingEpochDurationSeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVotingEpochDurationSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5a832088.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (fUpdater *FUpdater) TryPackVotingEpochDurationSeconds() ([]byte, error) {
	return fUpdater.abi.Pack("votingEpochDurationSeconds")
}

// UnpackVotingEpochDurationSeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (fUpdater *FUpdater) UnpackVotingEpochDurationSeconds(data []byte) (uint64, error) {
	out, err := fUpdater.abi.Unpack("votingEpochDurationSeconds", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// FUpdaterFastUpdateFeedRemoved represents a FastUpdateFeedRemoved event raised by the FUpdater contract.
type FUpdaterFastUpdateFeedRemoved struct {
	Index *big.Int
	Raw   *types.Log // Blockchain specific contextual infos
}

const FUpdaterFastUpdateFeedRemovedEventName = "FastUpdateFeedRemoved"

// ContractEventName returns the user-defined event name.
func (FUpdaterFastUpdateFeedRemoved) ContractEventName() string {
	return FUpdaterFastUpdateFeedRemovedEventName
}

// UnpackFastUpdateFeedRemovedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FastUpdateFeedRemoved(uint256 indexed index)
func (fUpdater *FUpdater) UnpackFastUpdateFeedRemovedEvent(log *types.Log) (*FUpdaterFastUpdateFeedRemoved, error) {
	event := "FastUpdateFeedRemoved"
	if len(log.Topics) == 0 || log.Topics[0] != fUpdater.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUpdaterFastUpdateFeedRemoved)
	if len(log.Data) > 0 {
		if err := fUpdater.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUpdater.abi.Events[event].Inputs {
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

// FUpdaterFastUpdateFeedReset represents a FastUpdateFeedReset event raised by the FUpdater contract.
type FUpdaterFastUpdateFeedReset struct {
	VotingRoundId *big.Int
	Index         *big.Int
	Id            [21]byte
	Value         *big.Int
	Decimals      int8
	Raw           *types.Log // Blockchain specific contextual infos
}

const FUpdaterFastUpdateFeedResetEventName = "FastUpdateFeedReset"

// ContractEventName returns the user-defined event name.
func (FUpdaterFastUpdateFeedReset) ContractEventName() string {
	return FUpdaterFastUpdateFeedResetEventName
}

// UnpackFastUpdateFeedResetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FastUpdateFeedReset(uint256 indexed votingRoundId, uint256 indexed index, bytes21 indexed id, uint256 value, int8 decimals)
func (fUpdater *FUpdater) UnpackFastUpdateFeedResetEvent(log *types.Log) (*FUpdaterFastUpdateFeedReset, error) {
	event := "FastUpdateFeedReset"
	if len(log.Topics) == 0 || log.Topics[0] != fUpdater.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUpdaterFastUpdateFeedReset)
	if len(log.Data) > 0 {
		if err := fUpdater.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUpdater.abi.Events[event].Inputs {
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

// FUpdaterFastUpdateFeeds represents a FastUpdateFeeds event raised by the FUpdater contract.
type FUpdaterFastUpdateFeeds struct {
	VotingEpochId *big.Int
	Feeds         []*big.Int
	Decimals      []int8
	Raw           *types.Log // Blockchain specific contextual infos
}

const FUpdaterFastUpdateFeedsEventName = "FastUpdateFeeds"

// ContractEventName returns the user-defined event name.
func (FUpdaterFastUpdateFeeds) ContractEventName() string {
	return FUpdaterFastUpdateFeedsEventName
}

// UnpackFastUpdateFeedsEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FastUpdateFeeds(uint256 indexed votingEpochId, uint256[] feeds, int8[] decimals)
func (fUpdater *FUpdater) UnpackFastUpdateFeedsEvent(log *types.Log) (*FUpdaterFastUpdateFeeds, error) {
	event := "FastUpdateFeeds"
	if len(log.Topics) == 0 || log.Topics[0] != fUpdater.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUpdaterFastUpdateFeeds)
	if len(log.Data) > 0 {
		if err := fUpdater.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUpdater.abi.Events[event].Inputs {
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

// FUpdaterFastUpdateFeedsSubmitted represents a FastUpdateFeedsSubmitted event raised by the FUpdater contract.
type FUpdaterFastUpdateFeedsSubmitted struct {
	VotingRoundId        uint32
	SigningPolicyAddress common.Address
	Raw                  *types.Log // Blockchain specific contextual infos
}

const FUpdaterFastUpdateFeedsSubmittedEventName = "FastUpdateFeedsSubmitted"

// ContractEventName returns the user-defined event name.
func (FUpdaterFastUpdateFeedsSubmitted) ContractEventName() string {
	return FUpdaterFastUpdateFeedsSubmittedEventName
}

// UnpackFastUpdateFeedsSubmittedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FastUpdateFeedsSubmitted(uint32 indexed votingRoundId, address indexed signingPolicyAddress)
func (fUpdater *FUpdater) UnpackFastUpdateFeedsSubmittedEvent(log *types.Log) (*FUpdaterFastUpdateFeedsSubmitted, error) {
	event := "FastUpdateFeedsSubmitted"
	if len(log.Topics) == 0 || log.Topics[0] != fUpdater.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUpdaterFastUpdateFeedsSubmitted)
	if len(log.Data) > 0 {
		if err := fUpdater.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUpdater.abi.Events[event].Inputs {
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

// FUpdaterGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FUpdater contract.
type FUpdaterGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   *types.Log // Blockchain specific contextual infos
}

const FUpdaterGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (FUpdaterGovernanceCallTimelocked) ContractEventName() string {
	return FUpdaterGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (fUpdater *FUpdater) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*FUpdaterGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != fUpdater.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUpdaterGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := fUpdater.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUpdater.abi.Events[event].Inputs {
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

// FUpdaterGovernanceInitialised represents a GovernanceInitialised event raised by the FUpdater contract.
type FUpdaterGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const FUpdaterGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (FUpdaterGovernanceInitialised) ContractEventName() string {
	return FUpdaterGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (fUpdater *FUpdater) UnpackGovernanceInitialisedEvent(log *types.Log) (*FUpdaterGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != fUpdater.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUpdaterGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := fUpdater.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUpdater.abi.Events[event].Inputs {
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

// FUpdaterGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FUpdater contract.
type FUpdaterGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const FUpdaterGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (FUpdaterGovernedProductionModeEntered) ContractEventName() string {
	return FUpdaterGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (fUpdater *FUpdater) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*FUpdaterGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != fUpdater.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUpdaterGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := fUpdater.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUpdater.abi.Events[event].Inputs {
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

// FUpdaterTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FUpdater contract.
type FUpdaterTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const FUpdaterTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (FUpdaterTimelockedGovernanceCallCanceled) ContractEventName() string {
	return FUpdaterTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (fUpdater *FUpdater) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*FUpdaterTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != fUpdater.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUpdaterTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := fUpdater.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUpdater.abi.Events[event].Inputs {
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

// FUpdaterTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FUpdater contract.
type FUpdaterTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const FUpdaterTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (FUpdaterTimelockedGovernanceCallExecuted) ContractEventName() string {
	return FUpdaterTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (fUpdater *FUpdater) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*FUpdaterTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != fUpdater.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FUpdaterTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := fUpdater.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range fUpdater.abi.Events[event].Inputs {
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
func (fUpdater *FUpdater) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], fUpdater.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return fUpdater.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], fUpdater.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return fUpdater.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], fUpdater.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return fUpdater.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// FUpdaterECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the FUpdater contract.
type FUpdaterECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func FUpdaterECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (fUpdater *FUpdater) UnpackECDSAInvalidSignatureError(raw []byte) (*FUpdaterECDSAInvalidSignature, error) {
	out := new(FUpdaterECDSAInvalidSignature)
	if err := fUpdater.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// FUpdaterECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the FUpdater contract.
type FUpdaterECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func FUpdaterECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (fUpdater *FUpdater) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*FUpdaterECDSAInvalidSignatureLength, error) {
	out := new(FUpdaterECDSAInvalidSignatureLength)
	if err := fUpdater.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// FUpdaterECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the FUpdater contract.
type FUpdaterECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func FUpdaterECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (fUpdater *FUpdater) UnpackECDSAInvalidSignatureSError(raw []byte) (*FUpdaterECDSAInvalidSignatureS, error) {
	out := new(FUpdaterECDSAInvalidSignatureS)
	if err := fUpdater.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}
