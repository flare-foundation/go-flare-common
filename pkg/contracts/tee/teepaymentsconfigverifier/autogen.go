// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepaymentsconfigverifier

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

// IPMWMultisigAccountConfiguredProof is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IPMWMultisigAccountConfiguredRequestBody
	ResponseBody IPMWMultisigAccountConfiguredResponseBody
}

// IPMWMultisigAccountConfiguredRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredRequestBody struct {
	AccountAddress string
	PublicKeys     [][]byte
	Threshold      uint64
}

// IPMWMultisigAccountConfiguredResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigAccountConfiguredResponseBody struct {
	Status   uint8
	Sequence uint64
}

// IPMWMultisigUtxoConfiguredAnchor is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigUtxoConfiguredAnchor struct {
	GenesisAnchorTxid [32]byte
	GenesisAnchorVout uint32
}

// IPMWMultisigUtxoConfiguredProof is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigUtxoConfiguredProof struct {
	Signatures   IFdc2VerificationFdc2Signatures
	Header       IFdc2HubFdc2ResponseHeader
	RequestBody  IPMWMultisigUtxoConfiguredRequestBody
	ResponseBody IPMWMultisigUtxoConfiguredResponseBody
}

// IPMWMultisigUtxoConfiguredRequestBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigUtxoConfiguredRequestBody struct {
	AccountIndex uint32
	PublicKeys   [][]byte
	Threshold    uint64
	Anchors      []IPMWMultisigUtxoConfiguredAnchor
}

// IPMWMultisigUtxoConfiguredResponseBody is an auto generated low-level Go binding around an user-defined struct.
type IPMWMultisigUtxoConfiguredResponseBody struct {
	Status          uint8
	AccountAddress  string
	AnchorAddresses []string
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// TeePaymentsConfigVerifierMetaData contains all meta data concerning the TeePaymentsConfigVerifier contract.
var TeePaymentsConfigVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountAddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorSetEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestBody\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockInvalidSelector\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2Hub\",\"outputs\":[{\"internalType\":\"contractIFdc2Hub\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2Verification\",\"outputs\":[{\"internalType\":\"contractIFdc2Verification\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"contractIIFlareTeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_accountAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proofOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestAccountConfiguredAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_accountIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Anchor[]\",\"name\":\"_anchors\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proofOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestUtxoConfiguredAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsRegistry\",\"outputs\":[{\"internalType\":\"contractITeePaymentsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigAccountConfigured.PMWMultisigAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAccountConfiguredProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Anchor[]\",\"name\":\"anchors\",\"type\":\"tuple[]\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigUtxoConfigured.PMWMultisigUtxoStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"anchorAddresses\",\"type\":\"string[]\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyUtxoConfiguredProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TeePaymentsConfigVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use TeePaymentsConfigVerifierMetaData.ABI instead.
var TeePaymentsConfigVerifierABI = TeePaymentsConfigVerifierMetaData.ABI

// TeePaymentsConfigVerifier is an auto generated Go binding around an Ethereum contract.
type TeePaymentsConfigVerifier struct {
	TeePaymentsConfigVerifierCaller     // Read-only binding to the contract
	TeePaymentsConfigVerifierTransactor // Write-only binding to the contract
	TeePaymentsConfigVerifierFilterer   // Log filterer for contract events
}

// TeePaymentsConfigVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeePaymentsConfigVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsConfigVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeePaymentsConfigVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsConfigVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeePaymentsConfigVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeePaymentsConfigVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeePaymentsConfigVerifierSession struct {
	Contract     *TeePaymentsConfigVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TeePaymentsConfigVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeePaymentsConfigVerifierCallerSession struct {
	Contract *TeePaymentsConfigVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// TeePaymentsConfigVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeePaymentsConfigVerifierTransactorSession struct {
	Contract     *TeePaymentsConfigVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// TeePaymentsConfigVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeePaymentsConfigVerifierRaw struct {
	Contract *TeePaymentsConfigVerifier // Generic contract binding to access the raw methods on
}

// TeePaymentsConfigVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeePaymentsConfigVerifierCallerRaw struct {
	Contract *TeePaymentsConfigVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// TeePaymentsConfigVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeePaymentsConfigVerifierTransactorRaw struct {
	Contract *TeePaymentsConfigVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeePaymentsConfigVerifier creates a new instance of TeePaymentsConfigVerifier, bound to a specific deployed contract.
func NewTeePaymentsConfigVerifier(address common.Address, backend bind.ContractBackend) (*TeePaymentsConfigVerifier, error) {
	contract, err := bindTeePaymentsConfigVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsConfigVerifier{TeePaymentsConfigVerifierCaller: TeePaymentsConfigVerifierCaller{contract: contract}, TeePaymentsConfigVerifierTransactor: TeePaymentsConfigVerifierTransactor{contract: contract}, TeePaymentsConfigVerifierFilterer: TeePaymentsConfigVerifierFilterer{contract: contract}}, nil
}

// NewTeePaymentsConfigVerifierCaller creates a new read-only instance of TeePaymentsConfigVerifier, bound to a specific deployed contract.
func NewTeePaymentsConfigVerifierCaller(address common.Address, caller bind.ContractCaller) (*TeePaymentsConfigVerifierCaller, error) {
	contract, err := bindTeePaymentsConfigVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsConfigVerifierCaller{contract: contract}, nil
}

// NewTeePaymentsConfigVerifierTransactor creates a new write-only instance of TeePaymentsConfigVerifier, bound to a specific deployed contract.
func NewTeePaymentsConfigVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*TeePaymentsConfigVerifierTransactor, error) {
	contract, err := bindTeePaymentsConfigVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsConfigVerifierTransactor{contract: contract}, nil
}

