// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package system

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

// FlareSystemsManagerInitialSettings is an auto generated low-level Go binding around an user-defined struct.
type FlareSystemsManagerInitialSettings struct {
	InitialRandomVotePowerBlockSelectionSize uint16
	InitialRewardEpochId                     *big.Int
	InitialRewardEpochThreshold              uint16
}

// FlareSystemsManagerSettings is an auto generated low-level Go binding around an user-defined struct.
type FlareSystemsManagerSettings struct {
	RandomAcquisitionMaxDurationSeconds          uint16
	RandomAcquisitionMaxDurationBlocks           uint16
	NewSigningPolicyInitializationStartSeconds   uint16
	NewSigningPolicyMinNumberOfVotingRoundsDelay uint8
	VoterRegistrationMinDurationSeconds          uint16
	VoterRegistrationMinDurationBlocks           uint16
	SubmitUptimeVoteMinDurationSeconds           uint16
	SubmitUptimeVoteMinDurationBlocks            uint16
	SigningPolicyThresholdPPM                    *big.Int
	SigningPolicyMinNumberOfVoters               uint16
	RewardExpiryOffsetSeconds                    uint32
}

// IFlareSystemsManagerNumberOfWeightBasedClaims is an auto generated low-level Go binding around an user-defined struct.
type IFlareSystemsManagerNumberOfWeightBasedClaims struct {
	RewardManagerId       *big.Int
	NoOfWeightBasedClaims *big.Int
}

// IFlareSystemsManagerSignature is an auto generated low-level Go binding around an user-defined struct.
type IFlareSystemsManagerSignature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// FlareSystemsManagerMetaData contains all meta data concerning the FlareSystemsManager contract.
var FlareSystemsManagerMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_flareDaemon\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"randomAcquisitionMaxDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"randomAcquisitionMaxDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newSigningPolicyInitializationStartSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"newSigningPolicyMinNumberOfVotingRoundsDelay\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"voterRegistrationMinDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"voterRegistrationMinDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"submitUptimeVoteMinDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"submitUptimeVoteMinDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"signingPolicyThresholdPPM\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"signingPolicyMinNumberOfVoters\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"rewardExpiryOffsetSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structFlareSystemsManager.Settings\",\"name\":\"_settings\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"_firstVotingRoundStartTs\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"_votingEpochDurationSeconds\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"_firstRewardEpochStartVotingRoundId\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"_rewardEpochDurationInVotingEpochs\",\"type\":\"uint16\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"initialRandomVotePowerBlockSelectionSize\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"initialRewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"initialRewardEpochThreshold\",\"type\":\"uint16\"}],\"internalType\":\"structFlareSystemsManager.InitialSettings\",\"name\":\"_initialSettings\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"ClosingExpiredRewardEpochFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"RandomAcquisitionStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startVotingRoundId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"RewardEpochStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rewardsHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rewardManagerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noOfWeightBasedClaims\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structIFlareSystemsManager.NumberOfWeightBasedClaims[]\",\"name\":\"noOfWeightBasedClaims\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"RewardsSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"name\":\"SettingCleanUpBlockNumberFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"SignUptimeVoteEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"SigningPolicySigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"TriggeringVoterRegistrationFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"uptimeVoteHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"thresholdReached\",\"type\":\"bool\"}],\"name\":\"UptimeVoteSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signingPolicyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes20[]\",\"name\":\"nodeIds\",\"type\":\"bytes20[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"UptimeVoteSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"votePowerBlock\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"name\":\"VotePowerBlockSelected\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cleanupBlockNumberManager\",\"outputs\":[{\"internalType\":\"contractIICleanupBlockNumberManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentRewardEpochExpectedEndTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daemonize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstRewardEpochStartTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstVotingRoundStartTs\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareDaemon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRewardEpoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentRewardEpochId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentVotingEpochId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getRandomAcquisitionInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_randomAcquisitionStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_randomAcquisitionStartBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_randomAcquisitionEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_randomAcquisitionEndBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getRewardEpochStartInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_rewardEpochStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardEpochStartBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardEpochSwitchoverTriggerContracts\",\"outputs\":[{\"internalType\":\"contractIIRewardEpochSwitchoverTrigger[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getRewardsSignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_rewardsSignStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardsSignStartBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardsSignEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardsSignEndBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getSeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getSigningPolicySignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignStartBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignEndTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignEndBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getStartVotingRoundId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getThreshold\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"getUptimeVoteSignStartInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSignStartTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSignStartBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getVotePowerBlock\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_votePowerBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"getVoterRegistrationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_votePowerBlock\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterRewardsSignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_rewardsSignTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_rewardsSignBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterSigningPolicySignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_signingPolicySignBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterUptimeVoteSignInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSignTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSignBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVoterUptimeVoteSubmitInfo\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSubmitTs\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_uptimeVoteSubmitBlock\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialRandomVotePowerBlockSelectionSize\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"}],\"name\":\"initialise\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isVoterRegistrationEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastInitializedVotingRoundId\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newSigningPolicyInitializationStartSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newSigningPolicyMinNumberOfVotingRoundsDelay\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardManagerId\",\"type\":\"uint256\"}],\"name\":\"noOfWeightBasedClaims\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"noOfWeightBasedClaimsHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomAcquisitionMaxDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"randomAcquisitionMaxDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relay\",\"outputs\":[{\"internalType\":\"contractIIRelay\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardEpochIdToExpireNext\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardExpiryOffsetSeconds\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"rewardsHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIRewardEpochSwitchoverTrigger[]\",\"name\":\"_contracts\",\"type\":\"address[]\"}],\"name\":\"setRewardEpochSwitchoverTriggerContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rewardManagerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noOfWeightBasedClaims\",\"type\":\"uint256\"}],\"internalType\":\"structIFlareSystemsManager.NumberOfWeightBasedClaims[]\",\"name\":\"_noOfWeightBasedClaims\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_rewardsHash\",\"type\":\"bytes32\"}],\"name\":\"setRewardsData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_submit3Aligned\",\"type\":\"bool\"}],\"name\":\"setSubmit3Aligned\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_triggerExpirationAndCleanup\",\"type\":\"bool\"}],\"name\":\"setTriggerExpirationAndCleanup\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIIVoterRegistrationTrigger\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setVoterRegistrationTriggerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_newSigningPolicyHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemsManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signNewSigningPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"rewardManagerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noOfWeightBasedClaims\",\"type\":\"uint256\"}],\"internalType\":\"structIFlareSystemsManager.NumberOfWeightBasedClaims[]\",\"name\":\"_noOfWeightBasedClaims\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_rewardsHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemsManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes32\",\"name\":\"_uptimeVoteHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemsManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"signUptimeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicyMinNumberOfVoters\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signingPolicyThresholdPPM\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submission\",\"outputs\":[{\"internalType\":\"contractIISubmission\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submit3Aligned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"_rewardEpochId\",\"type\":\"uint24\"},{\"internalType\":\"bytes20[]\",\"name\":\"_nodeIds\",\"type\":\"bytes20[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structIFlareSystemsManager.Signature\",\"name\":\"_signature\",\"type\":\"tuple\"}],\"name\":\"submitUptimeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitUptimeVoteMinDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"submitUptimeVoteMinDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToFallbackMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"}],\"name\":\"timelockedCalls\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerExpirationAndCleanup\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"randomAcquisitionMaxDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"randomAcquisitionMaxDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"newSigningPolicyInitializationStartSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"newSigningPolicyMinNumberOfVotingRoundsDelay\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"voterRegistrationMinDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"voterRegistrationMinDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"submitUptimeVoteMinDurationSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"submitUptimeVoteMinDurationBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"signingPolicyThresholdPPM\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"signingPolicyMinNumberOfVoters\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"rewardExpiryOffsetSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structFlareSystemsManager.Settings\",\"name\":\"_settings\",\"type\":\"tuple\"}],\"name\":\"updateSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewardEpochId\",\"type\":\"uint256\"}],\"name\":\"uptimeVoteHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistrationMinDurationBlocks\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistrationMinDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistrationTriggerContract\",\"outputs\":[{\"internalType\":\"contractIIVoterRegistrationTrigger\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voterRegistry\",\"outputs\":[{\"internalType\":\"contractIIVoterRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingEpochDurationSeconds\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	ID:  "FlareSystemsManager",
}

// FlareSystemsManager is an auto generated Go binding around an Ethereum contract.
type FlareSystemsManager struct {
	abi abi.ABI
}

