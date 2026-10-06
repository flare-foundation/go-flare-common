// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package walletpayments

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

// IBtcAccountsBtcConfirmedAttemptCommitment is an auto generated low-level Go binding around an user-defined struct.
type IBtcAccountsBtcConfirmedAttemptCommitment struct {
	Attempt uint32
	Txid    [32]byte
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
	RuleVersion                uint8
	MinPaymentAmountSat        uint32
	DustRelayFeeSatPerKvB      uint16
	BatchDeadlineSeconds       uint32
	MaxOutputsPerPayment       uint8
	MaxChangeOutputs           uint8
	ScoreFreeChangeOutputs     uint8
	MaxSweptInputs             uint8
	MinFundingConfirmations    uint8
	AnchorSpendDepth           uint8
	FeeFloorBIPS               uint16
	FeeEstimateTargetBlocks    uint16
	FeeRateMinBIPS             uint16
	FeeRateMaxBIPS             uint16
	EscrowReclaimMarginSeconds uint32
}

// IBtcParamsBtcWalletParams is an auto generated low-level Go binding around an user-defined struct.
type IBtcParamsBtcWalletParams struct {
	BatchDeadlineSeconds    uint32
	MaxOutputsPerPayment    uint8
	MaxChangeOutputs        uint8
	ScoreFreeChangeOutputs  uint8
	MaxSweptInputs          uint8
	MinFundingConfirmations uint8
	AnchorSpendDepth        uint8
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
	FinalizationGraceSeconds uint16
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
	Kind                uint8
	PackageHash         [32]byte
	ChainCommitmentHash [32]byte
	Proposer            common.Address
	Score               uint64
	ConfirmedAttempt    uint32
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

// IFdc2ProofPolicyProofPolicy is an auto generated low-level Go binding around an user-defined struct.
type IFdc2ProofPolicyProofPolicy struct {
	RequiredTeeSignatures uint16
	ThresholdBIPS         uint16
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

// IWalletFeeProofProof is an auto generated low-level Go binding around an user-defined struct.
type IWalletFeeProofProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IWalletFeeProofRequestBody
	ResponseBody IWalletFeeProofResponseBody
}

// IWalletFeeProofRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IWalletFeeProofRequestBody struct {
	OpType                [32]byte
	SenderAddress         string
	FirstSequencePosition uint64
	Count                 uint64
}

// IWalletFeeProofResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IWalletFeeProofResponseBody struct {
	ActualFee     *big.Int
	AuthorisedFee *big.Int
}

// IWalletOperationStatusProof is an auto generated low-level Go binding around an user-defined struct.
type IWalletOperationStatusProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IWalletOperationStatusRequestBody
	ResponseBody IWalletOperationStatusResponseBody
}

// IWalletOperationStatusRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IWalletOperationStatusRequestBody struct {
	OpType                [32]byte
	SenderAddress         string
	PaymentId             uint64
	SpendingTransactionId [32]byte
}

// IWalletOperationStatusResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IWalletOperationStatusResponseBody struct {
	Kind                  uint8
	SequencePosition      uint64
	ConfirmedAttempt      uint32
	Nullified             bool
	TransactionId         [32]byte
	TransactionFee        *big.Int
	BlockNumber           uint64
	BlockTimestamp        uint64
	EscrowCreated         bool
	EscrowTxid            [32]byte
	EscrowVout            uint32
	HtlcSpent             bool
	SpendPath             uint8
	SpendingTransactionId [32]byte
	ReclaimFee            *big.Int
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
var WalletPaymentsMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AccountAddressAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountIndexAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountIndexMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AccountNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorIndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorNonceRegression\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorSetEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AttemptAlreadyConfirmed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AttemptNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AttemptSuperseded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuthorizationAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BtcConsensusParamsNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BtcEscrowAlreadyReclaimed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BtcEscrowNotExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BtcEscrowNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BtcEscrowNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotNullifyPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CommitmentMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ConfirmedAttemptMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CspAccountNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"delay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDelay\",\"type\":\"uint256\"}],\"name\":\"DelayTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DuplicateAnchor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyProposerList\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyScheduleNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EscrowProfileNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EscrowsNotEnabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Fdc2ProofPolicyNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeProofRangeInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"FeeScheduleConfigNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeScheduleNotSet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeSchedulesUnsupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GraceClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GraceOpen\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"provided\",\"type\":\"uint256\"}],\"name\":\"InsufficientFee\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientTeeSignatures\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBtcConsensusParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBtcEscrowProfile\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBtcEscrowTerms\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBtcWalletParams\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCspSourceSettings\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFdc2ProofPolicy\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeDelay\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"InvalidFeeFactor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"}],\"name\":\"InvalidFeeScheduleConfig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidGrace\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentInstructionCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestBody\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidXrpEscrowTerms\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LaneMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LeaderExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxFeeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoNewAnchors\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoPendingReset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonceMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEligible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotLaneTip\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToSettle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullifiedPaymentsUnsupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProjectOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationAlreadyNullified\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationNotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OperationStatusMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverrideAlreadyPending\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentAmountTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentRangeInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProposerNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReissueDeclarationMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceChainKeyTypeMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceDisabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceIdZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceKeyTypeNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceKeyTypeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceLimitsNotConfigured\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceModelChainMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceNetworkKeyTypeMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"SourceNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourceOpTypeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"SourcePaymentModelUnknown\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StoredAnchorsChanged\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockCallNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockValueNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenIdUnsupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManySchedules\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"registry\",\"type\":\"address\"}],\"name\":\"UnknownWalletRegistry\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"UnsupportedPaymentModel\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UsePaymentReissue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValueNotExpected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"expectedNetwork\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"actualNetwork\",\"type\":\"uint8\"}],\"name\":\"WalletNetworkMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyType\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"XrpEscrowNotFound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxSweptInputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minFundingConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"anchorSpendDepth\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIBtcParams.BtcWalletParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"AccountBtcWalletParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"}],\"name\":\"AccountFeeScheduleCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"accountHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIFeeSchedules.FeeSchedule[]\",\"name\":\"schedule\",\"type\":\"tuple[]\"}],\"name\":\"AccountFeeScheduleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"proposers\",\"type\":\"address[]\"}],\"name\":\"AccountProposersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"walletRegistry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"anchorCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"}],\"name\":\"BtcAccountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"walletRegistry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"anchorCount\",\"type\":\"uint32\"}],\"name\":\"BtcAnchorsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"anchorIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"}],\"name\":\"BtcAttemptSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"ruleVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"minPaymentAmountSat\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"dustRelayFeeSatPerKvB\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxSweptInputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minFundingConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"anchorSpendDepth\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"feeFloorBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeEstimateTargetBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeRateMinBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeRateMaxBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"escrowReclaimMarginSeconds\",\"type\":\"uint32\"}],\"indexed\":false,\"internalType\":\"structIBtcParams.BtcConsensusParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"BtcConsensusParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumICspInstructions.InstructionKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"}],\"name\":\"BtcEscrowEmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"counterpartyPubKey\",\"type\":\"bytes\"}],\"name\":\"BtcEscrowProfileSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"createPaymentId\",\"type\":\"uint64\"}],\"name\":\"BtcEscrowReclaimEmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"supersededAttempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedAttemptSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"claimBackAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"rewardEpochId\",\"type\":\"uint24\"}],\"name\":\"CspFeeForwarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"finalizationGraceSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"proposerRedundancy\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"proposerFeeWei\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structICspProposals.CspSourceSettings\",\"name\":\"settings\",\"type\":\"tuple\"}],\"name\":\"CspSourceSettingsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"walletRegistry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"opCommand\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"CustodianInstructionIssued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamond.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structIDiamond.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"generation\",\"type\":\"uint64\"}],\"name\":\"EligibleAdvanced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"requiredTeeSignatures\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIFdc2ProofPolicy.ProofPolicy\",\"name\":\"policy\",\"type\":\"tuple\"}],\"name\":\"Fdc2ProofPolicySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"FeeScheduleConfigsCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"indexed\":false,\"internalType\":\"structIFeeSchedules.FeeScheduleConfigInput[]\",\"name\":\"configs\",\"type\":\"tuple[]\"}],\"name\":\"FeeScheduleConfigsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"firstSequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"accountedPosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"authorisedFee\",\"type\":\"uint256\"}],\"name\":\"FeesAccounted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"enumICspInstructions.InstructionKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"emittedAt\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"}],\"name\":\"InstructionEmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"settler\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"InstructionSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"walletRegistry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"initialNonce\",\"type\":\"uint64\"}],\"name\":\"NativeNonceAccountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"}],\"name\":\"OperationNullified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"PaymentBatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"name\":\"PaymentQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxSweptInputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minFundingConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"anchorSpendDepth\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structIBtcParams.BtcWalletParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"ProjectBtcWalletParamsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"name\":\"ProjectFeeScheduleCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"indexed\":false,\"internalType\":\"structIFeeSchedules.FeeSchedule[]\",\"name\":\"schedule\",\"type\":\"tuple[]\"}],\"name\":\"ProjectFeeScheduleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"projectId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"proposers\",\"type\":\"address[]\"}],\"name\":\"ProjectProposersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"}],\"name\":\"ProposalContended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"graceEndsAt\",\"type\":\"uint64\"}],\"name\":\"ProposalLeading\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"}],\"name\":\"ProposerUrlSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint64[]\",\"name\":\"nullifiedPaymentIds\",\"type\":\"uint64[]\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"}],\"name\":\"ReissueEmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"instructionId\",\"type\":\"bytes32\"}],\"name\":\"SigningRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"custodianWallets\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"teeWallets\",\"type\":\"bool\"}],\"name\":\"SourceEscrowsEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"sourceIds\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"SourcesEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"enumPaymentModel\",\"name\":\"paymentModel\",\"type\":\"uint8\"},{\"internalType\":\"enumISourceConfig.ChainKind\",\"name\":\"chainKind\",\"type\":\"uint8\"}],\"indexed\":false,\"internalType\":\"structISourceConfig.SourceRegistration[]\",\"name\":\"registrations\",\"type\":\"tuple[]\"}],\"name\":\"SourcesRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reissueNumber\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"XrpEscrowInstructed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"reissueNumber\",\"type\":\"uint64\"}],\"name\":\"XrpEscrowNullified\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"}],\"name\":\"XrpEscrowProfileSet\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccountConfigured.Anchor[]\",\"name\":\"anchors\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBtcAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIBtcAccountConfigured.BtcAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structIBtcAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIBtcAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"addAnchors\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccountConfigured.Anchor[]\",\"name\":\"anchors\",\"type\":\"tuple[]\"}],\"internalType\":\"structIBtcAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIBtcAccountConfigured.BtcAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structIBtcAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIBtcAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"addBtcAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structINativeNonceAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumINativeNonceAccountConfigured.NativeNonceAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structINativeNonceAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structINativeNonceAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"addNativeNonceAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"firstSequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"count\",\"type\":\"uint64\"}],\"internalType\":\"structIWalletFeeProof.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"actualFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"authorisedFee\",\"type\":\"uint256\"}],\"internalType\":\"structIWalletFeeProof.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIWalletFeeProof.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"advanceAccountedPosition\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"clearAccountFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_sourceIds\",\"type\":\"bytes32[]\"}],\"name\":\"clearFeeScheduleConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"clearProjectFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"enumICspInstructions.EmissionMode\",\"name\":\"_mode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_maxFee\",\"type\":\"uint256\"}],\"name\":\"consolidate\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"}],\"internalType\":\"structIEscrows.EscrowTerms\",\"name\":\"_terms\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_maxFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"createEscrow\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"custodianWalletManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumIDiamond.FacetCutAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamond.FacetCut[]\",\"name\":\"_diamondCut\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_init\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_functionSelector\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facetAddress_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facetAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_facet\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"facetFunctionSelectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structIDiamondLoupe.Facet[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2Hub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2RequestFeeConfigurations\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2Verification\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"eligibleGeneration\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"}],\"internalType\":\"structICspProposalCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"proposerAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"paymentCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"}],\"internalType\":\"structICspProposalCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structICspProposalCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"}],\"internalType\":\"structIBtcAccounts.BtcConfirmedAttemptCommitment\",\"name\":\"_commitment\",\"type\":\"tuple\"}],\"name\":\"finalizeConfirmedAttempt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"eligibleGeneration\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"}],\"internalType\":\"structICspProposalCheck.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"proposerAddress\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"paymentCount\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"}],\"internalType\":\"structICspProposalCheck.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structICspProposalCheck.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"anchorIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"}],\"internalType\":\"structIBtcAccounts.BtcProposalCommitment\",\"name\":\"_commitment\",\"type\":\"tuple\"}],\"name\":\"finalizeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAccountBtcWalletParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxSweptInputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minFundingConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"anchorSpendDepth\",\"type\":\"uint8\"}],\"internalType\":\"structIBtcParams.BtcWalletParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAccountFeeSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structIFeeSchedules.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAccountIndex\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"_accountIndex\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAccountProposers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_proposers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAccountedPosition\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_accountedPosition\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"_anchorIndex\",\"type\":\"uint32\"}],\"name\":\"getAnchor\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisVout\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nextNonce\",\"type\":\"uint64\"}],\"internalType\":\"structIBtcAccounts.BtcAnchor\",\"name\":\"_anchor\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAnchorCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"_anchorCount\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_attempt\",\"type\":\"uint32\"}],\"name\":\"getAttemptNullifiedPaymentIds\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"_paymentIds\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getAttempts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"settledAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"}],\"internalType\":\"structICspInstructions.Attempt[]\",\"name\":\"_attempts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getBatchPaymentId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_batchPaymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getBtcAttemptCommitments\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccounts.BtcAttemptCommitment[]\",\"name\":\"_commitments\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getBtcAttempts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"settledAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"}],\"internalType\":\"structICspInstructions.Attempt[]\",\"name\":\"_attempts\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccounts.BtcAttemptCommitment[]\",\"name\":\"_commitments\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getBtcConsensusParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"ruleVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"minPaymentAmountSat\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"dustRelayFeeSatPerKvB\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxSweptInputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minFundingConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"anchorSpendDepth\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"feeFloorBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeEstimateTargetBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeRateMinBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeRateMaxBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"escrowReclaimMarginSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcParams.BtcConsensusParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getBtcEscrow\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"counterpartyPubKey\",\"type\":\"bytes\"}],\"internalType\":\"structIBtcEscrows.BtcEscrowTerms\",\"name\":\"_terms\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_reclaimSequencePosition\",\"type\":\"uint64\"}],\"name\":\"getBtcEscrowCreate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_found\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_createSequencePosition\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getBtcEscrowProfile\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_counterpartyPubKey\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_createPaymentId\",\"type\":\"uint64\"}],\"name\":\"getBtcEscrowReclaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_reclaimed\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_reclaimPaymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getBtcEscrowTerms\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"counterpartyPubKey\",\"type\":\"bytes\"}],\"internalType\":\"structIBtcEscrows.BtcEscrowTerms\",\"name\":\"_terms\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getBtcInstructionNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_generation\",\"type\":\"uint64\"}],\"name\":\"getBtcLeader\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lane\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"graceEndsAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"paymentCount\",\"type\":\"uint32\"},{\"internalType\":\"enumICspProposals.ProposalKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"confirmedAttempt\",\"type\":\"uint32\"}],\"internalType\":\"structICspProposals.Leader\",\"name\":\"_leader\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"anchorIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"}],\"internalType\":\"structIBtcAccounts.BtcProposalCommitment\",\"name\":\"_commitment\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getBtcParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"ruleVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"minPaymentAmountSat\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"dustRelayFeeSatPerKvB\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxSweptInputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minFundingConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"anchorSpendDepth\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"feeFloorBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeEstimateTargetBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeRateMinBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeRateMaxBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"escrowReclaimMarginSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcParams.BtcConsensusParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_generation\",\"type\":\"uint64\"}],\"name\":\"getBtcProposalCommitment\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"anchorIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"nextAnchorVout\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"txid\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nextAnchorTxid\",\"type\":\"bytes32\"}],\"internalType\":\"structIBtcAccounts.BtcProposalCommitment\",\"name\":\"_commitment\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getConfirmedAttempt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_confirmed\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_attempt\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_accountIndex\",\"type\":\"uint32\"}],\"name\":\"getCspAccount\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getCspAccountWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_accountIndex\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getCspSourceSettings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"finalizationGraceSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"proposerRedundancy\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"proposerFeeWei\",\"type\":\"uint128\"}],\"internalType\":\"structICspProposals.CspSourceSettings\",\"name\":\"_settings\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_accountHash\",\"type\":\"bytes32\"}],\"name\":\"getEffectiveSchedule\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"_feeSchedule\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getEligible\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"generation\",\"type\":\"uint64\"}],\"internalType\":\"structICspInstructions.Eligible\",\"name\":\"_eligible\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getEscrow\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"}],\"internalType\":\"structIEscrows.EscrowTerms\",\"name\":\"_terms\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFdc2ProofPolicy\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"requiredTeeSignatures\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"}],\"internalType\":\"structIFdc2ProofPolicy.ProofPolicy\",\"name\":\"_policy\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getFeeScheduleConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structIFeeSchedules.FeeScheduleConfig\",\"name\":\"_config\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_generation\",\"type\":\"uint64\"}],\"name\":\"getFinalizedHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_packageHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getInitialNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_initialNonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getInstruction\",\"outputs\":[{\"components\":[{\"internalType\":\"enumICspInstructions.InstructionKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lane\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"emittedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"fromPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"settled\",\"type\":\"bool\"}],\"internalType\":\"structICspInstructions.InstructionRecord\",\"name\":\"_instruction\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"_lane\",\"type\":\"uint32\"}],\"name\":\"getLaneLastSettled\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"used\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"}],\"internalType\":\"structICspInstructions.LaneUse\",\"name\":\"_use\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"getLatestAttemptNullifiedPaymentIds\",\"outputs\":[{\"internalType\":\"uint64[]\",\"name\":\"_nullifiedPaymentIds\",\"type\":\"uint64[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_generation\",\"type\":\"uint64\"}],\"name\":\"getLeader\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lane\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"graceEndsAt\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"paymentCount\",\"type\":\"uint32\"},{\"internalType\":\"enumICspProposals.ProposalKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"packageHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"chainCommitmentHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"score\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"confirmedAttempt\",\"type\":\"uint32\"}],\"internalType\":\"structICspProposals.Leader\",\"name\":\"_leader\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNetwork\",\"outputs\":[{\"internalType\":\"enumISourceConfig.Network\",\"name\":\"_network\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getNextPaymentId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nextPaymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getNextSequencePosition\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nextSequencePosition\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getNextUnconsumedPaymentId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nextUnconsumedPaymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getOperationSequencePosition\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isOperation\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"}],\"name\":\"getPaymentFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getPaymentHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_paymentHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getPendingResets\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"attempt\",\"type\":\"uint32\"}],\"internalType\":\"structICspInstructions.PendingReset[]\",\"name\":\"_pendingResets\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getProjectBtcWalletParams\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxSweptInputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minFundingConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"anchorSpendDepth\",\"type\":\"uint8\"}],\"internalType\":\"structIBtcParams.BtcWalletParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getProjectFeeSchedule\",\"outputs\":[{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structIFeeSchedules.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"getProjectProposers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_proposers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"}],\"name\":\"getProposerUrl\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getQueuedPayment\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"amount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"maxFee\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structICspQueue.QueuedPayment\",\"name\":\"_payment\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegisteredSourceIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_sourceIds\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getSettledBatch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"fromPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"}],\"internalType\":\"structICspQueue.SettledBatch\",\"name\":\"_batch\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getSettledBatches\",\"outputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"fromPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"toPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"}],\"internalType\":\"structICspQueue.SettledBatch[]\",\"name\":\"_batches\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_generation\",\"type\":\"uint64\"}],\"name\":\"getSettlementCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_cost\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"getSourceConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"enumPaymentModel\",\"name\":\"paymentModel\",\"type\":\"uint8\"},{\"internalType\":\"enumISourceConfig.ChainKind\",\"name\":\"chainKind\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"custodianEscrows\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"teeEscrows\",\"type\":\"bool\"}],\"internalType\":\"structISourceConfig.Source\",\"name\":\"_source\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount[]\",\"name\":\"_walletAccounts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getWalletId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getWalletRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registry\",\"type\":\"address\"}],\"name\":\"getWalletRegistryKind\",\"outputs\":[{\"internalType\":\"enumRegistryKind\",\"name\":\"_kind\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getXrpEscrow\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"preimageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"expiresAt\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"destination\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"nullified\",\"type\":\"bool\"}],\"internalType\":\"structIXrpEscrows.XrpEscrowTerms\",\"name\":\"_terms\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getXrpEscrowProfile\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_destination\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_proposer\",\"type\":\"address\"}],\"name\":\"isAllowedProposer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"isOperationPaymentId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_isOperation\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"isSourceEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"}],\"name\":\"isSourceRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_registered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_address\",\"type\":\"string\"}],\"name\":\"isValidAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_maxFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"nullifyOperation\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"_attempt\",\"type\":\"uint32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structPaymentInstruction\",\"name\":\"_paymentInstruction\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_createPaymentId\",\"type\":\"uint64\"},{\"internalType\":\"enumICspInstructions.EmissionMode\",\"name\":\"_mode\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_maxFee\",\"type\":\"uint256\"}],\"name\":\"reclaimEscrow\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"keyType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"enumPaymentModel\",\"name\":\"paymentModel\",\"type\":\"uint8\"},{\"internalType\":\"enumISourceConfig.ChainKind\",\"name\":\"chainKind\",\"type\":\"uint8\"}],\"internalType\":\"structISourceConfig.SourceRegistration[]\",\"name\":\"_registrations\",\"type\":\"tuple[]\"}],\"name\":\"registerSources\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structPaymentInstruction[]\",\"name\":\"_retainedInstructions\",\"type\":\"tuple[]\"},{\"internalType\":\"uint64[]\",\"name\":\"_nullifiedPaymentIds\",\"type\":\"uint64[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"maxFeePerPayment\",\"type\":\"uint256[]\"},{\"internalType\":\"int16[][]\",\"name\":\"factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"delaysSeconds\",\"type\":\"uint16[]\"}],\"internalType\":\"structReissueFeeParams\",\"name\":\"_reissueFeeParams\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"reissue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_finalized\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"_maxFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"reissueOperation\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"_attempt\",\"type\":\"uint32\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"removeAccountBtcWalletParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"removeAccountProposers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"removeProjectBtcWalletParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"}],\"name\":\"removeProjectProposers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_accountIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcAccountConfigured.Anchor[]\",\"name\":\"_anchors\",\"type\":\"tuple[]\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_proofOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestBtcAccountConfiguredAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_walletRegistry\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_accountAddress\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_teeIds\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_proofOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestNativeNonceAccountConfiguredAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"requestSigning\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxSweptInputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minFundingConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"anchorSpendDepth\",\"type\":\"uint8\"}],\"internalType\":\"structIBtcParams.BtcWalletParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"setAccountBtcWalletParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structIFeeSchedules.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"name\":\"setAccountFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"address[]\",\"name\":\"_proposers\",\"type\":\"address[]\"}],\"name\":\"setAccountProposers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"ruleVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"minPaymentAmountSat\",\"type\":\"uint32\"},{\"internalType\":\"uint16\",\"name\":\"dustRelayFeeSatPerKvB\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxSweptInputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minFundingConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"anchorSpendDepth\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"feeFloorBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeEstimateTargetBlocks\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeRateMinBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"feeRateMaxBIPS\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"escrowReclaimMarginSeconds\",\"type\":\"uint32\"}],\"internalType\":\"structIBtcParams.BtcConsensusParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"setBtcConsensusParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"_counterpartyPubKey\",\"type\":\"bytes\"}],\"name\":\"setBtcEscrowProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"finalizationGraceSeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"proposerRedundancy\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"proposerFeeWei\",\"type\":\"uint128\"}],\"internalType\":\"structICspProposals.CspSourceSettings\",\"name\":\"_settings\",\"type\":\"tuple\"}],\"name\":\"setCspSourceSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"requiredTeeSignatures\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"}],\"internalType\":\"structIFdc2ProofPolicy.ProofPolicy\",\"name\":\"_policy\",\"type\":\"tuple\"}],\"name\":\"setFdc2ProofPolicy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"maxDelaySeconds\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"maxSchedules\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"}],\"internalType\":\"structIFeeSchedules.FeeScheduleConfigInput[]\",\"name\":\"_configs\",\"type\":\"tuple[]\"}],\"name\":\"setFeeScheduleConfigs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"batchDeadlineSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maxOutputsPerPayment\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"scoreFreeChangeOutputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"maxSweptInputs\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"minFundingConfirmations\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"anchorSpendDepth\",\"type\":\"uint8\"}],\"internalType\":\"structIBtcParams.BtcWalletParams\",\"name\":\"_params\",\"type\":\"tuple\"}],\"name\":\"setProjectBtcWalletParams\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"int16\",\"name\":\"factorBIPS\",\"type\":\"int16\"},{\"internalType\":\"uint16\",\"name\":\"delaySeconds\",\"type\":\"uint16\"}],\"internalType\":\"structIFeeSchedules.FeeSchedule[]\",\"name\":\"_schedule\",\"type\":\"tuple[]\"}],\"name\":\"setProjectFeeSchedule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_projectId\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_proposers\",\"type\":\"address[]\"}],\"name\":\"setProjectProposers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_url\",\"type\":\"string\"}],\"name\":\"setProposerUrl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"_custodianWallets\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"_teeWallets\",\"type\":\"bool\"}],\"name\":\"setSourceEscrowsEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_sourceIds\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setSourcesEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"_destination\",\"type\":\"string\"}],\"name\":\"setXrpEscrowProfile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structWalletAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_sequencePosition\",\"type\":\"uint64\"}],\"name\":\"settle\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"int16[][]\",\"name\":\"_factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"_delaysSeconds\",\"type\":\"uint16[]\"}],\"name\":\"validateAndEncodeSchedules\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_encodedPerPayment\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"opType\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"senderAddress\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"paymentId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"spendingTransactionId\",\"type\":\"bytes32\"}],\"internalType\":\"structIWalletOperationStatus.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIWalletOperationStatus.OperationKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequencePosition\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"confirmedAttempt\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"nullified\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"transactionFee\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"blockTimestamp\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"escrowCreated\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"escrowTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"escrowVout\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"htlcSpent\",\"type\":\"bool\"},{\"internalType\":\"enumIWalletOperationStatus.SpendPath\",\"name\":\"spendPath\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"spendingTransactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"reclaimFee\",\"type\":\"uint256\"}],\"internalType\":\"structIWalletOperationStatus.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIWalletOperationStatus.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyWalletOperationStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_valid\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "WalletPayments",
}

// WalletPayments is an auto generated Go binding around an Ethereum contract.
type WalletPayments struct {
	abi abi.ABI
}