// NewTeePaymentsConfigVerifierFilterer creates a new log filterer instance of TeePaymentsConfigVerifier, bound to a specific deployed contract.
func NewTeePaymentsConfigVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*TeePaymentsConfigVerifierFilterer, error) {
	contract, err := bindTeePaymentsConfigVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsConfigVerifierFilterer{contract: contract}, nil
}

// bindTeePaymentsConfigVerifier binds a generic wrapper to an already deployed contract.
func bindTeePaymentsConfigVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TeePaymentsConfigVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePaymentsConfigVerifier.Contract.TeePaymentsConfigVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.TeePaymentsConfigVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.TeePaymentsConfigVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeePaymentsConfigVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePaymentsConfigVerifier.Contract.UPGRADEINTERFACEVERSION(&_TeePaymentsConfigVerifier.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _TeePaymentsConfigVerifier.Contract.UPGRADEINTERFACEVERSION(&_TeePaymentsConfigVerifier.CallOpts)
}

// Fdc2Hub is a free data retrieval call binding the contract method 0xa7566ff3.
//
// Solidity: function fdc2Hub() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) Fdc2Hub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "fdc2Hub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fdc2Hub is a free data retrieval call binding the contract method 0xa7566ff3.
//
// Solidity: function fdc2Hub() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) Fdc2Hub() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.Fdc2Hub(&_TeePaymentsConfigVerifier.CallOpts)
}

// Fdc2Hub is a free data retrieval call binding the contract method 0xa7566ff3.
//
// Solidity: function fdc2Hub() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) Fdc2Hub() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.Fdc2Hub(&_TeePaymentsConfigVerifier.CallOpts)
}

// Fdc2Verification is a free data retrieval call binding the contract method 0xbf2e9839.
//
// Solidity: function fdc2Verification() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) Fdc2Verification(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "fdc2Verification")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Fdc2Verification is a free data retrieval call binding the contract method 0xbf2e9839.
//
// Solidity: function fdc2Verification() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) Fdc2Verification() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.Fdc2Verification(&_TeePaymentsConfigVerifier.CallOpts)
}

