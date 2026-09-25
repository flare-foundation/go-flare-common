// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletpayments

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

// IBtcAccountConfiguredAnchor is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountConfiguredAnchor struct {
	GenesisAnchorTxid [32]byte
	GenesisAnchorVout uint32
}

// IBtcAccountConfiguredProof is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountConfiguredProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IBtcAccountConfiguredRequestBody
	ResponseBody IBtcAccountConfiguredResponseBody
}

// IBtcAccountConfiguredRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountConfiguredRequestBody struct {
	AccountIndex uint32
	PublicKeys   [][]byte
	Threshold    uint64
	Anchors      []IBtcAccountConfiguredAnchor
}

// IBtcAccountConfiguredResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountConfiguredResponseBody struct {
	Status         uint8
	AccountAddress string
}

// IBtcAccountsBtcAnchor is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountsBtcAnchor struct {
	GenesisTxid [32]byte
	GenesisVout uint32
	NextNonce   uint64
}

// IBtcAccountsBtcAttemptCommitment is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountsBtcAttemptCommitment struct {
	Txid           [32]byte
	NextAnchorTxid [32]byte
	NextAnchorVout uint32
}

// IBtcAccountsBtcProposalCommitment is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountsBtcProposalCommitment struct {
	AnchorIndex    uint32
	Nonce          uint64
	NextAnchorVout uint32
	Txid           [32]byte
	NextAnchorTxid [32]byte
}

// IBtcEscrowsBtcEscrowTerms is an auto generated low-level Go binding around an user-defined struct.
type IBtcEscrowsBtcEscrowTerms struct {
	PreimageHash       [32]byte
	Amount             uint64
	ExpiresAt          uint64
	CounterpartyPubKey []byte
}

// IBtcParamsBtcConsensusParams is an auto generated low-level Go binding around an user-defined struct.
type IBtcParamsBtcConsensusParams struct {
	RuleVersion             uint16
	MinPaymentAmountSat     uint64
	BatchDeadlineSeconds    uint64
	DustRelayFeeSatPerKvB   uint32
	MaxChangeOutputs        uint8
	ScoreImprovementBias    uint8
	ScoreImprovementClamp   uint8
	ScoreFreeChangeOutputs  uint8
	MaxSweptInputs          uint16
	MinFundingConfirmations uint32
	AnchorSpendDepth        uint32
	FeeRateFloorBIPS        uint32
	FeeRateCeilingBIPS      uint32
	MaxOutputsPerPayment    uint32
}

// IBtcParamsBtcWalletParams is an auto generated low-level Go binding around an user-defined struct.
type IBtcParamsBtcWalletParams struct {
	BatchDeadlineSeconds   uint64
	MaxOutputsPerPayment   uint32
	MaxChangeOutputs       uint8
	ScoreFreeChangeOutputs uint8
}

// ICspInstructionsAttempt is an auto generated low-level Go binding around an user-defined struct.
type ICspInstructionsAttempt struct {
	PackageHash         [32]byte
	ChainCommitmentHash [32]byte
	Proposer            common.Address
	SettledAt           uint64
	MaxFee              uint64
}

// ICspInstructionsEligible is an auto generated low-level Go binding around an user-defined struct.
type ICspInstructionsEligible struct {
	SequencePosition uint64
	Attempt          uint32
	Generation       uint64
}

// ICspInstructionsInstructionRecord is an auto generated low-level Go binding around an user-defined struct.
type ICspInstructionsInstructionRecord struct {
	Kind          uint8
	Lane          uint32
	EmittedAt     uint64
	FromPaymentId uint64
	ToPaymentId   uint64
	Settled       bool
}

// ICspInstructionsLaneUse is an auto generated low-level Go binding around an user-defined struct.
type ICspInstructionsLaneUse struct {
	Used             bool
	SequencePosition uint64
}

// ICspInstructionsPendingReset is an auto generated low-level Go binding around an user-defined struct.
type ICspInstructionsPendingReset struct {
	SequencePosition uint64
	Attempt          uint32
}

// ICspProposalCheckProof is an auto generated low-level Go binding around an user-defined struct.
type ICspProposalCheckProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  ICspProposalCheckRequestBody
	ResponseBody ICspProposalCheckResponseBody
}

// ICspProposalCheckRequestBody is an auto generated low-level Go binding around an user-defined struct.
type ICspProposalCheckRequestBody struct {
	WalletRegistry     common.Address
	WalletId           [32]byte
	AccountIndex       uint32
	SequencePosition   uint64
	Attempt            uint32
	EligibleGeneration uint64
	PackageHash        [32]byte
}

// ICspProposalCheckResponseBody is an auto generated low-level Go binding around an user-defined struct.
type ICspProposalCheckResponseBody struct {
	ProposerAddress     common.Address
	Score               uint64
	PaymentCount        uint32
	ChainCommitmentHash [32]byte
}

// ICspProposalsCspSourceSettings is an auto generated low-level Go binding around an user-defined struct.
type ICspProposalsCspSourceSettings struct {
	FinalizationGraceSeconds uint64
	RequiredTeeSignatures    uint16
	ProposalThresholdBIPS    uint16
	ProposerRedundancy       uint8
	ProposerFeeWei           *big.Int
}

// ICspProposalsLeader is an auto generated low-level Go binding around an user-defined struct.
type ICspProposalsLeader struct {
	Exists              bool
	SequencePosition    uint64
	Attempt             uint32
	Lane                uint32
	GraceEndsAt         uint64
	PaymentCount        uint32
	PackageHash         [32]byte
	ChainCommitmentHash [32]byte
	Proposer            common.Address
	Score               uint64
}

// ICspQueueQueuedPayment is an auto generated low-level Go binding around an user-defined struct.
type ICspQueueQueuedPayment struct {
	RecipientAddress string
	Amount           uint64
	MaxFee           uint64
	PaymentReference [32]byte
}

// ICspQueueSettledBatch is an auto generated low-level Go binding around an user-defined struct.
type ICspQueueSettledBatch struct {
	FromPaymentId    uint64
	ToPaymentId      uint64
	SequencePosition uint64
}

// IDiamondFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// IDiamondLoupeFacet is an auto generated low-level Go binding around an user-defined struct.
type IDiamondLoupeFacet struct {
	FacetAddress      common.Address
	FunctionSelectors [][4]byte
}

// IEscrowsEscrowTerms is an auto generated low-level Go binding around an user-defined struct.
type IEscrowsEscrowTerms struct {
	PreimageHash [32]byte
	Amount       *big.Int
	ExpiresAt    uint64
}

// IFdc2HubFdc2ResponseHeader is an auto generated low-level Go binding around an user-defined struct.
type IFdc2HubFdc2ResponseHeader struct {
	AttestationType    [32]byte
	SourceId           [32]byte
	ThresholdBIPS      uint16
	ProofOwner         common.Address
	Cosigners          []common.Address
	CosignersThreshold uint64
	Timestamp          uint64
}

// IFdc2VerificationFdc2Signatures is an auto generated low-level Go binding around an user-defined struct.
type IFdc2VerificationFdc2Signatures struct {
	SigningPolicySignatures []byte
	TeeSignatures           []Signature
	CosignerSignatures      []Signature
}

// IFeeSchedulesFeeSchedule is an auto generated low-level Go binding around an user-defined struct.
type IFeeSchedulesFeeSchedule struct {
	FactorBIPS   int16
	DelaySeconds uint16
}

// IFeeSchedulesFeeScheduleConfig is an auto generated low-level Go binding around an user-defined struct.
type IFeeSchedulesFeeScheduleConfig struct {
	MaxSchedules    uint8
	MaxDelaySeconds uint16
}

// IFeeSchedulesFeeScheduleConfigInput is an auto generated low-level Go binding around an user-defined struct.
type IFeeSchedulesFeeScheduleConfigInput struct {
	MaxDelaySeconds uint16
	MaxSchedules    uint8
	SourceId        [32]byte
}

// INativeNonceAccountConfiguredProof is an auto generated low-level Go binding around an user-defined struct.
type INativeNonceAccountConfiguredProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  INativeNonceAccountConfiguredRequestBody
	ResponseBody INativeNonceAccountConfiguredResponseBody
}

// INativeNonceAccountConfiguredRequestBody is an auto generated low-level Go binding around an user-defined struct.
type INativeNonceAccountConfiguredRequestBody struct {
	AccountAddress string
	PublicKeys     [][]byte
	Threshold      uint64
}

// INativeNonceAccountConfiguredResponseBody is an auto generated low-level Go binding around an user-defined struct.
type INativeNonceAccountConfiguredResponseBody struct {
	Status   uint8
	Sequence uint64
}

// ISourceConfigSource is an auto generated low-level Go binding around an user-defined struct.
type ISourceConfigSource struct {
	KeyType          [32]byte
	OpType           [32]byte
	PaymentModel     uint8
	ChainKind        uint8
	Registered       bool
	Enabled          bool
	CustodianEscrows bool
	TeeEscrows       bool
}

// ISourceConfigSourceRegistration is an auto generated low-level Go binding around an user-defined struct.
type ISourceConfigSourceRegistration struct {
	SourceId     [32]byte
	KeyType      [32]byte
	OpType       [32]byte
	PaymentModel uint8
	ChainKind    uint8
}

// IXrpEscrowsXrpEscrowTerms is an auto generated low-level Go binding around an user-defined struct.
type IXrpEscrowsXrpEscrowTerms struct {
	PreimageHash [32]byte
	Amount       *big.Int
	ExpiresAt    uint64
	Destination  string
	Nullified    bool
}

// PaymentInstruction is an auto generated low-level Go binding around an user-defined struct.
type PaymentInstruction struct {
	RecipientAddress string
	TokenId          []byte
	Amount           *big.Int
	MaxFee           *big.Int
	PaymentReference [32]byte
}

// ReissueFeeParams is an auto generated low-level Go binding around an user-defined struct.
type ReissueFeeParams struct {
	MaxFeePerPayment      []*big.Int
	FactorsBIPSPerPayment [][]int16
	DelaysSeconds         []uint16
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// WalletAccount is an auto generated low-level Go binding around an user-defined struct.
type WalletAccount struct {
	SourceId       [32]byte
	AccountAddress string
}

// WalletPaymentsMetaData contains all meta data concerning the WalletPayments contract.
var WalletPaymentsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccountAddressAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountIndexAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountIndexMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorIndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorNonceRegression\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorSetEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AttemptSuperseded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuthorizationAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BtcConsensusParamsNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BtcEscrowAlreadyReclaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BtcEscrowNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BtcEscrowNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BtcEscrowNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotNullifyPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitmentMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CspAccountNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelay\",\"type\":\"uint256\"}],\"name\":\"DelayTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateAnchor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyProposerList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyScheduleNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EscrowProfileNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EscrowsNotEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"FeeScheduleConfigNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeScheduleNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeSchedulesUnsupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GraceClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GraceOpen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"}],\"name\":\"InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientTeeSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBtcConsensusParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBtcEscrowProfile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBtcEscrowTerms\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBtcWalletParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCspSourceSettings\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeFactor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"}],\"name\":\"InvalidFeeScheduleConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGrace\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentInstructionCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestBody\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidXrpEscrowTerms\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LaneMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeaderExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNewAnchors\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonceMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEligible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLaneTip\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToSettle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullifiedPaymentsUnsupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProjectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationAlreadyNullified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverrideAlreadyPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentRangeInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposerNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReissueDeclarationMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceChainKeyTypeMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceIdZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceKeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceKeyTypeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceLimitsNotConfigured\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceModelChainMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceOpTypeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourcePaymentModelUnknown\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredAnchorsChanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockCallNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockValueNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenIdUnsupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySchedules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"UnknownWalletRegistry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"UnsupportedPaymentModel\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UsePaymentReissue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotExpected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"expectedNetwork\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"actualNetwork\",\"type\":\"uint8\"}],\"name\":\"WalletNetworkMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XrpEscrowNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIBtcParams.BtcWalletParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"AccountBtcWalletParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"}],\"name\":\"AccountFeeScheduleCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIFeeSchedules.FeeSchedule[]\",\"name\":\"schedule\",\"type\":\"tuple[]\"}],\"name\":\"AccountFeeScheduleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"proposers\",\"type\":\"address[]\"}],\"name\":\"AccountProposersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"walletRegistry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"anchorCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"}],\"name\":\"BtcAccountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"walletRegistry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"anchorCount\",\"type\":\"uint32\"}],\"name\":\"BtcAnchorsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"anchorIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"}],\"name\":\"BtcAttemptSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ruleVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minPaymentAmountSat\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"dustRelayFeeSatPerKvB\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreImprovementBias\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreImprovementClamp\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"maxSweptInputs\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"minFundingConfirmations\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"anchorSpendDepth\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeRateFloorBIPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeRateCeilingBIPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIBtcParams.BtcConsensusParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"BtcConsensusParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumICspInstructions.InstructionKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"}],\"name\":\"BtcEscrowEmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"counterpartyPubKey\",\"type\":\"bytes\"}],\"name\":\"BtcEscrowProfileSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"createPaymentId\",\"type\":\"uint64\"}],\"name\":\"BtcEscrowReclaimEmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"CspFeeForwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"finalizationGraceSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"requiredTeeSignatures\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"proposalThresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"proposerRedundancy\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"proposerFeeWei\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structICspProposals.CspSourceSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"name\":\"CspSourceSettingsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"walletRegistry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"CustodianInstructionIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamond.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamond.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"generation\",\"type\":\"uint64\"}],\"name\":\"EligibleAdvanced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"FeeScheduleConfigsCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structIFeeSchedules.FeeScheduleConfigInput[]\",\"name\":\"configs\",\"type\":\"tuple[]\"}],\"name\":\"FeeScheduleConfigsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumICspInstructions.InstructionKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"emittedAt\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"}],\"name\":\"InstructionEmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"InstructionSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"walletRegistry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"initialNonce\",\"type\":\"uint64\"}],\"name\":\"NativeNonceAccountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"}],\"name\":\"OperationNullified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"PaymentBatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"name\":\"PaymentQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIBtcParams.BtcWalletParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ProjectBtcWalletParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"ProjectFeeScheduleCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIFeeSchedules.FeeSchedule[]\",\"name\":\"schedule\",\"type\":\"tuple[]\"}],\"name\":\"ProjectFeeScheduleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"proposers\",\"type\":\"address[]\"}],\"name\":\"ProjectProposersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"}],\"name\":\"ProposalContended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"graceEndsAt\",\"type\":\"uint64\"}],\"name\":\"ProposalLeading\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"ProposerUrlSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"nullifiedPaymentIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"}],\"name\":\"ReissueEmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"name\":\"SigningRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"custodianWallets\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"teeWallets\",\"type\":\"bool\"}],\"name\":\"SourceEscrowsEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"sourceIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SourcesEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"enumPaymentModel\",\"name\":\"paymentModel\",\"type\":\"uint8\"},{\"internalType\":\"enumISourceConfig.ChainKind\",\"name\":\"chainKind\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structISourceConfig.SourceRegistration[]\",\"name\":\"registrations\",\"type\":\"tuple[]\"}],\"name\":\"SourcesRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reissueNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"XrpEscrowInstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reissueNumber\",\"type\":\"uint64\"}],\"name\":\"XrpEscrowNullified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"}],\"name\":\"XrpEscrowProfileSet\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccountConfigured.Anchor[]\",\"name\":\"anchors\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBtcAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIBtcAccountConfigured.BtcAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structIBtcAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIBtcAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"addAnchors\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccountConfigured.Anchor[]\",\"name\":\"anchors\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBtcAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIBtcAccountConfigured.BtcAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structIBtcAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIBtcAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"addBtcAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structINativeNonceAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumINativeNonceAccountConfigured.NativeNonceAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structINativeNonceAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structINativeNonceAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"addNativeNonceAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"clearAccountFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"clearFeeScheduleConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"clearProjectFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"enumICspInstructions.EmissionMode\",\"name\":\"_mode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_maxFee\",\"type\":\"uint256\"}],\"name\":\"consolidate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"}],\"internalType\":\"structIEscrows.EscrowTerms\",\"name\":\"_terms\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_maxFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"createEscrow\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"custodianWalletManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamond.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamond.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"facetFunctionSelectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2Hub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2RequestFeeConfigurations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2Verification\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"eligibleGeneration\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"}],\"internalType\":\"structICspProposalCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"proposerAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"paymentCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"}],\"internalType\":\"structICspProposalCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structICspProposalCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"anchorIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"}],\"internalType\":\"structIBtcAccounts.BtcProposalCommitment\",\"name\":\"_commitment\",\"type\":\"tuple\"}],\"name\":\"finalizeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAccountBtcWalletParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"}],\"internalType\":\"structIBtcParams.BtcWalletParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAccountFeeSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structIFeeSchedules.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAccountIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"_accountIndex\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAccountProposers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_proposers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"_anchorIndex\",\"type\":\"uint32\"}],\"name\":\"getAnchor\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisVout\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nextNonce\",\"type\":\"uint64\"}],\"internalType\":\"structIBtcAccounts.BtcAnchor\",\"name\":\"_anchor\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAnchorCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"_anchorCount\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getAttempts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"settledAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"}],\"internalType\":\"structICspInstructions.Attempt[]\",\"name\":\"_attempts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getBatchPaymentId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchPaymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getBtcAttemptCommitments\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccounts.BtcAttemptCommitment[]\",\"name\":\"_commitments\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getBtcAttempts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"settledAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"}],\"internalType\":\"structICspInstructions.Attempt[]\",\"name\":\"_attempts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccounts.BtcAttemptCommitment[]\",\"name\":\"_commitments\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getBtcConsensusParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ruleVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minPaymentAmountSat\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"dustRelayFeeSatPerKvB\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreImprovementBias\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreImprovementClamp\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"maxSweptInputs\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"minFundingConfirmations\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"anchorSpendDepth\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeRateFloorBIPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeRateCeilingBIPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcParams.BtcConsensusParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getBtcEscrow\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"counterpartyPubKey\",\"type\":\"bytes\"}],\"internalType\":\"structIBtcEscrows.BtcEscrowTerms\",\"name\":\"_terms\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getBtcEscrowProfile\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_counterpartyPubKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_createPaymentId\",\"type\":\"uint64\"}],\"name\":\"getBtcEscrowReclaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_reclaimed\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_reclaimPaymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getBtcEscrowTerms\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"counterpartyPubKey\",\"type\":\"bytes\"}],\"internalType\":\"structIBtcEscrows.BtcEscrowTerms\",\"name\":\"_terms\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getBtcInstructionNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_generation\",\"type\":\"uint64\"}],\"name\":\"getBtcLeader\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lane\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"graceEndsAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"paymentCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"}],\"internalType\":\"structICspProposals.Leader\",\"name\":\"_leader\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"anchorIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"}],\"internalType\":\"structIBtcAccounts.BtcProposalCommitment\",\"name\":\"_commitment\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getBtcParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ruleVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minPaymentAmountSat\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"dustRelayFeeSatPerKvB\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreImprovementBias\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreImprovementClamp\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"maxSweptInputs\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"minFundingConfirmations\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"anchorSpendDepth\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeRateFloorBIPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeRateCeilingBIPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcParams.BtcConsensusParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_generation\",\"type\":\"uint64\"}],\"name\":\"getBtcProposalCommitment\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"anchorIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"}],\"internalType\":\"structIBtcAccounts.BtcProposalCommitment\",\"name\":\"_commitment\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_accountIndex\",\"type\":\"uint32\"}],\"name\":\"getCspAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getCspSourceSettings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"finalizationGraceSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"requiredTeeSignatures\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"proposalThresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"proposerRedundancy\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"proposerFeeWei\",\"type\":\"uint128\"}],\"internalType\":\"structICspProposals.CspSourceSettings\",\"name\":\"_settings\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_accountHash\",\"type\":\"bytes32\"}],\"name\":\"getEffectiveSchedule\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_feeSchedule\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getEligible\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"generation\",\"type\":\"uint64\"}],\"internalType\":\"structICspInstructions.Eligible\",\"name\":\"_eligible\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getEscrow\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"}],\"internalType\":\"structIEscrows.EscrowTerms\",\"name\":\"_terms\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getFeeScheduleConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structIFeeSchedules.FeeScheduleConfig\",\"name\":\"_config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_generation\",\"type\":\"uint64\"}],\"name\":\"getFinalizedHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_packageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getInitialNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_initialNonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getInstruction\",\"outputs\":[{\"components\":[{\"internalType\":\"enumICspInstructions.InstructionKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lane\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"emittedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fromPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"internalType\":\"structICspInstructions.InstructionRecord\",\"name\":\"_instruction\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"_lane\",\"type\":\"uint32\"}],\"name\":\"getLaneLastSettled\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"used\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"}],\"internalType\":\"structICspInstructions.LaneUse\",\"name\":\"_use\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_generation\",\"type\":\"uint64\"}],\"name\":\"getLeader\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lane\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"graceEndsAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"paymentCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"}],\"internalType\":\"structICspProposals.Leader\",\"name\":\"_leader\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNetwork\",\"outputs\":[{\"internalType\":\"enumISourceConfig.Network\",\"name\":\"_network\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getNextPaymentId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nextPaymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getNextSequencePosition\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nextSequencePosition\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getNextUnconsumedPaymentId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nextUnconsumedPaymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getNullifiedPaymentIds\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"_nullifiedPaymentIds\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getOperationSequencePosition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isOperation\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"}],\"name\":\"getPaymentFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getPaymentHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_paymentHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getPendingResets\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"}],\"internalType\":\"structICspInstructions.PendingReset[]\",\"name\":\"_pendingResets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getProjectBtcWalletParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"}],\"internalType\":\"structIBtcParams.BtcWalletParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getProjectFeeSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structIFeeSchedules.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getProjectProposers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_proposers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"}],\"name\":\"getProposerUrl\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getQueuedPayment\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structICspQueue.QueuedPayment\",\"name\":\"_payment\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredSourceIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_sourceIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getSettledBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"fromPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"}],\"internalType\":\"structICspQueue.SettledBatch\",\"name\":\"_batch\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getSettledBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"fromPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"}],\"internalType\":\"structICspQueue.SettledBatch[]\",\"name\":\"_batches\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_generation\",\"type\":\"uint64\"}],\"name\":\"getSettlementCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getSourceConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"enumPaymentModel\",\"name\":\"paymentModel\",\"type\":\"uint8\"},{\"internalType\":\"enumISourceConfig.ChainKind\",\"name\":\"chainKind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"custodianEscrows\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"teeEscrows\",\"type\":\"bool\"}],\"internalType\":\"structISourceConfig.Source\",\"name\":\"_source\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount[]\",\"name\":\"_walletAccounts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getWalletId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getWalletRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"getWalletRegistryKind\",\"outputs\":[{\"internalType\":\"enumRegistryKind\",\"name\":\"_kind\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getXrpEscrow\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"nullified\",\"type\":\"bool\"}],\"internalType\":\"structIXrpEscrows.XrpEscrowTerms\",\"name\":\"_terms\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getXrpEscrowProfile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_destination\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"}],\"name\":\"isAllowedProposer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"isOperationPaymentId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isOperation\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"isSourceEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"isSourceRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_registered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"isValidAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_maxFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"nullifyOperation\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"_attempt\",\"type\":\"uint32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structPaymentInstruction\",\"name\":\"_paymentInstruction\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_createPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"enumICspInstructions.EmissionMode\",\"name\":\"_mode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_maxFee\",\"type\":\"uint256\"}],\"name\":\"reclaimEscrow\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"enumPaymentModel\",\"name\":\"paymentModel\",\"type\":\"uint8\"},{\"internalType\":\"enumISourceConfig.ChainKind\",\"name\":\"chainKind\",\"type\":\"uint8\"}],\"internalType\":\"structISourceConfig.SourceRegistration[]\",\"name\":\"_registrations\",\"type\":\"tuple[]\"}],\"name\":\"registerSources\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structPaymentInstruction[]\",\"name\":\"_retainedInstructions\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_nullifiedPaymentIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"maxFeePerPayment\",\"type\":\"uint256[]\"},{\"internalType\":\"int16[][]\",\"name\":\"factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"delaysSeconds\",\"type\":\"uint16[]\"}],\"internalType\":\"structReissueFeeParams\",\"name\":\"_reissueFeeParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"reissue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_finalized\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_maxFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"reissueOperation\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"_attempt\",\"type\":\"uint32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"removeAccountBtcWalletParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"removeAccountProposers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"removeProjectBtcWalletParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"removeProjectProposers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_accountIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccountConfigured.Anchor[]\",\"name\":\"_anchors\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proofOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestBtcAccountConfiguredAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_accountAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proofOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestNativeNonceAccountConfiguredAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"requestSigning\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"}],\"internalType\":\"structIBtcParams.BtcWalletParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"setAccountBtcWalletParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structIFeeSchedules.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"name\":\"setAccountFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"_proposers\",\"type\":\"address[]\"}],\"name\":\"setAccountProposers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"ruleVersion\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"minPaymentAmountSat\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"dustRelayFeeSatPerKvB\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreImprovementBias\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreImprovementClamp\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"maxSweptInputs\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"minFundingConfirmations\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"anchorSpendDepth\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeRateFloorBIPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"feeRateCeilingBIPS\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcParams.BtcConsensusParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"setBtcConsensusParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_counterpartyPubKey\",\"type\":\"bytes\"}],\"name\":\"setBtcEscrowProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"finalizationGraceSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"requiredTeeSignatures\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"proposalThresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"proposerRedundancy\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"proposerFeeWei\",\"type\":\"uint128\"}],\"internalType\":\"structICspProposals.CspSourceSettings\",\"name\":\"_settings\",\"type\":\"tuple\"}],\"name\":\"setCspSourceSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"internalType\":\"structIFeeSchedules.FeeScheduleConfigInput[]\",\"name\":\"_configs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeScheduleConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"}],\"internalType\":\"structIBtcParams.BtcWalletParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"setProjectBtcWalletParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structIFeeSchedules.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"name\":\"setProjectFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_proposers\",\"type\":\"address[]\"}],\"name\":\"setProjectProposers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"setProposerUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_custodianWallets\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_teeWallets\",\"type\":\"bool\"}],\"name\":\"setSourceEscrowsEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_sourceIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setSourcesEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_destination\",\"type\":\"string\"}],\"name\":\"setXrpEscrowProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"int16[][]\",\"name\":\"_factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"_delaysSeconds\",\"type\":\"uint16[]\"}],\"name\":\"validateAndEncodeSchedules\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_encodedPerPayment\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WalletPaymentsABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletPaymentsMetaData.ABI instead.
var WalletPaymentsABI = WalletPaymentsMetaData.ABI

// WalletPayments is an auto generated Go binding around an Ethereum contract.
type WalletPayments struct {
	WalletPaymentsCaller     // Read-only binding to the contract
	WalletPaymentsTransactor // Write-only binding to the contract
	WalletPaymentsFilterer   // Log filterer for contract events
}

// WalletPaymentsCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletPaymentsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletPaymentsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletPaymentsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletPaymentsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletPaymentsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletPaymentsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletPaymentsSession struct {
	Contract     *WalletPayments   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletPaymentsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletPaymentsCallerSession struct {
	Contract *WalletPaymentsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WalletPaymentsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletPaymentsTransactorSession struct {
	Contract     *WalletPaymentsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WalletPaymentsRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletPaymentsRaw struct {
	Contract *WalletPayments // Generic contract binding to access the raw methods on
}

// WalletPaymentsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletPaymentsCallerRaw struct {
	Contract *WalletPaymentsCaller // Generic read-only contract binding to access the raw methods on
}

// WalletPaymentsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletPaymentsTransactorRaw struct {
	Contract *WalletPaymentsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletPayments creates a new instance of WalletPayments, bound to a specific deployed contract.
func NewWalletPayments(address common.Address, backend bind.ContractBackend) (*WalletPayments, error) {
	contract, err := bindWalletPayments(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletPayments{WalletPaymentsCaller: WalletPaymentsCaller{contract: contract}, WalletPaymentsTransactor: WalletPaymentsTransactor{contract: contract}, WalletPaymentsFilterer: WalletPaymentsFilterer{contract: contract}}, nil
}

// NewWalletPaymentsCaller creates a new read-only instance of WalletPayments, bound to a specific deployed contract.
func NewWalletPaymentsCaller(address common.Address, caller bind.ContractCaller) (*WalletPaymentsCaller, error) {
	contract, err := bindWalletPayments(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsCaller{contract: contract}, nil
}

// NewWalletPaymentsTransactor creates a new write-only instance of WalletPayments, bound to a specific deployed contract.
func NewWalletPaymentsTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletPaymentsTransactor, error) {
	contract, err := bindWalletPayments(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsTransactor{contract: contract}, nil
}

// NewWalletPaymentsFilterer creates a new log filterer instance of WalletPayments, bound to a specific deployed contract.
func NewWalletPaymentsFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletPaymentsFilterer, error) {
	contract, err := bindWalletPayments(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsFilterer{contract: contract}, nil
}

// bindWalletPayments binds a generic wrapper to an already deployed contract.
func bindWalletPayments(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletPaymentsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletPayments *WalletPaymentsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletPayments.Contract.WalletPaymentsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletPayments *WalletPaymentsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletPayments.Contract.WalletPaymentsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletPayments *WalletPaymentsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletPayments.Contract.WalletPaymentsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletPayments *WalletPaymentsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletPayments.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletPayments *WalletPaymentsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletPayments.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletPayments *WalletPaymentsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletPayments.Contract.contract.Transact(opts, method, params...)
}

// CustodianWalletManager is a free data retrieval call binding the contract method 0xbb04ccfe.
//
// Solidity: function custodianWalletManager() view returns(address)
func (_WalletPayments *WalletPaymentsCaller) CustodianWalletManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "custodianWalletManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CustodianWalletManager is a free data retrieval call binding the contract method 0xbb04ccfe.
//
// Solidity: function custodianWalletManager() view returns(address)
func (_WalletPayments *WalletPaymentsSession) CustodianWalletManager() (common.Address, error) {
	return _WalletPayments.Contract.CustodianWalletManager(&_WalletPayments.CallOpts)
}

// CustodianWalletManager is a free data retrieval call binding the contract method 0xbb04ccfe.
//
// Solidity: function custodianWalletManager() view returns(address)
func (_WalletPayments *WalletPaymentsCallerSession) CustodianWalletManager() (common.Address, error) {
	return _WalletPayments.Contract.CustodianWalletManager(&_WalletPayments.CallOpts)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_WalletPayments *WalletPaymentsCaller) FacetAddress(opts *bind.CallOpts, _functionSelector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "facetAddress", _functionSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_WalletPayments *WalletPaymentsSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _WalletPayments.Contract.FacetAddress(&_WalletPayments.CallOpts, _functionSelector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_WalletPayments *WalletPaymentsCallerSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _WalletPayments.Contract.FacetAddress(&_WalletPayments.CallOpts, _functionSelector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_WalletPayments *WalletPaymentsCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_WalletPayments *WalletPaymentsSession) FacetAddresses() ([]common.Address, error) {
	return _WalletPayments.Contract.FacetAddresses(&_WalletPayments.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_WalletPayments *WalletPaymentsCallerSession) FacetAddresses() ([]common.Address, error) {
	return _WalletPayments.Contract.FacetAddresses(&_WalletPayments.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_WalletPayments *WalletPaymentsCaller) FacetFunctionSelectors(opts *bind.CallOpts, _facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "facetFunctionSelectors", _facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_WalletPayments *WalletPaymentsSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _WalletPayments.Contract.FacetFunctionSelectors(&_WalletPayments.CallOpts, _facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (_WalletPayments *WalletPaymentsCallerSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _WalletPayments.Contract.FacetFunctionSelectors(&_WalletPayments.CallOpts, _facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_WalletPayments *WalletPaymentsCaller) Facets(opts *bind.CallOpts) ([]IDiamondLoupeFacet, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_WalletPayments *WalletPaymentsSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _WalletPayments.Contract.Facets(&_WalletPayments.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_WalletPayments *WalletPaymentsCallerSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _WalletPayments.Contract.Facets(&_WalletPayments.CallOpts)
}

// Fdc2Hub is a free data retrieval call binding the contract method 0xa7566ff3.
//
// Solidity: function fdc2Hub() view returns(address)
func (_WalletPayments *WalletPaymentsCaller) Fdc2Hub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "fdc2Hub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fdc2Hub is a free data retrieval call binding the contract method 0xa7566ff3.
//
// Solidity: function fdc2Hub() view returns(address)
func (_WalletPayments *WalletPaymentsSession) Fdc2Hub() (common.Address, error) {
	return _WalletPayments.Contract.Fdc2Hub(&_WalletPayments.CallOpts)
}

// Fdc2Hub is a free data retrieval call binding the contract method 0xa7566ff3.
//
// Solidity: function fdc2Hub() view returns(address)
func (_WalletPayments *WalletPaymentsCallerSession) Fdc2Hub() (common.Address, error) {
	return _WalletPayments.Contract.Fdc2Hub(&_WalletPayments.CallOpts)
}

// Fdc2RequestFeeConfigurations is a free data retrieval call binding the contract method 0x285434a2.
//
// Solidity: function fdc2RequestFeeConfigurations() view returns(address)
func (_WalletPayments *WalletPaymentsCaller) Fdc2RequestFeeConfigurations(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "fdc2RequestFeeConfigurations")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fdc2RequestFeeConfigurations is a free data retrieval call binding the contract method 0x285434a2.
//
// Solidity: function fdc2RequestFeeConfigurations() view returns(address)
func (_WalletPayments *WalletPaymentsSession) Fdc2RequestFeeConfigurations() (common.Address, error) {
	return _WalletPayments.Contract.Fdc2RequestFeeConfigurations(&_WalletPayments.CallOpts)
}

// Fdc2RequestFeeConfigurations is a free data retrieval call binding the contract method 0x285434a2.
//
// Solidity: function fdc2RequestFeeConfigurations() view returns(address)
func (_WalletPayments *WalletPaymentsCallerSession) Fdc2RequestFeeConfigurations() (common.Address, error) {
	return _WalletPayments.Contract.Fdc2RequestFeeConfigurations(&_WalletPayments.CallOpts)
}

// Fdc2Verification is a free data retrieval call binding the contract method 0xbf2e9839.
//
// Solidity: function fdc2Verification() view returns(address)
func (_WalletPayments *WalletPaymentsCaller) Fdc2Verification(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "fdc2Verification")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fdc2Verification is a free data retrieval call binding the contract method 0xbf2e9839.
//
// Solidity: function fdc2Verification() view returns(address)
func (_WalletPayments *WalletPaymentsSession) Fdc2Verification() (common.Address, error) {
	return _WalletPayments.Contract.Fdc2Verification(&_WalletPayments.CallOpts)
}

// Fdc2Verification is a free data retrieval call binding the contract method 0xbf2e9839.
//
// Solidity: function fdc2Verification() view returns(address)
func (_WalletPayments *WalletPaymentsCallerSession) Fdc2Verification() (common.Address, error) {
	return _WalletPayments.Contract.Fdc2Verification(&_WalletPayments.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_WalletPayments *WalletPaymentsCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_WalletPayments *WalletPaymentsSession) FlareSystemsManager() (common.Address, error) {
	return _WalletPayments.Contract.FlareSystemsManager(&_WalletPayments.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_WalletPayments *WalletPaymentsCallerSession) FlareSystemsManager() (common.Address, error) {
	return _WalletPayments.Contract.FlareSystemsManager(&_WalletPayments.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_WalletPayments *WalletPaymentsCaller) FlareTeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "flareTeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_WalletPayments *WalletPaymentsSession) FlareTeeManager() (common.Address, error) {
	return _WalletPayments.Contract.FlareTeeManager(&_WalletPayments.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_WalletPayments *WalletPaymentsCallerSession) FlareTeeManager() (common.Address, error) {
	return _WalletPayments.Contract.FlareTeeManager(&_WalletPayments.CallOpts)
}

// GetAccountBtcWalletParams is a free data retrieval call binding the contract method 0x32b4af53.
//
// Solidity: function getAccountBtcWalletParams((bytes32,string) _account) view returns((uint64,uint32,uint8,uint8) _params)
func (_WalletPayments *WalletPaymentsCaller) GetAccountBtcWalletParams(opts *bind.CallOpts, _account WalletAccount) (IBtcParamsBtcWalletParams, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getAccountBtcWalletParams", _account)

	if err != nil {
		return *new(IBtcParamsBtcWalletParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IBtcParamsBtcWalletParams)).(*IBtcParamsBtcWalletParams)

	return out0, err

}

// GetAccountBtcWalletParams is a free data retrieval call binding the contract method 0x32b4af53.
//
// Solidity: function getAccountBtcWalletParams((bytes32,string) _account) view returns((uint64,uint32,uint8,uint8) _params)
func (_WalletPayments *WalletPaymentsSession) GetAccountBtcWalletParams(_account WalletAccount) (IBtcParamsBtcWalletParams, error) {
	return _WalletPayments.Contract.GetAccountBtcWalletParams(&_WalletPayments.CallOpts, _account)
}

// GetAccountBtcWalletParams is a free data retrieval call binding the contract method 0x32b4af53.
//
// Solidity: function getAccountBtcWalletParams((bytes32,string) _account) view returns((uint64,uint32,uint8,uint8) _params)
func (_WalletPayments *WalletPaymentsCallerSession) GetAccountBtcWalletParams(_account WalletAccount) (IBtcParamsBtcWalletParams, error) {
	return _WalletPayments.Contract.GetAccountBtcWalletParams(&_WalletPayments.CallOpts, _account)
}

// GetAccountFeeSchedule is a free data retrieval call binding the contract method 0x3c9e0c3f.
//
// Solidity: function getAccountFeeSchedule((bytes32,string) _account) view returns((int16,uint16)[] _schedule)
func (_WalletPayments *WalletPaymentsCaller) GetAccountFeeSchedule(opts *bind.CallOpts, _account WalletAccount) ([]IFeeSchedulesFeeSchedule, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getAccountFeeSchedule", _account)

	if err != nil {
		return *new([]IFeeSchedulesFeeSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new([]IFeeSchedulesFeeSchedule)).(*[]IFeeSchedulesFeeSchedule)

	return out0, err

}

// GetAccountFeeSchedule is a free data retrieval call binding the contract method 0x3c9e0c3f.
//
// Solidity: function getAccountFeeSchedule((bytes32,string) _account) view returns((int16,uint16)[] _schedule)
func (_WalletPayments *WalletPaymentsSession) GetAccountFeeSchedule(_account WalletAccount) ([]IFeeSchedulesFeeSchedule, error) {
	return _WalletPayments.Contract.GetAccountFeeSchedule(&_WalletPayments.CallOpts, _account)
}

// GetAccountFeeSchedule is a free data retrieval call binding the contract method 0x3c9e0c3f.
//
// Solidity: function getAccountFeeSchedule((bytes32,string) _account) view returns((int16,uint16)[] _schedule)
func (_WalletPayments *WalletPaymentsCallerSession) GetAccountFeeSchedule(_account WalletAccount) ([]IFeeSchedulesFeeSchedule, error) {
	return _WalletPayments.Contract.GetAccountFeeSchedule(&_WalletPayments.CallOpts, _account)
}

// GetAccountIndex is a free data retrieval call binding the contract method 0x92865318.
//
// Solidity: function getAccountIndex((bytes32,string) _account) view returns(uint32 _accountIndex)
func (_WalletPayments *WalletPaymentsCaller) GetAccountIndex(opts *bind.CallOpts, _account WalletAccount) (uint32, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getAccountIndex", _account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetAccountIndex is a free data retrieval call binding the contract method 0x92865318.
//
// Solidity: function getAccountIndex((bytes32,string) _account) view returns(uint32 _accountIndex)
func (_WalletPayments *WalletPaymentsSession) GetAccountIndex(_account WalletAccount) (uint32, error) {
	return _WalletPayments.Contract.GetAccountIndex(&_WalletPayments.CallOpts, _account)
}

// GetAccountIndex is a free data retrieval call binding the contract method 0x92865318.
//
// Solidity: function getAccountIndex((bytes32,string) _account) view returns(uint32 _accountIndex)
func (_WalletPayments *WalletPaymentsCallerSession) GetAccountIndex(_account WalletAccount) (uint32, error) {
	return _WalletPayments.Contract.GetAccountIndex(&_WalletPayments.CallOpts, _account)
}

// GetAccountProposers is a free data retrieval call binding the contract method 0x9438ac6b.
//
// Solidity: function getAccountProposers((bytes32,string) _account) view returns(address[] _proposers)
func (_WalletPayments *WalletPaymentsCaller) GetAccountProposers(opts *bind.CallOpts, _account WalletAccount) ([]common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getAccountProposers", _account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAccountProposers is a free data retrieval call binding the contract method 0x9438ac6b.
//
// Solidity: function getAccountProposers((bytes32,string) _account) view returns(address[] _proposers)
func (_WalletPayments *WalletPaymentsSession) GetAccountProposers(_account WalletAccount) ([]common.Address, error) {
	return _WalletPayments.Contract.GetAccountProposers(&_WalletPayments.CallOpts, _account)
}

// GetAccountProposers is a free data retrieval call binding the contract method 0x9438ac6b.
//
// Solidity: function getAccountProposers((bytes32,string) _account) view returns(address[] _proposers)
func (_WalletPayments *WalletPaymentsCallerSession) GetAccountProposers(_account WalletAccount) ([]common.Address, error) {
	return _WalletPayments.Contract.GetAccountProposers(&_WalletPayments.CallOpts, _account)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address)
func (_WalletPayments *WalletPaymentsCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address)
func (_WalletPayments *WalletPaymentsSession) GetAddressUpdater() (common.Address, error) {
	return _WalletPayments.Contract.GetAddressUpdater(&_WalletPayments.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address)
func (_WalletPayments *WalletPaymentsCallerSession) GetAddressUpdater() (common.Address, error) {
	return _WalletPayments.Contract.GetAddressUpdater(&_WalletPayments.CallOpts)
}

// GetAnchor is a free data retrieval call binding the contract method 0x207e9d61.
//
// Solidity: function getAnchor((bytes32,string) _account, uint32 _anchorIndex) view returns((bytes32,uint32,uint64) _anchor)
func (_WalletPayments *WalletPaymentsCaller) GetAnchor(opts *bind.CallOpts, _account WalletAccount, _anchorIndex uint32) (IBtcAccountsBtcAnchor, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getAnchor", _account, _anchorIndex)

	if err != nil {
		return *new(IBtcAccountsBtcAnchor), err
	}

	out0 := *abi.ConvertType(out[0], new(IBtcAccountsBtcAnchor)).(*IBtcAccountsBtcAnchor)

	return out0, err

}

// GetAnchor is a free data retrieval call binding the contract method 0x207e9d61.
//
// Solidity: function getAnchor((bytes32,string) _account, uint32 _anchorIndex) view returns((bytes32,uint32,uint64) _anchor)
func (_WalletPayments *WalletPaymentsSession) GetAnchor(_account WalletAccount, _anchorIndex uint32) (IBtcAccountsBtcAnchor, error) {
	return _WalletPayments.Contract.GetAnchor(&_WalletPayments.CallOpts, _account, _anchorIndex)
}

// GetAnchor is a free data retrieval call binding the contract method 0x207e9d61.
//
// Solidity: function getAnchor((bytes32,string) _account, uint32 _anchorIndex) view returns((bytes32,uint32,uint64) _anchor)
func (_WalletPayments *WalletPaymentsCallerSession) GetAnchor(_account WalletAccount, _anchorIndex uint32) (IBtcAccountsBtcAnchor, error) {
	return _WalletPayments.Contract.GetAnchor(&_WalletPayments.CallOpts, _account, _anchorIndex)
}

// GetAnchorCount is a free data retrieval call binding the contract method 0x00dc90fd.
//
// Solidity: function getAnchorCount((bytes32,string) _account) view returns(uint32 _anchorCount)
func (_WalletPayments *WalletPaymentsCaller) GetAnchorCount(opts *bind.CallOpts, _account WalletAccount) (uint32, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getAnchorCount", _account)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetAnchorCount is a free data retrieval call binding the contract method 0x00dc90fd.
//
// Solidity: function getAnchorCount((bytes32,string) _account) view returns(uint32 _anchorCount)
func (_WalletPayments *WalletPaymentsSession) GetAnchorCount(_account WalletAccount) (uint32, error) {
	return _WalletPayments.Contract.GetAnchorCount(&_WalletPayments.CallOpts, _account)
}

// GetAnchorCount is a free data retrieval call binding the contract method 0x00dc90fd.
//
// Solidity: function getAnchorCount((bytes32,string) _account) view returns(uint32 _anchorCount)
func (_WalletPayments *WalletPaymentsCallerSession) GetAnchorCount(_account WalletAccount) (uint32, error) {
	return _WalletPayments.Contract.GetAnchorCount(&_WalletPayments.CallOpts, _account)
}

// GetAttempts is a free data retrieval call binding the contract method 0xde93b092.
//
// Solidity: function getAttempts((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,address,uint64,uint64)[] _attempts)
func (_WalletPayments *WalletPaymentsCaller) GetAttempts(opts *bind.CallOpts, _account WalletAccount, _sequencePosition uint64) ([]ICspInstructionsAttempt, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getAttempts", _account, _sequencePosition)

	if err != nil {
		return *new([]ICspInstructionsAttempt), err
	}

	out0 := *abi.ConvertType(out[0], new([]ICspInstructionsAttempt)).(*[]ICspInstructionsAttempt)

	return out0, err

}

// GetAttempts is a free data retrieval call binding the contract method 0xde93b092.
//
// Solidity: function getAttempts((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,address,uint64,uint64)[] _attempts)
func (_WalletPayments *WalletPaymentsSession) GetAttempts(_account WalletAccount, _sequencePosition uint64) ([]ICspInstructionsAttempt, error) {
	return _WalletPayments.Contract.GetAttempts(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetAttempts is a free data retrieval call binding the contract method 0xde93b092.
//
// Solidity: function getAttempts((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,address,uint64,uint64)[] _attempts)
func (_WalletPayments *WalletPaymentsCallerSession) GetAttempts(_account WalletAccount, _sequencePosition uint64) ([]ICspInstructionsAttempt, error) {
	return _WalletPayments.Contract.GetAttempts(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (_WalletPayments *WalletPaymentsCaller) GetAuthorizationAddress(opts *bind.CallOpts, _account WalletAccount) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getAuthorizationAddress", _account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (_WalletPayments *WalletPaymentsSession) GetAuthorizationAddress(_account WalletAccount) (common.Address, error) {
	return _WalletPayments.Contract.GetAuthorizationAddress(&_WalletPayments.CallOpts, _account)
}

// GetAuthorizationAddress is a free data retrieval call binding the contract method 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (_WalletPayments *WalletPaymentsCallerSession) GetAuthorizationAddress(_account WalletAccount) (common.Address, error) {
	return _WalletPayments.Contract.GetAuthorizationAddress(&_WalletPayments.CallOpts, _account)
}

// GetBatchPaymentId is a free data retrieval call binding the contract method 0x41952376.
//
// Solidity: function getBatchPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(uint64 _batchPaymentId)
func (_WalletPayments *WalletPaymentsCaller) GetBatchPaymentId(opts *bind.CallOpts, _account WalletAccount, _paymentId uint64) (uint64, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getBatchPaymentId", _account, _paymentId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBatchPaymentId is a free data retrieval call binding the contract method 0x41952376.
//
// Solidity: function getBatchPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(uint64 _batchPaymentId)
func (_WalletPayments *WalletPaymentsSession) GetBatchPaymentId(_account WalletAccount, _paymentId uint64) (uint64, error) {
	return _WalletPayments.Contract.GetBatchPaymentId(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetBatchPaymentId is a free data retrieval call binding the contract method 0x41952376.
//
// Solidity: function getBatchPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(uint64 _batchPaymentId)
func (_WalletPayments *WalletPaymentsCallerSession) GetBatchPaymentId(_account WalletAccount, _paymentId uint64) (uint64, error) {
	return _WalletPayments.Contract.GetBatchPaymentId(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetBtcAttemptCommitments is a free data retrieval call binding the contract method 0x3abce1e5.
//
// Solidity: function getBtcAttemptCommitments((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,uint32)[] _commitments)
func (_WalletPayments *WalletPaymentsCaller) GetBtcAttemptCommitments(opts *bind.CallOpts, _account WalletAccount, _sequencePosition uint64) ([]IBtcAccountsBtcAttemptCommitment, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getBtcAttemptCommitments", _account, _sequencePosition)

	if err != nil {
		return *new([]IBtcAccountsBtcAttemptCommitment), err
	}

	out0 := *abi.ConvertType(out[0], new([]IBtcAccountsBtcAttemptCommitment)).(*[]IBtcAccountsBtcAttemptCommitment)

	return out0, err

}

// GetBtcAttemptCommitments is a free data retrieval call binding the contract method 0x3abce1e5.
//
// Solidity: function getBtcAttemptCommitments((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,uint32)[] _commitments)
func (_WalletPayments *WalletPaymentsSession) GetBtcAttemptCommitments(_account WalletAccount, _sequencePosition uint64) ([]IBtcAccountsBtcAttemptCommitment, error) {
	return _WalletPayments.Contract.GetBtcAttemptCommitments(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetBtcAttemptCommitments is a free data retrieval call binding the contract method 0x3abce1e5.
//
// Solidity: function getBtcAttemptCommitments((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,uint32)[] _commitments)
func (_WalletPayments *WalletPaymentsCallerSession) GetBtcAttemptCommitments(_account WalletAccount, _sequencePosition uint64) ([]IBtcAccountsBtcAttemptCommitment, error) {
	return _WalletPayments.Contract.GetBtcAttemptCommitments(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetBtcAttempts is a free data retrieval call binding the contract method 0x94d7a05b.
//
// Solidity: function getBtcAttempts((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,address,uint64,uint64)[] _attempts, (bytes32,bytes32,uint32)[] _commitments)
func (_WalletPayments *WalletPaymentsCaller) GetBtcAttempts(opts *bind.CallOpts, _account WalletAccount, _sequencePosition uint64) (struct {
	Attempts    []ICspInstructionsAttempt
	Commitments []IBtcAccountsBtcAttemptCommitment
}, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getBtcAttempts", _account, _sequencePosition)

	outstruct := new(struct {
		Attempts    []ICspInstructionsAttempt
		Commitments []IBtcAccountsBtcAttemptCommitment
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Attempts = *abi.ConvertType(out[0], new([]ICspInstructionsAttempt)).(*[]ICspInstructionsAttempt)
	outstruct.Commitments = *abi.ConvertType(out[1], new([]IBtcAccountsBtcAttemptCommitment)).(*[]IBtcAccountsBtcAttemptCommitment)

	return *outstruct, err

}

// GetBtcAttempts is a free data retrieval call binding the contract method 0x94d7a05b.
//
// Solidity: function getBtcAttempts((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,address,uint64,uint64)[] _attempts, (bytes32,bytes32,uint32)[] _commitments)
func (_WalletPayments *WalletPaymentsSession) GetBtcAttempts(_account WalletAccount, _sequencePosition uint64) (struct {
	Attempts    []ICspInstructionsAttempt
	Commitments []IBtcAccountsBtcAttemptCommitment
}, error) {
	return _WalletPayments.Contract.GetBtcAttempts(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetBtcAttempts is a free data retrieval call binding the contract method 0x94d7a05b.
//
// Solidity: function getBtcAttempts((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,address,uint64,uint64)[] _attempts, (bytes32,bytes32,uint32)[] _commitments)
func (_WalletPayments *WalletPaymentsCallerSession) GetBtcAttempts(_account WalletAccount, _sequencePosition uint64) (struct {
	Attempts    []ICspInstructionsAttempt
	Commitments []IBtcAccountsBtcAttemptCommitment
}, error) {
	return _WalletPayments.Contract.GetBtcAttempts(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetBtcConsensusParams is a free data retrieval call binding the contract method 0x62eae977.
//
// Solidity: function getBtcConsensusParams(bytes32 _sourceId) view returns((uint16,uint64,uint64,uint32,uint8,uint8,uint8,uint8,uint16,uint32,uint32,uint32,uint32,uint32) _params)
func (_WalletPayments *WalletPaymentsCaller) GetBtcConsensusParams(opts *bind.CallOpts, _sourceId [32]byte) (IBtcParamsBtcConsensusParams, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getBtcConsensusParams", _sourceId)

	if err != nil {
		return *new(IBtcParamsBtcConsensusParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IBtcParamsBtcConsensusParams)).(*IBtcParamsBtcConsensusParams)

	return out0, err

}

// GetBtcConsensusParams is a free data retrieval call binding the contract method 0x62eae977.
//
// Solidity: function getBtcConsensusParams(bytes32 _sourceId) view returns((uint16,uint64,uint64,uint32,uint8,uint8,uint8,uint8,uint16,uint32,uint32,uint32,uint32,uint32) _params)
func (_WalletPayments *WalletPaymentsSession) GetBtcConsensusParams(_sourceId [32]byte) (IBtcParamsBtcConsensusParams, error) {
	return _WalletPayments.Contract.GetBtcConsensusParams(&_WalletPayments.CallOpts, _sourceId)
}

// GetBtcConsensusParams is a free data retrieval call binding the contract method 0x62eae977.
//
// Solidity: function getBtcConsensusParams(bytes32 _sourceId) view returns((uint16,uint64,uint64,uint32,uint8,uint8,uint8,uint8,uint16,uint32,uint32,uint32,uint32,uint32) _params)
func (_WalletPayments *WalletPaymentsCallerSession) GetBtcConsensusParams(_sourceId [32]byte) (IBtcParamsBtcConsensusParams, error) {
	return _WalletPayments.Contract.GetBtcConsensusParams(&_WalletPayments.CallOpts, _sourceId)
}

// GetBtcEscrow is a free data retrieval call binding the contract method 0x648d78de.
//
// Solidity: function getBtcEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint64,uint64,bytes) _terms)
func (_WalletPayments *WalletPaymentsCaller) GetBtcEscrow(opts *bind.CallOpts, _account WalletAccount, _paymentId uint64) (IBtcEscrowsBtcEscrowTerms, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getBtcEscrow", _account, _paymentId)

	if err != nil {
		return *new(IBtcEscrowsBtcEscrowTerms), err
	}

	out0 := *abi.ConvertType(out[0], new(IBtcEscrowsBtcEscrowTerms)).(*IBtcEscrowsBtcEscrowTerms)

	return out0, err

}

// GetBtcEscrow is a free data retrieval call binding the contract method 0x648d78de.
//
// Solidity: function getBtcEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint64,uint64,bytes) _terms)
func (_WalletPayments *WalletPaymentsSession) GetBtcEscrow(_account WalletAccount, _paymentId uint64) (IBtcEscrowsBtcEscrowTerms, error) {
	return _WalletPayments.Contract.GetBtcEscrow(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetBtcEscrow is a free data retrieval call binding the contract method 0x648d78de.
//
// Solidity: function getBtcEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint64,uint64,bytes) _terms)
func (_WalletPayments *WalletPaymentsCallerSession) GetBtcEscrow(_account WalletAccount, _paymentId uint64) (IBtcEscrowsBtcEscrowTerms, error) {
	return _WalletPayments.Contract.GetBtcEscrow(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetBtcEscrowProfile is a free data retrieval call binding the contract method 0x465f2f9a.
//
// Solidity: function getBtcEscrowProfile((bytes32,string) _account) view returns(bytes _counterpartyPubKey)
func (_WalletPayments *WalletPaymentsCaller) GetBtcEscrowProfile(opts *bind.CallOpts, _account WalletAccount) ([]byte, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getBtcEscrowProfile", _account)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetBtcEscrowProfile is a free data retrieval call binding the contract method 0x465f2f9a.
//
// Solidity: function getBtcEscrowProfile((bytes32,string) _account) view returns(bytes _counterpartyPubKey)
func (_WalletPayments *WalletPaymentsSession) GetBtcEscrowProfile(_account WalletAccount) ([]byte, error) {
	return _WalletPayments.Contract.GetBtcEscrowProfile(&_WalletPayments.CallOpts, _account)
}

// GetBtcEscrowProfile is a free data retrieval call binding the contract method 0x465f2f9a.
//
// Solidity: function getBtcEscrowProfile((bytes32,string) _account) view returns(bytes _counterpartyPubKey)
func (_WalletPayments *WalletPaymentsCallerSession) GetBtcEscrowProfile(_account WalletAccount) ([]byte, error) {
	return _WalletPayments.Contract.GetBtcEscrowProfile(&_WalletPayments.CallOpts, _account)
}

// GetBtcEscrowReclaim is a free data retrieval call binding the contract method 0x1aca3e0b.
//
// Solidity: function getBtcEscrowReclaim((bytes32,string) _account, uint64 _createPaymentId) view returns(bool _reclaimed, uint64 _reclaimPaymentId)
func (_WalletPayments *WalletPaymentsCaller) GetBtcEscrowReclaim(opts *bind.CallOpts, _account WalletAccount, _createPaymentId uint64) (struct {
	Reclaimed        bool
	ReclaimPaymentId uint64
}, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getBtcEscrowReclaim", _account, _createPaymentId)

	outstruct := new(struct {
		Reclaimed        bool
		ReclaimPaymentId uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reclaimed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ReclaimPaymentId = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetBtcEscrowReclaim is a free data retrieval call binding the contract method 0x1aca3e0b.
//
// Solidity: function getBtcEscrowReclaim((bytes32,string) _account, uint64 _createPaymentId) view returns(bool _reclaimed, uint64 _reclaimPaymentId)
func (_WalletPayments *WalletPaymentsSession) GetBtcEscrowReclaim(_account WalletAccount, _createPaymentId uint64) (struct {
	Reclaimed        bool
	ReclaimPaymentId uint64
}, error) {
	return _WalletPayments.Contract.GetBtcEscrowReclaim(&_WalletPayments.CallOpts, _account, _createPaymentId)
}

// GetBtcEscrowReclaim is a free data retrieval call binding the contract method 0x1aca3e0b.
//
// Solidity: function getBtcEscrowReclaim((bytes32,string) _account, uint64 _createPaymentId) view returns(bool _reclaimed, uint64 _reclaimPaymentId)
func (_WalletPayments *WalletPaymentsCallerSession) GetBtcEscrowReclaim(_account WalletAccount, _createPaymentId uint64) (struct {
	Reclaimed        bool
	ReclaimPaymentId uint64
}, error) {
	return _WalletPayments.Contract.GetBtcEscrowReclaim(&_WalletPayments.CallOpts, _account, _createPaymentId)
}

// GetBtcEscrowTerms is a free data retrieval call binding the contract method 0x4557c259.
//
// Solidity: function getBtcEscrowTerms((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,uint64,uint64,bytes) _terms)
func (_WalletPayments *WalletPaymentsCaller) GetBtcEscrowTerms(opts *bind.CallOpts, _account WalletAccount, _sequencePosition uint64) (IBtcEscrowsBtcEscrowTerms, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getBtcEscrowTerms", _account, _sequencePosition)

	if err != nil {
		return *new(IBtcEscrowsBtcEscrowTerms), err
	}

	out0 := *abi.ConvertType(out[0], new(IBtcEscrowsBtcEscrowTerms)).(*IBtcEscrowsBtcEscrowTerms)

	return out0, err

}

// GetBtcEscrowTerms is a free data retrieval call binding the contract method 0x4557c259.
//
// Solidity: function getBtcEscrowTerms((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,uint64,uint64,bytes) _terms)
func (_WalletPayments *WalletPaymentsSession) GetBtcEscrowTerms(_account WalletAccount, _sequencePosition uint64) (IBtcEscrowsBtcEscrowTerms, error) {
	return _WalletPayments.Contract.GetBtcEscrowTerms(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetBtcEscrowTerms is a free data retrieval call binding the contract method 0x4557c259.
//
// Solidity: function getBtcEscrowTerms((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,uint64,uint64,bytes) _terms)
func (_WalletPayments *WalletPaymentsCallerSession) GetBtcEscrowTerms(_account WalletAccount, _sequencePosition uint64) (IBtcEscrowsBtcEscrowTerms, error) {
	return _WalletPayments.Contract.GetBtcEscrowTerms(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetBtcInstructionNonce is a free data retrieval call binding the contract method 0x196153d7.
//
// Solidity: function getBtcInstructionNonce((bytes32,string) _account, uint64 _sequencePosition) view returns(uint64 _nonce)
func (_WalletPayments *WalletPaymentsCaller) GetBtcInstructionNonce(opts *bind.CallOpts, _account WalletAccount, _sequencePosition uint64) (uint64, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getBtcInstructionNonce", _account, _sequencePosition)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetBtcInstructionNonce is a free data retrieval call binding the contract method 0x196153d7.
//
// Solidity: function getBtcInstructionNonce((bytes32,string) _account, uint64 _sequencePosition) view returns(uint64 _nonce)
func (_WalletPayments *WalletPaymentsSession) GetBtcInstructionNonce(_account WalletAccount, _sequencePosition uint64) (uint64, error) {
	return _WalletPayments.Contract.GetBtcInstructionNonce(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetBtcInstructionNonce is a free data retrieval call binding the contract method 0x196153d7.
//
// Solidity: function getBtcInstructionNonce((bytes32,string) _account, uint64 _sequencePosition) view returns(uint64 _nonce)
func (_WalletPayments *WalletPaymentsCallerSession) GetBtcInstructionNonce(_account WalletAccount, _sequencePosition uint64) (uint64, error) {
	return _WalletPayments.Contract.GetBtcInstructionNonce(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetBtcLeader is a free data retrieval call binding the contract method 0xbd1e31bd.
//
// Solidity: function getBtcLeader((bytes32,string) _account, uint64 _generation) view returns((bool,uint64,uint32,uint32,uint64,uint32,bytes32,bytes32,address,uint64) _leader, (uint32,uint64,uint32,bytes32,bytes32) _commitment)
func (_WalletPayments *WalletPaymentsCaller) GetBtcLeader(opts *bind.CallOpts, _account WalletAccount, _generation uint64) (struct {
	Leader     ICspProposalsLeader
	Commitment IBtcAccountsBtcProposalCommitment
}, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getBtcLeader", _account, _generation)

	outstruct := new(struct {
		Leader     ICspProposalsLeader
		Commitment IBtcAccountsBtcProposalCommitment
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Leader = *abi.ConvertType(out[0], new(ICspProposalsLeader)).(*ICspProposalsLeader)
	outstruct.Commitment = *abi.ConvertType(out[1], new(IBtcAccountsBtcProposalCommitment)).(*IBtcAccountsBtcProposalCommitment)

	return *outstruct, err

}

// GetBtcLeader is a free data retrieval call binding the contract method 0xbd1e31bd.
//
// Solidity: function getBtcLeader((bytes32,string) _account, uint64 _generation) view returns((bool,uint64,uint32,uint32,uint64,uint32,bytes32,bytes32,address,uint64) _leader, (uint32,uint64,uint32,bytes32,bytes32) _commitment)
func (_WalletPayments *WalletPaymentsSession) GetBtcLeader(_account WalletAccount, _generation uint64) (struct {
	Leader     ICspProposalsLeader
	Commitment IBtcAccountsBtcProposalCommitment
}, error) {
	return _WalletPayments.Contract.GetBtcLeader(&_WalletPayments.CallOpts, _account, _generation)
}

// GetBtcLeader is a free data retrieval call binding the contract method 0xbd1e31bd.
//
// Solidity: function getBtcLeader((bytes32,string) _account, uint64 _generation) view returns((bool,uint64,uint32,uint32,uint64,uint32,bytes32,bytes32,address,uint64) _leader, (uint32,uint64,uint32,bytes32,bytes32) _commitment)
func (_WalletPayments *WalletPaymentsCallerSession) GetBtcLeader(_account WalletAccount, _generation uint64) (struct {
	Leader     ICspProposalsLeader
	Commitment IBtcAccountsBtcProposalCommitment
}, error) {
	return _WalletPayments.Contract.GetBtcLeader(&_WalletPayments.CallOpts, _account, _generation)
}

// GetBtcParams is a free data retrieval call binding the contract method 0x5d2cf8c9.
//
// Solidity: function getBtcParams((bytes32,string) _account) view returns((uint16,uint64,uint64,uint32,uint8,uint8,uint8,uint8,uint16,uint32,uint32,uint32,uint32,uint32) _params)
func (_WalletPayments *WalletPaymentsCaller) GetBtcParams(opts *bind.CallOpts, _account WalletAccount) (IBtcParamsBtcConsensusParams, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getBtcParams", _account)

	if err != nil {
		return *new(IBtcParamsBtcConsensusParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IBtcParamsBtcConsensusParams)).(*IBtcParamsBtcConsensusParams)

	return out0, err

}

// GetBtcParams is a free data retrieval call binding the contract method 0x5d2cf8c9.
//
// Solidity: function getBtcParams((bytes32,string) _account) view returns((uint16,uint64,uint64,uint32,uint8,uint8,uint8,uint8,uint16,uint32,uint32,uint32,uint32,uint32) _params)
func (_WalletPayments *WalletPaymentsSession) GetBtcParams(_account WalletAccount) (IBtcParamsBtcConsensusParams, error) {
	return _WalletPayments.Contract.GetBtcParams(&_WalletPayments.CallOpts, _account)
}

// GetBtcParams is a free data retrieval call binding the contract method 0x5d2cf8c9.
//
// Solidity: function getBtcParams((bytes32,string) _account) view returns((uint16,uint64,uint64,uint32,uint8,uint8,uint8,uint8,uint16,uint32,uint32,uint32,uint32,uint32) _params)
func (_WalletPayments *WalletPaymentsCallerSession) GetBtcParams(_account WalletAccount) (IBtcParamsBtcConsensusParams, error) {
	return _WalletPayments.Contract.GetBtcParams(&_WalletPayments.CallOpts, _account)
}

// GetBtcProposalCommitment is a free data retrieval call binding the contract method 0xa27afb27.
//
// Solidity: function getBtcProposalCommitment((bytes32,string) _account, uint64 _generation) view returns((uint32,uint64,uint32,bytes32,bytes32) _commitment)
func (_WalletPayments *WalletPaymentsCaller) GetBtcProposalCommitment(opts *bind.CallOpts, _account WalletAccount, _generation uint64) (IBtcAccountsBtcProposalCommitment, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getBtcProposalCommitment", _account, _generation)

	if err != nil {
		return *new(IBtcAccountsBtcProposalCommitment), err
	}

	out0 := *abi.ConvertType(out[0], new(IBtcAccountsBtcProposalCommitment)).(*IBtcAccountsBtcProposalCommitment)

	return out0, err

}

// GetBtcProposalCommitment is a free data retrieval call binding the contract method 0xa27afb27.
//
// Solidity: function getBtcProposalCommitment((bytes32,string) _account, uint64 _generation) view returns((uint32,uint64,uint32,bytes32,bytes32) _commitment)
func (_WalletPayments *WalletPaymentsSession) GetBtcProposalCommitment(_account WalletAccount, _generation uint64) (IBtcAccountsBtcProposalCommitment, error) {
	return _WalletPayments.Contract.GetBtcProposalCommitment(&_WalletPayments.CallOpts, _account, _generation)
}

// GetBtcProposalCommitment is a free data retrieval call binding the contract method 0xa27afb27.
//
// Solidity: function getBtcProposalCommitment((bytes32,string) _account, uint64 _generation) view returns((uint32,uint64,uint32,bytes32,bytes32) _commitment)
func (_WalletPayments *WalletPaymentsCallerSession) GetBtcProposalCommitment(_account WalletAccount, _generation uint64) (IBtcAccountsBtcProposalCommitment, error) {
	return _WalletPayments.Contract.GetBtcProposalCommitment(&_WalletPayments.CallOpts, _account, _generation)
}

// GetCspAccount is a free data retrieval call binding the contract method 0xc154c258.
//
// Solidity: function getCspAccount(address _walletRegistry, bytes32 _walletId, uint32 _accountIndex) view returns((bytes32,string) _account)
func (_WalletPayments *WalletPaymentsCaller) GetCspAccount(opts *bind.CallOpts, _walletRegistry common.Address, _walletId [32]byte, _accountIndex uint32) (WalletAccount, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getCspAccount", _walletRegistry, _walletId, _accountIndex)

	if err != nil {
		return *new(WalletAccount), err
	}

	out0 := *abi.ConvertType(out[0], new(WalletAccount)).(*WalletAccount)

	return out0, err

}

// GetCspAccount is a free data retrieval call binding the contract method 0xc154c258.
//
// Solidity: function getCspAccount(address _walletRegistry, bytes32 _walletId, uint32 _accountIndex) view returns((bytes32,string) _account)
func (_WalletPayments *WalletPaymentsSession) GetCspAccount(_walletRegistry common.Address, _walletId [32]byte, _accountIndex uint32) (WalletAccount, error) {
	return _WalletPayments.Contract.GetCspAccount(&_WalletPayments.CallOpts, _walletRegistry, _walletId, _accountIndex)
}

// GetCspAccount is a free data retrieval call binding the contract method 0xc154c258.
//
// Solidity: function getCspAccount(address _walletRegistry, bytes32 _walletId, uint32 _accountIndex) view returns((bytes32,string) _account)
func (_WalletPayments *WalletPaymentsCallerSession) GetCspAccount(_walletRegistry common.Address, _walletId [32]byte, _accountIndex uint32) (WalletAccount, error) {
	return _WalletPayments.Contract.GetCspAccount(&_WalletPayments.CallOpts, _walletRegistry, _walletId, _accountIndex)
}

// GetCspSourceSettings is a free data retrieval call binding the contract method 0xed856dd2.
//
// Solidity: function getCspSourceSettings(bytes32 _sourceId) view returns((uint64,uint16,uint16,uint8,uint128) _settings)
func (_WalletPayments *WalletPaymentsCaller) GetCspSourceSettings(opts *bind.CallOpts, _sourceId [32]byte) (ICspProposalsCspSourceSettings, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getCspSourceSettings", _sourceId)

	if err != nil {
		return *new(ICspProposalsCspSourceSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(ICspProposalsCspSourceSettings)).(*ICspProposalsCspSourceSettings)

	return out0, err

}

// GetCspSourceSettings is a free data retrieval call binding the contract method 0xed856dd2.
//
// Solidity: function getCspSourceSettings(bytes32 _sourceId) view returns((uint64,uint16,uint16,uint8,uint128) _settings)
func (_WalletPayments *WalletPaymentsSession) GetCspSourceSettings(_sourceId [32]byte) (ICspProposalsCspSourceSettings, error) {
	return _WalletPayments.Contract.GetCspSourceSettings(&_WalletPayments.CallOpts, _sourceId)
}

// GetCspSourceSettings is a free data retrieval call binding the contract method 0xed856dd2.
//
// Solidity: function getCspSourceSettings(bytes32 _sourceId) view returns((uint64,uint16,uint16,uint8,uint128) _settings)
func (_WalletPayments *WalletPaymentsCallerSession) GetCspSourceSettings(_sourceId [32]byte) (ICspProposalsCspSourceSettings, error) {
	return _WalletPayments.Contract.GetCspSourceSettings(&_WalletPayments.CallOpts, _sourceId)
}

// GetEffectiveSchedule is a free data retrieval call binding the contract method 0xed8b6f51.
//
// Solidity: function getEffectiveSchedule(bytes32 _projectId, bytes32 _sourceId, bytes32 _accountHash) view returns(bytes _feeSchedule)
func (_WalletPayments *WalletPaymentsCaller) GetEffectiveSchedule(opts *bind.CallOpts, _projectId [32]byte, _sourceId [32]byte, _accountHash [32]byte) ([]byte, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getEffectiveSchedule", _projectId, _sourceId, _accountHash)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEffectiveSchedule is a free data retrieval call binding the contract method 0xed8b6f51.
//
// Solidity: function getEffectiveSchedule(bytes32 _projectId, bytes32 _sourceId, bytes32 _accountHash) view returns(bytes _feeSchedule)
func (_WalletPayments *WalletPaymentsSession) GetEffectiveSchedule(_projectId [32]byte, _sourceId [32]byte, _accountHash [32]byte) ([]byte, error) {
	return _WalletPayments.Contract.GetEffectiveSchedule(&_WalletPayments.CallOpts, _projectId, _sourceId, _accountHash)
}

// GetEffectiveSchedule is a free data retrieval call binding the contract method 0xed8b6f51.
//
// Solidity: function getEffectiveSchedule(bytes32 _projectId, bytes32 _sourceId, bytes32 _accountHash) view returns(bytes _feeSchedule)
func (_WalletPayments *WalletPaymentsCallerSession) GetEffectiveSchedule(_projectId [32]byte, _sourceId [32]byte, _accountHash [32]byte) ([]byte, error) {
	return _WalletPayments.Contract.GetEffectiveSchedule(&_WalletPayments.CallOpts, _projectId, _sourceId, _accountHash)
}

// GetEligible is a free data retrieval call binding the contract method 0xc475a77a.
//
// Solidity: function getEligible((bytes32,string) _account) view returns((uint64,uint32,uint64) _eligible)
func (_WalletPayments *WalletPaymentsCaller) GetEligible(opts *bind.CallOpts, _account WalletAccount) (ICspInstructionsEligible, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getEligible", _account)

	if err != nil {
		return *new(ICspInstructionsEligible), err
	}

	out0 := *abi.ConvertType(out[0], new(ICspInstructionsEligible)).(*ICspInstructionsEligible)

	return out0, err

}

// GetEligible is a free data retrieval call binding the contract method 0xc475a77a.
//
// Solidity: function getEligible((bytes32,string) _account) view returns((uint64,uint32,uint64) _eligible)
func (_WalletPayments *WalletPaymentsSession) GetEligible(_account WalletAccount) (ICspInstructionsEligible, error) {
	return _WalletPayments.Contract.GetEligible(&_WalletPayments.CallOpts, _account)
}

// GetEligible is a free data retrieval call binding the contract method 0xc475a77a.
//
// Solidity: function getEligible((bytes32,string) _account) view returns((uint64,uint32,uint64) _eligible)
func (_WalletPayments *WalletPaymentsCallerSession) GetEligible(_account WalletAccount) (ICspInstructionsEligible, error) {
	return _WalletPayments.Contract.GetEligible(&_WalletPayments.CallOpts, _account)
}

// GetEscrow is a free data retrieval call binding the contract method 0x8bce253d.
//
// Solidity: function getEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint256,uint64) _terms)
func (_WalletPayments *WalletPaymentsCaller) GetEscrow(opts *bind.CallOpts, _account WalletAccount, _paymentId uint64) (IEscrowsEscrowTerms, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getEscrow", _account, _paymentId)

	if err != nil {
		return *new(IEscrowsEscrowTerms), err
	}

	out0 := *abi.ConvertType(out[0], new(IEscrowsEscrowTerms)).(*IEscrowsEscrowTerms)

	return out0, err

}

// GetEscrow is a free data retrieval call binding the contract method 0x8bce253d.
//
// Solidity: function getEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint256,uint64) _terms)
func (_WalletPayments *WalletPaymentsSession) GetEscrow(_account WalletAccount, _paymentId uint64) (IEscrowsEscrowTerms, error) {
	return _WalletPayments.Contract.GetEscrow(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetEscrow is a free data retrieval call binding the contract method 0x8bce253d.
//
// Solidity: function getEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint256,uint64) _terms)
func (_WalletPayments *WalletPaymentsCallerSession) GetEscrow(_account WalletAccount, _paymentId uint64) (IEscrowsEscrowTerms, error) {
	return _WalletPayments.Contract.GetEscrow(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetFeeScheduleConfig is a free data retrieval call binding the contract method 0x16743f21.
//
// Solidity: function getFeeScheduleConfig(bytes32 _sourceId) view returns((uint8,uint16) _config)
func (_WalletPayments *WalletPaymentsCaller) GetFeeScheduleConfig(opts *bind.CallOpts, _sourceId [32]byte) (IFeeSchedulesFeeScheduleConfig, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getFeeScheduleConfig", _sourceId)

	if err != nil {
		return *new(IFeeSchedulesFeeScheduleConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(IFeeSchedulesFeeScheduleConfig)).(*IFeeSchedulesFeeScheduleConfig)

	return out0, err

}

// GetFeeScheduleConfig is a free data retrieval call binding the contract method 0x16743f21.
//
// Solidity: function getFeeScheduleConfig(bytes32 _sourceId) view returns((uint8,uint16) _config)
func (_WalletPayments *WalletPaymentsSession) GetFeeScheduleConfig(_sourceId [32]byte) (IFeeSchedulesFeeScheduleConfig, error) {
	return _WalletPayments.Contract.GetFeeScheduleConfig(&_WalletPayments.CallOpts, _sourceId)
}

// GetFeeScheduleConfig is a free data retrieval call binding the contract method 0x16743f21.
//
// Solidity: function getFeeScheduleConfig(bytes32 _sourceId) view returns((uint8,uint16) _config)
func (_WalletPayments *WalletPaymentsCallerSession) GetFeeScheduleConfig(_sourceId [32]byte) (IFeeSchedulesFeeScheduleConfig, error) {
	return _WalletPayments.Contract.GetFeeScheduleConfig(&_WalletPayments.CallOpts, _sourceId)
}

// GetFinalizedHash is a free data retrieval call binding the contract method 0x113b0c23.
//
// Solidity: function getFinalizedHash((bytes32,string) _account, uint64 _generation) view returns(bytes32 _packageHash)
func (_WalletPayments *WalletPaymentsCaller) GetFinalizedHash(opts *bind.CallOpts, _account WalletAccount, _generation uint64) ([32]byte, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getFinalizedHash", _account, _generation)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetFinalizedHash is a free data retrieval call binding the contract method 0x113b0c23.
//
// Solidity: function getFinalizedHash((bytes32,string) _account, uint64 _generation) view returns(bytes32 _packageHash)
func (_WalletPayments *WalletPaymentsSession) GetFinalizedHash(_account WalletAccount, _generation uint64) ([32]byte, error) {
	return _WalletPayments.Contract.GetFinalizedHash(&_WalletPayments.CallOpts, _account, _generation)
}

// GetFinalizedHash is a free data retrieval call binding the contract method 0x113b0c23.
//
// Solidity: function getFinalizedHash((bytes32,string) _account, uint64 _generation) view returns(bytes32 _packageHash)
func (_WalletPayments *WalletPaymentsCallerSession) GetFinalizedHash(_account WalletAccount, _generation uint64) ([32]byte, error) {
	return _WalletPayments.Contract.GetFinalizedHash(&_WalletPayments.CallOpts, _account, _generation)
}

// GetInitialNonce is a free data retrieval call binding the contract method 0x2bbba299.
//
// Solidity: function getInitialNonce((bytes32,string) _account) view returns(uint64 _initialNonce)
func (_WalletPayments *WalletPaymentsCaller) GetInitialNonce(opts *bind.CallOpts, _account WalletAccount) (uint64, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getInitialNonce", _account)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetInitialNonce is a free data retrieval call binding the contract method 0x2bbba299.
//
// Solidity: function getInitialNonce((bytes32,string) _account) view returns(uint64 _initialNonce)
func (_WalletPayments *WalletPaymentsSession) GetInitialNonce(_account WalletAccount) (uint64, error) {
	return _WalletPayments.Contract.GetInitialNonce(&_WalletPayments.CallOpts, _account)
}

// GetInitialNonce is a free data retrieval call binding the contract method 0x2bbba299.
//
// Solidity: function getInitialNonce((bytes32,string) _account) view returns(uint64 _initialNonce)
func (_WalletPayments *WalletPaymentsCallerSession) GetInitialNonce(_account WalletAccount) (uint64, error) {
	return _WalletPayments.Contract.GetInitialNonce(&_WalletPayments.CallOpts, _account)
}

// GetInstruction is a free data retrieval call binding the contract method 0x87e7e4ea.
//
// Solidity: function getInstruction((bytes32,string) _account, uint64 _sequencePosition) view returns((uint8,uint32,uint64,uint64,uint64,bool) _instruction)
func (_WalletPayments *WalletPaymentsCaller) GetInstruction(opts *bind.CallOpts, _account WalletAccount, _sequencePosition uint64) (ICspInstructionsInstructionRecord, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getInstruction", _account, _sequencePosition)

	if err != nil {
		return *new(ICspInstructionsInstructionRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(ICspInstructionsInstructionRecord)).(*ICspInstructionsInstructionRecord)

	return out0, err

}

// GetInstruction is a free data retrieval call binding the contract method 0x87e7e4ea.
//
// Solidity: function getInstruction((bytes32,string) _account, uint64 _sequencePosition) view returns((uint8,uint32,uint64,uint64,uint64,bool) _instruction)
func (_WalletPayments *WalletPaymentsSession) GetInstruction(_account WalletAccount, _sequencePosition uint64) (ICspInstructionsInstructionRecord, error) {
	return _WalletPayments.Contract.GetInstruction(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetInstruction is a free data retrieval call binding the contract method 0x87e7e4ea.
//
// Solidity: function getInstruction((bytes32,string) _account, uint64 _sequencePosition) view returns((uint8,uint32,uint64,uint64,uint64,bool) _instruction)
func (_WalletPayments *WalletPaymentsCallerSession) GetInstruction(_account WalletAccount, _sequencePosition uint64) (ICspInstructionsInstructionRecord, error) {
	return _WalletPayments.Contract.GetInstruction(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetLaneLastSettled is a free data retrieval call binding the contract method 0x5c9c71f1.
//
// Solidity: function getLaneLastSettled((bytes32,string) _account, uint32 _lane) view returns((bool,uint64) _use)
func (_WalletPayments *WalletPaymentsCaller) GetLaneLastSettled(opts *bind.CallOpts, _account WalletAccount, _lane uint32) (ICspInstructionsLaneUse, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getLaneLastSettled", _account, _lane)

	if err != nil {
		return *new(ICspInstructionsLaneUse), err
	}

	out0 := *abi.ConvertType(out[0], new(ICspInstructionsLaneUse)).(*ICspInstructionsLaneUse)

	return out0, err

}

// GetLaneLastSettled is a free data retrieval call binding the contract method 0x5c9c71f1.
//
// Solidity: function getLaneLastSettled((bytes32,string) _account, uint32 _lane) view returns((bool,uint64) _use)
func (_WalletPayments *WalletPaymentsSession) GetLaneLastSettled(_account WalletAccount, _lane uint32) (ICspInstructionsLaneUse, error) {
	return _WalletPayments.Contract.GetLaneLastSettled(&_WalletPayments.CallOpts, _account, _lane)
}

// GetLaneLastSettled is a free data retrieval call binding the contract method 0x5c9c71f1.
//
// Solidity: function getLaneLastSettled((bytes32,string) _account, uint32 _lane) view returns((bool,uint64) _use)
func (_WalletPayments *WalletPaymentsCallerSession) GetLaneLastSettled(_account WalletAccount, _lane uint32) (ICspInstructionsLaneUse, error) {
	return _WalletPayments.Contract.GetLaneLastSettled(&_WalletPayments.CallOpts, _account, _lane)
}

// GetLeader is a free data retrieval call binding the contract method 0x9db9ddf1.
//
// Solidity: function getLeader((bytes32,string) _account, uint64 _generation) view returns((bool,uint64,uint32,uint32,uint64,uint32,bytes32,bytes32,address,uint64) _leader)
func (_WalletPayments *WalletPaymentsCaller) GetLeader(opts *bind.CallOpts, _account WalletAccount, _generation uint64) (ICspProposalsLeader, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getLeader", _account, _generation)

	if err != nil {
		return *new(ICspProposalsLeader), err
	}

	out0 := *abi.ConvertType(out[0], new(ICspProposalsLeader)).(*ICspProposalsLeader)

	return out0, err

}

// GetLeader is a free data retrieval call binding the contract method 0x9db9ddf1.
//
// Solidity: function getLeader((bytes32,string) _account, uint64 _generation) view returns((bool,uint64,uint32,uint32,uint64,uint32,bytes32,bytes32,address,uint64) _leader)
func (_WalletPayments *WalletPaymentsSession) GetLeader(_account WalletAccount, _generation uint64) (ICspProposalsLeader, error) {
	return _WalletPayments.Contract.GetLeader(&_WalletPayments.CallOpts, _account, _generation)
}

// GetLeader is a free data retrieval call binding the contract method 0x9db9ddf1.
//
// Solidity: function getLeader((bytes32,string) _account, uint64 _generation) view returns((bool,uint64,uint32,uint32,uint64,uint32,bytes32,bytes32,address,uint64) _leader)
func (_WalletPayments *WalletPaymentsCallerSession) GetLeader(_account WalletAccount, _generation uint64) (ICspProposalsLeader, error) {
	return _WalletPayments.Contract.GetLeader(&_WalletPayments.CallOpts, _account, _generation)
}

// GetNetwork is a free data retrieval call binding the contract method 0x07e4b7e9.
//
// Solidity: function getNetwork() view returns(uint8 _network)
func (_WalletPayments *WalletPaymentsCaller) GetNetwork(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getNetwork")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetNetwork is a free data retrieval call binding the contract method 0x07e4b7e9.
//
// Solidity: function getNetwork() view returns(uint8 _network)
func (_WalletPayments *WalletPaymentsSession) GetNetwork() (uint8, error) {
	return _WalletPayments.Contract.GetNetwork(&_WalletPayments.CallOpts)
}

// GetNetwork is a free data retrieval call binding the contract method 0x07e4b7e9.
//
// Solidity: function getNetwork() view returns(uint8 _network)
func (_WalletPayments *WalletPaymentsCallerSession) GetNetwork() (uint8, error) {
	return _WalletPayments.Contract.GetNetwork(&_WalletPayments.CallOpts)
}

// GetNextPaymentId is a free data retrieval call binding the contract method 0xfb49ac30.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (_WalletPayments *WalletPaymentsCaller) GetNextPaymentId(opts *bind.CallOpts, _account WalletAccount) (uint64, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getNextPaymentId", _account)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextPaymentId is a free data retrieval call binding the contract method 0xfb49ac30.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (_WalletPayments *WalletPaymentsSession) GetNextPaymentId(_account WalletAccount) (uint64, error) {
	return _WalletPayments.Contract.GetNextPaymentId(&_WalletPayments.CallOpts, _account)
}

// GetNextPaymentId is a free data retrieval call binding the contract method 0xfb49ac30.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (_WalletPayments *WalletPaymentsCallerSession) GetNextPaymentId(_account WalletAccount) (uint64, error) {
	return _WalletPayments.Contract.GetNextPaymentId(&_WalletPayments.CallOpts, _account)
}

// GetNextSequencePosition is a free data retrieval call binding the contract method 0x7efb8ae6.
//
// Solidity: function getNextSequencePosition((bytes32,string) _account) view returns(uint64 _nextSequencePosition)
func (_WalletPayments *WalletPaymentsCaller) GetNextSequencePosition(opts *bind.CallOpts, _account WalletAccount) (uint64, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getNextSequencePosition", _account)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextSequencePosition is a free data retrieval call binding the contract method 0x7efb8ae6.
//
// Solidity: function getNextSequencePosition((bytes32,string) _account) view returns(uint64 _nextSequencePosition)
func (_WalletPayments *WalletPaymentsSession) GetNextSequencePosition(_account WalletAccount) (uint64, error) {
	return _WalletPayments.Contract.GetNextSequencePosition(&_WalletPayments.CallOpts, _account)
}

// GetNextSequencePosition is a free data retrieval call binding the contract method 0x7efb8ae6.
//
// Solidity: function getNextSequencePosition((bytes32,string) _account) view returns(uint64 _nextSequencePosition)
func (_WalletPayments *WalletPaymentsCallerSession) GetNextSequencePosition(_account WalletAccount) (uint64, error) {
	return _WalletPayments.Contract.GetNextSequencePosition(&_WalletPayments.CallOpts, _account)
}

// GetNextUnconsumedPaymentId is a free data retrieval call binding the contract method 0xf05af1b1.
//
// Solidity: function getNextUnconsumedPaymentId((bytes32,string) _account) view returns(uint64 _nextUnconsumedPaymentId)
func (_WalletPayments *WalletPaymentsCaller) GetNextUnconsumedPaymentId(opts *bind.CallOpts, _account WalletAccount) (uint64, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getNextUnconsumedPaymentId", _account)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNextUnconsumedPaymentId is a free data retrieval call binding the contract method 0xf05af1b1.
//
// Solidity: function getNextUnconsumedPaymentId((bytes32,string) _account) view returns(uint64 _nextUnconsumedPaymentId)
func (_WalletPayments *WalletPaymentsSession) GetNextUnconsumedPaymentId(_account WalletAccount) (uint64, error) {
	return _WalletPayments.Contract.GetNextUnconsumedPaymentId(&_WalletPayments.CallOpts, _account)
}

// GetNextUnconsumedPaymentId is a free data retrieval call binding the contract method 0xf05af1b1.
//
// Solidity: function getNextUnconsumedPaymentId((bytes32,string) _account) view returns(uint64 _nextUnconsumedPaymentId)
func (_WalletPayments *WalletPaymentsCallerSession) GetNextUnconsumedPaymentId(_account WalletAccount) (uint64, error) {
	return _WalletPayments.Contract.GetNextUnconsumedPaymentId(&_WalletPayments.CallOpts, _account)
}

// GetNullifiedPaymentIds is a free data retrieval call binding the contract method 0x67135907.
//
// Solidity: function getNullifiedPaymentIds((bytes32,string) _account, uint64 _sequencePosition) view returns(uint64[] _nullifiedPaymentIds)
func (_WalletPayments *WalletPaymentsCaller) GetNullifiedPaymentIds(opts *bind.CallOpts, _account WalletAccount, _sequencePosition uint64) ([]uint64, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getNullifiedPaymentIds", _account, _sequencePosition)

	if err != nil {
		return *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)

	return out0, err

}

// GetNullifiedPaymentIds is a free data retrieval call binding the contract method 0x67135907.
//
// Solidity: function getNullifiedPaymentIds((bytes32,string) _account, uint64 _sequencePosition) view returns(uint64[] _nullifiedPaymentIds)
func (_WalletPayments *WalletPaymentsSession) GetNullifiedPaymentIds(_account WalletAccount, _sequencePosition uint64) ([]uint64, error) {
	return _WalletPayments.Contract.GetNullifiedPaymentIds(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetNullifiedPaymentIds is a free data retrieval call binding the contract method 0x67135907.
//
// Solidity: function getNullifiedPaymentIds((bytes32,string) _account, uint64 _sequencePosition) view returns(uint64[] _nullifiedPaymentIds)
func (_WalletPayments *WalletPaymentsCallerSession) GetNullifiedPaymentIds(_account WalletAccount, _sequencePosition uint64) ([]uint64, error) {
	return _WalletPayments.Contract.GetNullifiedPaymentIds(&_WalletPayments.CallOpts, _account, _sequencePosition)
}

// GetOperationSequencePosition is a free data retrieval call binding the contract method 0x09e616d8.
//
// Solidity: function getOperationSequencePosition((bytes32,string) _account, uint64 _paymentId) view returns(bool _isOperation, uint64 _sequencePosition)
func (_WalletPayments *WalletPaymentsCaller) GetOperationSequencePosition(opts *bind.CallOpts, _account WalletAccount, _paymentId uint64) (struct {
	IsOperation      bool
	SequencePosition uint64
}, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getOperationSequencePosition", _account, _paymentId)

	outstruct := new(struct {
		IsOperation      bool
		SequencePosition uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsOperation = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SequencePosition = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetOperationSequencePosition is a free data retrieval call binding the contract method 0x09e616d8.
//
// Solidity: function getOperationSequencePosition((bytes32,string) _account, uint64 _paymentId) view returns(bool _isOperation, uint64 _sequencePosition)
func (_WalletPayments *WalletPaymentsSession) GetOperationSequencePosition(_account WalletAccount, _paymentId uint64) (struct {
	IsOperation      bool
	SequencePosition uint64
}, error) {
	return _WalletPayments.Contract.GetOperationSequencePosition(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetOperationSequencePosition is a free data retrieval call binding the contract method 0x09e616d8.
//
// Solidity: function getOperationSequencePosition((bytes32,string) _account, uint64 _paymentId) view returns(bool _isOperation, uint64 _sequencePosition)
func (_WalletPayments *WalletPaymentsCallerSession) GetOperationSequencePosition(_account WalletAccount, _paymentId uint64) (struct {
	IsOperation      bool
	SequencePosition uint64
}, error) {
	return _WalletPayments.Contract.GetOperationSequencePosition(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetPaymentFee is a free data retrieval call binding the contract method 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (_WalletPayments *WalletPaymentsCaller) GetPaymentFee(opts *bind.CallOpts, _account WalletAccount, _opCommand [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getPaymentFee", _account, _opCommand)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPaymentFee is a free data retrieval call binding the contract method 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (_WalletPayments *WalletPaymentsSession) GetPaymentFee(_account WalletAccount, _opCommand [32]byte) (*big.Int, error) {
	return _WalletPayments.Contract.GetPaymentFee(&_WalletPayments.CallOpts, _account, _opCommand)
}

// GetPaymentFee is a free data retrieval call binding the contract method 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (_WalletPayments *WalletPaymentsCallerSession) GetPaymentFee(_account WalletAccount, _opCommand [32]byte) (*big.Int, error) {
	return _WalletPayments.Contract.GetPaymentFee(&_WalletPayments.CallOpts, _account, _opCommand)
}

// GetPaymentHash is a free data retrieval call binding the contract method 0x0079dfe8.
//
// Solidity: function getPaymentHash((bytes32,string) _account, uint64 _paymentId) view returns(bytes32 _paymentHash)
func (_WalletPayments *WalletPaymentsCaller) GetPaymentHash(opts *bind.CallOpts, _account WalletAccount, _paymentId uint64) ([32]byte, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getPaymentHash", _account, _paymentId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetPaymentHash is a free data retrieval call binding the contract method 0x0079dfe8.
//
// Solidity: function getPaymentHash((bytes32,string) _account, uint64 _paymentId) view returns(bytes32 _paymentHash)
func (_WalletPayments *WalletPaymentsSession) GetPaymentHash(_account WalletAccount, _paymentId uint64) ([32]byte, error) {
	return _WalletPayments.Contract.GetPaymentHash(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetPaymentHash is a free data retrieval call binding the contract method 0x0079dfe8.
//
// Solidity: function getPaymentHash((bytes32,string) _account, uint64 _paymentId) view returns(bytes32 _paymentHash)
func (_WalletPayments *WalletPaymentsCallerSession) GetPaymentHash(_account WalletAccount, _paymentId uint64) ([32]byte, error) {
	return _WalletPayments.Contract.GetPaymentHash(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetPendingResets is a free data retrieval call binding the contract method 0x807a0ade.
//
// Solidity: function getPendingResets((bytes32,string) _account) view returns((uint64,uint32)[] _pendingResets)
func (_WalletPayments *WalletPaymentsCaller) GetPendingResets(opts *bind.CallOpts, _account WalletAccount) ([]ICspInstructionsPendingReset, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getPendingResets", _account)

	if err != nil {
		return *new([]ICspInstructionsPendingReset), err
	}

	out0 := *abi.ConvertType(out[0], new([]ICspInstructionsPendingReset)).(*[]ICspInstructionsPendingReset)

	return out0, err

}

// GetPendingResets is a free data retrieval call binding the contract method 0x807a0ade.
//
// Solidity: function getPendingResets((bytes32,string) _account) view returns((uint64,uint32)[] _pendingResets)
func (_WalletPayments *WalletPaymentsSession) GetPendingResets(_account WalletAccount) ([]ICspInstructionsPendingReset, error) {
	return _WalletPayments.Contract.GetPendingResets(&_WalletPayments.CallOpts, _account)
}

// GetPendingResets is a free data retrieval call binding the contract method 0x807a0ade.
//
// Solidity: function getPendingResets((bytes32,string) _account) view returns((uint64,uint32)[] _pendingResets)
func (_WalletPayments *WalletPaymentsCallerSession) GetPendingResets(_account WalletAccount) ([]ICspInstructionsPendingReset, error) {
	return _WalletPayments.Contract.GetPendingResets(&_WalletPayments.CallOpts, _account)
}

// GetProjectBtcWalletParams is a free data retrieval call binding the contract method 0xa466c45f.
//
// Solidity: function getProjectBtcWalletParams(bytes32 _projectId) view returns((uint64,uint32,uint8,uint8) _params)
func (_WalletPayments *WalletPaymentsCaller) GetProjectBtcWalletParams(opts *bind.CallOpts, _projectId [32]byte) (IBtcParamsBtcWalletParams, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getProjectBtcWalletParams", _projectId)

	if err != nil {
		return *new(IBtcParamsBtcWalletParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IBtcParamsBtcWalletParams)).(*IBtcParamsBtcWalletParams)

	return out0, err

}

// GetProjectBtcWalletParams is a free data retrieval call binding the contract method 0xa466c45f.
//
// Solidity: function getProjectBtcWalletParams(bytes32 _projectId) view returns((uint64,uint32,uint8,uint8) _params)
func (_WalletPayments *WalletPaymentsSession) GetProjectBtcWalletParams(_projectId [32]byte) (IBtcParamsBtcWalletParams, error) {
	return _WalletPayments.Contract.GetProjectBtcWalletParams(&_WalletPayments.CallOpts, _projectId)
}

// GetProjectBtcWalletParams is a free data retrieval call binding the contract method 0xa466c45f.
//
// Solidity: function getProjectBtcWalletParams(bytes32 _projectId) view returns((uint64,uint32,uint8,uint8) _params)
func (_WalletPayments *WalletPaymentsCallerSession) GetProjectBtcWalletParams(_projectId [32]byte) (IBtcParamsBtcWalletParams, error) {
	return _WalletPayments.Contract.GetProjectBtcWalletParams(&_WalletPayments.CallOpts, _projectId)
}

// GetProjectFeeSchedule is a free data retrieval call binding the contract method 0x30d1bb93.
//
// Solidity: function getProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) view returns((int16,uint16)[] _schedule)
func (_WalletPayments *WalletPaymentsCaller) GetProjectFeeSchedule(opts *bind.CallOpts, _projectId [32]byte, _sourceId [32]byte) ([]IFeeSchedulesFeeSchedule, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getProjectFeeSchedule", _projectId, _sourceId)

	if err != nil {
		return *new([]IFeeSchedulesFeeSchedule), err
	}

	out0 := *abi.ConvertType(out[0], new([]IFeeSchedulesFeeSchedule)).(*[]IFeeSchedulesFeeSchedule)

	return out0, err

}

// GetProjectFeeSchedule is a free data retrieval call binding the contract method 0x30d1bb93.
//
// Solidity: function getProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) view returns((int16,uint16)[] _schedule)
func (_WalletPayments *WalletPaymentsSession) GetProjectFeeSchedule(_projectId [32]byte, _sourceId [32]byte) ([]IFeeSchedulesFeeSchedule, error) {
	return _WalletPayments.Contract.GetProjectFeeSchedule(&_WalletPayments.CallOpts, _projectId, _sourceId)
}

// GetProjectFeeSchedule is a free data retrieval call binding the contract method 0x30d1bb93.
//
// Solidity: function getProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) view returns((int16,uint16)[] _schedule)
func (_WalletPayments *WalletPaymentsCallerSession) GetProjectFeeSchedule(_projectId [32]byte, _sourceId [32]byte) ([]IFeeSchedulesFeeSchedule, error) {
	return _WalletPayments.Contract.GetProjectFeeSchedule(&_WalletPayments.CallOpts, _projectId, _sourceId)
}

// GetProjectProposers is a free data retrieval call binding the contract method 0xc07f32fe.
//
// Solidity: function getProjectProposers(bytes32 _projectId) view returns(address[] _proposers)
func (_WalletPayments *WalletPaymentsCaller) GetProjectProposers(opts *bind.CallOpts, _projectId [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getProjectProposers", _projectId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetProjectProposers is a free data retrieval call binding the contract method 0xc07f32fe.
//
// Solidity: function getProjectProposers(bytes32 _projectId) view returns(address[] _proposers)
func (_WalletPayments *WalletPaymentsSession) GetProjectProposers(_projectId [32]byte) ([]common.Address, error) {
	return _WalletPayments.Contract.GetProjectProposers(&_WalletPayments.CallOpts, _projectId)
}

// GetProjectProposers is a free data retrieval call binding the contract method 0xc07f32fe.
//
// Solidity: function getProjectProposers(bytes32 _projectId) view returns(address[] _proposers)
func (_WalletPayments *WalletPaymentsCallerSession) GetProjectProposers(_projectId [32]byte) ([]common.Address, error) {
	return _WalletPayments.Contract.GetProjectProposers(&_WalletPayments.CallOpts, _projectId)
}

// GetProposerUrl is a free data retrieval call binding the contract method 0x6cb39f40.
//
// Solidity: function getProposerUrl(address _proposer) view returns(string _url)
func (_WalletPayments *WalletPaymentsCaller) GetProposerUrl(opts *bind.CallOpts, _proposer common.Address) (string, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getProposerUrl", _proposer)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetProposerUrl is a free data retrieval call binding the contract method 0x6cb39f40.
//
// Solidity: function getProposerUrl(address _proposer) view returns(string _url)
func (_WalletPayments *WalletPaymentsSession) GetProposerUrl(_proposer common.Address) (string, error) {
	return _WalletPayments.Contract.GetProposerUrl(&_WalletPayments.CallOpts, _proposer)
}

// GetProposerUrl is a free data retrieval call binding the contract method 0x6cb39f40.
//
// Solidity: function getProposerUrl(address _proposer) view returns(string _url)
func (_WalletPayments *WalletPaymentsCallerSession) GetProposerUrl(_proposer common.Address) (string, error) {
	return _WalletPayments.Contract.GetProposerUrl(&_WalletPayments.CallOpts, _proposer)
}

// GetQueuedPayment is a free data retrieval call binding the contract method 0x532de34f.
//
// Solidity: function getQueuedPayment((bytes32,string) _account, uint64 _paymentId) view returns((string,uint64,uint64,bytes32) _payment)
func (_WalletPayments *WalletPaymentsCaller) GetQueuedPayment(opts *bind.CallOpts, _account WalletAccount, _paymentId uint64) (ICspQueueQueuedPayment, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getQueuedPayment", _account, _paymentId)

	if err != nil {
		return *new(ICspQueueQueuedPayment), err
	}

	out0 := *abi.ConvertType(out[0], new(ICspQueueQueuedPayment)).(*ICspQueueQueuedPayment)

	return out0, err

}

// GetQueuedPayment is a free data retrieval call binding the contract method 0x532de34f.
//
// Solidity: function getQueuedPayment((bytes32,string) _account, uint64 _paymentId) view returns((string,uint64,uint64,bytes32) _payment)
func (_WalletPayments *WalletPaymentsSession) GetQueuedPayment(_account WalletAccount, _paymentId uint64) (ICspQueueQueuedPayment, error) {
	return _WalletPayments.Contract.GetQueuedPayment(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetQueuedPayment is a free data retrieval call binding the contract method 0x532de34f.
//
// Solidity: function getQueuedPayment((bytes32,string) _account, uint64 _paymentId) view returns((string,uint64,uint64,bytes32) _payment)
func (_WalletPayments *WalletPaymentsCallerSession) GetQueuedPayment(_account WalletAccount, _paymentId uint64) (ICspQueueQueuedPayment, error) {
	return _WalletPayments.Contract.GetQueuedPayment(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetRegisteredSourceIds is a free data retrieval call binding the contract method 0xec02c8e0.
//
// Solidity: function getRegisteredSourceIds() view returns(bytes32[] _sourceIds)
func (_WalletPayments *WalletPaymentsCaller) GetRegisteredSourceIds(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getRegisteredSourceIds")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetRegisteredSourceIds is a free data retrieval call binding the contract method 0xec02c8e0.
//
// Solidity: function getRegisteredSourceIds() view returns(bytes32[] _sourceIds)
func (_WalletPayments *WalletPaymentsSession) GetRegisteredSourceIds() ([][32]byte, error) {
	return _WalletPayments.Contract.GetRegisteredSourceIds(&_WalletPayments.CallOpts)
}

// GetRegisteredSourceIds is a free data retrieval call binding the contract method 0xec02c8e0.
//
// Solidity: function getRegisteredSourceIds() view returns(bytes32[] _sourceIds)
func (_WalletPayments *WalletPaymentsCallerSession) GetRegisteredSourceIds() ([][32]byte, error) {
	return _WalletPayments.Contract.GetRegisteredSourceIds(&_WalletPayments.CallOpts)
}

// GetSettledBatch is a free data retrieval call binding the contract method 0xa30e1509.
//
// Solidity: function getSettledBatch((bytes32,string) _account, uint64 _paymentId) view returns((uint64,uint64,uint64) _batch)
func (_WalletPayments *WalletPaymentsCaller) GetSettledBatch(opts *bind.CallOpts, _account WalletAccount, _paymentId uint64) (ICspQueueSettledBatch, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getSettledBatch", _account, _paymentId)

	if err != nil {
		return *new(ICspQueueSettledBatch), err
	}

	out0 := *abi.ConvertType(out[0], new(ICspQueueSettledBatch)).(*ICspQueueSettledBatch)

	return out0, err

}

// GetSettledBatch is a free data retrieval call binding the contract method 0xa30e1509.
//
// Solidity: function getSettledBatch((bytes32,string) _account, uint64 _paymentId) view returns((uint64,uint64,uint64) _batch)
func (_WalletPayments *WalletPaymentsSession) GetSettledBatch(_account WalletAccount, _paymentId uint64) (ICspQueueSettledBatch, error) {
	return _WalletPayments.Contract.GetSettledBatch(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetSettledBatch is a free data retrieval call binding the contract method 0xa30e1509.
//
// Solidity: function getSettledBatch((bytes32,string) _account, uint64 _paymentId) view returns((uint64,uint64,uint64) _batch)
func (_WalletPayments *WalletPaymentsCallerSession) GetSettledBatch(_account WalletAccount, _paymentId uint64) (ICspQueueSettledBatch, error) {
	return _WalletPayments.Contract.GetSettledBatch(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetSettledBatches is a free data retrieval call binding the contract method 0x2a29805c.
//
// Solidity: function getSettledBatches((bytes32,string) _account) view returns((uint64,uint64,uint64)[] _batches)
func (_WalletPayments *WalletPaymentsCaller) GetSettledBatches(opts *bind.CallOpts, _account WalletAccount) ([]ICspQueueSettledBatch, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getSettledBatches", _account)

	if err != nil {
		return *new([]ICspQueueSettledBatch), err
	}

	out0 := *abi.ConvertType(out[0], new([]ICspQueueSettledBatch)).(*[]ICspQueueSettledBatch)

	return out0, err

}

// GetSettledBatches is a free data retrieval call binding the contract method 0x2a29805c.
//
// Solidity: function getSettledBatches((bytes32,string) _account) view returns((uint64,uint64,uint64)[] _batches)
func (_WalletPayments *WalletPaymentsSession) GetSettledBatches(_account WalletAccount) ([]ICspQueueSettledBatch, error) {
	return _WalletPayments.Contract.GetSettledBatches(&_WalletPayments.CallOpts, _account)
}

// GetSettledBatches is a free data retrieval call binding the contract method 0x2a29805c.
//
// Solidity: function getSettledBatches((bytes32,string) _account) view returns((uint64,uint64,uint64)[] _batches)
func (_WalletPayments *WalletPaymentsCallerSession) GetSettledBatches(_account WalletAccount) ([]ICspQueueSettledBatch, error) {
	return _WalletPayments.Contract.GetSettledBatches(&_WalletPayments.CallOpts, _account)
}

// GetSettlementCost is a free data retrieval call binding the contract method 0x52072400.
//
// Solidity: function getSettlementCost((bytes32,string) _account, uint64 _generation) view returns(uint256 _cost)
func (_WalletPayments *WalletPaymentsCaller) GetSettlementCost(opts *bind.CallOpts, _account WalletAccount, _generation uint64) (*big.Int, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getSettlementCost", _account, _generation)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSettlementCost is a free data retrieval call binding the contract method 0x52072400.
//
// Solidity: function getSettlementCost((bytes32,string) _account, uint64 _generation) view returns(uint256 _cost)
func (_WalletPayments *WalletPaymentsSession) GetSettlementCost(_account WalletAccount, _generation uint64) (*big.Int, error) {
	return _WalletPayments.Contract.GetSettlementCost(&_WalletPayments.CallOpts, _account, _generation)
}

// GetSettlementCost is a free data retrieval call binding the contract method 0x52072400.
//
// Solidity: function getSettlementCost((bytes32,string) _account, uint64 _generation) view returns(uint256 _cost)
func (_WalletPayments *WalletPaymentsCallerSession) GetSettlementCost(_account WalletAccount, _generation uint64) (*big.Int, error) {
	return _WalletPayments.Contract.GetSettlementCost(&_WalletPayments.CallOpts, _account, _generation)
}

// GetSourceConfig is a free data retrieval call binding the contract method 0x9a5b5cfd.
//
// Solidity: function getSourceConfig(bytes32 _sourceId) view returns((bytes32,bytes32,uint8,uint8,bool,bool,bool,bool) _source)
func (_WalletPayments *WalletPaymentsCaller) GetSourceConfig(opts *bind.CallOpts, _sourceId [32]byte) (ISourceConfigSource, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getSourceConfig", _sourceId)

	if err != nil {
		return *new(ISourceConfigSource), err
	}

	out0 := *abi.ConvertType(out[0], new(ISourceConfigSource)).(*ISourceConfigSource)

	return out0, err

}

// GetSourceConfig is a free data retrieval call binding the contract method 0x9a5b5cfd.
//
// Solidity: function getSourceConfig(bytes32 _sourceId) view returns((bytes32,bytes32,uint8,uint8,bool,bool,bool,bool) _source)
func (_WalletPayments *WalletPaymentsSession) GetSourceConfig(_sourceId [32]byte) (ISourceConfigSource, error) {
	return _WalletPayments.Contract.GetSourceConfig(&_WalletPayments.CallOpts, _sourceId)
}

// GetSourceConfig is a free data retrieval call binding the contract method 0x9a5b5cfd.
//
// Solidity: function getSourceConfig(bytes32 _sourceId) view returns((bytes32,bytes32,uint8,uint8,bool,bool,bool,bool) _source)
func (_WalletPayments *WalletPaymentsCallerSession) GetSourceConfig(_sourceId [32]byte) (ISourceConfigSource, error) {
	return _WalletPayments.Contract.GetSourceConfig(&_WalletPayments.CallOpts, _sourceId)
}

// GetWalletAccounts is a free data retrieval call binding the contract method 0xbd0b152c.
//
// Solidity: function getWalletAccounts(address _walletRegistry, bytes32 _walletId) view returns((bytes32,string)[] _walletAccounts)
func (_WalletPayments *WalletPaymentsCaller) GetWalletAccounts(opts *bind.CallOpts, _walletRegistry common.Address, _walletId [32]byte) ([]WalletAccount, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getWalletAccounts", _walletRegistry, _walletId)

	if err != nil {
		return *new([]WalletAccount), err
	}

	out0 := *abi.ConvertType(out[0], new([]WalletAccount)).(*[]WalletAccount)

	return out0, err

}

// GetWalletAccounts is a free data retrieval call binding the contract method 0xbd0b152c.
//
// Solidity: function getWalletAccounts(address _walletRegistry, bytes32 _walletId) view returns((bytes32,string)[] _walletAccounts)
func (_WalletPayments *WalletPaymentsSession) GetWalletAccounts(_walletRegistry common.Address, _walletId [32]byte) ([]WalletAccount, error) {
	return _WalletPayments.Contract.GetWalletAccounts(&_WalletPayments.CallOpts, _walletRegistry, _walletId)
}

// GetWalletAccounts is a free data retrieval call binding the contract method 0xbd0b152c.
//
// Solidity: function getWalletAccounts(address _walletRegistry, bytes32 _walletId) view returns((bytes32,string)[] _walletAccounts)
func (_WalletPayments *WalletPaymentsCallerSession) GetWalletAccounts(_walletRegistry common.Address, _walletId [32]byte) ([]WalletAccount, error) {
	return _WalletPayments.Contract.GetWalletAccounts(&_WalletPayments.CallOpts, _walletRegistry, _walletId)
}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32 _walletId)
func (_WalletPayments *WalletPaymentsCaller) GetWalletId(opts *bind.CallOpts, _account WalletAccount) ([32]byte, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getWalletId", _account)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32 _walletId)
func (_WalletPayments *WalletPaymentsSession) GetWalletId(_account WalletAccount) ([32]byte, error) {
	return _WalletPayments.Contract.GetWalletId(&_WalletPayments.CallOpts, _account)
}

// GetWalletId is a free data retrieval call binding the contract method 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32 _walletId)
func (_WalletPayments *WalletPaymentsCallerSession) GetWalletId(_account WalletAccount) ([32]byte, error) {
	return _WalletPayments.Contract.GetWalletId(&_WalletPayments.CallOpts, _account)
}

// GetWalletRegistry is a free data retrieval call binding the contract method 0x454e7714.
//
// Solidity: function getWalletRegistry((bytes32,string) _account) view returns(address _walletRegistry)
func (_WalletPayments *WalletPaymentsCaller) GetWalletRegistry(opts *bind.CallOpts, _account WalletAccount) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getWalletRegistry", _account)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWalletRegistry is a free data retrieval call binding the contract method 0x454e7714.
//
// Solidity: function getWalletRegistry((bytes32,string) _account) view returns(address _walletRegistry)
func (_WalletPayments *WalletPaymentsSession) GetWalletRegistry(_account WalletAccount) (common.Address, error) {
	return _WalletPayments.Contract.GetWalletRegistry(&_WalletPayments.CallOpts, _account)
}

// GetWalletRegistry is a free data retrieval call binding the contract method 0x454e7714.
//
// Solidity: function getWalletRegistry((bytes32,string) _account) view returns(address _walletRegistry)
func (_WalletPayments *WalletPaymentsCallerSession) GetWalletRegistry(_account WalletAccount) (common.Address, error) {
	return _WalletPayments.Contract.GetWalletRegistry(&_WalletPayments.CallOpts, _account)
}

// GetWalletRegistryKind is a free data retrieval call binding the contract method 0xfcd97026.
//
// Solidity: function getWalletRegistryKind(address _registry) view returns(uint8 _kind)
func (_WalletPayments *WalletPaymentsCaller) GetWalletRegistryKind(opts *bind.CallOpts, _registry common.Address) (uint8, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getWalletRegistryKind", _registry)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetWalletRegistryKind is a free data retrieval call binding the contract method 0xfcd97026.
//
// Solidity: function getWalletRegistryKind(address _registry) view returns(uint8 _kind)
func (_WalletPayments *WalletPaymentsSession) GetWalletRegistryKind(_registry common.Address) (uint8, error) {
	return _WalletPayments.Contract.GetWalletRegistryKind(&_WalletPayments.CallOpts, _registry)
}

// GetWalletRegistryKind is a free data retrieval call binding the contract method 0xfcd97026.
//
// Solidity: function getWalletRegistryKind(address _registry) view returns(uint8 _kind)
func (_WalletPayments *WalletPaymentsCallerSession) GetWalletRegistryKind(_registry common.Address) (uint8, error) {
	return _WalletPayments.Contract.GetWalletRegistryKind(&_WalletPayments.CallOpts, _registry)
}

// GetXrpEscrow is a free data retrieval call binding the contract method 0xf38a8837.
//
// Solidity: function getXrpEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint256,uint64,string,bool) _terms)
func (_WalletPayments *WalletPaymentsCaller) GetXrpEscrow(opts *bind.CallOpts, _account WalletAccount, _paymentId uint64) (IXrpEscrowsXrpEscrowTerms, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getXrpEscrow", _account, _paymentId)

	if err != nil {
		return *new(IXrpEscrowsXrpEscrowTerms), err
	}

	out0 := *abi.ConvertType(out[0], new(IXrpEscrowsXrpEscrowTerms)).(*IXrpEscrowsXrpEscrowTerms)

	return out0, err

}

// GetXrpEscrow is a free data retrieval call binding the contract method 0xf38a8837.
//
// Solidity: function getXrpEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint256,uint64,string,bool) _terms)
func (_WalletPayments *WalletPaymentsSession) GetXrpEscrow(_account WalletAccount, _paymentId uint64) (IXrpEscrowsXrpEscrowTerms, error) {
	return _WalletPayments.Contract.GetXrpEscrow(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetXrpEscrow is a free data retrieval call binding the contract method 0xf38a8837.
//
// Solidity: function getXrpEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint256,uint64,string,bool) _terms)
func (_WalletPayments *WalletPaymentsCallerSession) GetXrpEscrow(_account WalletAccount, _paymentId uint64) (IXrpEscrowsXrpEscrowTerms, error) {
	return _WalletPayments.Contract.GetXrpEscrow(&_WalletPayments.CallOpts, _account, _paymentId)
}

// GetXrpEscrowProfile is a free data retrieval call binding the contract method 0x235fb6a3.
//
// Solidity: function getXrpEscrowProfile((bytes32,string) _account) view returns(string _destination)
func (_WalletPayments *WalletPaymentsCaller) GetXrpEscrowProfile(opts *bind.CallOpts, _account WalletAccount) (string, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "getXrpEscrowProfile", _account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetXrpEscrowProfile is a free data retrieval call binding the contract method 0x235fb6a3.
//
// Solidity: function getXrpEscrowProfile((bytes32,string) _account) view returns(string _destination)
func (_WalletPayments *WalletPaymentsSession) GetXrpEscrowProfile(_account WalletAccount) (string, error) {
	return _WalletPayments.Contract.GetXrpEscrowProfile(&_WalletPayments.CallOpts, _account)
}

// GetXrpEscrowProfile is a free data retrieval call binding the contract method 0x235fb6a3.
//
// Solidity: function getXrpEscrowProfile((bytes32,string) _account) view returns(string _destination)
func (_WalletPayments *WalletPaymentsCallerSession) GetXrpEscrowProfile(_account WalletAccount) (string, error) {
	return _WalletPayments.Contract.GetXrpEscrowProfile(&_WalletPayments.CallOpts, _account)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_WalletPayments *WalletPaymentsCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_WalletPayments *WalletPaymentsSession) Governance() (common.Address, error) {
	return _WalletPayments.Contract.Governance(&_WalletPayments.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_WalletPayments *WalletPaymentsCallerSession) Governance() (common.Address, error) {
	return _WalletPayments.Contract.Governance(&_WalletPayments.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_WalletPayments *WalletPaymentsCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_WalletPayments *WalletPaymentsSession) GovernanceSettings() (common.Address, error) {
	return _WalletPayments.Contract.GovernanceSettings(&_WalletPayments.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_WalletPayments *WalletPaymentsCallerSession) GovernanceSettings() (common.Address, error) {
	return _WalletPayments.Contract.GovernanceSettings(&_WalletPayments.CallOpts)
}

// IsAllowedProposer is a free data retrieval call binding the contract method 0x41cac86b.
//
// Solidity: function isAllowedProposer((bytes32,string) _account, address _proposer) view returns(bool)
func (_WalletPayments *WalletPaymentsCaller) IsAllowedProposer(opts *bind.CallOpts, _account WalletAccount, _proposer common.Address) (bool, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "isAllowedProposer", _account, _proposer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedProposer is a free data retrieval call binding the contract method 0x41cac86b.
//
// Solidity: function isAllowedProposer((bytes32,string) _account, address _proposer) view returns(bool)
func (_WalletPayments *WalletPaymentsSession) IsAllowedProposer(_account WalletAccount, _proposer common.Address) (bool, error) {
	return _WalletPayments.Contract.IsAllowedProposer(&_WalletPayments.CallOpts, _account, _proposer)
}

// IsAllowedProposer is a free data retrieval call binding the contract method 0x41cac86b.
//
// Solidity: function isAllowedProposer((bytes32,string) _account, address _proposer) view returns(bool)
func (_WalletPayments *WalletPaymentsCallerSession) IsAllowedProposer(_account WalletAccount, _proposer common.Address) (bool, error) {
	return _WalletPayments.Contract.IsAllowedProposer(&_WalletPayments.CallOpts, _account, _proposer)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_WalletPayments *WalletPaymentsCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_WalletPayments *WalletPaymentsSession) IsExecutor(_address common.Address) (bool, error) {
	return _WalletPayments.Contract.IsExecutor(&_WalletPayments.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_WalletPayments *WalletPaymentsCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _WalletPayments.Contract.IsExecutor(&_WalletPayments.CallOpts, _address)
}

// IsOperationPaymentId is a free data retrieval call binding the contract method 0xd1230187.
//
// Solidity: function isOperationPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(bool _isOperation)
func (_WalletPayments *WalletPaymentsCaller) IsOperationPaymentId(opts *bind.CallOpts, _account WalletAccount, _paymentId uint64) (bool, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "isOperationPaymentId", _account, _paymentId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperationPaymentId is a free data retrieval call binding the contract method 0xd1230187.
//
// Solidity: function isOperationPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(bool _isOperation)
func (_WalletPayments *WalletPaymentsSession) IsOperationPaymentId(_account WalletAccount, _paymentId uint64) (bool, error) {
	return _WalletPayments.Contract.IsOperationPaymentId(&_WalletPayments.CallOpts, _account, _paymentId)
}

// IsOperationPaymentId is a free data retrieval call binding the contract method 0xd1230187.
//
// Solidity: function isOperationPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(bool _isOperation)
func (_WalletPayments *WalletPaymentsCallerSession) IsOperationPaymentId(_account WalletAccount, _paymentId uint64) (bool, error) {
	return _WalletPayments.Contract.IsOperationPaymentId(&_WalletPayments.CallOpts, _account, _paymentId)
}

// IsSourceEnabled is a free data retrieval call binding the contract method 0xf6d2abbb.
//
// Solidity: function isSourceEnabled(bytes32 _sourceId) view returns(bool _enabled)
func (_WalletPayments *WalletPaymentsCaller) IsSourceEnabled(opts *bind.CallOpts, _sourceId [32]byte) (bool, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "isSourceEnabled", _sourceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSourceEnabled is a free data retrieval call binding the contract method 0xf6d2abbb.
//
// Solidity: function isSourceEnabled(bytes32 _sourceId) view returns(bool _enabled)
func (_WalletPayments *WalletPaymentsSession) IsSourceEnabled(_sourceId [32]byte) (bool, error) {
	return _WalletPayments.Contract.IsSourceEnabled(&_WalletPayments.CallOpts, _sourceId)
}

// IsSourceEnabled is a free data retrieval call binding the contract method 0xf6d2abbb.
//
// Solidity: function isSourceEnabled(bytes32 _sourceId) view returns(bool _enabled)
func (_WalletPayments *WalletPaymentsCallerSession) IsSourceEnabled(_sourceId [32]byte) (bool, error) {
	return _WalletPayments.Contract.IsSourceEnabled(&_WalletPayments.CallOpts, _sourceId)
}

// IsSourceRegistered is a free data retrieval call binding the contract method 0x6848bfa4.
//
// Solidity: function isSourceRegistered(bytes32 _sourceId) view returns(bool _registered)
func (_WalletPayments *WalletPaymentsCaller) IsSourceRegistered(opts *bind.CallOpts, _sourceId [32]byte) (bool, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "isSourceRegistered", _sourceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSourceRegistered is a free data retrieval call binding the contract method 0x6848bfa4.
//
// Solidity: function isSourceRegistered(bytes32 _sourceId) view returns(bool _registered)
func (_WalletPayments *WalletPaymentsSession) IsSourceRegistered(_sourceId [32]byte) (bool, error) {
	return _WalletPayments.Contract.IsSourceRegistered(&_WalletPayments.CallOpts, _sourceId)
}

// IsSourceRegistered is a free data retrieval call binding the contract method 0x6848bfa4.
//
// Solidity: function isSourceRegistered(bytes32 _sourceId) view returns(bool _registered)
func (_WalletPayments *WalletPaymentsCallerSession) IsSourceRegistered(_sourceId [32]byte) (bool, error) {
	return _WalletPayments.Contract.IsSourceRegistered(&_WalletPayments.CallOpts, _sourceId)
}

// IsValidAddress is a free data retrieval call binding the contract method 0xdebecae8.
//
// Solidity: function isValidAddress(bytes32 _sourceId, string _address) view returns(bool _valid)
func (_WalletPayments *WalletPaymentsCaller) IsValidAddress(opts *bind.CallOpts, _sourceId [32]byte, _address string) (bool, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "isValidAddress", _sourceId, _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidAddress is a free data retrieval call binding the contract method 0xdebecae8.
//
// Solidity: function isValidAddress(bytes32 _sourceId, string _address) view returns(bool _valid)
func (_WalletPayments *WalletPaymentsSession) IsValidAddress(_sourceId [32]byte, _address string) (bool, error) {
	return _WalletPayments.Contract.IsValidAddress(&_WalletPayments.CallOpts, _sourceId, _address)
}

// IsValidAddress is a free data retrieval call binding the contract method 0xdebecae8.
//
// Solidity: function isValidAddress(bytes32 _sourceId, string _address) view returns(bool _valid)
func (_WalletPayments *WalletPaymentsCallerSession) IsValidAddress(_sourceId [32]byte, _address string) (bool, error) {
	return _WalletPayments.Contract.IsValidAddress(&_WalletPayments.CallOpts, _sourceId, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_WalletPayments *WalletPaymentsCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_WalletPayments *WalletPaymentsSession) ProductionMode() (bool, error) {
	return _WalletPayments.Contract.ProductionMode(&_WalletPayments.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_WalletPayments *WalletPaymentsCallerSession) ProductionMode() (bool, error) {
	return _WalletPayments.Contract.ProductionMode(&_WalletPayments.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_WalletPayments *WalletPaymentsCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_WalletPayments *WalletPaymentsSession) RewardManager() (common.Address, error) {
	return _WalletPayments.Contract.RewardManager(&_WalletPayments.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_WalletPayments *WalletPaymentsCallerSession) RewardManager() (common.Address, error) {
	return _WalletPayments.Contract.RewardManager(&_WalletPayments.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WalletPayments *WalletPaymentsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WalletPayments *WalletPaymentsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WalletPayments.Contract.SupportsInterface(&_WalletPayments.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_WalletPayments *WalletPaymentsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _WalletPayments.Contract.SupportsInterface(&_WalletPayments.CallOpts, interfaceId)
}

// ValidateAndEncodeSchedules is a free data retrieval call binding the contract method 0x3d699135.
//
// Solidity: function validateAndEncodeSchedules(bytes32 _sourceId, int16[][] _factorsBIPSPerPayment, uint16[] _delaysSeconds) view returns(bytes[] _encodedPerPayment)
func (_WalletPayments *WalletPaymentsCaller) ValidateAndEncodeSchedules(opts *bind.CallOpts, _sourceId [32]byte, _factorsBIPSPerPayment [][]int16, _delaysSeconds []uint16) ([][]byte, error) {
	var out []interface{}
	err := _WalletPayments.contract.Call(opts, &out, "validateAndEncodeSchedules", _sourceId, _factorsBIPSPerPayment, _delaysSeconds)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// ValidateAndEncodeSchedules is a free data retrieval call binding the contract method 0x3d699135.
//
// Solidity: function validateAndEncodeSchedules(bytes32 _sourceId, int16[][] _factorsBIPSPerPayment, uint16[] _delaysSeconds) view returns(bytes[] _encodedPerPayment)
func (_WalletPayments *WalletPaymentsSession) ValidateAndEncodeSchedules(_sourceId [32]byte, _factorsBIPSPerPayment [][]int16, _delaysSeconds []uint16) ([][]byte, error) {
	return _WalletPayments.Contract.ValidateAndEncodeSchedules(&_WalletPayments.CallOpts, _sourceId, _factorsBIPSPerPayment, _delaysSeconds)
}

// ValidateAndEncodeSchedules is a free data retrieval call binding the contract method 0x3d699135.
//
// Solidity: function validateAndEncodeSchedules(bytes32 _sourceId, int16[][] _factorsBIPSPerPayment, uint16[] _delaysSeconds) view returns(bytes[] _encodedPerPayment)
func (_WalletPayments *WalletPaymentsCallerSession) ValidateAndEncodeSchedules(_sourceId [32]byte, _factorsBIPSPerPayment [][]int16, _delaysSeconds []uint16) ([][]byte, error) {
	return _WalletPayments.Contract.ValidateAndEncodeSchedules(&_WalletPayments.CallOpts, _sourceId, _factorsBIPSPerPayment, _delaysSeconds)
}

// AddAnchors is a paid mutator transaction binding the contract method 0x22344637.
//
// Solidity: function addAnchors(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof) returns()
func (_WalletPayments *WalletPaymentsTransactor) AddAnchors(opts *bind.TransactOpts, _proof IBtcAccountConfiguredProof) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "addAnchors", _proof)
}

// AddAnchors is a paid mutator transaction binding the contract method 0x22344637.
//
// Solidity: function addAnchors(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof) returns()
func (_WalletPayments *WalletPaymentsSession) AddAnchors(_proof IBtcAccountConfiguredProof) (*types.Transaction, error) {
	return _WalletPayments.Contract.AddAnchors(&_WalletPayments.TransactOpts, _proof)
}

// AddAnchors is a paid mutator transaction binding the contract method 0x22344637.
//
// Solidity: function addAnchors(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) AddAnchors(_proof IBtcAccountConfiguredProof) (*types.Transaction, error) {
	return _WalletPayments.Contract.AddAnchors(&_WalletPayments.TransactOpts, _proof)
}

// AddBtcAccount is a paid mutator transaction binding the contract method 0x79cf216a.
//
// Solidity: function addBtcAccount(address _walletRegistry, bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof, address _authorizationAddress) returns()
func (_WalletPayments *WalletPaymentsTransactor) AddBtcAccount(opts *bind.TransactOpts, _walletRegistry common.Address, _walletId [32]byte, _proof IBtcAccountConfiguredProof, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "addBtcAccount", _walletRegistry, _walletId, _proof, _authorizationAddress)
}

// AddBtcAccount is a paid mutator transaction binding the contract method 0x79cf216a.
//
// Solidity: function addBtcAccount(address _walletRegistry, bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof, address _authorizationAddress) returns()
func (_WalletPayments *WalletPaymentsSession) AddBtcAccount(_walletRegistry common.Address, _walletId [32]byte, _proof IBtcAccountConfiguredProof, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.AddBtcAccount(&_WalletPayments.TransactOpts, _walletRegistry, _walletId, _proof, _authorizationAddress)
}

// AddBtcAccount is a paid mutator transaction binding the contract method 0x79cf216a.
//
// Solidity: function addBtcAccount(address _walletRegistry, bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof, address _authorizationAddress) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) AddBtcAccount(_walletRegistry common.Address, _walletId [32]byte, _proof IBtcAccountConfiguredProof, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.AddBtcAccount(&_WalletPayments.TransactOpts, _walletRegistry, _walletId, _proof, _authorizationAddress)
}

// AddNativeNonceAccount is a paid mutator transaction binding the contract method 0xb9820a1e.
//
// Solidity: function addNativeNonceAccount(address _walletRegistry, bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof, address _authorizationAddress) returns()
func (_WalletPayments *WalletPaymentsTransactor) AddNativeNonceAccount(opts *bind.TransactOpts, _walletRegistry common.Address, _walletId [32]byte, _proof INativeNonceAccountConfiguredProof, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "addNativeNonceAccount", _walletRegistry, _walletId, _proof, _authorizationAddress)
}

// AddNativeNonceAccount is a paid mutator transaction binding the contract method 0xb9820a1e.
//
// Solidity: function addNativeNonceAccount(address _walletRegistry, bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof, address _authorizationAddress) returns()
func (_WalletPayments *WalletPaymentsSession) AddNativeNonceAccount(_walletRegistry common.Address, _walletId [32]byte, _proof INativeNonceAccountConfiguredProof, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.AddNativeNonceAccount(&_WalletPayments.TransactOpts, _walletRegistry, _walletId, _proof, _authorizationAddress)
}

// AddNativeNonceAccount is a paid mutator transaction binding the contract method 0xb9820a1e.
//
// Solidity: function addNativeNonceAccount(address _walletRegistry, bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof, address _authorizationAddress) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) AddNativeNonceAccount(_walletRegistry common.Address, _walletId [32]byte, _proof INativeNonceAccountConfiguredProof, _authorizationAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.AddNativeNonceAccount(&_WalletPayments.TransactOpts, _walletRegistry, _walletId, _proof, _authorizationAddress)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_WalletPayments *WalletPaymentsTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _encodedCall []byte) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "cancelGovernanceCall", _encodedCall)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_WalletPayments *WalletPaymentsSession) CancelGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.CancelGovernanceCall(&_WalletPayments.TransactOpts, _encodedCall)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) CancelGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.CancelGovernanceCall(&_WalletPayments.TransactOpts, _encodedCall)
}

// ClearAccountFeeSchedule is a paid mutator transaction binding the contract method 0xf7443f84.
//
// Solidity: function clearAccountFeeSchedule((bytes32,string) _account) returns()
func (_WalletPayments *WalletPaymentsTransactor) ClearAccountFeeSchedule(opts *bind.TransactOpts, _account WalletAccount) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "clearAccountFeeSchedule", _account)
}

// ClearAccountFeeSchedule is a paid mutator transaction binding the contract method 0xf7443f84.
//
// Solidity: function clearAccountFeeSchedule((bytes32,string) _account) returns()
func (_WalletPayments *WalletPaymentsSession) ClearAccountFeeSchedule(_account WalletAccount) (*types.Transaction, error) {
	return _WalletPayments.Contract.ClearAccountFeeSchedule(&_WalletPayments.TransactOpts, _account)
}

// ClearAccountFeeSchedule is a paid mutator transaction binding the contract method 0xf7443f84.
//
// Solidity: function clearAccountFeeSchedule((bytes32,string) _account) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) ClearAccountFeeSchedule(_account WalletAccount) (*types.Transaction, error) {
	return _WalletPayments.Contract.ClearAccountFeeSchedule(&_WalletPayments.TransactOpts, _account)
}

// ClearFeeScheduleConfigs is a paid mutator transaction binding the contract method 0x11a5c047.
//
// Solidity: function clearFeeScheduleConfigs(bytes32[] _sourceIds) returns()
func (_WalletPayments *WalletPaymentsTransactor) ClearFeeScheduleConfigs(opts *bind.TransactOpts, _sourceIds [][32]byte) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "clearFeeScheduleConfigs", _sourceIds)
}

// ClearFeeScheduleConfigs is a paid mutator transaction binding the contract method 0x11a5c047.
//
// Solidity: function clearFeeScheduleConfigs(bytes32[] _sourceIds) returns()
func (_WalletPayments *WalletPaymentsSession) ClearFeeScheduleConfigs(_sourceIds [][32]byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.ClearFeeScheduleConfigs(&_WalletPayments.TransactOpts, _sourceIds)
}

// ClearFeeScheduleConfigs is a paid mutator transaction binding the contract method 0x11a5c047.
//
// Solidity: function clearFeeScheduleConfigs(bytes32[] _sourceIds) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) ClearFeeScheduleConfigs(_sourceIds [][32]byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.ClearFeeScheduleConfigs(&_WalletPayments.TransactOpts, _sourceIds)
}

// ClearProjectFeeSchedule is a paid mutator transaction binding the contract method 0x7aecd6f7.
//
// Solidity: function clearProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) returns()
func (_WalletPayments *WalletPaymentsTransactor) ClearProjectFeeSchedule(opts *bind.TransactOpts, _projectId [32]byte, _sourceId [32]byte) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "clearProjectFeeSchedule", _projectId, _sourceId)
}

// ClearProjectFeeSchedule is a paid mutator transaction binding the contract method 0x7aecd6f7.
//
// Solidity: function clearProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) returns()
func (_WalletPayments *WalletPaymentsSession) ClearProjectFeeSchedule(_projectId [32]byte, _sourceId [32]byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.ClearProjectFeeSchedule(&_WalletPayments.TransactOpts, _projectId, _sourceId)
}

// ClearProjectFeeSchedule is a paid mutator transaction binding the contract method 0x7aecd6f7.
//
// Solidity: function clearProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) ClearProjectFeeSchedule(_projectId [32]byte, _sourceId [32]byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.ClearProjectFeeSchedule(&_WalletPayments.TransactOpts, _projectId, _sourceId)
}

// Consolidate is a paid mutator transaction binding the contract method 0x1e5e66ac.
//
// Solidity: function consolidate((bytes32,string) _account, uint8 _mode, uint256 _maxFee) returns(uint64 _paymentId)
func (_WalletPayments *WalletPaymentsTransactor) Consolidate(opts *bind.TransactOpts, _account WalletAccount, _mode uint8, _maxFee *big.Int) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "consolidate", _account, _mode, _maxFee)
}

// Consolidate is a paid mutator transaction binding the contract method 0x1e5e66ac.
//
// Solidity: function consolidate((bytes32,string) _account, uint8 _mode, uint256 _maxFee) returns(uint64 _paymentId)
func (_WalletPayments *WalletPaymentsSession) Consolidate(_account WalletAccount, _mode uint8, _maxFee *big.Int) (*types.Transaction, error) {
	return _WalletPayments.Contract.Consolidate(&_WalletPayments.TransactOpts, _account, _mode, _maxFee)
}

// Consolidate is a paid mutator transaction binding the contract method 0x1e5e66ac.
//
// Solidity: function consolidate((bytes32,string) _account, uint8 _mode, uint256 _maxFee) returns(uint64 _paymentId)
func (_WalletPayments *WalletPaymentsTransactorSession) Consolidate(_account WalletAccount, _mode uint8, _maxFee *big.Int) (*types.Transaction, error) {
	return _WalletPayments.Contract.Consolidate(&_WalletPayments.TransactOpts, _account, _mode, _maxFee)
}

// CreateEscrow is a paid mutator transaction binding the contract method 0xd40003db.
//
// Solidity: function createEscrow((bytes32,string) _account, (bytes32,uint256,uint64) _terms, uint256 _maxFee, address _claimBackAddress) payable returns(uint64 _paymentId)
func (_WalletPayments *WalletPaymentsTransactor) CreateEscrow(opts *bind.TransactOpts, _account WalletAccount, _terms IEscrowsEscrowTerms, _maxFee *big.Int, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "createEscrow", _account, _terms, _maxFee, _claimBackAddress)
}

// CreateEscrow is a paid mutator transaction binding the contract method 0xd40003db.
//
// Solidity: function createEscrow((bytes32,string) _account, (bytes32,uint256,uint64) _terms, uint256 _maxFee, address _claimBackAddress) payable returns(uint64 _paymentId)
func (_WalletPayments *WalletPaymentsSession) CreateEscrow(_account WalletAccount, _terms IEscrowsEscrowTerms, _maxFee *big.Int, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.CreateEscrow(&_WalletPayments.TransactOpts, _account, _terms, _maxFee, _claimBackAddress)
}

// CreateEscrow is a paid mutator transaction binding the contract method 0xd40003db.
//
// Solidity: function createEscrow((bytes32,string) _account, (bytes32,uint256,uint64) _terms, uint256 _maxFee, address _claimBackAddress) payable returns(uint64 _paymentId)
func (_WalletPayments *WalletPaymentsTransactorSession) CreateEscrow(_account WalletAccount, _terms IEscrowsEscrowTerms, _maxFee *big.Int, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.CreateEscrow(&_WalletPayments.TransactOpts, _account, _terms, _maxFee, _claimBackAddress)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_WalletPayments *WalletPaymentsTransactor) DiamondCut(opts *bind.TransactOpts, _diamondCut []IDiamondFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "diamondCut", _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_WalletPayments *WalletPaymentsSession) DiamondCut(_diamondCut []IDiamondFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.DiamondCut(&_WalletPayments.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) DiamondCut(_diamondCut []IDiamondFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.DiamondCut(&_WalletPayments.TransactOpts, _diamondCut, _init, _calldata)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) payable returns()
func (_WalletPayments *WalletPaymentsTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _encodedCall []byte) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "executeGovernanceCall", _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) payable returns()
func (_WalletPayments *WalletPaymentsSession) ExecuteGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.ExecuteGovernanceCall(&_WalletPayments.TransactOpts, _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) payable returns()
func (_WalletPayments *WalletPaymentsTransactorSession) ExecuteGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.ExecuteGovernanceCall(&_WalletPayments.TransactOpts, _encodedCall)
}

// FinalizeProposal is a paid mutator transaction binding the contract method 0xd5c4c7f3.
//
// Solidity: function finalizeProposal(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint64,uint32,uint64,bytes32),(address,uint64,uint32,bytes32)) _proof, (uint32,uint64,uint32,bytes32,bytes32) _commitment) returns()
func (_WalletPayments *WalletPaymentsTransactor) FinalizeProposal(opts *bind.TransactOpts, _proof ICspProposalCheckProof, _commitment IBtcAccountsBtcProposalCommitment) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "finalizeProposal", _proof, _commitment)
}

// FinalizeProposal is a paid mutator transaction binding the contract method 0xd5c4c7f3.
//
// Solidity: function finalizeProposal(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint64,uint32,uint64,bytes32),(address,uint64,uint32,bytes32)) _proof, (uint32,uint64,uint32,bytes32,bytes32) _commitment) returns()
func (_WalletPayments *WalletPaymentsSession) FinalizeProposal(_proof ICspProposalCheckProof, _commitment IBtcAccountsBtcProposalCommitment) (*types.Transaction, error) {
	return _WalletPayments.Contract.FinalizeProposal(&_WalletPayments.TransactOpts, _proof, _commitment)
}

// FinalizeProposal is a paid mutator transaction binding the contract method 0xd5c4c7f3.
//
// Solidity: function finalizeProposal(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint64,uint32,uint64,bytes32),(address,uint64,uint32,bytes32)) _proof, (uint32,uint64,uint32,bytes32,bytes32) _commitment) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) FinalizeProposal(_proof ICspProposalCheckProof, _commitment IBtcAccountsBtcProposalCommitment) (*types.Transaction, error) {
	return _WalletPayments.Contract.FinalizeProposal(&_WalletPayments.TransactOpts, _proof, _commitment)
}

// NullifyOperation is a paid mutator transaction binding the contract method 0x6a33d505.
//
// Solidity: function nullifyOperation((bytes32,string) _account, uint64 _paymentId, uint256 _maxFee, address _claimBackAddress) payable returns(uint32 _attempt)
func (_WalletPayments *WalletPaymentsTransactor) NullifyOperation(opts *bind.TransactOpts, _account WalletAccount, _paymentId uint64, _maxFee *big.Int, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "nullifyOperation", _account, _paymentId, _maxFee, _claimBackAddress)
}

// NullifyOperation is a paid mutator transaction binding the contract method 0x6a33d505.
//
// Solidity: function nullifyOperation((bytes32,string) _account, uint64 _paymentId, uint256 _maxFee, address _claimBackAddress) payable returns(uint32 _attempt)
func (_WalletPayments *WalletPaymentsSession) NullifyOperation(_account WalletAccount, _paymentId uint64, _maxFee *big.Int, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.NullifyOperation(&_WalletPayments.TransactOpts, _account, _paymentId, _maxFee, _claimBackAddress)
}

// NullifyOperation is a paid mutator transaction binding the contract method 0x6a33d505.
//
// Solidity: function nullifyOperation((bytes32,string) _account, uint64 _paymentId, uint256 _maxFee, address _claimBackAddress) payable returns(uint32 _attempt)
func (_WalletPayments *WalletPaymentsTransactorSession) NullifyOperation(_account WalletAccount, _paymentId uint64, _maxFee *big.Int, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.NullifyOperation(&_WalletPayments.TransactOpts, _account, _paymentId, _maxFee, _claimBackAddress)
}

// Pay is a paid mutator transaction binding the contract method 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (_WalletPayments *WalletPaymentsTransactor) Pay(opts *bind.TransactOpts, _account WalletAccount, _paymentInstruction PaymentInstruction, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "pay", _account, _paymentInstruction, _claimBackAddress)
}

// Pay is a paid mutator transaction binding the contract method 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (_WalletPayments *WalletPaymentsSession) Pay(_account WalletAccount, _paymentInstruction PaymentInstruction, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.Pay(&_WalletPayments.TransactOpts, _account, _paymentInstruction, _claimBackAddress)
}

// Pay is a paid mutator transaction binding the contract method 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (_WalletPayments *WalletPaymentsTransactorSession) Pay(_account WalletAccount, _paymentInstruction PaymentInstruction, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.Pay(&_WalletPayments.TransactOpts, _account, _paymentInstruction, _claimBackAddress)
}

// ReclaimEscrow is a paid mutator transaction binding the contract method 0xa6bf742f.
//
// Solidity: function reclaimEscrow((bytes32,string) _account, uint64 _createPaymentId, uint8 _mode, uint256 _maxFee) returns(uint64 _paymentId)
func (_WalletPayments *WalletPaymentsTransactor) ReclaimEscrow(opts *bind.TransactOpts, _account WalletAccount, _createPaymentId uint64, _mode uint8, _maxFee *big.Int) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "reclaimEscrow", _account, _createPaymentId, _mode, _maxFee)
}

// ReclaimEscrow is a paid mutator transaction binding the contract method 0xa6bf742f.
//
// Solidity: function reclaimEscrow((bytes32,string) _account, uint64 _createPaymentId, uint8 _mode, uint256 _maxFee) returns(uint64 _paymentId)
func (_WalletPayments *WalletPaymentsSession) ReclaimEscrow(_account WalletAccount, _createPaymentId uint64, _mode uint8, _maxFee *big.Int) (*types.Transaction, error) {
	return _WalletPayments.Contract.ReclaimEscrow(&_WalletPayments.TransactOpts, _account, _createPaymentId, _mode, _maxFee)
}

// ReclaimEscrow is a paid mutator transaction binding the contract method 0xa6bf742f.
//
// Solidity: function reclaimEscrow((bytes32,string) _account, uint64 _createPaymentId, uint8 _mode, uint256 _maxFee) returns(uint64 _paymentId)
func (_WalletPayments *WalletPaymentsTransactorSession) ReclaimEscrow(_account WalletAccount, _createPaymentId uint64, _mode uint8, _maxFee *big.Int) (*types.Transaction, error) {
	return _WalletPayments.Contract.ReclaimEscrow(&_WalletPayments.TransactOpts, _account, _createPaymentId, _mode, _maxFee)
}

// RegisterSources is a paid mutator transaction binding the contract method 0xf8c90fb6.
//
// Solidity: function registerSources((bytes32,bytes32,bytes32,uint8,uint8)[] _registrations) returns()
func (_WalletPayments *WalletPaymentsTransactor) RegisterSources(opts *bind.TransactOpts, _registrations []ISourceConfigSourceRegistration) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "registerSources", _registrations)
}

// RegisterSources is a paid mutator transaction binding the contract method 0xf8c90fb6.
//
// Solidity: function registerSources((bytes32,bytes32,bytes32,uint8,uint8)[] _registrations) returns()
func (_WalletPayments *WalletPaymentsSession) RegisterSources(_registrations []ISourceConfigSourceRegistration) (*types.Transaction, error) {
	return _WalletPayments.Contract.RegisterSources(&_WalletPayments.TransactOpts, _registrations)
}

// RegisterSources is a paid mutator transaction binding the contract method 0xf8c90fb6.
//
// Solidity: function registerSources((bytes32,bytes32,bytes32,uint8,uint8)[] _registrations) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) RegisterSources(_registrations []ISourceConfigSourceRegistration) (*types.Transaction, error) {
	return _WalletPayments.Contract.RegisterSources(&_WalletPayments.TransactOpts, _registrations)
}

// Reissue is a paid mutator transaction binding the contract method 0x5cc0a260.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _paymentId, (string,bytes,uint256,uint256,bytes32)[] _retainedInstructions, uint64[] _nullifiedPaymentIds, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns(bool _finalized)
func (_WalletPayments *WalletPaymentsTransactor) Reissue(opts *bind.TransactOpts, _account WalletAccount, _paymentId uint64, _retainedInstructions []PaymentInstruction, _nullifiedPaymentIds []uint64, _reissueFeeParams ReissueFeeParams, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "reissue", _account, _paymentId, _retainedInstructions, _nullifiedPaymentIds, _reissueFeeParams, _claimBackAddress)
}

// Reissue is a paid mutator transaction binding the contract method 0x5cc0a260.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _paymentId, (string,bytes,uint256,uint256,bytes32)[] _retainedInstructions, uint64[] _nullifiedPaymentIds, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns(bool _finalized)
func (_WalletPayments *WalletPaymentsSession) Reissue(_account WalletAccount, _paymentId uint64, _retainedInstructions []PaymentInstruction, _nullifiedPaymentIds []uint64, _reissueFeeParams ReissueFeeParams, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.Reissue(&_WalletPayments.TransactOpts, _account, _paymentId, _retainedInstructions, _nullifiedPaymentIds, _reissueFeeParams, _claimBackAddress)
}

// Reissue is a paid mutator transaction binding the contract method 0x5cc0a260.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _paymentId, (string,bytes,uint256,uint256,bytes32)[] _retainedInstructions, uint64[] _nullifiedPaymentIds, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns(bool _finalized)
func (_WalletPayments *WalletPaymentsTransactorSession) Reissue(_account WalletAccount, _paymentId uint64, _retainedInstructions []PaymentInstruction, _nullifiedPaymentIds []uint64, _reissueFeeParams ReissueFeeParams, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.Reissue(&_WalletPayments.TransactOpts, _account, _paymentId, _retainedInstructions, _nullifiedPaymentIds, _reissueFeeParams, _claimBackAddress)
}

// ReissueOperation is a paid mutator transaction binding the contract method 0xf10cc52c.
//
// Solidity: function reissueOperation((bytes32,string) _account, uint64 _paymentId, uint256 _maxFee, address _claimBackAddress) payable returns(uint32 _attempt)
func (_WalletPayments *WalletPaymentsTransactor) ReissueOperation(opts *bind.TransactOpts, _account WalletAccount, _paymentId uint64, _maxFee *big.Int, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "reissueOperation", _account, _paymentId, _maxFee, _claimBackAddress)
}

// ReissueOperation is a paid mutator transaction binding the contract method 0xf10cc52c.
//
// Solidity: function reissueOperation((bytes32,string) _account, uint64 _paymentId, uint256 _maxFee, address _claimBackAddress) payable returns(uint32 _attempt)
func (_WalletPayments *WalletPaymentsSession) ReissueOperation(_account WalletAccount, _paymentId uint64, _maxFee *big.Int, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.ReissueOperation(&_WalletPayments.TransactOpts, _account, _paymentId, _maxFee, _claimBackAddress)
}

// ReissueOperation is a paid mutator transaction binding the contract method 0xf10cc52c.
//
// Solidity: function reissueOperation((bytes32,string) _account, uint64 _paymentId, uint256 _maxFee, address _claimBackAddress) payable returns(uint32 _attempt)
func (_WalletPayments *WalletPaymentsTransactorSession) ReissueOperation(_account WalletAccount, _paymentId uint64, _maxFee *big.Int, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.ReissueOperation(&_WalletPayments.TransactOpts, _account, _paymentId, _maxFee, _claimBackAddress)
}

// RemoveAccountBtcWalletParams is a paid mutator transaction binding the contract method 0x1c54c7cd.
//
// Solidity: function removeAccountBtcWalletParams((bytes32,string) _account) returns()
func (_WalletPayments *WalletPaymentsTransactor) RemoveAccountBtcWalletParams(opts *bind.TransactOpts, _account WalletAccount) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "removeAccountBtcWalletParams", _account)
}

// RemoveAccountBtcWalletParams is a paid mutator transaction binding the contract method 0x1c54c7cd.
//
// Solidity: function removeAccountBtcWalletParams((bytes32,string) _account) returns()
func (_WalletPayments *WalletPaymentsSession) RemoveAccountBtcWalletParams(_account WalletAccount) (*types.Transaction, error) {
	return _WalletPayments.Contract.RemoveAccountBtcWalletParams(&_WalletPayments.TransactOpts, _account)
}

// RemoveAccountBtcWalletParams is a paid mutator transaction binding the contract method 0x1c54c7cd.
//
// Solidity: function removeAccountBtcWalletParams((bytes32,string) _account) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) RemoveAccountBtcWalletParams(_account WalletAccount) (*types.Transaction, error) {
	return _WalletPayments.Contract.RemoveAccountBtcWalletParams(&_WalletPayments.TransactOpts, _account)
}

// RemoveAccountProposers is a paid mutator transaction binding the contract method 0x379ae241.
//
// Solidity: function removeAccountProposers((bytes32,string) _account) returns()
func (_WalletPayments *WalletPaymentsTransactor) RemoveAccountProposers(opts *bind.TransactOpts, _account WalletAccount) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "removeAccountProposers", _account)
}

// RemoveAccountProposers is a paid mutator transaction binding the contract method 0x379ae241.
//
// Solidity: function removeAccountProposers((bytes32,string) _account) returns()
func (_WalletPayments *WalletPaymentsSession) RemoveAccountProposers(_account WalletAccount) (*types.Transaction, error) {
	return _WalletPayments.Contract.RemoveAccountProposers(&_WalletPayments.TransactOpts, _account)
}

// RemoveAccountProposers is a paid mutator transaction binding the contract method 0x379ae241.
//
// Solidity: function removeAccountProposers((bytes32,string) _account) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) RemoveAccountProposers(_account WalletAccount) (*types.Transaction, error) {
	return _WalletPayments.Contract.RemoveAccountProposers(&_WalletPayments.TransactOpts, _account)
}

// RemoveProjectBtcWalletParams is a paid mutator transaction binding the contract method 0x65250799.
//
// Solidity: function removeProjectBtcWalletParams(bytes32 _projectId) returns()
func (_WalletPayments *WalletPaymentsTransactor) RemoveProjectBtcWalletParams(opts *bind.TransactOpts, _projectId [32]byte) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "removeProjectBtcWalletParams", _projectId)
}

// RemoveProjectBtcWalletParams is a paid mutator transaction binding the contract method 0x65250799.
//
// Solidity: function removeProjectBtcWalletParams(bytes32 _projectId) returns()
func (_WalletPayments *WalletPaymentsSession) RemoveProjectBtcWalletParams(_projectId [32]byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.RemoveProjectBtcWalletParams(&_WalletPayments.TransactOpts, _projectId)
}

// RemoveProjectBtcWalletParams is a paid mutator transaction binding the contract method 0x65250799.
//
// Solidity: function removeProjectBtcWalletParams(bytes32 _projectId) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) RemoveProjectBtcWalletParams(_projectId [32]byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.RemoveProjectBtcWalletParams(&_WalletPayments.TransactOpts, _projectId)
}

// RemoveProjectProposers is a paid mutator transaction binding the contract method 0x3304f19c.
//
// Solidity: function removeProjectProposers(bytes32 _projectId) returns()
func (_WalletPayments *WalletPaymentsTransactor) RemoveProjectProposers(opts *bind.TransactOpts, _projectId [32]byte) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "removeProjectProposers", _projectId)
}

// RemoveProjectProposers is a paid mutator transaction binding the contract method 0x3304f19c.
//
// Solidity: function removeProjectProposers(bytes32 _projectId) returns()
func (_WalletPayments *WalletPaymentsSession) RemoveProjectProposers(_projectId [32]byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.RemoveProjectProposers(&_WalletPayments.TransactOpts, _projectId)
}

// RemoveProjectProposers is a paid mutator transaction binding the contract method 0x3304f19c.
//
// Solidity: function removeProjectProposers(bytes32 _projectId) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) RemoveProjectProposers(_projectId [32]byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.RemoveProjectProposers(&_WalletPayments.TransactOpts, _projectId)
}

// RequestBtcAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0xf2945a64.
//
// Solidity: function requestBtcAccountConfiguredAttestation(address _walletRegistry, bytes32 _walletId, bytes32 _sourceId, uint32 _accountIndex, (bytes32,uint32)[] _anchors, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_WalletPayments *WalletPaymentsTransactor) RequestBtcAccountConfiguredAttestation(opts *bind.TransactOpts, _walletRegistry common.Address, _walletId [32]byte, _sourceId [32]byte, _accountIndex uint32, _anchors []IBtcAccountConfiguredAnchor, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "requestBtcAccountConfiguredAttestation", _walletRegistry, _walletId, _sourceId, _accountIndex, _anchors, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestBtcAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0xf2945a64.
//
// Solidity: function requestBtcAccountConfiguredAttestation(address _walletRegistry, bytes32 _walletId, bytes32 _sourceId, uint32 _accountIndex, (bytes32,uint32)[] _anchors, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_WalletPayments *WalletPaymentsSession) RequestBtcAccountConfiguredAttestation(_walletRegistry common.Address, _walletId [32]byte, _sourceId [32]byte, _accountIndex uint32, _anchors []IBtcAccountConfiguredAnchor, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.RequestBtcAccountConfiguredAttestation(&_WalletPayments.TransactOpts, _walletRegistry, _walletId, _sourceId, _accountIndex, _anchors, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestBtcAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0xf2945a64.
//
// Solidity: function requestBtcAccountConfiguredAttestation(address _walletRegistry, bytes32 _walletId, bytes32 _sourceId, uint32 _accountIndex, (bytes32,uint32)[] _anchors, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_WalletPayments *WalletPaymentsTransactorSession) RequestBtcAccountConfiguredAttestation(_walletRegistry common.Address, _walletId [32]byte, _sourceId [32]byte, _accountIndex uint32, _anchors []IBtcAccountConfiguredAnchor, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.RequestBtcAccountConfiguredAttestation(&_WalletPayments.TransactOpts, _walletRegistry, _walletId, _sourceId, _accountIndex, _anchors, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestNativeNonceAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0x975aaa50.
//
// Solidity: function requestNativeNonceAccountConfiguredAttestation(address _walletRegistry, bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_WalletPayments *WalletPaymentsTransactor) RequestNativeNonceAccountConfiguredAttestation(opts *bind.TransactOpts, _walletRegistry common.Address, _walletId [32]byte, _sourceId [32]byte, _accountAddress string, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "requestNativeNonceAccountConfiguredAttestation", _walletRegistry, _walletId, _sourceId, _accountAddress, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestNativeNonceAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0x975aaa50.
//
// Solidity: function requestNativeNonceAccountConfiguredAttestation(address _walletRegistry, bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_WalletPayments *WalletPaymentsSession) RequestNativeNonceAccountConfiguredAttestation(_walletRegistry common.Address, _walletId [32]byte, _sourceId [32]byte, _accountAddress string, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.RequestNativeNonceAccountConfiguredAttestation(&_WalletPayments.TransactOpts, _walletRegistry, _walletId, _sourceId, _accountAddress, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestNativeNonceAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0x975aaa50.
//
// Solidity: function requestNativeNonceAccountConfiguredAttestation(address _walletRegistry, bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_WalletPayments *WalletPaymentsTransactorSession) RequestNativeNonceAccountConfiguredAttestation(_walletRegistry common.Address, _walletId [32]byte, _sourceId [32]byte, _accountAddress string, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.RequestNativeNonceAccountConfiguredAttestation(&_WalletPayments.TransactOpts, _walletRegistry, _walletId, _sourceId, _accountAddress, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestSigning is a paid mutator transaction binding the contract method 0xc26145d2.
//
// Solidity: function requestSigning((bytes32,string) _account, uint64 _sequencePosition) payable returns()
func (_WalletPayments *WalletPaymentsTransactor) RequestSigning(opts *bind.TransactOpts, _account WalletAccount, _sequencePosition uint64) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "requestSigning", _account, _sequencePosition)
}

// RequestSigning is a paid mutator transaction binding the contract method 0xc26145d2.
//
// Solidity: function requestSigning((bytes32,string) _account, uint64 _sequencePosition) payable returns()
func (_WalletPayments *WalletPaymentsSession) RequestSigning(_account WalletAccount, _sequencePosition uint64) (*types.Transaction, error) {
	return _WalletPayments.Contract.RequestSigning(&_WalletPayments.TransactOpts, _account, _sequencePosition)
}

// RequestSigning is a paid mutator transaction binding the contract method 0xc26145d2.
//
// Solidity: function requestSigning((bytes32,string) _account, uint64 _sequencePosition) payable returns()
func (_WalletPayments *WalletPaymentsTransactorSession) RequestSigning(_account WalletAccount, _sequencePosition uint64) (*types.Transaction, error) {
	return _WalletPayments.Contract.RequestSigning(&_WalletPayments.TransactOpts, _account, _sequencePosition)
}

// SetAccountBtcWalletParams is a paid mutator transaction binding the contract method 0x95841d68.
//
// Solidity: function setAccountBtcWalletParams((bytes32,string) _account, (uint64,uint32,uint8,uint8) _params) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetAccountBtcWalletParams(opts *bind.TransactOpts, _account WalletAccount, _params IBtcParamsBtcWalletParams) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setAccountBtcWalletParams", _account, _params)
}

// SetAccountBtcWalletParams is a paid mutator transaction binding the contract method 0x95841d68.
//
// Solidity: function setAccountBtcWalletParams((bytes32,string) _account, (uint64,uint32,uint8,uint8) _params) returns()
func (_WalletPayments *WalletPaymentsSession) SetAccountBtcWalletParams(_account WalletAccount, _params IBtcParamsBtcWalletParams) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetAccountBtcWalletParams(&_WalletPayments.TransactOpts, _account, _params)
}

// SetAccountBtcWalletParams is a paid mutator transaction binding the contract method 0x95841d68.
//
// Solidity: function setAccountBtcWalletParams((bytes32,string) _account, (uint64,uint32,uint8,uint8) _params) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetAccountBtcWalletParams(_account WalletAccount, _params IBtcParamsBtcWalletParams) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetAccountBtcWalletParams(&_WalletPayments.TransactOpts, _account, _params)
}

// SetAccountFeeSchedule is a paid mutator transaction binding the contract method 0x32ab4577.
//
// Solidity: function setAccountFeeSchedule((bytes32,string) _account, (int16,uint16)[] _schedule) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetAccountFeeSchedule(opts *bind.TransactOpts, _account WalletAccount, _schedule []IFeeSchedulesFeeSchedule) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setAccountFeeSchedule", _account, _schedule)
}

// SetAccountFeeSchedule is a paid mutator transaction binding the contract method 0x32ab4577.
//
// Solidity: function setAccountFeeSchedule((bytes32,string) _account, (int16,uint16)[] _schedule) returns()
func (_WalletPayments *WalletPaymentsSession) SetAccountFeeSchedule(_account WalletAccount, _schedule []IFeeSchedulesFeeSchedule) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetAccountFeeSchedule(&_WalletPayments.TransactOpts, _account, _schedule)
}

// SetAccountFeeSchedule is a paid mutator transaction binding the contract method 0x32ab4577.
//
// Solidity: function setAccountFeeSchedule((bytes32,string) _account, (int16,uint16)[] _schedule) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetAccountFeeSchedule(_account WalletAccount, _schedule []IFeeSchedulesFeeSchedule) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetAccountFeeSchedule(&_WalletPayments.TransactOpts, _account, _schedule)
}

// SetAccountProposers is a paid mutator transaction binding the contract method 0x23f3826f.
//
// Solidity: function setAccountProposers((bytes32,string) _account, address[] _proposers) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetAccountProposers(opts *bind.TransactOpts, _account WalletAccount, _proposers []common.Address) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setAccountProposers", _account, _proposers)
}

// SetAccountProposers is a paid mutator transaction binding the contract method 0x23f3826f.
//
// Solidity: function setAccountProposers((bytes32,string) _account, address[] _proposers) returns()
func (_WalletPayments *WalletPaymentsSession) SetAccountProposers(_account WalletAccount, _proposers []common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetAccountProposers(&_WalletPayments.TransactOpts, _account, _proposers)
}

// SetAccountProposers is a paid mutator transaction binding the contract method 0x23f3826f.
//
// Solidity: function setAccountProposers((bytes32,string) _account, address[] _proposers) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetAccountProposers(_account WalletAccount, _proposers []common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetAccountProposers(&_WalletPayments.TransactOpts, _account, _proposers)
}

// SetBtcConsensusParams is a paid mutator transaction binding the contract method 0x29935ff6.
//
// Solidity: function setBtcConsensusParams(bytes32 _sourceId, (uint16,uint64,uint64,uint32,uint8,uint8,uint8,uint8,uint16,uint32,uint32,uint32,uint32,uint32) _params) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetBtcConsensusParams(opts *bind.TransactOpts, _sourceId [32]byte, _params IBtcParamsBtcConsensusParams) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setBtcConsensusParams", _sourceId, _params)
}

// SetBtcConsensusParams is a paid mutator transaction binding the contract method 0x29935ff6.
//
// Solidity: function setBtcConsensusParams(bytes32 _sourceId, (uint16,uint64,uint64,uint32,uint8,uint8,uint8,uint8,uint16,uint32,uint32,uint32,uint32,uint32) _params) returns()
func (_WalletPayments *WalletPaymentsSession) SetBtcConsensusParams(_sourceId [32]byte, _params IBtcParamsBtcConsensusParams) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetBtcConsensusParams(&_WalletPayments.TransactOpts, _sourceId, _params)
}

// SetBtcConsensusParams is a paid mutator transaction binding the contract method 0x29935ff6.
//
// Solidity: function setBtcConsensusParams(bytes32 _sourceId, (uint16,uint64,uint64,uint32,uint8,uint8,uint8,uint8,uint16,uint32,uint32,uint32,uint32,uint32) _params) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetBtcConsensusParams(_sourceId [32]byte, _params IBtcParamsBtcConsensusParams) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetBtcConsensusParams(&_WalletPayments.TransactOpts, _sourceId, _params)
}

// SetBtcEscrowProfile is a paid mutator transaction binding the contract method 0x0fd69b17.
//
// Solidity: function setBtcEscrowProfile((bytes32,string) _account, bytes _counterpartyPubKey) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetBtcEscrowProfile(opts *bind.TransactOpts, _account WalletAccount, _counterpartyPubKey []byte) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setBtcEscrowProfile", _account, _counterpartyPubKey)
}

// SetBtcEscrowProfile is a paid mutator transaction binding the contract method 0x0fd69b17.
//
// Solidity: function setBtcEscrowProfile((bytes32,string) _account, bytes _counterpartyPubKey) returns()
func (_WalletPayments *WalletPaymentsSession) SetBtcEscrowProfile(_account WalletAccount, _counterpartyPubKey []byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetBtcEscrowProfile(&_WalletPayments.TransactOpts, _account, _counterpartyPubKey)
}

// SetBtcEscrowProfile is a paid mutator transaction binding the contract method 0x0fd69b17.
//
// Solidity: function setBtcEscrowProfile((bytes32,string) _account, bytes _counterpartyPubKey) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetBtcEscrowProfile(_account WalletAccount, _counterpartyPubKey []byte) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetBtcEscrowProfile(&_WalletPayments.TransactOpts, _account, _counterpartyPubKey)
}

// SetCspSourceSettings is a paid mutator transaction binding the contract method 0x37d330a0.
//
// Solidity: function setCspSourceSettings(bytes32 _sourceId, (uint64,uint16,uint16,uint8,uint128) _settings) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetCspSourceSettings(opts *bind.TransactOpts, _sourceId [32]byte, _settings ICspProposalsCspSourceSettings) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setCspSourceSettings", _sourceId, _settings)
}

// SetCspSourceSettings is a paid mutator transaction binding the contract method 0x37d330a0.
//
// Solidity: function setCspSourceSettings(bytes32 _sourceId, (uint64,uint16,uint16,uint8,uint128) _settings) returns()
func (_WalletPayments *WalletPaymentsSession) SetCspSourceSettings(_sourceId [32]byte, _settings ICspProposalsCspSourceSettings) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetCspSourceSettings(&_WalletPayments.TransactOpts, _sourceId, _settings)
}

// SetCspSourceSettings is a paid mutator transaction binding the contract method 0x37d330a0.
//
// Solidity: function setCspSourceSettings(bytes32 _sourceId, (uint64,uint16,uint16,uint8,uint128) _settings) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetCspSourceSettings(_sourceId [32]byte, _settings ICspProposalsCspSourceSettings) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetCspSourceSettings(&_WalletPayments.TransactOpts, _sourceId, _settings)
}

// SetFeeScheduleConfigs is a paid mutator transaction binding the contract method 0x9b7f8f29.
//
// Solidity: function setFeeScheduleConfigs((uint16,uint8,bytes32)[] _configs) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetFeeScheduleConfigs(opts *bind.TransactOpts, _configs []IFeeSchedulesFeeScheduleConfigInput) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setFeeScheduleConfigs", _configs)
}

// SetFeeScheduleConfigs is a paid mutator transaction binding the contract method 0x9b7f8f29.
//
// Solidity: function setFeeScheduleConfigs((uint16,uint8,bytes32)[] _configs) returns()
func (_WalletPayments *WalletPaymentsSession) SetFeeScheduleConfigs(_configs []IFeeSchedulesFeeScheduleConfigInput) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetFeeScheduleConfigs(&_WalletPayments.TransactOpts, _configs)
}

// SetFeeScheduleConfigs is a paid mutator transaction binding the contract method 0x9b7f8f29.
//
// Solidity: function setFeeScheduleConfigs((uint16,uint8,bytes32)[] _configs) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetFeeScheduleConfigs(_configs []IFeeSchedulesFeeScheduleConfigInput) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetFeeScheduleConfigs(&_WalletPayments.TransactOpts, _configs)
}

// SetProjectBtcWalletParams is a paid mutator transaction binding the contract method 0xf688654f.
//
// Solidity: function setProjectBtcWalletParams(bytes32 _projectId, (uint64,uint32,uint8,uint8) _params) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetProjectBtcWalletParams(opts *bind.TransactOpts, _projectId [32]byte, _params IBtcParamsBtcWalletParams) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setProjectBtcWalletParams", _projectId, _params)
}

// SetProjectBtcWalletParams is a paid mutator transaction binding the contract method 0xf688654f.
//
// Solidity: function setProjectBtcWalletParams(bytes32 _projectId, (uint64,uint32,uint8,uint8) _params) returns()
func (_WalletPayments *WalletPaymentsSession) SetProjectBtcWalletParams(_projectId [32]byte, _params IBtcParamsBtcWalletParams) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetProjectBtcWalletParams(&_WalletPayments.TransactOpts, _projectId, _params)
}

// SetProjectBtcWalletParams is a paid mutator transaction binding the contract method 0xf688654f.
//
// Solidity: function setProjectBtcWalletParams(bytes32 _projectId, (uint64,uint32,uint8,uint8) _params) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetProjectBtcWalletParams(_projectId [32]byte, _params IBtcParamsBtcWalletParams) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetProjectBtcWalletParams(&_WalletPayments.TransactOpts, _projectId, _params)
}

// SetProjectFeeSchedule is a paid mutator transaction binding the contract method 0xe27d19bd.
//
// Solidity: function setProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId, (int16,uint16)[] _schedule) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetProjectFeeSchedule(opts *bind.TransactOpts, _projectId [32]byte, _sourceId [32]byte, _schedule []IFeeSchedulesFeeSchedule) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setProjectFeeSchedule", _projectId, _sourceId, _schedule)
}

// SetProjectFeeSchedule is a paid mutator transaction binding the contract method 0xe27d19bd.
//
// Solidity: function setProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId, (int16,uint16)[] _schedule) returns()
func (_WalletPayments *WalletPaymentsSession) SetProjectFeeSchedule(_projectId [32]byte, _sourceId [32]byte, _schedule []IFeeSchedulesFeeSchedule) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetProjectFeeSchedule(&_WalletPayments.TransactOpts, _projectId, _sourceId, _schedule)
}

// SetProjectFeeSchedule is a paid mutator transaction binding the contract method 0xe27d19bd.
//
// Solidity: function setProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId, (int16,uint16)[] _schedule) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetProjectFeeSchedule(_projectId [32]byte, _sourceId [32]byte, _schedule []IFeeSchedulesFeeSchedule) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetProjectFeeSchedule(&_WalletPayments.TransactOpts, _projectId, _sourceId, _schedule)
}

// SetProjectProposers is a paid mutator transaction binding the contract method 0xe902206d.
//
// Solidity: function setProjectProposers(bytes32 _projectId, address[] _proposers) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetProjectProposers(opts *bind.TransactOpts, _projectId [32]byte, _proposers []common.Address) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setProjectProposers", _projectId, _proposers)
}

// SetProjectProposers is a paid mutator transaction binding the contract method 0xe902206d.
//
// Solidity: function setProjectProposers(bytes32 _projectId, address[] _proposers) returns()
func (_WalletPayments *WalletPaymentsSession) SetProjectProposers(_projectId [32]byte, _proposers []common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetProjectProposers(&_WalletPayments.TransactOpts, _projectId, _proposers)
}

// SetProjectProposers is a paid mutator transaction binding the contract method 0xe902206d.
//
// Solidity: function setProjectProposers(bytes32 _projectId, address[] _proposers) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetProjectProposers(_projectId [32]byte, _proposers []common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetProjectProposers(&_WalletPayments.TransactOpts, _projectId, _proposers)
}

// SetProposerUrl is a paid mutator transaction binding the contract method 0xa1808899.
//
// Solidity: function setProposerUrl(string _url) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetProposerUrl(opts *bind.TransactOpts, _url string) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setProposerUrl", _url)
}

// SetProposerUrl is a paid mutator transaction binding the contract method 0xa1808899.
//
// Solidity: function setProposerUrl(string _url) returns()
func (_WalletPayments *WalletPaymentsSession) SetProposerUrl(_url string) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetProposerUrl(&_WalletPayments.TransactOpts, _url)
}

// SetProposerUrl is a paid mutator transaction binding the contract method 0xa1808899.
//
// Solidity: function setProposerUrl(string _url) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetProposerUrl(_url string) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetProposerUrl(&_WalletPayments.TransactOpts, _url)
}

// SetSourceEscrowsEnabled is a paid mutator transaction binding the contract method 0xdb6411e9.
//
// Solidity: function setSourceEscrowsEnabled(bytes32 _sourceId, bool _custodianWallets, bool _teeWallets) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetSourceEscrowsEnabled(opts *bind.TransactOpts, _sourceId [32]byte, _custodianWallets bool, _teeWallets bool) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setSourceEscrowsEnabled", _sourceId, _custodianWallets, _teeWallets)
}

// SetSourceEscrowsEnabled is a paid mutator transaction binding the contract method 0xdb6411e9.
//
// Solidity: function setSourceEscrowsEnabled(bytes32 _sourceId, bool _custodianWallets, bool _teeWallets) returns()
func (_WalletPayments *WalletPaymentsSession) SetSourceEscrowsEnabled(_sourceId [32]byte, _custodianWallets bool, _teeWallets bool) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetSourceEscrowsEnabled(&_WalletPayments.TransactOpts, _sourceId, _custodianWallets, _teeWallets)
}

// SetSourceEscrowsEnabled is a paid mutator transaction binding the contract method 0xdb6411e9.
//
// Solidity: function setSourceEscrowsEnabled(bytes32 _sourceId, bool _custodianWallets, bool _teeWallets) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetSourceEscrowsEnabled(_sourceId [32]byte, _custodianWallets bool, _teeWallets bool) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetSourceEscrowsEnabled(&_WalletPayments.TransactOpts, _sourceId, _custodianWallets, _teeWallets)
}

// SetSourcesEnabled is a paid mutator transaction binding the contract method 0x62825a23.
//
// Solidity: function setSourcesEnabled(bytes32[] _sourceIds, bool _enabled) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetSourcesEnabled(opts *bind.TransactOpts, _sourceIds [][32]byte, _enabled bool) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setSourcesEnabled", _sourceIds, _enabled)
}

// SetSourcesEnabled is a paid mutator transaction binding the contract method 0x62825a23.
//
// Solidity: function setSourcesEnabled(bytes32[] _sourceIds, bool _enabled) returns()
func (_WalletPayments *WalletPaymentsSession) SetSourcesEnabled(_sourceIds [][32]byte, _enabled bool) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetSourcesEnabled(&_WalletPayments.TransactOpts, _sourceIds, _enabled)
}

// SetSourcesEnabled is a paid mutator transaction binding the contract method 0x62825a23.
//
// Solidity: function setSourcesEnabled(bytes32[] _sourceIds, bool _enabled) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetSourcesEnabled(_sourceIds [][32]byte, _enabled bool) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetSourcesEnabled(&_WalletPayments.TransactOpts, _sourceIds, _enabled)
}

// SetXrpEscrowProfile is a paid mutator transaction binding the contract method 0xd2c0c050.
//
// Solidity: function setXrpEscrowProfile((bytes32,string) _account, string _destination) returns()
func (_WalletPayments *WalletPaymentsTransactor) SetXrpEscrowProfile(opts *bind.TransactOpts, _account WalletAccount, _destination string) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "setXrpEscrowProfile", _account, _destination)
}

// SetXrpEscrowProfile is a paid mutator transaction binding the contract method 0xd2c0c050.
//
// Solidity: function setXrpEscrowProfile((bytes32,string) _account, string _destination) returns()
func (_WalletPayments *WalletPaymentsSession) SetXrpEscrowProfile(_account WalletAccount, _destination string) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetXrpEscrowProfile(&_WalletPayments.TransactOpts, _account, _destination)
}

// SetXrpEscrowProfile is a paid mutator transaction binding the contract method 0xd2c0c050.
//
// Solidity: function setXrpEscrowProfile((bytes32,string) _account, string _destination) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SetXrpEscrowProfile(_account WalletAccount, _destination string) (*types.Transaction, error) {
	return _WalletPayments.Contract.SetXrpEscrowProfile(&_WalletPayments.TransactOpts, _account, _destination)
}

// Settle is a paid mutator transaction binding the contract method 0xf5eb3083.
//
// Solidity: function settle((bytes32,string) _account, uint64 _sequencePosition) payable returns()
func (_WalletPayments *WalletPaymentsTransactor) Settle(opts *bind.TransactOpts, _account WalletAccount, _sequencePosition uint64) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "settle", _account, _sequencePosition)
}

// Settle is a paid mutator transaction binding the contract method 0xf5eb3083.
//
// Solidity: function settle((bytes32,string) _account, uint64 _sequencePosition) payable returns()
func (_WalletPayments *WalletPaymentsSession) Settle(_account WalletAccount, _sequencePosition uint64) (*types.Transaction, error) {
	return _WalletPayments.Contract.Settle(&_WalletPayments.TransactOpts, _account, _sequencePosition)
}

// Settle is a paid mutator transaction binding the contract method 0xf5eb3083.
//
// Solidity: function settle((bytes32,string) _account, uint64 _sequencePosition) payable returns()
func (_WalletPayments *WalletPaymentsTransactorSession) Settle(_account WalletAccount, _sequencePosition uint64) (*types.Transaction, error) {
	return _WalletPayments.Contract.Settle(&_WalletPayments.TransactOpts, _account, _sequencePosition)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_WalletPayments *WalletPaymentsTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_WalletPayments *WalletPaymentsSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _WalletPayments.Contract.SwitchToProductionMode(&_WalletPayments.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_WalletPayments *WalletPaymentsTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _WalletPayments.Contract.SwitchToProductionMode(&_WalletPayments.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_WalletPayments *WalletPaymentsTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _WalletPayments.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_WalletPayments *WalletPaymentsSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.UpdateContractAddresses(&_WalletPayments.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_WalletPayments *WalletPaymentsTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _WalletPayments.Contract.UpdateContractAddresses(&_WalletPayments.TransactOpts, _contractNameHashes, _contractAddresses)
}

// WalletPaymentsAccountBtcWalletParamsSetIterator is returned from FilterAccountBtcWalletParamsSet and is used to iterate over the raw logs and unpacked data for AccountBtcWalletParamsSet events raised by the WalletPayments contract.
type WalletPaymentsAccountBtcWalletParamsSetIterator struct {
	Event *WalletPaymentsAccountBtcWalletParamsSet // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsAccountBtcWalletParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsAccountBtcWalletParamsSet)
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
		it.Event = new(WalletPaymentsAccountBtcWalletParamsSet)
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
func (it *WalletPaymentsAccountBtcWalletParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsAccountBtcWalletParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsAccountBtcWalletParamsSet represents a AccountBtcWalletParamsSet event raised by the WalletPayments contract.
type WalletPaymentsAccountBtcWalletParamsSet struct {
	WalletId     [32]byte
	AccountIndex uint32
	Params       IBtcParamsBtcWalletParams
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccountBtcWalletParamsSet is a free log retrieval operation binding the contract event 0xa78a36e246e3e524f6e695848ede9cfaa0be12a536ee80793b5a797897c278c5.
//
// Solidity: event AccountBtcWalletParamsSet(bytes32 indexed walletId, uint32 indexed accountIndex, (uint64,uint32,uint8,uint8) params)
func (_WalletPayments *WalletPaymentsFilterer) FilterAccountBtcWalletParamsSet(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32) (*WalletPaymentsAccountBtcWalletParamsSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "AccountBtcWalletParamsSet", walletIdRule, accountIndexRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsAccountBtcWalletParamsSetIterator{contract: _WalletPayments.contract, event: "AccountBtcWalletParamsSet", logs: logs, sub: sub}, nil
}

// WatchAccountBtcWalletParamsSet is a free log subscription operation binding the contract event 0xa78a36e246e3e524f6e695848ede9cfaa0be12a536ee80793b5a797897c278c5.
//
// Solidity: event AccountBtcWalletParamsSet(bytes32 indexed walletId, uint32 indexed accountIndex, (uint64,uint32,uint8,uint8) params)
func (_WalletPayments *WalletPaymentsFilterer) WatchAccountBtcWalletParamsSet(opts *bind.WatchOpts, sink chan<- *WalletPaymentsAccountBtcWalletParamsSet, walletId [][32]byte, accountIndex []uint32) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "AccountBtcWalletParamsSet", walletIdRule, accountIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsAccountBtcWalletParamsSet)
				if err := _WalletPayments.contract.UnpackLog(event, "AccountBtcWalletParamsSet", log); err != nil {
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

// ParseAccountBtcWalletParamsSet is a log parse operation binding the contract event 0xa78a36e246e3e524f6e695848ede9cfaa0be12a536ee80793b5a797897c278c5.
//
// Solidity: event AccountBtcWalletParamsSet(bytes32 indexed walletId, uint32 indexed accountIndex, (uint64,uint32,uint8,uint8) params)
func (_WalletPayments *WalletPaymentsFilterer) ParseAccountBtcWalletParamsSet(log types.Log) (*WalletPaymentsAccountBtcWalletParamsSet, error) {
	event := new(WalletPaymentsAccountBtcWalletParamsSet)
	if err := _WalletPayments.contract.UnpackLog(event, "AccountBtcWalletParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsAccountFeeScheduleClearedIterator is returned from FilterAccountFeeScheduleCleared and is used to iterate over the raw logs and unpacked data for AccountFeeScheduleCleared events raised by the WalletPayments contract.
type WalletPaymentsAccountFeeScheduleClearedIterator struct {
	Event *WalletPaymentsAccountFeeScheduleCleared // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsAccountFeeScheduleClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsAccountFeeScheduleCleared)
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
		it.Event = new(WalletPaymentsAccountFeeScheduleCleared)
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
func (it *WalletPaymentsAccountFeeScheduleClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsAccountFeeScheduleClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsAccountFeeScheduleCleared represents a AccountFeeScheduleCleared event raised by the WalletPayments contract.
type WalletPaymentsAccountFeeScheduleCleared struct {
	ProjectId      [32]byte
	SourceId       [32]byte
	AccountAddress string
	AccountHash    [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountFeeScheduleCleared is a free log retrieval operation binding the contract event 0x2426e08a74c1afcafde131339b2e2fa2289d879322564f7e49d16bfe6bb0fd7b.
//
// Solidity: event AccountFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash)
func (_WalletPayments *WalletPaymentsFilterer) FilterAccountFeeScheduleCleared(opts *bind.FilterOpts, projectId [][32]byte, sourceId [][32]byte, accountHash [][32]byte) (*WalletPaymentsAccountFeeScheduleClearedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	var accountHashRule []interface{}
	for _, accountHashItem := range accountHash {
		accountHashRule = append(accountHashRule, accountHashItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "AccountFeeScheduleCleared", projectIdRule, sourceIdRule, accountHashRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsAccountFeeScheduleClearedIterator{contract: _WalletPayments.contract, event: "AccountFeeScheduleCleared", logs: logs, sub: sub}, nil
}

// WatchAccountFeeScheduleCleared is a free log subscription operation binding the contract event 0x2426e08a74c1afcafde131339b2e2fa2289d879322564f7e49d16bfe6bb0fd7b.
//
// Solidity: event AccountFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash)
func (_WalletPayments *WalletPaymentsFilterer) WatchAccountFeeScheduleCleared(opts *bind.WatchOpts, sink chan<- *WalletPaymentsAccountFeeScheduleCleared, projectId [][32]byte, sourceId [][32]byte, accountHash [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	var accountHashRule []interface{}
	for _, accountHashItem := range accountHash {
		accountHashRule = append(accountHashRule, accountHashItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "AccountFeeScheduleCleared", projectIdRule, sourceIdRule, accountHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsAccountFeeScheduleCleared)
				if err := _WalletPayments.contract.UnpackLog(event, "AccountFeeScheduleCleared", log); err != nil {
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

// ParseAccountFeeScheduleCleared is a log parse operation binding the contract event 0x2426e08a74c1afcafde131339b2e2fa2289d879322564f7e49d16bfe6bb0fd7b.
//
// Solidity: event AccountFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash)
func (_WalletPayments *WalletPaymentsFilterer) ParseAccountFeeScheduleCleared(log types.Log) (*WalletPaymentsAccountFeeScheduleCleared, error) {
	event := new(WalletPaymentsAccountFeeScheduleCleared)
	if err := _WalletPayments.contract.UnpackLog(event, "AccountFeeScheduleCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsAccountFeeScheduleSetIterator is returned from FilterAccountFeeScheduleSet and is used to iterate over the raw logs and unpacked data for AccountFeeScheduleSet events raised by the WalletPayments contract.
type WalletPaymentsAccountFeeScheduleSetIterator struct {
	Event *WalletPaymentsAccountFeeScheduleSet // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsAccountFeeScheduleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsAccountFeeScheduleSet)
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
		it.Event = new(WalletPaymentsAccountFeeScheduleSet)
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
func (it *WalletPaymentsAccountFeeScheduleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsAccountFeeScheduleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsAccountFeeScheduleSet represents a AccountFeeScheduleSet event raised by the WalletPayments contract.
type WalletPaymentsAccountFeeScheduleSet struct {
	ProjectId      [32]byte
	SourceId       [32]byte
	AccountAddress string
	AccountHash    [32]byte
	Schedule       []IFeeSchedulesFeeSchedule
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAccountFeeScheduleSet is a free log retrieval operation binding the contract event 0x84562203b55c4f8e7837356a2b059908a1da339c2081ce5d934d3e0fb119b98f.
//
// Solidity: event AccountFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash, (int16,uint16)[] schedule)
func (_WalletPayments *WalletPaymentsFilterer) FilterAccountFeeScheduleSet(opts *bind.FilterOpts, projectId [][32]byte, sourceId [][32]byte, accountHash [][32]byte) (*WalletPaymentsAccountFeeScheduleSetIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	var accountHashRule []interface{}
	for _, accountHashItem := range accountHash {
		accountHashRule = append(accountHashRule, accountHashItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "AccountFeeScheduleSet", projectIdRule, sourceIdRule, accountHashRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsAccountFeeScheduleSetIterator{contract: _WalletPayments.contract, event: "AccountFeeScheduleSet", logs: logs, sub: sub}, nil
}

// WatchAccountFeeScheduleSet is a free log subscription operation binding the contract event 0x84562203b55c4f8e7837356a2b059908a1da339c2081ce5d934d3e0fb119b98f.
//
// Solidity: event AccountFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash, (int16,uint16)[] schedule)
func (_WalletPayments *WalletPaymentsFilterer) WatchAccountFeeScheduleSet(opts *bind.WatchOpts, sink chan<- *WalletPaymentsAccountFeeScheduleSet, projectId [][32]byte, sourceId [][32]byte, accountHash [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	var accountHashRule []interface{}
	for _, accountHashItem := range accountHash {
		accountHashRule = append(accountHashRule, accountHashItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "AccountFeeScheduleSet", projectIdRule, sourceIdRule, accountHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsAccountFeeScheduleSet)
				if err := _WalletPayments.contract.UnpackLog(event, "AccountFeeScheduleSet", log); err != nil {
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

// ParseAccountFeeScheduleSet is a log parse operation binding the contract event 0x84562203b55c4f8e7837356a2b059908a1da339c2081ce5d934d3e0fb119b98f.
//
// Solidity: event AccountFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash, (int16,uint16)[] schedule)
func (_WalletPayments *WalletPaymentsFilterer) ParseAccountFeeScheduleSet(log types.Log) (*WalletPaymentsAccountFeeScheduleSet, error) {
	event := new(WalletPaymentsAccountFeeScheduleSet)
	if err := _WalletPayments.contract.UnpackLog(event, "AccountFeeScheduleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsAccountProposersSetIterator is returned from FilterAccountProposersSet and is used to iterate over the raw logs and unpacked data for AccountProposersSet events raised by the WalletPayments contract.
type WalletPaymentsAccountProposersSetIterator struct {
	Event *WalletPaymentsAccountProposersSet // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsAccountProposersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsAccountProposersSet)
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
		it.Event = new(WalletPaymentsAccountProposersSet)
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
func (it *WalletPaymentsAccountProposersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsAccountProposersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsAccountProposersSet represents a AccountProposersSet event raised by the WalletPayments contract.
type WalletPaymentsAccountProposersSet struct {
	WalletId     [32]byte
	AccountIndex uint32
	Proposers    []common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAccountProposersSet is a free log retrieval operation binding the contract event 0x4681b3bfc634a38f7e29054f1c08f0f4d3ea4898159444ccea0a9eaab1d95a02.
//
// Solidity: event AccountProposersSet(bytes32 indexed walletId, uint32 indexed accountIndex, address[] proposers)
func (_WalletPayments *WalletPaymentsFilterer) FilterAccountProposersSet(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32) (*WalletPaymentsAccountProposersSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "AccountProposersSet", walletIdRule, accountIndexRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsAccountProposersSetIterator{contract: _WalletPayments.contract, event: "AccountProposersSet", logs: logs, sub: sub}, nil
}

// WatchAccountProposersSet is a free log subscription operation binding the contract event 0x4681b3bfc634a38f7e29054f1c08f0f4d3ea4898159444ccea0a9eaab1d95a02.
//
// Solidity: event AccountProposersSet(bytes32 indexed walletId, uint32 indexed accountIndex, address[] proposers)
func (_WalletPayments *WalletPaymentsFilterer) WatchAccountProposersSet(opts *bind.WatchOpts, sink chan<- *WalletPaymentsAccountProposersSet, walletId [][32]byte, accountIndex []uint32) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "AccountProposersSet", walletIdRule, accountIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsAccountProposersSet)
				if err := _WalletPayments.contract.UnpackLog(event, "AccountProposersSet", log); err != nil {
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

// ParseAccountProposersSet is a log parse operation binding the contract event 0x4681b3bfc634a38f7e29054f1c08f0f4d3ea4898159444ccea0a9eaab1d95a02.
//
// Solidity: event AccountProposersSet(bytes32 indexed walletId, uint32 indexed accountIndex, address[] proposers)
func (_WalletPayments *WalletPaymentsFilterer) ParseAccountProposersSet(log types.Log) (*WalletPaymentsAccountProposersSet, error) {
	event := new(WalletPaymentsAccountProposersSet)
	if err := _WalletPayments.contract.UnpackLog(event, "AccountProposersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsBtcAccountAddedIterator is returned from FilterBtcAccountAdded and is used to iterate over the raw logs and unpacked data for BtcAccountAdded events raised by the WalletPayments contract.
type WalletPaymentsBtcAccountAddedIterator struct {
	Event *WalletPaymentsBtcAccountAdded // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsBtcAccountAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsBtcAccountAdded)
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
		it.Event = new(WalletPaymentsBtcAccountAdded)
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
func (it *WalletPaymentsBtcAccountAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsBtcAccountAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsBtcAccountAdded represents a BtcAccountAdded event raised by the WalletPayments contract.
type WalletPaymentsBtcAccountAdded struct {
	WalletRegistry       common.Address
	WalletId             [32]byte
	SourceId             [32]byte
	AccountAddress       string
	AccountIndex         uint32
	AnchorCount          uint32
	AuthorizationAddress common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterBtcAccountAdded is a free log retrieval operation binding the contract event 0x7edb2f7accbc76d9e9e7d2a4349105afecbe1d249963dc9c7db912b9e5a0c44b.
//
// Solidity: event BtcAccountAdded(address indexed walletRegistry, bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount, address authorizationAddress)
func (_WalletPayments *WalletPaymentsFilterer) FilterBtcAccountAdded(opts *bind.FilterOpts, walletRegistry []common.Address, walletId [][32]byte) (*WalletPaymentsBtcAccountAddedIterator, error) {

	var walletRegistryRule []interface{}
	for _, walletRegistryItem := range walletRegistry {
		walletRegistryRule = append(walletRegistryRule, walletRegistryItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "BtcAccountAdded", walletRegistryRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsBtcAccountAddedIterator{contract: _WalletPayments.contract, event: "BtcAccountAdded", logs: logs, sub: sub}, nil
}

// WatchBtcAccountAdded is a free log subscription operation binding the contract event 0x7edb2f7accbc76d9e9e7d2a4349105afecbe1d249963dc9c7db912b9e5a0c44b.
//
// Solidity: event BtcAccountAdded(address indexed walletRegistry, bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount, address authorizationAddress)
func (_WalletPayments *WalletPaymentsFilterer) WatchBtcAccountAdded(opts *bind.WatchOpts, sink chan<- *WalletPaymentsBtcAccountAdded, walletRegistry []common.Address, walletId [][32]byte) (event.Subscription, error) {

	var walletRegistryRule []interface{}
	for _, walletRegistryItem := range walletRegistry {
		walletRegistryRule = append(walletRegistryRule, walletRegistryItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "BtcAccountAdded", walletRegistryRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsBtcAccountAdded)
				if err := _WalletPayments.contract.UnpackLog(event, "BtcAccountAdded", log); err != nil {
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

// ParseBtcAccountAdded is a log parse operation binding the contract event 0x7edb2f7accbc76d9e9e7d2a4349105afecbe1d249963dc9c7db912b9e5a0c44b.
//
// Solidity: event BtcAccountAdded(address indexed walletRegistry, bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount, address authorizationAddress)
func (_WalletPayments *WalletPaymentsFilterer) ParseBtcAccountAdded(log types.Log) (*WalletPaymentsBtcAccountAdded, error) {
	event := new(WalletPaymentsBtcAccountAdded)
	if err := _WalletPayments.contract.UnpackLog(event, "BtcAccountAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsBtcAnchorsAddedIterator is returned from FilterBtcAnchorsAdded and is used to iterate over the raw logs and unpacked data for BtcAnchorsAdded events raised by the WalletPayments contract.
type WalletPaymentsBtcAnchorsAddedIterator struct {
	Event *WalletPaymentsBtcAnchorsAdded // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsBtcAnchorsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsBtcAnchorsAdded)
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
		it.Event = new(WalletPaymentsBtcAnchorsAdded)
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
func (it *WalletPaymentsBtcAnchorsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsBtcAnchorsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsBtcAnchorsAdded represents a BtcAnchorsAdded event raised by the WalletPayments contract.
type WalletPaymentsBtcAnchorsAdded struct {
	WalletRegistry common.Address
	WalletId       [32]byte
	SourceId       [32]byte
	AccountAddress string
	AccountIndex   uint32
	AnchorCount    uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBtcAnchorsAdded is a free log retrieval operation binding the contract event 0x5315fabc859831b6483d8460fff21f81566822556da663245b85b2ec189c0577.
//
// Solidity: event BtcAnchorsAdded(address indexed walletRegistry, bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount)
func (_WalletPayments *WalletPaymentsFilterer) FilterBtcAnchorsAdded(opts *bind.FilterOpts, walletRegistry []common.Address, walletId [][32]byte) (*WalletPaymentsBtcAnchorsAddedIterator, error) {

	var walletRegistryRule []interface{}
	for _, walletRegistryItem := range walletRegistry {
		walletRegistryRule = append(walletRegistryRule, walletRegistryItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "BtcAnchorsAdded", walletRegistryRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsBtcAnchorsAddedIterator{contract: _WalletPayments.contract, event: "BtcAnchorsAdded", logs: logs, sub: sub}, nil
}

// WatchBtcAnchorsAdded is a free log subscription operation binding the contract event 0x5315fabc859831b6483d8460fff21f81566822556da663245b85b2ec189c0577.
//
// Solidity: event BtcAnchorsAdded(address indexed walletRegistry, bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount)
func (_WalletPayments *WalletPaymentsFilterer) WatchBtcAnchorsAdded(opts *bind.WatchOpts, sink chan<- *WalletPaymentsBtcAnchorsAdded, walletRegistry []common.Address, walletId [][32]byte) (event.Subscription, error) {

	var walletRegistryRule []interface{}
	for _, walletRegistryItem := range walletRegistry {
		walletRegistryRule = append(walletRegistryRule, walletRegistryItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "BtcAnchorsAdded", walletRegistryRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsBtcAnchorsAdded)
				if err := _WalletPayments.contract.UnpackLog(event, "BtcAnchorsAdded", log); err != nil {
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

// ParseBtcAnchorsAdded is a log parse operation binding the contract event 0x5315fabc859831b6483d8460fff21f81566822556da663245b85b2ec189c0577.
//
// Solidity: event BtcAnchorsAdded(address indexed walletRegistry, bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount)
func (_WalletPayments *WalletPaymentsFilterer) ParseBtcAnchorsAdded(log types.Log) (*WalletPaymentsBtcAnchorsAdded, error) {
	event := new(WalletPaymentsBtcAnchorsAdded)
	if err := _WalletPayments.contract.UnpackLog(event, "BtcAnchorsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsBtcAttemptSettledIterator is returned from FilterBtcAttemptSettled and is used to iterate over the raw logs and unpacked data for BtcAttemptSettled events raised by the WalletPayments contract.
type WalletPaymentsBtcAttemptSettledIterator struct {
	Event *WalletPaymentsBtcAttemptSettled // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsBtcAttemptSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsBtcAttemptSettled)
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
		it.Event = new(WalletPaymentsBtcAttemptSettled)
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
func (it *WalletPaymentsBtcAttemptSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsBtcAttemptSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsBtcAttemptSettled represents a BtcAttemptSettled event raised by the WalletPayments contract.
type WalletPaymentsBtcAttemptSettled struct {
	WalletId         [32]byte
	AccountIndex     uint32
	SequencePosition uint64
	Attempt          uint32
	AnchorIndex      uint32
	Nonce            uint64
	Txid             [32]byte
	NextAnchorTxid   [32]byte
	NextAnchorVout   uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBtcAttemptSettled is a free log retrieval operation binding the contract event 0x1bd7e0a3a9f760b6ce48058b6250ad1331b487fdf2549fc6c6525ca51b850e4e.
//
// Solidity: event BtcAttemptSettled(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, uint32 anchorIndex, uint64 nonce, bytes32 txid, bytes32 nextAnchorTxid, uint32 nextAnchorVout)
func (_WalletPayments *WalletPaymentsFilterer) FilterBtcAttemptSettled(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (*WalletPaymentsBtcAttemptSettledIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "BtcAttemptSettled", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsBtcAttemptSettledIterator{contract: _WalletPayments.contract, event: "BtcAttemptSettled", logs: logs, sub: sub}, nil
}

// WatchBtcAttemptSettled is a free log subscription operation binding the contract event 0x1bd7e0a3a9f760b6ce48058b6250ad1331b487fdf2549fc6c6525ca51b850e4e.
//
// Solidity: event BtcAttemptSettled(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, uint32 anchorIndex, uint64 nonce, bytes32 txid, bytes32 nextAnchorTxid, uint32 nextAnchorVout)
func (_WalletPayments *WalletPaymentsFilterer) WatchBtcAttemptSettled(opts *bind.WatchOpts, sink chan<- *WalletPaymentsBtcAttemptSettled, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "BtcAttemptSettled", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsBtcAttemptSettled)
				if err := _WalletPayments.contract.UnpackLog(event, "BtcAttemptSettled", log); err != nil {
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

// ParseBtcAttemptSettled is a log parse operation binding the contract event 0x1bd7e0a3a9f760b6ce48058b6250ad1331b487fdf2549fc6c6525ca51b850e4e.
//
// Solidity: event BtcAttemptSettled(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, uint32 anchorIndex, uint64 nonce, bytes32 txid, bytes32 nextAnchorTxid, uint32 nextAnchorVout)
func (_WalletPayments *WalletPaymentsFilterer) ParseBtcAttemptSettled(log types.Log) (*WalletPaymentsBtcAttemptSettled, error) {
	event := new(WalletPaymentsBtcAttemptSettled)
	if err := _WalletPayments.contract.UnpackLog(event, "BtcAttemptSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsBtcConsensusParamsSetIterator is returned from FilterBtcConsensusParamsSet and is used to iterate over the raw logs and unpacked data for BtcConsensusParamsSet events raised by the WalletPayments contract.
type WalletPaymentsBtcConsensusParamsSetIterator struct {
	Event *WalletPaymentsBtcConsensusParamsSet // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsBtcConsensusParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsBtcConsensusParamsSet)
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
		it.Event = new(WalletPaymentsBtcConsensusParamsSet)
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
func (it *WalletPaymentsBtcConsensusParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsBtcConsensusParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsBtcConsensusParamsSet represents a BtcConsensusParamsSet event raised by the WalletPayments contract.
type WalletPaymentsBtcConsensusParamsSet struct {
	SourceId [32]byte
	Params   IBtcParamsBtcConsensusParams
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBtcConsensusParamsSet is a free log retrieval operation binding the contract event 0x0c6f5bc4448213f91aa6d4a79354f537c61111de1eb30309644e69096ccaccb3.
//
// Solidity: event BtcConsensusParamsSet(bytes32 indexed sourceId, (uint16,uint64,uint64,uint32,uint8,uint8,uint8,uint8,uint16,uint32,uint32,uint32,uint32,uint32) params)
func (_WalletPayments *WalletPaymentsFilterer) FilterBtcConsensusParamsSet(opts *bind.FilterOpts, sourceId [][32]byte) (*WalletPaymentsBtcConsensusParamsSetIterator, error) {

	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "BtcConsensusParamsSet", sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsBtcConsensusParamsSetIterator{contract: _WalletPayments.contract, event: "BtcConsensusParamsSet", logs: logs, sub: sub}, nil
}

// WatchBtcConsensusParamsSet is a free log subscription operation binding the contract event 0x0c6f5bc4448213f91aa6d4a79354f537c61111de1eb30309644e69096ccaccb3.
//
// Solidity: event BtcConsensusParamsSet(bytes32 indexed sourceId, (uint16,uint64,uint64,uint32,uint8,uint8,uint8,uint8,uint16,uint32,uint32,uint32,uint32,uint32) params)
func (_WalletPayments *WalletPaymentsFilterer) WatchBtcConsensusParamsSet(opts *bind.WatchOpts, sink chan<- *WalletPaymentsBtcConsensusParamsSet, sourceId [][32]byte) (event.Subscription, error) {

	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "BtcConsensusParamsSet", sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsBtcConsensusParamsSet)
				if err := _WalletPayments.contract.UnpackLog(event, "BtcConsensusParamsSet", log); err != nil {
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

// ParseBtcConsensusParamsSet is a log parse operation binding the contract event 0x0c6f5bc4448213f91aa6d4a79354f537c61111de1eb30309644e69096ccaccb3.
//
// Solidity: event BtcConsensusParamsSet(bytes32 indexed sourceId, (uint16,uint64,uint64,uint32,uint8,uint8,uint8,uint8,uint16,uint32,uint32,uint32,uint32,uint32) params)
func (_WalletPayments *WalletPaymentsFilterer) ParseBtcConsensusParamsSet(log types.Log) (*WalletPaymentsBtcConsensusParamsSet, error) {
	event := new(WalletPaymentsBtcConsensusParamsSet)
	if err := _WalletPayments.contract.UnpackLog(event, "BtcConsensusParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsBtcEscrowEmittedIterator is returned from FilterBtcEscrowEmitted and is used to iterate over the raw logs and unpacked data for BtcEscrowEmitted events raised by the WalletPayments contract.
type WalletPaymentsBtcEscrowEmittedIterator struct {
	Event *WalletPaymentsBtcEscrowEmitted // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsBtcEscrowEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsBtcEscrowEmitted)
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
		it.Event = new(WalletPaymentsBtcEscrowEmitted)
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
func (it *WalletPaymentsBtcEscrowEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsBtcEscrowEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsBtcEscrowEmitted represents a BtcEscrowEmitted event raised by the WalletPayments contract.
type WalletPaymentsBtcEscrowEmitted struct {
	WalletId         [32]byte
	AccountIndex     uint32
	PaymentId        uint64
	SequencePosition uint64
	Kind             uint8
	PreimageHash     [32]byte
	Amount           uint64
	ExpiresAt        uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBtcEscrowEmitted is a free log retrieval operation binding the contract event 0xcdeaf40af840e3981d311b0486b6350e459197d0bfec07f602eb8dfbd5fe1b3d.
//
// Solidity: event BtcEscrowEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed paymentId, uint64 sequencePosition, uint8 kind, bytes32 preimageHash, uint64 amount, uint64 expiresAt)
func (_WalletPayments *WalletPaymentsFilterer) FilterBtcEscrowEmitted(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32, paymentId []uint64) (*WalletPaymentsBtcEscrowEmittedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "BtcEscrowEmitted", walletIdRule, accountIndexRule, paymentIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsBtcEscrowEmittedIterator{contract: _WalletPayments.contract, event: "BtcEscrowEmitted", logs: logs, sub: sub}, nil
}

// WatchBtcEscrowEmitted is a free log subscription operation binding the contract event 0xcdeaf40af840e3981d311b0486b6350e459197d0bfec07f602eb8dfbd5fe1b3d.
//
// Solidity: event BtcEscrowEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed paymentId, uint64 sequencePosition, uint8 kind, bytes32 preimageHash, uint64 amount, uint64 expiresAt)
func (_WalletPayments *WalletPaymentsFilterer) WatchBtcEscrowEmitted(opts *bind.WatchOpts, sink chan<- *WalletPaymentsBtcEscrowEmitted, walletId [][32]byte, accountIndex []uint32, paymentId []uint64) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "BtcEscrowEmitted", walletIdRule, accountIndexRule, paymentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsBtcEscrowEmitted)
				if err := _WalletPayments.contract.UnpackLog(event, "BtcEscrowEmitted", log); err != nil {
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

// ParseBtcEscrowEmitted is a log parse operation binding the contract event 0xcdeaf40af840e3981d311b0486b6350e459197d0bfec07f602eb8dfbd5fe1b3d.
//
// Solidity: event BtcEscrowEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed paymentId, uint64 sequencePosition, uint8 kind, bytes32 preimageHash, uint64 amount, uint64 expiresAt)
func (_WalletPayments *WalletPaymentsFilterer) ParseBtcEscrowEmitted(log types.Log) (*WalletPaymentsBtcEscrowEmitted, error) {
	event := new(WalletPaymentsBtcEscrowEmitted)
	if err := _WalletPayments.contract.UnpackLog(event, "BtcEscrowEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsBtcEscrowProfileSetIterator is returned from FilterBtcEscrowProfileSet and is used to iterate over the raw logs and unpacked data for BtcEscrowProfileSet events raised by the WalletPayments contract.
type WalletPaymentsBtcEscrowProfileSetIterator struct {
	Event *WalletPaymentsBtcEscrowProfileSet // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsBtcEscrowProfileSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsBtcEscrowProfileSet)
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
		it.Event = new(WalletPaymentsBtcEscrowProfileSet)
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
func (it *WalletPaymentsBtcEscrowProfileSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsBtcEscrowProfileSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsBtcEscrowProfileSet represents a BtcEscrowProfileSet event raised by the WalletPayments contract.
type WalletPaymentsBtcEscrowProfileSet struct {
	WalletId           [32]byte
	AccountIndex       uint32
	CounterpartyPubKey []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBtcEscrowProfileSet is a free log retrieval operation binding the contract event 0xdcbde836bb21153a86c06d7580d72441dd03e2cb0530f06759f32633ec1eb650.
//
// Solidity: event BtcEscrowProfileSet(bytes32 indexed walletId, uint32 indexed accountIndex, bytes counterpartyPubKey)
func (_WalletPayments *WalletPaymentsFilterer) FilterBtcEscrowProfileSet(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32) (*WalletPaymentsBtcEscrowProfileSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "BtcEscrowProfileSet", walletIdRule, accountIndexRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsBtcEscrowProfileSetIterator{contract: _WalletPayments.contract, event: "BtcEscrowProfileSet", logs: logs, sub: sub}, nil
}

// WatchBtcEscrowProfileSet is a free log subscription operation binding the contract event 0xdcbde836bb21153a86c06d7580d72441dd03e2cb0530f06759f32633ec1eb650.
//
// Solidity: event BtcEscrowProfileSet(bytes32 indexed walletId, uint32 indexed accountIndex, bytes counterpartyPubKey)
func (_WalletPayments *WalletPaymentsFilterer) WatchBtcEscrowProfileSet(opts *bind.WatchOpts, sink chan<- *WalletPaymentsBtcEscrowProfileSet, walletId [][32]byte, accountIndex []uint32) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "BtcEscrowProfileSet", walletIdRule, accountIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsBtcEscrowProfileSet)
				if err := _WalletPayments.contract.UnpackLog(event, "BtcEscrowProfileSet", log); err != nil {
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

// ParseBtcEscrowProfileSet is a log parse operation binding the contract event 0xdcbde836bb21153a86c06d7580d72441dd03e2cb0530f06759f32633ec1eb650.
//
// Solidity: event BtcEscrowProfileSet(bytes32 indexed walletId, uint32 indexed accountIndex, bytes counterpartyPubKey)
func (_WalletPayments *WalletPaymentsFilterer) ParseBtcEscrowProfileSet(log types.Log) (*WalletPaymentsBtcEscrowProfileSet, error) {
	event := new(WalletPaymentsBtcEscrowProfileSet)
	if err := _WalletPayments.contract.UnpackLog(event, "BtcEscrowProfileSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsBtcEscrowReclaimEmittedIterator is returned from FilterBtcEscrowReclaimEmitted and is used to iterate over the raw logs and unpacked data for BtcEscrowReclaimEmitted events raised by the WalletPayments contract.
type WalletPaymentsBtcEscrowReclaimEmittedIterator struct {
	Event *WalletPaymentsBtcEscrowReclaimEmitted // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsBtcEscrowReclaimEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsBtcEscrowReclaimEmitted)
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
		it.Event = new(WalletPaymentsBtcEscrowReclaimEmitted)
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
func (it *WalletPaymentsBtcEscrowReclaimEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsBtcEscrowReclaimEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsBtcEscrowReclaimEmitted represents a BtcEscrowReclaimEmitted event raised by the WalletPayments contract.
type WalletPaymentsBtcEscrowReclaimEmitted struct {
	WalletId        [32]byte
	AccountIndex    uint32
	PaymentId       uint64
	CreatePaymentId uint64
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBtcEscrowReclaimEmitted is a free log retrieval operation binding the contract event 0x9b14da8ae08717a7867c015a92cca74e5ea34f6f818cec109ee3369248aec68b.
//
// Solidity: event BtcEscrowReclaimEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed paymentId, uint64 createPaymentId)
func (_WalletPayments *WalletPaymentsFilterer) FilterBtcEscrowReclaimEmitted(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32, paymentId []uint64) (*WalletPaymentsBtcEscrowReclaimEmittedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "BtcEscrowReclaimEmitted", walletIdRule, accountIndexRule, paymentIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsBtcEscrowReclaimEmittedIterator{contract: _WalletPayments.contract, event: "BtcEscrowReclaimEmitted", logs: logs, sub: sub}, nil
}

// WatchBtcEscrowReclaimEmitted is a free log subscription operation binding the contract event 0x9b14da8ae08717a7867c015a92cca74e5ea34f6f818cec109ee3369248aec68b.
//
// Solidity: event BtcEscrowReclaimEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed paymentId, uint64 createPaymentId)
func (_WalletPayments *WalletPaymentsFilterer) WatchBtcEscrowReclaimEmitted(opts *bind.WatchOpts, sink chan<- *WalletPaymentsBtcEscrowReclaimEmitted, walletId [][32]byte, accountIndex []uint32, paymentId []uint64) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "BtcEscrowReclaimEmitted", walletIdRule, accountIndexRule, paymentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsBtcEscrowReclaimEmitted)
				if err := _WalletPayments.contract.UnpackLog(event, "BtcEscrowReclaimEmitted", log); err != nil {
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

// ParseBtcEscrowReclaimEmitted is a log parse operation binding the contract event 0x9b14da8ae08717a7867c015a92cca74e5ea34f6f818cec109ee3369248aec68b.
//
// Solidity: event BtcEscrowReclaimEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed paymentId, uint64 createPaymentId)
func (_WalletPayments *WalletPaymentsFilterer) ParseBtcEscrowReclaimEmitted(log types.Log) (*WalletPaymentsBtcEscrowReclaimEmitted, error) {
	event := new(WalletPaymentsBtcEscrowReclaimEmitted)
	if err := _WalletPayments.contract.UnpackLog(event, "BtcEscrowReclaimEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsCspFeeForwardedIterator is returned from FilterCspFeeForwarded and is used to iterate over the raw logs and unpacked data for CspFeeForwarded events raised by the WalletPayments contract.
type WalletPaymentsCspFeeForwardedIterator struct {
	Event *WalletPaymentsCspFeeForwarded // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsCspFeeForwardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsCspFeeForwarded)
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
		it.Event = new(WalletPaymentsCspFeeForwarded)
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
func (it *WalletPaymentsCspFeeForwardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsCspFeeForwardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsCspFeeForwarded represents a CspFeeForwarded event raised by the WalletPayments contract.
type WalletPaymentsCspFeeForwarded struct {
	WalletId         [32]byte
	AccountIndex     uint32
	OpCommand        [32]byte
	PaymentId        uint64
	ClaimBackAddress common.Address
	Value            *big.Int
	RewardEpochId    *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterCspFeeForwarded is a free log retrieval operation binding the contract event 0x23b023e616d9f2fc3336d1761d5048e13dc27ed2c1e7b23ebdad04aa57c7e2ea.
//
// Solidity: event CspFeeForwarded(bytes32 indexed walletId, uint32 indexed accountIndex, bytes32 indexed opCommand, uint64 paymentId, address claimBackAddress, uint256 value, uint24 rewardEpochId)
func (_WalletPayments *WalletPaymentsFilterer) FilterCspFeeForwarded(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32, opCommand [][32]byte) (*WalletPaymentsCspFeeForwardedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var opCommandRule []interface{}
	for _, opCommandItem := range opCommand {
		opCommandRule = append(opCommandRule, opCommandItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "CspFeeForwarded", walletIdRule, accountIndexRule, opCommandRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsCspFeeForwardedIterator{contract: _WalletPayments.contract, event: "CspFeeForwarded", logs: logs, sub: sub}, nil
}

// WatchCspFeeForwarded is a free log subscription operation binding the contract event 0x23b023e616d9f2fc3336d1761d5048e13dc27ed2c1e7b23ebdad04aa57c7e2ea.
//
// Solidity: event CspFeeForwarded(bytes32 indexed walletId, uint32 indexed accountIndex, bytes32 indexed opCommand, uint64 paymentId, address claimBackAddress, uint256 value, uint24 rewardEpochId)
func (_WalletPayments *WalletPaymentsFilterer) WatchCspFeeForwarded(opts *bind.WatchOpts, sink chan<- *WalletPaymentsCspFeeForwarded, walletId [][32]byte, accountIndex []uint32, opCommand [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var opCommandRule []interface{}
	for _, opCommandItem := range opCommand {
		opCommandRule = append(opCommandRule, opCommandItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "CspFeeForwarded", walletIdRule, accountIndexRule, opCommandRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsCspFeeForwarded)
				if err := _WalletPayments.contract.UnpackLog(event, "CspFeeForwarded", log); err != nil {
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

// ParseCspFeeForwarded is a log parse operation binding the contract event 0x23b023e616d9f2fc3336d1761d5048e13dc27ed2c1e7b23ebdad04aa57c7e2ea.
//
// Solidity: event CspFeeForwarded(bytes32 indexed walletId, uint32 indexed accountIndex, bytes32 indexed opCommand, uint64 paymentId, address claimBackAddress, uint256 value, uint24 rewardEpochId)
func (_WalletPayments *WalletPaymentsFilterer) ParseCspFeeForwarded(log types.Log) (*WalletPaymentsCspFeeForwarded, error) {
	event := new(WalletPaymentsCspFeeForwarded)
	if err := _WalletPayments.contract.UnpackLog(event, "CspFeeForwarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsCspSourceSettingsSetIterator is returned from FilterCspSourceSettingsSet and is used to iterate over the raw logs and unpacked data for CspSourceSettingsSet events raised by the WalletPayments contract.
type WalletPaymentsCspSourceSettingsSetIterator struct {
	Event *WalletPaymentsCspSourceSettingsSet // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsCspSourceSettingsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsCspSourceSettingsSet)
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
		it.Event = new(WalletPaymentsCspSourceSettingsSet)
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
func (it *WalletPaymentsCspSourceSettingsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsCspSourceSettingsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsCspSourceSettingsSet represents a CspSourceSettingsSet event raised by the WalletPayments contract.
type WalletPaymentsCspSourceSettingsSet struct {
	SourceId [32]byte
	Settings ICspProposalsCspSourceSettings
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCspSourceSettingsSet is a free log retrieval operation binding the contract event 0x743f87af3b1953facf06417af2e8db7e222a428a0466d074c6360b17ae5f058f.
//
// Solidity: event CspSourceSettingsSet(bytes32 indexed sourceId, (uint64,uint16,uint16,uint8,uint128) settings)
func (_WalletPayments *WalletPaymentsFilterer) FilterCspSourceSettingsSet(opts *bind.FilterOpts, sourceId [][32]byte) (*WalletPaymentsCspSourceSettingsSetIterator, error) {

	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "CspSourceSettingsSet", sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsCspSourceSettingsSetIterator{contract: _WalletPayments.contract, event: "CspSourceSettingsSet", logs: logs, sub: sub}, nil
}

// WatchCspSourceSettingsSet is a free log subscription operation binding the contract event 0x743f87af3b1953facf06417af2e8db7e222a428a0466d074c6360b17ae5f058f.
//
// Solidity: event CspSourceSettingsSet(bytes32 indexed sourceId, (uint64,uint16,uint16,uint8,uint128) settings)
func (_WalletPayments *WalletPaymentsFilterer) WatchCspSourceSettingsSet(opts *bind.WatchOpts, sink chan<- *WalletPaymentsCspSourceSettingsSet, sourceId [][32]byte) (event.Subscription, error) {

	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "CspSourceSettingsSet", sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsCspSourceSettingsSet)
				if err := _WalletPayments.contract.UnpackLog(event, "CspSourceSettingsSet", log); err != nil {
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

// ParseCspSourceSettingsSet is a log parse operation binding the contract event 0x743f87af3b1953facf06417af2e8db7e222a428a0466d074c6360b17ae5f058f.
//
// Solidity: event CspSourceSettingsSet(bytes32 indexed sourceId, (uint64,uint16,uint16,uint8,uint128) settings)
func (_WalletPayments *WalletPaymentsFilterer) ParseCspSourceSettingsSet(log types.Log) (*WalletPaymentsCspSourceSettingsSet, error) {
	event := new(WalletPaymentsCspSourceSettingsSet)
	if err := _WalletPayments.contract.UnpackLog(event, "CspSourceSettingsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsCustodianInstructionIssuedIterator is returned from FilterCustodianInstructionIssued and is used to iterate over the raw logs and unpacked data for CustodianInstructionIssued events raised by the WalletPayments contract.
type WalletPaymentsCustodianInstructionIssuedIterator struct {
	Event *WalletPaymentsCustodianInstructionIssued // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsCustodianInstructionIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsCustodianInstructionIssued)
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
		it.Event = new(WalletPaymentsCustodianInstructionIssued)
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
func (it *WalletPaymentsCustodianInstructionIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsCustodianInstructionIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsCustodianInstructionIssued represents a CustodianInstructionIssued event raised by the WalletPayments contract.
type WalletPaymentsCustodianInstructionIssued struct {
	WalletRegistry common.Address
	WalletId       [32]byte
	InstructionId  [32]byte
	OpType         [32]byte
	OpCommand      [32]byte
	Message        []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCustodianInstructionIssued is a free log retrieval operation binding the contract event 0x359b53606cb8cc4b2ad91fa298aed01e6402d3ccc6b1616f966b267b7ef6d708.
//
// Solidity: event CustodianInstructionIssued(address indexed walletRegistry, bytes32 indexed walletId, bytes32 indexed instructionId, bytes32 opType, bytes32 opCommand, bytes message)
func (_WalletPayments *WalletPaymentsFilterer) FilterCustodianInstructionIssued(opts *bind.FilterOpts, walletRegistry []common.Address, walletId [][32]byte, instructionId [][32]byte) (*WalletPaymentsCustodianInstructionIssuedIterator, error) {

	var walletRegistryRule []interface{}
	for _, walletRegistryItem := range walletRegistry {
		walletRegistryRule = append(walletRegistryRule, walletRegistryItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "CustodianInstructionIssued", walletRegistryRule, walletIdRule, instructionIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsCustodianInstructionIssuedIterator{contract: _WalletPayments.contract, event: "CustodianInstructionIssued", logs: logs, sub: sub}, nil
}

// WatchCustodianInstructionIssued is a free log subscription operation binding the contract event 0x359b53606cb8cc4b2ad91fa298aed01e6402d3ccc6b1616f966b267b7ef6d708.
//
// Solidity: event CustodianInstructionIssued(address indexed walletRegistry, bytes32 indexed walletId, bytes32 indexed instructionId, bytes32 opType, bytes32 opCommand, bytes message)
func (_WalletPayments *WalletPaymentsFilterer) WatchCustodianInstructionIssued(opts *bind.WatchOpts, sink chan<- *WalletPaymentsCustodianInstructionIssued, walletRegistry []common.Address, walletId [][32]byte, instructionId [][32]byte) (event.Subscription, error) {

	var walletRegistryRule []interface{}
	for _, walletRegistryItem := range walletRegistry {
		walletRegistryRule = append(walletRegistryRule, walletRegistryItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "CustodianInstructionIssued", walletRegistryRule, walletIdRule, instructionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsCustodianInstructionIssued)
				if err := _WalletPayments.contract.UnpackLog(event, "CustodianInstructionIssued", log); err != nil {
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

// ParseCustodianInstructionIssued is a log parse operation binding the contract event 0x359b53606cb8cc4b2ad91fa298aed01e6402d3ccc6b1616f966b267b7ef6d708.
//
// Solidity: event CustodianInstructionIssued(address indexed walletRegistry, bytes32 indexed walletId, bytes32 indexed instructionId, bytes32 opType, bytes32 opCommand, bytes message)
func (_WalletPayments *WalletPaymentsFilterer) ParseCustodianInstructionIssued(log types.Log) (*WalletPaymentsCustodianInstructionIssued, error) {
	event := new(WalletPaymentsCustodianInstructionIssued)
	if err := _WalletPayments.contract.UnpackLog(event, "CustodianInstructionIssued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the WalletPayments contract.
type WalletPaymentsDiamondCutIterator struct {
	Event *WalletPaymentsDiamondCut // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsDiamondCut)
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
		it.Event = new(WalletPaymentsDiamondCut)
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
func (it *WalletPaymentsDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsDiamondCut represents a DiamondCut event raised by the WalletPayments contract.
type WalletPaymentsDiamondCut struct {
	DiamondCut []IDiamondFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_WalletPayments *WalletPaymentsFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*WalletPaymentsDiamondCutIterator, error) {

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsDiamondCutIterator{contract: _WalletPayments.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_WalletPayments *WalletPaymentsFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *WalletPaymentsDiamondCut) (event.Subscription, error) {

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsDiamondCut)
				if err := _WalletPayments.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_WalletPayments *WalletPaymentsFilterer) ParseDiamondCut(log types.Log) (*WalletPaymentsDiamondCut, error) {
	event := new(WalletPaymentsDiamondCut)
	if err := _WalletPayments.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsEligibleAdvancedIterator is returned from FilterEligibleAdvanced and is used to iterate over the raw logs and unpacked data for EligibleAdvanced events raised by the WalletPayments contract.
type WalletPaymentsEligibleAdvancedIterator struct {
	Event *WalletPaymentsEligibleAdvanced // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsEligibleAdvancedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsEligibleAdvanced)
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
		it.Event = new(WalletPaymentsEligibleAdvanced)
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
func (it *WalletPaymentsEligibleAdvancedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsEligibleAdvancedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsEligibleAdvanced represents a EligibleAdvanced event raised by the WalletPayments contract.
type WalletPaymentsEligibleAdvanced struct {
	WalletId         [32]byte
	AccountIndex     uint32
	SequencePosition uint64
	Attempt          uint32
	Generation       uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterEligibleAdvanced is a free log retrieval operation binding the contract event 0x963dd061e388c9d694ffb537491dbba5ff9da375c2884ab3311afc1e65b9846e.
//
// Solidity: event EligibleAdvanced(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 sequencePosition, uint32 attempt, uint64 generation)
func (_WalletPayments *WalletPaymentsFilterer) FilterEligibleAdvanced(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32) (*WalletPaymentsEligibleAdvancedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "EligibleAdvanced", walletIdRule, accountIndexRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsEligibleAdvancedIterator{contract: _WalletPayments.contract, event: "EligibleAdvanced", logs: logs, sub: sub}, nil
}

// WatchEligibleAdvanced is a free log subscription operation binding the contract event 0x963dd061e388c9d694ffb537491dbba5ff9da375c2884ab3311afc1e65b9846e.
//
// Solidity: event EligibleAdvanced(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 sequencePosition, uint32 attempt, uint64 generation)
func (_WalletPayments *WalletPaymentsFilterer) WatchEligibleAdvanced(opts *bind.WatchOpts, sink chan<- *WalletPaymentsEligibleAdvanced, walletId [][32]byte, accountIndex []uint32) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "EligibleAdvanced", walletIdRule, accountIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsEligibleAdvanced)
				if err := _WalletPayments.contract.UnpackLog(event, "EligibleAdvanced", log); err != nil {
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

// ParseEligibleAdvanced is a log parse operation binding the contract event 0x963dd061e388c9d694ffb537491dbba5ff9da375c2884ab3311afc1e65b9846e.
//
// Solidity: event EligibleAdvanced(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 sequencePosition, uint32 attempt, uint64 generation)
func (_WalletPayments *WalletPaymentsFilterer) ParseEligibleAdvanced(log types.Log) (*WalletPaymentsEligibleAdvanced, error) {
	event := new(WalletPaymentsEligibleAdvanced)
	if err := _WalletPayments.contract.UnpackLog(event, "EligibleAdvanced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsFeeScheduleConfigsClearedIterator is returned from FilterFeeScheduleConfigsCleared and is used to iterate over the raw logs and unpacked data for FeeScheduleConfigsCleared events raised by the WalletPayments contract.
type WalletPaymentsFeeScheduleConfigsClearedIterator struct {
	Event *WalletPaymentsFeeScheduleConfigsCleared // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsFeeScheduleConfigsClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsFeeScheduleConfigsCleared)
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
		it.Event = new(WalletPaymentsFeeScheduleConfigsCleared)
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
func (it *WalletPaymentsFeeScheduleConfigsClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsFeeScheduleConfigsClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsFeeScheduleConfigsCleared represents a FeeScheduleConfigsCleared event raised by the WalletPayments contract.
type WalletPaymentsFeeScheduleConfigsCleared struct {
	SourceIds [][32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeScheduleConfigsCleared is a free log retrieval operation binding the contract event 0xa65a10cbf89a889142f7004143171a2381448115ee506c5a8a1c4c8444f76016.
//
// Solidity: event FeeScheduleConfigsCleared(bytes32[] sourceIds)
func (_WalletPayments *WalletPaymentsFilterer) FilterFeeScheduleConfigsCleared(opts *bind.FilterOpts) (*WalletPaymentsFeeScheduleConfigsClearedIterator, error) {

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "FeeScheduleConfigsCleared")
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsFeeScheduleConfigsClearedIterator{contract: _WalletPayments.contract, event: "FeeScheduleConfigsCleared", logs: logs, sub: sub}, nil
}

// WatchFeeScheduleConfigsCleared is a free log subscription operation binding the contract event 0xa65a10cbf89a889142f7004143171a2381448115ee506c5a8a1c4c8444f76016.
//
// Solidity: event FeeScheduleConfigsCleared(bytes32[] sourceIds)
func (_WalletPayments *WalletPaymentsFilterer) WatchFeeScheduleConfigsCleared(opts *bind.WatchOpts, sink chan<- *WalletPaymentsFeeScheduleConfigsCleared) (event.Subscription, error) {

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "FeeScheduleConfigsCleared")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsFeeScheduleConfigsCleared)
				if err := _WalletPayments.contract.UnpackLog(event, "FeeScheduleConfigsCleared", log); err != nil {
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

// ParseFeeScheduleConfigsCleared is a log parse operation binding the contract event 0xa65a10cbf89a889142f7004143171a2381448115ee506c5a8a1c4c8444f76016.
//
// Solidity: event FeeScheduleConfigsCleared(bytes32[] sourceIds)
func (_WalletPayments *WalletPaymentsFilterer) ParseFeeScheduleConfigsCleared(log types.Log) (*WalletPaymentsFeeScheduleConfigsCleared, error) {
	event := new(WalletPaymentsFeeScheduleConfigsCleared)
	if err := _WalletPayments.contract.UnpackLog(event, "FeeScheduleConfigsCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsFeeScheduleConfigsSetIterator is returned from FilterFeeScheduleConfigsSet and is used to iterate over the raw logs and unpacked data for FeeScheduleConfigsSet events raised by the WalletPayments contract.
type WalletPaymentsFeeScheduleConfigsSetIterator struct {
	Event *WalletPaymentsFeeScheduleConfigsSet // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsFeeScheduleConfigsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsFeeScheduleConfigsSet)
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
		it.Event = new(WalletPaymentsFeeScheduleConfigsSet)
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
func (it *WalletPaymentsFeeScheduleConfigsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsFeeScheduleConfigsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsFeeScheduleConfigsSet represents a FeeScheduleConfigsSet event raised by the WalletPayments contract.
type WalletPaymentsFeeScheduleConfigsSet struct {
	Configs []IFeeSchedulesFeeScheduleConfigInput
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeScheduleConfigsSet is a free log retrieval operation binding the contract event 0x6a0718b26b38a09770c97425ea4ecdb442e8d8ab07960c3e00fd32c6397646c4.
//
// Solidity: event FeeScheduleConfigsSet((uint16,uint8,bytes32)[] configs)
func (_WalletPayments *WalletPaymentsFilterer) FilterFeeScheduleConfigsSet(opts *bind.FilterOpts) (*WalletPaymentsFeeScheduleConfigsSetIterator, error) {

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "FeeScheduleConfigsSet")
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsFeeScheduleConfigsSetIterator{contract: _WalletPayments.contract, event: "FeeScheduleConfigsSet", logs: logs, sub: sub}, nil
}

// WatchFeeScheduleConfigsSet is a free log subscription operation binding the contract event 0x6a0718b26b38a09770c97425ea4ecdb442e8d8ab07960c3e00fd32c6397646c4.
//
// Solidity: event FeeScheduleConfigsSet((uint16,uint8,bytes32)[] configs)
func (_WalletPayments *WalletPaymentsFilterer) WatchFeeScheduleConfigsSet(opts *bind.WatchOpts, sink chan<- *WalletPaymentsFeeScheduleConfigsSet) (event.Subscription, error) {

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "FeeScheduleConfigsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsFeeScheduleConfigsSet)
				if err := _WalletPayments.contract.UnpackLog(event, "FeeScheduleConfigsSet", log); err != nil {
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

// ParseFeeScheduleConfigsSet is a log parse operation binding the contract event 0x6a0718b26b38a09770c97425ea4ecdb442e8d8ab07960c3e00fd32c6397646c4.
//
// Solidity: event FeeScheduleConfigsSet((uint16,uint8,bytes32)[] configs)
func (_WalletPayments *WalletPaymentsFilterer) ParseFeeScheduleConfigsSet(log types.Log) (*WalletPaymentsFeeScheduleConfigsSet, error) {
	event := new(WalletPaymentsFeeScheduleConfigsSet)
	if err := _WalletPayments.contract.UnpackLog(event, "FeeScheduleConfigsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the WalletPayments contract.
type WalletPaymentsGovernanceCallTimelockedIterator struct {
	Event *WalletPaymentsGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsGovernanceCallTimelocked)
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
		it.Event = new(WalletPaymentsGovernanceCallTimelocked)
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
func (it *WalletPaymentsGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the WalletPayments contract.
type WalletPaymentsGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_WalletPayments *WalletPaymentsFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*WalletPaymentsGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsGovernanceCallTimelockedIterator{contract: _WalletPayments.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_WalletPayments *WalletPaymentsFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *WalletPaymentsGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsGovernanceCallTimelocked)
				if err := _WalletPayments.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_WalletPayments *WalletPaymentsFilterer) ParseGovernanceCallTimelocked(log types.Log) (*WalletPaymentsGovernanceCallTimelocked, error) {
	event := new(WalletPaymentsGovernanceCallTimelocked)
	if err := _WalletPayments.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the WalletPayments contract.
type WalletPaymentsGovernanceInitialisedIterator struct {
	Event *WalletPaymentsGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsGovernanceInitialised)
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
		it.Event = new(WalletPaymentsGovernanceInitialised)
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
func (it *WalletPaymentsGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsGovernanceInitialised represents a GovernanceInitialised event raised by the WalletPayments contract.
type WalletPaymentsGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_WalletPayments *WalletPaymentsFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*WalletPaymentsGovernanceInitialisedIterator, error) {

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsGovernanceInitialisedIterator{contract: _WalletPayments.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_WalletPayments *WalletPaymentsFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *WalletPaymentsGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsGovernanceInitialised)
				if err := _WalletPayments.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_WalletPayments *WalletPaymentsFilterer) ParseGovernanceInitialised(log types.Log) (*WalletPaymentsGovernanceInitialised, error) {
	event := new(WalletPaymentsGovernanceInitialised)
	if err := _WalletPayments.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the WalletPayments contract.
type WalletPaymentsGovernedProductionModeEnteredIterator struct {
	Event *WalletPaymentsGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsGovernedProductionModeEntered)
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
		it.Event = new(WalletPaymentsGovernedProductionModeEntered)
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
func (it *WalletPaymentsGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the WalletPayments contract.
type WalletPaymentsGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_WalletPayments *WalletPaymentsFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*WalletPaymentsGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsGovernedProductionModeEnteredIterator{contract: _WalletPayments.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_WalletPayments *WalletPaymentsFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *WalletPaymentsGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsGovernedProductionModeEntered)
				if err := _WalletPayments.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_WalletPayments *WalletPaymentsFilterer) ParseGovernedProductionModeEntered(log types.Log) (*WalletPaymentsGovernedProductionModeEntered, error) {
	event := new(WalletPaymentsGovernedProductionModeEntered)
	if err := _WalletPayments.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsInstructionEmittedIterator is returned from FilterInstructionEmitted and is used to iterate over the raw logs and unpacked data for InstructionEmitted events raised by the WalletPayments contract.
type WalletPaymentsInstructionEmittedIterator struct {
	Event *WalletPaymentsInstructionEmitted // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsInstructionEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsInstructionEmitted)
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
		it.Event = new(WalletPaymentsInstructionEmitted)
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
func (it *WalletPaymentsInstructionEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsInstructionEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsInstructionEmitted represents a InstructionEmitted event raised by the WalletPayments contract.
type WalletPaymentsInstructionEmitted struct {
	WalletId         [32]byte
	AccountIndex     uint32
	SequencePosition uint64
	Kind             uint8
	PaymentId        uint64
	EmittedAt        uint64
	MaxFee           uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInstructionEmitted is a free log retrieval operation binding the contract event 0xa1eb98293d3e67bee56d1f8d1f6d3187e00caf47fb50b319f4ec50578eb20a77.
//
// Solidity: event InstructionEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint8 kind, uint64 paymentId, uint64 emittedAt, uint64 maxFee)
func (_WalletPayments *WalletPaymentsFilterer) FilterInstructionEmitted(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (*WalletPaymentsInstructionEmittedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "InstructionEmitted", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsInstructionEmittedIterator{contract: _WalletPayments.contract, event: "InstructionEmitted", logs: logs, sub: sub}, nil
}

// WatchInstructionEmitted is a free log subscription operation binding the contract event 0xa1eb98293d3e67bee56d1f8d1f6d3187e00caf47fb50b319f4ec50578eb20a77.
//
// Solidity: event InstructionEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint8 kind, uint64 paymentId, uint64 emittedAt, uint64 maxFee)
func (_WalletPayments *WalletPaymentsFilterer) WatchInstructionEmitted(opts *bind.WatchOpts, sink chan<- *WalletPaymentsInstructionEmitted, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "InstructionEmitted", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsInstructionEmitted)
				if err := _WalletPayments.contract.UnpackLog(event, "InstructionEmitted", log); err != nil {
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

// ParseInstructionEmitted is a log parse operation binding the contract event 0xa1eb98293d3e67bee56d1f8d1f6d3187e00caf47fb50b319f4ec50578eb20a77.
//
// Solidity: event InstructionEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint8 kind, uint64 paymentId, uint64 emittedAt, uint64 maxFee)
func (_WalletPayments *WalletPaymentsFilterer) ParseInstructionEmitted(log types.Log) (*WalletPaymentsInstructionEmitted, error) {
	event := new(WalletPaymentsInstructionEmitted)
	if err := _WalletPayments.contract.UnpackLog(event, "InstructionEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsInstructionSettledIterator is returned from FilterInstructionSettled and is used to iterate over the raw logs and unpacked data for InstructionSettled events raised by the WalletPayments contract.
type WalletPaymentsInstructionSettledIterator struct {
	Event *WalletPaymentsInstructionSettled // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsInstructionSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsInstructionSettled)
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
		it.Event = new(WalletPaymentsInstructionSettled)
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
func (it *WalletPaymentsInstructionSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsInstructionSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsInstructionSettled represents a InstructionSettled event raised by the WalletPayments contract.
type WalletPaymentsInstructionSettled struct {
	WalletId            [32]byte
	AccountIndex        uint32
	SequencePosition    uint64
	Attempt             uint32
	PackageHash         [32]byte
	ChainCommitmentHash [32]byte
	Settler             common.Address
	Message             []byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterInstructionSettled is a free log retrieval operation binding the contract event 0x8db15c3745d9e21da77aa3fe06a7495d4ddaeee70c075cecd0f3151308bd38e5.
//
// Solidity: event InstructionSettled(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, bytes32 chainCommitmentHash, address settler, bytes message)
func (_WalletPayments *WalletPaymentsFilterer) FilterInstructionSettled(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (*WalletPaymentsInstructionSettledIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "InstructionSettled", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsInstructionSettledIterator{contract: _WalletPayments.contract, event: "InstructionSettled", logs: logs, sub: sub}, nil
}

// WatchInstructionSettled is a free log subscription operation binding the contract event 0x8db15c3745d9e21da77aa3fe06a7495d4ddaeee70c075cecd0f3151308bd38e5.
//
// Solidity: event InstructionSettled(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, bytes32 chainCommitmentHash, address settler, bytes message)
func (_WalletPayments *WalletPaymentsFilterer) WatchInstructionSettled(opts *bind.WatchOpts, sink chan<- *WalletPaymentsInstructionSettled, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "InstructionSettled", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsInstructionSettled)
				if err := _WalletPayments.contract.UnpackLog(event, "InstructionSettled", log); err != nil {
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

// ParseInstructionSettled is a log parse operation binding the contract event 0x8db15c3745d9e21da77aa3fe06a7495d4ddaeee70c075cecd0f3151308bd38e5.
//
// Solidity: event InstructionSettled(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, bytes32 chainCommitmentHash, address settler, bytes message)
func (_WalletPayments *WalletPaymentsFilterer) ParseInstructionSettled(log types.Log) (*WalletPaymentsInstructionSettled, error) {
	event := new(WalletPaymentsInstructionSettled)
	if err := _WalletPayments.contract.UnpackLog(event, "InstructionSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsNativeNonceAccountAddedIterator is returned from FilterNativeNonceAccountAdded and is used to iterate over the raw logs and unpacked data for NativeNonceAccountAdded events raised by the WalletPayments contract.
type WalletPaymentsNativeNonceAccountAddedIterator struct {
	Event *WalletPaymentsNativeNonceAccountAdded // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsNativeNonceAccountAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsNativeNonceAccountAdded)
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
		it.Event = new(WalletPaymentsNativeNonceAccountAdded)
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
func (it *WalletPaymentsNativeNonceAccountAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsNativeNonceAccountAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsNativeNonceAccountAdded represents a NativeNonceAccountAdded event raised by the WalletPayments contract.
type WalletPaymentsNativeNonceAccountAdded struct {
	WalletRegistry       common.Address
	WalletId             [32]byte
	SourceId             [32]byte
	AccountAddress       string
	AuthorizationAddress common.Address
	InitialNonce         uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNativeNonceAccountAdded is a free log retrieval operation binding the contract event 0xa5b1d8da8d1af98b4f7a270eb9c3ca8def1a048115a12bf1a30757f4ca6eef31.
//
// Solidity: event NativeNonceAccountAdded(address indexed walletRegistry, bytes32 indexed walletId, bytes32 sourceId, string accountAddress, address authorizationAddress, uint64 initialNonce)
func (_WalletPayments *WalletPaymentsFilterer) FilterNativeNonceAccountAdded(opts *bind.FilterOpts, walletRegistry []common.Address, walletId [][32]byte) (*WalletPaymentsNativeNonceAccountAddedIterator, error) {

	var walletRegistryRule []interface{}
	for _, walletRegistryItem := range walletRegistry {
		walletRegistryRule = append(walletRegistryRule, walletRegistryItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "NativeNonceAccountAdded", walletRegistryRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsNativeNonceAccountAddedIterator{contract: _WalletPayments.contract, event: "NativeNonceAccountAdded", logs: logs, sub: sub}, nil
}

// WatchNativeNonceAccountAdded is a free log subscription operation binding the contract event 0xa5b1d8da8d1af98b4f7a270eb9c3ca8def1a048115a12bf1a30757f4ca6eef31.
//
// Solidity: event NativeNonceAccountAdded(address indexed walletRegistry, bytes32 indexed walletId, bytes32 sourceId, string accountAddress, address authorizationAddress, uint64 initialNonce)
func (_WalletPayments *WalletPaymentsFilterer) WatchNativeNonceAccountAdded(opts *bind.WatchOpts, sink chan<- *WalletPaymentsNativeNonceAccountAdded, walletRegistry []common.Address, walletId [][32]byte) (event.Subscription, error) {

	var walletRegistryRule []interface{}
	for _, walletRegistryItem := range walletRegistry {
		walletRegistryRule = append(walletRegistryRule, walletRegistryItem)
	}
	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "NativeNonceAccountAdded", walletRegistryRule, walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsNativeNonceAccountAdded)
				if err := _WalletPayments.contract.UnpackLog(event, "NativeNonceAccountAdded", log); err != nil {
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

// ParseNativeNonceAccountAdded is a log parse operation binding the contract event 0xa5b1d8da8d1af98b4f7a270eb9c3ca8def1a048115a12bf1a30757f4ca6eef31.
//
// Solidity: event NativeNonceAccountAdded(address indexed walletRegistry, bytes32 indexed walletId, bytes32 sourceId, string accountAddress, address authorizationAddress, uint64 initialNonce)
func (_WalletPayments *WalletPaymentsFilterer) ParseNativeNonceAccountAdded(log types.Log) (*WalletPaymentsNativeNonceAccountAdded, error) {
	event := new(WalletPaymentsNativeNonceAccountAdded)
	if err := _WalletPayments.contract.UnpackLog(event, "NativeNonceAccountAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsOperationNullifiedIterator is returned from FilterOperationNullified and is used to iterate over the raw logs and unpacked data for OperationNullified events raised by the WalletPayments contract.
type WalletPaymentsOperationNullifiedIterator struct {
	Event *WalletPaymentsOperationNullified // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsOperationNullifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsOperationNullified)
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
		it.Event = new(WalletPaymentsOperationNullified)
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
func (it *WalletPaymentsOperationNullifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsOperationNullifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsOperationNullified represents a OperationNullified event raised by the WalletPayments contract.
type WalletPaymentsOperationNullified struct {
	WalletId         [32]byte
	AccountIndex     uint32
	SequencePosition uint64
	PaymentId        uint64
	Attempt          uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOperationNullified is a free log retrieval operation binding the contract event 0xb02c63b61a14ac9af2bb575bc014b3c8225edc3f94ed467e31ce3040fcf92cc1.
//
// Solidity: event OperationNullified(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint64 paymentId, uint32 attempt)
func (_WalletPayments *WalletPaymentsFilterer) FilterOperationNullified(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (*WalletPaymentsOperationNullifiedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "OperationNullified", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsOperationNullifiedIterator{contract: _WalletPayments.contract, event: "OperationNullified", logs: logs, sub: sub}, nil
}

// WatchOperationNullified is a free log subscription operation binding the contract event 0xb02c63b61a14ac9af2bb575bc014b3c8225edc3f94ed467e31ce3040fcf92cc1.
//
// Solidity: event OperationNullified(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint64 paymentId, uint32 attempt)
func (_WalletPayments *WalletPaymentsFilterer) WatchOperationNullified(opts *bind.WatchOpts, sink chan<- *WalletPaymentsOperationNullified, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "OperationNullified", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsOperationNullified)
				if err := _WalletPayments.contract.UnpackLog(event, "OperationNullified", log); err != nil {
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

// ParseOperationNullified is a log parse operation binding the contract event 0xb02c63b61a14ac9af2bb575bc014b3c8225edc3f94ed467e31ce3040fcf92cc1.
//
// Solidity: event OperationNullified(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint64 paymentId, uint32 attempt)
func (_WalletPayments *WalletPaymentsFilterer) ParseOperationNullified(log types.Log) (*WalletPaymentsOperationNullified, error) {
	event := new(WalletPaymentsOperationNullified)
	if err := _WalletPayments.contract.UnpackLog(event, "OperationNullified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsPaymentBatchedIterator is returned from FilterPaymentBatched and is used to iterate over the raw logs and unpacked data for PaymentBatched events raised by the WalletPayments contract.
type WalletPaymentsPaymentBatchedIterator struct {
	Event *WalletPaymentsPaymentBatched // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsPaymentBatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsPaymentBatched)
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
		it.Event = new(WalletPaymentsPaymentBatched)
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
func (it *WalletPaymentsPaymentBatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsPaymentBatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsPaymentBatched represents a PaymentBatched event raised by the WalletPayments contract.
type WalletPaymentsPaymentBatched struct {
	InstructionId [32]byte
	PaymentId     uint64
	Message       []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPaymentBatched is a free log retrieval operation binding the contract event 0xb4fab23eb1b19257eafe05b07b472cc11ad71f934749abf78c95253866cc24d3.
//
// Solidity: event PaymentBatched(bytes32 indexed instructionId, uint64 indexed paymentId, bytes message)
func (_WalletPayments *WalletPaymentsFilterer) FilterPaymentBatched(opts *bind.FilterOpts, instructionId [][32]byte, paymentId []uint64) (*WalletPaymentsPaymentBatchedIterator, error) {

	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "PaymentBatched", instructionIdRule, paymentIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsPaymentBatchedIterator{contract: _WalletPayments.contract, event: "PaymentBatched", logs: logs, sub: sub}, nil
}

// WatchPaymentBatched is a free log subscription operation binding the contract event 0xb4fab23eb1b19257eafe05b07b472cc11ad71f934749abf78c95253866cc24d3.
//
// Solidity: event PaymentBatched(bytes32 indexed instructionId, uint64 indexed paymentId, bytes message)
func (_WalletPayments *WalletPaymentsFilterer) WatchPaymentBatched(opts *bind.WatchOpts, sink chan<- *WalletPaymentsPaymentBatched, instructionId [][32]byte, paymentId []uint64) (event.Subscription, error) {

	var instructionIdRule []interface{}
	for _, instructionIdItem := range instructionId {
		instructionIdRule = append(instructionIdRule, instructionIdItem)
	}
	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "PaymentBatched", instructionIdRule, paymentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsPaymentBatched)
				if err := _WalletPayments.contract.UnpackLog(event, "PaymentBatched", log); err != nil {
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

// ParsePaymentBatched is a log parse operation binding the contract event 0xb4fab23eb1b19257eafe05b07b472cc11ad71f934749abf78c95253866cc24d3.
//
// Solidity: event PaymentBatched(bytes32 indexed instructionId, uint64 indexed paymentId, bytes message)
func (_WalletPayments *WalletPaymentsFilterer) ParsePaymentBatched(log types.Log) (*WalletPaymentsPaymentBatched, error) {
	event := new(WalletPaymentsPaymentBatched)
	if err := _WalletPayments.contract.UnpackLog(event, "PaymentBatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsPaymentQueuedIterator is returned from FilterPaymentQueued and is used to iterate over the raw logs and unpacked data for PaymentQueued events raised by the WalletPayments contract.
type WalletPaymentsPaymentQueuedIterator struct {
	Event *WalletPaymentsPaymentQueued // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsPaymentQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsPaymentQueued)
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
		it.Event = new(WalletPaymentsPaymentQueued)
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
func (it *WalletPaymentsPaymentQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsPaymentQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsPaymentQueued represents a PaymentQueued event raised by the WalletPayments contract.
type WalletPaymentsPaymentQueued struct {
	WalletId         [32]byte
	AccountIndex     uint32
	PaymentId        uint64
	RecipientAddress string
	Amount           uint64
	MaxFee           uint64
	PaymentReference [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPaymentQueued is a free log retrieval operation binding the contract event 0x66d09eecbea30423568b5f143e3c73e10c8391255675768cc000e02cd5b29163.
//
// Solidity: event PaymentQueued(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 paymentId, string recipientAddress, uint64 amount, uint64 maxFee, bytes32 paymentReference)
func (_WalletPayments *WalletPaymentsFilterer) FilterPaymentQueued(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32) (*WalletPaymentsPaymentQueuedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "PaymentQueued", walletIdRule, accountIndexRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsPaymentQueuedIterator{contract: _WalletPayments.contract, event: "PaymentQueued", logs: logs, sub: sub}, nil
}

// WatchPaymentQueued is a free log subscription operation binding the contract event 0x66d09eecbea30423568b5f143e3c73e10c8391255675768cc000e02cd5b29163.
//
// Solidity: event PaymentQueued(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 paymentId, string recipientAddress, uint64 amount, uint64 maxFee, bytes32 paymentReference)
func (_WalletPayments *WalletPaymentsFilterer) WatchPaymentQueued(opts *bind.WatchOpts, sink chan<- *WalletPaymentsPaymentQueued, walletId [][32]byte, accountIndex []uint32) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "PaymentQueued", walletIdRule, accountIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsPaymentQueued)
				if err := _WalletPayments.contract.UnpackLog(event, "PaymentQueued", log); err != nil {
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

// ParsePaymentQueued is a log parse operation binding the contract event 0x66d09eecbea30423568b5f143e3c73e10c8391255675768cc000e02cd5b29163.
//
// Solidity: event PaymentQueued(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 paymentId, string recipientAddress, uint64 amount, uint64 maxFee, bytes32 paymentReference)
func (_WalletPayments *WalletPaymentsFilterer) ParsePaymentQueued(log types.Log) (*WalletPaymentsPaymentQueued, error) {
	event := new(WalletPaymentsPaymentQueued)
	if err := _WalletPayments.contract.UnpackLog(event, "PaymentQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsProjectBtcWalletParamsSetIterator is returned from FilterProjectBtcWalletParamsSet and is used to iterate over the raw logs and unpacked data for ProjectBtcWalletParamsSet events raised by the WalletPayments contract.
type WalletPaymentsProjectBtcWalletParamsSetIterator struct {
	Event *WalletPaymentsProjectBtcWalletParamsSet // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsProjectBtcWalletParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsProjectBtcWalletParamsSet)
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
		it.Event = new(WalletPaymentsProjectBtcWalletParamsSet)
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
func (it *WalletPaymentsProjectBtcWalletParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsProjectBtcWalletParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsProjectBtcWalletParamsSet represents a ProjectBtcWalletParamsSet event raised by the WalletPayments contract.
type WalletPaymentsProjectBtcWalletParamsSet struct {
	ProjectId [32]byte
	Params    IBtcParamsBtcWalletParams
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectBtcWalletParamsSet is a free log retrieval operation binding the contract event 0x3a59b720930ba55818131366bf6be373ce5c06e0ceff9332bc303a7a12e368ce.
//
// Solidity: event ProjectBtcWalletParamsSet(bytes32 indexed projectId, (uint64,uint32,uint8,uint8) params)
func (_WalletPayments *WalletPaymentsFilterer) FilterProjectBtcWalletParamsSet(opts *bind.FilterOpts, projectId [][32]byte) (*WalletPaymentsProjectBtcWalletParamsSetIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "ProjectBtcWalletParamsSet", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsProjectBtcWalletParamsSetIterator{contract: _WalletPayments.contract, event: "ProjectBtcWalletParamsSet", logs: logs, sub: sub}, nil
}

// WatchProjectBtcWalletParamsSet is a free log subscription operation binding the contract event 0x3a59b720930ba55818131366bf6be373ce5c06e0ceff9332bc303a7a12e368ce.
//
// Solidity: event ProjectBtcWalletParamsSet(bytes32 indexed projectId, (uint64,uint32,uint8,uint8) params)
func (_WalletPayments *WalletPaymentsFilterer) WatchProjectBtcWalletParamsSet(opts *bind.WatchOpts, sink chan<- *WalletPaymentsProjectBtcWalletParamsSet, projectId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "ProjectBtcWalletParamsSet", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsProjectBtcWalletParamsSet)
				if err := _WalletPayments.contract.UnpackLog(event, "ProjectBtcWalletParamsSet", log); err != nil {
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

// ParseProjectBtcWalletParamsSet is a log parse operation binding the contract event 0x3a59b720930ba55818131366bf6be373ce5c06e0ceff9332bc303a7a12e368ce.
//
// Solidity: event ProjectBtcWalletParamsSet(bytes32 indexed projectId, (uint64,uint32,uint8,uint8) params)
func (_WalletPayments *WalletPaymentsFilterer) ParseProjectBtcWalletParamsSet(log types.Log) (*WalletPaymentsProjectBtcWalletParamsSet, error) {
	event := new(WalletPaymentsProjectBtcWalletParamsSet)
	if err := _WalletPayments.contract.UnpackLog(event, "ProjectBtcWalletParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsProjectFeeScheduleClearedIterator is returned from FilterProjectFeeScheduleCleared and is used to iterate over the raw logs and unpacked data for ProjectFeeScheduleCleared events raised by the WalletPayments contract.
type WalletPaymentsProjectFeeScheduleClearedIterator struct {
	Event *WalletPaymentsProjectFeeScheduleCleared // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsProjectFeeScheduleClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsProjectFeeScheduleCleared)
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
		it.Event = new(WalletPaymentsProjectFeeScheduleCleared)
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
func (it *WalletPaymentsProjectFeeScheduleClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsProjectFeeScheduleClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsProjectFeeScheduleCleared represents a ProjectFeeScheduleCleared event raised by the WalletPayments contract.
type WalletPaymentsProjectFeeScheduleCleared struct {
	ProjectId [32]byte
	SourceId  [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectFeeScheduleCleared is a free log retrieval operation binding the contract event 0xe3f22e2cb60f792203c33a8b3ce8faae7df971c5bd56aa1beca573c904909aa2.
//
// Solidity: event ProjectFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId)
func (_WalletPayments *WalletPaymentsFilterer) FilterProjectFeeScheduleCleared(opts *bind.FilterOpts, projectId [][32]byte, sourceId [][32]byte) (*WalletPaymentsProjectFeeScheduleClearedIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "ProjectFeeScheduleCleared", projectIdRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsProjectFeeScheduleClearedIterator{contract: _WalletPayments.contract, event: "ProjectFeeScheduleCleared", logs: logs, sub: sub}, nil
}

// WatchProjectFeeScheduleCleared is a free log subscription operation binding the contract event 0xe3f22e2cb60f792203c33a8b3ce8faae7df971c5bd56aa1beca573c904909aa2.
//
// Solidity: event ProjectFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId)
func (_WalletPayments *WalletPaymentsFilterer) WatchProjectFeeScheduleCleared(opts *bind.WatchOpts, sink chan<- *WalletPaymentsProjectFeeScheduleCleared, projectId [][32]byte, sourceId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "ProjectFeeScheduleCleared", projectIdRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsProjectFeeScheduleCleared)
				if err := _WalletPayments.contract.UnpackLog(event, "ProjectFeeScheduleCleared", log); err != nil {
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

// ParseProjectFeeScheduleCleared is a log parse operation binding the contract event 0xe3f22e2cb60f792203c33a8b3ce8faae7df971c5bd56aa1beca573c904909aa2.
//
// Solidity: event ProjectFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId)
func (_WalletPayments *WalletPaymentsFilterer) ParseProjectFeeScheduleCleared(log types.Log) (*WalletPaymentsProjectFeeScheduleCleared, error) {
	event := new(WalletPaymentsProjectFeeScheduleCleared)
	if err := _WalletPayments.contract.UnpackLog(event, "ProjectFeeScheduleCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsProjectFeeScheduleSetIterator is returned from FilterProjectFeeScheduleSet and is used to iterate over the raw logs and unpacked data for ProjectFeeScheduleSet events raised by the WalletPayments contract.
type WalletPaymentsProjectFeeScheduleSetIterator struct {
	Event *WalletPaymentsProjectFeeScheduleSet // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsProjectFeeScheduleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsProjectFeeScheduleSet)
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
		it.Event = new(WalletPaymentsProjectFeeScheduleSet)
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
func (it *WalletPaymentsProjectFeeScheduleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsProjectFeeScheduleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsProjectFeeScheduleSet represents a ProjectFeeScheduleSet event raised by the WalletPayments contract.
type WalletPaymentsProjectFeeScheduleSet struct {
	ProjectId [32]byte
	SourceId  [32]byte
	Schedule  []IFeeSchedulesFeeSchedule
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectFeeScheduleSet is a free log retrieval operation binding the contract event 0xbf2c68fcd2f4ce03009cd21ea065e36803d64647a5840a1b97cdf4bdfc1dacfe.
//
// Solidity: event ProjectFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, (int16,uint16)[] schedule)
func (_WalletPayments *WalletPaymentsFilterer) FilterProjectFeeScheduleSet(opts *bind.FilterOpts, projectId [][32]byte, sourceId [][32]byte) (*WalletPaymentsProjectFeeScheduleSetIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "ProjectFeeScheduleSet", projectIdRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsProjectFeeScheduleSetIterator{contract: _WalletPayments.contract, event: "ProjectFeeScheduleSet", logs: logs, sub: sub}, nil
}

// WatchProjectFeeScheduleSet is a free log subscription operation binding the contract event 0xbf2c68fcd2f4ce03009cd21ea065e36803d64647a5840a1b97cdf4bdfc1dacfe.
//
// Solidity: event ProjectFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, (int16,uint16)[] schedule)
func (_WalletPayments *WalletPaymentsFilterer) WatchProjectFeeScheduleSet(opts *bind.WatchOpts, sink chan<- *WalletPaymentsProjectFeeScheduleSet, projectId [][32]byte, sourceId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}
	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "ProjectFeeScheduleSet", projectIdRule, sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsProjectFeeScheduleSet)
				if err := _WalletPayments.contract.UnpackLog(event, "ProjectFeeScheduleSet", log); err != nil {
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

// ParseProjectFeeScheduleSet is a log parse operation binding the contract event 0xbf2c68fcd2f4ce03009cd21ea065e36803d64647a5840a1b97cdf4bdfc1dacfe.
//
// Solidity: event ProjectFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, (int16,uint16)[] schedule)
func (_WalletPayments *WalletPaymentsFilterer) ParseProjectFeeScheduleSet(log types.Log) (*WalletPaymentsProjectFeeScheduleSet, error) {
	event := new(WalletPaymentsProjectFeeScheduleSet)
	if err := _WalletPayments.contract.UnpackLog(event, "ProjectFeeScheduleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsProjectProposersSetIterator is returned from FilterProjectProposersSet and is used to iterate over the raw logs and unpacked data for ProjectProposersSet events raised by the WalletPayments contract.
type WalletPaymentsProjectProposersSetIterator struct {
	Event *WalletPaymentsProjectProposersSet // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsProjectProposersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsProjectProposersSet)
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
		it.Event = new(WalletPaymentsProjectProposersSet)
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
func (it *WalletPaymentsProjectProposersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsProjectProposersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsProjectProposersSet represents a ProjectProposersSet event raised by the WalletPayments contract.
type WalletPaymentsProjectProposersSet struct {
	ProjectId [32]byte
	Proposers []common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProjectProposersSet is a free log retrieval operation binding the contract event 0x1027e8086fa8808793e837def6796fbf017796c26fdd96dadf5790efe171d3f5.
//
// Solidity: event ProjectProposersSet(bytes32 indexed projectId, address[] proposers)
func (_WalletPayments *WalletPaymentsFilterer) FilterProjectProposersSet(opts *bind.FilterOpts, projectId [][32]byte) (*WalletPaymentsProjectProposersSetIterator, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "ProjectProposersSet", projectIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsProjectProposersSetIterator{contract: _WalletPayments.contract, event: "ProjectProposersSet", logs: logs, sub: sub}, nil
}

// WatchProjectProposersSet is a free log subscription operation binding the contract event 0x1027e8086fa8808793e837def6796fbf017796c26fdd96dadf5790efe171d3f5.
//
// Solidity: event ProjectProposersSet(bytes32 indexed projectId, address[] proposers)
func (_WalletPayments *WalletPaymentsFilterer) WatchProjectProposersSet(opts *bind.WatchOpts, sink chan<- *WalletPaymentsProjectProposersSet, projectId [][32]byte) (event.Subscription, error) {

	var projectIdRule []interface{}
	for _, projectIdItem := range projectId {
		projectIdRule = append(projectIdRule, projectIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "ProjectProposersSet", projectIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsProjectProposersSet)
				if err := _WalletPayments.contract.UnpackLog(event, "ProjectProposersSet", log); err != nil {
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

// ParseProjectProposersSet is a log parse operation binding the contract event 0x1027e8086fa8808793e837def6796fbf017796c26fdd96dadf5790efe171d3f5.
//
// Solidity: event ProjectProposersSet(bytes32 indexed projectId, address[] proposers)
func (_WalletPayments *WalletPaymentsFilterer) ParseProjectProposersSet(log types.Log) (*WalletPaymentsProjectProposersSet, error) {
	event := new(WalletPaymentsProjectProposersSet)
	if err := _WalletPayments.contract.UnpackLog(event, "ProjectProposersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsProposalContendedIterator is returned from FilterProposalContended and is used to iterate over the raw logs and unpacked data for ProposalContended events raised by the WalletPayments contract.
type WalletPaymentsProposalContendedIterator struct {
	Event *WalletPaymentsProposalContended // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsProposalContendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsProposalContended)
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
		it.Event = new(WalletPaymentsProposalContended)
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
func (it *WalletPaymentsProposalContendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsProposalContendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsProposalContended represents a ProposalContended event raised by the WalletPayments contract.
type WalletPaymentsProposalContended struct {
	WalletId         [32]byte
	AccountIndex     uint32
	SequencePosition uint64
	Attempt          uint32
	PackageHash      [32]byte
	Proposer         common.Address
	Score            uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProposalContended is a free log retrieval operation binding the contract event 0xb5abc58896177f8f66bba6d6f987eb63c23f3f0ab65afa7671a1348dcabdb76f.
//
// Solidity: event ProposalContended(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, address proposer, uint64 score)
func (_WalletPayments *WalletPaymentsFilterer) FilterProposalContended(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (*WalletPaymentsProposalContendedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "ProposalContended", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsProposalContendedIterator{contract: _WalletPayments.contract, event: "ProposalContended", logs: logs, sub: sub}, nil
}

// WatchProposalContended is a free log subscription operation binding the contract event 0xb5abc58896177f8f66bba6d6f987eb63c23f3f0ab65afa7671a1348dcabdb76f.
//
// Solidity: event ProposalContended(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, address proposer, uint64 score)
func (_WalletPayments *WalletPaymentsFilterer) WatchProposalContended(opts *bind.WatchOpts, sink chan<- *WalletPaymentsProposalContended, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "ProposalContended", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsProposalContended)
				if err := _WalletPayments.contract.UnpackLog(event, "ProposalContended", log); err != nil {
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

// ParseProposalContended is a log parse operation binding the contract event 0xb5abc58896177f8f66bba6d6f987eb63c23f3f0ab65afa7671a1348dcabdb76f.
//
// Solidity: event ProposalContended(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, address proposer, uint64 score)
func (_WalletPayments *WalletPaymentsFilterer) ParseProposalContended(log types.Log) (*WalletPaymentsProposalContended, error) {
	event := new(WalletPaymentsProposalContended)
	if err := _WalletPayments.contract.UnpackLog(event, "ProposalContended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsProposalLeadingIterator is returned from FilterProposalLeading and is used to iterate over the raw logs and unpacked data for ProposalLeading events raised by the WalletPayments contract.
type WalletPaymentsProposalLeadingIterator struct {
	Event *WalletPaymentsProposalLeading // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsProposalLeadingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsProposalLeading)
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
		it.Event = new(WalletPaymentsProposalLeading)
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
func (it *WalletPaymentsProposalLeadingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsProposalLeadingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsProposalLeading represents a ProposalLeading event raised by the WalletPayments contract.
type WalletPaymentsProposalLeading struct {
	WalletId         [32]byte
	AccountIndex     uint32
	SequencePosition uint64
	Attempt          uint32
	PackageHash      [32]byte
	Proposer         common.Address
	Score            uint64
	GraceEndsAt      uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProposalLeading is a free log retrieval operation binding the contract event 0x9bfef8292f23fc1c0b1ca49d0c7f44635f250998e574ebd9c992339ac5457736.
//
// Solidity: event ProposalLeading(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, address proposer, uint64 score, uint64 graceEndsAt)
func (_WalletPayments *WalletPaymentsFilterer) FilterProposalLeading(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (*WalletPaymentsProposalLeadingIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "ProposalLeading", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsProposalLeadingIterator{contract: _WalletPayments.contract, event: "ProposalLeading", logs: logs, sub: sub}, nil
}

// WatchProposalLeading is a free log subscription operation binding the contract event 0x9bfef8292f23fc1c0b1ca49d0c7f44635f250998e574ebd9c992339ac5457736.
//
// Solidity: event ProposalLeading(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, address proposer, uint64 score, uint64 graceEndsAt)
func (_WalletPayments *WalletPaymentsFilterer) WatchProposalLeading(opts *bind.WatchOpts, sink chan<- *WalletPaymentsProposalLeading, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "ProposalLeading", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsProposalLeading)
				if err := _WalletPayments.contract.UnpackLog(event, "ProposalLeading", log); err != nil {
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

// ParseProposalLeading is a log parse operation binding the contract event 0x9bfef8292f23fc1c0b1ca49d0c7f44635f250998e574ebd9c992339ac5457736.
//
// Solidity: event ProposalLeading(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, address proposer, uint64 score, uint64 graceEndsAt)
func (_WalletPayments *WalletPaymentsFilterer) ParseProposalLeading(log types.Log) (*WalletPaymentsProposalLeading, error) {
	event := new(WalletPaymentsProposalLeading)
	if err := _WalletPayments.contract.UnpackLog(event, "ProposalLeading", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsProposerUrlSetIterator is returned from FilterProposerUrlSet and is used to iterate over the raw logs and unpacked data for ProposerUrlSet events raised by the WalletPayments contract.
type WalletPaymentsProposerUrlSetIterator struct {
	Event *WalletPaymentsProposerUrlSet // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsProposerUrlSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsProposerUrlSet)
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
		it.Event = new(WalletPaymentsProposerUrlSet)
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
func (it *WalletPaymentsProposerUrlSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsProposerUrlSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsProposerUrlSet represents a ProposerUrlSet event raised by the WalletPayments contract.
type WalletPaymentsProposerUrlSet struct {
	Proposer common.Address
	Url      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProposerUrlSet is a free log retrieval operation binding the contract event 0xee095b0a35b4b51332b007a1aad7a32e468c7db3c59888b420e208dcfbf11054.
//
// Solidity: event ProposerUrlSet(address indexed proposer, string url)
func (_WalletPayments *WalletPaymentsFilterer) FilterProposerUrlSet(opts *bind.FilterOpts, proposer []common.Address) (*WalletPaymentsProposerUrlSetIterator, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "ProposerUrlSet", proposerRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsProposerUrlSetIterator{contract: _WalletPayments.contract, event: "ProposerUrlSet", logs: logs, sub: sub}, nil
}

// WatchProposerUrlSet is a free log subscription operation binding the contract event 0xee095b0a35b4b51332b007a1aad7a32e468c7db3c59888b420e208dcfbf11054.
//
// Solidity: event ProposerUrlSet(address indexed proposer, string url)
func (_WalletPayments *WalletPaymentsFilterer) WatchProposerUrlSet(opts *bind.WatchOpts, sink chan<- *WalletPaymentsProposerUrlSet, proposer []common.Address) (event.Subscription, error) {

	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "ProposerUrlSet", proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsProposerUrlSet)
				if err := _WalletPayments.contract.UnpackLog(event, "ProposerUrlSet", log); err != nil {
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

// ParseProposerUrlSet is a log parse operation binding the contract event 0xee095b0a35b4b51332b007a1aad7a32e468c7db3c59888b420e208dcfbf11054.
//
// Solidity: event ProposerUrlSet(address indexed proposer, string url)
func (_WalletPayments *WalletPaymentsFilterer) ParseProposerUrlSet(log types.Log) (*WalletPaymentsProposerUrlSet, error) {
	event := new(WalletPaymentsProposerUrlSet)
	if err := _WalletPayments.contract.UnpackLog(event, "ProposerUrlSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsReissueEmittedIterator is returned from FilterReissueEmitted and is used to iterate over the raw logs and unpacked data for ReissueEmitted events raised by the WalletPayments contract.
type WalletPaymentsReissueEmittedIterator struct {
	Event *WalletPaymentsReissueEmitted // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsReissueEmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsReissueEmitted)
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
		it.Event = new(WalletPaymentsReissueEmitted)
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
func (it *WalletPaymentsReissueEmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsReissueEmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsReissueEmitted represents a ReissueEmitted event raised by the WalletPayments contract.
type WalletPaymentsReissueEmitted struct {
	WalletId            [32]byte
	AccountIndex        uint32
	SequencePosition    uint64
	Attempt             uint32
	NullifiedPaymentIds []uint64
	MaxFee              uint64
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReissueEmitted is a free log retrieval operation binding the contract event 0x32bc078622da3e0394697770c4d37b0d182f825b15f851074eb2002611fcb735.
//
// Solidity: event ReissueEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, uint64[] nullifiedPaymentIds, uint64 maxFee)
func (_WalletPayments *WalletPaymentsFilterer) FilterReissueEmitted(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (*WalletPaymentsReissueEmittedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "ReissueEmitted", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsReissueEmittedIterator{contract: _WalletPayments.contract, event: "ReissueEmitted", logs: logs, sub: sub}, nil
}

// WatchReissueEmitted is a free log subscription operation binding the contract event 0x32bc078622da3e0394697770c4d37b0d182f825b15f851074eb2002611fcb735.
//
// Solidity: event ReissueEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, uint64[] nullifiedPaymentIds, uint64 maxFee)
func (_WalletPayments *WalletPaymentsFilterer) WatchReissueEmitted(opts *bind.WatchOpts, sink chan<- *WalletPaymentsReissueEmitted, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "ReissueEmitted", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsReissueEmitted)
				if err := _WalletPayments.contract.UnpackLog(event, "ReissueEmitted", log); err != nil {
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

// ParseReissueEmitted is a log parse operation binding the contract event 0x32bc078622da3e0394697770c4d37b0d182f825b15f851074eb2002611fcb735.
//
// Solidity: event ReissueEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, uint64[] nullifiedPaymentIds, uint64 maxFee)
func (_WalletPayments *WalletPaymentsFilterer) ParseReissueEmitted(log types.Log) (*WalletPaymentsReissueEmitted, error) {
	event := new(WalletPaymentsReissueEmitted)
	if err := _WalletPayments.contract.UnpackLog(event, "ReissueEmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsSigningRequestedIterator is returned from FilterSigningRequested and is used to iterate over the raw logs and unpacked data for SigningRequested events raised by the WalletPayments contract.
type WalletPaymentsSigningRequestedIterator struct {
	Event *WalletPaymentsSigningRequested // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsSigningRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsSigningRequested)
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
		it.Event = new(WalletPaymentsSigningRequested)
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
func (it *WalletPaymentsSigningRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsSigningRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsSigningRequested represents a SigningRequested event raised by the WalletPayments contract.
type WalletPaymentsSigningRequested struct {
	WalletId            [32]byte
	AccountIndex        uint32
	SequencePosition    uint64
	Attempt             uint32
	PackageHash         [32]byte
	ChainCommitmentHash [32]byte
	InstructionId       [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSigningRequested is a free log retrieval operation binding the contract event 0x0a57e873974bd2b9015c221e7c9ceea0b491d8f7a78ad6bd3fba5d11fcfa5894.
//
// Solidity: event SigningRequested(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, bytes32 chainCommitmentHash, bytes32 instructionId)
func (_WalletPayments *WalletPaymentsFilterer) FilterSigningRequested(opts *bind.FilterOpts, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (*WalletPaymentsSigningRequestedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "SigningRequested", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsSigningRequestedIterator{contract: _WalletPayments.contract, event: "SigningRequested", logs: logs, sub: sub}, nil
}

// WatchSigningRequested is a free log subscription operation binding the contract event 0x0a57e873974bd2b9015c221e7c9ceea0b491d8f7a78ad6bd3fba5d11fcfa5894.
//
// Solidity: event SigningRequested(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, bytes32 chainCommitmentHash, bytes32 instructionId)
func (_WalletPayments *WalletPaymentsFilterer) WatchSigningRequested(opts *bind.WatchOpts, sink chan<- *WalletPaymentsSigningRequested, walletId [][32]byte, accountIndex []uint32, sequencePosition []uint64) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var accountIndexRule []interface{}
	for _, accountIndexItem := range accountIndex {
		accountIndexRule = append(accountIndexRule, accountIndexItem)
	}
	var sequencePositionRule []interface{}
	for _, sequencePositionItem := range sequencePosition {
		sequencePositionRule = append(sequencePositionRule, sequencePositionItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "SigningRequested", walletIdRule, accountIndexRule, sequencePositionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsSigningRequested)
				if err := _WalletPayments.contract.UnpackLog(event, "SigningRequested", log); err != nil {
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

// ParseSigningRequested is a log parse operation binding the contract event 0x0a57e873974bd2b9015c221e7c9ceea0b491d8f7a78ad6bd3fba5d11fcfa5894.
//
// Solidity: event SigningRequested(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, bytes32 chainCommitmentHash, bytes32 instructionId)
func (_WalletPayments *WalletPaymentsFilterer) ParseSigningRequested(log types.Log) (*WalletPaymentsSigningRequested, error) {
	event := new(WalletPaymentsSigningRequested)
	if err := _WalletPayments.contract.UnpackLog(event, "SigningRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsSourceEscrowsEnabledIterator is returned from FilterSourceEscrowsEnabled and is used to iterate over the raw logs and unpacked data for SourceEscrowsEnabled events raised by the WalletPayments contract.
type WalletPaymentsSourceEscrowsEnabledIterator struct {
	Event *WalletPaymentsSourceEscrowsEnabled // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsSourceEscrowsEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsSourceEscrowsEnabled)
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
		it.Event = new(WalletPaymentsSourceEscrowsEnabled)
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
func (it *WalletPaymentsSourceEscrowsEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsSourceEscrowsEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsSourceEscrowsEnabled represents a SourceEscrowsEnabled event raised by the WalletPayments contract.
type WalletPaymentsSourceEscrowsEnabled struct {
	SourceId         [32]byte
	CustodianWallets bool
	TeeWallets       bool
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSourceEscrowsEnabled is a free log retrieval operation binding the contract event 0x1d0c82ef72cce46a7bcbb20cad7bce29f5c3636df2a301f4344bfb7d71aebeed.
//
// Solidity: event SourceEscrowsEnabled(bytes32 indexed sourceId, bool custodianWallets, bool teeWallets)
func (_WalletPayments *WalletPaymentsFilterer) FilterSourceEscrowsEnabled(opts *bind.FilterOpts, sourceId [][32]byte) (*WalletPaymentsSourceEscrowsEnabledIterator, error) {

	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "SourceEscrowsEnabled", sourceIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsSourceEscrowsEnabledIterator{contract: _WalletPayments.contract, event: "SourceEscrowsEnabled", logs: logs, sub: sub}, nil
}

// WatchSourceEscrowsEnabled is a free log subscription operation binding the contract event 0x1d0c82ef72cce46a7bcbb20cad7bce29f5c3636df2a301f4344bfb7d71aebeed.
//
// Solidity: event SourceEscrowsEnabled(bytes32 indexed sourceId, bool custodianWallets, bool teeWallets)
func (_WalletPayments *WalletPaymentsFilterer) WatchSourceEscrowsEnabled(opts *bind.WatchOpts, sink chan<- *WalletPaymentsSourceEscrowsEnabled, sourceId [][32]byte) (event.Subscription, error) {

	var sourceIdRule []interface{}
	for _, sourceIdItem := range sourceId {
		sourceIdRule = append(sourceIdRule, sourceIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "SourceEscrowsEnabled", sourceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsSourceEscrowsEnabled)
				if err := _WalletPayments.contract.UnpackLog(event, "SourceEscrowsEnabled", log); err != nil {
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

// ParseSourceEscrowsEnabled is a log parse operation binding the contract event 0x1d0c82ef72cce46a7bcbb20cad7bce29f5c3636df2a301f4344bfb7d71aebeed.
//
// Solidity: event SourceEscrowsEnabled(bytes32 indexed sourceId, bool custodianWallets, bool teeWallets)
func (_WalletPayments *WalletPaymentsFilterer) ParseSourceEscrowsEnabled(log types.Log) (*WalletPaymentsSourceEscrowsEnabled, error) {
	event := new(WalletPaymentsSourceEscrowsEnabled)
	if err := _WalletPayments.contract.UnpackLog(event, "SourceEscrowsEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsSourcesEnabledIterator is returned from FilterSourcesEnabled and is used to iterate over the raw logs and unpacked data for SourcesEnabled events raised by the WalletPayments contract.
type WalletPaymentsSourcesEnabledIterator struct {
	Event *WalletPaymentsSourcesEnabled // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsSourcesEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsSourcesEnabled)
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
		it.Event = new(WalletPaymentsSourcesEnabled)
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
func (it *WalletPaymentsSourcesEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsSourcesEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsSourcesEnabled represents a SourcesEnabled event raised by the WalletPayments contract.
type WalletPaymentsSourcesEnabled struct {
	SourceIds [][32]byte
	Enabled   bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSourcesEnabled is a free log retrieval operation binding the contract event 0x70b7ae228ddc8d2b2d780456599c4cd77c53c076f88bf452f12ec5a50f225d6a.
//
// Solidity: event SourcesEnabled(bytes32[] sourceIds, bool enabled)
func (_WalletPayments *WalletPaymentsFilterer) FilterSourcesEnabled(opts *bind.FilterOpts) (*WalletPaymentsSourcesEnabledIterator, error) {

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "SourcesEnabled")
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsSourcesEnabledIterator{contract: _WalletPayments.contract, event: "SourcesEnabled", logs: logs, sub: sub}, nil
}

// WatchSourcesEnabled is a free log subscription operation binding the contract event 0x70b7ae228ddc8d2b2d780456599c4cd77c53c076f88bf452f12ec5a50f225d6a.
//
// Solidity: event SourcesEnabled(bytes32[] sourceIds, bool enabled)
func (_WalletPayments *WalletPaymentsFilterer) WatchSourcesEnabled(opts *bind.WatchOpts, sink chan<- *WalletPaymentsSourcesEnabled) (event.Subscription, error) {

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "SourcesEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsSourcesEnabled)
				if err := _WalletPayments.contract.UnpackLog(event, "SourcesEnabled", log); err != nil {
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

// ParseSourcesEnabled is a log parse operation binding the contract event 0x70b7ae228ddc8d2b2d780456599c4cd77c53c076f88bf452f12ec5a50f225d6a.
//
// Solidity: event SourcesEnabled(bytes32[] sourceIds, bool enabled)
func (_WalletPayments *WalletPaymentsFilterer) ParseSourcesEnabled(log types.Log) (*WalletPaymentsSourcesEnabled, error) {
	event := new(WalletPaymentsSourcesEnabled)
	if err := _WalletPayments.contract.UnpackLog(event, "SourcesEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsSourcesRegisteredIterator is returned from FilterSourcesRegistered and is used to iterate over the raw logs and unpacked data for SourcesRegistered events raised by the WalletPayments contract.
type WalletPaymentsSourcesRegisteredIterator struct {
	Event *WalletPaymentsSourcesRegistered // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsSourcesRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsSourcesRegistered)
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
		it.Event = new(WalletPaymentsSourcesRegistered)
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
func (it *WalletPaymentsSourcesRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsSourcesRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsSourcesRegistered represents a SourcesRegistered event raised by the WalletPayments contract.
type WalletPaymentsSourcesRegistered struct {
	Registrations []ISourceConfigSourceRegistration
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSourcesRegistered is a free log retrieval operation binding the contract event 0xf985fef73e1a4e9b26adfb93d7561232237050e9131cc92986ead5ec1c066405.
//
// Solidity: event SourcesRegistered((bytes32,bytes32,bytes32,uint8,uint8)[] registrations)
func (_WalletPayments *WalletPaymentsFilterer) FilterSourcesRegistered(opts *bind.FilterOpts) (*WalletPaymentsSourcesRegisteredIterator, error) {

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "SourcesRegistered")
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsSourcesRegisteredIterator{contract: _WalletPayments.contract, event: "SourcesRegistered", logs: logs, sub: sub}, nil
}

// WatchSourcesRegistered is a free log subscription operation binding the contract event 0xf985fef73e1a4e9b26adfb93d7561232237050e9131cc92986ead5ec1c066405.
//
// Solidity: event SourcesRegistered((bytes32,bytes32,bytes32,uint8,uint8)[] registrations)
func (_WalletPayments *WalletPaymentsFilterer) WatchSourcesRegistered(opts *bind.WatchOpts, sink chan<- *WalletPaymentsSourcesRegistered) (event.Subscription, error) {

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "SourcesRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsSourcesRegistered)
				if err := _WalletPayments.contract.UnpackLog(event, "SourcesRegistered", log); err != nil {
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

// ParseSourcesRegistered is a log parse operation binding the contract event 0xf985fef73e1a4e9b26adfb93d7561232237050e9131cc92986ead5ec1c066405.
//
// Solidity: event SourcesRegistered((bytes32,bytes32,bytes32,uint8,uint8)[] registrations)
func (_WalletPayments *WalletPaymentsFilterer) ParseSourcesRegistered(log types.Log) (*WalletPaymentsSourcesRegistered, error) {
	event := new(WalletPaymentsSourcesRegistered)
	if err := _WalletPayments.contract.UnpackLog(event, "SourcesRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the WalletPayments contract.
type WalletPaymentsTimelockedGovernanceCallCanceledIterator struct {
	Event *WalletPaymentsTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsTimelockedGovernanceCallCanceled)
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
		it.Event = new(WalletPaymentsTimelockedGovernanceCallCanceled)
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
func (it *WalletPaymentsTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the WalletPayments contract.
type WalletPaymentsTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_WalletPayments *WalletPaymentsFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*WalletPaymentsTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsTimelockedGovernanceCallCanceledIterator{contract: _WalletPayments.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_WalletPayments *WalletPaymentsFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *WalletPaymentsTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsTimelockedGovernanceCallCanceled)
				if err := _WalletPayments.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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

// ParseTimelockedGovernanceCallCanceled is a log parse operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_WalletPayments *WalletPaymentsFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*WalletPaymentsTimelockedGovernanceCallCanceled, error) {
	event := new(WalletPaymentsTimelockedGovernanceCallCanceled)
	if err := _WalletPayments.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the WalletPayments contract.
type WalletPaymentsTimelockedGovernanceCallExecutedIterator struct {
	Event *WalletPaymentsTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsTimelockedGovernanceCallExecuted)
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
		it.Event = new(WalletPaymentsTimelockedGovernanceCallExecuted)
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
func (it *WalletPaymentsTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the WalletPayments contract.
type WalletPaymentsTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_WalletPayments *WalletPaymentsFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*WalletPaymentsTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsTimelockedGovernanceCallExecutedIterator{contract: _WalletPayments.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_WalletPayments *WalletPaymentsFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *WalletPaymentsTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsTimelockedGovernanceCallExecuted)
				if err := _WalletPayments.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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

// ParseTimelockedGovernanceCallExecuted is a log parse operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_WalletPayments *WalletPaymentsFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*WalletPaymentsTimelockedGovernanceCallExecuted, error) {
	event := new(WalletPaymentsTimelockedGovernanceCallExecuted)
	if err := _WalletPayments.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsXrpEscrowInstructedIterator is returned from FilterXrpEscrowInstructed and is used to iterate over the raw logs and unpacked data for XrpEscrowInstructed events raised by the WalletPayments contract.
type WalletPaymentsXrpEscrowInstructedIterator struct {
	Event *WalletPaymentsXrpEscrowInstructed // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsXrpEscrowInstructedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsXrpEscrowInstructed)
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
		it.Event = new(WalletPaymentsXrpEscrowInstructed)
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
func (it *WalletPaymentsXrpEscrowInstructedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsXrpEscrowInstructedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsXrpEscrowInstructed represents a XrpEscrowInstructed event raised by the WalletPayments contract.
type WalletPaymentsXrpEscrowInstructed struct {
	WalletId      [32]byte
	PaymentId     uint64
	ReissueNumber uint64
	MaxFee        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterXrpEscrowInstructed is a free log retrieval operation binding the contract event 0x9115dc50889c6fc31bac5404320792694dc61307725c712d28fb6d45cf0308d6.
//
// Solidity: event XrpEscrowInstructed(bytes32 indexed walletId, uint64 indexed paymentId, uint64 reissueNumber, uint256 maxFee)
func (_WalletPayments *WalletPaymentsFilterer) FilterXrpEscrowInstructed(opts *bind.FilterOpts, walletId [][32]byte, paymentId []uint64) (*WalletPaymentsXrpEscrowInstructedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "XrpEscrowInstructed", walletIdRule, paymentIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsXrpEscrowInstructedIterator{contract: _WalletPayments.contract, event: "XrpEscrowInstructed", logs: logs, sub: sub}, nil
}

// WatchXrpEscrowInstructed is a free log subscription operation binding the contract event 0x9115dc50889c6fc31bac5404320792694dc61307725c712d28fb6d45cf0308d6.
//
// Solidity: event XrpEscrowInstructed(bytes32 indexed walletId, uint64 indexed paymentId, uint64 reissueNumber, uint256 maxFee)
func (_WalletPayments *WalletPaymentsFilterer) WatchXrpEscrowInstructed(opts *bind.WatchOpts, sink chan<- *WalletPaymentsXrpEscrowInstructed, walletId [][32]byte, paymentId []uint64) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "XrpEscrowInstructed", walletIdRule, paymentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsXrpEscrowInstructed)
				if err := _WalletPayments.contract.UnpackLog(event, "XrpEscrowInstructed", log); err != nil {
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

// ParseXrpEscrowInstructed is a log parse operation binding the contract event 0x9115dc50889c6fc31bac5404320792694dc61307725c712d28fb6d45cf0308d6.
//
// Solidity: event XrpEscrowInstructed(bytes32 indexed walletId, uint64 indexed paymentId, uint64 reissueNumber, uint256 maxFee)
func (_WalletPayments *WalletPaymentsFilterer) ParseXrpEscrowInstructed(log types.Log) (*WalletPaymentsXrpEscrowInstructed, error) {
	event := new(WalletPaymentsXrpEscrowInstructed)
	if err := _WalletPayments.contract.UnpackLog(event, "XrpEscrowInstructed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsXrpEscrowNullifiedIterator is returned from FilterXrpEscrowNullified and is used to iterate over the raw logs and unpacked data for XrpEscrowNullified events raised by the WalletPayments contract.
type WalletPaymentsXrpEscrowNullifiedIterator struct {
	Event *WalletPaymentsXrpEscrowNullified // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsXrpEscrowNullifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsXrpEscrowNullified)
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
		it.Event = new(WalletPaymentsXrpEscrowNullified)
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
func (it *WalletPaymentsXrpEscrowNullifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsXrpEscrowNullifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsXrpEscrowNullified represents a XrpEscrowNullified event raised by the WalletPayments contract.
type WalletPaymentsXrpEscrowNullified struct {
	WalletId      [32]byte
	PaymentId     uint64
	ReissueNumber uint64
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterXrpEscrowNullified is a free log retrieval operation binding the contract event 0xbf2a5468ab060717b74288db429fe3d2004699f1020f9f8800ec30bf2f2ac902.
//
// Solidity: event XrpEscrowNullified(bytes32 indexed walletId, uint64 indexed paymentId, uint64 reissueNumber)
func (_WalletPayments *WalletPaymentsFilterer) FilterXrpEscrowNullified(opts *bind.FilterOpts, walletId [][32]byte, paymentId []uint64) (*WalletPaymentsXrpEscrowNullifiedIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "XrpEscrowNullified", walletIdRule, paymentIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsXrpEscrowNullifiedIterator{contract: _WalletPayments.contract, event: "XrpEscrowNullified", logs: logs, sub: sub}, nil
}

// WatchXrpEscrowNullified is a free log subscription operation binding the contract event 0xbf2a5468ab060717b74288db429fe3d2004699f1020f9f8800ec30bf2f2ac902.
//
// Solidity: event XrpEscrowNullified(bytes32 indexed walletId, uint64 indexed paymentId, uint64 reissueNumber)
func (_WalletPayments *WalletPaymentsFilterer) WatchXrpEscrowNullified(opts *bind.WatchOpts, sink chan<- *WalletPaymentsXrpEscrowNullified, walletId [][32]byte, paymentId []uint64) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}
	var paymentIdRule []interface{}
	for _, paymentIdItem := range paymentId {
		paymentIdRule = append(paymentIdRule, paymentIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "XrpEscrowNullified", walletIdRule, paymentIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsXrpEscrowNullified)
				if err := _WalletPayments.contract.UnpackLog(event, "XrpEscrowNullified", log); err != nil {
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

// ParseXrpEscrowNullified is a log parse operation binding the contract event 0xbf2a5468ab060717b74288db429fe3d2004699f1020f9f8800ec30bf2f2ac902.
//
// Solidity: event XrpEscrowNullified(bytes32 indexed walletId, uint64 indexed paymentId, uint64 reissueNumber)
func (_WalletPayments *WalletPaymentsFilterer) ParseXrpEscrowNullified(log types.Log) (*WalletPaymentsXrpEscrowNullified, error) {
	event := new(WalletPaymentsXrpEscrowNullified)
	if err := _WalletPayments.contract.UnpackLog(event, "XrpEscrowNullified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletPaymentsXrpEscrowProfileSetIterator is returned from FilterXrpEscrowProfileSet and is used to iterate over the raw logs and unpacked data for XrpEscrowProfileSet events raised by the WalletPayments contract.
type WalletPaymentsXrpEscrowProfileSetIterator struct {
	Event *WalletPaymentsXrpEscrowProfileSet // Event containing the contract specifics and raw log

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
func (it *WalletPaymentsXrpEscrowProfileSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletPaymentsXrpEscrowProfileSet)
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
		it.Event = new(WalletPaymentsXrpEscrowProfileSet)
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
func (it *WalletPaymentsXrpEscrowProfileSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletPaymentsXrpEscrowProfileSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletPaymentsXrpEscrowProfileSet represents a XrpEscrowProfileSet event raised by the WalletPayments contract.
type WalletPaymentsXrpEscrowProfileSet struct {
	WalletId    [32]byte
	Destination string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterXrpEscrowProfileSet is a free log retrieval operation binding the contract event 0x844db090424904b284ad6b65b3c6ea6d72e4e4b57bbdf512d4824aea5b1f96cb.
//
// Solidity: event XrpEscrowProfileSet(bytes32 indexed walletId, string destination)
func (_WalletPayments *WalletPaymentsFilterer) FilterXrpEscrowProfileSet(opts *bind.FilterOpts, walletId [][32]byte) (*WalletPaymentsXrpEscrowProfileSetIterator, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletPayments.contract.FilterLogs(opts, "XrpEscrowProfileSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return &WalletPaymentsXrpEscrowProfileSetIterator{contract: _WalletPayments.contract, event: "XrpEscrowProfileSet", logs: logs, sub: sub}, nil
}

// WatchXrpEscrowProfileSet is a free log subscription operation binding the contract event 0x844db090424904b284ad6b65b3c6ea6d72e4e4b57bbdf512d4824aea5b1f96cb.
//
// Solidity: event XrpEscrowProfileSet(bytes32 indexed walletId, string destination)
func (_WalletPayments *WalletPaymentsFilterer) WatchXrpEscrowProfileSet(opts *bind.WatchOpts, sink chan<- *WalletPaymentsXrpEscrowProfileSet, walletId [][32]byte) (event.Subscription, error) {

	var walletIdRule []interface{}
	for _, walletIdItem := range walletId {
		walletIdRule = append(walletIdRule, walletIdItem)
	}

	logs, sub, err := _WalletPayments.contract.WatchLogs(opts, "XrpEscrowProfileSet", walletIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletPaymentsXrpEscrowProfileSet)
				if err := _WalletPayments.contract.UnpackLog(event, "XrpEscrowProfileSet", log); err != nil {
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

// ParseXrpEscrowProfileSet is a log parse operation binding the contract event 0x844db090424904b284ad6b65b3c6ea6d72e4e4b57bbdf512d4824aea5b1f96cb.
//
// Solidity: event XrpEscrowProfileSet(bytes32 indexed walletId, string destination)
func (_WalletPayments *WalletPaymentsFilterer) ParseXrpEscrowProfileSet(log types.Log) (*WalletPaymentsXrpEscrowProfileSet, error) {
	event := new(WalletPaymentsXrpEscrowProfileSet)
	if err := _WalletPayments.contract.UnpackLog(event, "XrpEscrowProfileSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