// NewWalletPayments creates a new instance of WalletPayments.
func NewWalletPayments() *WalletPayments {
	parsed, err := WalletPaymentsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &WalletPayments{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *WalletPayments) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackAddAnchors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22344637.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addAnchors(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof) returns()
func (walletPayments *WalletPayments) PackAddAnchors(proof IBtcAccountConfiguredProof) []byte {
	enc, err := walletPayments.abi.Pack("addAnchors", proof)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddAnchors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x22344637.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addAnchors(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof) returns()
func (walletPayments *WalletPayments) TryPackAddAnchors(proof IBtcAccountConfiguredProof) ([]byte, error) {
	return walletPayments.abi.Pack("addAnchors", proof)
}

// PackAddBtcAccount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cf216a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addBtcAccount(address _walletRegistry, bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof, address _authorizationAddress) returns()
func (walletPayments *WalletPayments) PackAddBtcAccount(walletRegistry common.Address, walletId [32]byte, proof IBtcAccountConfiguredProof, authorizationAddress common.Address) []byte {
	enc, err := walletPayments.abi.Pack("addBtcAccount", walletRegistry, walletId, proof, authorizationAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddBtcAccount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x79cf216a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addBtcAccount(address _walletRegistry, bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof, address _authorizationAddress) returns()
func (walletPayments *WalletPayments) TryPackAddBtcAccount(walletRegistry common.Address, walletId [32]byte, proof IBtcAccountConfiguredProof, authorizationAddress common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("addBtcAccount", walletRegistry, walletId, proof, authorizationAddress)
}

// PackAddNativeNonceAccount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb9820a1e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addNativeNonceAccount(address _walletRegistry, bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof, address _authorizationAddress) returns()
func (walletPayments *WalletPayments) PackAddNativeNonceAccount(walletRegistry common.Address, walletId [32]byte, proof INativeNonceAccountConfiguredProof, authorizationAddress common.Address) []byte {
	enc, err := walletPayments.abi.Pack("addNativeNonceAccount", walletRegistry, walletId, proof, authorizationAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddNativeNonceAccount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb9820a1e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addNativeNonceAccount(address _walletRegistry, bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof, address _authorizationAddress) returns()
func (walletPayments *WalletPayments) TryPackAddNativeNonceAccount(walletRegistry common.Address, walletId [32]byte, proof INativeNonceAccountConfiguredProof, authorizationAddress common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("addNativeNonceAccount", walletRegistry, walletId, proof, authorizationAddress)
}

// PackAdvanceAccountedPosition is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46506dd8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function advanceAccountedPosition(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,uint64),(uint256,uint256)) _proof) returns()
func (walletPayments *WalletPayments) PackAdvanceAccountedPosition(proof IWalletFeeProofProof) []byte {
	enc, err := walletPayments.abi.Pack("advanceAccountedPosition", proof)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAdvanceAccountedPosition is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46506dd8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function advanceAccountedPosition(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,uint64),(uint256,uint256)) _proof) returns()
func (walletPayments *WalletPayments) TryPackAdvanceAccountedPosition(proof IWalletFeeProofProof) ([]byte, error) {
	return walletPayments.abi.Pack("advanceAccountedPosition", proof)
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16fc2f6d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (walletPayments *WalletPayments) PackCancelGovernanceCall(encodedCall []byte) []byte {
	enc, err := walletPayments.abi.Pack("cancelGovernanceCall", encodedCall)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16fc2f6d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (walletPayments *WalletPayments) TryPackCancelGovernanceCall(encodedCall []byte) ([]byte, error) {
	return walletPayments.abi.Pack("cancelGovernanceCall", encodedCall)
}

// PackClearAccountFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf7443f84.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clearAccountFeeSchedule((bytes32,string) _account) returns()
func (walletPayments *WalletPayments) PackClearAccountFeeSchedule(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("clearAccountFeeSchedule", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClearAccountFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf7443f84.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clearAccountFeeSchedule((bytes32,string) _account) returns()
func (walletPayments *WalletPayments) TryPackClearAccountFeeSchedule(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("clearAccountFeeSchedule", account)
}

// PackClearFeeScheduleConfigs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11a5c047.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clearFeeScheduleConfigs(bytes32[] _sourceIds) returns()
func (walletPayments *WalletPayments) PackClearFeeScheduleConfigs(sourceIds [][32]byte) []byte {
	enc, err := walletPayments.abi.Pack("clearFeeScheduleConfigs", sourceIds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClearFeeScheduleConfigs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x11a5c047.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clearFeeScheduleConfigs(bytes32[] _sourceIds) returns()
func (walletPayments *WalletPayments) TryPackClearFeeScheduleConfigs(sourceIds [][32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("clearFeeScheduleConfigs", sourceIds)
}

// PackClearProjectFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7aecd6f7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function clearProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) returns()
func (walletPayments *WalletPayments) PackClearProjectFeeSchedule(projectId [32]byte, sourceId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("clearProjectFeeSchedule", projectId, sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackClearProjectFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7aecd6f7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function clearProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) returns()
func (walletPayments *WalletPayments) TryPackClearProjectFeeSchedule(projectId [32]byte, sourceId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("clearProjectFeeSchedule", projectId, sourceId)
}

// PackConsolidate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e5e66ac.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function consolidate((bytes32,string) _account, uint8 _mode, uint256 _maxFee) returns(uint64 _paymentId)
func (walletPayments *WalletPayments) PackConsolidate(account WalletAccount, mode uint8, maxFee *big.Int) []byte {
	enc, err := walletPayments.abi.Pack("consolidate", account, mode, maxFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackConsolidate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1e5e66ac.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function consolidate((bytes32,string) _account, uint8 _mode, uint256 _maxFee) returns(uint64 _paymentId)
func (walletPayments *WalletPayments) TryPackConsolidate(account WalletAccount, mode uint8, maxFee *big.Int) ([]byte, error) {
	return walletPayments.abi.Pack("consolidate", account, mode, maxFee)
}

// UnpackConsolidate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1e5e66ac.
//
// Solidity: function consolidate((bytes32,string) _account, uint8 _mode, uint256 _maxFee) returns(uint64 _paymentId)
func (walletPayments *WalletPayments) UnpackConsolidate(data []byte) (uint64, error) {
	out, err := walletPayments.abi.Unpack("consolidate", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackCreateEscrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd40003db.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function createEscrow((bytes32,string) _account, (bytes32,uint256,uint64) _terms, uint256 _maxFee, address _claimBackAddress) payable returns(uint64 _paymentId)
func (walletPayments *WalletPayments) PackCreateEscrow(account WalletAccount, terms IEscrowsEscrowTerms, maxFee *big.Int, claimBackAddress common.Address) []byte {
	enc, err := walletPayments.abi.Pack("createEscrow", account, terms, maxFee, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCreateEscrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd40003db.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function createEscrow((bytes32,string) _account, (bytes32,uint256,uint64) _terms, uint256 _maxFee, address _claimBackAddress) payable returns(uint64 _paymentId)
func (walletPayments *WalletPayments) TryPackCreateEscrow(account WalletAccount, terms IEscrowsEscrowTerms, maxFee *big.Int, claimBackAddress common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("createEscrow", account, terms, maxFee, claimBackAddress)
}

// UnpackCreateEscrow is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd40003db.
//
// Solidity: function createEscrow((bytes32,string) _account, (bytes32,uint256,uint64) _terms, uint256 _maxFee, address _claimBackAddress) payable returns(uint64 _paymentId)
func (walletPayments *WalletPayments) UnpackCreateEscrow(data []byte) (uint64, error) {
	out, err := walletPayments.abi.Unpack("createEscrow", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackCustodianWalletManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb04ccfe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function custodianWalletManager() view returns(address)
func (walletPayments *WalletPayments) PackCustodianWalletManager() []byte {
	enc, err := walletPayments.abi.Pack("custodianWalletManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackCustodianWalletManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb04ccfe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function custodianWalletManager() view returns(address)
func (walletPayments *WalletPayments) TryPackCustodianWalletManager() ([]byte, error) {
	return walletPayments.abi.Pack("custodianWalletManager")
}

// UnpackCustodianWalletManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbb04ccfe.
//
// Solidity: function custodianWalletManager() view returns(address)
func (walletPayments *WalletPayments) UnpackCustodianWalletManager(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("custodianWalletManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackDiamondCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f931c1c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (walletPayments *WalletPayments) PackDiamondCut(diamondCut []IDiamondFacetCut, init common.Address, calldata []byte) []byte {
	enc, err := walletPayments.abi.Pack("diamondCut", diamondCut, init, calldata)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackDiamondCut is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1f931c1c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (walletPayments *WalletPayments) TryPackDiamondCut(diamondCut []IDiamondFacetCut, init common.Address, calldata []byte) ([]byte, error) {
	return walletPayments.abi.Pack("diamondCut", diamondCut, init, calldata)
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) payable returns()
func (walletPayments *WalletPayments) PackExecuteGovernanceCall(encodedCall []byte) []byte {
	enc, err := walletPayments.abi.Pack("executeGovernanceCall", encodedCall)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) payable returns()
func (walletPayments *WalletPayments) TryPackExecuteGovernanceCall(encodedCall []byte) ([]byte, error) {
	return walletPayments.abi.Pack("executeGovernanceCall", encodedCall)
}

// PackFacetAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcdffacc6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (walletPayments *WalletPayments) PackFacetAddress(functionSelector [4]byte) []byte {
	enc, err := walletPayments.abi.Pack("facetAddress", functionSelector)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFacetAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xcdffacc6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (walletPayments *WalletPayments) TryPackFacetAddress(functionSelector [4]byte) ([]byte, error) {
	return walletPayments.abi.Pack("facetAddress", functionSelector)
}

// UnpackFacetAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (walletPayments *WalletPayments) UnpackFacetAddress(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("facetAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFacetAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52ef6b2c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (walletPayments *WalletPayments) PackFacetAddresses() []byte {
	enc, err := walletPayments.abi.Pack("facetAddresses")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFacetAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52ef6b2c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (walletPayments *WalletPayments) TryPackFacetAddresses() ([]byte, error) {
	return walletPayments.abi.Pack("facetAddresses")
}

// UnpackFacetAddresses is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (walletPayments *WalletPayments) UnpackFacetAddresses(data []byte) ([]common.Address, error) {
	out, err := walletPayments.abi.Unpack("facetAddresses", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackFacetFunctionSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xadfca15e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (walletPayments *WalletPayments) PackFacetFunctionSelectors(facet common.Address) []byte {
	enc, err := walletPayments.abi.Pack("facetFunctionSelectors", facet)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFacetFunctionSelectors is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xadfca15e.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (walletPayments *WalletPayments) TryPackFacetFunctionSelectors(facet common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("facetFunctionSelectors", facet)
}

// UnpackFacetFunctionSelectors is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] facetFunctionSelectors_)
func (walletPayments *WalletPayments) UnpackFacetFunctionSelectors(data []byte) ([][4]byte, error) {
	out, err := walletPayments.abi.Unpack("facetFunctionSelectors", data)
	if err != nil {
		return *new([][4]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)
	return out0, nil
}

// PackFacets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a0ed627.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (walletPayments *WalletPayments) PackFacets() []byte {
	enc, err := walletPayments.abi.Pack("facets")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFacets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7a0ed627.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (walletPayments *WalletPayments) TryPackFacets() ([]byte, error) {
	return walletPayments.abi.Pack("facets")
}

// UnpackFacets is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (walletPayments *WalletPayments) UnpackFacets(data []byte) ([]IDiamondLoupeFacet, error) {
	out, err := walletPayments.abi.Unpack("facets", data)
	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}
	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)
	return out0, nil
}

// PackFdc2Hub is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7566ff3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fdc2Hub() view returns(address)
func (walletPayments *WalletPayments) PackFdc2Hub() []byte {
	enc, err := walletPayments.abi.Pack("fdc2Hub")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFdc2Hub is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7566ff3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fdc2Hub() view returns(address)
func (walletPayments *WalletPayments) TryPackFdc2Hub() ([]byte, error) {
	return walletPayments.abi.Pack("fdc2Hub")
}

// UnpackFdc2Hub is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa7566ff3.
//
// Solidity: function fdc2Hub() view returns(address)
func (walletPayments *WalletPayments) UnpackFdc2Hub(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("fdc2Hub", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFdc2RequestFeeConfigurations is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x285434a2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fdc2RequestFeeConfigurations() view returns(address)
func (walletPayments *WalletPayments) PackFdc2RequestFeeConfigurations() []byte {
	enc, err := walletPayments.abi.Pack("fdc2RequestFeeConfigurations")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFdc2RequestFeeConfigurations is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x285434a2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fdc2RequestFeeConfigurations() view returns(address)
func (walletPayments *WalletPayments) TryPackFdc2RequestFeeConfigurations() ([]byte, error) {
	return walletPayments.abi.Pack("fdc2RequestFeeConfigurations")
}

// UnpackFdc2RequestFeeConfigurations is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x285434a2.
//
// Solidity: function fdc2RequestFeeConfigurations() view returns(address)
func (walletPayments *WalletPayments) UnpackFdc2RequestFeeConfigurations(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("fdc2RequestFeeConfigurations", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFdc2Verification is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf2e9839.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fdc2Verification() view returns(address)
func (walletPayments *WalletPayments) PackFdc2Verification() []byte {
	enc, err := walletPayments.abi.Pack("fdc2Verification")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFdc2Verification is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbf2e9839.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function fdc2Verification() view returns(address)
func (walletPayments *WalletPayments) TryPackFdc2Verification() ([]byte, error) {
	return walletPayments.abi.Pack("fdc2Verification")
}

// UnpackFdc2Verification is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf2e9839.
//
// Solidity: function fdc2Verification() view returns(address)
func (walletPayments *WalletPayments) UnpackFdc2Verification(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("fdc2Verification", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFinalizeConfirmedAttempt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd29e627c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function finalizeConfirmedAttempt(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint64,uint32,uint64,bytes32),(address,uint64,uint32,bytes32)) _proof, (uint32,bytes32) _commitment) returns()
func (walletPayments *WalletPayments) PackFinalizeConfirmedAttempt(proof ICspProposalCheckProof, commitment IBtcAccountsBtcConfirmedAttemptCommitment) []byte {
	enc, err := walletPayments.abi.Pack("finalizeConfirmedAttempt", proof, commitment)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFinalizeConfirmedAttempt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd29e627c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function finalizeConfirmedAttempt(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint64,uint32,uint64,bytes32),(address,uint64,uint32,bytes32)) _proof, (uint32,bytes32) _commitment) returns()
func (walletPayments *WalletPayments) TryPackFinalizeConfirmedAttempt(proof ICspProposalCheckProof, commitment IBtcAccountsBtcConfirmedAttemptCommitment) ([]byte, error) {
	return walletPayments.abi.Pack("finalizeConfirmedAttempt", proof, commitment)
}

// PackFinalizeProposal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5c4c7f3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function finalizeProposal(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint64,uint32,uint64,bytes32),(address,uint64,uint32,bytes32)) _proof, (uint32,uint64,uint32,bytes32,bytes32) _commitment) returns()
func (walletPayments *WalletPayments) PackFinalizeProposal(proof ICspProposalCheckProof, commitment IBtcAccountsBtcProposalCommitment) []byte {
	enc, err := walletPayments.abi.Pack("finalizeProposal", proof, commitment)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFinalizeProposal is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd5c4c7f3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function finalizeProposal(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(address,bytes32,uint32,uint64,uint32,uint64,bytes32),(address,uint64,uint32,bytes32)) _proof, (uint32,uint64,uint32,bytes32,bytes32) _commitment) returns()
func (walletPayments *WalletPayments) TryPackFinalizeProposal(proof ICspProposalCheckProof, commitment IBtcAccountsBtcProposalCommitment) ([]byte, error) {
	return walletPayments.abi.Pack("finalizeProposal", proof, commitment)
}

// PackFlareSystemsManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfaae7fc9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareSystemsManager() view returns(address)
func (walletPayments *WalletPayments) PackFlareSystemsManager() []byte {
	enc, err := walletPayments.abi.Pack("flareSystemsManager")
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
func (walletPayments *WalletPayments) TryPackFlareSystemsManager() ([]byte, error) {
	return walletPayments.abi.Pack("flareSystemsManager")
}

// UnpackFlareSystemsManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (walletPayments *WalletPayments) UnpackFlareSystemsManager(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("flareSystemsManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackFlareTeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x453f7ab4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareTeeManager() view returns(address)
func (walletPayments *WalletPayments) PackFlareTeeManager() []byte {
	enc, err := walletPayments.abi.Pack("flareTeeManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackFlareTeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x453f7ab4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function flareTeeManager() view returns(address)
func (walletPayments *WalletPayments) TryPackFlareTeeManager() ([]byte, error) {
	return walletPayments.abi.Pack("flareTeeManager")
}

// UnpackFlareTeeManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (walletPayments *WalletPayments) UnpackFlareTeeManager(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("flareTeeManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAccountBtcWalletParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32b4af53.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAccountBtcWalletParams((bytes32,string) _account) view returns((uint32,uint8,uint8,uint8,uint8,uint8,uint8) _params)
func (walletPayments *WalletPayments) PackGetAccountBtcWalletParams(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getAccountBtcWalletParams", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAccountBtcWalletParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32b4af53.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAccountBtcWalletParams((bytes32,string) _account) view returns((uint32,uint8,uint8,uint8,uint8,uint8,uint8) _params)
func (walletPayments *WalletPayments) TryPackGetAccountBtcWalletParams(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getAccountBtcWalletParams", account)
}

// UnpackGetAccountBtcWalletParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x32b4af53.
//
// Solidity: function getAccountBtcWalletParams((bytes32,string) _account) view returns((uint32,uint8,uint8,uint8,uint8,uint8,uint8) _params)
func (walletPayments *WalletPayments) UnpackGetAccountBtcWalletParams(data []byte) (IBtcParamsBtcWalletParams, error) {
	out, err := walletPayments.abi.Unpack("getAccountBtcWalletParams", data)
	if err != nil {
		return *new(IBtcParamsBtcWalletParams), err
	}
	out0 := *abi.ConvertType(out[0], new(IBtcParamsBtcWalletParams)).(*IBtcParamsBtcWalletParams)
	return out0, nil
}

// PackGetAccountFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c9e0c3f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAccountFeeSchedule((bytes32,string) _account) view returns((int16,uint16)[] _schedule)
func (walletPayments *WalletPayments) PackGetAccountFeeSchedule(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getAccountFeeSchedule", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAccountFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3c9e0c3f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAccountFeeSchedule((bytes32,string) _account) view returns((int16,uint16)[] _schedule)
func (walletPayments *WalletPayments) TryPackGetAccountFeeSchedule(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getAccountFeeSchedule", account)
}

// UnpackGetAccountFeeSchedule is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3c9e0c3f.
//
// Solidity: function getAccountFeeSchedule((bytes32,string) _account) view returns((int16,uint16)[] _schedule)
func (walletPayments *WalletPayments) UnpackGetAccountFeeSchedule(data []byte) ([]IFeeSchedulesFeeSchedule, error) {
	out, err := walletPayments.abi.Unpack("getAccountFeeSchedule", data)
	if err != nil {
		return *new([]IFeeSchedulesFeeSchedule), err
	}
	out0 := *abi.ConvertType(out[0], new([]IFeeSchedulesFeeSchedule)).(*[]IFeeSchedulesFeeSchedule)
	return out0, nil
}

// PackGetAccountIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92865318.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAccountIndex((bytes32,string) _account) view returns(uint32 _accountIndex)
func (walletPayments *WalletPayments) PackGetAccountIndex(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getAccountIndex", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAccountIndex is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x92865318.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAccountIndex((bytes32,string) _account) view returns(uint32 _accountIndex)
func (walletPayments *WalletPayments) TryPackGetAccountIndex(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getAccountIndex", account)
}

// UnpackGetAccountIndex is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x92865318.
//
// Solidity: function getAccountIndex((bytes32,string) _account) view returns(uint32 _accountIndex)
func (walletPayments *WalletPayments) UnpackGetAccountIndex(data []byte) (uint32, error) {
	out, err := walletPayments.abi.Unpack("getAccountIndex", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackGetAccountProposers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9438ac6b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAccountProposers((bytes32,string) _account) view returns(address[] _proposers)
func (walletPayments *WalletPayments) PackGetAccountProposers(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getAccountProposers", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAccountProposers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9438ac6b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAccountProposers((bytes32,string) _account) view returns(address[] _proposers)
func (walletPayments *WalletPayments) TryPackGetAccountProposers(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getAccountProposers", account)
}

// UnpackGetAccountProposers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9438ac6b.
//
// Solidity: function getAccountProposers((bytes32,string) _account) view returns(address[] _proposers)
func (walletPayments *WalletPayments) UnpackGetAccountProposers(data []byte) ([]common.Address, error) {
	out, err := walletPayments.abi.Unpack("getAccountProposers", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetAccountedPosition is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ceca9f8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAccountedPosition((bytes32,string) _account) view returns(uint64 _accountedPosition)
func (walletPayments *WalletPayments) PackGetAccountedPosition(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getAccountedPosition", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAccountedPosition is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9ceca9f8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAccountedPosition((bytes32,string) _account) view returns(uint64 _accountedPosition)
func (walletPayments *WalletPayments) TryPackGetAccountedPosition(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getAccountedPosition", account)
}

// UnpackGetAccountedPosition is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9ceca9f8.
//
// Solidity: function getAccountedPosition((bytes32,string) _account) view returns(uint64 _accountedPosition)
func (walletPayments *WalletPayments) UnpackGetAccountedPosition(data []byte) (uint64, error) {
	out, err := walletPayments.abi.Unpack("getAccountedPosition", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAddressUpdater() view returns(address)
func (walletPayments *WalletPayments) PackGetAddressUpdater() []byte {
	enc, err := walletPayments.abi.Pack("getAddressUpdater")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAddressUpdater is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5267a15d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAddressUpdater() view returns(address)
func (walletPayments *WalletPayments) TryPackGetAddressUpdater() ([]byte, error) {
	return walletPayments.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address)
func (walletPayments *WalletPayments) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAnchor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x207e9d61.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAnchor((bytes32,string) _account, uint32 _anchorIndex) view returns((bytes32,uint32,uint64) _anchor)
func (walletPayments *WalletPayments) PackGetAnchor(account WalletAccount, anchorIndex uint32) []byte {
	enc, err := walletPayments.abi.Pack("getAnchor", account, anchorIndex)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAnchor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x207e9d61.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAnchor((bytes32,string) _account, uint32 _anchorIndex) view returns((bytes32,uint32,uint64) _anchor)
func (walletPayments *WalletPayments) TryPackGetAnchor(account WalletAccount, anchorIndex uint32) ([]byte, error) {
	return walletPayments.abi.Pack("getAnchor", account, anchorIndex)
}

// UnpackGetAnchor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x207e9d61.
//
// Solidity: function getAnchor((bytes32,string) _account, uint32 _anchorIndex) view returns((bytes32,uint32,uint64) _anchor)
func (walletPayments *WalletPayments) UnpackGetAnchor(data []byte) (IBtcAccountsBtcAnchor, error) {
	out, err := walletPayments.abi.Unpack("getAnchor", data)
	if err != nil {
		return *new(IBtcAccountsBtcAnchor), err
	}
	out0 := *abi.ConvertType(out[0], new(IBtcAccountsBtcAnchor)).(*IBtcAccountsBtcAnchor)
	return out0, nil
}

// PackGetAnchorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00dc90fd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAnchorCount((bytes32,string) _account) view returns(uint32 _anchorCount)
func (walletPayments *WalletPayments) PackGetAnchorCount(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getAnchorCount", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAnchorCount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x00dc90fd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAnchorCount((bytes32,string) _account) view returns(uint32 _anchorCount)
func (walletPayments *WalletPayments) TryPackGetAnchorCount(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getAnchorCount", account)
}

// UnpackGetAnchorCount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x00dc90fd.
//
// Solidity: function getAnchorCount((bytes32,string) _account) view returns(uint32 _anchorCount)
func (walletPayments *WalletPayments) UnpackGetAnchorCount(data []byte) (uint32, error) {
	out, err := walletPayments.abi.Unpack("getAnchorCount", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackGetAttemptNullifiedPaymentIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb661c031.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAttemptNullifiedPaymentIds((bytes32,string) _account, uint64 _sequencePosition, uint32 _attempt) view returns(uint64[] _paymentIds)
func (walletPayments *WalletPayments) PackGetAttemptNullifiedPaymentIds(account WalletAccount, sequencePosition uint64, attempt uint32) []byte {
	enc, err := walletPayments.abi.Pack("getAttemptNullifiedPaymentIds", account, sequencePosition, attempt)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAttemptNullifiedPaymentIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb661c031.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAttemptNullifiedPaymentIds((bytes32,string) _account, uint64 _sequencePosition, uint32 _attempt) view returns(uint64[] _paymentIds)
func (walletPayments *WalletPayments) TryPackGetAttemptNullifiedPaymentIds(account WalletAccount, sequencePosition uint64, attempt uint32) ([]byte, error) {
	return walletPayments.abi.Pack("getAttemptNullifiedPaymentIds", account, sequencePosition, attempt)
}

// UnpackGetAttemptNullifiedPaymentIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xb661c031.
//
// Solidity: function getAttemptNullifiedPaymentIds((bytes32,string) _account, uint64 _sequencePosition, uint32 _attempt) view returns(uint64[] _paymentIds)
func (walletPayments *WalletPayments) UnpackGetAttemptNullifiedPaymentIds(data []byte) ([]uint64, error) {
	out, err := walletPayments.abi.Unpack("getAttemptNullifiedPaymentIds", data)
	if err != nil {
		return *new([]uint64), err
	}
	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	return out0, nil
}

// PackGetAttempts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde93b092.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAttempts((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,address,uint64,uint64)[] _attempts)
func (walletPayments *WalletPayments) PackGetAttempts(account WalletAccount, sequencePosition uint64) []byte {
	enc, err := walletPayments.abi.Pack("getAttempts", account, sequencePosition)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAttempts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xde93b092.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAttempts((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,address,uint64,uint64)[] _attempts)
func (walletPayments *WalletPayments) TryPackGetAttempts(account WalletAccount, sequencePosition uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getAttempts", account, sequencePosition)
}

// UnpackGetAttempts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xde93b092.
//
// Solidity: function getAttempts((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,address,uint64,uint64)[] _attempts)
func (walletPayments *WalletPayments) UnpackGetAttempts(data []byte) ([]ICspInstructionsAttempt, error) {
	out, err := walletPayments.abi.Unpack("getAttempts", data)
	if err != nil {
		return *new([]ICspInstructionsAttempt), err
	}
	out0 := *abi.ConvertType(out[0], new([]ICspInstructionsAttempt)).(*[]ICspInstructionsAttempt)
	return out0, nil
}

// PackGetAuthorizationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x410642e0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (walletPayments *WalletPayments) PackGetAuthorizationAddress(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getAuthorizationAddress", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetAuthorizationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x410642e0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (walletPayments *WalletPayments) TryPackGetAuthorizationAddress(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getAuthorizationAddress", account)
}

// UnpackGetAuthorizationAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (walletPayments *WalletPayments) UnpackGetAuthorizationAddress(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("getAuthorizationAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetBatchPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41952376.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBatchPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(uint64 _batchPaymentId)
func (walletPayments *WalletPayments) PackGetBatchPaymentId(account WalletAccount, paymentId uint64) []byte {
	enc, err := walletPayments.abi.Pack("getBatchPaymentId", account, paymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBatchPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41952376.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBatchPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(uint64 _batchPaymentId)
func (walletPayments *WalletPayments) TryPackGetBatchPaymentId(account WalletAccount, paymentId uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getBatchPaymentId", account, paymentId)
}

// UnpackGetBatchPaymentId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x41952376.
//
// Solidity: function getBatchPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(uint64 _batchPaymentId)
func (walletPayments *WalletPayments) UnpackGetBatchPaymentId(data []byte) (uint64, error) {
	out, err := walletPayments.abi.Unpack("getBatchPaymentId", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetBtcAttemptCommitments is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3abce1e5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBtcAttemptCommitments((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,uint32)[] _commitments)
func (walletPayments *WalletPayments) PackGetBtcAttemptCommitments(account WalletAccount, sequencePosition uint64) []byte {
	enc, err := walletPayments.abi.Pack("getBtcAttemptCommitments", account, sequencePosition)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBtcAttemptCommitments is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3abce1e5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBtcAttemptCommitments((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,uint32)[] _commitments)
func (walletPayments *WalletPayments) TryPackGetBtcAttemptCommitments(account WalletAccount, sequencePosition uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getBtcAttemptCommitments", account, sequencePosition)
}

// UnpackGetBtcAttemptCommitments is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3abce1e5.
//
// Solidity: function getBtcAttemptCommitments((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,uint32)[] _commitments)
func (walletPayments *WalletPayments) UnpackGetBtcAttemptCommitments(data []byte) ([]IBtcAccountsBtcAttemptCommitment, error) {
	out, err := walletPayments.abi.Unpack("getBtcAttemptCommitments", data)
	if err != nil {
		return *new([]IBtcAccountsBtcAttemptCommitment), err
	}
	out0 := *abi.ConvertType(out[0], new([]IBtcAccountsBtcAttemptCommitment)).(*[]IBtcAccountsBtcAttemptCommitment)
	return out0, nil
}

// PackGetBtcAttempts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x94d7a05b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBtcAttempts((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,address,uint64,uint64)[] _attempts, (bytes32,bytes32,uint32)[] _commitments)
func (walletPayments *WalletPayments) PackGetBtcAttempts(account WalletAccount, sequencePosition uint64) []byte {
	enc, err := walletPayments.abi.Pack("getBtcAttempts", account, sequencePosition)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBtcAttempts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x94d7a05b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBtcAttempts((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,address,uint64,uint64)[] _attempts, (bytes32,bytes32,uint32)[] _commitments)
func (walletPayments *WalletPayments) TryPackGetBtcAttempts(account WalletAccount, sequencePosition uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getBtcAttempts", account, sequencePosition)
}

// GetBtcAttemptsOutput serves as a container for the return parameters of contract
// method GetBtcAttempts.
type GetBtcAttemptsOutput struct {
	Attempts    []ICspInstructionsAttempt
	Commitments []IBtcAccountsBtcAttemptCommitment
}

// UnpackGetBtcAttempts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x94d7a05b.
//
// Solidity: function getBtcAttempts((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,bytes32,address,uint64,uint64)[] _attempts, (bytes32,bytes32,uint32)[] _commitments)
func (walletPayments *WalletPayments) UnpackGetBtcAttempts(data []byte) (GetBtcAttemptsOutput, error) {
	out, err := walletPayments.abi.Unpack("getBtcAttempts", data)
	outstruct := new(GetBtcAttemptsOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Attempts = *abi.ConvertType(out[0], new([]ICspInstructionsAttempt)).(*[]ICspInstructionsAttempt)
	outstruct.Commitments = *abi.ConvertType(out[1], new([]IBtcAccountsBtcAttemptCommitment)).(*[]IBtcAccountsBtcAttemptCommitment)
	return *outstruct, nil
}

// PackGetBtcConsensusParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62eae977.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBtcConsensusParams(bytes32 _sourceId) view returns((uint8,uint32,uint16,uint32,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint16,uint16,uint16,uint32) _params)
func (walletPayments *WalletPayments) PackGetBtcConsensusParams(sourceId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("getBtcConsensusParams", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBtcConsensusParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62eae977.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBtcConsensusParams(bytes32 _sourceId) view returns((uint8,uint32,uint16,uint32,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint16,uint16,uint16,uint32) _params)
func (walletPayments *WalletPayments) TryPackGetBtcConsensusParams(sourceId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("getBtcConsensusParams", sourceId)
}

// UnpackGetBtcConsensusParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62eae977.
//
// Solidity: function getBtcConsensusParams(bytes32 _sourceId) view returns((uint8,uint32,uint16,uint32,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint16,uint16,uint16,uint32) _params)
func (walletPayments *WalletPayments) UnpackGetBtcConsensusParams(data []byte) (IBtcParamsBtcConsensusParams, error) {
	out, err := walletPayments.abi.Unpack("getBtcConsensusParams", data)
	if err != nil {
		return *new(IBtcParamsBtcConsensusParams), err
	}
	out0 := *abi.ConvertType(out[0], new(IBtcParamsBtcConsensusParams)).(*IBtcParamsBtcConsensusParams)
	return out0, nil
}

// PackGetBtcEscrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x648d78de.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBtcEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint64,uint64,bytes) _terms)
func (walletPayments *WalletPayments) PackGetBtcEscrow(account WalletAccount, paymentId uint64) []byte {
	enc, err := walletPayments.abi.Pack("getBtcEscrow", account, paymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBtcEscrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x648d78de.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBtcEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint64,uint64,bytes) _terms)
func (walletPayments *WalletPayments) TryPackGetBtcEscrow(account WalletAccount, paymentId uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getBtcEscrow", account, paymentId)
}

// UnpackGetBtcEscrow is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x648d78de.
//
// Solidity: function getBtcEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint64,uint64,bytes) _terms)
func (walletPayments *WalletPayments) UnpackGetBtcEscrow(data []byte) (IBtcEscrowsBtcEscrowTerms, error) {
	out, err := walletPayments.abi.Unpack("getBtcEscrow", data)
	if err != nil {
		return *new(IBtcEscrowsBtcEscrowTerms), err
	}
	out0 := *abi.ConvertType(out[0], new(IBtcEscrowsBtcEscrowTerms)).(*IBtcEscrowsBtcEscrowTerms)
	return out0, nil
}

// PackGetBtcEscrowCreate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeca465b3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBtcEscrowCreate((bytes32,string) _account, uint64 _reclaimSequencePosition) view returns(bool _found, uint64 _createSequencePosition)
func (walletPayments *WalletPayments) PackGetBtcEscrowCreate(account WalletAccount, reclaimSequencePosition uint64) []byte {
	enc, err := walletPayments.abi.Pack("getBtcEscrowCreate", account, reclaimSequencePosition)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBtcEscrowCreate is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xeca465b3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBtcEscrowCreate((bytes32,string) _account, uint64 _reclaimSequencePosition) view returns(bool _found, uint64 _createSequencePosition)
func (walletPayments *WalletPayments) TryPackGetBtcEscrowCreate(account WalletAccount, reclaimSequencePosition uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getBtcEscrowCreate", account, reclaimSequencePosition)
}

// GetBtcEscrowCreateOutput serves as a container for the return parameters of contract
// method GetBtcEscrowCreate.
type GetBtcEscrowCreateOutput struct {
	Found                  bool
	CreateSequencePosition uint64
}

// UnpackGetBtcEscrowCreate is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xeca465b3.
//
// Solidity: function getBtcEscrowCreate((bytes32,string) _account, uint64 _reclaimSequencePosition) view returns(bool _found, uint64 _createSequencePosition)
func (walletPayments *WalletPayments) UnpackGetBtcEscrowCreate(data []byte) (GetBtcEscrowCreateOutput, error) {
	out, err := walletPayments.abi.Unpack("getBtcEscrowCreate", data)
	outstruct := new(GetBtcEscrowCreateOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Found = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CreateSequencePosition = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetBtcEscrowProfile is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x465f2f9a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBtcEscrowProfile((bytes32,string) _account) view returns(bytes _counterpartyPubKey)
func (walletPayments *WalletPayments) PackGetBtcEscrowProfile(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getBtcEscrowProfile", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBtcEscrowProfile is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x465f2f9a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBtcEscrowProfile((bytes32,string) _account) view returns(bytes _counterpartyPubKey)
func (walletPayments *WalletPayments) TryPackGetBtcEscrowProfile(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getBtcEscrowProfile", account)
}

// UnpackGetBtcEscrowProfile is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x465f2f9a.
//
// Solidity: function getBtcEscrowProfile((bytes32,string) _account) view returns(bytes _counterpartyPubKey)
func (walletPayments *WalletPayments) UnpackGetBtcEscrowProfile(data []byte) ([]byte, error) {
	out, err := walletPayments.abi.Unpack("getBtcEscrowProfile", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackGetBtcEscrowReclaim is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1aca3e0b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBtcEscrowReclaim((bytes32,string) _account, uint64 _createPaymentId) view returns(bool _reclaimed, uint64 _reclaimPaymentId)
func (walletPayments *WalletPayments) PackGetBtcEscrowReclaim(account WalletAccount, createPaymentId uint64) []byte {
	enc, err := walletPayments.abi.Pack("getBtcEscrowReclaim", account, createPaymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBtcEscrowReclaim is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1aca3e0b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBtcEscrowReclaim((bytes32,string) _account, uint64 _createPaymentId) view returns(bool _reclaimed, uint64 _reclaimPaymentId)
func (walletPayments *WalletPayments) TryPackGetBtcEscrowReclaim(account WalletAccount, createPaymentId uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getBtcEscrowReclaim", account, createPaymentId)
}

// GetBtcEscrowReclaimOutput serves as a container for the return parameters of contract
// method GetBtcEscrowReclaim.
type GetBtcEscrowReclaimOutput struct {
	Reclaimed        bool
	ReclaimPaymentId uint64
}

// UnpackGetBtcEscrowReclaim is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1aca3e0b.
//
// Solidity: function getBtcEscrowReclaim((bytes32,string) _account, uint64 _createPaymentId) view returns(bool _reclaimed, uint64 _reclaimPaymentId)
func (walletPayments *WalletPayments) UnpackGetBtcEscrowReclaim(data []byte) (GetBtcEscrowReclaimOutput, error) {
	out, err := walletPayments.abi.Unpack("getBtcEscrowReclaim", data)
	outstruct := new(GetBtcEscrowReclaimOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Reclaimed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.ReclaimPaymentId = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetBtcEscrowTerms is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4557c259.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBtcEscrowTerms((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,uint64,uint64,bytes) _terms)
func (walletPayments *WalletPayments) PackGetBtcEscrowTerms(account WalletAccount, sequencePosition uint64) []byte {
	enc, err := walletPayments.abi.Pack("getBtcEscrowTerms", account, sequencePosition)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBtcEscrowTerms is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4557c259.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBtcEscrowTerms((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,uint64,uint64,bytes) _terms)
func (walletPayments *WalletPayments) TryPackGetBtcEscrowTerms(account WalletAccount, sequencePosition uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getBtcEscrowTerms", account, sequencePosition)
}

// UnpackGetBtcEscrowTerms is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x4557c259.
//
// Solidity: function getBtcEscrowTerms((bytes32,string) _account, uint64 _sequencePosition) view returns((bytes32,uint64,uint64,bytes) _terms)
func (walletPayments *WalletPayments) UnpackGetBtcEscrowTerms(data []byte) (IBtcEscrowsBtcEscrowTerms, error) {
	out, err := walletPayments.abi.Unpack("getBtcEscrowTerms", data)
	if err != nil {
		return *new(IBtcEscrowsBtcEscrowTerms), err
	}
	out0 := *abi.ConvertType(out[0], new(IBtcEscrowsBtcEscrowTerms)).(*IBtcEscrowsBtcEscrowTerms)
	return out0, nil
}

// PackGetBtcInstructionNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x196153d7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBtcInstructionNonce((bytes32,string) _account, uint64 _sequencePosition) view returns(uint64 _nonce)
func (walletPayments *WalletPayments) PackGetBtcInstructionNonce(account WalletAccount, sequencePosition uint64) []byte {
	enc, err := walletPayments.abi.Pack("getBtcInstructionNonce", account, sequencePosition)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBtcInstructionNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x196153d7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBtcInstructionNonce((bytes32,string) _account, uint64 _sequencePosition) view returns(uint64 _nonce)
func (walletPayments *WalletPayments) TryPackGetBtcInstructionNonce(account WalletAccount, sequencePosition uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getBtcInstructionNonce", account, sequencePosition)
}

// UnpackGetBtcInstructionNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x196153d7.
//
// Solidity: function getBtcInstructionNonce((bytes32,string) _account, uint64 _sequencePosition) view returns(uint64 _nonce)
func (walletPayments *WalletPayments) UnpackGetBtcInstructionNonce(data []byte) (uint64, error) {
	out, err := walletPayments.abi.Unpack("getBtcInstructionNonce", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetBtcLeader is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd1e31bd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBtcLeader((bytes32,string) _account, uint64 _generation) view returns((bool,uint64,uint32,uint32,uint64,uint32,uint8,bytes32,bytes32,address,uint64,uint32) _leader, (uint32,uint64,uint32,bytes32,bytes32) _commitment)
func (walletPayments *WalletPayments) PackGetBtcLeader(account WalletAccount, generation uint64) []byte {
	enc, err := walletPayments.abi.Pack("getBtcLeader", account, generation)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBtcLeader is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd1e31bd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBtcLeader((bytes32,string) _account, uint64 _generation) view returns((bool,uint64,uint32,uint32,uint64,uint32,uint8,bytes32,bytes32,address,uint64,uint32) _leader, (uint32,uint64,uint32,bytes32,bytes32) _commitment)
func (walletPayments *WalletPayments) TryPackGetBtcLeader(account WalletAccount, generation uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getBtcLeader", account, generation)
}

// GetBtcLeaderOutput serves as a container for the return parameters of contract
// method GetBtcLeader.
type GetBtcLeaderOutput struct {
	Leader     ICspProposalsLeader
	Commitment IBtcAccountsBtcProposalCommitment
}

// UnpackGetBtcLeader is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbd1e31bd.
//
// Solidity: function getBtcLeader((bytes32,string) _account, uint64 _generation) view returns((bool,uint64,uint32,uint32,uint64,uint32,uint8,bytes32,bytes32,address,uint64,uint32) _leader, (uint32,uint64,uint32,bytes32,bytes32) _commitment)
func (walletPayments *WalletPayments) UnpackGetBtcLeader(data []byte) (GetBtcLeaderOutput, error) {
	out, err := walletPayments.abi.Unpack("getBtcLeader", data)
	outstruct := new(GetBtcLeaderOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Leader = *abi.ConvertType(out[0], new(ICspProposalsLeader)).(*ICspProposalsLeader)
	outstruct.Commitment = *abi.ConvertType(out[1], new(IBtcAccountsBtcProposalCommitment)).(*IBtcAccountsBtcProposalCommitment)
	return *outstruct, nil
}

// PackGetBtcParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d2cf8c9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBtcParams((bytes32,string) _account) view returns((uint8,uint32,uint16,uint32,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint16,uint16,uint16,uint32) _params)
func (walletPayments *WalletPayments) PackGetBtcParams(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getBtcParams", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBtcParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5d2cf8c9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBtcParams((bytes32,string) _account) view returns((uint8,uint32,uint16,uint32,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint16,uint16,uint16,uint32) _params)
func (walletPayments *WalletPayments) TryPackGetBtcParams(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getBtcParams", account)
}

// UnpackGetBtcParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5d2cf8c9.
//
// Solidity: function getBtcParams((bytes32,string) _account) view returns((uint8,uint32,uint16,uint32,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint16,uint16,uint16,uint32) _params)
func (walletPayments *WalletPayments) UnpackGetBtcParams(data []byte) (IBtcParamsBtcConsensusParams, error) {
	out, err := walletPayments.abi.Unpack("getBtcParams", data)
	if err != nil {
		return *new(IBtcParamsBtcConsensusParams), err
	}
	out0 := *abi.ConvertType(out[0], new(IBtcParamsBtcConsensusParams)).(*IBtcParamsBtcConsensusParams)
	return out0, nil
}

// PackGetBtcProposalCommitment is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa27afb27.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getBtcProposalCommitment((bytes32,string) _account, uint64 _generation) view returns((uint32,uint64,uint32,bytes32,bytes32) _commitment)
func (walletPayments *WalletPayments) PackGetBtcProposalCommitment(account WalletAccount, generation uint64) []byte {
	enc, err := walletPayments.abi.Pack("getBtcProposalCommitment", account, generation)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetBtcProposalCommitment is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa27afb27.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getBtcProposalCommitment((bytes32,string) _account, uint64 _generation) view returns((uint32,uint64,uint32,bytes32,bytes32) _commitment)
func (walletPayments *WalletPayments) TryPackGetBtcProposalCommitment(account WalletAccount, generation uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getBtcProposalCommitment", account, generation)
}

// UnpackGetBtcProposalCommitment is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa27afb27.
//
// Solidity: function getBtcProposalCommitment((bytes32,string) _account, uint64 _generation) view returns((uint32,uint64,uint32,bytes32,bytes32) _commitment)
func (walletPayments *WalletPayments) UnpackGetBtcProposalCommitment(data []byte) (IBtcAccountsBtcProposalCommitment, error) {
	out, err := walletPayments.abi.Unpack("getBtcProposalCommitment", data)
	if err != nil {
		return *new(IBtcAccountsBtcProposalCommitment), err
	}
	out0 := *abi.ConvertType(out[0], new(IBtcAccountsBtcProposalCommitment)).(*IBtcAccountsBtcProposalCommitment)
	return out0, nil
}

// PackGetConfirmedAttempt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8c7c8f18.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getConfirmedAttempt((bytes32,string) _account, uint64 _sequencePosition) view returns(bool _confirmed, uint32 _attempt)
func (walletPayments *WalletPayments) PackGetConfirmedAttempt(account WalletAccount, sequencePosition uint64) []byte {
	enc, err := walletPayments.abi.Pack("getConfirmedAttempt", account, sequencePosition)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetConfirmedAttempt is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8c7c8f18.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getConfirmedAttempt((bytes32,string) _account, uint64 _sequencePosition) view returns(bool _confirmed, uint32 _attempt)
func (walletPayments *WalletPayments) TryPackGetConfirmedAttempt(account WalletAccount, sequencePosition uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getConfirmedAttempt", account, sequencePosition)
}

// GetConfirmedAttemptOutput serves as a container for the return parameters of contract
// method GetConfirmedAttempt.
type GetConfirmedAttemptOutput struct {
	Confirmed bool
	Attempt   uint32
}

// UnpackGetConfirmedAttempt is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8c7c8f18.
//
// Solidity: function getConfirmedAttempt((bytes32,string) _account, uint64 _sequencePosition) view returns(bool _confirmed, uint32 _attempt)
func (walletPayments *WalletPayments) UnpackGetConfirmedAttempt(data []byte) (GetConfirmedAttemptOutput, error) {
	out, err := walletPayments.abi.Unpack("getConfirmedAttempt", data)
	outstruct := new(GetConfirmedAttemptOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.Confirmed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Attempt = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	return *outstruct, nil
}

// PackGetCspAccount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc154c258.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCspAccount(address _walletRegistry, bytes32 _walletId, uint32 _accountIndex) view returns((bytes32,string) _account)
func (walletPayments *WalletPayments) PackGetCspAccount(walletRegistry common.Address, walletId [32]byte, accountIndex uint32) []byte {
	enc, err := walletPayments.abi.Pack("getCspAccount", walletRegistry, walletId, accountIndex)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCspAccount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc154c258.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCspAccount(address _walletRegistry, bytes32 _walletId, uint32 _accountIndex) view returns((bytes32,string) _account)
func (walletPayments *WalletPayments) TryPackGetCspAccount(walletRegistry common.Address, walletId [32]byte, accountIndex uint32) ([]byte, error) {
	return walletPayments.abi.Pack("getCspAccount", walletRegistry, walletId, accountIndex)
}

// UnpackGetCspAccount is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc154c258.
//
// Solidity: function getCspAccount(address _walletRegistry, bytes32 _walletId, uint32 _accountIndex) view returns((bytes32,string) _account)
func (walletPayments *WalletPayments) UnpackGetCspAccount(data []byte) (WalletAccount, error) {
	out, err := walletPayments.abi.Unpack("getCspAccount", data)
	if err != nil {
		return *new(WalletAccount), err
	}
	out0 := *abi.ConvertType(out[0], new(WalletAccount)).(*WalletAccount)
	return out0, nil
}

// PackGetCspAccountWallet is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x409bf1c5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCspAccountWallet((bytes32,string) _account) view returns(address _walletRegistry, bytes32 _walletId, uint32 _accountIndex)
func (walletPayments *WalletPayments) PackGetCspAccountWallet(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getCspAccountWallet", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCspAccountWallet is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x409bf1c5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCspAccountWallet((bytes32,string) _account) view returns(address _walletRegistry, bytes32 _walletId, uint32 _accountIndex)
func (walletPayments *WalletPayments) TryPackGetCspAccountWallet(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getCspAccountWallet", account)
}

// GetCspAccountWalletOutput serves as a container for the return parameters of contract
// method GetCspAccountWallet.
type GetCspAccountWalletOutput struct {
	WalletRegistry common.Address
	WalletId       [32]byte
	AccountIndex   uint32
}

// UnpackGetCspAccountWallet is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x409bf1c5.
//
// Solidity: function getCspAccountWallet((bytes32,string) _account) view returns(address _walletRegistry, bytes32 _walletId, uint32 _accountIndex)
func (walletPayments *WalletPayments) UnpackGetCspAccountWallet(data []byte) (GetCspAccountWalletOutput, error) {
	out, err := walletPayments.abi.Unpack("getCspAccountWallet", data)
	outstruct := new(GetCspAccountWalletOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.WalletRegistry = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.WalletId = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.AccountIndex = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	return *outstruct, nil
}

// PackGetCspSourceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed856dd2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getCspSourceSettings(bytes32 _sourceId) view returns((uint16,uint8,uint128) _settings)
func (walletPayments *WalletPayments) PackGetCspSourceSettings(sourceId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("getCspSourceSettings", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetCspSourceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed856dd2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getCspSourceSettings(bytes32 _sourceId) view returns((uint16,uint8,uint128) _settings)
func (walletPayments *WalletPayments) TryPackGetCspSourceSettings(sourceId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("getCspSourceSettings", sourceId)
}

// UnpackGetCspSourceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xed856dd2.
//
// Solidity: function getCspSourceSettings(bytes32 _sourceId) view returns((uint16,uint8,uint128) _settings)
func (walletPayments *WalletPayments) UnpackGetCspSourceSettings(data []byte) (ICspProposalsCspSourceSettings, error) {
	out, err := walletPayments.abi.Unpack("getCspSourceSettings", data)
	if err != nil {
		return *new(ICspProposalsCspSourceSettings), err
	}
	out0 := *abi.ConvertType(out[0], new(ICspProposalsCspSourceSettings)).(*ICspProposalsCspSourceSettings)
	return out0, nil
}

// PackGetEffectiveSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed8b6f51.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getEffectiveSchedule(bytes32 _projectId, bytes32 _sourceId, bytes32 _accountHash) view returns(bytes _feeSchedule)
func (walletPayments *WalletPayments) PackGetEffectiveSchedule(projectId [32]byte, sourceId [32]byte, accountHash [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("getEffectiveSchedule", projectId, sourceId, accountHash)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetEffectiveSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xed8b6f51.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getEffectiveSchedule(bytes32 _projectId, bytes32 _sourceId, bytes32 _accountHash) view returns(bytes _feeSchedule)
func (walletPayments *WalletPayments) TryPackGetEffectiveSchedule(projectId [32]byte, sourceId [32]byte, accountHash [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("getEffectiveSchedule", projectId, sourceId, accountHash)
}

// UnpackGetEffectiveSchedule is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xed8b6f51.
//
// Solidity: function getEffectiveSchedule(bytes32 _projectId, bytes32 _sourceId, bytes32 _accountHash) view returns(bytes _feeSchedule)
func (walletPayments *WalletPayments) UnpackGetEffectiveSchedule(data []byte) ([]byte, error) {
	out, err := walletPayments.abi.Unpack("getEffectiveSchedule", data)
	if err != nil {
		return *new([]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	return out0, nil
}

// PackGetEligible is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc475a77a.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getEligible((bytes32,string) _account) view returns((uint64,uint32,uint64) _eligible)
func (walletPayments *WalletPayments) PackGetEligible(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getEligible", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetEligible is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc475a77a.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getEligible((bytes32,string) _account) view returns((uint64,uint32,uint64) _eligible)
func (walletPayments *WalletPayments) TryPackGetEligible(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getEligible", account)
}

// UnpackGetEligible is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc475a77a.
//
// Solidity: function getEligible((bytes32,string) _account) view returns((uint64,uint32,uint64) _eligible)
func (walletPayments *WalletPayments) UnpackGetEligible(data []byte) (ICspInstructionsEligible, error) {
	out, err := walletPayments.abi.Unpack("getEligible", data)
	if err != nil {
		return *new(ICspInstructionsEligible), err
	}
	out0 := *abi.ConvertType(out[0], new(ICspInstructionsEligible)).(*ICspInstructionsEligible)
	return out0, nil
}

// PackGetEscrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8bce253d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint256,uint64) _terms)
func (walletPayments *WalletPayments) PackGetEscrow(account WalletAccount, paymentId uint64) []byte {
	enc, err := walletPayments.abi.Pack("getEscrow", account, paymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetEscrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8bce253d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint256,uint64) _terms)
func (walletPayments *WalletPayments) TryPackGetEscrow(account WalletAccount, paymentId uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getEscrow", account, paymentId)
}

// UnpackGetEscrow is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x8bce253d.
//
// Solidity: function getEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint256,uint64) _terms)
func (walletPayments *WalletPayments) UnpackGetEscrow(data []byte) (IEscrowsEscrowTerms, error) {
	out, err := walletPayments.abi.Unpack("getEscrow", data)
	if err != nil {
		return *new(IEscrowsEscrowTerms), err
	}
	out0 := *abi.ConvertType(out[0], new(IEscrowsEscrowTerms)).(*IEscrowsEscrowTerms)
	return out0, nil
}

// PackGetFdc2ProofPolicy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3457d77.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getFdc2ProofPolicy() view returns((uint16,uint16) _policy)
func (walletPayments *WalletPayments) PackGetFdc2ProofPolicy() []byte {
	enc, err := walletPayments.abi.Pack("getFdc2ProofPolicy")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetFdc2ProofPolicy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa3457d77.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getFdc2ProofPolicy() view returns((uint16,uint16) _policy)
func (walletPayments *WalletPayments) TryPackGetFdc2ProofPolicy() ([]byte, error) {
	return walletPayments.abi.Pack("getFdc2ProofPolicy")
}

// UnpackGetFdc2ProofPolicy is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa3457d77.
//
// Solidity: function getFdc2ProofPolicy() view returns((uint16,uint16) _policy)
func (walletPayments *WalletPayments) UnpackGetFdc2ProofPolicy(data []byte) (IFdc2ProofPolicyProofPolicy, error) {
	out, err := walletPayments.abi.Unpack("getFdc2ProofPolicy", data)
	if err != nil {
		return *new(IFdc2ProofPolicyProofPolicy), err
	}
	out0 := *abi.ConvertType(out[0], new(IFdc2ProofPolicyProofPolicy)).(*IFdc2ProofPolicyProofPolicy)
	return out0, nil
}

// PackGetFeeScheduleConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16743f21.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getFeeScheduleConfig(bytes32 _sourceId) view returns((uint8,uint16) _config)
func (walletPayments *WalletPayments) PackGetFeeScheduleConfig(sourceId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("getFeeScheduleConfig", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetFeeScheduleConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16743f21.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getFeeScheduleConfig(bytes32 _sourceId) view returns((uint8,uint16) _config)
func (walletPayments *WalletPayments) TryPackGetFeeScheduleConfig(sourceId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("getFeeScheduleConfig", sourceId)
}

// UnpackGetFeeScheduleConfig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x16743f21.
//
// Solidity: function getFeeScheduleConfig(bytes32 _sourceId) view returns((uint8,uint16) _config)
func (walletPayments *WalletPayments) UnpackGetFeeScheduleConfig(data []byte) (IFeeSchedulesFeeScheduleConfig, error) {
	out, err := walletPayments.abi.Unpack("getFeeScheduleConfig", data)
	if err != nil {
		return *new(IFeeSchedulesFeeScheduleConfig), err
	}
	out0 := *abi.ConvertType(out[0], new(IFeeSchedulesFeeScheduleConfig)).(*IFeeSchedulesFeeScheduleConfig)
	return out0, nil
}

// PackGetFinalizedHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x113b0c23.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getFinalizedHash((bytes32,string) _account, uint64 _generation) view returns(bytes32 _packageHash)
func (walletPayments *WalletPayments) PackGetFinalizedHash(account WalletAccount, generation uint64) []byte {
	enc, err := walletPayments.abi.Pack("getFinalizedHash", account, generation)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetFinalizedHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x113b0c23.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getFinalizedHash((bytes32,string) _account, uint64 _generation) view returns(bytes32 _packageHash)
func (walletPayments *WalletPayments) TryPackGetFinalizedHash(account WalletAccount, generation uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getFinalizedHash", account, generation)
}

// UnpackGetFinalizedHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x113b0c23.
//
// Solidity: function getFinalizedHash((bytes32,string) _account, uint64 _generation) view returns(bytes32 _packageHash)
func (walletPayments *WalletPayments) UnpackGetFinalizedHash(data []byte) ([32]byte, error) {
	out, err := walletPayments.abi.Unpack("getFinalizedHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetInitialNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2bbba299.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getInitialNonce((bytes32,string) _account) view returns(uint64 _initialNonce)
func (walletPayments *WalletPayments) PackGetInitialNonce(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getInitialNonce", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetInitialNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2bbba299.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getInitialNonce((bytes32,string) _account) view returns(uint64 _initialNonce)
func (walletPayments *WalletPayments) TryPackGetInitialNonce(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getInitialNonce", account)
}

// UnpackGetInitialNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2bbba299.
//
// Solidity: function getInitialNonce((bytes32,string) _account) view returns(uint64 _initialNonce)
func (walletPayments *WalletPayments) UnpackGetInitialNonce(data []byte) (uint64, error) {
	out, err := walletPayments.abi.Unpack("getInitialNonce", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetInstruction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87e7e4ea.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getInstruction((bytes32,string) _account, uint64 _sequencePosition) view returns((uint8,uint32,uint64,uint64,uint64,bool) _instruction)
func (walletPayments *WalletPayments) PackGetInstruction(account WalletAccount, sequencePosition uint64) []byte {
	enc, err := walletPayments.abi.Pack("getInstruction", account, sequencePosition)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetInstruction is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87e7e4ea.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getInstruction((bytes32,string) _account, uint64 _sequencePosition) view returns((uint8,uint32,uint64,uint64,uint64,bool) _instruction)
func (walletPayments *WalletPayments) TryPackGetInstruction(account WalletAccount, sequencePosition uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getInstruction", account, sequencePosition)
}

// UnpackGetInstruction is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x87e7e4ea.
//
// Solidity: function getInstruction((bytes32,string) _account, uint64 _sequencePosition) view returns((uint8,uint32,uint64,uint64,uint64,bool) _instruction)
func (walletPayments *WalletPayments) UnpackGetInstruction(data []byte) (ICspInstructionsInstructionRecord, error) {
	out, err := walletPayments.abi.Unpack("getInstruction", data)
	if err != nil {
		return *new(ICspInstructionsInstructionRecord), err
	}
	out0 := *abi.ConvertType(out[0], new(ICspInstructionsInstructionRecord)).(*ICspInstructionsInstructionRecord)
	return out0, nil
}

// PackGetLaneLastSettled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c9c71f1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLaneLastSettled((bytes32,string) _account, uint32 _lane) view returns((bool,uint64) _use)
func (walletPayments *WalletPayments) PackGetLaneLastSettled(account WalletAccount, lane uint32) []byte {
	enc, err := walletPayments.abi.Pack("getLaneLastSettled", account, lane)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLaneLastSettled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c9c71f1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLaneLastSettled((bytes32,string) _account, uint32 _lane) view returns((bool,uint64) _use)
func (walletPayments *WalletPayments) TryPackGetLaneLastSettled(account WalletAccount, lane uint32) ([]byte, error) {
	return walletPayments.abi.Pack("getLaneLastSettled", account, lane)
}

// UnpackGetLaneLastSettled is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c9c71f1.
//
// Solidity: function getLaneLastSettled((bytes32,string) _account, uint32 _lane) view returns((bool,uint64) _use)
func (walletPayments *WalletPayments) UnpackGetLaneLastSettled(data []byte) (ICspInstructionsLaneUse, error) {
	out, err := walletPayments.abi.Unpack("getLaneLastSettled", data)
	if err != nil {
		return *new(ICspInstructionsLaneUse), err
	}
	out0 := *abi.ConvertType(out[0], new(ICspInstructionsLaneUse)).(*ICspInstructionsLaneUse)
	return out0, nil
}

// PackGetLatestAttemptNullifiedPaymentIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2cda059c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLatestAttemptNullifiedPaymentIds((bytes32,string) _account, uint64 _sequencePosition) view returns(uint64[] _nullifiedPaymentIds)
func (walletPayments *WalletPayments) PackGetLatestAttemptNullifiedPaymentIds(account WalletAccount, sequencePosition uint64) []byte {
	enc, err := walletPayments.abi.Pack("getLatestAttemptNullifiedPaymentIds", account, sequencePosition)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLatestAttemptNullifiedPaymentIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2cda059c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLatestAttemptNullifiedPaymentIds((bytes32,string) _account, uint64 _sequencePosition) view returns(uint64[] _nullifiedPaymentIds)
func (walletPayments *WalletPayments) TryPackGetLatestAttemptNullifiedPaymentIds(account WalletAccount, sequencePosition uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getLatestAttemptNullifiedPaymentIds", account, sequencePosition)
}

// UnpackGetLatestAttemptNullifiedPaymentIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2cda059c.
//
// Solidity: function getLatestAttemptNullifiedPaymentIds((bytes32,string) _account, uint64 _sequencePosition) view returns(uint64[] _nullifiedPaymentIds)
func (walletPayments *WalletPayments) UnpackGetLatestAttemptNullifiedPaymentIds(data []byte) ([]uint64, error) {
	out, err := walletPayments.abi.Unpack("getLatestAttemptNullifiedPaymentIds", data)
	if err != nil {
		return *new([]uint64), err
	}
	out0 := *abi.ConvertType(out[0], new([]uint64)).(*[]uint64)
	return out0, nil
}

// PackGetLeader is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9db9ddf1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getLeader((bytes32,string) _account, uint64 _generation) view returns((bool,uint64,uint32,uint32,uint64,uint32,uint8,bytes32,bytes32,address,uint64,uint32) _leader)
func (walletPayments *WalletPayments) PackGetLeader(account WalletAccount, generation uint64) []byte {
	enc, err := walletPayments.abi.Pack("getLeader", account, generation)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetLeader is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9db9ddf1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getLeader((bytes32,string) _account, uint64 _generation) view returns((bool,uint64,uint32,uint32,uint64,uint32,uint8,bytes32,bytes32,address,uint64,uint32) _leader)
func (walletPayments *WalletPayments) TryPackGetLeader(account WalletAccount, generation uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getLeader", account, generation)
}

// UnpackGetLeader is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9db9ddf1.
//
// Solidity: function getLeader((bytes32,string) _account, uint64 _generation) view returns((bool,uint64,uint32,uint32,uint64,uint32,uint8,bytes32,bytes32,address,uint64,uint32) _leader)
func (walletPayments *WalletPayments) UnpackGetLeader(data []byte) (ICspProposalsLeader, error) {
	out, err := walletPayments.abi.Unpack("getLeader", data)
	if err != nil {
		return *new(ICspProposalsLeader), err
	}
	out0 := *abi.ConvertType(out[0], new(ICspProposalsLeader)).(*ICspProposalsLeader)
	return out0, nil
}

// PackGetNetwork is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07e4b7e9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNetwork() view returns(uint8 _network)
func (walletPayments *WalletPayments) PackGetNetwork() []byte {
	enc, err := walletPayments.abi.Pack("getNetwork")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNetwork is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x07e4b7e9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNetwork() view returns(uint8 _network)
func (walletPayments *WalletPayments) TryPackGetNetwork() ([]byte, error) {
	return walletPayments.abi.Pack("getNetwork")
}

// UnpackGetNetwork is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x07e4b7e9.
//
// Solidity: function getNetwork() view returns(uint8 _network)
func (walletPayments *WalletPayments) UnpackGetNetwork(data []byte) (uint8, error) {
	out, err := walletPayments.abi.Unpack("getNetwork", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackGetNextPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfb49ac30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (walletPayments *WalletPayments) PackGetNextPaymentId(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getNextPaymentId", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNextPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfb49ac30.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (walletPayments *WalletPayments) TryPackGetNextPaymentId(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getNextPaymentId", account)
}

// UnpackGetNextPaymentId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfb49ac30.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (walletPayments *WalletPayments) UnpackGetNextPaymentId(data []byte) (uint64, error) {
	out, err := walletPayments.abi.Unpack("getNextPaymentId", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetNextSequencePosition is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7efb8ae6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNextSequencePosition((bytes32,string) _account) view returns(uint64 _nextSequencePosition)
func (walletPayments *WalletPayments) PackGetNextSequencePosition(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getNextSequencePosition", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNextSequencePosition is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7efb8ae6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNextSequencePosition((bytes32,string) _account) view returns(uint64 _nextSequencePosition)
func (walletPayments *WalletPayments) TryPackGetNextSequencePosition(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getNextSequencePosition", account)
}

// UnpackGetNextSequencePosition is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x7efb8ae6.
//
// Solidity: function getNextSequencePosition((bytes32,string) _account) view returns(uint64 _nextSequencePosition)
func (walletPayments *WalletPayments) UnpackGetNextSequencePosition(data []byte) (uint64, error) {
	out, err := walletPayments.abi.Unpack("getNextSequencePosition", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetNextUnconsumedPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf05af1b1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNextUnconsumedPaymentId((bytes32,string) _account) view returns(uint64 _nextUnconsumedPaymentId)
func (walletPayments *WalletPayments) PackGetNextUnconsumedPaymentId(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getNextUnconsumedPaymentId", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetNextUnconsumedPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf05af1b1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getNextUnconsumedPaymentId((bytes32,string) _account) view returns(uint64 _nextUnconsumedPaymentId)
func (walletPayments *WalletPayments) TryPackGetNextUnconsumedPaymentId(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getNextUnconsumedPaymentId", account)
}

// UnpackGetNextUnconsumedPaymentId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf05af1b1.
//
// Solidity: function getNextUnconsumedPaymentId((bytes32,string) _account) view returns(uint64 _nextUnconsumedPaymentId)
func (walletPayments *WalletPayments) UnpackGetNextUnconsumedPaymentId(data []byte) (uint64, error) {
	out, err := walletPayments.abi.Unpack("getNextUnconsumedPaymentId", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetOperationSequencePosition is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09e616d8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getOperationSequencePosition((bytes32,string) _account, uint64 _paymentId) view returns(bool _isOperation, uint64 _sequencePosition)
func (walletPayments *WalletPayments) PackGetOperationSequencePosition(account WalletAccount, paymentId uint64) []byte {
	enc, err := walletPayments.abi.Pack("getOperationSequencePosition", account, paymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetOperationSequencePosition is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x09e616d8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getOperationSequencePosition((bytes32,string) _account, uint64 _paymentId) view returns(bool _isOperation, uint64 _sequencePosition)
func (walletPayments *WalletPayments) TryPackGetOperationSequencePosition(account WalletAccount, paymentId uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getOperationSequencePosition", account, paymentId)
}

// GetOperationSequencePositionOutput serves as a container for the return parameters of contract
// method GetOperationSequencePosition.
type GetOperationSequencePositionOutput struct {
	IsOperation      bool
	SequencePosition uint64
}

// UnpackGetOperationSequencePosition is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x09e616d8.
//
// Solidity: function getOperationSequencePosition((bytes32,string) _account, uint64 _paymentId) view returns(bool _isOperation, uint64 _sequencePosition)
func (walletPayments *WalletPayments) UnpackGetOperationSequencePosition(data []byte) (GetOperationSequencePositionOutput, error) {
	out, err := walletPayments.abi.Unpack("getOperationSequencePosition", data)
	outstruct := new(GetOperationSequencePositionOutput)
	if err != nil {
		return *outstruct, err
	}
	outstruct.IsOperation = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.SequencePosition = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	return *outstruct, nil
}

// PackGetPaymentFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57abf78b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (walletPayments *WalletPayments) PackGetPaymentFee(account WalletAccount, opCommand [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("getPaymentFee", account, opCommand)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPaymentFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57abf78b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (walletPayments *WalletPayments) TryPackGetPaymentFee(account WalletAccount, opCommand [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("getPaymentFee", account, opCommand)
}

// UnpackGetPaymentFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (walletPayments *WalletPayments) UnpackGetPaymentFee(data []byte) (*big.Int, error) {
	out, err := walletPayments.abi.Unpack("getPaymentFee", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetPaymentHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0079dfe8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPaymentHash((bytes32,string) _account, uint64 _paymentId) view returns(bytes32 _paymentHash)
func (walletPayments *WalletPayments) PackGetPaymentHash(account WalletAccount, paymentId uint64) []byte {
	enc, err := walletPayments.abi.Pack("getPaymentHash", account, paymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPaymentHash is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0079dfe8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPaymentHash((bytes32,string) _account, uint64 _paymentId) view returns(bytes32 _paymentHash)
func (walletPayments *WalletPayments) TryPackGetPaymentHash(account WalletAccount, paymentId uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getPaymentHash", account, paymentId)
}

// UnpackGetPaymentHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0079dfe8.
//
// Solidity: function getPaymentHash((bytes32,string) _account, uint64 _paymentId) view returns(bytes32 _paymentHash)
func (walletPayments *WalletPayments) UnpackGetPaymentHash(data []byte) ([32]byte, error) {
	out, err := walletPayments.abi.Unpack("getPaymentHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetPendingResets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x807a0ade.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPendingResets((bytes32,string) _account) view returns((uint64,uint32)[] _pendingResets)
func (walletPayments *WalletPayments) PackGetPendingResets(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getPendingResets", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetPendingResets is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x807a0ade.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getPendingResets((bytes32,string) _account) view returns((uint64,uint32)[] _pendingResets)
func (walletPayments *WalletPayments) TryPackGetPendingResets(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getPendingResets", account)
}

// UnpackGetPendingResets is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x807a0ade.
//
// Solidity: function getPendingResets((bytes32,string) _account) view returns((uint64,uint32)[] _pendingResets)
func (walletPayments *WalletPayments) UnpackGetPendingResets(data []byte) ([]ICspInstructionsPendingReset, error) {
	out, err := walletPayments.abi.Unpack("getPendingResets", data)
	if err != nil {
		return *new([]ICspInstructionsPendingReset), err
	}
	out0 := *abi.ConvertType(out[0], new([]ICspInstructionsPendingReset)).(*[]ICspInstructionsPendingReset)
	return out0, nil
}

// PackGetProjectBtcWalletParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa466c45f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getProjectBtcWalletParams(bytes32 _projectId) view returns((uint32,uint8,uint8,uint8,uint8,uint8,uint8) _params)
func (walletPayments *WalletPayments) PackGetProjectBtcWalletParams(projectId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("getProjectBtcWalletParams", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetProjectBtcWalletParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa466c45f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getProjectBtcWalletParams(bytes32 _projectId) view returns((uint32,uint8,uint8,uint8,uint8,uint8,uint8) _params)
func (walletPayments *WalletPayments) TryPackGetProjectBtcWalletParams(projectId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("getProjectBtcWalletParams", projectId)
}

// UnpackGetProjectBtcWalletParams is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa466c45f.
//
// Solidity: function getProjectBtcWalletParams(bytes32 _projectId) view returns((uint32,uint8,uint8,uint8,uint8,uint8,uint8) _params)
func (walletPayments *WalletPayments) UnpackGetProjectBtcWalletParams(data []byte) (IBtcParamsBtcWalletParams, error) {
	out, err := walletPayments.abi.Unpack("getProjectBtcWalletParams", data)
	if err != nil {
		return *new(IBtcParamsBtcWalletParams), err
	}
	out0 := *abi.ConvertType(out[0], new(IBtcParamsBtcWalletParams)).(*IBtcParamsBtcWalletParams)
	return out0, nil
}

// PackGetProjectFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x30d1bb93.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) view returns((int16,uint16)[] _schedule)
func (walletPayments *WalletPayments) PackGetProjectFeeSchedule(projectId [32]byte, sourceId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("getProjectFeeSchedule", projectId, sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetProjectFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x30d1bb93.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) view returns((int16,uint16)[] _schedule)
func (walletPayments *WalletPayments) TryPackGetProjectFeeSchedule(projectId [32]byte, sourceId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("getProjectFeeSchedule", projectId, sourceId)
}

// UnpackGetProjectFeeSchedule is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x30d1bb93.
//
// Solidity: function getProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId) view returns((int16,uint16)[] _schedule)
func (walletPayments *WalletPayments) UnpackGetProjectFeeSchedule(data []byte) ([]IFeeSchedulesFeeSchedule, error) {
	out, err := walletPayments.abi.Unpack("getProjectFeeSchedule", data)
	if err != nil {
		return *new([]IFeeSchedulesFeeSchedule), err
	}
	out0 := *abi.ConvertType(out[0], new([]IFeeSchedulesFeeSchedule)).(*[]IFeeSchedulesFeeSchedule)
	return out0, nil
}

// PackGetProjectProposers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc07f32fe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getProjectProposers(bytes32 _projectId) view returns(address[] _proposers)
func (walletPayments *WalletPayments) PackGetProjectProposers(projectId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("getProjectProposers", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetProjectProposers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc07f32fe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getProjectProposers(bytes32 _projectId) view returns(address[] _proposers)
func (walletPayments *WalletPayments) TryPackGetProjectProposers(projectId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("getProjectProposers", projectId)
}

// UnpackGetProjectProposers is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xc07f32fe.
//
// Solidity: function getProjectProposers(bytes32 _projectId) view returns(address[] _proposers)
func (walletPayments *WalletPayments) UnpackGetProjectProposers(data []byte) ([]common.Address, error) {
	out, err := walletPayments.abi.Unpack("getProjectProposers", data)
	if err != nil {
		return *new([]common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	return out0, nil
}

// PackGetProposerUrl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6cb39f40.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getProposerUrl(address _proposer) view returns(string _url)
func (walletPayments *WalletPayments) PackGetProposerUrl(proposer common.Address) []byte {
	enc, err := walletPayments.abi.Pack("getProposerUrl", proposer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetProposerUrl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6cb39f40.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getProposerUrl(address _proposer) view returns(string _url)
func (walletPayments *WalletPayments) TryPackGetProposerUrl(proposer common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("getProposerUrl", proposer)
}

// UnpackGetProposerUrl is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6cb39f40.
//
// Solidity: function getProposerUrl(address _proposer) view returns(string _url)
func (walletPayments *WalletPayments) UnpackGetProposerUrl(data []byte) (string, error) {
	out, err := walletPayments.abi.Unpack("getProposerUrl", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackGetQueuedPayment is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x532de34f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getQueuedPayment((bytes32,string) _account, uint64 _paymentId) view returns((string,uint64,uint64,bytes32) _payment)
func (walletPayments *WalletPayments) PackGetQueuedPayment(account WalletAccount, paymentId uint64) []byte {
	enc, err := walletPayments.abi.Pack("getQueuedPayment", account, paymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetQueuedPayment is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x532de34f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getQueuedPayment((bytes32,string) _account, uint64 _paymentId) view returns((string,uint64,uint64,bytes32) _payment)
func (walletPayments *WalletPayments) TryPackGetQueuedPayment(account WalletAccount, paymentId uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getQueuedPayment", account, paymentId)
}

// UnpackGetQueuedPayment is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x532de34f.
//
// Solidity: function getQueuedPayment((bytes32,string) _account, uint64 _paymentId) view returns((string,uint64,uint64,bytes32) _payment)
func (walletPayments *WalletPayments) UnpackGetQueuedPayment(data []byte) (ICspQueueQueuedPayment, error) {
	out, err := walletPayments.abi.Unpack("getQueuedPayment", data)
	if err != nil {
		return *new(ICspQueueQueuedPayment), err
	}
	out0 := *abi.ConvertType(out[0], new(ICspQueueQueuedPayment)).(*ICspQueueQueuedPayment)
	return out0, nil
}

// PackGetRegisteredSourceIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec02c8e0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getRegisteredSourceIds() view returns(bytes32[] _sourceIds)
func (walletPayments *WalletPayments) PackGetRegisteredSourceIds() []byte {
	enc, err := walletPayments.abi.Pack("getRegisteredSourceIds")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetRegisteredSourceIds is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xec02c8e0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getRegisteredSourceIds() view returns(bytes32[] _sourceIds)
func (walletPayments *WalletPayments) TryPackGetRegisteredSourceIds() ([]byte, error) {
	return walletPayments.abi.Pack("getRegisteredSourceIds")
}

// UnpackGetRegisteredSourceIds is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xec02c8e0.
//
// Solidity: function getRegisteredSourceIds() view returns(bytes32[] _sourceIds)
func (walletPayments *WalletPayments) UnpackGetRegisteredSourceIds(data []byte) ([][32]byte, error) {
	out, err := walletPayments.abi.Unpack("getRegisteredSourceIds", data)
	if err != nil {
		return *new([][32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	return out0, nil
}

// PackGetSettledBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa30e1509.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSettledBatch((bytes32,string) _account, uint64 _paymentId) view returns((uint64,uint64,uint64) _batch)
func (walletPayments *WalletPayments) PackGetSettledBatch(account WalletAccount, paymentId uint64) []byte {
	enc, err := walletPayments.abi.Pack("getSettledBatch", account, paymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSettledBatch is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa30e1509.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSettledBatch((bytes32,string) _account, uint64 _paymentId) view returns((uint64,uint64,uint64) _batch)
func (walletPayments *WalletPayments) TryPackGetSettledBatch(account WalletAccount, paymentId uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getSettledBatch", account, paymentId)
}

// UnpackGetSettledBatch is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa30e1509.
//
// Solidity: function getSettledBatch((bytes32,string) _account, uint64 _paymentId) view returns((uint64,uint64,uint64) _batch)
func (walletPayments *WalletPayments) UnpackGetSettledBatch(data []byte) (ICspQueueSettledBatch, error) {
	out, err := walletPayments.abi.Unpack("getSettledBatch", data)
	if err != nil {
		return *new(ICspQueueSettledBatch), err
	}
	out0 := *abi.ConvertType(out[0], new(ICspQueueSettledBatch)).(*ICspQueueSettledBatch)
	return out0, nil
}

// PackGetSettledBatches is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2a29805c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSettledBatches((bytes32,string) _account) view returns((uint64,uint64,uint64)[] _batches)
func (walletPayments *WalletPayments) PackGetSettledBatches(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getSettledBatches", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSettledBatches is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2a29805c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSettledBatches((bytes32,string) _account) view returns((uint64,uint64,uint64)[] _batches)
func (walletPayments *WalletPayments) TryPackGetSettledBatches(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getSettledBatches", account)
}

// UnpackGetSettledBatches is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2a29805c.
//
// Solidity: function getSettledBatches((bytes32,string) _account) view returns((uint64,uint64,uint64)[] _batches)
func (walletPayments *WalletPayments) UnpackGetSettledBatches(data []byte) ([]ICspQueueSettledBatch, error) {
	out, err := walletPayments.abi.Unpack("getSettledBatches", data)
	if err != nil {
		return *new([]ICspQueueSettledBatch), err
	}
	out0 := *abi.ConvertType(out[0], new([]ICspQueueSettledBatch)).(*[]ICspQueueSettledBatch)
	return out0, nil
}

// PackGetSettlementCost is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52072400.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSettlementCost((bytes32,string) _account, uint64 _generation) view returns(uint256 _cost)
func (walletPayments *WalletPayments) PackGetSettlementCost(account WalletAccount, generation uint64) []byte {
	enc, err := walletPayments.abi.Pack("getSettlementCost", account, generation)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSettlementCost is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52072400.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSettlementCost((bytes32,string) _account, uint64 _generation) view returns(uint256 _cost)
func (walletPayments *WalletPayments) TryPackGetSettlementCost(account WalletAccount, generation uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getSettlementCost", account, generation)
}

// UnpackGetSettlementCost is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52072400.
//
// Solidity: function getSettlementCost((bytes32,string) _account, uint64 _generation) view returns(uint256 _cost)
func (walletPayments *WalletPayments) UnpackGetSettlementCost(data []byte) (*big.Int, error) {
	out, err := walletPayments.abi.Unpack("getSettlementCost", data)
	if err != nil {
		return new(big.Int), err
	}
	out0 := abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return out0, nil
}

// PackGetSourceConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a5b5cfd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getSourceConfig(bytes32 _sourceId) view returns((bytes32,bytes32,uint8,uint8,bool,bool,bool,bool) _source)
func (walletPayments *WalletPayments) PackGetSourceConfig(sourceId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("getSourceConfig", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetSourceConfig is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9a5b5cfd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getSourceConfig(bytes32 _sourceId) view returns((bytes32,bytes32,uint8,uint8,bool,bool,bool,bool) _source)
func (walletPayments *WalletPayments) TryPackGetSourceConfig(sourceId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("getSourceConfig", sourceId)
}

// UnpackGetSourceConfig is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x9a5b5cfd.
//
// Solidity: function getSourceConfig(bytes32 _sourceId) view returns((bytes32,bytes32,uint8,uint8,bool,bool,bool,bool) _source)
func (walletPayments *WalletPayments) UnpackGetSourceConfig(data []byte) (ISourceConfigSource, error) {
	out, err := walletPayments.abi.Unpack("getSourceConfig", data)
	if err != nil {
		return *new(ISourceConfigSource), err
	}
	out0 := *abi.ConvertType(out[0], new(ISourceConfigSource)).(*ISourceConfigSource)
	return out0, nil
}

// PackGetWalletAccounts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd0b152c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletAccounts(address _walletRegistry, bytes32 _walletId) view returns((bytes32,string)[] _walletAccounts)
func (walletPayments *WalletPayments) PackGetWalletAccounts(walletRegistry common.Address, walletId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("getWalletAccounts", walletRegistry, walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletAccounts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbd0b152c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletAccounts(address _walletRegistry, bytes32 _walletId) view returns((bytes32,string)[] _walletAccounts)
func (walletPayments *WalletPayments) TryPackGetWalletAccounts(walletRegistry common.Address, walletId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("getWalletAccounts", walletRegistry, walletId)
}

// UnpackGetWalletAccounts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbd0b152c.
//
// Solidity: function getWalletAccounts(address _walletRegistry, bytes32 _walletId) view returns((bytes32,string)[] _walletAccounts)
func (walletPayments *WalletPayments) UnpackGetWalletAccounts(data []byte) ([]WalletAccount, error) {
	out, err := walletPayments.abi.Unpack("getWalletAccounts", data)
	if err != nil {
		return *new([]WalletAccount), err
	}
	out0 := *abi.ConvertType(out[0], new([]WalletAccount)).(*[]WalletAccount)
	return out0, nil
}

// PackGetWalletId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5623b3f5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32 _walletId)
func (walletPayments *WalletPayments) PackGetWalletId(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getWalletId", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5623b3f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32 _walletId)
func (walletPayments *WalletPayments) TryPackGetWalletId(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getWalletId", account)
}

// UnpackGetWalletId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32 _walletId)
func (walletPayments *WalletPayments) UnpackGetWalletId(data []byte) ([32]byte, error) {
	out, err := walletPayments.abi.Unpack("getWalletId", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetWalletRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x454e7714.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletRegistry((bytes32,string) _account) view returns(address _walletRegistry)
func (walletPayments *WalletPayments) PackGetWalletRegistry(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getWalletRegistry", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x454e7714.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletRegistry((bytes32,string) _account) view returns(address _walletRegistry)
func (walletPayments *WalletPayments) TryPackGetWalletRegistry(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getWalletRegistry", account)
}

// UnpackGetWalletRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x454e7714.
//
// Solidity: function getWalletRegistry((bytes32,string) _account) view returns(address _walletRegistry)
func (walletPayments *WalletPayments) UnpackGetWalletRegistry(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("getWalletRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetWalletRegistryKind is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfcd97026.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletRegistryKind(address _registry) view returns(uint8 _kind)
func (walletPayments *WalletPayments) PackGetWalletRegistryKind(registry common.Address) []byte {
	enc, err := walletPayments.abi.Pack("getWalletRegistryKind", registry)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletRegistryKind is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfcd97026.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletRegistryKind(address _registry) view returns(uint8 _kind)
func (walletPayments *WalletPayments) TryPackGetWalletRegistryKind(registry common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("getWalletRegistryKind", registry)
}

// UnpackGetWalletRegistryKind is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfcd97026.
//
// Solidity: function getWalletRegistryKind(address _registry) view returns(uint8 _kind)
func (walletPayments *WalletPayments) UnpackGetWalletRegistryKind(data []byte) (uint8, error) {
	out, err := walletPayments.abi.Unpack("getWalletRegistryKind", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackGetXrpEscrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf38a8837.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getXrpEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint256,uint64,string,bool) _terms)
func (walletPayments *WalletPayments) PackGetXrpEscrow(account WalletAccount, paymentId uint64) []byte {
	enc, err := walletPayments.abi.Pack("getXrpEscrow", account, paymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetXrpEscrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf38a8837.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getXrpEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint256,uint64,string,bool) _terms)
func (walletPayments *WalletPayments) TryPackGetXrpEscrow(account WalletAccount, paymentId uint64) ([]byte, error) {
	return walletPayments.abi.Pack("getXrpEscrow", account, paymentId)
}

// UnpackGetXrpEscrow is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf38a8837.
//
// Solidity: function getXrpEscrow((bytes32,string) _account, uint64 _paymentId) view returns((bytes32,uint256,uint64,string,bool) _terms)
func (walletPayments *WalletPayments) UnpackGetXrpEscrow(data []byte) (IXrpEscrowsXrpEscrowTerms, error) {
	out, err := walletPayments.abi.Unpack("getXrpEscrow", data)
	if err != nil {
		return *new(IXrpEscrowsXrpEscrowTerms), err
	}
	out0 := *abi.ConvertType(out[0], new(IXrpEscrowsXrpEscrowTerms)).(*IXrpEscrowsXrpEscrowTerms)
	return out0, nil
}

// PackGetXrpEscrowProfile is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x235fb6a3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getXrpEscrowProfile((bytes32,string) _account) view returns(string _destination)
func (walletPayments *WalletPayments) PackGetXrpEscrowProfile(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("getXrpEscrowProfile", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetXrpEscrowProfile is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x235fb6a3.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getXrpEscrowProfile((bytes32,string) _account) view returns(string _destination)
func (walletPayments *WalletPayments) TryPackGetXrpEscrowProfile(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("getXrpEscrowProfile", account)
}

// UnpackGetXrpEscrowProfile is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x235fb6a3.
//
// Solidity: function getXrpEscrowProfile((bytes32,string) _account) view returns(string _destination)
func (walletPayments *WalletPayments) UnpackGetXrpEscrowProfile(data []byte) (string, error) {
	out, err := walletPayments.abi.Unpack("getXrpEscrowProfile", data)
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
func (walletPayments *WalletPayments) PackGovernance() []byte {
	enc, err := walletPayments.abi.Pack("governance")
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
func (walletPayments *WalletPayments) TryPackGovernance() ([]byte, error) {
	return walletPayments.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (walletPayments *WalletPayments) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("governance", data)
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
func (walletPayments *WalletPayments) PackGovernanceSettings() []byte {
	enc, err := walletPayments.abi.Pack("governanceSettings")
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
func (walletPayments *WalletPayments) TryPackGovernanceSettings() ([]byte, error) {
	return walletPayments.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (walletPayments *WalletPayments) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("governanceSettings", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackIsAllowedProposer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41cac86b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isAllowedProposer((bytes32,string) _account, address _proposer) view returns(bool)
func (walletPayments *WalletPayments) PackIsAllowedProposer(account WalletAccount, proposer common.Address) []byte {
	enc, err := walletPayments.abi.Pack("isAllowedProposer", account, proposer)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsAllowedProposer is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x41cac86b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isAllowedProposer((bytes32,string) _account, address _proposer) view returns(bool)
func (walletPayments *WalletPayments) TryPackIsAllowedProposer(account WalletAccount, proposer common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("isAllowedProposer", account, proposer)
}

// UnpackIsAllowedProposer is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x41cac86b.
//
// Solidity: function isAllowedProposer((bytes32,string) _account, address _proposer) view returns(bool)
func (walletPayments *WalletPayments) UnpackIsAllowedProposer(data []byte) (bool, error) {
	out, err := walletPayments.abi.Unpack("isAllowedProposer", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (walletPayments *WalletPayments) PackIsExecutor(address common.Address) []byte {
	enc, err := walletPayments.abi.Pack("isExecutor", address)
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
func (walletPayments *WalletPayments) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (walletPayments *WalletPayments) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := walletPayments.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsOperationPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1230187.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isOperationPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(bool _isOperation)
func (walletPayments *WalletPayments) PackIsOperationPaymentId(account WalletAccount, paymentId uint64) []byte {
	enc, err := walletPayments.abi.Pack("isOperationPaymentId", account, paymentId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsOperationPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd1230187.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isOperationPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(bool _isOperation)
func (walletPayments *WalletPayments) TryPackIsOperationPaymentId(account WalletAccount, paymentId uint64) ([]byte, error) {
	return walletPayments.abi.Pack("isOperationPaymentId", account, paymentId)
}

// UnpackIsOperationPaymentId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xd1230187.
//
// Solidity: function isOperationPaymentId((bytes32,string) _account, uint64 _paymentId) view returns(bool _isOperation)
func (walletPayments *WalletPayments) UnpackIsOperationPaymentId(data []byte) (bool, error) {
	out, err := walletPayments.abi.Unpack("isOperationPaymentId", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsSourceEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf6d2abbb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isSourceEnabled(bytes32 _sourceId) view returns(bool _enabled)
func (walletPayments *WalletPayments) PackIsSourceEnabled(sourceId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("isSourceEnabled", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsSourceEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf6d2abbb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isSourceEnabled(bytes32 _sourceId) view returns(bool _enabled)
func (walletPayments *WalletPayments) TryPackIsSourceEnabled(sourceId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("isSourceEnabled", sourceId)
}

// UnpackIsSourceEnabled is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf6d2abbb.
//
// Solidity: function isSourceEnabled(bytes32 _sourceId) view returns(bool _enabled)
func (walletPayments *WalletPayments) UnpackIsSourceEnabled(data []byte) (bool, error) {
	out, err := walletPayments.abi.Unpack("isSourceEnabled", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsSourceRegistered is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6848bfa4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isSourceRegistered(bytes32 _sourceId) view returns(bool _registered)
func (walletPayments *WalletPayments) PackIsSourceRegistered(sourceId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("isSourceRegistered", sourceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsSourceRegistered is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6848bfa4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isSourceRegistered(bytes32 _sourceId) view returns(bool _registered)
func (walletPayments *WalletPayments) TryPackIsSourceRegistered(sourceId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("isSourceRegistered", sourceId)
}

// UnpackIsSourceRegistered is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6848bfa4.
//
// Solidity: function isSourceRegistered(bytes32 _sourceId) view returns(bool _registered)
func (walletPayments *WalletPayments) UnpackIsSourceRegistered(data []byte) (bool, error) {
	out, err := walletPayments.abi.Unpack("isSourceRegistered", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackIsValidAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebecae8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isValidAddress(bytes32 _sourceId, string _address) view returns(bool _valid)
func (walletPayments *WalletPayments) PackIsValidAddress(sourceId [32]byte, address string) []byte {
	enc, err := walletPayments.abi.Pack("isValidAddress", sourceId, address)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackIsValidAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebecae8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function isValidAddress(bytes32 _sourceId, string _address) view returns(bool _valid)
func (walletPayments *WalletPayments) TryPackIsValidAddress(sourceId [32]byte, address string) ([]byte, error) {
	return walletPayments.abi.Pack("isValidAddress", sourceId, address)
}

// UnpackIsValidAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebecae8.
//
// Solidity: function isValidAddress(bytes32 _sourceId, string _address) view returns(bool _valid)
func (walletPayments *WalletPayments) UnpackIsValidAddress(data []byte) (bool, error) {
	out, err := walletPayments.abi.Unpack("isValidAddress", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackNullifyOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6a33d505.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function nullifyOperation((bytes32,string) _account, uint64 _paymentId, uint256 _maxFee, address _claimBackAddress) payable returns(uint32 _attempt)
func (walletPayments *WalletPayments) PackNullifyOperation(account WalletAccount, paymentId uint64, maxFee *big.Int, claimBackAddress common.Address) []byte {
	enc, err := walletPayments.abi.Pack("nullifyOperation", account, paymentId, maxFee, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackNullifyOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x6a33d505.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function nullifyOperation((bytes32,string) _account, uint64 _paymentId, uint256 _maxFee, address _claimBackAddress) payable returns(uint32 _attempt)
func (walletPayments *WalletPayments) TryPackNullifyOperation(account WalletAccount, paymentId uint64, maxFee *big.Int, claimBackAddress common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("nullifyOperation", account, paymentId, maxFee, claimBackAddress)
}

// UnpackNullifyOperation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x6a33d505.
//
// Solidity: function nullifyOperation((bytes32,string) _account, uint64 _paymentId, uint256 _maxFee, address _claimBackAddress) payable returns(uint32 _attempt)
func (walletPayments *WalletPayments) UnpackNullifyOperation(data []byte) (uint32, error) {
	out, err := walletPayments.abi.Unpack("nullifyOperation", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackPay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x009ce938.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (walletPayments *WalletPayments) PackPay(account WalletAccount, paymentInstruction PaymentInstruction, claimBackAddress common.Address) []byte {
	enc, err := walletPayments.abi.Pack("pay", account, paymentInstruction, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x009ce938.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (walletPayments *WalletPayments) TryPackPay(account WalletAccount, paymentInstruction PaymentInstruction, claimBackAddress common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("pay", account, paymentInstruction, claimBackAddress)
}

// UnpackPay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (walletPayments *WalletPayments) UnpackPay(data []byte) (uint64, error) {
	out, err := walletPayments.abi.Unpack("pay", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (walletPayments *WalletPayments) PackProductionMode() []byte {
	enc, err := walletPayments.abi.Pack("productionMode")
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
func (walletPayments *WalletPayments) TryPackProductionMode() ([]byte, error) {
	return walletPayments.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (walletPayments *WalletPayments) UnpackProductionMode(data []byte) (bool, error) {
	out, err := walletPayments.abi.Unpack("productionMode", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackReclaimEscrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa6bf742f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reclaimEscrow((bytes32,string) _account, uint64 _createPaymentId, uint8 _mode, uint256 _maxFee) returns(uint64 _paymentId)
func (walletPayments *WalletPayments) PackReclaimEscrow(account WalletAccount, createPaymentId uint64, mode uint8, maxFee *big.Int) []byte {
	enc, err := walletPayments.abi.Pack("reclaimEscrow", account, createPaymentId, mode, maxFee)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReclaimEscrow is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa6bf742f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reclaimEscrow((bytes32,string) _account, uint64 _createPaymentId, uint8 _mode, uint256 _maxFee) returns(uint64 _paymentId)
func (walletPayments *WalletPayments) TryPackReclaimEscrow(account WalletAccount, createPaymentId uint64, mode uint8, maxFee *big.Int) ([]byte, error) {
	return walletPayments.abi.Pack("reclaimEscrow", account, createPaymentId, mode, maxFee)
}

// UnpackReclaimEscrow is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa6bf742f.
//
// Solidity: function reclaimEscrow((bytes32,string) _account, uint64 _createPaymentId, uint8 _mode, uint256 _maxFee) returns(uint64 _paymentId)
func (walletPayments *WalletPayments) UnpackReclaimEscrow(data []byte) (uint64, error) {
	out, err := walletPayments.abi.Unpack("reclaimEscrow", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackRegisterSources is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8c90fb6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function registerSources((bytes32,bytes32,bytes32,uint8,uint8)[] _registrations) returns()
func (walletPayments *WalletPayments) PackRegisterSources(registrations []ISourceConfigSourceRegistration) []byte {
	enc, err := walletPayments.abi.Pack("registerSources", registrations)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRegisterSources is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf8c90fb6.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function registerSources((bytes32,bytes32,bytes32,uint8,uint8)[] _registrations) returns()
func (walletPayments *WalletPayments) TryPackRegisterSources(registrations []ISourceConfigSourceRegistration) ([]byte, error) {
	return walletPayments.abi.Pack("registerSources", registrations)
}

// PackReissue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5cc0a260.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _paymentId, (string,bytes,uint256,uint256,bytes32)[] _retainedInstructions, uint64[] _nullifiedPaymentIds, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns(bool _finalized)
func (walletPayments *WalletPayments) PackReissue(account WalletAccount, paymentId uint64, retainedInstructions []PaymentInstruction, nullifiedPaymentIds []uint64, reissueFeeParams ReissueFeeParams, claimBackAddress common.Address) []byte {
	enc, err := walletPayments.abi.Pack("reissue", account, paymentId, retainedInstructions, nullifiedPaymentIds, reissueFeeParams, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReissue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5cc0a260.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _paymentId, (string,bytes,uint256,uint256,bytes32)[] _retainedInstructions, uint64[] _nullifiedPaymentIds, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns(bool _finalized)
func (walletPayments *WalletPayments) TryPackReissue(account WalletAccount, paymentId uint64, retainedInstructions []PaymentInstruction, nullifiedPaymentIds []uint64, reissueFeeParams ReissueFeeParams, claimBackAddress common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("reissue", account, paymentId, retainedInstructions, nullifiedPaymentIds, reissueFeeParams, claimBackAddress)
}

// UnpackReissue is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5cc0a260.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _paymentId, (string,bytes,uint256,uint256,bytes32)[] _retainedInstructions, uint64[] _nullifiedPaymentIds, (uint256[],int16[][],uint16[]) _reissueFeeParams, address _claimBackAddress) payable returns(bool _finalized)
func (walletPayments *WalletPayments) UnpackReissue(data []byte) (bool, error) {
	out, err := walletPayments.abi.Unpack("reissue", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackReissueOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf10cc52c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reissueOperation((bytes32,string) _account, uint64 _paymentId, uint256 _maxFee, address _claimBackAddress) payable returns(uint32 _attempt)
func (walletPayments *WalletPayments) PackReissueOperation(account WalletAccount, paymentId uint64, maxFee *big.Int, claimBackAddress common.Address) []byte {
	enc, err := walletPayments.abi.Pack("reissueOperation", account, paymentId, maxFee, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReissueOperation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf10cc52c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reissueOperation((bytes32,string) _account, uint64 _paymentId, uint256 _maxFee, address _claimBackAddress) payable returns(uint32 _attempt)
func (walletPayments *WalletPayments) TryPackReissueOperation(account WalletAccount, paymentId uint64, maxFee *big.Int, claimBackAddress common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("reissueOperation", account, paymentId, maxFee, claimBackAddress)
}

// UnpackReissueOperation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf10cc52c.
//
// Solidity: function reissueOperation((bytes32,string) _account, uint64 _paymentId, uint256 _maxFee, address _claimBackAddress) payable returns(uint32 _attempt)
func (walletPayments *WalletPayments) UnpackReissueOperation(data []byte) (uint32, error) {
	out, err := walletPayments.abi.Unpack("reissueOperation", data)
	if err != nil {
		return *new(uint32), err
	}
	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	return out0, nil
}

// PackRemoveAccountBtcWalletParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c54c7cd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeAccountBtcWalletParams((bytes32,string) _account) returns()
func (walletPayments *WalletPayments) PackRemoveAccountBtcWalletParams(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("removeAccountBtcWalletParams", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveAccountBtcWalletParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1c54c7cd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeAccountBtcWalletParams((bytes32,string) _account) returns()
func (walletPayments *WalletPayments) TryPackRemoveAccountBtcWalletParams(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("removeAccountBtcWalletParams", account)
}

// PackRemoveAccountProposers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x379ae241.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeAccountProposers((bytes32,string) _account) returns()
func (walletPayments *WalletPayments) PackRemoveAccountProposers(account WalletAccount) []byte {
	enc, err := walletPayments.abi.Pack("removeAccountProposers", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveAccountProposers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x379ae241.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeAccountProposers((bytes32,string) _account) returns()
func (walletPayments *WalletPayments) TryPackRemoveAccountProposers(account WalletAccount) ([]byte, error) {
	return walletPayments.abi.Pack("removeAccountProposers", account)
}

// PackRemoveProjectBtcWalletParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x65250799.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeProjectBtcWalletParams(bytes32 _projectId) returns()
func (walletPayments *WalletPayments) PackRemoveProjectBtcWalletParams(projectId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("removeProjectBtcWalletParams", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveProjectBtcWalletParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x65250799.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeProjectBtcWalletParams(bytes32 _projectId) returns()
func (walletPayments *WalletPayments) TryPackRemoveProjectBtcWalletParams(projectId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("removeProjectBtcWalletParams", projectId)
}

// PackRemoveProjectProposers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3304f19c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function removeProjectProposers(bytes32 _projectId) returns()
func (walletPayments *WalletPayments) PackRemoveProjectProposers(projectId [32]byte) []byte {
	enc, err := walletPayments.abi.Pack("removeProjectProposers", projectId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRemoveProjectProposers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3304f19c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function removeProjectProposers(bytes32 _projectId) returns()
func (walletPayments *WalletPayments) TryPackRemoveProjectProposers(projectId [32]byte) ([]byte, error) {
	return walletPayments.abi.Pack("removeProjectProposers", projectId)
}

// PackRequestBtcAccountConfiguredAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x58a3b1da.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function requestBtcAccountConfiguredAttestation(address _walletRegistry, bytes32 _walletId, bytes32 _sourceId, uint32 _accountIndex, (bytes32,uint32)[] _anchors, address[] _teeIds, address _proofOwner, address _claimBackAddress) payable returns()
func (walletPayments *WalletPayments) PackRequestBtcAccountConfiguredAttestation(walletRegistry common.Address, walletId [32]byte, sourceId [32]byte, accountIndex uint32, anchors []IBtcAccountConfiguredAnchor, teeIds []common.Address, proofOwner common.Address, claimBackAddress common.Address) []byte {
	enc, err := walletPayments.abi.Pack("requestBtcAccountConfiguredAttestation", walletRegistry, walletId, sourceId, accountIndex, anchors, teeIds, proofOwner, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRequestBtcAccountConfiguredAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x58a3b1da.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function requestBtcAccountConfiguredAttestation(address _walletRegistry, bytes32 _walletId, bytes32 _sourceId, uint32 _accountIndex, (bytes32,uint32)[] _anchors, address[] _teeIds, address _proofOwner, address _claimBackAddress) payable returns()
func (walletPayments *WalletPayments) TryPackRequestBtcAccountConfiguredAttestation(walletRegistry common.Address, walletId [32]byte, sourceId [32]byte, accountIndex uint32, anchors []IBtcAccountConfiguredAnchor, teeIds []common.Address, proofOwner common.Address, claimBackAddress common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("requestBtcAccountConfiguredAttestation", walletRegistry, walletId, sourceId, accountIndex, anchors, teeIds, proofOwner, claimBackAddress)
}

// PackRequestNativeNonceAccountConfiguredAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc262cf6b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function requestNativeNonceAccountConfiguredAttestation(address _walletRegistry, bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address[] _teeIds, address _proofOwner, address _claimBackAddress) payable returns()
func (walletPayments *WalletPayments) PackRequestNativeNonceAccountConfiguredAttestation(walletRegistry common.Address, walletId [32]byte, sourceId [32]byte, accountAddress string, teeIds []common.Address, proofOwner common.Address, claimBackAddress common.Address) []byte {
	enc, err := walletPayments.abi.Pack("requestNativeNonceAccountConfiguredAttestation", walletRegistry, walletId, sourceId, accountAddress, teeIds, proofOwner, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRequestNativeNonceAccountConfiguredAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc262cf6b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function requestNativeNonceAccountConfiguredAttestation(address _walletRegistry, bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address[] _teeIds, address _proofOwner, address _claimBackAddress) payable returns()
func (walletPayments *WalletPayments) TryPackRequestNativeNonceAccountConfiguredAttestation(walletRegistry common.Address, walletId [32]byte, sourceId [32]byte, accountAddress string, teeIds []common.Address, proofOwner common.Address, claimBackAddress common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("requestNativeNonceAccountConfiguredAttestation", walletRegistry, walletId, sourceId, accountAddress, teeIds, proofOwner, claimBackAddress)
}

// PackRequestSigning is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc26145d2.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function requestSigning((bytes32,string) _account, uint64 _sequencePosition) payable returns()
func (walletPayments *WalletPayments) PackRequestSigning(account WalletAccount, sequencePosition uint64) []byte {
	enc, err := walletPayments.abi.Pack("requestSigning", account, sequencePosition)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRequestSigning is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc26145d2.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function requestSigning((bytes32,string) _account, uint64 _sequencePosition) payable returns()
func (walletPayments *WalletPayments) TryPackRequestSigning(account WalletAccount, sequencePosition uint64) ([]byte, error) {
	return walletPayments.abi.Pack("requestSigning", account, sequencePosition)
}

// PackRewardManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f4ef8a6.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function rewardManager() view returns(address)
func (walletPayments *WalletPayments) PackRewardManager() []byte {
	enc, err := walletPayments.abi.Pack("rewardManager")
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
func (walletPayments *WalletPayments) TryPackRewardManager() ([]byte, error) {
	return walletPayments.abi.Pack("rewardManager")
}

// UnpackRewardManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (walletPayments *WalletPayments) UnpackRewardManager(data []byte) (common.Address, error) {
	out, err := walletPayments.abi.Unpack("rewardManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackSetAccountBtcWalletParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4a5b8992.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setAccountBtcWalletParams((bytes32,string) _account, (uint32,uint8,uint8,uint8,uint8,uint8,uint8) _params) returns()
func (walletPayments *WalletPayments) PackSetAccountBtcWalletParams(account WalletAccount, params IBtcParamsBtcWalletParams) []byte {
	enc, err := walletPayments.abi.Pack("setAccountBtcWalletParams", account, params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetAccountBtcWalletParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4a5b8992.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setAccountBtcWalletParams((bytes32,string) _account, (uint32,uint8,uint8,uint8,uint8,uint8,uint8) _params) returns()
func (walletPayments *WalletPayments) TryPackSetAccountBtcWalletParams(account WalletAccount, params IBtcParamsBtcWalletParams) ([]byte, error) {
	return walletPayments.abi.Pack("setAccountBtcWalletParams", account, params)
}

// PackSetAccountFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32ab4577.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setAccountFeeSchedule((bytes32,string) _account, (int16,uint16)[] _schedule) returns()
func (walletPayments *WalletPayments) PackSetAccountFeeSchedule(account WalletAccount, schedule []IFeeSchedulesFeeSchedule) []byte {
	enc, err := walletPayments.abi.Pack("setAccountFeeSchedule", account, schedule)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetAccountFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x32ab4577.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setAccountFeeSchedule((bytes32,string) _account, (int16,uint16)[] _schedule) returns()
func (walletPayments *WalletPayments) TryPackSetAccountFeeSchedule(account WalletAccount, schedule []IFeeSchedulesFeeSchedule) ([]byte, error) {
	return walletPayments.abi.Pack("setAccountFeeSchedule", account, schedule)
}

// PackSetAccountProposers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23f3826f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setAccountProposers((bytes32,string) _account, address[] _proposers) returns()
func (walletPayments *WalletPayments) PackSetAccountProposers(account WalletAccount, proposers []common.Address) []byte {
	enc, err := walletPayments.abi.Pack("setAccountProposers", account, proposers)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetAccountProposers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x23f3826f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setAccountProposers((bytes32,string) _account, address[] _proposers) returns()
func (walletPayments *WalletPayments) TryPackSetAccountProposers(account WalletAccount, proposers []common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("setAccountProposers", account, proposers)
}

// PackSetBtcConsensusParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7497207f.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setBtcConsensusParams(bytes32 _sourceId, (uint8,uint32,uint16,uint32,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint16,uint16,uint16,uint32) _params) returns()
func (walletPayments *WalletPayments) PackSetBtcConsensusParams(sourceId [32]byte, params IBtcParamsBtcConsensusParams) []byte {
	enc, err := walletPayments.abi.Pack("setBtcConsensusParams", sourceId, params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetBtcConsensusParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x7497207f.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setBtcConsensusParams(bytes32 _sourceId, (uint8,uint32,uint16,uint32,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint16,uint16,uint16,uint32) _params) returns()
func (walletPayments *WalletPayments) TryPackSetBtcConsensusParams(sourceId [32]byte, params IBtcParamsBtcConsensusParams) ([]byte, error) {
	return walletPayments.abi.Pack("setBtcConsensusParams", sourceId, params)
}

// PackSetBtcEscrowProfile is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0fd69b17.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setBtcEscrowProfile((bytes32,string) _account, bytes _counterpartyPubKey) returns()
func (walletPayments *WalletPayments) PackSetBtcEscrowProfile(account WalletAccount, counterpartyPubKey []byte) []byte {
	enc, err := walletPayments.abi.Pack("setBtcEscrowProfile", account, counterpartyPubKey)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetBtcEscrowProfile is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0fd69b17.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setBtcEscrowProfile((bytes32,string) _account, bytes _counterpartyPubKey) returns()
func (walletPayments *WalletPayments) TryPackSetBtcEscrowProfile(account WalletAccount, counterpartyPubKey []byte) ([]byte, error) {
	return walletPayments.abi.Pack("setBtcEscrowProfile", account, counterpartyPubKey)
}

// PackSetCspSourceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x067dc6f1.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setCspSourceSettings(bytes32 _sourceId, (uint16,uint8,uint128) _settings) returns()
func (walletPayments *WalletPayments) PackSetCspSourceSettings(sourceId [32]byte, settings ICspProposalsCspSourceSettings) []byte {
	enc, err := walletPayments.abi.Pack("setCspSourceSettings", sourceId, settings)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetCspSourceSettings is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x067dc6f1.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setCspSourceSettings(bytes32 _sourceId, (uint16,uint8,uint128) _settings) returns()
func (walletPayments *WalletPayments) TryPackSetCspSourceSettings(sourceId [32]byte, settings ICspProposalsCspSourceSettings) ([]byte, error) {
	return walletPayments.abi.Pack("setCspSourceSettings", sourceId, settings)
}

// PackSetFdc2ProofPolicy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd7b3c3cd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFdc2ProofPolicy((uint16,uint16) _policy) returns()
func (walletPayments *WalletPayments) PackSetFdc2ProofPolicy(policy IFdc2ProofPolicyProofPolicy) []byte {
	enc, err := walletPayments.abi.Pack("setFdc2ProofPolicy", policy)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFdc2ProofPolicy is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd7b3c3cd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFdc2ProofPolicy((uint16,uint16) _policy) returns()
func (walletPayments *WalletPayments) TryPackSetFdc2ProofPolicy(policy IFdc2ProofPolicyProofPolicy) ([]byte, error) {
	return walletPayments.abi.Pack("setFdc2ProofPolicy", policy)
}

// PackSetFeeScheduleConfigs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b7f8f29.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setFeeScheduleConfigs((uint16,uint8,bytes32)[] _configs) returns()
func (walletPayments *WalletPayments) PackSetFeeScheduleConfigs(configs []IFeeSchedulesFeeScheduleConfigInput) []byte {
	enc, err := walletPayments.abi.Pack("setFeeScheduleConfigs", configs)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetFeeScheduleConfigs is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x9b7f8f29.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setFeeScheduleConfigs((uint16,uint8,bytes32)[] _configs) returns()
func (walletPayments *WalletPayments) TryPackSetFeeScheduleConfigs(configs []IFeeSchedulesFeeScheduleConfigInput) ([]byte, error) {
	return walletPayments.abi.Pack("setFeeScheduleConfigs", configs)
}

// PackSetProjectBtcWalletParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87915b8d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setProjectBtcWalletParams(bytes32 _projectId, (uint32,uint8,uint8,uint8,uint8,uint8,uint8) _params) returns()
func (walletPayments *WalletPayments) PackSetProjectBtcWalletParams(projectId [32]byte, params IBtcParamsBtcWalletParams) []byte {
	enc, err := walletPayments.abi.Pack("setProjectBtcWalletParams", projectId, params)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetProjectBtcWalletParams is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x87915b8d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setProjectBtcWalletParams(bytes32 _projectId, (uint32,uint8,uint8,uint8,uint8,uint8,uint8) _params) returns()
func (walletPayments *WalletPayments) TryPackSetProjectBtcWalletParams(projectId [32]byte, params IBtcParamsBtcWalletParams) ([]byte, error) {
	return walletPayments.abi.Pack("setProjectBtcWalletParams", projectId, params)
}

// PackSetProjectFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe27d19bd.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId, (int16,uint16)[] _schedule) returns()
func (walletPayments *WalletPayments) PackSetProjectFeeSchedule(projectId [32]byte, sourceId [32]byte, schedule []IFeeSchedulesFeeSchedule) []byte {
	enc, err := walletPayments.abi.Pack("setProjectFeeSchedule", projectId, sourceId, schedule)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetProjectFeeSchedule is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe27d19bd.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setProjectFeeSchedule(bytes32 _projectId, bytes32 _sourceId, (int16,uint16)[] _schedule) returns()
func (walletPayments *WalletPayments) TryPackSetProjectFeeSchedule(projectId [32]byte, sourceId [32]byte, schedule []IFeeSchedulesFeeSchedule) ([]byte, error) {
	return walletPayments.abi.Pack("setProjectFeeSchedule", projectId, sourceId, schedule)
}

// PackSetProjectProposers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe902206d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setProjectProposers(bytes32 _projectId, address[] _proposers) returns()
func (walletPayments *WalletPayments) PackSetProjectProposers(projectId [32]byte, proposers []common.Address) []byte {
	enc, err := walletPayments.abi.Pack("setProjectProposers", projectId, proposers)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetProjectProposers is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe902206d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setProjectProposers(bytes32 _projectId, address[] _proposers) returns()
func (walletPayments *WalletPayments) TryPackSetProjectProposers(projectId [32]byte, proposers []common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("setProjectProposers", projectId, proposers)
}

// PackSetProposerUrl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1808899.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setProposerUrl(string _url) returns()
func (walletPayments *WalletPayments) PackSetProposerUrl(url string) []byte {
	enc, err := walletPayments.abi.Pack("setProposerUrl", url)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetProposerUrl is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa1808899.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setProposerUrl(string _url) returns()
func (walletPayments *WalletPayments) TryPackSetProposerUrl(url string) ([]byte, error) {
	return walletPayments.abi.Pack("setProposerUrl", url)
}

// PackSetSourceEscrowsEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb6411e9.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSourceEscrowsEnabled(bytes32 _sourceId, bool _custodianWallets, bool _teeWallets) returns()
func (walletPayments *WalletPayments) PackSetSourceEscrowsEnabled(sourceId [32]byte, custodianWallets bool, teeWallets bool) []byte {
	enc, err := walletPayments.abi.Pack("setSourceEscrowsEnabled", sourceId, custodianWallets, teeWallets)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSourceEscrowsEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdb6411e9.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSourceEscrowsEnabled(bytes32 _sourceId, bool _custodianWallets, bool _teeWallets) returns()
func (walletPayments *WalletPayments) TryPackSetSourceEscrowsEnabled(sourceId [32]byte, custodianWallets bool, teeWallets bool) ([]byte, error) {
	return walletPayments.abi.Pack("setSourceEscrowsEnabled", sourceId, custodianWallets, teeWallets)
}

// PackSetSourcesEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62825a23.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setSourcesEnabled(bytes32[] _sourceIds, bool _enabled) returns()
func (walletPayments *WalletPayments) PackSetSourcesEnabled(sourceIds [][32]byte, enabled bool) []byte {
	enc, err := walletPayments.abi.Pack("setSourcesEnabled", sourceIds, enabled)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetSourcesEnabled is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x62825a23.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setSourcesEnabled(bytes32[] _sourceIds, bool _enabled) returns()
func (walletPayments *WalletPayments) TryPackSetSourcesEnabled(sourceIds [][32]byte, enabled bool) ([]byte, error) {
	return walletPayments.abi.Pack("setSourcesEnabled", sourceIds, enabled)
}

// PackSetXrpEscrowProfile is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2c0c050.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function setXrpEscrowProfile((bytes32,string) _account, string _destination) returns()
func (walletPayments *WalletPayments) PackSetXrpEscrowProfile(account WalletAccount, destination string) []byte {
	enc, err := walletPayments.abi.Pack("setXrpEscrowProfile", account, destination)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSetXrpEscrowProfile is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xd2c0c050.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function setXrpEscrowProfile((bytes32,string) _account, string _destination) returns()
func (walletPayments *WalletPayments) TryPackSetXrpEscrowProfile(account WalletAccount, destination string) ([]byte, error) {
	return walletPayments.abi.Pack("setXrpEscrowProfile", account, destination)
}

// PackSettle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5eb3083.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function settle((bytes32,string) _account, uint64 _sequencePosition) payable returns()
func (walletPayments *WalletPayments) PackSettle(account WalletAccount, sequencePosition uint64) []byte {
	enc, err := walletPayments.abi.Pack("settle", account, sequencePosition)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSettle is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5eb3083.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function settle((bytes32,string) _account, uint64 _sequencePosition) payable returns()
func (walletPayments *WalletPayments) TryPackSettle(account WalletAccount, sequencePosition uint64) ([]byte, error) {
	return walletPayments.abi.Pack("settle", account, sequencePosition)
}

// PackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (walletPayments *WalletPayments) PackSupportsInterface(interfaceId [4]byte) []byte {
	enc, err := walletPayments.abi.Pack("supportsInterface", interfaceId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackSupportsInterface is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x01ffc9a7.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (walletPayments *WalletPayments) TryPackSupportsInterface(interfaceId [4]byte) ([]byte, error) {
	return walletPayments.abi.Pack("supportsInterface", interfaceId)
}

// UnpackSupportsInterface is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (walletPayments *WalletPayments) UnpackSupportsInterface(data []byte) (bool, error) {
	out, err := walletPayments.abi.Unpack("supportsInterface", data)
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
func (walletPayments *WalletPayments) PackSwitchToProductionMode() []byte {
	enc, err := walletPayments.abi.Pack("switchToProductionMode")
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
func (walletPayments *WalletPayments) TryPackSwitchToProductionMode() ([]byte, error) {
	return walletPayments.abi.Pack("switchToProductionMode")
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (walletPayments *WalletPayments) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := walletPayments.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (walletPayments *WalletPayments) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return walletPayments.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackValidateAndEncodeSchedules is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3d699135.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function validateAndEncodeSchedules(bytes32 _sourceId, int16[][] _factorsBIPSPerPayment, uint16[] _delaysSeconds) view returns(bytes[] _encodedPerPayment)
func (walletPayments *WalletPayments) PackValidateAndEncodeSchedules(sourceId [32]byte, factorsBIPSPerPayment [][]int16, delaysSeconds []uint16) []byte {
	enc, err := walletPayments.abi.Pack("validateAndEncodeSchedules", sourceId, factorsBIPSPerPayment, delaysSeconds)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackValidateAndEncodeSchedules is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3d699135.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function validateAndEncodeSchedules(bytes32 _sourceId, int16[][] _factorsBIPSPerPayment, uint16[] _delaysSeconds) view returns(bytes[] _encodedPerPayment)
func (walletPayments *WalletPayments) TryPackValidateAndEncodeSchedules(sourceId [32]byte, factorsBIPSPerPayment [][]int16, delaysSeconds []uint16) ([]byte, error) {
	return walletPayments.abi.Pack("validateAndEncodeSchedules", sourceId, factorsBIPSPerPayment, delaysSeconds)
}

// UnpackValidateAndEncodeSchedules is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3d699135.
//
// Solidity: function validateAndEncodeSchedules(bytes32 _sourceId, int16[][] _factorsBIPSPerPayment, uint16[] _delaysSeconds) view returns(bytes[] _encodedPerPayment)
func (walletPayments *WalletPayments) UnpackValidateAndEncodeSchedules(data []byte) ([][]byte, error) {
	out, err := walletPayments.abi.Unpack("validateAndEncodeSchedules", data)
	if err != nil {
		return *new([][]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)
	return out0, nil
}

// PackVerifyWalletOperationStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1ccfd841.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function verifyWalletOperationStatus(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,bytes32),(uint8,uint64,uint32,bool,bytes32,uint256,uint64,uint64,bool,bytes32,uint32,bool,uint8,bytes32,uint256)) _proof) returns(bool _valid)
func (walletPayments *WalletPayments) PackVerifyWalletOperationStatus(proof IWalletOperationStatusProof) []byte {
	enc, err := walletPayments.abi.Pack("verifyWalletOperationStatus", proof)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVerifyWalletOperationStatus is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1ccfd841.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function verifyWalletOperationStatus(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,bytes32),(uint8,uint64,uint32,bool,bytes32,uint256,uint64,uint64,bool,bytes32,uint32,bool,uint8,bytes32,uint256)) _proof) returns(bool _valid)
func (walletPayments *WalletPayments) TryPackVerifyWalletOperationStatus(proof IWalletOperationStatusProof) ([]byte, error) {
	return walletPayments.abi.Pack("verifyWalletOperationStatus", proof)
}

// UnpackVerifyWalletOperationStatus is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1ccfd841.
//
// Solidity: function verifyWalletOperationStatus(((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(bytes32,string,uint64,bytes32),(uint8,uint64,uint32,bool,bytes32,uint256,uint64,uint64,bool,bytes32,uint32,bool,uint8,bytes32,uint256)) _proof) returns(bool _valid)
func (walletPayments *WalletPayments) UnpackVerifyWalletOperationStatus(data []byte) (bool, error) {
	out, err := walletPayments.abi.Unpack("verifyWalletOperationStatus", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// WalletPaymentsAccountBtcWalletParamsSet represents a AccountBtcWalletParamsSet event raised by the WalletPayments contract.
type WalletPaymentsAccountBtcWalletParamsSet struct {
	WalletId     [32]byte
	AccountIndex uint32
	Params       IBtcParamsBtcWalletParams
	Raw          *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsAccountBtcWalletParamsSetEventName = "AccountBtcWalletParamsSet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsAccountBtcWalletParamsSet) ContractEventName() string {
	return WalletPaymentsAccountBtcWalletParamsSetEventName
}

// UnpackAccountBtcWalletParamsSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AccountBtcWalletParamsSet(bytes32 indexed walletId, uint32 indexed accountIndex, (uint32,uint8,uint8,uint8,uint8,uint8,uint8) params)
func (walletPayments *WalletPayments) UnpackAccountBtcWalletParamsSetEvent(log *types.Log) (*WalletPaymentsAccountBtcWalletParamsSet, error) {
	event := "AccountBtcWalletParamsSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsAccountBtcWalletParamsSet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsAccountFeeScheduleCleared represents a AccountFeeScheduleCleared event raised by the WalletPayments contract.
type WalletPaymentsAccountFeeScheduleCleared struct {
	ProjectId      [32]byte
	SourceId       [32]byte
	AccountAddress string
	AccountHash    [32]byte
	Raw            *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsAccountFeeScheduleClearedEventName = "AccountFeeScheduleCleared"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsAccountFeeScheduleCleared) ContractEventName() string {
	return WalletPaymentsAccountFeeScheduleClearedEventName
}

// UnpackAccountFeeScheduleClearedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AccountFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash)
func (walletPayments *WalletPayments) UnpackAccountFeeScheduleClearedEvent(log *types.Log) (*WalletPaymentsAccountFeeScheduleCleared, error) {
	event := "AccountFeeScheduleCleared"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsAccountFeeScheduleCleared)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsAccountFeeScheduleSet represents a AccountFeeScheduleSet event raised by the WalletPayments contract.
type WalletPaymentsAccountFeeScheduleSet struct {
	ProjectId      [32]byte
	SourceId       [32]byte
	AccountAddress string
	AccountHash    [32]byte
	Schedule       []IFeeSchedulesFeeSchedule
	Raw            *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsAccountFeeScheduleSetEventName = "AccountFeeScheduleSet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsAccountFeeScheduleSet) ContractEventName() string {
	return WalletPaymentsAccountFeeScheduleSetEventName
}

// UnpackAccountFeeScheduleSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AccountFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, string accountAddress, bytes32 indexed accountHash, (int16,uint16)[] schedule)
func (walletPayments *WalletPayments) UnpackAccountFeeScheduleSetEvent(log *types.Log) (*WalletPaymentsAccountFeeScheduleSet, error) {
	event := "AccountFeeScheduleSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsAccountFeeScheduleSet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsAccountProposersSet represents a AccountProposersSet event raised by the WalletPayments contract.
type WalletPaymentsAccountProposersSet struct {
	WalletId     [32]byte
	AccountIndex uint32
	Proposers    []common.Address
	Raw          *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsAccountProposersSetEventName = "AccountProposersSet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsAccountProposersSet) ContractEventName() string {
	return WalletPaymentsAccountProposersSetEventName
}

// UnpackAccountProposersSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event AccountProposersSet(bytes32 indexed walletId, uint32 indexed accountIndex, address[] proposers)
func (walletPayments *WalletPayments) UnpackAccountProposersSetEvent(log *types.Log) (*WalletPaymentsAccountProposersSet, error) {
	event := "AccountProposersSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsAccountProposersSet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsBtcAccountAdded represents a BtcAccountAdded event raised by the WalletPayments contract.
type WalletPaymentsBtcAccountAdded struct {
	WalletRegistry       common.Address
	WalletId             [32]byte
	SourceId             [32]byte
	AccountAddress       string
	AccountIndex         uint32
	AnchorCount          uint32
	AuthorizationAddress common.Address
	Raw                  *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsBtcAccountAddedEventName = "BtcAccountAdded"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsBtcAccountAdded) ContractEventName() string {
	return WalletPaymentsBtcAccountAddedEventName
}

// UnpackBtcAccountAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BtcAccountAdded(address indexed walletRegistry, bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount, address authorizationAddress)
func (walletPayments *WalletPayments) UnpackBtcAccountAddedEvent(log *types.Log) (*WalletPaymentsBtcAccountAdded, error) {
	event := "BtcAccountAdded"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsBtcAccountAdded)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsBtcAnchorsAdded represents a BtcAnchorsAdded event raised by the WalletPayments contract.
type WalletPaymentsBtcAnchorsAdded struct {
	WalletRegistry common.Address
	WalletId       [32]byte
	SourceId       [32]byte
	AccountAddress string
	AccountIndex   uint32
	AnchorCount    uint32
	Raw            *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsBtcAnchorsAddedEventName = "BtcAnchorsAdded"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsBtcAnchorsAdded) ContractEventName() string {
	return WalletPaymentsBtcAnchorsAddedEventName
}

// UnpackBtcAnchorsAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BtcAnchorsAdded(address indexed walletRegistry, bytes32 indexed walletId, bytes32 sourceId, string accountAddress, uint32 accountIndex, uint32 anchorCount)
func (walletPayments *WalletPayments) UnpackBtcAnchorsAddedEvent(log *types.Log) (*WalletPaymentsBtcAnchorsAdded, error) {
	event := "BtcAnchorsAdded"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsBtcAnchorsAdded)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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
	Raw              *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsBtcAttemptSettledEventName = "BtcAttemptSettled"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsBtcAttemptSettled) ContractEventName() string {
	return WalletPaymentsBtcAttemptSettledEventName
}

// UnpackBtcAttemptSettledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BtcAttemptSettled(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, uint32 anchorIndex, uint64 nonce, bytes32 txid, bytes32 nextAnchorTxid, uint32 nextAnchorVout)
func (walletPayments *WalletPayments) UnpackBtcAttemptSettledEvent(log *types.Log) (*WalletPaymentsBtcAttemptSettled, error) {
	event := "BtcAttemptSettled"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsBtcAttemptSettled)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsBtcConsensusParamsSet represents a BtcConsensusParamsSet event raised by the WalletPayments contract.
type WalletPaymentsBtcConsensusParamsSet struct {
	SourceId [32]byte
	Params   IBtcParamsBtcConsensusParams
	Raw      *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsBtcConsensusParamsSetEventName = "BtcConsensusParamsSet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsBtcConsensusParamsSet) ContractEventName() string {
	return WalletPaymentsBtcConsensusParamsSetEventName
}

// UnpackBtcConsensusParamsSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BtcConsensusParamsSet(bytes32 indexed sourceId, (uint8,uint32,uint16,uint32,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint16,uint16,uint16,uint32) params)
func (walletPayments *WalletPayments) UnpackBtcConsensusParamsSetEvent(log *types.Log) (*WalletPaymentsBtcConsensusParamsSet, error) {
	event := "BtcConsensusParamsSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsBtcConsensusParamsSet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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
	Raw              *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsBtcEscrowEmittedEventName = "BtcEscrowEmitted"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsBtcEscrowEmitted) ContractEventName() string {
	return WalletPaymentsBtcEscrowEmittedEventName
}

// UnpackBtcEscrowEmittedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BtcEscrowEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed paymentId, uint64 sequencePosition, uint8 kind, bytes32 preimageHash, uint64 amount, uint64 expiresAt)
func (walletPayments *WalletPayments) UnpackBtcEscrowEmittedEvent(log *types.Log) (*WalletPaymentsBtcEscrowEmitted, error) {
	event := "BtcEscrowEmitted"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsBtcEscrowEmitted)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsBtcEscrowProfileSet represents a BtcEscrowProfileSet event raised by the WalletPayments contract.
type WalletPaymentsBtcEscrowProfileSet struct {
	WalletId           [32]byte
	AccountIndex       uint32
	CounterpartyPubKey []byte
	Raw                *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsBtcEscrowProfileSetEventName = "BtcEscrowProfileSet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsBtcEscrowProfileSet) ContractEventName() string {
	return WalletPaymentsBtcEscrowProfileSetEventName
}

// UnpackBtcEscrowProfileSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BtcEscrowProfileSet(bytes32 indexed walletId, uint32 indexed accountIndex, bytes counterpartyPubKey)
func (walletPayments *WalletPayments) UnpackBtcEscrowProfileSetEvent(log *types.Log) (*WalletPaymentsBtcEscrowProfileSet, error) {
	event := "BtcEscrowProfileSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsBtcEscrowProfileSet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsBtcEscrowReclaimEmitted represents a BtcEscrowReclaimEmitted event raised by the WalletPayments contract.
type WalletPaymentsBtcEscrowReclaimEmitted struct {
	WalletId        [32]byte
	AccountIndex    uint32
	PaymentId       uint64
	CreatePaymentId uint64
	Raw             *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsBtcEscrowReclaimEmittedEventName = "BtcEscrowReclaimEmitted"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsBtcEscrowReclaimEmitted) ContractEventName() string {
	return WalletPaymentsBtcEscrowReclaimEmittedEventName
}

// UnpackBtcEscrowReclaimEmittedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event BtcEscrowReclaimEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed paymentId, uint64 createPaymentId)
func (walletPayments *WalletPayments) UnpackBtcEscrowReclaimEmittedEvent(log *types.Log) (*WalletPaymentsBtcEscrowReclaimEmitted, error) {
	event := "BtcEscrowReclaimEmitted"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsBtcEscrowReclaimEmitted)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsConfirmedAttemptSettled represents a ConfirmedAttemptSettled event raised by the WalletPayments contract.
type WalletPaymentsConfirmedAttemptSettled struct {
	WalletId          [32]byte
	AccountIndex      uint32
	SequencePosition  uint64
	Attempt           uint32
	SupersededAttempt uint32
	Txid              [32]byte
	Raw               *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsConfirmedAttemptSettledEventName = "ConfirmedAttemptSettled"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsConfirmedAttemptSettled) ContractEventName() string {
	return WalletPaymentsConfirmedAttemptSettledEventName
}

// UnpackConfirmedAttemptSettledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ConfirmedAttemptSettled(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, uint32 supersededAttempt, bytes32 txid)
func (walletPayments *WalletPayments) UnpackConfirmedAttemptSettledEvent(log *types.Log) (*WalletPaymentsConfirmedAttemptSettled, error) {
	event := "ConfirmedAttemptSettled"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsConfirmedAttemptSettled)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsCspFeeForwarded represents a CspFeeForwarded event raised by the WalletPayments contract.
type WalletPaymentsCspFeeForwarded struct {
	WalletId         [32]byte
	AccountIndex     uint32
	OpCommand        [32]byte
	PaymentId        uint64
	ClaimBackAddress common.Address
	Value            *big.Int
	RewardEpochId    *big.Int
	Raw              *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsCspFeeForwardedEventName = "CspFeeForwarded"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsCspFeeForwarded) ContractEventName() string {
	return WalletPaymentsCspFeeForwardedEventName
}

// UnpackCspFeeForwardedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CspFeeForwarded(bytes32 indexed walletId, uint32 indexed accountIndex, bytes32 indexed opCommand, uint64 paymentId, address claimBackAddress, uint256 value, uint24 rewardEpochId)
func (walletPayments *WalletPayments) UnpackCspFeeForwardedEvent(log *types.Log) (*WalletPaymentsCspFeeForwarded, error) {
	event := "CspFeeForwarded"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsCspFeeForwarded)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsCspSourceSettingsSet represents a CspSourceSettingsSet event raised by the WalletPayments contract.
type WalletPaymentsCspSourceSettingsSet struct {
	SourceId [32]byte
	Settings ICspProposalsCspSourceSettings
	Raw      *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsCspSourceSettingsSetEventName = "CspSourceSettingsSet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsCspSourceSettingsSet) ContractEventName() string {
	return WalletPaymentsCspSourceSettingsSetEventName
}

// UnpackCspSourceSettingsSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CspSourceSettingsSet(bytes32 indexed sourceId, (uint16,uint8,uint128) settings)
func (walletPayments *WalletPayments) UnpackCspSourceSettingsSetEvent(log *types.Log) (*WalletPaymentsCspSourceSettingsSet, error) {
	event := "CspSourceSettingsSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsCspSourceSettingsSet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsCustodianInstructionIssued represents a CustodianInstructionIssued event raised by the WalletPayments contract.
type WalletPaymentsCustodianInstructionIssued struct {
	WalletRegistry common.Address
	WalletId       [32]byte
	InstructionId  [32]byte
	OpType         [32]byte
	OpCommand      [32]byte
	Message        []byte
	Raw            *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsCustodianInstructionIssuedEventName = "CustodianInstructionIssued"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsCustodianInstructionIssued) ContractEventName() string {
	return WalletPaymentsCustodianInstructionIssuedEventName
}

// UnpackCustodianInstructionIssuedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event CustodianInstructionIssued(address indexed walletRegistry, bytes32 indexed walletId, bytes32 indexed instructionId, bytes32 opType, bytes32 opCommand, bytes message)
func (walletPayments *WalletPayments) UnpackCustodianInstructionIssuedEvent(log *types.Log) (*WalletPaymentsCustodianInstructionIssued, error) {
	event := "CustodianInstructionIssued"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsCustodianInstructionIssued)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsDiamondCut represents a DiamondCut event raised by the WalletPayments contract.
type WalletPaymentsDiamondCut struct {
	DiamondCut []IDiamondFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsDiamondCutEventName = "DiamondCut"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsDiamondCut) ContractEventName() string {
	return WalletPaymentsDiamondCutEventName
}

// UnpackDiamondCutEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (walletPayments *WalletPayments) UnpackDiamondCutEvent(log *types.Log) (*WalletPaymentsDiamondCut, error) {
	event := "DiamondCut"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsDiamondCut)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsEligibleAdvanced represents a EligibleAdvanced event raised by the WalletPayments contract.
type WalletPaymentsEligibleAdvanced struct {
	WalletId         [32]byte
	AccountIndex     uint32
	SequencePosition uint64
	Attempt          uint32
	Generation       uint64
	Raw              *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsEligibleAdvancedEventName = "EligibleAdvanced"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsEligibleAdvanced) ContractEventName() string {
	return WalletPaymentsEligibleAdvancedEventName
}

// UnpackEligibleAdvancedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event EligibleAdvanced(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 sequencePosition, uint32 attempt, uint64 generation)
func (walletPayments *WalletPayments) UnpackEligibleAdvancedEvent(log *types.Log) (*WalletPaymentsEligibleAdvanced, error) {
	event := "EligibleAdvanced"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsEligibleAdvanced)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsFdc2ProofPolicySet represents a Fdc2ProofPolicySet event raised by the WalletPayments contract.
type WalletPaymentsFdc2ProofPolicySet struct {
	Policy IFdc2ProofPolicyProofPolicy
	Raw    *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsFdc2ProofPolicySetEventName = "Fdc2ProofPolicySet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsFdc2ProofPolicySet) ContractEventName() string {
	return WalletPaymentsFdc2ProofPolicySetEventName
}

// UnpackFdc2ProofPolicySetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Fdc2ProofPolicySet((uint16,uint16) policy)
func (walletPayments *WalletPayments) UnpackFdc2ProofPolicySetEvent(log *types.Log) (*WalletPaymentsFdc2ProofPolicySet, error) {
	event := "Fdc2ProofPolicySet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsFdc2ProofPolicySet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsFeeScheduleConfigsCleared represents a FeeScheduleConfigsCleared event raised by the WalletPayments contract.
type WalletPaymentsFeeScheduleConfigsCleared struct {
	SourceIds [][32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsFeeScheduleConfigsClearedEventName = "FeeScheduleConfigsCleared"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsFeeScheduleConfigsCleared) ContractEventName() string {
	return WalletPaymentsFeeScheduleConfigsClearedEventName
}

// UnpackFeeScheduleConfigsClearedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeScheduleConfigsCleared(bytes32[] sourceIds)
func (walletPayments *WalletPayments) UnpackFeeScheduleConfigsClearedEvent(log *types.Log) (*WalletPaymentsFeeScheduleConfigsCleared, error) {
	event := "FeeScheduleConfigsCleared"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsFeeScheduleConfigsCleared)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsFeeScheduleConfigsSet represents a FeeScheduleConfigsSet event raised by the WalletPayments contract.
type WalletPaymentsFeeScheduleConfigsSet struct {
	Configs []IFeeSchedulesFeeScheduleConfigInput
	Raw     *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsFeeScheduleConfigsSetEventName = "FeeScheduleConfigsSet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsFeeScheduleConfigsSet) ContractEventName() string {
	return WalletPaymentsFeeScheduleConfigsSetEventName
}

// UnpackFeeScheduleConfigsSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeeScheduleConfigsSet((uint16,uint8,bytes32)[] configs)
func (walletPayments *WalletPayments) UnpackFeeScheduleConfigsSetEvent(log *types.Log) (*WalletPaymentsFeeScheduleConfigsSet, error) {
	event := "FeeScheduleConfigsSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsFeeScheduleConfigsSet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsFeesAccounted represents a FeesAccounted event raised by the WalletPayments contract.
type WalletPaymentsFeesAccounted struct {
	WalletId              [32]byte
	AccountIndex          uint32
	FirstSequencePosition uint64
	AccountedPosition     uint64
	ActualFee             *big.Int
	AuthorisedFee         *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsFeesAccountedEventName = "FeesAccounted"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsFeesAccounted) ContractEventName() string {
	return WalletPaymentsFeesAccountedEventName
}

// UnpackFeesAccountedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event FeesAccounted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed firstSequencePosition, uint64 accountedPosition, uint256 actualFee, uint256 authorisedFee)
func (walletPayments *WalletPayments) UnpackFeesAccountedEvent(log *types.Log) (*WalletPaymentsFeesAccounted, error) {
	event := "FeesAccounted"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsFeesAccounted)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the WalletPayments contract.
type WalletPaymentsGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsGovernanceCallTimelocked) ContractEventName() string {
	return WalletPaymentsGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (walletPayments *WalletPayments) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*WalletPaymentsGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsGovernanceInitialised represents a GovernanceInitialised event raised by the WalletPayments contract.
type WalletPaymentsGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsGovernanceInitialised) ContractEventName() string {
	return WalletPaymentsGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (walletPayments *WalletPayments) UnpackGovernanceInitialisedEvent(log *types.Log) (*WalletPaymentsGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the WalletPayments contract.
type WalletPaymentsGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsGovernedProductionModeEntered) ContractEventName() string {
	return WalletPaymentsGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (walletPayments *WalletPayments) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*WalletPaymentsGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsInstructionEmitted represents a InstructionEmitted event raised by the WalletPayments contract.
type WalletPaymentsInstructionEmitted struct {
	WalletId         [32]byte
	AccountIndex     uint32
	SequencePosition uint64
	Kind             uint8
	PaymentId        uint64
	EmittedAt        uint64
	MaxFee           uint64
	Raw              *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsInstructionEmittedEventName = "InstructionEmitted"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsInstructionEmitted) ContractEventName() string {
	return WalletPaymentsInstructionEmittedEventName
}

// UnpackInstructionEmittedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InstructionEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint8 kind, uint64 paymentId, uint64 emittedAt, uint64 maxFee)
func (walletPayments *WalletPayments) UnpackInstructionEmittedEvent(log *types.Log) (*WalletPaymentsInstructionEmitted, error) {
	event := "InstructionEmitted"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsInstructionEmitted)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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
	Raw                 *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsInstructionSettledEventName = "InstructionSettled"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsInstructionSettled) ContractEventName() string {
	return WalletPaymentsInstructionSettledEventName
}

// UnpackInstructionSettledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event InstructionSettled(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, bytes32 chainCommitmentHash, address settler, bytes message)
func (walletPayments *WalletPayments) UnpackInstructionSettledEvent(log *types.Log) (*WalletPaymentsInstructionSettled, error) {
	event := "InstructionSettled"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsInstructionSettled)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsNativeNonceAccountAdded represents a NativeNonceAccountAdded event raised by the WalletPayments contract.
type WalletPaymentsNativeNonceAccountAdded struct {
	WalletRegistry       common.Address
	WalletId             [32]byte
	SourceId             [32]byte
	AccountAddress       string
	AuthorizationAddress common.Address
	InitialNonce         uint64
	Raw                  *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsNativeNonceAccountAddedEventName = "NativeNonceAccountAdded"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsNativeNonceAccountAdded) ContractEventName() string {
	return WalletPaymentsNativeNonceAccountAddedEventName
}

// UnpackNativeNonceAccountAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event NativeNonceAccountAdded(address indexed walletRegistry, bytes32 indexed walletId, bytes32 sourceId, string accountAddress, address authorizationAddress, uint64 initialNonce)
func (walletPayments *WalletPayments) UnpackNativeNonceAccountAddedEvent(log *types.Log) (*WalletPaymentsNativeNonceAccountAdded, error) {
	event := "NativeNonceAccountAdded"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsNativeNonceAccountAdded)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsOperationNullified represents a OperationNullified event raised by the WalletPayments contract.
type WalletPaymentsOperationNullified struct {
	WalletId         [32]byte
	AccountIndex     uint32
	SequencePosition uint64
	PaymentId        uint64
	Attempt          uint32
	MaxFee           uint64
	Raw              *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsOperationNullifiedEventName = "OperationNullified"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsOperationNullified) ContractEventName() string {
	return WalletPaymentsOperationNullifiedEventName
}

// UnpackOperationNullifiedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event OperationNullified(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint64 paymentId, uint32 attempt, uint64 maxFee)
func (walletPayments *WalletPayments) UnpackOperationNullifiedEvent(log *types.Log) (*WalletPaymentsOperationNullified, error) {
	event := "OperationNullified"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsOperationNullified)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsPaymentBatched represents a PaymentBatched event raised by the WalletPayments contract.
type WalletPaymentsPaymentBatched struct {
	InstructionId [32]byte
	PaymentId     uint64
	Message       []byte
	Raw           *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsPaymentBatchedEventName = "PaymentBatched"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsPaymentBatched) ContractEventName() string {
	return WalletPaymentsPaymentBatchedEventName
}

// UnpackPaymentBatchedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentBatched(bytes32 indexed instructionId, uint64 indexed paymentId, bytes message)
func (walletPayments *WalletPayments) UnpackPaymentBatchedEvent(log *types.Log) (*WalletPaymentsPaymentBatched, error) {
	event := "PaymentBatched"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsPaymentBatched)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsPaymentQueued represents a PaymentQueued event raised by the WalletPayments contract.
type WalletPaymentsPaymentQueued struct {
	WalletId         [32]byte
	AccountIndex     uint32
	PaymentId        uint64
	RecipientAddress string
	Amount           uint64
	MaxFee           uint64
	PaymentReference [32]byte
	Raw              *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsPaymentQueuedEventName = "PaymentQueued"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsPaymentQueued) ContractEventName() string {
	return WalletPaymentsPaymentQueuedEventName
}

// UnpackPaymentQueuedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PaymentQueued(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 paymentId, string recipientAddress, uint64 amount, uint64 maxFee, bytes32 paymentReference)
func (walletPayments *WalletPayments) UnpackPaymentQueuedEvent(log *types.Log) (*WalletPaymentsPaymentQueued, error) {
	event := "PaymentQueued"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsPaymentQueued)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsProjectBtcWalletParamsSet represents a ProjectBtcWalletParamsSet event raised by the WalletPayments contract.
type WalletPaymentsProjectBtcWalletParamsSet struct {
	ProjectId [32]byte
	Params    IBtcParamsBtcWalletParams
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsProjectBtcWalletParamsSetEventName = "ProjectBtcWalletParamsSet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsProjectBtcWalletParamsSet) ContractEventName() string {
	return WalletPaymentsProjectBtcWalletParamsSetEventName
}

// UnpackProjectBtcWalletParamsSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProjectBtcWalletParamsSet(bytes32 indexed projectId, (uint32,uint8,uint8,uint8,uint8,uint8,uint8) params)
func (walletPayments *WalletPayments) UnpackProjectBtcWalletParamsSetEvent(log *types.Log) (*WalletPaymentsProjectBtcWalletParamsSet, error) {
	event := "ProjectBtcWalletParamsSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsProjectBtcWalletParamsSet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsProjectFeeScheduleCleared represents a ProjectFeeScheduleCleared event raised by the WalletPayments contract.
type WalletPaymentsProjectFeeScheduleCleared struct {
	ProjectId [32]byte
	SourceId  [32]byte
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsProjectFeeScheduleClearedEventName = "ProjectFeeScheduleCleared"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsProjectFeeScheduleCleared) ContractEventName() string {
	return WalletPaymentsProjectFeeScheduleClearedEventName
}

// UnpackProjectFeeScheduleClearedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProjectFeeScheduleCleared(bytes32 indexed projectId, bytes32 indexed sourceId)
func (walletPayments *WalletPayments) UnpackProjectFeeScheduleClearedEvent(log *types.Log) (*WalletPaymentsProjectFeeScheduleCleared, error) {
	event := "ProjectFeeScheduleCleared"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsProjectFeeScheduleCleared)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsProjectFeeScheduleSet represents a ProjectFeeScheduleSet event raised by the WalletPayments contract.
type WalletPaymentsProjectFeeScheduleSet struct {
	ProjectId [32]byte
	SourceId  [32]byte
	Schedule  []IFeeSchedulesFeeSchedule
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsProjectFeeScheduleSetEventName = "ProjectFeeScheduleSet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsProjectFeeScheduleSet) ContractEventName() string {
	return WalletPaymentsProjectFeeScheduleSetEventName
}

// UnpackProjectFeeScheduleSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProjectFeeScheduleSet(bytes32 indexed projectId, bytes32 indexed sourceId, (int16,uint16)[] schedule)
func (walletPayments *WalletPayments) UnpackProjectFeeScheduleSetEvent(log *types.Log) (*WalletPaymentsProjectFeeScheduleSet, error) {
	event := "ProjectFeeScheduleSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsProjectFeeScheduleSet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsProjectProposersSet represents a ProjectProposersSet event raised by the WalletPayments contract.
type WalletPaymentsProjectProposersSet struct {
	ProjectId [32]byte
	Proposers []common.Address
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsProjectProposersSetEventName = "ProjectProposersSet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsProjectProposersSet) ContractEventName() string {
	return WalletPaymentsProjectProposersSetEventName
}

// UnpackProjectProposersSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProjectProposersSet(bytes32 indexed projectId, address[] proposers)
func (walletPayments *WalletPayments) UnpackProjectProposersSetEvent(log *types.Log) (*WalletPaymentsProjectProposersSet, error) {
	event := "ProjectProposersSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsProjectProposersSet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsProposalContended represents a ProposalContended event raised by the WalletPayments contract.
type WalletPaymentsProposalContended struct {
	WalletId         [32]byte
	AccountIndex     uint32
	SequencePosition uint64
	Attempt          uint32
	PackageHash      [32]byte
	Proposer         common.Address
	Score            uint64
	Raw              *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsProposalContendedEventName = "ProposalContended"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsProposalContended) ContractEventName() string {
	return WalletPaymentsProposalContendedEventName
}

// UnpackProposalContendedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalContended(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, address proposer, uint64 score)
func (walletPayments *WalletPayments) UnpackProposalContendedEvent(log *types.Log) (*WalletPaymentsProposalContended, error) {
	event := "ProposalContended"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsProposalContended)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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
	Raw              *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsProposalLeadingEventName = "ProposalLeading"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsProposalLeading) ContractEventName() string {
	return WalletPaymentsProposalLeadingEventName
}

// UnpackProposalLeadingEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposalLeading(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, address proposer, uint64 score, uint64 graceEndsAt)
func (walletPayments *WalletPayments) UnpackProposalLeadingEvent(log *types.Log) (*WalletPaymentsProposalLeading, error) {
	event := "ProposalLeading"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsProposalLeading)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsProposerUrlSet represents a ProposerUrlSet event raised by the WalletPayments contract.
type WalletPaymentsProposerUrlSet struct {
	Proposer common.Address
	Url      string
	Raw      *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsProposerUrlSetEventName = "ProposerUrlSet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsProposerUrlSet) ContractEventName() string {
	return WalletPaymentsProposerUrlSetEventName
}

// UnpackProposerUrlSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ProposerUrlSet(address indexed proposer, string url)
func (walletPayments *WalletPayments) UnpackProposerUrlSetEvent(log *types.Log) (*WalletPaymentsProposerUrlSet, error) {
	event := "ProposerUrlSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsProposerUrlSet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsReissueEmitted represents a ReissueEmitted event raised by the WalletPayments contract.
type WalletPaymentsReissueEmitted struct {
	WalletId            [32]byte
	AccountIndex        uint32
	SequencePosition    uint64
	Attempt             uint32
	NullifiedPaymentIds []uint64
	MaxFee              uint64
	Raw                 *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsReissueEmittedEventName = "ReissueEmitted"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsReissueEmitted) ContractEventName() string {
	return WalletPaymentsReissueEmittedEventName
}

// UnpackReissueEmittedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event ReissueEmitted(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, uint64[] nullifiedPaymentIds, uint64 maxFee)
func (walletPayments *WalletPayments) UnpackReissueEmittedEvent(log *types.Log) (*WalletPaymentsReissueEmitted, error) {
	event := "ReissueEmitted"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsReissueEmitted)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsSigningRequested represents a SigningRequested event raised by the WalletPayments contract.
type WalletPaymentsSigningRequested struct {
	WalletId            [32]byte
	AccountIndex        uint32
	SequencePosition    uint64
	Attempt             uint32
	PackageHash         [32]byte
	ChainCommitmentHash [32]byte
	InstructionId       [32]byte
	Raw                 *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsSigningRequestedEventName = "SigningRequested"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsSigningRequested) ContractEventName() string {
	return WalletPaymentsSigningRequestedEventName
}

// UnpackSigningRequestedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SigningRequested(bytes32 indexed walletId, uint32 indexed accountIndex, uint64 indexed sequencePosition, uint32 attempt, bytes32 packageHash, bytes32 chainCommitmentHash, bytes32 instructionId)
func (walletPayments *WalletPayments) UnpackSigningRequestedEvent(log *types.Log) (*WalletPaymentsSigningRequested, error) {
	event := "SigningRequested"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsSigningRequested)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsSourceEscrowsEnabled represents a SourceEscrowsEnabled event raised by the WalletPayments contract.
type WalletPaymentsSourceEscrowsEnabled struct {
	SourceId         [32]byte
	CustodianWallets bool
	TeeWallets       bool
	Raw              *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsSourceEscrowsEnabledEventName = "SourceEscrowsEnabled"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsSourceEscrowsEnabled) ContractEventName() string {
	return WalletPaymentsSourceEscrowsEnabledEventName
}

// UnpackSourceEscrowsEnabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SourceEscrowsEnabled(bytes32 indexed sourceId, bool custodianWallets, bool teeWallets)
func (walletPayments *WalletPayments) UnpackSourceEscrowsEnabledEvent(log *types.Log) (*WalletPaymentsSourceEscrowsEnabled, error) {
	event := "SourceEscrowsEnabled"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsSourceEscrowsEnabled)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsSourcesEnabled represents a SourcesEnabled event raised by the WalletPayments contract.
type WalletPaymentsSourcesEnabled struct {
	SourceIds [][32]byte
	Enabled   bool
	Raw       *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsSourcesEnabledEventName = "SourcesEnabled"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsSourcesEnabled) ContractEventName() string {
	return WalletPaymentsSourcesEnabledEventName
}

// UnpackSourcesEnabledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SourcesEnabled(bytes32[] sourceIds, bool enabled)
func (walletPayments *WalletPayments) UnpackSourcesEnabledEvent(log *types.Log) (*WalletPaymentsSourcesEnabled, error) {
	event := "SourcesEnabled"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsSourcesEnabled)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsSourcesRegistered represents a SourcesRegistered event raised by the WalletPayments contract.
type WalletPaymentsSourcesRegistered struct {
	Registrations []ISourceConfigSourceRegistration
	Raw           *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsSourcesRegisteredEventName = "SourcesRegistered"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsSourcesRegistered) ContractEventName() string {
	return WalletPaymentsSourcesRegisteredEventName
}

// UnpackSourcesRegisteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event SourcesRegistered((bytes32,bytes32,bytes32,uint8,uint8)[] registrations)
func (walletPayments *WalletPayments) UnpackSourcesRegisteredEvent(log *types.Log) (*WalletPaymentsSourcesRegistered, error) {
	event := "SourcesRegistered"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsSourcesRegistered)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the WalletPayments contract.
type WalletPaymentsTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsTimelockedGovernanceCallCanceled) ContractEventName() string {
	return WalletPaymentsTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (walletPayments *WalletPayments) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*WalletPaymentsTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the WalletPayments contract.
type WalletPaymentsTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsTimelockedGovernanceCallExecuted) ContractEventName() string {
	return WalletPaymentsTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (walletPayments *WalletPayments) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*WalletPaymentsTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsXrpEscrowInstructed represents a XrpEscrowInstructed event raised by the WalletPayments contract.
type WalletPaymentsXrpEscrowInstructed struct {
	WalletId      [32]byte
	PaymentId     uint64
	ReissueNumber uint64
	MaxFee        *big.Int
	Raw           *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsXrpEscrowInstructedEventName = "XrpEscrowInstructed"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsXrpEscrowInstructed) ContractEventName() string {
	return WalletPaymentsXrpEscrowInstructedEventName
}

// UnpackXrpEscrowInstructedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event XrpEscrowInstructed(bytes32 indexed walletId, uint64 indexed paymentId, uint64 reissueNumber, uint256 maxFee)
func (walletPayments *WalletPayments) UnpackXrpEscrowInstructedEvent(log *types.Log) (*WalletPaymentsXrpEscrowInstructed, error) {
	event := "XrpEscrowInstructed"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsXrpEscrowInstructed)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsXrpEscrowNullified represents a XrpEscrowNullified event raised by the WalletPayments contract.
type WalletPaymentsXrpEscrowNullified struct {
	WalletId      [32]byte
	PaymentId     uint64
	ReissueNumber uint64
	Raw           *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsXrpEscrowNullifiedEventName = "XrpEscrowNullified"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsXrpEscrowNullified) ContractEventName() string {
	return WalletPaymentsXrpEscrowNullifiedEventName
}

// UnpackXrpEscrowNullifiedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event XrpEscrowNullified(bytes32 indexed walletId, uint64 indexed paymentId, uint64 reissueNumber)
func (walletPayments *WalletPayments) UnpackXrpEscrowNullifiedEvent(log *types.Log) (*WalletPaymentsXrpEscrowNullified, error) {
	event := "XrpEscrowNullified"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsXrpEscrowNullified)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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

// WalletPaymentsXrpEscrowProfileSet represents a XrpEscrowProfileSet event raised by the WalletPayments contract.
type WalletPaymentsXrpEscrowProfileSet struct {
	WalletId    [32]byte
	Destination string
	Raw         *types.Log // Blockchain specific contextual infos
}

const WalletPaymentsXrpEscrowProfileSetEventName = "XrpEscrowProfileSet"

// ContractEventName returns the user-defined event name.
func (WalletPaymentsXrpEscrowProfileSet) ContractEventName() string {
	return WalletPaymentsXrpEscrowProfileSetEventName
}

// UnpackXrpEscrowProfileSetEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event XrpEscrowProfileSet(bytes32 indexed walletId, string destination)
func (walletPayments *WalletPayments) UnpackXrpEscrowProfileSetEvent(log *types.Log) (*WalletPaymentsXrpEscrowProfileSet, error) {
	event := "XrpEscrowProfileSet"
	if len(log.Topics) == 0 || log.Topics[0] != walletPayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(WalletPaymentsXrpEscrowProfileSet)
	if len(log.Data) > 0 {
		if err := walletPayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range walletPayments.abi.Events[event].Inputs {
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
func (walletPayments *WalletPayments) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AccountAddressAlreadySet"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAccountAddressAlreadySetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AccountAddressZero"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAccountAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AccountIndexAlreadyUsed"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAccountIndexAlreadyUsedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AccountIndexMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAccountIndexMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AccountNotRegistered"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAccountNotRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AlreadyInProductionMode"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAlreadyInProductionModeError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AmountTooLarge"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAmountTooLargeError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AnchorIndexOutOfBounds"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAnchorIndexOutOfBoundsError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AnchorLimitExceeded"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAnchorLimitExceededError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AnchorNonceRegression"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAnchorNonceRegressionError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AnchorSetEmpty"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAnchorSetEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AttemptAlreadyConfirmed"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAttemptAlreadyConfirmedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AttemptNotSettled"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAttemptNotSettledError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AttemptSuperseded"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAttemptSupersededError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["AuthorizationAddressZero"].ID.Bytes()[:4]) {
		return walletPayments.UnpackAuthorizationAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["BtcConsensusParamsNotSet"].ID.Bytes()[:4]) {
		return walletPayments.UnpackBtcConsensusParamsNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["BtcEscrowAlreadyReclaimed"].ID.Bytes()[:4]) {
		return walletPayments.UnpackBtcEscrowAlreadyReclaimedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["BtcEscrowNotExpired"].ID.Bytes()[:4]) {
		return walletPayments.UnpackBtcEscrowNotExpiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["BtcEscrowNotFound"].ID.Bytes()[:4]) {
		return walletPayments.UnpackBtcEscrowNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["BtcEscrowNotSettled"].ID.Bytes()[:4]) {
		return walletPayments.UnpackBtcEscrowNotSettledError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["CannotNullifyPayment"].ID.Bytes()[:4]) {
		return walletPayments.UnpackCannotNullifyPaymentError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["CommitmentMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackCommitmentMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["ConfirmedAttemptMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackConfirmedAttemptMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["CspAccountNotRegistered"].ID.Bytes()[:4]) {
		return walletPayments.UnpackCspAccountNotRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["DelayTooLarge"].ID.Bytes()[:4]) {
		return walletPayments.UnpackDelayTooLargeError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["DuplicateAnchor"].ID.Bytes()[:4]) {
		return walletPayments.UnpackDuplicateAnchorError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["EmptyProposerList"].ID.Bytes()[:4]) {
		return walletPayments.UnpackEmptyProposerListError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["EmptyScheduleNotAllowed"].ID.Bytes()[:4]) {
		return walletPayments.UnpackEmptyScheduleNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["EscrowProfileNotSet"].ID.Bytes()[:4]) {
		return walletPayments.UnpackEscrowProfileNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["EscrowsNotEnabled"].ID.Bytes()[:4]) {
		return walletPayments.UnpackEscrowsNotEnabledError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["Fdc2ProofPolicyNotSet"].ID.Bytes()[:4]) {
		return walletPayments.UnpackFdc2ProofPolicyNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["FeeProofRangeInvalid"].ID.Bytes()[:4]) {
		return walletPayments.UnpackFeeProofRangeInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["FeeScheduleConfigNotSet"].ID.Bytes()[:4]) {
		return walletPayments.UnpackFeeScheduleConfigNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["FeeScheduleNotSet"].ID.Bytes()[:4]) {
		return walletPayments.UnpackFeeScheduleNotSetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["FeeSchedulesUnsupported"].ID.Bytes()[:4]) {
		return walletPayments.UnpackFeeSchedulesUnsupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["GovernedAddressZero"].ID.Bytes()[:4]) {
		return walletPayments.UnpackGovernedAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["GovernedAlreadyInitialized"].ID.Bytes()[:4]) {
		return walletPayments.UnpackGovernedAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["GraceClosed"].ID.Bytes()[:4]) {
		return walletPayments.UnpackGraceClosedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["GraceOpen"].ID.Bytes()[:4]) {
		return walletPayments.UnpackGraceOpenError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InsufficientFee"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInsufficientFeeError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InsufficientTeeSignatures"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInsufficientTeeSignaturesError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidAttestation"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidAttestationError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidBtcConsensusParams"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidBtcConsensusParamsError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidBtcEscrowProfile"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidBtcEscrowProfileError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidBtcEscrowTerms"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidBtcEscrowTermsError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidBtcWalletParams"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidBtcWalletParamsError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidCspSourceSettings"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidCspSourceSettingsError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidFdc2ProofPolicy"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidFdc2ProofPolicyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidFeeDelay"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidFeeDelayError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidFeeFactor"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidFeeFactorError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidFeeScheduleConfig"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidFeeScheduleConfigError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidGrace"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidGraceError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidPaymentCount"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidPaymentCountError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidPaymentId"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidPaymentIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidPaymentInstructionCount"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidPaymentInstructionCountError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidProof"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidProofError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidRecipientAddress"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidRecipientAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidRequestBody"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidRequestBodyError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["InvalidXrpEscrowTerms"].ID.Bytes()[:4]) {
		return walletPayments.UnpackInvalidXrpEscrowTermsError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["LaneMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackLaneMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["LeaderExists"].ID.Bytes()[:4]) {
		return walletPayments.UnpackLeaderExistsError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["MaxFeeTooLarge"].ID.Bytes()[:4]) {
		return walletPayments.UnpackMaxFeeTooLargeError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["MaxFeeZero"].ID.Bytes()[:4]) {
		return walletPayments.UnpackMaxFeeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["NoNewAnchors"].ID.Bytes()[:4]) {
		return walletPayments.UnpackNoNewAnchorsError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["NoPendingReset"].ID.Bytes()[:4]) {
		return walletPayments.UnpackNoPendingResetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["NonceMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackNonceMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["NotEligible"].ID.Bytes()[:4]) {
		return walletPayments.UnpackNotEligibleError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["NotLaneTip"].ID.Bytes()[:4]) {
		return walletPayments.UnpackNotLaneTipError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["NotSettled"].ID.Bytes()[:4]) {
		return walletPayments.UnpackNotSettledError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["NothingToSettle"].ID.Bytes()[:4]) {
		return walletPayments.UnpackNothingToSettleError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["NullifiedPaymentsUnsupported"].ID.Bytes()[:4]) {
		return walletPayments.UnpackNullifiedPaymentsUnsupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["OnlyAuthorizationAddress"].ID.Bytes()[:4]) {
		return walletPayments.UnpackOnlyAuthorizationAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["OnlyExecutor"].ID.Bytes()[:4]) {
		return walletPayments.UnpackOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return walletPayments.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return walletPayments.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["OnlyProjectOwner"].ID.Bytes()[:4]) {
		return walletPayments.UnpackOnlyProjectOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["OnlySystemExtensionId"].ID.Bytes()[:4]) {
		return walletPayments.UnpackOnlySystemExtensionIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["OnlyWalletOwner"].ID.Bytes()[:4]) {
		return walletPayments.UnpackOnlyWalletOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["OperationAlreadyNullified"].ID.Bytes()[:4]) {
		return walletPayments.UnpackOperationAlreadyNullifiedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["OperationNotSupported"].ID.Bytes()[:4]) {
		return walletPayments.UnpackOperationNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["OperationStatusMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackOperationStatusMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["OverrideAlreadyPending"].ID.Bytes()[:4]) {
		return walletPayments.UnpackOverrideAlreadyPendingError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["PaymentAmountTooLow"].ID.Bytes()[:4]) {
		return walletPayments.UnpackPaymentAmountTooLowError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["PaymentAmountZero"].ID.Bytes()[:4]) {
		return walletPayments.UnpackPaymentAmountZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["PaymentHashMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackPaymentHashMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["PaymentNotSettled"].ID.Bytes()[:4]) {
		return walletPayments.UnpackPaymentNotSettledError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["PaymentRangeInvalid"].ID.Bytes()[:4]) {
		return walletPayments.UnpackPaymentRangeInvalidError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["ProposerNotWhitelisted"].ID.Bytes()[:4]) {
		return walletPayments.UnpackProposerNotWhitelistedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["ReissueDeclarationMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackReissueDeclarationMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["SourceAlreadyRegistered"].ID.Bytes()[:4]) {
		return walletPayments.UnpackSourceAlreadyRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["SourceChainKeyTypeMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackSourceChainKeyTypeMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["SourceDisabled"].ID.Bytes()[:4]) {
		return walletPayments.UnpackSourceDisabledError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["SourceIdZero"].ID.Bytes()[:4]) {
		return walletPayments.UnpackSourceIdZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["SourceKeyTypeNotSupported"].ID.Bytes()[:4]) {
		return walletPayments.UnpackSourceKeyTypeNotSupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["SourceKeyTypeZero"].ID.Bytes()[:4]) {
		return walletPayments.UnpackSourceKeyTypeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["SourceLimitsNotConfigured"].ID.Bytes()[:4]) {
		return walletPayments.UnpackSourceLimitsNotConfiguredError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["SourceModelChainMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackSourceModelChainMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["SourceNetworkKeyTypeMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackSourceNetworkKeyTypeMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["SourceNotRegistered"].ID.Bytes()[:4]) {
		return walletPayments.UnpackSourceNotRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["SourceOpTypeZero"].ID.Bytes()[:4]) {
		return walletPayments.UnpackSourceOpTypeZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["SourcePaymentModelUnknown"].ID.Bytes()[:4]) {
		return walletPayments.UnpackSourcePaymentModelUnknownError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["StoredAnchorsChanged"].ID.Bytes()[:4]) {
		return walletPayments.UnpackStoredAnchorsChangedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["TimelockCallNotFound"].ID.Bytes()[:4]) {
		return walletPayments.UnpackTimelockCallNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["TimelockNotAllowedYet"].ID.Bytes()[:4]) {
		return walletPayments.UnpackTimelockNotAllowedYetError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["TimelockValueNotAllowed"].ID.Bytes()[:4]) {
		return walletPayments.UnpackTimelockValueNotAllowedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["TokenIdUnsupported"].ID.Bytes()[:4]) {
		return walletPayments.UnpackTokenIdUnsupportedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["TooManySchedules"].ID.Bytes()[:4]) {
		return walletPayments.UnpackTooManySchedulesError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["UnknownWalletRegistry"].ID.Bytes()[:4]) {
		return walletPayments.UnpackUnknownWalletRegistryError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["UnsupportedPaymentModel"].ID.Bytes()[:4]) {
		return walletPayments.UnpackUnsupportedPaymentModelError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["UnsupportedSourceId"].ID.Bytes()[:4]) {
		return walletPayments.UnpackUnsupportedSourceIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["UsePaymentReissue"].ID.Bytes()[:4]) {
		return walletPayments.UnpackUsePaymentReissueError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["ValueNotExpected"].ID.Bytes()[:4]) {
		return walletPayments.UnpackValueNotExpectedError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["WalletNetworkMismatch"].ID.Bytes()[:4]) {
		return walletPayments.UnpackWalletNetworkMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["WalletNotInProduction"].ID.Bytes()[:4]) {
		return walletPayments.UnpackWalletNotInProductionError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["WrongKeyType"].ID.Bytes()[:4]) {
		return walletPayments.UnpackWrongKeyTypeError(raw[4:])
	}
	if bytes.Equal(raw[:4], walletPayments.abi.Errors["XrpEscrowNotFound"].ID.Bytes()[:4]) {
		return walletPayments.UnpackXrpEscrowNotFoundError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// WalletPaymentsAccountAddressAlreadySet represents a AccountAddressAlreadySet error raised by the WalletPayments contract.
type WalletPaymentsAccountAddressAlreadySet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccountAddressAlreadySet()
func WalletPaymentsAccountAddressAlreadySetErrorID() common.Hash {
	return common.HexToHash("0x910c94b29215a7b943c8b5838e7d8ed384a0a75ca1a6386a4426752822ff0223")
}

// UnpackAccountAddressAlreadySetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccountAddressAlreadySet()
func (walletPayments *WalletPayments) UnpackAccountAddressAlreadySetError(raw []byte) (*WalletPaymentsAccountAddressAlreadySet, error) {
	out := new(WalletPaymentsAccountAddressAlreadySet)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AccountAddressAlreadySet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAccountAddressZero represents a AccountAddressZero error raised by the WalletPayments contract.
type WalletPaymentsAccountAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccountAddressZero()
func WalletPaymentsAccountAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x7d01fca993146ff8e429fe0a909628970d27f0058569d84d8ec792b0847e5bd7")
}

// UnpackAccountAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccountAddressZero()
func (walletPayments *WalletPayments) UnpackAccountAddressZeroError(raw []byte) (*WalletPaymentsAccountAddressZero, error) {
	out := new(WalletPaymentsAccountAddressZero)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AccountAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAccountIndexAlreadyUsed represents a AccountIndexAlreadyUsed error raised by the WalletPayments contract.
type WalletPaymentsAccountIndexAlreadyUsed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccountIndexAlreadyUsed()
func WalletPaymentsAccountIndexAlreadyUsedErrorID() common.Hash {
	return common.HexToHash("0x0cc694067a70839ce45579cc3e5631b36a4b299c99489227067fdf3839cd6757")
}

// UnpackAccountIndexAlreadyUsedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccountIndexAlreadyUsed()
func (walletPayments *WalletPayments) UnpackAccountIndexAlreadyUsedError(raw []byte) (*WalletPaymentsAccountIndexAlreadyUsed, error) {
	out := new(WalletPaymentsAccountIndexAlreadyUsed)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AccountIndexAlreadyUsed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAccountIndexMismatch represents a AccountIndexMismatch error raised by the WalletPayments contract.
type WalletPaymentsAccountIndexMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccountIndexMismatch()
func WalletPaymentsAccountIndexMismatchErrorID() common.Hash {
	return common.HexToHash("0x27313024826f37a5dc59237f5f36a6005c35c5374c7feb74c5211b91b777fe75")
}

// UnpackAccountIndexMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccountIndexMismatch()
func (walletPayments *WalletPayments) UnpackAccountIndexMismatchError(raw []byte) (*WalletPaymentsAccountIndexMismatch, error) {
	out := new(WalletPaymentsAccountIndexMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AccountIndexMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAccountNotRegistered represents a AccountNotRegistered error raised by the WalletPayments contract.
type WalletPaymentsAccountNotRegistered struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccountNotRegistered()
func WalletPaymentsAccountNotRegisteredErrorID() common.Hash {
	return common.HexToHash("0x9704dd5103977e212a4ccc5d11b13fb9a4e452fb5b457b2b39c0fcad36ab2cf0")
}

// UnpackAccountNotRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccountNotRegistered()
func (walletPayments *WalletPayments) UnpackAccountNotRegisteredError(raw []byte) (*WalletPaymentsAccountNotRegistered, error) {
	out := new(WalletPaymentsAccountNotRegistered)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AccountNotRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAlreadyInProductionMode represents a AlreadyInProductionMode error raised by the WalletPayments contract.
type WalletPaymentsAlreadyInProductionMode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInProductionMode()
func WalletPaymentsAlreadyInProductionModeErrorID() common.Hash {
	return common.HexToHash("0x72a23ec2b1fdbc3de7d5f411b572fcf8ad8409803196930f63ddf40a2a2b2594")
}

// UnpackAlreadyInProductionModeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInProductionMode()
func (walletPayments *WalletPayments) UnpackAlreadyInProductionModeError(raw []byte) (*WalletPaymentsAlreadyInProductionMode, error) {
	out := new(WalletPaymentsAlreadyInProductionMode)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AlreadyInProductionMode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAmountTooLarge represents a AmountTooLarge error raised by the WalletPayments contract.
type WalletPaymentsAmountTooLarge struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AmountTooLarge()
func WalletPaymentsAmountTooLargeErrorID() common.Hash {
	return common.HexToHash("0x06250401c072295744520da9e60a3a29ab28ceb9fda5d792222d5242d22e8a44")
}

// UnpackAmountTooLargeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AmountTooLarge()
func (walletPayments *WalletPayments) UnpackAmountTooLargeError(raw []byte) (*WalletPaymentsAmountTooLarge, error) {
	out := new(WalletPaymentsAmountTooLarge)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AmountTooLarge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAnchorIndexOutOfBounds represents a AnchorIndexOutOfBounds error raised by the WalletPayments contract.
type WalletPaymentsAnchorIndexOutOfBounds struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AnchorIndexOutOfBounds()
func WalletPaymentsAnchorIndexOutOfBoundsErrorID() common.Hash {
	return common.HexToHash("0x2de58279687b2e18582d62fa0fa14630123921cc78446039d46922005d9e20a6")
}

// UnpackAnchorIndexOutOfBoundsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AnchorIndexOutOfBounds()
func (walletPayments *WalletPayments) UnpackAnchorIndexOutOfBoundsError(raw []byte) (*WalletPaymentsAnchorIndexOutOfBounds, error) {
	out := new(WalletPaymentsAnchorIndexOutOfBounds)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AnchorIndexOutOfBounds", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAnchorLimitExceeded represents a AnchorLimitExceeded error raised by the WalletPayments contract.
type WalletPaymentsAnchorLimitExceeded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AnchorLimitExceeded()
func WalletPaymentsAnchorLimitExceededErrorID() common.Hash {
	return common.HexToHash("0x53d9693388bb3de1aa9c38a0e4b38f8bd50c274a2615f12bf28ed76d2566a5ee")
}

// UnpackAnchorLimitExceededError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AnchorLimitExceeded()
func (walletPayments *WalletPayments) UnpackAnchorLimitExceededError(raw []byte) (*WalletPaymentsAnchorLimitExceeded, error) {
	out := new(WalletPaymentsAnchorLimitExceeded)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AnchorLimitExceeded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAnchorNonceRegression represents a AnchorNonceRegression error raised by the WalletPayments contract.
type WalletPaymentsAnchorNonceRegression struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AnchorNonceRegression()
func WalletPaymentsAnchorNonceRegressionErrorID() common.Hash {
	return common.HexToHash("0xd127642ab884465d1c3ed1b56a729e01cdd0bc14c80d415b4ce79ca7977c8428")
}

// UnpackAnchorNonceRegressionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AnchorNonceRegression()
func (walletPayments *WalletPayments) UnpackAnchorNonceRegressionError(raw []byte) (*WalletPaymentsAnchorNonceRegression, error) {
	out := new(WalletPaymentsAnchorNonceRegression)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AnchorNonceRegression", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAnchorSetEmpty represents a AnchorSetEmpty error raised by the WalletPayments contract.
type WalletPaymentsAnchorSetEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AnchorSetEmpty()
func WalletPaymentsAnchorSetEmptyErrorID() common.Hash {
	return common.HexToHash("0xb06cd5887ea5632bef5615859f9d82914081bba13f3044570559e85b2214c80e")
}

// UnpackAnchorSetEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AnchorSetEmpty()
func (walletPayments *WalletPayments) UnpackAnchorSetEmptyError(raw []byte) (*WalletPaymentsAnchorSetEmpty, error) {
	out := new(WalletPaymentsAnchorSetEmpty)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AnchorSetEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAttemptAlreadyConfirmed represents a AttemptAlreadyConfirmed error raised by the WalletPayments contract.
type WalletPaymentsAttemptAlreadyConfirmed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AttemptAlreadyConfirmed()
func WalletPaymentsAttemptAlreadyConfirmedErrorID() common.Hash {
	return common.HexToHash("0xa9285fad276bdf30c328d8e8042a536a9562f3796f6197a32d883ad4287886da")
}

// UnpackAttemptAlreadyConfirmedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AttemptAlreadyConfirmed()
func (walletPayments *WalletPayments) UnpackAttemptAlreadyConfirmedError(raw []byte) (*WalletPaymentsAttemptAlreadyConfirmed, error) {
	out := new(WalletPaymentsAttemptAlreadyConfirmed)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AttemptAlreadyConfirmed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAttemptNotSettled represents a AttemptNotSettled error raised by the WalletPayments contract.
type WalletPaymentsAttemptNotSettled struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AttemptNotSettled()
func WalletPaymentsAttemptNotSettledErrorID() common.Hash {
	return common.HexToHash("0x6d60958fab6d4008c99033d18577ef220dd6d0ec64510ff7a7124aacd5f44b8d")
}

// UnpackAttemptNotSettledError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AttemptNotSettled()
func (walletPayments *WalletPayments) UnpackAttemptNotSettledError(raw []byte) (*WalletPaymentsAttemptNotSettled, error) {
	out := new(WalletPaymentsAttemptNotSettled)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AttemptNotSettled", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAttemptSuperseded represents a AttemptSuperseded error raised by the WalletPayments contract.
type WalletPaymentsAttemptSuperseded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AttemptSuperseded()
func WalletPaymentsAttemptSupersededErrorID() common.Hash {
	return common.HexToHash("0xcccb6dd9518672a6d803f115b016f733f4dfb5d50d8aae5446d5ed83f63f3ae4")
}

// UnpackAttemptSupersededError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AttemptSuperseded()
func (walletPayments *WalletPayments) UnpackAttemptSupersededError(raw []byte) (*WalletPaymentsAttemptSuperseded, error) {
	out := new(WalletPaymentsAttemptSuperseded)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AttemptSuperseded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsAuthorizationAddressZero represents a AuthorizationAddressZero error raised by the WalletPayments contract.
type WalletPaymentsAuthorizationAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AuthorizationAddressZero()
func WalletPaymentsAuthorizationAddressZeroErrorID() common.Hash {
	return common.HexToHash("0xe4472ca7994bab3d6f38c9bdc243c438c59c017951973a129c8674da85c299a2")
}

// UnpackAuthorizationAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AuthorizationAddressZero()
func (walletPayments *WalletPayments) UnpackAuthorizationAddressZeroError(raw []byte) (*WalletPaymentsAuthorizationAddressZero, error) {
	out := new(WalletPaymentsAuthorizationAddressZero)
	if err := walletPayments.abi.UnpackIntoInterface(out, "AuthorizationAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsBtcConsensusParamsNotSet represents a BtcConsensusParamsNotSet error raised by the WalletPayments contract.
type WalletPaymentsBtcConsensusParamsNotSet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BtcConsensusParamsNotSet()
func WalletPaymentsBtcConsensusParamsNotSetErrorID() common.Hash {
	return common.HexToHash("0xd60a95cdd9f84388f9f99936f7f0dc2f9193c9cf3421813e5619cff8efc20f13")
}

// UnpackBtcConsensusParamsNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BtcConsensusParamsNotSet()
func (walletPayments *WalletPayments) UnpackBtcConsensusParamsNotSetError(raw []byte) (*WalletPaymentsBtcConsensusParamsNotSet, error) {
	out := new(WalletPaymentsBtcConsensusParamsNotSet)
	if err := walletPayments.abi.UnpackIntoInterface(out, "BtcConsensusParamsNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsBtcEscrowAlreadyReclaimed represents a BtcEscrowAlreadyReclaimed error raised by the WalletPayments contract.
type WalletPaymentsBtcEscrowAlreadyReclaimed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BtcEscrowAlreadyReclaimed()
func WalletPaymentsBtcEscrowAlreadyReclaimedErrorID() common.Hash {
	return common.HexToHash("0xee8cefb48a94f8565bb7f3a4d1e35484849dd540be9ab1e7058f5e40ed94c818")
}

// UnpackBtcEscrowAlreadyReclaimedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BtcEscrowAlreadyReclaimed()
func (walletPayments *WalletPayments) UnpackBtcEscrowAlreadyReclaimedError(raw []byte) (*WalletPaymentsBtcEscrowAlreadyReclaimed, error) {
	out := new(WalletPaymentsBtcEscrowAlreadyReclaimed)
	if err := walletPayments.abi.UnpackIntoInterface(out, "BtcEscrowAlreadyReclaimed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsBtcEscrowNotExpired represents a BtcEscrowNotExpired error raised by the WalletPayments contract.
type WalletPaymentsBtcEscrowNotExpired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BtcEscrowNotExpired()
func WalletPaymentsBtcEscrowNotExpiredErrorID() common.Hash {
	return common.HexToHash("0x5ae9de488f6f928c4765228061d508b466b4923cfe9d199871963722174a1587")
}

// UnpackBtcEscrowNotExpiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BtcEscrowNotExpired()
func (walletPayments *WalletPayments) UnpackBtcEscrowNotExpiredError(raw []byte) (*WalletPaymentsBtcEscrowNotExpired, error) {
	out := new(WalletPaymentsBtcEscrowNotExpired)
	if err := walletPayments.abi.UnpackIntoInterface(out, "BtcEscrowNotExpired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsBtcEscrowNotFound represents a BtcEscrowNotFound error raised by the WalletPayments contract.
type WalletPaymentsBtcEscrowNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BtcEscrowNotFound()
func WalletPaymentsBtcEscrowNotFoundErrorID() common.Hash {
	return common.HexToHash("0xcc40b55ec4ee57fe911eb22d8475be550a1d0b61219375e8d77ce97fc5afcfa4")
}

// UnpackBtcEscrowNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BtcEscrowNotFound()
func (walletPayments *WalletPayments) UnpackBtcEscrowNotFoundError(raw []byte) (*WalletPaymentsBtcEscrowNotFound, error) {
	out := new(WalletPaymentsBtcEscrowNotFound)
	if err := walletPayments.abi.UnpackIntoInterface(out, "BtcEscrowNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsBtcEscrowNotSettled represents a BtcEscrowNotSettled error raised by the WalletPayments contract.
type WalletPaymentsBtcEscrowNotSettled struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error BtcEscrowNotSettled()
func WalletPaymentsBtcEscrowNotSettledErrorID() common.Hash {
	return common.HexToHash("0xe44df67f5fe003f88535d10d5e3c7e81ff72f26dbfe01749cde527cdb911197c")
}

// UnpackBtcEscrowNotSettledError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error BtcEscrowNotSettled()
func (walletPayments *WalletPayments) UnpackBtcEscrowNotSettledError(raw []byte) (*WalletPaymentsBtcEscrowNotSettled, error) {
	out := new(WalletPaymentsBtcEscrowNotSettled)
	if err := walletPayments.abi.UnpackIntoInterface(out, "BtcEscrowNotSettled", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsCannotNullifyPayment represents a CannotNullifyPayment error raised by the WalletPayments contract.
type WalletPaymentsCannotNullifyPayment struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CannotNullifyPayment()
func WalletPaymentsCannotNullifyPaymentErrorID() common.Hash {
	return common.HexToHash("0xfb4defac4d7650326d08a541df5fe4f356670a6daba1d77e3f0d645221b34397")
}

// UnpackCannotNullifyPaymentError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CannotNullifyPayment()
func (walletPayments *WalletPayments) UnpackCannotNullifyPaymentError(raw []byte) (*WalletPaymentsCannotNullifyPayment, error) {
	out := new(WalletPaymentsCannotNullifyPayment)
	if err := walletPayments.abi.UnpackIntoInterface(out, "CannotNullifyPayment", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsCommitmentMismatch represents a CommitmentMismatch error raised by the WalletPayments contract.
type WalletPaymentsCommitmentMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CommitmentMismatch()
func WalletPaymentsCommitmentMismatchErrorID() common.Hash {
	return common.HexToHash("0x5054097bb92104fdc128156d244c4d9d3a1008b45209d260b0cb38908eec9de3")
}

// UnpackCommitmentMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CommitmentMismatch()
func (walletPayments *WalletPayments) UnpackCommitmentMismatchError(raw []byte) (*WalletPaymentsCommitmentMismatch, error) {
	out := new(WalletPaymentsCommitmentMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "CommitmentMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsConfirmedAttemptMismatch represents a ConfirmedAttemptMismatch error raised by the WalletPayments contract.
type WalletPaymentsConfirmedAttemptMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ConfirmedAttemptMismatch()
func WalletPaymentsConfirmedAttemptMismatchErrorID() common.Hash {
	return common.HexToHash("0xdd16bc70e965ceaecde843767115816704aa6b9da240cab0b84e42ef38b2c7ec")
}

// UnpackConfirmedAttemptMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ConfirmedAttemptMismatch()
func (walletPayments *WalletPayments) UnpackConfirmedAttemptMismatchError(raw []byte) (*WalletPaymentsConfirmedAttemptMismatch, error) {
	out := new(WalletPaymentsConfirmedAttemptMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "ConfirmedAttemptMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsCspAccountNotRegistered represents a CspAccountNotRegistered error raised by the WalletPayments contract.
type WalletPaymentsCspAccountNotRegistered struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CspAccountNotRegistered()
func WalletPaymentsCspAccountNotRegisteredErrorID() common.Hash {
	return common.HexToHash("0x80d979dcf6bd7b4cb7eae7df27e0c72f894e9d2bc490d9e6ac8eb95b5c080e32")
}

// UnpackCspAccountNotRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CspAccountNotRegistered()
func (walletPayments *WalletPayments) UnpackCspAccountNotRegisteredError(raw []byte) (*WalletPaymentsCspAccountNotRegistered, error) {
	out := new(WalletPaymentsCspAccountNotRegistered)
	if err := walletPayments.abi.UnpackIntoInterface(out, "CspAccountNotRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsDelayTooLarge represents a DelayTooLarge error raised by the WalletPayments contract.
type WalletPaymentsDelayTooLarge struct {
	Delay    *big.Int
	MaxDelay *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DelayTooLarge(uint256 delay, uint256 maxDelay)
func WalletPaymentsDelayTooLargeErrorID() common.Hash {
	return common.HexToHash("0xf77cfffc23bacd92259f81786482f236072e2ca49bedf646dcdb1d309c5f0ff4")
}

// UnpackDelayTooLargeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DelayTooLarge(uint256 delay, uint256 maxDelay)
func (walletPayments *WalletPayments) UnpackDelayTooLargeError(raw []byte) (*WalletPaymentsDelayTooLarge, error) {
	out := new(WalletPaymentsDelayTooLarge)
	if err := walletPayments.abi.UnpackIntoInterface(out, "DelayTooLarge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsDuplicateAnchor represents a DuplicateAnchor error raised by the WalletPayments contract.
type WalletPaymentsDuplicateAnchor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error DuplicateAnchor()
func WalletPaymentsDuplicateAnchorErrorID() common.Hash {
	return common.HexToHash("0x5b7386047b7de0cf066981bf8a6c0dbc443bf7191f341f16a27adbad7317ae23")
}

// UnpackDuplicateAnchorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error DuplicateAnchor()
func (walletPayments *WalletPayments) UnpackDuplicateAnchorError(raw []byte) (*WalletPaymentsDuplicateAnchor, error) {
	out := new(WalletPaymentsDuplicateAnchor)
	if err := walletPayments.abi.UnpackIntoInterface(out, "DuplicateAnchor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsEmptyProposerList represents a EmptyProposerList error raised by the WalletPayments contract.
type WalletPaymentsEmptyProposerList struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmptyProposerList()
func WalletPaymentsEmptyProposerListErrorID() common.Hash {
	return common.HexToHash("0xe7d5612d4888cfbe936f39d03c31e06382d802a48dd8b89776a5b86b8c8d0c9d")
}

// UnpackEmptyProposerListError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmptyProposerList()
func (walletPayments *WalletPayments) UnpackEmptyProposerListError(raw []byte) (*WalletPaymentsEmptyProposerList, error) {
	out := new(WalletPaymentsEmptyProposerList)
	if err := walletPayments.abi.UnpackIntoInterface(out, "EmptyProposerList", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsEmptyScheduleNotAllowed represents a EmptyScheduleNotAllowed error raised by the WalletPayments contract.
type WalletPaymentsEmptyScheduleNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EmptyScheduleNotAllowed()
func WalletPaymentsEmptyScheduleNotAllowedErrorID() common.Hash {
	return common.HexToHash("0xcdc0f5e297e9ba099b219e61725f34d2cf720b212c9e7d68571ab793321ba39d")
}

// UnpackEmptyScheduleNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EmptyScheduleNotAllowed()
func (walletPayments *WalletPayments) UnpackEmptyScheduleNotAllowedError(raw []byte) (*WalletPaymentsEmptyScheduleNotAllowed, error) {
	out := new(WalletPaymentsEmptyScheduleNotAllowed)
	if err := walletPayments.abi.UnpackIntoInterface(out, "EmptyScheduleNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsEscrowProfileNotSet represents a EscrowProfileNotSet error raised by the WalletPayments contract.
type WalletPaymentsEscrowProfileNotSet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EscrowProfileNotSet()
func WalletPaymentsEscrowProfileNotSetErrorID() common.Hash {
	return common.HexToHash("0xdfb33a39d012c232b87653dedb7c0fd8de162eccd23ee82237fb8deb14be2b81")
}

// UnpackEscrowProfileNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EscrowProfileNotSet()
func (walletPayments *WalletPayments) UnpackEscrowProfileNotSetError(raw []byte) (*WalletPaymentsEscrowProfileNotSet, error) {
	out := new(WalletPaymentsEscrowProfileNotSet)
	if err := walletPayments.abi.UnpackIntoInterface(out, "EscrowProfileNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsEscrowsNotEnabled represents a EscrowsNotEnabled error raised by the WalletPayments contract.
type WalletPaymentsEscrowsNotEnabled struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error EscrowsNotEnabled()
func WalletPaymentsEscrowsNotEnabledErrorID() common.Hash {
	return common.HexToHash("0x5d2d740553ca80cb27f3cc4dee972ef687c7bc820f922e4997edb5d8e2207311")
}

// UnpackEscrowsNotEnabledError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error EscrowsNotEnabled()
func (walletPayments *WalletPayments) UnpackEscrowsNotEnabledError(raw []byte) (*WalletPaymentsEscrowsNotEnabled, error) {
	out := new(WalletPaymentsEscrowsNotEnabled)
	if err := walletPayments.abi.UnpackIntoInterface(out, "EscrowsNotEnabled", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsFdc2ProofPolicyNotSet represents a Fdc2ProofPolicyNotSet error raised by the WalletPayments contract.
type WalletPaymentsFdc2ProofPolicyNotSet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error Fdc2ProofPolicyNotSet()
func WalletPaymentsFdc2ProofPolicyNotSetErrorID() common.Hash {
	return common.HexToHash("0x9ef6c9f6be455405f6672cdabe85cbceca0bacdbcddf8597308ec96c9e1c1a1c")
}

// UnpackFdc2ProofPolicyNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error Fdc2ProofPolicyNotSet()
func (walletPayments *WalletPayments) UnpackFdc2ProofPolicyNotSetError(raw []byte) (*WalletPaymentsFdc2ProofPolicyNotSet, error) {
	out := new(WalletPaymentsFdc2ProofPolicyNotSet)
	if err := walletPayments.abi.UnpackIntoInterface(out, "Fdc2ProofPolicyNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsFeeProofRangeInvalid represents a FeeProofRangeInvalid error raised by the WalletPayments contract.
type WalletPaymentsFeeProofRangeInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeProofRangeInvalid()
func WalletPaymentsFeeProofRangeInvalidErrorID() common.Hash {
	return common.HexToHash("0xa1e9ddc3af3e7b5dd57101eacf859d4de21f5c68ae832a551e1b048cb46f2890")
}

// UnpackFeeProofRangeInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeProofRangeInvalid()
func (walletPayments *WalletPayments) UnpackFeeProofRangeInvalidError(raw []byte) (*WalletPaymentsFeeProofRangeInvalid, error) {
	out := new(WalletPaymentsFeeProofRangeInvalid)
	if err := walletPayments.abi.UnpackIntoInterface(out, "FeeProofRangeInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsFeeScheduleConfigNotSet represents a FeeScheduleConfigNotSet error raised by the WalletPayments contract.
type WalletPaymentsFeeScheduleConfigNotSet struct {
	SourceId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeScheduleConfigNotSet(bytes32 sourceId)
func WalletPaymentsFeeScheduleConfigNotSetErrorID() common.Hash {
	return common.HexToHash("0x42b58a254883973ca2465c0acbab0b789fd25da4ab5f1bea80bbfd12320073d5")
}

// UnpackFeeScheduleConfigNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeScheduleConfigNotSet(bytes32 sourceId)
func (walletPayments *WalletPayments) UnpackFeeScheduleConfigNotSetError(raw []byte) (*WalletPaymentsFeeScheduleConfigNotSet, error) {
	out := new(WalletPaymentsFeeScheduleConfigNotSet)
	if err := walletPayments.abi.UnpackIntoInterface(out, "FeeScheduleConfigNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsFeeScheduleNotSet represents a FeeScheduleNotSet error raised by the WalletPayments contract.
type WalletPaymentsFeeScheduleNotSet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeScheduleNotSet()
func WalletPaymentsFeeScheduleNotSetErrorID() common.Hash {
	return common.HexToHash("0x4f47ce423907d3f49d5eb099e6293bf6be71471950c89ba028907cabdaae2e73")
}

// UnpackFeeScheduleNotSetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeScheduleNotSet()
func (walletPayments *WalletPayments) UnpackFeeScheduleNotSetError(raw []byte) (*WalletPaymentsFeeScheduleNotSet, error) {
	out := new(WalletPaymentsFeeScheduleNotSet)
	if err := walletPayments.abi.UnpackIntoInterface(out, "FeeScheduleNotSet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsFeeSchedulesUnsupported represents a FeeSchedulesUnsupported error raised by the WalletPayments contract.
type WalletPaymentsFeeSchedulesUnsupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FeeSchedulesUnsupported()
func WalletPaymentsFeeSchedulesUnsupportedErrorID() common.Hash {
	return common.HexToHash("0x25892736a6635cda86b757cb19ca56229c54b5286e36fbda461bc67b37adbd4d")
}

// UnpackFeeSchedulesUnsupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FeeSchedulesUnsupported()
func (walletPayments *WalletPayments) UnpackFeeSchedulesUnsupportedError(raw []byte) (*WalletPaymentsFeeSchedulesUnsupported, error) {
	out := new(WalletPaymentsFeeSchedulesUnsupported)
	if err := walletPayments.abi.UnpackIntoInterface(out, "FeeSchedulesUnsupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsGovernedAddressZero represents a GovernedAddressZero error raised by the WalletPayments contract.
type WalletPaymentsGovernedAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAddressZero()
func WalletPaymentsGovernedAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x695725f6b429bac0cc5a2fa88214e31520e0c69f2fb0e2fb881a7130a564f49f")
}

// UnpackGovernedAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAddressZero()
func (walletPayments *WalletPayments) UnpackGovernedAddressZeroError(raw []byte) (*WalletPaymentsGovernedAddressZero, error) {
	out := new(WalletPaymentsGovernedAddressZero)
	if err := walletPayments.abi.UnpackIntoInterface(out, "GovernedAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsGovernedAlreadyInitialized represents a GovernedAlreadyInitialized error raised by the WalletPayments contract.
type WalletPaymentsGovernedAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAlreadyInitialized()
func WalletPaymentsGovernedAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0xb2f0162156d83828197776d56bbfba0a7a9d5936fad70f0a3014ae18cc7aa33c")
}

// UnpackGovernedAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAlreadyInitialized()
func (walletPayments *WalletPayments) UnpackGovernedAlreadyInitializedError(raw []byte) (*WalletPaymentsGovernedAlreadyInitialized, error) {
	out := new(WalletPaymentsGovernedAlreadyInitialized)
	if err := walletPayments.abi.UnpackIntoInterface(out, "GovernedAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsGraceClosed represents a GraceClosed error raised by the WalletPayments contract.
type WalletPaymentsGraceClosed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GraceClosed()
func WalletPaymentsGraceClosedErrorID() common.Hash {
	return common.HexToHash("0x5225b90c5c94194cb30caa07ef2e608a497bee5eecc30041095644bf6a5f72a5")
}

// UnpackGraceClosedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GraceClosed()
func (walletPayments *WalletPayments) UnpackGraceClosedError(raw []byte) (*WalletPaymentsGraceClosed, error) {
	out := new(WalletPaymentsGraceClosed)
	if err := walletPayments.abi.UnpackIntoInterface(out, "GraceClosed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsGraceOpen represents a GraceOpen error raised by the WalletPayments contract.
type WalletPaymentsGraceOpen struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GraceOpen()
func WalletPaymentsGraceOpenErrorID() common.Hash {
	return common.HexToHash("0xc23ee5e6fdb39cfdf31c02ef3ae230bb910791d369c7cc01535ccdff457dd678")
}

// UnpackGraceOpenError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GraceOpen()
func (walletPayments *WalletPayments) UnpackGraceOpenError(raw []byte) (*WalletPaymentsGraceOpen, error) {
	out := new(WalletPaymentsGraceOpen)
	if err := walletPayments.abi.UnpackIntoInterface(out, "GraceOpen", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInsufficientFee represents a InsufficientFee error raised by the WalletPayments contract.
type WalletPaymentsInsufficientFee struct {
	Required *big.Int
	Provided *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientFee(uint256 required, uint256 provided)
func WalletPaymentsInsufficientFeeErrorID() common.Hash {
	return common.HexToHash("0xa458261b25e6c4899486bdd60e6932f9a312c90ca996e36fe276f5ae4112202c")
}

// UnpackInsufficientFeeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientFee(uint256 required, uint256 provided)
func (walletPayments *WalletPayments) UnpackInsufficientFeeError(raw []byte) (*WalletPaymentsInsufficientFee, error) {
	out := new(WalletPaymentsInsufficientFee)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InsufficientFee", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInsufficientTeeSignatures represents a InsufficientTeeSignatures error raised by the WalletPayments contract.
type WalletPaymentsInsufficientTeeSignatures struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InsufficientTeeSignatures()
func WalletPaymentsInsufficientTeeSignaturesErrorID() common.Hash {
	return common.HexToHash("0xb2a597bd17627405b809a6f19e5c9f57958cdaede6356b492efd5149031b57ea")
}

// UnpackInsufficientTeeSignaturesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InsufficientTeeSignatures()
func (walletPayments *WalletPayments) UnpackInsufficientTeeSignaturesError(raw []byte) (*WalletPaymentsInsufficientTeeSignatures, error) {
	out := new(WalletPaymentsInsufficientTeeSignatures)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InsufficientTeeSignatures", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidAttestation represents a InvalidAttestation error raised by the WalletPayments contract.
type WalletPaymentsInvalidAttestation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAttestation()
func WalletPaymentsInvalidAttestationErrorID() common.Hash {
	return common.HexToHash("0xbd8ba84d4af301e1b110d1f93f8971934a3ba3afdf4200e7f884f76f2dd27077")
}

// UnpackInvalidAttestationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAttestation()
func (walletPayments *WalletPayments) UnpackInvalidAttestationError(raw []byte) (*WalletPaymentsInvalidAttestation, error) {
	out := new(WalletPaymentsInvalidAttestation)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidAttestation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidBtcConsensusParams represents a InvalidBtcConsensusParams error raised by the WalletPayments contract.
type WalletPaymentsInvalidBtcConsensusParams struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidBtcConsensusParams()
func WalletPaymentsInvalidBtcConsensusParamsErrorID() common.Hash {
	return common.HexToHash("0x7714190a43b48f2651073c4b3c8198eb8101174b57d155c766a671f84cc06d67")
}

// UnpackInvalidBtcConsensusParamsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidBtcConsensusParams()
func (walletPayments *WalletPayments) UnpackInvalidBtcConsensusParamsError(raw []byte) (*WalletPaymentsInvalidBtcConsensusParams, error) {
	out := new(WalletPaymentsInvalidBtcConsensusParams)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidBtcConsensusParams", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidBtcEscrowProfile represents a InvalidBtcEscrowProfile error raised by the WalletPayments contract.
type WalletPaymentsInvalidBtcEscrowProfile struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidBtcEscrowProfile()
func WalletPaymentsInvalidBtcEscrowProfileErrorID() common.Hash {
	return common.HexToHash("0x6749673c9d6a337ea7fe5a5c90943ecea1fbc1d88566e5a82460dc89b4b5845f")
}

// UnpackInvalidBtcEscrowProfileError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidBtcEscrowProfile()
func (walletPayments *WalletPayments) UnpackInvalidBtcEscrowProfileError(raw []byte) (*WalletPaymentsInvalidBtcEscrowProfile, error) {
	out := new(WalletPaymentsInvalidBtcEscrowProfile)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidBtcEscrowProfile", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidBtcEscrowTerms represents a InvalidBtcEscrowTerms error raised by the WalletPayments contract.
type WalletPaymentsInvalidBtcEscrowTerms struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidBtcEscrowTerms()
func WalletPaymentsInvalidBtcEscrowTermsErrorID() common.Hash {
	return common.HexToHash("0x0e62b7a97a34a7feeaf7ea3b689dcbda4b7505ff77cd26adc5f60fadb3795f6f")
}

// UnpackInvalidBtcEscrowTermsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidBtcEscrowTerms()
func (walletPayments *WalletPayments) UnpackInvalidBtcEscrowTermsError(raw []byte) (*WalletPaymentsInvalidBtcEscrowTerms, error) {
	out := new(WalletPaymentsInvalidBtcEscrowTerms)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidBtcEscrowTerms", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidBtcWalletParams represents a InvalidBtcWalletParams error raised by the WalletPayments contract.
type WalletPaymentsInvalidBtcWalletParams struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidBtcWalletParams()
func WalletPaymentsInvalidBtcWalletParamsErrorID() common.Hash {
	return common.HexToHash("0x6593c2e56ad40b0802d4648405b98a583ee98609ef2cbb8c94ab25288d4cc3ac")
}

// UnpackInvalidBtcWalletParamsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidBtcWalletParams()
func (walletPayments *WalletPayments) UnpackInvalidBtcWalletParamsError(raw []byte) (*WalletPaymentsInvalidBtcWalletParams, error) {
	out := new(WalletPaymentsInvalidBtcWalletParams)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidBtcWalletParams", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidCspSourceSettings represents a InvalidCspSourceSettings error raised by the WalletPayments contract.
type WalletPaymentsInvalidCspSourceSettings struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCspSourceSettings()
func WalletPaymentsInvalidCspSourceSettingsErrorID() common.Hash {
	return common.HexToHash("0xc70606aa293575c950ac5bb28d55a04cbd561a56b8604584e69d1fc80fcb3f39")
}

// UnpackInvalidCspSourceSettingsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCspSourceSettings()
func (walletPayments *WalletPayments) UnpackInvalidCspSourceSettingsError(raw []byte) (*WalletPaymentsInvalidCspSourceSettings, error) {
	out := new(WalletPaymentsInvalidCspSourceSettings)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidCspSourceSettings", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidFdc2ProofPolicy represents a InvalidFdc2ProofPolicy error raised by the WalletPayments contract.
type WalletPaymentsInvalidFdc2ProofPolicy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFdc2ProofPolicy()
func WalletPaymentsInvalidFdc2ProofPolicyErrorID() common.Hash {
	return common.HexToHash("0xd5704bd8e9dd304e933815944ed2d6eb2e8c02300c0b2231ff76216d488f5bbf")
}

// UnpackInvalidFdc2ProofPolicyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFdc2ProofPolicy()
func (walletPayments *WalletPayments) UnpackInvalidFdc2ProofPolicyError(raw []byte) (*WalletPaymentsInvalidFdc2ProofPolicy, error) {
	out := new(WalletPaymentsInvalidFdc2ProofPolicy)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidFdc2ProofPolicy", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidFeeDelay represents a InvalidFeeDelay error raised by the WalletPayments contract.
type WalletPaymentsInvalidFeeDelay struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeDelay(uint256 index)
func WalletPaymentsInvalidFeeDelayErrorID() common.Hash {
	return common.HexToHash("0x2a9c9c04432dee2444e275db8772f3255e5685176b3891e4a344c3c7c4854620")
}

// UnpackInvalidFeeDelayError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeDelay(uint256 index)
func (walletPayments *WalletPayments) UnpackInvalidFeeDelayError(raw []byte) (*WalletPaymentsInvalidFeeDelay, error) {
	out := new(WalletPaymentsInvalidFeeDelay)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidFeeDelay", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidFeeFactor represents a InvalidFeeFactor error raised by the WalletPayments contract.
type WalletPaymentsInvalidFeeFactor struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeFactor(uint256 index)
func WalletPaymentsInvalidFeeFactorErrorID() common.Hash {
	return common.HexToHash("0x89619400c26b8355799d991b2b4ac022351c0c4f21a4c0f3b4ee2ef70d0c8737")
}

// UnpackInvalidFeeFactorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeFactor(uint256 index)
func (walletPayments *WalletPayments) UnpackInvalidFeeFactorError(raw []byte) (*WalletPaymentsInvalidFeeFactor, error) {
	out := new(WalletPaymentsInvalidFeeFactor)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidFeeFactor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidFeeScheduleConfig represents a InvalidFeeScheduleConfig error raised by the WalletPayments contract.
type WalletPaymentsInvalidFeeScheduleConfig struct {
	SourceId        [32]byte
	MaxSchedules    uint8
	MaxDelaySeconds uint16
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidFeeScheduleConfig(bytes32 sourceId, uint8 maxSchedules, uint16 maxDelaySeconds)
func WalletPaymentsInvalidFeeScheduleConfigErrorID() common.Hash {
	return common.HexToHash("0x3749abf55999564327d44208b64402fca0d55fb8e7c8c042042af79acd07d579")
}

// UnpackInvalidFeeScheduleConfigError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidFeeScheduleConfig(bytes32 sourceId, uint8 maxSchedules, uint16 maxDelaySeconds)
func (walletPayments *WalletPayments) UnpackInvalidFeeScheduleConfigError(raw []byte) (*WalletPaymentsInvalidFeeScheduleConfig, error) {
	out := new(WalletPaymentsInvalidFeeScheduleConfig)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidFeeScheduleConfig", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidGrace represents a InvalidGrace error raised by the WalletPayments contract.
type WalletPaymentsInvalidGrace struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidGrace()
func WalletPaymentsInvalidGraceErrorID() common.Hash {
	return common.HexToHash("0x795ee5af27b0e919c848f6b99bf2048a0a5363368438175c8e937504ce06cc87")
}

// UnpackInvalidGraceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidGrace()
func (walletPayments *WalletPayments) UnpackInvalidGraceError(raw []byte) (*WalletPaymentsInvalidGrace, error) {
	out := new(WalletPaymentsInvalidGrace)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidGrace", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidPaymentCount represents a InvalidPaymentCount error raised by the WalletPayments contract.
type WalletPaymentsInvalidPaymentCount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPaymentCount()
func WalletPaymentsInvalidPaymentCountErrorID() common.Hash {
	return common.HexToHash("0x3ed4fa969d6a398b18db6b58c7759cf2310f1e51661f974f8881fec01d1bdad0")
}

// UnpackInvalidPaymentCountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPaymentCount()
func (walletPayments *WalletPayments) UnpackInvalidPaymentCountError(raw []byte) (*WalletPaymentsInvalidPaymentCount, error) {
	out := new(WalletPaymentsInvalidPaymentCount)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidPaymentCount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidPaymentId represents a InvalidPaymentId error raised by the WalletPayments contract.
type WalletPaymentsInvalidPaymentId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPaymentId()
func WalletPaymentsInvalidPaymentIdErrorID() common.Hash {
	return common.HexToHash("0xb11be8c4d37d5cf7116085b4f06afe80345db997f69552f4e6a826a12eea0c0f")
}

// UnpackInvalidPaymentIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPaymentId()
func (walletPayments *WalletPayments) UnpackInvalidPaymentIdError(raw []byte) (*WalletPaymentsInvalidPaymentId, error) {
	out := new(WalletPaymentsInvalidPaymentId)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidPaymentId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidPaymentInstructionCount represents a InvalidPaymentInstructionCount error raised by the WalletPayments contract.
type WalletPaymentsInvalidPaymentInstructionCount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPaymentInstructionCount()
func WalletPaymentsInvalidPaymentInstructionCountErrorID() common.Hash {
	return common.HexToHash("0xb944aca0ae5b99f4666f05d50060098c29717bc0c0e1c515201a3cee72b4f5c3")
}

// UnpackInvalidPaymentInstructionCountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPaymentInstructionCount()
func (walletPayments *WalletPayments) UnpackInvalidPaymentInstructionCountError(raw []byte) (*WalletPaymentsInvalidPaymentInstructionCount, error) {
	out := new(WalletPaymentsInvalidPaymentInstructionCount)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidPaymentInstructionCount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidProof represents a InvalidProof error raised by the WalletPayments contract.
type WalletPaymentsInvalidProof struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidProof()
func WalletPaymentsInvalidProofErrorID() common.Hash {
	return common.HexToHash("0x09bde339c6b182be216ee7ef8ccff6338c6ef7993445216112ae575c5438fd27")
}

// UnpackInvalidProofError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidProof()
func (walletPayments *WalletPayments) UnpackInvalidProofError(raw []byte) (*WalletPaymentsInvalidProof, error) {
	out := new(WalletPaymentsInvalidProof)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidProof", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidRecipientAddress represents a InvalidRecipientAddress error raised by the WalletPayments contract.
type WalletPaymentsInvalidRecipientAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRecipientAddress()
func WalletPaymentsInvalidRecipientAddressErrorID() common.Hash {
	return common.HexToHash("0x44d99fea8d272f0b82a2cdf27d16958dace6a85ad6302ff443201a39bfc4820a")
}

// UnpackInvalidRecipientAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRecipientAddress()
func (walletPayments *WalletPayments) UnpackInvalidRecipientAddressError(raw []byte) (*WalletPaymentsInvalidRecipientAddress, error) {
	out := new(WalletPaymentsInvalidRecipientAddress)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidRecipientAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidRequestBody represents a InvalidRequestBody error raised by the WalletPayments contract.
type WalletPaymentsInvalidRequestBody struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRequestBody()
func WalletPaymentsInvalidRequestBodyErrorID() common.Hash {
	return common.HexToHash("0xfe524bcc2b70aad90ebe31fd595732afd3bed1387a6e1141569409a40554173c")
}

// UnpackInvalidRequestBodyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRequestBody()
func (walletPayments *WalletPayments) UnpackInvalidRequestBodyError(raw []byte) (*WalletPaymentsInvalidRequestBody, error) {
	out := new(WalletPaymentsInvalidRequestBody)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidRequestBody", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsInvalidXrpEscrowTerms represents a InvalidXrpEscrowTerms error raised by the WalletPayments contract.
type WalletPaymentsInvalidXrpEscrowTerms struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidXrpEscrowTerms()
func WalletPaymentsInvalidXrpEscrowTermsErrorID() common.Hash {
	return common.HexToHash("0xf894bed3b3e2cb63440dc3056c8c9188e724bec004118c47740ee1bc4784335c")
}

// UnpackInvalidXrpEscrowTermsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidXrpEscrowTerms()
func (walletPayments *WalletPayments) UnpackInvalidXrpEscrowTermsError(raw []byte) (*WalletPaymentsInvalidXrpEscrowTerms, error) {
	out := new(WalletPaymentsInvalidXrpEscrowTerms)
	if err := walletPayments.abi.UnpackIntoInterface(out, "InvalidXrpEscrowTerms", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsLaneMismatch represents a LaneMismatch error raised by the WalletPayments contract.
type WalletPaymentsLaneMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LaneMismatch()
func WalletPaymentsLaneMismatchErrorID() common.Hash {
	return common.HexToHash("0x3829d65ed0e236edac9b604bc50cf4ec2c45c9c262deeed1a585adcfd5015670")
}

// UnpackLaneMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LaneMismatch()
func (walletPayments *WalletPayments) UnpackLaneMismatchError(raw []byte) (*WalletPaymentsLaneMismatch, error) {
	out := new(WalletPaymentsLaneMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "LaneMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsLeaderExists represents a LeaderExists error raised by the WalletPayments contract.
type WalletPaymentsLeaderExists struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LeaderExists()
func WalletPaymentsLeaderExistsErrorID() common.Hash {
	return common.HexToHash("0x0cbd735d7c6164a42a977a31bae66f43c13c80b2ae159605cdfd6a6bf9d6cd52")
}

// UnpackLeaderExistsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LeaderExists()
func (walletPayments *WalletPayments) UnpackLeaderExistsError(raw []byte) (*WalletPaymentsLeaderExists, error) {
	out := new(WalletPaymentsLeaderExists)
	if err := walletPayments.abi.UnpackIntoInterface(out, "LeaderExists", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsLengthsMismatch represents a LengthsMismatch error raised by the WalletPayments contract.
type WalletPaymentsLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func WalletPaymentsLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (walletPayments *WalletPayments) UnpackLengthsMismatchError(raw []byte) (*WalletPaymentsLengthsMismatch, error) {
	out := new(WalletPaymentsLengthsMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsMaxFeeTooLarge represents a MaxFeeTooLarge error raised by the WalletPayments contract.
type WalletPaymentsMaxFeeTooLarge struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MaxFeeTooLarge()
func WalletPaymentsMaxFeeTooLargeErrorID() common.Hash {
	return common.HexToHash("0x26361e6d5086b93e2ea6abfd6677d041dd6a5fe6efe03def26c63567195200b6")
}

// UnpackMaxFeeTooLargeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MaxFeeTooLarge()
func (walletPayments *WalletPayments) UnpackMaxFeeTooLargeError(raw []byte) (*WalletPaymentsMaxFeeTooLarge, error) {
	out := new(WalletPaymentsMaxFeeTooLarge)
	if err := walletPayments.abi.UnpackIntoInterface(out, "MaxFeeTooLarge", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsMaxFeeZero represents a MaxFeeZero error raised by the WalletPayments contract.
type WalletPaymentsMaxFeeZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error MaxFeeZero()
func WalletPaymentsMaxFeeZeroErrorID() common.Hash {
	return common.HexToHash("0xf63a2fe7757bc551db9e8a7f4a473bfafe9300a9d1012f9947ec01e434f38cef")
}

// UnpackMaxFeeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error MaxFeeZero()
func (walletPayments *WalletPayments) UnpackMaxFeeZeroError(raw []byte) (*WalletPaymentsMaxFeeZero, error) {
	out := new(WalletPaymentsMaxFeeZero)
	if err := walletPayments.abi.UnpackIntoInterface(out, "MaxFeeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsNoNewAnchors represents a NoNewAnchors error raised by the WalletPayments contract.
type WalletPaymentsNoNewAnchors struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoNewAnchors()
func WalletPaymentsNoNewAnchorsErrorID() common.Hash {
	return common.HexToHash("0xb01793c2dd316a7313f6307f05a3be47b867e01e1f48b301608f1f92885c2d6d")
}

// UnpackNoNewAnchorsError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoNewAnchors()
func (walletPayments *WalletPayments) UnpackNoNewAnchorsError(raw []byte) (*WalletPaymentsNoNewAnchors, error) {
	out := new(WalletPaymentsNoNewAnchors)
	if err := walletPayments.abi.UnpackIntoInterface(out, "NoNewAnchors", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsNoPendingReset represents a NoPendingReset error raised by the WalletPayments contract.
type WalletPaymentsNoPendingReset struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NoPendingReset()
func WalletPaymentsNoPendingResetErrorID() common.Hash {
	return common.HexToHash("0x01ab8ba0e78e139eebb25117033fbc0dc4040db8782a30b08e17d781fc4b3fe8")
}

// UnpackNoPendingResetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NoPendingReset()
func (walletPayments *WalletPayments) UnpackNoPendingResetError(raw []byte) (*WalletPaymentsNoPendingReset, error) {
	out := new(WalletPaymentsNoPendingReset)
	if err := walletPayments.abi.UnpackIntoInterface(out, "NoPendingReset", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsNonceMismatch represents a NonceMismatch error raised by the WalletPayments contract.
type WalletPaymentsNonceMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NonceMismatch()
func WalletPaymentsNonceMismatchErrorID() common.Hash {
	return common.HexToHash("0xe112ed91bb74dd69dfe256ccb69e567ab7aac0f92757a06850be5dadb6d2c097")
}

// UnpackNonceMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NonceMismatch()
func (walletPayments *WalletPayments) UnpackNonceMismatchError(raw []byte) (*WalletPaymentsNonceMismatch, error) {
	out := new(WalletPaymentsNonceMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "NonceMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsNotEligible represents a NotEligible error raised by the WalletPayments contract.
type WalletPaymentsNotEligible struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotEligible()
func WalletPaymentsNotEligibleErrorID() common.Hash {
	return common.HexToHash("0xf8eb54de284c97bfa6191d3cb5e2a0cef222e6ce29e2e7ed2edafab612c84e14")
}

// UnpackNotEligibleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotEligible()
func (walletPayments *WalletPayments) UnpackNotEligibleError(raw []byte) (*WalletPaymentsNotEligible, error) {
	out := new(WalletPaymentsNotEligible)
	if err := walletPayments.abi.UnpackIntoInterface(out, "NotEligible", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsNotLaneTip represents a NotLaneTip error raised by the WalletPayments contract.
type WalletPaymentsNotLaneTip struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotLaneTip()
func WalletPaymentsNotLaneTipErrorID() common.Hash {
	return common.HexToHash("0xf0c399505c67ab999b8c337188e89cd5e202fabc21d9c1af371c17ae4ece5a1e")
}

// UnpackNotLaneTipError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotLaneTip()
func (walletPayments *WalletPayments) UnpackNotLaneTipError(raw []byte) (*WalletPaymentsNotLaneTip, error) {
	out := new(WalletPaymentsNotLaneTip)
	if err := walletPayments.abi.UnpackIntoInterface(out, "NotLaneTip", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsNotSettled represents a NotSettled error raised by the WalletPayments contract.
type WalletPaymentsNotSettled struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotSettled()
func WalletPaymentsNotSettledErrorID() common.Hash {
	return common.HexToHash("0xba329a9ba88ea34e96b72c575515344032eb5ca11490926c149f35dc9dd0d722")
}

// UnpackNotSettledError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotSettled()
func (walletPayments *WalletPayments) UnpackNotSettledError(raw []byte) (*WalletPaymentsNotSettled, error) {
	out := new(WalletPaymentsNotSettled)
	if err := walletPayments.abi.UnpackIntoInterface(out, "NotSettled", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsNothingToSettle represents a NothingToSettle error raised by the WalletPayments contract.
type WalletPaymentsNothingToSettle struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NothingToSettle()
func WalletPaymentsNothingToSettleErrorID() common.Hash {
	return common.HexToHash("0x618104e821a31c94c175b9bf52827d750880c5f9e96152367da18c1fe40d3ee0")
}

// UnpackNothingToSettleError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NothingToSettle()
func (walletPayments *WalletPayments) UnpackNothingToSettleError(raw []byte) (*WalletPaymentsNothingToSettle, error) {
	out := new(WalletPaymentsNothingToSettle)
	if err := walletPayments.abi.UnpackIntoInterface(out, "NothingToSettle", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsNullifiedPaymentsUnsupported represents a NullifiedPaymentsUnsupported error raised by the WalletPayments contract.
type WalletPaymentsNullifiedPaymentsUnsupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NullifiedPaymentsUnsupported()
func WalletPaymentsNullifiedPaymentsUnsupportedErrorID() common.Hash {
	return common.HexToHash("0xe9943b24a570c58d8bb59eb634bc71f0f88d10b0dc9a2873a655f0769cd8d5fc")
}

// UnpackNullifiedPaymentsUnsupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NullifiedPaymentsUnsupported()
func (walletPayments *WalletPayments) UnpackNullifiedPaymentsUnsupportedError(raw []byte) (*WalletPaymentsNullifiedPaymentsUnsupported, error) {
	out := new(WalletPaymentsNullifiedPaymentsUnsupported)
	if err := walletPayments.abi.UnpackIntoInterface(out, "NullifiedPaymentsUnsupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsOnlyAuthorizationAddress represents a OnlyAuthorizationAddress error raised by the WalletPayments contract.
type WalletPaymentsOnlyAuthorizationAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyAuthorizationAddress()
func WalletPaymentsOnlyAuthorizationAddressErrorID() common.Hash {
	return common.HexToHash("0xdc65e71c09e77be7539dadee1dc86849e01c3e995f0c8d49e625476c2a1b1ed5")
}

// UnpackOnlyAuthorizationAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyAuthorizationAddress()
func (walletPayments *WalletPayments) UnpackOnlyAuthorizationAddressError(raw []byte) (*WalletPaymentsOnlyAuthorizationAddress, error) {
	out := new(WalletPaymentsOnlyAuthorizationAddress)
	if err := walletPayments.abi.UnpackIntoInterface(out, "OnlyAuthorizationAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsOnlyExecutor represents a OnlyExecutor error raised by the WalletPayments contract.
type WalletPaymentsOnlyExecutor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExecutor()
func WalletPaymentsOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x7fb6be02aaf40e420b0e7eb67e43a3feee949e8932a3c64e33b591a436670841")
}

// UnpackOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExecutor()
func (walletPayments *WalletPayments) UnpackOnlyExecutorError(raw []byte) (*WalletPaymentsOnlyExecutor, error) {
	out := new(WalletPaymentsOnlyExecutor)
	if err := walletPayments.abi.UnpackIntoInterface(out, "OnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsOnlyGovernance represents a OnlyGovernance error raised by the WalletPayments contract.
type WalletPaymentsOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func WalletPaymentsOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (walletPayments *WalletPayments) UnpackOnlyGovernanceError(raw []byte) (*WalletPaymentsOnlyGovernance, error) {
	out := new(WalletPaymentsOnlyGovernance)
	if err := walletPayments.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the WalletPayments contract.
type WalletPaymentsOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func WalletPaymentsOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (walletPayments *WalletPayments) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*WalletPaymentsOnlyProductionOrPausedStatus, error) {
	out := new(WalletPaymentsOnlyProductionOrPausedStatus)
	if err := walletPayments.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsOnlyProjectOwner represents a OnlyProjectOwner error raised by the WalletPayments contract.
type WalletPaymentsOnlyProjectOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProjectOwner()
func WalletPaymentsOnlyProjectOwnerErrorID() common.Hash {
	return common.HexToHash("0x9e140988363c55adef362d3d92091fcd2495ebb514cc5070cdc0816e6a555dd6")
}

// UnpackOnlyProjectOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProjectOwner()
func (walletPayments *WalletPayments) UnpackOnlyProjectOwnerError(raw []byte) (*WalletPaymentsOnlyProjectOwner, error) {
	out := new(WalletPaymentsOnlyProjectOwner)
	if err := walletPayments.abi.UnpackIntoInterface(out, "OnlyProjectOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsOnlySystemExtensionId represents a OnlySystemExtensionId error raised by the WalletPayments contract.
type WalletPaymentsOnlySystemExtensionId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemExtensionId()
func WalletPaymentsOnlySystemExtensionIdErrorID() common.Hash {
	return common.HexToHash("0xbd49094d5a86ee3f3c4b39274f5f98391dde2cd16618a917e121333ac3d43ee4")
}

// UnpackOnlySystemExtensionIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemExtensionId()
func (walletPayments *WalletPayments) UnpackOnlySystemExtensionIdError(raw []byte) (*WalletPaymentsOnlySystemExtensionId, error) {
	out := new(WalletPaymentsOnlySystemExtensionId)
	if err := walletPayments.abi.UnpackIntoInterface(out, "OnlySystemExtensionId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsOnlyWalletOwner represents a OnlyWalletOwner error raised by the WalletPayments contract.
type WalletPaymentsOnlyWalletOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyWalletOwner()
func WalletPaymentsOnlyWalletOwnerErrorID() common.Hash {
	return common.HexToHash("0x3132d62dfddd1d8b6a9dd805001453bbb6a84b746ebe51fd650df1a6704a8143")
}

// UnpackOnlyWalletOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyWalletOwner()
func (walletPayments *WalletPayments) UnpackOnlyWalletOwnerError(raw []byte) (*WalletPaymentsOnlyWalletOwner, error) {
	out := new(WalletPaymentsOnlyWalletOwner)
	if err := walletPayments.abi.UnpackIntoInterface(out, "OnlyWalletOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsOperationAlreadyNullified represents a OperationAlreadyNullified error raised by the WalletPayments contract.
type WalletPaymentsOperationAlreadyNullified struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationAlreadyNullified()
func WalletPaymentsOperationAlreadyNullifiedErrorID() common.Hash {
	return common.HexToHash("0xf6900be86e26247f674cfb877a1310eba3675fad2fe791ebf2b2ed70f62b1a8f")
}

// UnpackOperationAlreadyNullifiedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationAlreadyNullified()
func (walletPayments *WalletPayments) UnpackOperationAlreadyNullifiedError(raw []byte) (*WalletPaymentsOperationAlreadyNullified, error) {
	out := new(WalletPaymentsOperationAlreadyNullified)
	if err := walletPayments.abi.UnpackIntoInterface(out, "OperationAlreadyNullified", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsOperationNotSupported represents a OperationNotSupported error raised by the WalletPayments contract.
type WalletPaymentsOperationNotSupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationNotSupported()
func WalletPaymentsOperationNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x29a270f5e1b8b965905c03a15e79ccfc8e66183fd2489866b39ca3fd8bfc0ae6")
}

// UnpackOperationNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationNotSupported()
func (walletPayments *WalletPayments) UnpackOperationNotSupportedError(raw []byte) (*WalletPaymentsOperationNotSupported, error) {
	out := new(WalletPaymentsOperationNotSupported)
	if err := walletPayments.abi.UnpackIntoInterface(out, "OperationNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsOperationStatusMismatch represents a OperationStatusMismatch error raised by the WalletPayments contract.
type WalletPaymentsOperationStatusMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OperationStatusMismatch()
func WalletPaymentsOperationStatusMismatchErrorID() common.Hash {
	return common.HexToHash("0x74ecfcf677f879e6763d3be2779894e475943e33f9018309a225ab947e946a32")
}

// UnpackOperationStatusMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OperationStatusMismatch()
func (walletPayments *WalletPayments) UnpackOperationStatusMismatchError(raw []byte) (*WalletPaymentsOperationStatusMismatch, error) {
	out := new(WalletPaymentsOperationStatusMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "OperationStatusMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsOverrideAlreadyPending represents a OverrideAlreadyPending error raised by the WalletPayments contract.
type WalletPaymentsOverrideAlreadyPending struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OverrideAlreadyPending()
func WalletPaymentsOverrideAlreadyPendingErrorID() common.Hash {
	return common.HexToHash("0x8084e8c1eb0ff3429efc9f3a89c15e41c001262ec4790a03db4b08ffcacf3eff")
}

// UnpackOverrideAlreadyPendingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OverrideAlreadyPending()
func (walletPayments *WalletPayments) UnpackOverrideAlreadyPendingError(raw []byte) (*WalletPaymentsOverrideAlreadyPending, error) {
	out := new(WalletPaymentsOverrideAlreadyPending)
	if err := walletPayments.abi.UnpackIntoInterface(out, "OverrideAlreadyPending", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsPaymentAmountTooLow represents a PaymentAmountTooLow error raised by the WalletPayments contract.
type WalletPaymentsPaymentAmountTooLow struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentAmountTooLow()
func WalletPaymentsPaymentAmountTooLowErrorID() common.Hash {
	return common.HexToHash("0xb61990da5c4478b17cca22703360c6530a540819b8448b828580634b4051b23b")
}

// UnpackPaymentAmountTooLowError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentAmountTooLow()
func (walletPayments *WalletPayments) UnpackPaymentAmountTooLowError(raw []byte) (*WalletPaymentsPaymentAmountTooLow, error) {
	out := new(WalletPaymentsPaymentAmountTooLow)
	if err := walletPayments.abi.UnpackIntoInterface(out, "PaymentAmountTooLow", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsPaymentAmountZero represents a PaymentAmountZero error raised by the WalletPayments contract.
type WalletPaymentsPaymentAmountZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentAmountZero()
func WalletPaymentsPaymentAmountZeroErrorID() common.Hash {
	return common.HexToHash("0xa74a1c28a64fc70e7d718ae1b3af31f436e1c55b3754d607ef7ec829ec74034c")
}

// UnpackPaymentAmountZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentAmountZero()
func (walletPayments *WalletPayments) UnpackPaymentAmountZeroError(raw []byte) (*WalletPaymentsPaymentAmountZero, error) {
	out := new(WalletPaymentsPaymentAmountZero)
	if err := walletPayments.abi.UnpackIntoInterface(out, "PaymentAmountZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsPaymentHashMismatch represents a PaymentHashMismatch error raised by the WalletPayments contract.
type WalletPaymentsPaymentHashMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentHashMismatch()
func WalletPaymentsPaymentHashMismatchErrorID() common.Hash {
	return common.HexToHash("0x85628ea5f0ea30f18d9205be9d6ae9bd4b189b3dc00aeba346ab8f4c0e25674e")
}

// UnpackPaymentHashMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentHashMismatch()
func (walletPayments *WalletPayments) UnpackPaymentHashMismatchError(raw []byte) (*WalletPaymentsPaymentHashMismatch, error) {
	out := new(WalletPaymentsPaymentHashMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "PaymentHashMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsPaymentNotSettled represents a PaymentNotSettled error raised by the WalletPayments contract.
type WalletPaymentsPaymentNotSettled struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentNotSettled()
func WalletPaymentsPaymentNotSettledErrorID() common.Hash {
	return common.HexToHash("0xbe2148e8953fdd027dbf153a84c3b47e6e84254d3144dcb87a51e45da484a184")
}

// UnpackPaymentNotSettledError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentNotSettled()
func (walletPayments *WalletPayments) UnpackPaymentNotSettledError(raw []byte) (*WalletPaymentsPaymentNotSettled, error) {
	out := new(WalletPaymentsPaymentNotSettled)
	if err := walletPayments.abi.UnpackIntoInterface(out, "PaymentNotSettled", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsPaymentRangeInvalid represents a PaymentRangeInvalid error raised by the WalletPayments contract.
type WalletPaymentsPaymentRangeInvalid struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentRangeInvalid()
func WalletPaymentsPaymentRangeInvalidErrorID() common.Hash {
	return common.HexToHash("0x69612ee8e4d0de758e4740d1979e51ca2f865edca38dd42526a6becb0f1957f0")
}

// UnpackPaymentRangeInvalidError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentRangeInvalid()
func (walletPayments *WalletPayments) UnpackPaymentRangeInvalidError(raw []byte) (*WalletPaymentsPaymentRangeInvalid, error) {
	out := new(WalletPaymentsPaymentRangeInvalid)
	if err := walletPayments.abi.UnpackIntoInterface(out, "PaymentRangeInvalid", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsProposerNotWhitelisted represents a ProposerNotWhitelisted error raised by the WalletPayments contract.
type WalletPaymentsProposerNotWhitelisted struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ProposerNotWhitelisted()
func WalletPaymentsProposerNotWhitelistedErrorID() common.Hash {
	return common.HexToHash("0x6ffe5a95d0cd81f8c9e2f050b26c3923967e546ff30112c8d16ed680e628f46d")
}

// UnpackProposerNotWhitelistedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ProposerNotWhitelisted()
func (walletPayments *WalletPayments) UnpackProposerNotWhitelistedError(raw []byte) (*WalletPaymentsProposerNotWhitelisted, error) {
	out := new(WalletPaymentsProposerNotWhitelisted)
	if err := walletPayments.abi.UnpackIntoInterface(out, "ProposerNotWhitelisted", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsReissueDeclarationMismatch represents a ReissueDeclarationMismatch error raised by the WalletPayments contract.
type WalletPaymentsReissueDeclarationMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ReissueDeclarationMismatch()
func WalletPaymentsReissueDeclarationMismatchErrorID() common.Hash {
	return common.HexToHash("0x7855cb91733bb5da347d0e234fba3097b93888226f4c7dd6d9383e1111dc50dd")
}

// UnpackReissueDeclarationMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ReissueDeclarationMismatch()
func (walletPayments *WalletPayments) UnpackReissueDeclarationMismatchError(raw []byte) (*WalletPaymentsReissueDeclarationMismatch, error) {
	out := new(WalletPaymentsReissueDeclarationMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "ReissueDeclarationMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsSourceAlreadyRegistered represents a SourceAlreadyRegistered error raised by the WalletPayments contract.
type WalletPaymentsSourceAlreadyRegistered struct {
	SourceId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceAlreadyRegistered(bytes32 sourceId)
func WalletPaymentsSourceAlreadyRegisteredErrorID() common.Hash {
	return common.HexToHash("0x2b4d4b864fb5a9bb89780f72a01a026ded28f45eb2755f97d12c174e45cb94b9")
}

// UnpackSourceAlreadyRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceAlreadyRegistered(bytes32 sourceId)
func (walletPayments *WalletPayments) UnpackSourceAlreadyRegisteredError(raw []byte) (*WalletPaymentsSourceAlreadyRegistered, error) {
	out := new(WalletPaymentsSourceAlreadyRegistered)
	if err := walletPayments.abi.UnpackIntoInterface(out, "SourceAlreadyRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsSourceChainKeyTypeMismatch represents a SourceChainKeyTypeMismatch error raised by the WalletPayments contract.
type WalletPaymentsSourceChainKeyTypeMismatch struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceChainKeyTypeMismatch(uint256 index)
func WalletPaymentsSourceChainKeyTypeMismatchErrorID() common.Hash {
	return common.HexToHash("0xf0590cf7ca832901bb48a2adef8a795f94213c707def05182ae67edd562142a6")
}

// UnpackSourceChainKeyTypeMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceChainKeyTypeMismatch(uint256 index)
func (walletPayments *WalletPayments) UnpackSourceChainKeyTypeMismatchError(raw []byte) (*WalletPaymentsSourceChainKeyTypeMismatch, error) {
	out := new(WalletPaymentsSourceChainKeyTypeMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "SourceChainKeyTypeMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsSourceDisabled represents a SourceDisabled error raised by the WalletPayments contract.
type WalletPaymentsSourceDisabled struct {
	SourceId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceDisabled(bytes32 sourceId)
func WalletPaymentsSourceDisabledErrorID() common.Hash {
	return common.HexToHash("0x8b41b325938ef51b536cc31bf6bac9ea957ed9a3ae410800ba38584cf899655c")
}

// UnpackSourceDisabledError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceDisabled(bytes32 sourceId)
func (walletPayments *WalletPayments) UnpackSourceDisabledError(raw []byte) (*WalletPaymentsSourceDisabled, error) {
	out := new(WalletPaymentsSourceDisabled)
	if err := walletPayments.abi.UnpackIntoInterface(out, "SourceDisabled", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsSourceIdZero represents a SourceIdZero error raised by the WalletPayments contract.
type WalletPaymentsSourceIdZero struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceIdZero(uint256 index)
func WalletPaymentsSourceIdZeroErrorID() common.Hash {
	return common.HexToHash("0xd9950937ebabdb658eaf457d54197e233dd4841a64bb53f18118a853631876e7")
}

// UnpackSourceIdZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceIdZero(uint256 index)
func (walletPayments *WalletPayments) UnpackSourceIdZeroError(raw []byte) (*WalletPaymentsSourceIdZero, error) {
	out := new(WalletPaymentsSourceIdZero)
	if err := walletPayments.abi.UnpackIntoInterface(out, "SourceIdZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsSourceKeyTypeNotSupported represents a SourceKeyTypeNotSupported error raised by the WalletPayments contract.
type WalletPaymentsSourceKeyTypeNotSupported struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceKeyTypeNotSupported(uint256 index)
func WalletPaymentsSourceKeyTypeNotSupportedErrorID() common.Hash {
	return common.HexToHash("0x5725fbe89db4399180f35363d1903d3e5b6e2772a320911cbe0f41169693b4d3")
}

// UnpackSourceKeyTypeNotSupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceKeyTypeNotSupported(uint256 index)
func (walletPayments *WalletPayments) UnpackSourceKeyTypeNotSupportedError(raw []byte) (*WalletPaymentsSourceKeyTypeNotSupported, error) {
	out := new(WalletPaymentsSourceKeyTypeNotSupported)
	if err := walletPayments.abi.UnpackIntoInterface(out, "SourceKeyTypeNotSupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsSourceKeyTypeZero represents a SourceKeyTypeZero error raised by the WalletPayments contract.
type WalletPaymentsSourceKeyTypeZero struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceKeyTypeZero(uint256 index)
func WalletPaymentsSourceKeyTypeZeroErrorID() common.Hash {
	return common.HexToHash("0xeeb24635fe3f0a4f44f49b89472e645d49ec073ab20e68b83785acdbc3484f63")
}

// UnpackSourceKeyTypeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceKeyTypeZero(uint256 index)
func (walletPayments *WalletPayments) UnpackSourceKeyTypeZeroError(raw []byte) (*WalletPaymentsSourceKeyTypeZero, error) {
	out := new(WalletPaymentsSourceKeyTypeZero)
	if err := walletPayments.abi.UnpackIntoInterface(out, "SourceKeyTypeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsSourceLimitsNotConfigured represents a SourceLimitsNotConfigured error raised by the WalletPayments contract.
type WalletPaymentsSourceLimitsNotConfigured struct {
	SourceId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceLimitsNotConfigured(bytes32 sourceId)
func WalletPaymentsSourceLimitsNotConfiguredErrorID() common.Hash {
	return common.HexToHash("0x09198876d604df9387eeb94a3b2a10ecb3233765b8e0c28b96dd6be17e7bbe75")
}

// UnpackSourceLimitsNotConfiguredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceLimitsNotConfigured(bytes32 sourceId)
func (walletPayments *WalletPayments) UnpackSourceLimitsNotConfiguredError(raw []byte) (*WalletPaymentsSourceLimitsNotConfigured, error) {
	out := new(WalletPaymentsSourceLimitsNotConfigured)
	if err := walletPayments.abi.UnpackIntoInterface(out, "SourceLimitsNotConfigured", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsSourceModelChainMismatch represents a SourceModelChainMismatch error raised by the WalletPayments contract.
type WalletPaymentsSourceModelChainMismatch struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceModelChainMismatch(uint256 index)
func WalletPaymentsSourceModelChainMismatchErrorID() common.Hash {
	return common.HexToHash("0xe53bd840d3323885de702bbd0a7a735ddfee5a7c9ea8ca5f2775c8c4f78429da")
}

// UnpackSourceModelChainMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceModelChainMismatch(uint256 index)
func (walletPayments *WalletPayments) UnpackSourceModelChainMismatchError(raw []byte) (*WalletPaymentsSourceModelChainMismatch, error) {
	out := new(WalletPaymentsSourceModelChainMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "SourceModelChainMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsSourceNetworkKeyTypeMismatch represents a SourceNetworkKeyTypeMismatch error raised by the WalletPayments contract.
type WalletPaymentsSourceNetworkKeyTypeMismatch struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceNetworkKeyTypeMismatch(uint256 index)
func WalletPaymentsSourceNetworkKeyTypeMismatchErrorID() common.Hash {
	return common.HexToHash("0x439d614dfcb49f503fef9e4cfae0f75f4eec00286f481c57a08dd9de3824b7db")
}

// UnpackSourceNetworkKeyTypeMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceNetworkKeyTypeMismatch(uint256 index)
func (walletPayments *WalletPayments) UnpackSourceNetworkKeyTypeMismatchError(raw []byte) (*WalletPaymentsSourceNetworkKeyTypeMismatch, error) {
	out := new(WalletPaymentsSourceNetworkKeyTypeMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "SourceNetworkKeyTypeMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsSourceNotRegistered represents a SourceNotRegistered error raised by the WalletPayments contract.
type WalletPaymentsSourceNotRegistered struct {
	SourceId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceNotRegistered(bytes32 sourceId)
func WalletPaymentsSourceNotRegisteredErrorID() common.Hash {
	return common.HexToHash("0x321b06cf0196e728deff60d765e36d0ef5386b39e55e899738403838fb35f40a")
}

// UnpackSourceNotRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceNotRegistered(bytes32 sourceId)
func (walletPayments *WalletPayments) UnpackSourceNotRegisteredError(raw []byte) (*WalletPaymentsSourceNotRegistered, error) {
	out := new(WalletPaymentsSourceNotRegistered)
	if err := walletPayments.abi.UnpackIntoInterface(out, "SourceNotRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsSourceOpTypeZero represents a SourceOpTypeZero error raised by the WalletPayments contract.
type WalletPaymentsSourceOpTypeZero struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourceOpTypeZero(uint256 index)
func WalletPaymentsSourceOpTypeZeroErrorID() common.Hash {
	return common.HexToHash("0xef1189f6c1f88aa0a6fd81ffc2c6c160ee2269443b0c873bcf7ce43b2cfef8a7")
}

// UnpackSourceOpTypeZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourceOpTypeZero(uint256 index)
func (walletPayments *WalletPayments) UnpackSourceOpTypeZeroError(raw []byte) (*WalletPaymentsSourceOpTypeZero, error) {
	out := new(WalletPaymentsSourceOpTypeZero)
	if err := walletPayments.abi.UnpackIntoInterface(out, "SourceOpTypeZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsSourcePaymentModelUnknown represents a SourcePaymentModelUnknown error raised by the WalletPayments contract.
type WalletPaymentsSourcePaymentModelUnknown struct {
	Index *big.Int
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error SourcePaymentModelUnknown(uint256 index)
func WalletPaymentsSourcePaymentModelUnknownErrorID() common.Hash {
	return common.HexToHash("0x287d290df524bb6943a708cad1441bc00da738b9b88d19b86e8c763e72bb6706")
}

// UnpackSourcePaymentModelUnknownError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error SourcePaymentModelUnknown(uint256 index)
func (walletPayments *WalletPayments) UnpackSourcePaymentModelUnknownError(raw []byte) (*WalletPaymentsSourcePaymentModelUnknown, error) {
	out := new(WalletPaymentsSourcePaymentModelUnknown)
	if err := walletPayments.abi.UnpackIntoInterface(out, "SourcePaymentModelUnknown", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsStoredAnchorsChanged represents a StoredAnchorsChanged error raised by the WalletPayments contract.
type WalletPaymentsStoredAnchorsChanged struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StoredAnchorsChanged()
func WalletPaymentsStoredAnchorsChangedErrorID() common.Hash {
	return common.HexToHash("0xa75800371050bb34f61374ed78d303f3993a7da54c934674671aae03851dff77")
}

// UnpackStoredAnchorsChangedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StoredAnchorsChanged()
func (walletPayments *WalletPayments) UnpackStoredAnchorsChangedError(raw []byte) (*WalletPaymentsStoredAnchorsChanged, error) {
	out := new(WalletPaymentsStoredAnchorsChanged)
	if err := walletPayments.abi.UnpackIntoInterface(out, "StoredAnchorsChanged", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsTimelockCallNotFound represents a TimelockCallNotFound error raised by the WalletPayments contract.
type WalletPaymentsTimelockCallNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockCallNotFound()
func WalletPaymentsTimelockCallNotFoundErrorID() common.Hash {
	return common.HexToHash("0x262ac987dd09120380d393a0c52cdef1ca58949d1107dc26d8da9f7b527eb33d")
}

// UnpackTimelockCallNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockCallNotFound()
func (walletPayments *WalletPayments) UnpackTimelockCallNotFoundError(raw []byte) (*WalletPaymentsTimelockCallNotFound, error) {
	out := new(WalletPaymentsTimelockCallNotFound)
	if err := walletPayments.abi.UnpackIntoInterface(out, "TimelockCallNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsTimelockNotAllowedYet represents a TimelockNotAllowedYet error raised by the WalletPayments contract.
type WalletPaymentsTimelockNotAllowedYet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockNotAllowedYet()
func WalletPaymentsTimelockNotAllowedYetErrorID() common.Hash {
	return common.HexToHash("0x6124e5c2c545bec1d02048c68500eebdabeae0cc69995a629ebccf8a62f69083")
}

// UnpackTimelockNotAllowedYetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockNotAllowedYet()
func (walletPayments *WalletPayments) UnpackTimelockNotAllowedYetError(raw []byte) (*WalletPaymentsTimelockNotAllowedYet, error) {
	out := new(WalletPaymentsTimelockNotAllowedYet)
	if err := walletPayments.abi.UnpackIntoInterface(out, "TimelockNotAllowedYet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsTimelockValueNotAllowed represents a TimelockValueNotAllowed error raised by the WalletPayments contract.
type WalletPaymentsTimelockValueNotAllowed struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockValueNotAllowed()
func WalletPaymentsTimelockValueNotAllowedErrorID() common.Hash {
	return common.HexToHash("0xe8de44898b890bc2e885e91c46f33b0f00b758d608db20a3d83afd946a5cd46a")
}

// UnpackTimelockValueNotAllowedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockValueNotAllowed()
func (walletPayments *WalletPayments) UnpackTimelockValueNotAllowedError(raw []byte) (*WalletPaymentsTimelockValueNotAllowed, error) {
	out := new(WalletPaymentsTimelockValueNotAllowed)
	if err := walletPayments.abi.UnpackIntoInterface(out, "TimelockValueNotAllowed", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsTokenIdUnsupported represents a TokenIdUnsupported error raised by the WalletPayments contract.
type WalletPaymentsTokenIdUnsupported struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TokenIdUnsupported()
func WalletPaymentsTokenIdUnsupportedErrorID() common.Hash {
	return common.HexToHash("0xd06a2d44d16c1a63f786a8117479c4e15e923472d03585fd154a75e6b888330b")
}

// UnpackTokenIdUnsupportedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TokenIdUnsupported()
func (walletPayments *WalletPayments) UnpackTokenIdUnsupportedError(raw []byte) (*WalletPaymentsTokenIdUnsupported, error) {
	out := new(WalletPaymentsTokenIdUnsupported)
	if err := walletPayments.abi.UnpackIntoInterface(out, "TokenIdUnsupported", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsTooManySchedules represents a TooManySchedules error raised by the WalletPayments contract.
type WalletPaymentsTooManySchedules struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TooManySchedules()
func WalletPaymentsTooManySchedulesErrorID() common.Hash {
	return common.HexToHash("0x6dd8671e1b420f7157257ab2728fbfd77b6ef516af0a461b4affaaee232305eb")
}

// UnpackTooManySchedulesError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TooManySchedules()
func (walletPayments *WalletPayments) UnpackTooManySchedulesError(raw []byte) (*WalletPaymentsTooManySchedules, error) {
	out := new(WalletPaymentsTooManySchedules)
	if err := walletPayments.abi.UnpackIntoInterface(out, "TooManySchedules", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsUnknownWalletRegistry represents a UnknownWalletRegistry error raised by the WalletPayments contract.
type WalletPaymentsUnknownWalletRegistry struct {
	Registry common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnknownWalletRegistry(address registry)
func WalletPaymentsUnknownWalletRegistryErrorID() common.Hash {
	return common.HexToHash("0x9462ef4acaec5e18a8ba594782a8163398ed1354ab655014691c467faaf01f0b")
}

// UnpackUnknownWalletRegistryError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnknownWalletRegistry(address registry)
func (walletPayments *WalletPayments) UnpackUnknownWalletRegistryError(raw []byte) (*WalletPaymentsUnknownWalletRegistry, error) {
	out := new(WalletPaymentsUnknownWalletRegistry)
	if err := walletPayments.abi.UnpackIntoInterface(out, "UnknownWalletRegistry", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsUnsupportedPaymentModel represents a UnsupportedPaymentModel error raised by the WalletPayments contract.
type WalletPaymentsUnsupportedPaymentModel struct {
	SourceId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedPaymentModel(bytes32 sourceId)
func WalletPaymentsUnsupportedPaymentModelErrorID() common.Hash {
	return common.HexToHash("0xa9302a2698d29b4767346820df0365cf999aaf8d17126f81c2690e6e098afc3a")
}

// UnpackUnsupportedPaymentModelError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedPaymentModel(bytes32 sourceId)
func (walletPayments *WalletPayments) UnpackUnsupportedPaymentModelError(raw []byte) (*WalletPaymentsUnsupportedPaymentModel, error) {
	out := new(WalletPaymentsUnsupportedPaymentModel)
	if err := walletPayments.abi.UnpackIntoInterface(out, "UnsupportedPaymentModel", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsUnsupportedSourceId represents a UnsupportedSourceId error raised by the WalletPayments contract.
type WalletPaymentsUnsupportedSourceId struct {
	SourceId [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedSourceId(bytes32 sourceId)
func WalletPaymentsUnsupportedSourceIdErrorID() common.Hash {
	return common.HexToHash("0xe9a2a60c442e03fe02a63f743afaa6a578a16c5c1cfe959ea4878b57953601ff")
}

// UnpackUnsupportedSourceIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedSourceId(bytes32 sourceId)
func (walletPayments *WalletPayments) UnpackUnsupportedSourceIdError(raw []byte) (*WalletPaymentsUnsupportedSourceId, error) {
	out := new(WalletPaymentsUnsupportedSourceId)
	if err := walletPayments.abi.UnpackIntoInterface(out, "UnsupportedSourceId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsUsePaymentReissue represents a UsePaymentReissue error raised by the WalletPayments contract.
type WalletPaymentsUsePaymentReissue struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UsePaymentReissue()
func WalletPaymentsUsePaymentReissueErrorID() common.Hash {
	return common.HexToHash("0x5ea80e63d37bcacf453e055a682833a0af98342052cfaef71c8c95caf7938a8b")
}

// UnpackUsePaymentReissueError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UsePaymentReissue()
func (walletPayments *WalletPayments) UnpackUsePaymentReissueError(raw []byte) (*WalletPaymentsUsePaymentReissue, error) {
	out := new(WalletPaymentsUsePaymentReissue)
	if err := walletPayments.abi.UnpackIntoInterface(out, "UsePaymentReissue", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsValueNotExpected represents a ValueNotExpected error raised by the WalletPayments contract.
type WalletPaymentsValueNotExpected struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ValueNotExpected()
func WalletPaymentsValueNotExpectedErrorID() common.Hash {
	return common.HexToHash("0x6bfec0fe0afdd62e04f5b01c37d09cb1875440f7b29420d7bc0f633fe137e868")
}

// UnpackValueNotExpectedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ValueNotExpected()
func (walletPayments *WalletPayments) UnpackValueNotExpectedError(raw []byte) (*WalletPaymentsValueNotExpected, error) {
	out := new(WalletPaymentsValueNotExpected)
	if err := walletPayments.abi.UnpackIntoInterface(out, "ValueNotExpected", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsWalletNetworkMismatch represents a WalletNetworkMismatch error raised by the WalletPayments contract.
type WalletPaymentsWalletNetworkMismatch struct {
	ExpectedNetwork uint8
	ActualNetwork   uint8
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WalletNetworkMismatch(uint8 expectedNetwork, uint8 actualNetwork)
func WalletPaymentsWalletNetworkMismatchErrorID() common.Hash {
	return common.HexToHash("0x45d5615160a9c5d11adba27e36191c9ec780afac7a7c1fd15e20627c177b063e")
}

// UnpackWalletNetworkMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WalletNetworkMismatch(uint8 expectedNetwork, uint8 actualNetwork)
func (walletPayments *WalletPayments) UnpackWalletNetworkMismatchError(raw []byte) (*WalletPaymentsWalletNetworkMismatch, error) {
	out := new(WalletPaymentsWalletNetworkMismatch)
	if err := walletPayments.abi.UnpackIntoInterface(out, "WalletNetworkMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsWalletNotInProduction represents a WalletNotInProduction error raised by the WalletPayments contract.
type WalletPaymentsWalletNotInProduction struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WalletNotInProduction()
func WalletPaymentsWalletNotInProductionErrorID() common.Hash {
	return common.HexToHash("0x2b7f97d6ba9d96a13708bf15f7c7ed50584fa90c4e3f979e82b455f65581f045")
}

// UnpackWalletNotInProductionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WalletNotInProduction()
func (walletPayments *WalletPayments) UnpackWalletNotInProductionError(raw []byte) (*WalletPaymentsWalletNotInProduction, error) {
	out := new(WalletPaymentsWalletNotInProduction)
	if err := walletPayments.abi.UnpackIntoInterface(out, "WalletNotInProduction", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsWrongKeyType represents a WrongKeyType error raised by the WalletPayments contract.
type WalletPaymentsWrongKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WrongKeyType()
func WalletPaymentsWrongKeyTypeErrorID() common.Hash {
	return common.HexToHash("0xfb224bce70b4062685f138c7eed15a72b903495475f7cd7843835692a7c58bbd")
}

// UnpackWrongKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WrongKeyType()
func (walletPayments *WalletPayments) UnpackWrongKeyTypeError(raw []byte) (*WalletPaymentsWrongKeyType, error) {
	out := new(WalletPaymentsWrongKeyType)
	if err := walletPayments.abi.UnpackIntoInterface(out, "WrongKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// WalletPaymentsXrpEscrowNotFound represents a XrpEscrowNotFound error raised by the WalletPayments contract.
type WalletPaymentsXrpEscrowNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error XrpEscrowNotFound()
func WalletPaymentsXrpEscrowNotFoundErrorID() common.Hash {
	return common.HexToHash("0x43c40ac4df78d5b68e830656175bb32ff4172c8aff1e8a6994766ccbe0520c2d")
}

// UnpackXrpEscrowNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error XrpEscrowNotFound()
func (walletPayments *WalletPayments) UnpackXrpEscrowNotFoundError(raw []byte) (*WalletPaymentsXrpEscrowNotFound, error) {
	out := new(WalletPaymentsXrpEscrowNotFound)
	if err := walletPayments.abi.UnpackIntoInterface(out, "XrpEscrowNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}