// Fdc2Verification is a free data retrieval call binding the contract method 0xbf2e9839.
//
// Solidity: function fdc2Verification() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) Fdc2Verification() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.Fdc2Verification(&_TeePaymentsConfigVerifier.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) FlareSystemsManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "flareSystemsManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) FlareSystemsManager() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.FlareSystemsManager(&_TeePaymentsConfigVerifier.CallOpts)
}

// FlareSystemsManager is a free data retrieval call binding the contract method 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) FlareSystemsManager() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.FlareSystemsManager(&_TeePaymentsConfigVerifier.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) FlareTeeManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "flareTeeManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) FlareTeeManager() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.FlareTeeManager(&_TeePaymentsConfigVerifier.CallOpts)
}

// FlareTeeManager is a free data retrieval call binding the contract method 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) FlareTeeManager() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.FlareTeeManager(&_TeePaymentsConfigVerifier.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) GetAddressUpdater(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "getAddressUpdater")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) GetAddressUpdater() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.GetAddressUpdater(&_TeePaymentsConfigVerifier.CallOpts)
}

// GetAddressUpdater is a free data retrieval call binding the contract method 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) GetAddressUpdater() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.GetAddressUpdater(&_TeePaymentsConfigVerifier.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "governance")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) Governance() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.Governance(&_TeePaymentsConfigVerifier.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) Governance() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.Governance(&_TeePaymentsConfigVerifier.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) GovernanceSettings(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "governanceSettings")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) GovernanceSettings() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.GovernanceSettings(&_TeePaymentsConfigVerifier.CallOpts)
}

// GovernanceSettings is a free data retrieval call binding the contract method 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) GovernanceSettings() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.GovernanceSettings(&_TeePaymentsConfigVerifier.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) Implementation() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.Implementation(&_TeePaymentsConfigVerifier.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) Implementation() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.Implementation(&_TeePaymentsConfigVerifier.CallOpts)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) IsExecutor(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "isExecutor", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePaymentsConfigVerifier.Contract.IsExecutor(&_TeePaymentsConfigVerifier.CallOpts, _address)
}

// IsExecutor is a free data retrieval call binding the contract method 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) IsExecutor(_address common.Address) (bool, error) {
	return _TeePaymentsConfigVerifier.Contract.IsExecutor(&_TeePaymentsConfigVerifier.CallOpts, _address)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) ProductionMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "productionMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) ProductionMode() (bool, error) {
	return _TeePaymentsConfigVerifier.Contract.ProductionMode(&_TeePaymentsConfigVerifier.CallOpts)
}

// ProductionMode is a free data retrieval call binding the contract method 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) ProductionMode() (bool, error) {
	return _TeePaymentsConfigVerifier.Contract.ProductionMode(&_TeePaymentsConfigVerifier.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) ProxiableUUID() ([32]byte, error) {
	return _TeePaymentsConfigVerifier.Contract.ProxiableUUID(&_TeePaymentsConfigVerifier.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TeePaymentsConfigVerifier.Contract.ProxiableUUID(&_TeePaymentsConfigVerifier.CallOpts)
}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCaller) TeePaymentsRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeePaymentsConfigVerifier.contract.Call(opts, &out, "teePaymentsRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) TeePaymentsRegistry() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.TeePaymentsRegistry(&_TeePaymentsConfigVerifier.CallOpts)
}

// TeePaymentsRegistry is a free data retrieval call binding the contract method 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierCallerSession) TeePaymentsRegistry() (common.Address, error) {
	return _TeePaymentsConfigVerifier.Contract.TeePaymentsRegistry(&_TeePaymentsConfigVerifier.CallOpts)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactor) CancelGovernanceCall(opts *bind.TransactOpts, _encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.contract.Transact(opts, "cancelGovernanceCall", _encodedCall)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) CancelGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.CancelGovernanceCall(&_TeePaymentsConfigVerifier.TransactOpts, _encodedCall)
}