// NewFlareSystemsManager creates a new instance of FlareSystemsManager.
func NewFlareSystemsManager() *FlareSystemsManager {
	parsed, err := FlareSystemsManagerMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &FlareSystemsManager{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *FlareSystemsManager) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackConstructor is the Go binding used to pack the parameters required for
// contract deployment.
//
// Solidity: constructor(address _governanceSettings, address _initialGovernance, address _addressUpdater, address _flareDaemon, (uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint24,uint16,uint32) _settings, uint32 _firstVotingRoundStartTs, uint8 _votingEpochDurationSeconds, uint32 _firstRewardEpochStartVotingRoundId, uint16 _rewardEpochDurationInVotingEpochs, (uint16,uint24,uint16) _initialSettings) returns()
func (flareSystemsManager *FlareSystemsManager) PackConstructor(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address, _flareDaemon common.Address, _settings FlareSystemsManagerSettings, _firstVotingRoundStartTs uint32, _votingEpochDurationSeconds uint8, _firstRewardEpochStartVotingRoundId uint32, _rewardEpochDurationInVotingEpochs uint16, _initialSettings FlareSystemsManagerInitialSettings) []byte {
	enc, err := flareSystemsManager.abi.Pack("", _governanceSettings, _initialGovernance, _addressUpdater, _flareDaemon, _settings, _firstVotingRoundStartTs, _votingEpochDurationSeconds, _firstRewardEpochStartVotingRoundId, _rewardEpochDurationInVotingEpochs, _initialSettings)
	if err != nil {
		panic(err)
	}
	return enc
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67fc4029.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes4 _selector) returns()
func (flareSystemsManager *FlareSystemsManager) PackCancelGovernanceCall(selector [4]byte) []byte {
	enc, err := flareSystemsManager.abi.Pack("cancelGovernanceCall", selector)
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
func (flareSystemsManager *FlareSystemsManager) TryPackCancelGovernanceCall(selector [4]byte) ([]byte, error) {
	return flareSystemsManager.abi.Pack("cancelGovernanceCall", selector)
}

// PackCleanupBlockNumberManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4eac870f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (flareSystemsManager *FlareSystemsManager) PackCleanupBlockNumberManager() []byte {
	enc, err := flareSystemsManager.abi.Pack("cleanupBlockNumberManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCleanupBlockNumberManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4eac870f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (flareSystemsManager *FlareSystemsManager) TryPackCleanupBlockNumberManager() ([]byte, error) {
	return flareSystemsManager.abi.Pack("cleanupBlockNumberManager")
}

// UnpackCleanupBlockNumberManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4eac870f.
//
// Solidity: function cleanupBlockNumberManager() view returns(address)
func (flareSystemsManager *FlareSystemsManager) UnpackCleanupBlockNumberManager(data []byte) (common.Address, error) {
	out, err := flareSystemsManager.abi.Unpack("cleanupBlockNumberManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCurrentRewardEpochExpectedEndTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed54fd63.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function currentRewardEpochExpectedEndTs() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) PackCurrentRewardEpochExpectedEndTs() []byte {
	enc, err := flareSystemsManager.abi.Pack("currentRewardEpochExpectedEndTs")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCurrentRewardEpochExpectedEndTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed54fd63.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function currentRewardEpochExpectedEndTs() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) TryPackCurrentRewardEpochExpectedEndTs() ([]byte, error) {
	return flareSystemsManager.abi.Pack("currentRewardEpochExpectedEndTs")
}

// UnpackCurrentRewardEpochExpectedEndTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xed54fd63.
//
// Solidity: function currentRewardEpochExpectedEndTs() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackCurrentRewardEpochExpectedEndTs(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("currentRewardEpochExpectedEndTs", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackDaemonize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6d0e8c34.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function daemonize() returns(bool)
func (flareSystemsManager *FlareSystemsManager) PackDaemonize() []byte {
	enc, err := flareSystemsManager.abi.Pack("daemonize")
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
func (flareSystemsManager *FlareSystemsManager) TryPackDaemonize() ([]byte, error) {
	return flareSystemsManager.abi.Pack("daemonize")
}

// UnpackDaemonize is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6d0e8c34.
//
// Solidity: function daemonize() returns(bool)
func (flareSystemsManager *FlareSystemsManager) UnpackDaemonize(data []byte) (bool, error) {
	out, err := flareSystemsManager.abi.Unpack("daemonize", data)
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
func (flareSystemsManager *FlareSystemsManager) PackExecuteGovernanceCall(selector [4]byte) []byte {
	enc, err := flareSystemsManager.abi.Pack("executeGovernanceCall", selector)
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
func (flareSystemsManager *FlareSystemsManager) TryPackExecuteGovernanceCall(selector [4]byte) ([]byte, error) {
	return flareSystemsManager.abi.Pack("executeGovernanceCall", selector)
}

// PackFirstRewardEpochStartTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79e047ed.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function firstRewardEpochStartTs() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) PackFirstRewardEpochStartTs() []byte {
	enc, err := flareSystemsManager.abi.Pack("firstRewardEpochStartTs")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFirstRewardEpochStartTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79e047ed.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function firstRewardEpochStartTs() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) TryPackFirstRewardEpochStartTs() ([]byte, error) {
	return flareSystemsManager.abi.Pack("firstRewardEpochStartTs")
}

// UnpackFirstRewardEpochStartTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x79e047ed.
//
// Solidity: function firstRewardEpochStartTs() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackFirstRewardEpochStartTs(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("firstRewardEpochStartTs", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackFirstVotingRoundStartTs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe8d0e70a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) PackFirstVotingRoundStartTs() []byte {
	enc, err := flareSystemsManager.abi.Pack("firstVotingRoundStartTs")
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
func (flareSystemsManager *FlareSystemsManager) TryPackFirstVotingRoundStartTs() ([]byte, error) {
	return flareSystemsManager.abi.Pack("firstVotingRoundStartTs")
}

// UnpackFirstVotingRoundStartTs is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe8d0e70a.
//
// Solidity: function firstVotingRoundStartTs() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackFirstVotingRoundStartTs(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("firstVotingRoundStartTs", data)
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
func (flareSystemsManager *FlareSystemsManager) PackFlareDaemon() []byte {
	enc, err := flareSystemsManager.abi.Pack("flareDaemon")
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
func (flareSystemsManager *FlareSystemsManager) TryPackFlareDaemon() ([]byte, error) {
	return flareSystemsManager.abi.Pack("flareDaemon")
}

// UnpackFlareDaemon is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa1077532.
//
// Solidity: function flareDaemon() view returns(address)
func (flareSystemsManager *FlareSystemsManager) UnpackFlareDaemon(data []byte) (common.Address, error) {
	out, err := flareSystemsManager.abi.Unpack("flareDaemon", data)
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
func (flareSystemsManager *FlareSystemsManager) PackGetAddressUpdater() []byte {
	enc, err := flareSystemsManager.abi.Pack("getAddressUpdater")
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
func (flareSystemsManager *FlareSystemsManager) TryPackGetAddressUpdater() ([]byte, error) {
	return flareSystemsManager.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (flareSystemsManager *FlareSystemsManager) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := flareSystemsManager.abi.Unpack("getAddressUpdater", data)
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
func (flareSystemsManager *FlareSystemsManager) PackGetContractName() []byte {
	enc, err := flareSystemsManager.abi.Pack("getContractName")
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
func (flareSystemsManager *FlareSystemsManager) TryPackGetContractName() ([]byte, error) {
	return flareSystemsManager.abi.Pack("getContractName")
}

// UnpackGetContractName is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf5f5ba72.
//
// Solidity: function getContractName() pure returns(string)
func (flareSystemsManager *FlareSystemsManager) UnpackGetContractName(data []byte) (string, error) {
	out, err := flareSystemsManager.abi.Unpack("getContractName", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackGetCurrentRewardEpoch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe7c830d4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCurrentRewardEpoch() view returns(uint256)
func (flareSystemsManager *FlareSystemsManager) PackGetCurrentRewardEpoch() []byte {
	enc, err := flareSystemsManager.abi.Pack("getCurrentRewardEpoch")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCurrentRewardEpoch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe7c830d4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCurrentRewardEpoch() view returns(uint256)
func (flareSystemsManager *FlareSystemsManager) TryPackGetCurrentRewardEpoch() ([]byte, error) {
	return flareSystemsManager.abi.Pack("getCurrentRewardEpoch")
}

// UnpackGetCurrentRewardEpoch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe7c830d4.
//
// Solidity: function getCurrentRewardEpoch() view returns(uint256)
func (flareSystemsManager *FlareSystemsManager) UnpackGetCurrentRewardEpoch(data []byte) (*big.Int, error) {
	out, err := flareSystemsManager.abi.Unpack("getCurrentRewardEpoch", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetCurrentRewardEpochId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70562697.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (flareSystemsManager *FlareSystemsManager) PackGetCurrentRewardEpochId() []byte {
	enc, err := flareSystemsManager.abi.Pack("getCurrentRewardEpochId")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCurrentRewardEpochId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x70562697.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (flareSystemsManager *FlareSystemsManager) TryPackGetCurrentRewardEpochId() ([]byte, error) {
	return flareSystemsManager.abi.Pack("getCurrentRewardEpochId")
}

// UnpackGetCurrentRewardEpochId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x70562697.
//
// Solidity: function getCurrentRewardEpochId() view returns(uint24)
func (flareSystemsManager *FlareSystemsManager) UnpackGetCurrentRewardEpochId(data []byte) (*big.Int, error) {
	out, err := flareSystemsManager.abi.Unpack("getCurrentRewardEpochId", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetCurrentVotingEpochId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4134520b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCurrentVotingEpochId() view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) PackGetCurrentVotingEpochId() []byte {
	enc, err := flareSystemsManager.abi.Pack("getCurrentVotingEpochId")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCurrentVotingEpochId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4134520b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCurrentVotingEpochId() view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) TryPackGetCurrentVotingEpochId() ([]byte, error) {
	return flareSystemsManager.abi.Pack("getCurrentVotingEpochId")
}

// UnpackGetCurrentVotingEpochId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4134520b.
//
// Solidity: function getCurrentVotingEpochId() view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) UnpackGetCurrentVotingEpochId(data []byte) (uint32, error) {
	out, err := flareSystemsManager.abi.Unpack("getCurrentVotingEpochId", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackGetRandomAcquisitionInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f8f9f3a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRandomAcquisitionInfo(uint24 _rewardEpochId) view returns(uint64 _randomAcquisitionStartTs, uint64 _randomAcquisitionStartBlock, uint64 _randomAcquisitionEndTs, uint64 _randomAcquisitionEndBlock)
func (flareSystemsManager *FlareSystemsManager) PackGetRandomAcquisitionInfo(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("getRandomAcquisitionInfo", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRandomAcquisitionInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8f8f9f3a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRandomAcquisitionInfo(uint24 _rewardEpochId) view returns(uint64 _randomAcquisitionStartTs, uint64 _randomAcquisitionStartBlock, uint64 _randomAcquisitionEndTs, uint64 _randomAcquisitionEndBlock)
func (flareSystemsManager *FlareSystemsManager) TryPackGetRandomAcquisitionInfo(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getRandomAcquisitionInfo", rewardEpochId)
}

// GetRandomAcquisitionInfoOutput serves as a container for the return parameters of contract
// method GetRandomAcquisitionInfo.
type GetRandomAcquisitionInfoOutput struct {
	RandomAcquisitionStartTs    uint64
	RandomAcquisitionStartBlock uint64
	RandomAcquisitionEndTs      uint64
	RandomAcquisitionEndBlock   uint64
}

// UnpackGetRandomAcquisitionInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8f8f9f3a.
//
// Solidity: function getRandomAcquisitionInfo(uint24 _rewardEpochId) view returns(uint64 _randomAcquisitionStartTs, uint64 _randomAcquisitionStartBlock, uint64 _randomAcquisitionEndTs, uint64 _randomAcquisitionEndBlock)
func (flareSystemsManager *FlareSystemsManager) UnpackGetRandomAcquisitionInfo(data []byte) (GetRandomAcquisitionInfoOutput, error) {
	out, err := flareSystemsManager.abi.Unpack("getRandomAcquisitionInfo", data)
	outstruct := new(GetRandomAcquisitionInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.RandomAcquisitionStartTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.RandomAcquisitionStartBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.RandomAcquisitionEndTs = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.RandomAcquisitionEndBlock = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetRewardEpochStartInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00ddae53.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRewardEpochStartInfo(uint24 _rewardEpochId) view returns(uint64 _rewardEpochStartTs, uint64 _rewardEpochStartBlock)
func (flareSystemsManager *FlareSystemsManager) PackGetRewardEpochStartInfo(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("getRewardEpochStartInfo", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRewardEpochStartInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00ddae53.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRewardEpochStartInfo(uint24 _rewardEpochId) view returns(uint64 _rewardEpochStartTs, uint64 _rewardEpochStartBlock)
func (flareSystemsManager *FlareSystemsManager) TryPackGetRewardEpochStartInfo(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getRewardEpochStartInfo", rewardEpochId)
}

// GetRewardEpochStartInfoOutput serves as a container for the return parameters of contract
// method GetRewardEpochStartInfo.
type GetRewardEpochStartInfoOutput struct {
	RewardEpochStartTs    uint64
	RewardEpochStartBlock uint64
}

// UnpackGetRewardEpochStartInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x00ddae53.
//
// Solidity: function getRewardEpochStartInfo(uint24 _rewardEpochId) view returns(uint64 _rewardEpochStartTs, uint64 _rewardEpochStartBlock)
func (flareSystemsManager *FlareSystemsManager) UnpackGetRewardEpochStartInfo(data []byte) (GetRewardEpochStartInfoOutput, error) {
	out, err := flareSystemsManager.abi.Unpack("getRewardEpochStartInfo", data)
	outstruct := new(GetRewardEpochStartInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.RewardEpochStartTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.RewardEpochStartBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetRewardEpochSwitchoverTriggerContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46831531.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRewardEpochSwitchoverTriggerContracts() view returns(address[])
func (flareSystemsManager *FlareSystemsManager) PackGetRewardEpochSwitchoverTriggerContracts() []byte {
	enc, err := flareSystemsManager.abi.Pack("getRewardEpochSwitchoverTriggerContracts")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRewardEpochSwitchoverTriggerContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46831531.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRewardEpochSwitchoverTriggerContracts() view returns(address[])
func (flareSystemsManager *FlareSystemsManager) TryPackGetRewardEpochSwitchoverTriggerContracts() ([]byte, error) {
	return flareSystemsManager.abi.Pack("getRewardEpochSwitchoverTriggerContracts")
}

// UnpackGetRewardEpochSwitchoverTriggerContracts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x46831531.
//
// Solidity: function getRewardEpochSwitchoverTriggerContracts() view returns(address[])
func (flareSystemsManager *FlareSystemsManager) UnpackGetRewardEpochSwitchoverTriggerContracts(data []byte) ([]common.Address, error) {
	out, err := flareSystemsManager.abi.Unpack("getRewardEpochSwitchoverTriggerContracts", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetRewardsSignInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb6c25af0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRewardsSignInfo(uint24 _rewardEpochId) view returns(uint64 _rewardsSignStartTs, uint64 _rewardsSignStartBlock, uint64 _rewardsSignEndTs, uint64 _rewardsSignEndBlock)
func (flareSystemsManager *FlareSystemsManager) PackGetRewardsSignInfo(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("getRewardsSignInfo", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRewardsSignInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb6c25af0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRewardsSignInfo(uint24 _rewardEpochId) view returns(uint64 _rewardsSignStartTs, uint64 _rewardsSignStartBlock, uint64 _rewardsSignEndTs, uint64 _rewardsSignEndBlock)
func (flareSystemsManager *FlareSystemsManager) TryPackGetRewardsSignInfo(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getRewardsSignInfo", rewardEpochId)
}

// GetRewardsSignInfoOutput serves as a container for the return parameters of contract
// method GetRewardsSignInfo.
type GetRewardsSignInfoOutput struct {
	RewardsSignStartTs    uint64
	RewardsSignStartBlock uint64
	RewardsSignEndTs      uint64
	RewardsSignEndBlock   uint64
}

// UnpackGetRewardsSignInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb6c25af0.
//
// Solidity: function getRewardsSignInfo(uint24 _rewardEpochId) view returns(uint64 _rewardsSignStartTs, uint64 _rewardsSignStartBlock, uint64 _rewardsSignEndTs, uint64 _rewardsSignEndBlock)
func (flareSystemsManager *FlareSystemsManager) UnpackGetRewardsSignInfo(data []byte) (GetRewardsSignInfoOutput, error) {
	out, err := flareSystemsManager.abi.Unpack("getRewardsSignInfo", data)
	outstruct := new(GetRewardsSignInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.RewardsSignStartTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.RewardsSignStartBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.RewardsSignEndTs = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.RewardsSignEndBlock = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetSeed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe0d4ea37.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSeed(uint256 _rewardEpochId) view returns(uint256)
func (flareSystemsManager *FlareSystemsManager) PackGetSeed(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("getSeed", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSeed is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe0d4ea37.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSeed(uint256 _rewardEpochId) view returns(uint256)
func (flareSystemsManager *FlareSystemsManager) TryPackGetSeed(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getSeed", rewardEpochId)
}

// UnpackGetSeed is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe0d4ea37.
//
// Solidity: function getSeed(uint256 _rewardEpochId) view returns(uint256)
func (flareSystemsManager *FlareSystemsManager) UnpackGetSeed(data []byte) (*big.Int, error) {
	out, err := flareSystemsManager.abi.Unpack("getSeed", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetSigningPolicySignInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2e9ad71.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSigningPolicySignInfo(uint24 _rewardEpochId) view returns(uint64 _signingPolicySignStartTs, uint64 _signingPolicySignStartBlock, uint64 _signingPolicySignEndTs, uint64 _signingPolicySignEndBlock)
func (flareSystemsManager *FlareSystemsManager) PackGetSigningPolicySignInfo(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("getSigningPolicySignInfo", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSigningPolicySignInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2e9ad71.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSigningPolicySignInfo(uint24 _rewardEpochId) view returns(uint64 _signingPolicySignStartTs, uint64 _signingPolicySignStartBlock, uint64 _signingPolicySignEndTs, uint64 _signingPolicySignEndBlock)
func (flareSystemsManager *FlareSystemsManager) TryPackGetSigningPolicySignInfo(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getSigningPolicySignInfo", rewardEpochId)
}

// GetSigningPolicySignInfoOutput serves as a container for the return parameters of contract
// method GetSigningPolicySignInfo.
type GetSigningPolicySignInfoOutput struct {
	SigningPolicySignStartTs    uint64
	SigningPolicySignStartBlock uint64
	SigningPolicySignEndTs      uint64
	SigningPolicySignEndBlock   uint64
}

// UnpackGetSigningPolicySignInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd2e9ad71.
//
// Solidity: function getSigningPolicySignInfo(uint24 _rewardEpochId) view returns(uint64 _signingPolicySignStartTs, uint64 _signingPolicySignStartBlock, uint64 _signingPolicySignEndTs, uint64 _signingPolicySignEndBlock)
func (flareSystemsManager *FlareSystemsManager) UnpackGetSigningPolicySignInfo(data []byte) (GetSigningPolicySignInfoOutput, error) {
	out, err := flareSystemsManager.abi.Unpack("getSigningPolicySignInfo", data)
	outstruct := new(GetSigningPolicySignInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.SigningPolicySignStartTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.SigningPolicySignStartBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.SigningPolicySignEndTs = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.SigningPolicySignEndBlock = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetStartVotingRoundId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75d2187a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getStartVotingRoundId(uint256 _rewardEpochId) view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) PackGetStartVotingRoundId(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("getStartVotingRoundId", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetStartVotingRoundId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x75d2187a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getStartVotingRoundId(uint256 _rewardEpochId) view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) TryPackGetStartVotingRoundId(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getStartVotingRoundId", rewardEpochId)
}

// UnpackGetStartVotingRoundId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x75d2187a.
//
// Solidity: function getStartVotingRoundId(uint256 _rewardEpochId) view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) UnpackGetStartVotingRoundId(data []byte) (uint32, error) {
	out, err := flareSystemsManager.abi.Unpack("getStartVotingRoundId", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackGetThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4615d5e9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getThreshold(uint256 _rewardEpochId) view returns(uint16)
func (flareSystemsManager *FlareSystemsManager) PackGetThreshold(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("getThreshold", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetThreshold is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4615d5e9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getThreshold(uint256 _rewardEpochId) view returns(uint16)
func (flareSystemsManager *FlareSystemsManager) TryPackGetThreshold(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getThreshold", rewardEpochId)
}

// UnpackGetThreshold is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4615d5e9.
//
// Solidity: function getThreshold(uint256 _rewardEpochId) view returns(uint16)
func (flareSystemsManager *FlareSystemsManager) UnpackGetThreshold(data []byte) (uint16, error) {
	out, err := flareSystemsManager.abi.Unpack("getThreshold", data)
	if err != nil {
		return *new(uint16), err
	}
	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	return out0, nil
}

// PackGetUptimeVoteSignStartInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc9f1d2aa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getUptimeVoteSignStartInfo(uint24 _rewardEpochId) view returns(uint64 _uptimeVoteSignStartTs, uint64 _uptimeVoteSignStartBlock)
func (flareSystemsManager *FlareSystemsManager) PackGetUptimeVoteSignStartInfo(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("getUptimeVoteSignStartInfo", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetUptimeVoteSignStartInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc9f1d2aa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getUptimeVoteSignStartInfo(uint24 _rewardEpochId) view returns(uint64 _uptimeVoteSignStartTs, uint64 _uptimeVoteSignStartBlock)
func (flareSystemsManager *FlareSystemsManager) TryPackGetUptimeVoteSignStartInfo(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getUptimeVoteSignStartInfo", rewardEpochId)
}

// GetUptimeVoteSignStartInfoOutput serves as a container for the return parameters of contract
// method GetUptimeVoteSignStartInfo.
type GetUptimeVoteSignStartInfoOutput struct {
	UptimeVoteSignStartTs    uint64
	UptimeVoteSignStartBlock uint64
}

// UnpackGetUptimeVoteSignStartInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc9f1d2aa.
//
// Solidity: function getUptimeVoteSignStartInfo(uint24 _rewardEpochId) view returns(uint64 _uptimeVoteSignStartTs, uint64 _uptimeVoteSignStartBlock)
func (flareSystemsManager *FlareSystemsManager) UnpackGetUptimeVoteSignStartInfo(data []byte) (GetUptimeVoteSignStartInfoOutput, error) {
	out, err := flareSystemsManager.abi.Unpack("getUptimeVoteSignStartInfo", data)
	outstruct := new(GetUptimeVoteSignStartInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.UptimeVoteSignStartTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.UptimeVoteSignStartBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetVotePowerBlock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc2632216.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpochId) view returns(uint64 _votePowerBlock)
func (flareSystemsManager *FlareSystemsManager) PackGetVotePowerBlock(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("getVotePowerBlock", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVotePowerBlock is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc2632216.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpochId) view returns(uint64 _votePowerBlock)
func (flareSystemsManager *FlareSystemsManager) TryPackGetVotePowerBlock(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getVotePowerBlock", rewardEpochId)
}

// UnpackGetVotePowerBlock is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc2632216.
//
// Solidity: function getVotePowerBlock(uint256 _rewardEpochId) view returns(uint64 _votePowerBlock)
func (flareSystemsManager *FlareSystemsManager) UnpackGetVotePowerBlock(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("getVotePowerBlock", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetVoterRegistrationData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1703a788.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpochId) view returns(uint256 _votePowerBlock, bool _enabled)
func (flareSystemsManager *FlareSystemsManager) PackGetVoterRegistrationData(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("getVoterRegistrationData", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterRegistrationData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1703a788.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpochId) view returns(uint256 _votePowerBlock, bool _enabled)
func (flareSystemsManager *FlareSystemsManager) TryPackGetVoterRegistrationData(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getVoterRegistrationData", rewardEpochId)
}

// GetVoterRegistrationDataOutput serves as a container for the return parameters of contract
// method GetVoterRegistrationData.
type GetVoterRegistrationDataOutput struct {
	VotePowerBlock *big.Int
	Enabled        bool
}

// UnpackGetVoterRegistrationData is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1703a788.
//
// Solidity: function getVoterRegistrationData(uint256 _rewardEpochId) view returns(uint256 _votePowerBlock, bool _enabled)
func (flareSystemsManager *FlareSystemsManager) UnpackGetVoterRegistrationData(data []byte) (GetVoterRegistrationDataOutput, error) {
	out, err := flareSystemsManager.abi.Unpack("getVoterRegistrationData", data)
	outstruct := new(GetVoterRegistrationDataOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.VotePowerBlock = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.Enabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	return *outstruct, nil
}

// PackGetVoterRewardsSignInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1916e915.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterRewardsSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _rewardsSignTs, uint64 _rewardsSignBlock)
func (flareSystemsManager *FlareSystemsManager) PackGetVoterRewardsSignInfo(rewardEpochId *big.Int, voter common.Address) []byte {
	enc, err := flareSystemsManager.abi.Pack("getVoterRewardsSignInfo", rewardEpochId, voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterRewardsSignInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1916e915.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterRewardsSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _rewardsSignTs, uint64 _rewardsSignBlock)
func (flareSystemsManager *FlareSystemsManager) TryPackGetVoterRewardsSignInfo(rewardEpochId *big.Int, voter common.Address) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getVoterRewardsSignInfo", rewardEpochId, voter)
}

// GetVoterRewardsSignInfoOutput serves as a container for the return parameters of contract
// method GetVoterRewardsSignInfo.
type GetVoterRewardsSignInfoOutput struct {
	RewardsSignTs    uint64
	RewardsSignBlock uint64
}

// UnpackGetVoterRewardsSignInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1916e915.
//
// Solidity: function getVoterRewardsSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _rewardsSignTs, uint64 _rewardsSignBlock)
func (flareSystemsManager *FlareSystemsManager) UnpackGetVoterRewardsSignInfo(data []byte) (GetVoterRewardsSignInfoOutput, error) {
	out, err := flareSystemsManager.abi.Unpack("getVoterRewardsSignInfo", data)
	outstruct := new(GetVoterRewardsSignInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.RewardsSignTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.RewardsSignBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetVoterSigningPolicySignInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdac4319d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterSigningPolicySignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _signingPolicySignTs, uint64 _signingPolicySignBlock)
func (flareSystemsManager *FlareSystemsManager) PackGetVoterSigningPolicySignInfo(rewardEpochId *big.Int, voter common.Address) []byte {
	enc, err := flareSystemsManager.abi.Pack("getVoterSigningPolicySignInfo", rewardEpochId, voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterSigningPolicySignInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdac4319d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterSigningPolicySignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _signingPolicySignTs, uint64 _signingPolicySignBlock)
func (flareSystemsManager *FlareSystemsManager) TryPackGetVoterSigningPolicySignInfo(rewardEpochId *big.Int, voter common.Address) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getVoterSigningPolicySignInfo", rewardEpochId, voter)
}

// GetVoterSigningPolicySignInfoOutput serves as a container for the return parameters of contract
// method GetVoterSigningPolicySignInfo.
type GetVoterSigningPolicySignInfoOutput struct {
	SigningPolicySignTs    uint64
	SigningPolicySignBlock uint64
}

// UnpackGetVoterSigningPolicySignInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdac4319d.
//
// Solidity: function getVoterSigningPolicySignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _signingPolicySignTs, uint64 _signingPolicySignBlock)
func (flareSystemsManager *FlareSystemsManager) UnpackGetVoterSigningPolicySignInfo(data []byte) (GetVoterSigningPolicySignInfoOutput, error) {
	out, err := flareSystemsManager.abi.Unpack("getVoterSigningPolicySignInfo", data)
	outstruct := new(GetVoterSigningPolicySignInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.SigningPolicySignTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.SigningPolicySignBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetVoterUptimeVoteSignInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41c05ad5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterUptimeVoteSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSignTs, uint64 _uptimeVoteSignBlock)
func (flareSystemsManager *FlareSystemsManager) PackGetVoterUptimeVoteSignInfo(rewardEpochId *big.Int, voter common.Address) []byte {
	enc, err := flareSystemsManager.abi.Pack("getVoterUptimeVoteSignInfo", rewardEpochId, voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterUptimeVoteSignInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41c05ad5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterUptimeVoteSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSignTs, uint64 _uptimeVoteSignBlock)
func (flareSystemsManager *FlareSystemsManager) TryPackGetVoterUptimeVoteSignInfo(rewardEpochId *big.Int, voter common.Address) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getVoterUptimeVoteSignInfo", rewardEpochId, voter)
}

// GetVoterUptimeVoteSignInfoOutput serves as a container for the return parameters of contract
// method GetVoterUptimeVoteSignInfo.
type GetVoterUptimeVoteSignInfoOutput struct {
	UptimeVoteSignTs    uint64
	UptimeVoteSignBlock uint64
}

// UnpackGetVoterUptimeVoteSignInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x41c05ad5.
//
// Solidity: function getVoterUptimeVoteSignInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSignTs, uint64 _uptimeVoteSignBlock)
func (flareSystemsManager *FlareSystemsManager) UnpackGetVoterUptimeVoteSignInfo(data []byte) (GetVoterUptimeVoteSignInfoOutput, error) {
	out, err := flareSystemsManager.abi.Unpack("getVoterUptimeVoteSignInfo", data)
	outstruct := new(GetVoterUptimeVoteSignInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.UptimeVoteSignTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.UptimeVoteSignBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetVoterUptimeVoteSubmitInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x59db0e2f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getVoterUptimeVoteSubmitInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSubmitTs, uint64 _uptimeVoteSubmitBlock)
func (flareSystemsManager *FlareSystemsManager) PackGetVoterUptimeVoteSubmitInfo(rewardEpochId *big.Int, voter common.Address) []byte {
	enc, err := flareSystemsManager.abi.Pack("getVoterUptimeVoteSubmitInfo", rewardEpochId, voter)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetVoterUptimeVoteSubmitInfo is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x59db0e2f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getVoterUptimeVoteSubmitInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSubmitTs, uint64 _uptimeVoteSubmitBlock)
func (flareSystemsManager *FlareSystemsManager) TryPackGetVoterUptimeVoteSubmitInfo(rewardEpochId *big.Int, voter common.Address) ([]byte, error) {
	return flareSystemsManager.abi.Pack("getVoterUptimeVoteSubmitInfo", rewardEpochId, voter)
}

// GetVoterUptimeVoteSubmitInfoOutput serves as a container for the return parameters of contract
// method GetVoterUptimeVoteSubmitInfo.
type GetVoterUptimeVoteSubmitInfoOutput struct {
	UptimeVoteSubmitTs    uint64
	UptimeVoteSubmitBlock uint64
}

// UnpackGetVoterUptimeVoteSubmitInfo is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x59db0e2f.
//
// Solidity: function getVoterUptimeVoteSubmitInfo(uint24 _rewardEpochId, address _voter) view returns(uint64 _uptimeVoteSubmitTs, uint64 _uptimeVoteSubmitBlock)
func (flareSystemsManager *FlareSystemsManager) UnpackGetVoterUptimeVoteSubmitInfo(data []byte) (GetVoterUptimeVoteSubmitInfoOutput, error) {
	out, err := flareSystemsManager.abi.Unpack("getVoterUptimeVoteSubmitInfo", data)
	outstruct := new(GetVoterUptimeVoteSubmitInfoOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.UptimeVoteSubmitTs = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.UptimeVoteSubmitBlock = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governance() view returns(address)
func (flareSystemsManager *FlareSystemsManager) PackGovernance() []byte {
	enc, err := flareSystemsManager.abi.Pack("governance")
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
func (flareSystemsManager *FlareSystemsManager) TryPackGovernance() ([]byte, error) {
	return flareSystemsManager.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (flareSystemsManager *FlareSystemsManager) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := flareSystemsManager.abi.Unpack("governance", data)
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
func (flareSystemsManager *FlareSystemsManager) PackGovernanceSettings() []byte {
	enc, err := flareSystemsManager.abi.Pack("governanceSettings")
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
func (flareSystemsManager *FlareSystemsManager) TryPackGovernanceSettings() ([]byte, error) {
	return flareSystemsManager.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (flareSystemsManager *FlareSystemsManager) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := flareSystemsManager.abi.Unpack("governanceSettings", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitialRandomVotePowerBlockSelectionSize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xded7c4b8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialRandomVotePowerBlockSelectionSize() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) PackInitialRandomVotePowerBlockSelectionSize() []byte {
	enc, err := flareSystemsManager.abi.Pack("initialRandomVotePowerBlockSelectionSize")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialRandomVotePowerBlockSelectionSize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xded7c4b8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialRandomVotePowerBlockSelectionSize() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) TryPackInitialRandomVotePowerBlockSelectionSize() ([]byte, error) {
	return flareSystemsManager.abi.Pack("initialRandomVotePowerBlockSelectionSize")
}

// UnpackInitialRandomVotePowerBlockSelectionSize is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xded7c4b8.
//
// Solidity: function initialRandomVotePowerBlockSelectionSize() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackInitialRandomVotePowerBlockSelectionSize(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("initialRandomVotePowerBlockSelectionSize", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackInitialise is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xef88bf13.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialise(address _governanceSettings, address _initialGovernance) returns()
func (flareSystemsManager *FlareSystemsManager) PackInitialise(governanceSettings common.Address, initialGovernance common.Address) []byte {
	enc, err := flareSystemsManager.abi.Pack("initialise", governanceSettings, initialGovernance)
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
func (flareSystemsManager *FlareSystemsManager) TryPackInitialise(governanceSettings common.Address, initialGovernance common.Address) ([]byte, error) {
	return flareSystemsManager.abi.Pack("initialise", governanceSettings, initialGovernance)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (flareSystemsManager *FlareSystemsManager) PackIsExecutor(address common.Address) []byte {
	enc, err := flareSystemsManager.abi.Pack("isExecutor", address)
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
func (flareSystemsManager *FlareSystemsManager) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return flareSystemsManager.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (flareSystemsManager *FlareSystemsManager) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := flareSystemsManager.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsVoterRegistrationEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09505d25.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isVoterRegistrationEnabled() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) PackIsVoterRegistrationEnabled() []byte {
	enc, err := flareSystemsManager.abi.Pack("isVoterRegistrationEnabled")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsVoterRegistrationEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09505d25.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isVoterRegistrationEnabled() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) TryPackIsVoterRegistrationEnabled() ([]byte, error) {
	return flareSystemsManager.abi.Pack("isVoterRegistrationEnabled")
}

// UnpackIsVoterRegistrationEnabled is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x09505d25.
//
// Solidity: function isVoterRegistrationEnabled() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) UnpackIsVoterRegistrationEnabled(data []byte) (bool, error) {
	out, err := flareSystemsManager.abi.Unpack("isVoterRegistrationEnabled", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackLastInitializedVotingRoundId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f923d37.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function lastInitializedVotingRoundId() view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) PackLastInitializedVotingRoundId() []byte {
	enc, err := flareSystemsManager.abi.Pack("lastInitializedVotingRoundId")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackLastInitializedVotingRoundId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f923d37.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function lastInitializedVotingRoundId() view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) TryPackLastInitializedVotingRoundId() ([]byte, error) {
	return flareSystemsManager.abi.Pack("lastInitializedVotingRoundId")
}

// UnpackLastInitializedVotingRoundId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4f923d37.
//
// Solidity: function lastInitializedVotingRoundId() view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) UnpackLastInitializedVotingRoundId(data []byte) (uint32, error) {
	out, err := flareSystemsManager.abi.Unpack("lastInitializedVotingRoundId", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackNewSigningPolicyInitializationStartSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6aeffddc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function newSigningPolicyInitializationStartSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) PackNewSigningPolicyInitializationStartSeconds() []byte {
	enc, err := flareSystemsManager.abi.Pack("newSigningPolicyInitializationStartSeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNewSigningPolicyInitializationStartSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6aeffddc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function newSigningPolicyInitializationStartSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) TryPackNewSigningPolicyInitializationStartSeconds() ([]byte, error) {
	return flareSystemsManager.abi.Pack("newSigningPolicyInitializationStartSeconds")
}

// UnpackNewSigningPolicyInitializationStartSeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6aeffddc.
//
// Solidity: function newSigningPolicyInitializationStartSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackNewSigningPolicyInitializationStartSeconds(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("newSigningPolicyInitializationStartSeconds", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackNewSigningPolicyMinNumberOfVotingRoundsDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa733d54b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function newSigningPolicyMinNumberOfVotingRoundsDelay() view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) PackNewSigningPolicyMinNumberOfVotingRoundsDelay() []byte {
	enc, err := flareSystemsManager.abi.Pack("newSigningPolicyMinNumberOfVotingRoundsDelay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNewSigningPolicyMinNumberOfVotingRoundsDelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa733d54b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function newSigningPolicyMinNumberOfVotingRoundsDelay() view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) TryPackNewSigningPolicyMinNumberOfVotingRoundsDelay() ([]byte, error) {
	return flareSystemsManager.abi.Pack("newSigningPolicyMinNumberOfVotingRoundsDelay")
}

// UnpackNewSigningPolicyMinNumberOfVotingRoundsDelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa733d54b.
//
// Solidity: function newSigningPolicyMinNumberOfVotingRoundsDelay() view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) UnpackNewSigningPolicyMinNumberOfVotingRoundsDelay(data []byte) (uint32, error) {
	out, err := flareSystemsManager.abi.Unpack("newSigningPolicyMinNumberOfVotingRoundsDelay", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackNoOfWeightBasedClaims is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc581e791.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function noOfWeightBasedClaims(uint256 rewardEpochId, uint256 rewardManagerId) view returns(uint256)
func (flareSystemsManager *FlareSystemsManager) PackNoOfWeightBasedClaims(rewardEpochId *big.Int, rewardManagerId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("noOfWeightBasedClaims", rewardEpochId, rewardManagerId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNoOfWeightBasedClaims is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc581e791.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function noOfWeightBasedClaims(uint256 rewardEpochId, uint256 rewardManagerId) view returns(uint256)
func (flareSystemsManager *FlareSystemsManager) TryPackNoOfWeightBasedClaims(rewardEpochId *big.Int, rewardManagerId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("noOfWeightBasedClaims", rewardEpochId, rewardManagerId)
}

// UnpackNoOfWeightBasedClaims is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc581e791.
//
// Solidity: function noOfWeightBasedClaims(uint256 rewardEpochId, uint256 rewardManagerId) view returns(uint256)
func (flareSystemsManager *FlareSystemsManager) UnpackNoOfWeightBasedClaims(data []byte) (*big.Int, error) {
	out, err := flareSystemsManager.abi.Unpack("noOfWeightBasedClaims", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackNoOfWeightBasedClaimsHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x787173e7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function noOfWeightBasedClaimsHash(uint256 rewardEpochId) view returns(bytes32)
func (flareSystemsManager *FlareSystemsManager) PackNoOfWeightBasedClaimsHash(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("noOfWeightBasedClaimsHash", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNoOfWeightBasedClaimsHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x787173e7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function noOfWeightBasedClaimsHash(uint256 rewardEpochId) view returns(bytes32)
func (flareSystemsManager *FlareSystemsManager) TryPackNoOfWeightBasedClaimsHash(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("noOfWeightBasedClaimsHash", rewardEpochId)
}

// UnpackNoOfWeightBasedClaimsHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x787173e7.
//
// Solidity: function noOfWeightBasedClaimsHash(uint256 rewardEpochId) view returns(bytes32)
func (flareSystemsManager *FlareSystemsManager) UnpackNoOfWeightBasedClaimsHash(data []byte) ([32]byte, error) {
	out, err := flareSystemsManager.abi.Unpack("noOfWeightBasedClaimsHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) PackProductionMode() []byte {
	enc, err := flareSystemsManager.abi.Pack("productionMode")
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
func (flareSystemsManager *FlareSystemsManager) TryPackProductionMode() ([]byte, error) {
	return flareSystemsManager.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) UnpackProductionMode(data []byte) (bool, error) {
	out, err := flareSystemsManager.abi.Unpack("productionMode", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackRandomAcquisitionMaxDurationBlocks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x490344f4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function randomAcquisitionMaxDurationBlocks() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) PackRandomAcquisitionMaxDurationBlocks() []byte {
	enc, err := flareSystemsManager.abi.Pack("randomAcquisitionMaxDurationBlocks")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRandomAcquisitionMaxDurationBlocks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x490344f4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function randomAcquisitionMaxDurationBlocks() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) TryPackRandomAcquisitionMaxDurationBlocks() ([]byte, error) {
	return flareSystemsManager.abi.Pack("randomAcquisitionMaxDurationBlocks")
}

// UnpackRandomAcquisitionMaxDurationBlocks is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x490344f4.
//
// Solidity: function randomAcquisitionMaxDurationBlocks() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackRandomAcquisitionMaxDurationBlocks(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("randomAcquisitionMaxDurationBlocks", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackRandomAcquisitionMaxDurationSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x098e7ff6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function randomAcquisitionMaxDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) PackRandomAcquisitionMaxDurationSeconds() []byte {
	enc, err := flareSystemsManager.abi.Pack("randomAcquisitionMaxDurationSeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRandomAcquisitionMaxDurationSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x098e7ff6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function randomAcquisitionMaxDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) TryPackRandomAcquisitionMaxDurationSeconds() ([]byte, error) {
	return flareSystemsManager.abi.Pack("randomAcquisitionMaxDurationSeconds")
}

// UnpackRandomAcquisitionMaxDurationSeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x098e7ff6.
//
// Solidity: function randomAcquisitionMaxDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackRandomAcquisitionMaxDurationSeconds(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("randomAcquisitionMaxDurationSeconds", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb59589d1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function relay() view returns(address)
func (flareSystemsManager *FlareSystemsManager) PackRelay() []byte {
	enc, err := flareSystemsManager.abi.Pack("relay")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRelay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb59589d1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function relay() view returns(address)
func (flareSystemsManager *FlareSystemsManager) TryPackRelay() ([]byte, error) {
	return flareSystemsManager.abi.Pack("relay")
}

// UnpackRelay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb59589d1.
//
// Solidity: function relay() view returns(address)
func (flareSystemsManager *FlareSystemsManager) UnpackRelay(data []byte) (common.Address, error) {
	out, err := flareSystemsManager.abi.Unpack("relay", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRewardEpochDurationSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85f3c9c9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) PackRewardEpochDurationSeconds() []byte {
	enc, err := flareSystemsManager.abi.Pack("rewardEpochDurationSeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRewardEpochDurationSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x85f3c9c9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) TryPackRewardEpochDurationSeconds() ([]byte, error) {
	return flareSystemsManager.abi.Pack("rewardEpochDurationSeconds")
}

// UnpackRewardEpochDurationSeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x85f3c9c9.
//
// Solidity: function rewardEpochDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackRewardEpochDurationSeconds(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("rewardEpochDurationSeconds", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackRewardEpochIdToExpireNext is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaec84ab6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardEpochIdToExpireNext() view returns(uint24)
func (flareSystemsManager *FlareSystemsManager) PackRewardEpochIdToExpireNext() []byte {
	enc, err := flareSystemsManager.abi.Pack("rewardEpochIdToExpireNext")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRewardEpochIdToExpireNext is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaec84ab6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rewardEpochIdToExpireNext() view returns(uint24)
func (flareSystemsManager *FlareSystemsManager) TryPackRewardEpochIdToExpireNext() ([]byte, error) {
	return flareSystemsManager.abi.Pack("rewardEpochIdToExpireNext")
}

// UnpackRewardEpochIdToExpireNext is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaec84ab6.
//
// Solidity: function rewardEpochIdToExpireNext() view returns(uint24)
func (flareSystemsManager *FlareSystemsManager) UnpackRewardEpochIdToExpireNext(data []byte) (*big.Int, error) {
	out, err := flareSystemsManager.abi.Unpack("rewardEpochIdToExpireNext", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackRewardExpiryOffsetSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4eaee307.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardExpiryOffsetSeconds() view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) PackRewardExpiryOffsetSeconds() []byte {
	enc, err := flareSystemsManager.abi.Pack("rewardExpiryOffsetSeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRewardExpiryOffsetSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4eaee307.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rewardExpiryOffsetSeconds() view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) TryPackRewardExpiryOffsetSeconds() ([]byte, error) {
	return flareSystemsManager.abi.Pack("rewardExpiryOffsetSeconds")
}

// UnpackRewardExpiryOffsetSeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4eaee307.
//
// Solidity: function rewardExpiryOffsetSeconds() view returns(uint32)
func (flareSystemsManager *FlareSystemsManager) UnpackRewardExpiryOffsetSeconds(data []byte) (uint32, error) {
	out, err := flareSystemsManager.abi.Unpack("rewardExpiryOffsetSeconds", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackRewardManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f4ef8a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardManager() view returns(address)
func (flareSystemsManager *FlareSystemsManager) PackRewardManager() []byte {
	enc, err := flareSystemsManager.abi.Pack("rewardManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRewardManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f4ef8a6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rewardManager() view returns(address)
func (flareSystemsManager *FlareSystemsManager) TryPackRewardManager() ([]byte, error) {
	return flareSystemsManager.abi.Pack("rewardManager")
}

// UnpackRewardManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (flareSystemsManager *FlareSystemsManager) UnpackRewardManager(data []byte) (common.Address, error) {
	out, err := flareSystemsManager.abi.Unpack("rewardManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackRewardsHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x647006e2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardsHash(uint256 rewardEpochId) view returns(bytes32)
func (flareSystemsManager *FlareSystemsManager) PackRewardsHash(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("rewardsHash", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRewardsHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x647006e2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function rewardsHash(uint256 rewardEpochId) view returns(bytes32)
func (flareSystemsManager *FlareSystemsManager) TryPackRewardsHash(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("rewardsHash", rewardEpochId)
}

// UnpackRewardsHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x647006e2.
//
// Solidity: function rewardsHash(uint256 rewardEpochId) view returns(bytes32)
func (flareSystemsManager *FlareSystemsManager) UnpackRewardsHash(data []byte) ([32]byte, error) {
	out, err := flareSystemsManager.abi.Unpack("rewardsHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackSetRewardEpochSwitchoverTriggerContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06886f41.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setRewardEpochSwitchoverTriggerContracts(address[] _contracts) returns()
func (flareSystemsManager *FlareSystemsManager) PackSetRewardEpochSwitchoverTriggerContracts(contracts []common.Address) []byte {
	enc, err := flareSystemsManager.abi.Pack("setRewardEpochSwitchoverTriggerContracts", contracts)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetRewardEpochSwitchoverTriggerContracts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x06886f41.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setRewardEpochSwitchoverTriggerContracts(address[] _contracts) returns()
func (flareSystemsManager *FlareSystemsManager) TryPackSetRewardEpochSwitchoverTriggerContracts(contracts []common.Address) ([]byte, error) {
	return flareSystemsManager.abi.Pack("setRewardEpochSwitchoverTriggerContracts", contracts)
}

// PackSetRewardsData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd4be6faa.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setRewardsData(uint24 _rewardEpochId, (uint256,uint256)[] _noOfWeightBasedClaims, bytes32 _rewardsHash) returns()
func (flareSystemsManager *FlareSystemsManager) PackSetRewardsData(rewardEpochId *big.Int, noOfWeightBasedClaims []IFlareSystemsManagerNumberOfWeightBasedClaims, rewardsHash [32]byte) []byte {
	enc, err := flareSystemsManager.abi.Pack("setRewardsData", rewardEpochId, noOfWeightBasedClaims, rewardsHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetRewardsData is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd4be6faa.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setRewardsData(uint24 _rewardEpochId, (uint256,uint256)[] _noOfWeightBasedClaims, bytes32 _rewardsHash) returns()
func (flareSystemsManager *FlareSystemsManager) TryPackSetRewardsData(rewardEpochId *big.Int, noOfWeightBasedClaims []IFlareSystemsManagerNumberOfWeightBasedClaims, rewardsHash [32]byte) ([]byte, error) {
	return flareSystemsManager.abi.Pack("setRewardsData", rewardEpochId, noOfWeightBasedClaims, rewardsHash)
}

// PackSetSubmit3Aligned is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa72b826e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSubmit3Aligned(bool _submit3Aligned) returns()
func (flareSystemsManager *FlareSystemsManager) PackSetSubmit3Aligned(submit3Aligned bool) []byte {
	enc, err := flareSystemsManager.abi.Pack("setSubmit3Aligned", submit3Aligned)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSubmit3Aligned is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa72b826e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSubmit3Aligned(bool _submit3Aligned) returns()
func (flareSystemsManager *FlareSystemsManager) TryPackSetSubmit3Aligned(submit3Aligned bool) ([]byte, error) {
	return flareSystemsManager.abi.Pack("setSubmit3Aligned", submit3Aligned)
}

// PackSetTriggerExpirationAndCleanup is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67daec89.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setTriggerExpirationAndCleanup(bool _triggerExpirationAndCleanup) returns()
func (flareSystemsManager *FlareSystemsManager) PackSetTriggerExpirationAndCleanup(triggerExpirationAndCleanup bool) []byte {
	enc, err := flareSystemsManager.abi.Pack("setTriggerExpirationAndCleanup", triggerExpirationAndCleanup)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetTriggerExpirationAndCleanup is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x67daec89.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setTriggerExpirationAndCleanup(bool _triggerExpirationAndCleanup) returns()
func (flareSystemsManager *FlareSystemsManager) TryPackSetTriggerExpirationAndCleanup(triggerExpirationAndCleanup bool) ([]byte, error) {
	return flareSystemsManager.abi.Pack("setTriggerExpirationAndCleanup", triggerExpirationAndCleanup)
}

// PackSetVoterRegistrationTriggerContract is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24eb64de.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setVoterRegistrationTriggerContract(address _contract) returns()
func (flareSystemsManager *FlareSystemsManager) PackSetVoterRegistrationTriggerContract(contract common.Address) []byte {
	enc, err := flareSystemsManager.abi.Pack("setVoterRegistrationTriggerContract", contract)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetVoterRegistrationTriggerContract is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x24eb64de.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setVoterRegistrationTriggerContract(address _contract) returns()
func (flareSystemsManager *FlareSystemsManager) TryPackSetVoterRegistrationTriggerContract(contract common.Address) ([]byte, error) {
	return flareSystemsManager.abi.Pack("setVoterRegistrationTriggerContract", contract)
}

// PackSignNewSigningPolicy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6b4c7bd6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (flareSystemsManager *FlareSystemsManager) PackSignNewSigningPolicy(rewardEpochId *big.Int, newSigningPolicyHash [32]byte, signature IFlareSystemsManagerSignature) []byte {
	enc, err := flareSystemsManager.abi.Pack("signNewSigningPolicy", rewardEpochId, newSigningPolicyHash, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSignNewSigningPolicy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6b4c7bd6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signNewSigningPolicy(uint24 _rewardEpochId, bytes32 _newSigningPolicyHash, (uint8,bytes32,bytes32) _signature) returns()
func (flareSystemsManager *FlareSystemsManager) TryPackSignNewSigningPolicy(rewardEpochId *big.Int, newSigningPolicyHash [32]byte, signature IFlareSystemsManagerSignature) ([]byte, error) {
	return flareSystemsManager.abi.Pack("signNewSigningPolicy", rewardEpochId, newSigningPolicyHash, signature)
}

// PackSignRewards is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc00a1a97.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signRewards(uint24 _rewardEpochId, (uint256,uint256)[] _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (flareSystemsManager *FlareSystemsManager) PackSignRewards(rewardEpochId *big.Int, noOfWeightBasedClaims []IFlareSystemsManagerNumberOfWeightBasedClaims, rewardsHash [32]byte, signature IFlareSystemsManagerSignature) []byte {
	enc, err := flareSystemsManager.abi.Pack("signRewards", rewardEpochId, noOfWeightBasedClaims, rewardsHash, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSignRewards is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc00a1a97.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signRewards(uint24 _rewardEpochId, (uint256,uint256)[] _noOfWeightBasedClaims, bytes32 _rewardsHash, (uint8,bytes32,bytes32) _signature) returns()
func (flareSystemsManager *FlareSystemsManager) TryPackSignRewards(rewardEpochId *big.Int, noOfWeightBasedClaims []IFlareSystemsManagerNumberOfWeightBasedClaims, rewardsHash [32]byte, signature IFlareSystemsManagerSignature) ([]byte, error) {
	return flareSystemsManager.abi.Pack("signRewards", rewardEpochId, noOfWeightBasedClaims, rewardsHash, signature)
}

// PackSignUptimeVote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdc5a4225.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (flareSystemsManager *FlareSystemsManager) PackSignUptimeVote(rewardEpochId *big.Int, uptimeVoteHash [32]byte, signature IFlareSystemsManagerSignature) []byte {
	enc, err := flareSystemsManager.abi.Pack("signUptimeVote", rewardEpochId, uptimeVoteHash, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSignUptimeVote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdc5a4225.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signUptimeVote(uint24 _rewardEpochId, bytes32 _uptimeVoteHash, (uint8,bytes32,bytes32) _signature) returns()
func (flareSystemsManager *FlareSystemsManager) TryPackSignUptimeVote(rewardEpochId *big.Int, uptimeVoteHash [32]byte, signature IFlareSystemsManagerSignature) ([]byte, error) {
	return flareSystemsManager.abi.Pack("signUptimeVote", rewardEpochId, uptimeVoteHash, signature)
}

// PackSigningPolicyMinNumberOfVoters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2e3645f8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signingPolicyMinNumberOfVoters() view returns(uint16)
func (flareSystemsManager *FlareSystemsManager) PackSigningPolicyMinNumberOfVoters() []byte {
	enc, err := flareSystemsManager.abi.Pack("signingPolicyMinNumberOfVoters")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSigningPolicyMinNumberOfVoters is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2e3645f8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signingPolicyMinNumberOfVoters() view returns(uint16)
func (flareSystemsManager *FlareSystemsManager) TryPackSigningPolicyMinNumberOfVoters() ([]byte, error) {
	return flareSystemsManager.abi.Pack("signingPolicyMinNumberOfVoters")
}

// UnpackSigningPolicyMinNumberOfVoters is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2e3645f8.
//
// Solidity: function signingPolicyMinNumberOfVoters() view returns(uint16)
func (flareSystemsManager *FlareSystemsManager) UnpackSigningPolicyMinNumberOfVoters(data []byte) (uint16, error) {
	out, err := flareSystemsManager.abi.Unpack("signingPolicyMinNumberOfVoters", data)
	if err != nil {
		return *new(uint16), err
	}
	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	return out0, nil
}

// PackSigningPolicyThresholdPPM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf21d6304.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function signingPolicyThresholdPPM() view returns(uint24)
func (flareSystemsManager *FlareSystemsManager) PackSigningPolicyThresholdPPM() []byte {
	enc, err := flareSystemsManager.abi.Pack("signingPolicyThresholdPPM")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSigningPolicyThresholdPPM is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf21d6304.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function signingPolicyThresholdPPM() view returns(uint24)
func (flareSystemsManager *FlareSystemsManager) TryPackSigningPolicyThresholdPPM() ([]byte, error) {
	return flareSystemsManager.abi.Pack("signingPolicyThresholdPPM")
}

// UnpackSigningPolicyThresholdPPM is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf21d6304.
//
// Solidity: function signingPolicyThresholdPPM() view returns(uint24)
func (flareSystemsManager *FlareSystemsManager) UnpackSigningPolicyThresholdPPM(data []byte) (*big.Int, error) {
	out, err := flareSystemsManager.abi.Unpack("signingPolicyThresholdPPM", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackSubmission is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a759097.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submission() view returns(address)
func (flareSystemsManager *FlareSystemsManager) PackSubmission() []byte {
	enc, err := flareSystemsManager.abi.Pack("submission")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmission is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a759097.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submission() view returns(address)
func (flareSystemsManager *FlareSystemsManager) TryPackSubmission() ([]byte, error) {
	return flareSystemsManager.abi.Pack("submission")
}

// UnpackSubmission is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9a759097.
//
// Solidity: function submission() view returns(address)
func (flareSystemsManager *FlareSystemsManager) UnpackSubmission(data []byte) (common.Address, error) {
	out, err := flareSystemsManager.abi.Unpack("submission", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSubmit3Aligned is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x107d8ffb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submit3Aligned() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) PackSubmit3Aligned() []byte {
	enc, err := flareSystemsManager.abi.Pack("submit3Aligned")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmit3Aligned is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x107d8ffb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submit3Aligned() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) TryPackSubmit3Aligned() ([]byte, error) {
	return flareSystemsManager.abi.Pack("submit3Aligned")
}

// UnpackSubmit3Aligned is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x107d8ffb.
//
// Solidity: function submit3Aligned() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) UnpackSubmit3Aligned(data []byte) (bool, error) {
	out, err := flareSystemsManager.abi.Unpack("submit3Aligned", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackSubmitUptimeVote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dd6850f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submitUptimeVote(uint24 _rewardEpochId, bytes20[] _nodeIds, (uint8,bytes32,bytes32) _signature) returns()
func (flareSystemsManager *FlareSystemsManager) PackSubmitUptimeVote(rewardEpochId *big.Int, nodeIds [][20]byte, signature IFlareSystemsManagerSignature) []byte {
	enc, err := flareSystemsManager.abi.Pack("submitUptimeVote", rewardEpochId, nodeIds, signature)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmitUptimeVote is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9dd6850f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submitUptimeVote(uint24 _rewardEpochId, bytes20[] _nodeIds, (uint8,bytes32,bytes32) _signature) returns()
func (flareSystemsManager *FlareSystemsManager) TryPackSubmitUptimeVote(rewardEpochId *big.Int, nodeIds [][20]byte, signature IFlareSystemsManagerSignature) ([]byte, error) {
	return flareSystemsManager.abi.Pack("submitUptimeVote", rewardEpochId, nodeIds, signature)
}

// PackSubmitUptimeVoteMinDurationBlocks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd8a01a0a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submitUptimeVoteMinDurationBlocks() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) PackSubmitUptimeVoteMinDurationBlocks() []byte {
	enc, err := flareSystemsManager.abi.Pack("submitUptimeVoteMinDurationBlocks")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmitUptimeVoteMinDurationBlocks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd8a01a0a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submitUptimeVoteMinDurationBlocks() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) TryPackSubmitUptimeVoteMinDurationBlocks() ([]byte, error) {
	return flareSystemsManager.abi.Pack("submitUptimeVoteMinDurationBlocks")
}

// UnpackSubmitUptimeVoteMinDurationBlocks is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd8a01a0a.
//
// Solidity: function submitUptimeVoteMinDurationBlocks() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackSubmitUptimeVoteMinDurationBlocks(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("submitUptimeVoteMinDurationBlocks", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackSubmitUptimeVoteMinDurationSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4c528765.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function submitUptimeVoteMinDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) PackSubmitUptimeVoteMinDurationSeconds() []byte {
	enc, err := flareSystemsManager.abi.Pack("submitUptimeVoteMinDurationSeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSubmitUptimeVoteMinDurationSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4c528765.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function submitUptimeVoteMinDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) TryPackSubmitUptimeVoteMinDurationSeconds() ([]byte, error) {
	return flareSystemsManager.abi.Pack("submitUptimeVoteMinDurationSeconds")
}

// UnpackSubmitUptimeVoteMinDurationSeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4c528765.
//
// Solidity: function submitUptimeVoteMinDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackSubmitUptimeVoteMinDurationSeconds(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("submitUptimeVoteMinDurationSeconds", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackSwitchToFallbackMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe22fdece.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) PackSwitchToFallbackMode() []byte {
	enc, err := flareSystemsManager.abi.Pack("switchToFallbackMode")
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
func (flareSystemsManager *FlareSystemsManager) TryPackSwitchToFallbackMode() ([]byte, error) {
	return flareSystemsManager.abi.Pack("switchToFallbackMode")
}

// UnpackSwitchToFallbackMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe22fdece.
//
// Solidity: function switchToFallbackMode() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) UnpackSwitchToFallbackMode(data []byte) (bool, error) {
	out, err := flareSystemsManager.abi.Unpack("switchToFallbackMode", data)
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
func (flareSystemsManager *FlareSystemsManager) PackSwitchToProductionMode() []byte {
	enc, err := flareSystemsManager.abi.Pack("switchToProductionMode")
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
func (flareSystemsManager *FlareSystemsManager) TryPackSwitchToProductionMode() ([]byte, error) {
	return flareSystemsManager.abi.Pack("switchToProductionMode")
}

// PackTimelockedCalls is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x74e6310e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function timelockedCalls(bytes4 selector) view returns(uint256 allowedAfterTimestamp, bytes encodedCall)
func (flareSystemsManager *FlareSystemsManager) PackTimelockedCalls(selector [4]byte) []byte {
	enc, err := flareSystemsManager.abi.Pack("timelockedCalls", selector)
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
func (flareSystemsManager *FlareSystemsManager) TryPackTimelockedCalls(selector [4]byte) ([]byte, error) {
	return flareSystemsManager.abi.Pack("timelockedCalls", selector)
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
func (flareSystemsManager *FlareSystemsManager) UnpackTimelockedCalls(data []byte) (TimelockedCallsOutput, error) {
	out, err := flareSystemsManager.abi.Unpack("timelockedCalls", data)
	outstruct := new(TimelockedCallsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.AllowedAfterTimestamp = abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	outstruct.EncodedCall = *abi.ConvertType(out[1], new([]byte)).(*[]byte)
	return *outstruct, nil
}

// PackTriggerExpirationAndCleanup is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b760d13.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function triggerExpirationAndCleanup() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) PackTriggerExpirationAndCleanup() []byte {
	enc, err := flareSystemsManager.abi.Pack("triggerExpirationAndCleanup")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTriggerExpirationAndCleanup is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b760d13.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function triggerExpirationAndCleanup() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) TryPackTriggerExpirationAndCleanup() ([]byte, error) {
	return flareSystemsManager.abi.Pack("triggerExpirationAndCleanup")
}

// UnpackTriggerExpirationAndCleanup is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9b760d13.
//
// Solidity: function triggerExpirationAndCleanup() view returns(bool)
func (flareSystemsManager *FlareSystemsManager) UnpackTriggerExpirationAndCleanup(data []byte) (bool, error) {
	out, err := flareSystemsManager.abi.Unpack("triggerExpirationAndCleanup", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (flareSystemsManager *FlareSystemsManager) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := flareSystemsManager.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (flareSystemsManager *FlareSystemsManager) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return flareSystemsManager.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackUpdateSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8fbaf860.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateSettings((uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint24,uint16,uint32) _settings) returns()
func (flareSystemsManager *FlareSystemsManager) PackUpdateSettings(settings FlareSystemsManagerSettings) []byte {
	enc, err := flareSystemsManager.abi.Pack("updateSettings", settings)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpdateSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8fbaf860.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function updateSettings((uint16,uint16,uint16,uint8,uint16,uint16,uint16,uint16,uint24,uint16,uint32) _settings) returns()
func (flareSystemsManager *FlareSystemsManager) TryPackUpdateSettings(settings FlareSystemsManagerSettings) ([]byte, error) {
	return flareSystemsManager.abi.Pack("updateSettings", settings)
}

// PackUptimeVoteHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3466911.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function uptimeVoteHash(uint256 rewardEpochId) view returns(bytes32)
func (flareSystemsManager *FlareSystemsManager) PackUptimeVoteHash(rewardEpochId *big.Int) []byte {
	enc, err := flareSystemsManager.abi.Pack("uptimeVoteHash", rewardEpochId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUptimeVoteHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd3466911.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function uptimeVoteHash(uint256 rewardEpochId) view returns(bytes32)
func (flareSystemsManager *FlareSystemsManager) TryPackUptimeVoteHash(rewardEpochId *big.Int) ([]byte, error) {
	return flareSystemsManager.abi.Pack("uptimeVoteHash", rewardEpochId)
}

// UnpackUptimeVoteHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd3466911.
//
// Solidity: function uptimeVoteHash(uint256 rewardEpochId) view returns(bytes32)
func (flareSystemsManager *FlareSystemsManager) UnpackUptimeVoteHash(data []byte) ([32]byte, error) {
	out, err := flareSystemsManager.abi.Unpack("uptimeVoteHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackVoterRegistrationMinDurationBlocks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd10e807f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function voterRegistrationMinDurationBlocks() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) PackVoterRegistrationMinDurationBlocks() []byte {
	enc, err := flareSystemsManager.abi.Pack("voterRegistrationMinDurationBlocks")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVoterRegistrationMinDurationBlocks is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd10e807f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function voterRegistrationMinDurationBlocks() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) TryPackVoterRegistrationMinDurationBlocks() ([]byte, error) {
	return flareSystemsManager.abi.Pack("voterRegistrationMinDurationBlocks")
}

// UnpackVoterRegistrationMinDurationBlocks is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd10e807f.
//
// Solidity: function voterRegistrationMinDurationBlocks() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackVoterRegistrationMinDurationBlocks(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("voterRegistrationMinDurationBlocks", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackVoterRegistrationMinDurationSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa219fe02.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function voterRegistrationMinDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) PackVoterRegistrationMinDurationSeconds() []byte {
	enc, err := flareSystemsManager.abi.Pack("voterRegistrationMinDurationSeconds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVoterRegistrationMinDurationSeconds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa219fe02.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function voterRegistrationMinDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) TryPackVoterRegistrationMinDurationSeconds() ([]byte, error) {
	return flareSystemsManager.abi.Pack("voterRegistrationMinDurationSeconds")
}

// UnpackVoterRegistrationMinDurationSeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa219fe02.
//
// Solidity: function voterRegistrationMinDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackVoterRegistrationMinDurationSeconds(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("voterRegistrationMinDurationSeconds", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackVoterRegistrationTriggerContract is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x88e49ac7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function voterRegistrationTriggerContract() view returns(address)
func (flareSystemsManager *FlareSystemsManager) PackVoterRegistrationTriggerContract() []byte {
	enc, err := flareSystemsManager.abi.Pack("voterRegistrationTriggerContract")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVoterRegistrationTriggerContract is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x88e49ac7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function voterRegistrationTriggerContract() view returns(address)
func (flareSystemsManager *FlareSystemsManager) TryPackVoterRegistrationTriggerContract() ([]byte, error) {
	return flareSystemsManager.abi.Pack("voterRegistrationTriggerContract")
}

// UnpackVoterRegistrationTriggerContract is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x88e49ac7.
//
// Solidity: function voterRegistrationTriggerContract() view returns(address)
func (flareSystemsManager *FlareSystemsManager) UnpackVoterRegistrationTriggerContract(data []byte) (common.Address, error) {
	out, err := flareSystemsManager.abi.Unpack("voterRegistrationTriggerContract", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackVoterRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbe60040e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function voterRegistry() view returns(address)
func (flareSystemsManager *FlareSystemsManager) PackVoterRegistry() []byte {
	enc, err := flareSystemsManager.abi.Pack("voterRegistry")
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
func (flareSystemsManager *FlareSystemsManager) TryPackVoterRegistry() ([]byte, error) {
	return flareSystemsManager.abi.Pack("voterRegistry")
}

// UnpackVoterRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbe60040e.
//
// Solidity: function voterRegistry() view returns(address)
func (flareSystemsManager *FlareSystemsManager) UnpackVoterRegistry(data []byte) (common.Address, error) {
	out, err := flareSystemsManager.abi.Unpack("voterRegistry", data)
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
func (flareSystemsManager *FlareSystemsManager) PackVotingEpochDurationSeconds() []byte {
	enc, err := flareSystemsManager.abi.Pack("votingEpochDurationSeconds")
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
func (flareSystemsManager *FlareSystemsManager) TryPackVotingEpochDurationSeconds() ([]byte, error) {
	return flareSystemsManager.abi.Pack("votingEpochDurationSeconds")
}

// UnpackVotingEpochDurationSeconds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5a832088.
//
// Solidity: function votingEpochDurationSeconds() view returns(uint64)
func (flareSystemsManager *FlareSystemsManager) UnpackVotingEpochDurationSeconds(data []byte) (uint64, error) {
	out, err := flareSystemsManager.abi.Unpack("votingEpochDurationSeconds", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// FlareSystemsManagerClosingExpiredRewardEpochFailed represents a ClosingExpiredRewardEpochFailed event raised by the FlareSystemsManager contract.
type FlareSystemsManagerClosingExpiredRewardEpochFailed struct {
	RewardEpochId *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerClosingExpiredRewardEpochFailedEventName = "ClosingExpiredRewardEpochFailed"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerClosingExpiredRewardEpochFailed) ContractEventName() string {
	return FlareSystemsManagerClosingExpiredRewardEpochFailedEventName
}

// UnpackClosingExpiredRewardEpochFailedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ClosingExpiredRewardEpochFailed(uint24 rewardEpochId)
func (flareSystemsManager *FlareSystemsManager) UnpackClosingExpiredRewardEpochFailedEvent(log *types.Log) (*FlareSystemsManagerClosingExpiredRewardEpochFailed, error) {
	event := "ClosingExpiredRewardEpochFailed"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerClosingExpiredRewardEpochFailed)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the FlareSystemsManager contract.
type FlareSystemsManagerGovernanceCallTimelocked struct {
	Selector              [4]byte
	AllowedAfterTimestamp *big.Int
	EncodedCall           []byte
	Raw                   *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerGovernanceCallTimelocked) ContractEventName() string {
	return FlareSystemsManagerGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes4 selector, uint256 allowedAfterTimestamp, bytes encodedCall)
func (flareSystemsManager *FlareSystemsManager) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*FlareSystemsManagerGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerGovernanceInitialised represents a GovernanceInitialised event raised by the FlareSystemsManager contract.
type FlareSystemsManagerGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerGovernanceInitialised) ContractEventName() string {
	return FlareSystemsManagerGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (flareSystemsManager *FlareSystemsManager) UnpackGovernanceInitialisedEvent(log *types.Log) (*FlareSystemsManagerGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the FlareSystemsManager contract.
type FlareSystemsManagerGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerGovernedProductionModeEntered) ContractEventName() string {
	return FlareSystemsManagerGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (flareSystemsManager *FlareSystemsManager) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*FlareSystemsManagerGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerRandomAcquisitionStarted represents a RandomAcquisitionStarted event raised by the FlareSystemsManager contract.
type FlareSystemsManagerRandomAcquisitionStarted struct {
	RewardEpochId *big.Int
	Timestamp     uint64
	Raw           *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerRandomAcquisitionStartedEventName = "RandomAcquisitionStarted"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerRandomAcquisitionStarted) ContractEventName() string {
	return FlareSystemsManagerRandomAcquisitionStartedEventName
}

// UnpackRandomAcquisitionStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RandomAcquisitionStarted(uint24 indexed rewardEpochId, uint64 timestamp)
func (flareSystemsManager *FlareSystemsManager) UnpackRandomAcquisitionStartedEvent(log *types.Log) (*FlareSystemsManagerRandomAcquisitionStarted, error) {
	event := "RandomAcquisitionStarted"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerRandomAcquisitionStarted)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerRewardEpochStarted represents a RewardEpochStarted event raised by the FlareSystemsManager contract.
type FlareSystemsManagerRewardEpochStarted struct {
	RewardEpochId      *big.Int
	StartVotingRoundId uint32
	Timestamp          uint64
	Raw                *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerRewardEpochStartedEventName = "RewardEpochStarted"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerRewardEpochStarted) ContractEventName() string {
	return FlareSystemsManagerRewardEpochStartedEventName
}

// UnpackRewardEpochStartedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RewardEpochStarted(uint24 indexed rewardEpochId, uint32 startVotingRoundId, uint64 timestamp)
func (flareSystemsManager *FlareSystemsManager) UnpackRewardEpochStartedEvent(log *types.Log) (*FlareSystemsManagerRewardEpochStarted, error) {
	event := "RewardEpochStarted"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerRewardEpochStarted)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerRewardsSigned represents a RewardsSigned event raised by the FlareSystemsManager contract.
type FlareSystemsManagerRewardsSigned struct {
	RewardEpochId         *big.Int
	SigningPolicyAddress  common.Address
	Voter                 common.Address
	RewardsHash           [32]byte
	NoOfWeightBasedClaims []IFlareSystemsManagerNumberOfWeightBasedClaims
	Timestamp             uint64
	ThresholdReached      bool
	Raw                   *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerRewardsSignedEventName = "RewardsSigned"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerRewardsSigned) ContractEventName() string {
	return FlareSystemsManagerRewardsSignedEventName
}

// UnpackRewardsSignedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event RewardsSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 rewardsHash, (uint256,uint256)[] noOfWeightBasedClaims, uint64 timestamp, bool thresholdReached)
func (flareSystemsManager *FlareSystemsManager) UnpackRewardsSignedEvent(log *types.Log) (*FlareSystemsManagerRewardsSigned, error) {
	event := "RewardsSigned"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerRewardsSigned)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerSettingCleanUpBlockNumberFailed represents a SettingCleanUpBlockNumberFailed event raised by the FlareSystemsManager contract.
type FlareSystemsManagerSettingCleanUpBlockNumberFailed struct {
	BlockNumber uint64
	Raw         *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerSettingCleanUpBlockNumberFailedEventName = "SettingCleanUpBlockNumberFailed"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerSettingCleanUpBlockNumberFailed) ContractEventName() string {
	return FlareSystemsManagerSettingCleanUpBlockNumberFailedEventName
}

// UnpackSettingCleanUpBlockNumberFailedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SettingCleanUpBlockNumberFailed(uint64 blockNumber)
func (flareSystemsManager *FlareSystemsManager) UnpackSettingCleanUpBlockNumberFailedEvent(log *types.Log) (*FlareSystemsManagerSettingCleanUpBlockNumberFailed, error) {
	event := "SettingCleanUpBlockNumberFailed"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerSettingCleanUpBlockNumberFailed)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerSignUptimeVoteEnabled represents a SignUptimeVoteEnabled event raised by the FlareSystemsManager contract.
type FlareSystemsManagerSignUptimeVoteEnabled struct {
	RewardEpochId *big.Int
	Timestamp     uint64
	Raw           *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerSignUptimeVoteEnabledEventName = "SignUptimeVoteEnabled"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerSignUptimeVoteEnabled) ContractEventName() string {
	return FlareSystemsManagerSignUptimeVoteEnabledEventName
}

// UnpackSignUptimeVoteEnabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SignUptimeVoteEnabled(uint24 indexed rewardEpochId, uint64 timestamp)
func (flareSystemsManager *FlareSystemsManager) UnpackSignUptimeVoteEnabledEvent(log *types.Log) (*FlareSystemsManagerSignUptimeVoteEnabled, error) {
	event := "SignUptimeVoteEnabled"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerSignUptimeVoteEnabled)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerSigningPolicySigned represents a SigningPolicySigned event raised by the FlareSystemsManager contract.
type FlareSystemsManagerSigningPolicySigned struct {
	RewardEpochId        *big.Int
	SigningPolicyAddress common.Address
	Voter                common.Address
	Timestamp            uint64
	ThresholdReached     bool
	Raw                  *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerSigningPolicySignedEventName = "SigningPolicySigned"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerSigningPolicySigned) ContractEventName() string {
	return FlareSystemsManagerSigningPolicySignedEventName
}

// UnpackSigningPolicySignedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SigningPolicySigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, uint64 timestamp, bool thresholdReached)
func (flareSystemsManager *FlareSystemsManager) UnpackSigningPolicySignedEvent(log *types.Log) (*FlareSystemsManagerSigningPolicySigned, error) {
	event := "SigningPolicySigned"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerSigningPolicySigned)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the FlareSystemsManager contract.
type FlareSystemsManagerTimelockedGovernanceCallCanceled struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerTimelockedGovernanceCallCanceled) ContractEventName() string {
	return FlareSystemsManagerTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes4 selector, uint256 timestamp)
func (flareSystemsManager *FlareSystemsManager) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*FlareSystemsManagerTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the FlareSystemsManager contract.
type FlareSystemsManagerTimelockedGovernanceCallExecuted struct {
	Selector  [4]byte
	Timestamp *big.Int
	Raw       *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerTimelockedGovernanceCallExecuted) ContractEventName() string {
	return FlareSystemsManagerTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes4 selector, uint256 timestamp)
func (flareSystemsManager *FlareSystemsManager) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*FlareSystemsManagerTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerTriggeringVoterRegistrationFailed represents a TriggeringVoterRegistrationFailed event raised by the FlareSystemsManager contract.
type FlareSystemsManagerTriggeringVoterRegistrationFailed struct {
	RewardEpochId *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerTriggeringVoterRegistrationFailedEventName = "TriggeringVoterRegistrationFailed"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerTriggeringVoterRegistrationFailed) ContractEventName() string {
	return FlareSystemsManagerTriggeringVoterRegistrationFailedEventName
}

// UnpackTriggeringVoterRegistrationFailedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TriggeringVoterRegistrationFailed(uint24 rewardEpochId)
func (flareSystemsManager *FlareSystemsManager) UnpackTriggeringVoterRegistrationFailedEvent(log *types.Log) (*FlareSystemsManagerTriggeringVoterRegistrationFailed, error) {
	event := "TriggeringVoterRegistrationFailed"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerTriggeringVoterRegistrationFailed)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerUptimeVoteSigned represents a UptimeVoteSigned event raised by the FlareSystemsManager contract.
type FlareSystemsManagerUptimeVoteSigned struct {
	RewardEpochId        *big.Int
	SigningPolicyAddress common.Address
	Voter                common.Address
	UptimeVoteHash       [32]byte
	Timestamp            uint64
	ThresholdReached     bool
	Raw                  *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerUptimeVoteSignedEventName = "UptimeVoteSigned"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerUptimeVoteSigned) ContractEventName() string {
	return FlareSystemsManagerUptimeVoteSignedEventName
}

// UnpackUptimeVoteSignedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UptimeVoteSigned(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes32 uptimeVoteHash, uint64 timestamp, bool thresholdReached)
func (flareSystemsManager *FlareSystemsManager) UnpackUptimeVoteSignedEvent(log *types.Log) (*FlareSystemsManagerUptimeVoteSigned, error) {
	event := "UptimeVoteSigned"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerUptimeVoteSigned)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerUptimeVoteSubmitted represents a UptimeVoteSubmitted event raised by the FlareSystemsManager contract.
type FlareSystemsManagerUptimeVoteSubmitted struct {
	RewardEpochId        *big.Int
	SigningPolicyAddress common.Address
	Voter                common.Address
	NodeIds              [][20]byte
	Timestamp            uint64
	Raw                  *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerUptimeVoteSubmittedEventName = "UptimeVoteSubmitted"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerUptimeVoteSubmitted) ContractEventName() string {
	return FlareSystemsManagerUptimeVoteSubmittedEventName
}

// UnpackUptimeVoteSubmittedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event UptimeVoteSubmitted(uint24 indexed rewardEpochId, address indexed signingPolicyAddress, address indexed voter, bytes20[] nodeIds, uint64 timestamp)
func (flareSystemsManager *FlareSystemsManager) UnpackUptimeVoteSubmittedEvent(log *types.Log) (*FlareSystemsManagerUptimeVoteSubmitted, error) {
	event := "UptimeVoteSubmitted"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerUptimeVoteSubmitted)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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

// FlareSystemsManagerVotePowerBlockSelected represents a VotePowerBlockSelected event raised by the FlareSystemsManager contract.
type FlareSystemsManagerVotePowerBlockSelected struct {
	RewardEpochId  *big.Int
	VotePowerBlock uint64
	Timestamp      uint64
	Raw            *types.Log // Blockchain specific contextual infos
}

const FlareSystemsManagerVotePowerBlockSelectedEventName = "VotePowerBlockSelected"

// ContractEventName returns the user-defined event name.
func (FlareSystemsManagerVotePowerBlockSelected) ContractEventName() string {
	return FlareSystemsManagerVotePowerBlockSelectedEventName
}

// UnpackVotePowerBlockSelectedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event VotePowerBlockSelected(uint24 indexed rewardEpochId, uint64 votePowerBlock, uint64 timestamp)
func (flareSystemsManager *FlareSystemsManager) UnpackVotePowerBlockSelectedEvent(log *types.Log) (*FlareSystemsManagerVotePowerBlockSelected, error) {
	event := "VotePowerBlockSelected"
	if len(log.Topics) == 0 || log.Topics[0] != flareSystemsManager.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(FlareSystemsManagerVotePowerBlockSelected)
	if len(log.Data) > 0 {
		if err := flareSystemsManager.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range flareSystemsManager.abi.Events[event].Inputs {
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
func (flareSystemsManager *FlareSystemsManager) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], flareSystemsManager.abi.Errors["ECDSAInvalidSignature"].ID.Bytes()[:4]) {
		return flareSystemsManager.UnpackECDSAInvalidSignatureError(raw[4:])
	}
	if bytes.Equal(raw[:4], flareSystemsManager.abi.Errors["ECDSAInvalidSignatureLength"].ID.Bytes()[:4]) {
		return flareSystemsManager.UnpackECDSAInvalidSignatureLengthError(raw[4:])
	}
	if bytes.Equal(raw[:4], flareSystemsManager.abi.Errors["ECDSAInvalidSignatureS"].ID.Bytes()[:4]) {
		return flareSystemsManager.UnpackECDSAInvalidSignatureSError(raw[4:])
	}
	if bytes.Equal(raw[:4], flareSystemsManager.abi.Errors["SafeCastOverflowedUintDowncast"].ID.Bytes()[:4]) {
		return flareSystemsManager.UnpackSafeCastOverflowedUintDowncastError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// FlareSystemsManagerECDSAInvalidSignature represents a ECDSAInvalidSignature error raised by the FlareSystemsManager contract.
type FlareSystemsManagerECDSAInvalidSignature struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignature()
func FlareSystemsManagerECDSAInvalidSignatureErrorID() common.Hash {
	return common.HexToHash("0xf645eedf0193584640b6b90cb9477e4c95b98636c148a891d4c0a146dc46e75a")
}

// UnpackECDSAInvalidSignatureError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignature()
func (flareSystemsManager *FlareSystemsManager) UnpackECDSAInvalidSignatureError(raw []byte) (*FlareSystemsManagerECDSAInvalidSignature, error) {
	out := new(FlareSystemsManagerECDSAInvalidSignature)
	if err := flareSystemsManager.abi.UnpackIntoInterface(out, "ECDSAInvalidSignature", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// FlareSystemsManagerECDSAInvalidSignatureLength represents a ECDSAInvalidSignatureLength error raised by the FlareSystemsManager contract.
type FlareSystemsManagerECDSAInvalidSignatureLength struct {
	Length *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func FlareSystemsManagerECDSAInvalidSignatureLengthErrorID() common.Hash {
	return common.HexToHash("0xfce698f7e8e5342cd615f641317bc45fe7e1e4a8b0a14dd1383ff8dc9c41917f")
}

// UnpackECDSAInvalidSignatureLengthError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureLength(uint256 length)
func (flareSystemsManager *FlareSystemsManager) UnpackECDSAInvalidSignatureLengthError(raw []byte) (*FlareSystemsManagerECDSAInvalidSignatureLength, error) {
	out := new(FlareSystemsManagerECDSAInvalidSignatureLength)
	if err := flareSystemsManager.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureLength", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// FlareSystemsManagerECDSAInvalidSignatureS represents a ECDSAInvalidSignatureS error raised by the FlareSystemsManager contract.
type FlareSystemsManagerECDSAInvalidSignatureS struct {
	S [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func FlareSystemsManagerECDSAInvalidSignatureSErrorID() common.Hash {
	return common.HexToHash("0xd78bce0cccb935155ed6428d1c13e50b7f3550fd2b66b9fe266006fea4a5e1eb")
}

// UnpackECDSAInvalidSignatureSError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ECDSAInvalidSignatureS(bytes32 s)
func (flareSystemsManager *FlareSystemsManager) UnpackECDSAInvalidSignatureSError(raw []byte) (*FlareSystemsManagerECDSAInvalidSignatureS, error) {
	out := new(FlareSystemsManagerECDSAInvalidSignatureS)
	if err := flareSystemsManager.abi.UnpackIntoInterface(out, "ECDSAInvalidSignatureS", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// FlareSystemsManagerSafeCastOverflowedUintDowncast represents a SafeCastOverflowedUintDowncast error raised by the FlareSystemsManager contract.
type FlareSystemsManagerSafeCastOverflowedUintDowncast struct {
	Bits  uint8
	Value *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func FlareSystemsManagerSafeCastOverflowedUintDowncastErrorID() common.Hash {
	return common.HexToHash("0x6dfcc6503a32754ce7a89698e18201fc5294fd4aad43edefee786f88423b1a12")
}

// UnpackSafeCastOverflowedUintDowncastError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SafeCastOverflowedUintDowncast(uint8 bits, uint256 value)
func (flareSystemsManager *FlareSystemsManager) UnpackSafeCastOverflowedUintDowncastError(raw []byte) (*FlareSystemsManagerSafeCastOverflowedUintDowncast, error) {
	out := new(FlareSystemsManagerSafeCastOverflowedUintDowncast)
	if err := flareSystemsManager.abi.UnpackIntoInterface(out, "SafeCastOverflowedUintDowncast", raw); err != nil {
		return nil, err
	}
	return out, nil
}