// CancelGovernanceCall is a paid mutator transaction binding the contract method 0x16fc2f6d.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactorSession) CancelGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.CancelGovernanceCall(&_TeePaymentsConfigVerifier.TransactOpts, _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactor) ExecuteGovernanceCall(opts *bind.TransactOpts, _encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.contract.Transact(opts, "executeGovernanceCall", _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) ExecuteGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.ExecuteGovernanceCall(&_TeePaymentsConfigVerifier.TransactOpts, _encodedCall)
}

// ExecuteGovernanceCall is a paid mutator transaction binding the contract method 0x20c5f99d.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactorSession) ExecuteGovernanceCall(_encodedCall []byte) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.ExecuteGovernanceCall(&_TeePaymentsConfigVerifier.TransactOpts, _encodedCall)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactor) Initialize(opts *bind.TransactOpts, _governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.contract.Transact(opts, "initialize", _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.Initialize(&_TeePaymentsConfigVerifier.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactorSession) Initialize(_governanceSettings common.Address, _initialGovernance common.Address, _addressUpdater common.Address) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.Initialize(&_TeePaymentsConfigVerifier.TransactOpts, _governanceSettings, _initialGovernance, _addressUpdater)
}

// RequestAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0x17005854.
//
// Solidity: function requestAccountConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactor) RequestAccountConfiguredAttestation(opts *bind.TransactOpts, _walletId [32]byte, _sourceId [32]byte, _accountAddress string, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.contract.Transact(opts, "requestAccountConfiguredAttestation", _walletId, _sourceId, _accountAddress, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0x17005854.
//
// Solidity: function requestAccountConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) RequestAccountConfiguredAttestation(_walletId [32]byte, _sourceId [32]byte, _accountAddress string, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.RequestAccountConfiguredAttestation(&_TeePaymentsConfigVerifier.TransactOpts, _walletId, _sourceId, _accountAddress, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestAccountConfiguredAttestation is a paid mutator transaction binding the contract method 0x17005854.
//
// Solidity: function requestAccountConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactorSession) RequestAccountConfiguredAttestation(_walletId [32]byte, _sourceId [32]byte, _accountAddress string, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.RequestAccountConfiguredAttestation(&_TeePaymentsConfigVerifier.TransactOpts, _walletId, _sourceId, _accountAddress, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestUtxoConfiguredAttestation is a paid mutator transaction binding the contract method 0x0f12700d.
//
// Solidity: function requestUtxoConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, uint32 _accountIndex, (bytes32,uint32)[] _anchors, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactor) RequestUtxoConfiguredAttestation(opts *bind.TransactOpts, _walletId [32]byte, _sourceId [32]byte, _accountIndex uint32, _anchors []IPMWMultisigUtxoConfiguredAnchor, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.contract.Transact(opts, "requestUtxoConfiguredAttestation", _walletId, _sourceId, _accountIndex, _anchors, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestUtxoConfiguredAttestation is a paid mutator transaction binding the contract method 0x0f12700d.
//
// Solidity: function requestUtxoConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, uint32 _accountIndex, (bytes32,uint32)[] _anchors, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) RequestUtxoConfiguredAttestation(_walletId [32]byte, _sourceId [32]byte, _accountIndex uint32, _anchors []IPMWMultisigUtxoConfiguredAnchor, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.RequestUtxoConfiguredAttestation(&_TeePaymentsConfigVerifier.TransactOpts, _walletId, _sourceId, _accountIndex, _anchors, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// RequestUtxoConfiguredAttestation is a paid mutator transaction binding the contract method 0x0f12700d.
//
// Solidity: function requestUtxoConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, uint32 _accountIndex, (bytes32,uint32)[] _anchors, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactorSession) RequestUtxoConfiguredAttestation(_walletId [32]byte, _sourceId [32]byte, _accountIndex uint32, _anchors []IPMWMultisigUtxoConfiguredAnchor, _testOnTeeId common.Address, _proofOwner common.Address, _claimBackAddress common.Address) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.RequestUtxoConfiguredAttestation(&_TeePaymentsConfigVerifier.TransactOpts, _walletId, _sourceId, _accountIndex, _anchors, _testOnTeeId, _proofOwner, _claimBackAddress)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactor) SwitchToProductionMode(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.contract.Transact(opts, "switchToProductionMode")
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.SwitchToProductionMode(&_TeePaymentsConfigVerifier.TransactOpts)
}

// SwitchToProductionMode is a paid mutator transaction binding the contract method 0xf5a98383.
//
// Solidity: function switchToProductionMode() returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactorSession) SwitchToProductionMode() (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.SwitchToProductionMode(&_TeePaymentsConfigVerifier.TransactOpts)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactor) UpdateContractAddresses(opts *bind.TransactOpts, _contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.contract.Transact(opts, "updateContractAddresses", _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.UpdateContractAddresses(&_TeePaymentsConfigVerifier.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpdateContractAddresses is a paid mutator transaction binding the contract method 0xb00c0b76.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactorSession) UpdateContractAddresses(_contractNameHashes [][32]byte, _contractAddresses []common.Address) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.UpdateContractAddresses(&_TeePaymentsConfigVerifier.TransactOpts, _contractNameHashes, _contractAddresses)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactor) UpgradeToAndCall(opts *bind.TransactOpts, _newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.contract.Transact(opts, "upgradeToAndCall", _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.UpgradeToAndCall(&_TeePaymentsConfigVerifier.TransactOpts, _newImplementation, _data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactorSession) UpgradeToAndCall(_newImplementation common.Address, _data []byte) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.UpgradeToAndCall(&_TeePaymentsConfigVerifier.TransactOpts, _newImplementation, _data)
}

// VerifyAccountConfiguredProof is a paid mutator transaction binding the contract method 0x8c876193.
//
// Solidity: function verifyAccountConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactor) VerifyAccountConfiguredProof(opts *bind.TransactOpts, _walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.contract.Transact(opts, "verifyAccountConfiguredProof", _walletId, _proof)
}

// VerifyAccountConfiguredProof is a paid mutator transaction binding the contract method 0x8c876193.
//
// Solidity: function verifyAccountConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) VerifyAccountConfiguredProof(_walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.VerifyAccountConfiguredProof(&_TeePaymentsConfigVerifier.TransactOpts, _walletId, _proof)
}

// VerifyAccountConfiguredProof is a paid mutator transaction binding the contract method 0x8c876193.
//
// Solidity: function verifyAccountConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactorSession) VerifyAccountConfiguredProof(_walletId [32]byte, _proof IPMWMultisigAccountConfiguredProof) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.VerifyAccountConfiguredProof(&_TeePaymentsConfigVerifier.TransactOpts, _walletId, _proof)
}

// VerifyUtxoConfiguredProof is a paid mutator transaction binding the contract method 0x6858376b.
//
// Solidity: function verifyUtxoConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string,string[])) _proof) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactor) VerifyUtxoConfiguredProof(opts *bind.TransactOpts, _walletId [32]byte, _proof IPMWMultisigUtxoConfiguredProof) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.contract.Transact(opts, "verifyUtxoConfiguredProof", _walletId, _proof)
}

// VerifyUtxoConfiguredProof is a paid mutator transaction binding the contract method 0x6858376b.
//
// Solidity: function verifyUtxoConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string,string[])) _proof) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierSession) VerifyUtxoConfiguredProof(_walletId [32]byte, _proof IPMWMultisigUtxoConfiguredProof) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.VerifyUtxoConfiguredProof(&_TeePaymentsConfigVerifier.TransactOpts, _walletId, _proof)
}

// VerifyUtxoConfiguredProof is a paid mutator transaction binding the contract method 0x6858376b.
//
// Solidity: function verifyUtxoConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string,string[])) _proof) returns()
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierTransactorSession) VerifyUtxoConfiguredProof(_walletId [32]byte, _proof IPMWMultisigUtxoConfiguredProof) (*types.Transaction, error) {
	return _TeePaymentsConfigVerifier.Contract.VerifyUtxoConfiguredProof(&_TeePaymentsConfigVerifier.TransactOpts, _walletId, _proof)
}

// TeePaymentsConfigVerifierGovernanceCallTimelockedIterator is returned from FilterGovernanceCallTimelocked and is used to iterate over the raw logs and unpacked data for GovernanceCallTimelocked events raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierGovernanceCallTimelockedIterator struct {
	Event *TeePaymentsConfigVerifierGovernanceCallTimelocked // Event containing the contract specifics and raw log

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
func (it *TeePaymentsConfigVerifierGovernanceCallTimelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsConfigVerifierGovernanceCallTimelocked)
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
		it.Event = new(TeePaymentsConfigVerifierGovernanceCallTimelocked)
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
func (it *TeePaymentsConfigVerifierGovernanceCallTimelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsConfigVerifierGovernanceCallTimelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsConfigVerifierGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterGovernanceCallTimelocked is a free log retrieval operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) FilterGovernanceCallTimelocked(opts *bind.FilterOpts) (*TeePaymentsConfigVerifierGovernanceCallTimelockedIterator, error) {

	logs, sub, err := _TeePaymentsConfigVerifier.contract.FilterLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsConfigVerifierGovernanceCallTimelockedIterator{contract: _TeePaymentsConfigVerifier.contract, event: "GovernanceCallTimelocked", logs: logs, sub: sub}, nil
}

// WatchGovernanceCallTimelocked is a free log subscription operation binding the contract event 0x8c02104dfc280f713854f25297de671710c544c58de69dbde8fb66974ce1ab9e.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) WatchGovernanceCallTimelocked(opts *bind.WatchOpts, sink chan<- *TeePaymentsConfigVerifierGovernanceCallTimelocked) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsConfigVerifier.contract.WatchLogs(opts, "GovernanceCallTimelocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsConfigVerifierGovernanceCallTimelocked)
				if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
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
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) ParseGovernanceCallTimelocked(log types.Log) (*TeePaymentsConfigVerifierGovernanceCallTimelocked, error) {
	event := new(TeePaymentsConfigVerifierGovernanceCallTimelocked)
	if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "GovernanceCallTimelocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsConfigVerifierGovernanceInitialisedIterator is returned from FilterGovernanceInitialised and is used to iterate over the raw logs and unpacked data for GovernanceInitialised events raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierGovernanceInitialisedIterator struct {
	Event *TeePaymentsConfigVerifierGovernanceInitialised // Event containing the contract specifics and raw log

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
func (it *TeePaymentsConfigVerifierGovernanceInitialisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsConfigVerifierGovernanceInitialised)
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
		it.Event = new(TeePaymentsConfigVerifierGovernanceInitialised)
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
func (it *TeePaymentsConfigVerifierGovernanceInitialisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsConfigVerifierGovernanceInitialisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsConfigVerifierGovernanceInitialised represents a GovernanceInitialised event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGovernanceInitialised is a free log retrieval operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) FilterGovernanceInitialised(opts *bind.FilterOpts) (*TeePaymentsConfigVerifierGovernanceInitialisedIterator, error) {

	logs, sub, err := _TeePaymentsConfigVerifier.contract.FilterLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsConfigVerifierGovernanceInitialisedIterator{contract: _TeePaymentsConfigVerifier.contract, event: "GovernanceInitialised", logs: logs, sub: sub}, nil
}

// WatchGovernanceInitialised is a free log subscription operation binding the contract event 0x9789733827840833afc031fb2ef9ab6894271f77bad2085687cf4ae5c7bee4db.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) WatchGovernanceInitialised(opts *bind.WatchOpts, sink chan<- *TeePaymentsConfigVerifierGovernanceInitialised) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsConfigVerifier.contract.WatchLogs(opts, "GovernanceInitialised")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsConfigVerifierGovernanceInitialised)
				if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
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
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) ParseGovernanceInitialised(log types.Log) (*TeePaymentsConfigVerifierGovernanceInitialised, error) {
	event := new(TeePaymentsConfigVerifierGovernanceInitialised)
	if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "GovernanceInitialised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsConfigVerifierGovernedProductionModeEnteredIterator is returned from FilterGovernedProductionModeEntered and is used to iterate over the raw logs and unpacked data for GovernedProductionModeEntered events raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierGovernedProductionModeEnteredIterator struct {
	Event *TeePaymentsConfigVerifierGovernedProductionModeEntered // Event containing the contract specifics and raw log

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
func (it *TeePaymentsConfigVerifierGovernedProductionModeEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsConfigVerifierGovernedProductionModeEntered)
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
		it.Event = new(TeePaymentsConfigVerifierGovernedProductionModeEntered)
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
func (it *TeePaymentsConfigVerifierGovernedProductionModeEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsConfigVerifierGovernedProductionModeEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsConfigVerifierGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterGovernedProductionModeEntered is a free log retrieval operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) FilterGovernedProductionModeEntered(opts *bind.FilterOpts) (*TeePaymentsConfigVerifierGovernedProductionModeEnteredIterator, error) {

	logs, sub, err := _TeePaymentsConfigVerifier.contract.FilterLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsConfigVerifierGovernedProductionModeEnteredIterator{contract: _TeePaymentsConfigVerifier.contract, event: "GovernedProductionModeEntered", logs: logs, sub: sub}, nil
}

// WatchGovernedProductionModeEntered is a free log subscription operation binding the contract event 0x83af113638b5422f9e977cebc0aaf0eaf2188eb9a8baae7f9d46c42b33a1560c.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) WatchGovernedProductionModeEntered(opts *bind.WatchOpts, sink chan<- *TeePaymentsConfigVerifierGovernedProductionModeEntered) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsConfigVerifier.contract.WatchLogs(opts, "GovernedProductionModeEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsConfigVerifierGovernedProductionModeEntered)
				if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
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
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) ParseGovernedProductionModeEntered(log types.Log) (*TeePaymentsConfigVerifierGovernedProductionModeEntered, error) {
	event := new(TeePaymentsConfigVerifierGovernedProductionModeEntered)
	if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "GovernedProductionModeEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsConfigVerifierInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierInitializedIterator struct {
	Event *TeePaymentsConfigVerifierInitialized // Event containing the contract specifics and raw log

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
func (it *TeePaymentsConfigVerifierInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsConfigVerifierInitialized)
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
		it.Event = new(TeePaymentsConfigVerifierInitialized)
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
func (it *TeePaymentsConfigVerifierInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsConfigVerifierInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsConfigVerifierInitialized represents a Initialized event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) FilterInitialized(opts *bind.FilterOpts) (*TeePaymentsConfigVerifierInitializedIterator, error) {

	logs, sub, err := _TeePaymentsConfigVerifier.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsConfigVerifierInitializedIterator{contract: _TeePaymentsConfigVerifier.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TeePaymentsConfigVerifierInitialized) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsConfigVerifier.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsConfigVerifierInitialized)
				if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) ParseInitialized(log types.Log) (*TeePaymentsConfigVerifierInitialized, error) {
	event := new(TeePaymentsConfigVerifierInitialized)
	if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsConfigVerifierTimelockedGovernanceCallCanceledIterator is returned from FilterTimelockedGovernanceCallCanceled and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallCanceled events raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierTimelockedGovernanceCallCanceledIterator struct {
	Event *TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled // Event containing the contract specifics and raw log

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
func (it *TeePaymentsConfigVerifierTimelockedGovernanceCallCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled)
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
		it.Event = new(TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled)
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
func (it *TeePaymentsConfigVerifierTimelockedGovernanceCallCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsConfigVerifierTimelockedGovernanceCallCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallCanceled is a free log retrieval operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) FilterTimelockedGovernanceCallCanceled(opts *bind.FilterOpts) (*TeePaymentsConfigVerifierTimelockedGovernanceCallCanceledIterator, error) {

	logs, sub, err := _TeePaymentsConfigVerifier.contract.FilterLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsConfigVerifierTimelockedGovernanceCallCanceledIterator{contract: _TeePaymentsConfigVerifier.contract, event: "TimelockedGovernanceCallCanceled", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallCanceled is a free log subscription operation binding the contract event 0x69b058d6225c01c1f2a25801ca5b05705fa2e9118e93d518390ba804398c87b1.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) WatchTimelockedGovernanceCallCanceled(opts *bind.WatchOpts, sink chan<- *TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsConfigVerifier.contract.WatchLogs(opts, "TimelockedGovernanceCallCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled)
				if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
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
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) ParseTimelockedGovernanceCallCanceled(log types.Log) (*TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled, error) {
	event := new(TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled)
	if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "TimelockedGovernanceCallCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsConfigVerifierTimelockedGovernanceCallExecutedIterator is returned from FilterTimelockedGovernanceCallExecuted and is used to iterate over the raw logs and unpacked data for TimelockedGovernanceCallExecuted events raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierTimelockedGovernanceCallExecutedIterator struct {
	Event *TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted // Event containing the contract specifics and raw log

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
func (it *TeePaymentsConfigVerifierTimelockedGovernanceCallExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted)
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
		it.Event = new(TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted)
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
func (it *TeePaymentsConfigVerifierTimelockedGovernanceCallExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsConfigVerifierTimelockedGovernanceCallExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimelockedGovernanceCallExecuted is a free log retrieval operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) FilterTimelockedGovernanceCallExecuted(opts *bind.FilterOpts) (*TeePaymentsConfigVerifierTimelockedGovernanceCallExecutedIterator, error) {

	logs, sub, err := _TeePaymentsConfigVerifier.contract.FilterLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return &TeePaymentsConfigVerifierTimelockedGovernanceCallExecutedIterator{contract: _TeePaymentsConfigVerifier.contract, event: "TimelockedGovernanceCallExecuted", logs: logs, sub: sub}, nil
}

// WatchTimelockedGovernanceCallExecuted is a free log subscription operation binding the contract event 0xec1225e5a8a8acb91e03ce648c683c74f5d152a775b9715980999441d714c44f.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) WatchTimelockedGovernanceCallExecuted(opts *bind.WatchOpts, sink chan<- *TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted) (event.Subscription, error) {

	logs, sub, err := _TeePaymentsConfigVerifier.contract.WatchLogs(opts, "TimelockedGovernanceCallExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted)
				if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
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
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) ParseTimelockedGovernanceCallExecuted(log types.Log) (*TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted, error) {
	event := new(TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted)
	if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "TimelockedGovernanceCallExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeePaymentsConfigVerifierUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierUpgradedIterator struct {
	Event *TeePaymentsConfigVerifierUpgraded // Event containing the contract specifics and raw log

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
func (it *TeePaymentsConfigVerifierUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeePaymentsConfigVerifierUpgraded)
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
		it.Event = new(TeePaymentsConfigVerifierUpgraded)
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
func (it *TeePaymentsConfigVerifierUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeePaymentsConfigVerifierUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeePaymentsConfigVerifierUpgraded represents a Upgraded event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TeePaymentsConfigVerifierUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePaymentsConfigVerifier.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TeePaymentsConfigVerifierUpgradedIterator{contract: _TeePaymentsConfigVerifier.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TeePaymentsConfigVerifierUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TeePaymentsConfigVerifier.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeePaymentsConfigVerifierUpgraded)
				if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TeePaymentsConfigVerifier *TeePaymentsConfigVerifierFilterer) ParseUpgraded(log types.Log) (*TeePaymentsConfigVerifierUpgraded, error) {
	event := new(TeePaymentsConfigVerifierUpgraded)
	if err := _TeePaymentsConfigVerifier.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
