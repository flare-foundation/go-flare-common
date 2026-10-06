// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepayments

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

// ITeePaymentsBasePMWMultisigAccount is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsBasePMWMultisigAccount struct {
	SourceId       [32]byte
	AccountAddress string
}

// ITeePaymentsBasePaymentInstruction is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsBasePaymentInstruction struct {
	RecipientAddress string
	TokenId          []byte
	Amount           *big.Int
	MaxFee           *big.Int
	PaymentReference [32]byte
}

// ITeePaymentsBaseReissueFeeParams is an auto generated low-level Go binding around an user-defined struct.
type ITeePaymentsBaseReissueFeeParams struct {
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

// TeePaymentsMetaData contains all meta data concerning the TeePayments contract.
var TeePaymentsMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuthorizationAddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPaymentInstructionCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LengthsMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAuthorizationAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyWalletOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMWMultisigAccountAddressAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PMWMultisigAccountNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentAmountZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentHashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StartNewRequired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockCallNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WalletNotInProduction\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"walletId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"authorizationAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"initialNonce\",\"type\":\"uint64\"}],\"name\":\"PMWMultisigAccountAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigAccountConfigured.PMWMultisigAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"name\":\"addPMWMultisigAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addressValidator\",\"outputs\":[{\"internalType\":\"contractIAddressValidator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"contractIIFlareTeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getAuthorizationAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_authorizationAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getInitialNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_initialNonce\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getNextPaymentId\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nextPaymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"_opCommand\",\"type\":\"bytes32\"}],\"name\":\"getPaymentFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"name\":\"getPaymentHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_paymentHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"}],\"name\":\"getWalletAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"}],\"name\":\"getWalletId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsBase.PaymentInstruction\",\"name\":\"_paymentInstruction\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"pay\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentModel\",\"outputs\":[{\"internalType\":\"enumPaymentModel\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structITeePaymentsBase.PMWMultisigAccount\",\"name\":\"_account\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"_paymentId\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"recipientAddress\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"tokenId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"paymentReference\",\"type\":\"bytes32\"}],\"internalType\":\"structITeePaymentsBase.PaymentInstruction[]\",\"name\":\"_paymentInstructions\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"maxFeePerPayment\",\"type\":\"uint256[]\"},{\"internalType\":\"int16[][]\",\"name\":\"factorsBIPSPerPayment\",\"type\":\"int16[][]\"},{\"internalType\":\"uint16[]\",\"name\":\"delaysSeconds\",\"type\":\"uint16[]\"}],\"internalType\":\"structITeePaymentsBase.ReissueFeeParams\",\"name\":\"_reissueFeeParams\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_startNew\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"reissue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_finalized\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsConfigVerifier\",\"outputs\":[{\"internalType\":\"contractITeePaymentsConfigVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsFeeScheduleManager\",\"outputs\":[{\"internalType\":\"contractITeePaymentsFeeScheduleManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsRegistry\",\"outputs\":[{\"internalType\":\"contractITeePaymentsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	ID:  "TeePayments",
}

// TeePayments is an auto generated Go binding around an Ethereum contract.
type TeePayments struct {
	abi abi.ABI
}

// NewTeePayments creates a new instance of TeePayments.
func NewTeePayments() *TeePayments {
	parsed, err := TeePaymentsMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TeePayments{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TeePayments) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePayments *TeePayments) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := teePayments.abi.Pack("UPGRADE_INTERFACE_VERSION")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePayments *TeePayments) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return teePayments.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePayments *TeePayments) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := teePayments.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackAddPMWMultisigAccount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x895d1bb4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof, address _authorizationAddress) returns()
func (teePayments *TeePayments) PackAddPMWMultisigAccount(walletId [32]byte, proof IPMWMultisigAccountConfiguredProof, authorizationAddress common.Address) []byte {
	enc, err := teePayments.abi.Pack("addPMWMultisigAccount", walletId, proof, authorizationAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddPMWMultisigAccount is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x895d1bb4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addPMWMultisigAccount(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof, address _authorizationAddress) returns()
func (teePayments *TeePayments) TryPackAddPMWMultisigAccount(walletId [32]byte, proof IPMWMultisigAccountConfiguredProof, authorizationAddress common.Address) ([]byte, error) {
	return teePayments.abi.Pack("addPMWMultisigAccount", walletId, proof, authorizationAddress)
}

// PackAddressValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe37facf4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function addressValidator() view returns(address)
func (teePayments *TeePayments) PackAddressValidator() []byte {
	enc, err := teePayments.abi.Pack("addressValidator")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackAddressValidator is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe37facf4.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function addressValidator() view returns(address)
func (teePayments *TeePayments) TryPackAddressValidator() ([]byte, error) {
	return teePayments.abi.Pack("addressValidator")
}

// UnpackAddressValidator is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe37facf4.
//
// Solidity: function addressValidator() view returns(address)
func (teePayments *TeePayments) UnpackAddressValidator(data []byte) (common.Address, error) {
	out, err := teePayments.abi.Unpack("addressValidator", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16fc2f6d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (teePayments *TeePayments) PackCancelGovernanceCall(encodedCall []byte) []byte {
	enc, err := teePayments.abi.Pack("cancelGovernanceCall", encodedCall)
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
func (teePayments *TeePayments) TryPackCancelGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teePayments.abi.Pack("cancelGovernanceCall", encodedCall)
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (teePayments *TeePayments) PackExecuteGovernanceCall(encodedCall []byte) []byte {
	enc, err := teePayments.abi.Pack("executeGovernanceCall", encodedCall)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (teePayments *TeePayments) TryPackExecuteGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teePayments.abi.Pack("executeGovernanceCall", encodedCall)
}

// PackFlareTeeManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x453f7ab4.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function flareTeeManager() view returns(address)
func (teePayments *TeePayments) PackFlareTeeManager() []byte {
	enc, err := teePayments.abi.Pack("flareTeeManager")
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
func (teePayments *TeePayments) TryPackFlareTeeManager() ([]byte, error) {
	return teePayments.abi.Pack("flareTeeManager")
}

// UnpackFlareTeeManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (teePayments *TeePayments) UnpackFlareTeeManager(data []byte) (common.Address, error) {
	out, err := teePayments.abi.Unpack("flareTeeManager", data)
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
func (teePayments *TeePayments) PackGetAddressUpdater() []byte {
	enc, err := teePayments.abi.Pack("getAddressUpdater")
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
func (teePayments *TeePayments) TryPackGetAddressUpdater() ([]byte, error) {
	return teePayments.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (teePayments *TeePayments) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := teePayments.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetAuthorizationAddress is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x410642e0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (teePayments *TeePayments) PackGetAuthorizationAddress(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePayments.abi.Pack("getAuthorizationAddress", account)
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
func (teePayments *TeePayments) TryPackGetAuthorizationAddress(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePayments.abi.Pack("getAuthorizationAddress", account)
}

// UnpackGetAuthorizationAddress is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x410642e0.
//
// Solidity: function getAuthorizationAddress((bytes32,string) _account) view returns(address _authorizationAddress)
func (teePayments *TeePayments) UnpackGetAuthorizationAddress(data []byte) (common.Address, error) {
	out, err := teePayments.abi.Unpack("getAuthorizationAddress", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGetInitialNonce is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x2bbba299.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getInitialNonce((bytes32,string) _account) view returns(uint64 _initialNonce)
func (teePayments *TeePayments) PackGetInitialNonce(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePayments.abi.Pack("getInitialNonce", account)
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
func (teePayments *TeePayments) TryPackGetInitialNonce(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePayments.abi.Pack("getInitialNonce", account)
}

// UnpackGetInitialNonce is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x2bbba299.
//
// Solidity: function getInitialNonce((bytes32,string) _account) view returns(uint64 _initialNonce)
func (teePayments *TeePayments) UnpackGetInitialNonce(data []byte) (uint64, error) {
	out, err := teePayments.abi.Unpack("getInitialNonce", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetNextPaymentId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfb49ac30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (teePayments *TeePayments) PackGetNextPaymentId(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePayments.abi.Pack("getNextPaymentId", account)
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
func (teePayments *TeePayments) TryPackGetNextPaymentId(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePayments.abi.Pack("getNextPaymentId", account)
}

// UnpackGetNextPaymentId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfb49ac30.
//
// Solidity: function getNextPaymentId((bytes32,string) _account) view returns(uint64 _nextPaymentId)
func (teePayments *TeePayments) UnpackGetNextPaymentId(data []byte) (uint64, error) {
	out, err := teePayments.abi.Unpack("getNextPaymentId", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackGetPaymentFee is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x57abf78b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (teePayments *TeePayments) PackGetPaymentFee(account ITeePaymentsBasePMWMultisigAccount, opCommand [32]byte) []byte {
	enc, err := teePayments.abi.Pack("getPaymentFee", account, opCommand)
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
func (teePayments *TeePayments) TryPackGetPaymentFee(account ITeePaymentsBasePMWMultisigAccount, opCommand [32]byte) ([]byte, error) {
	return teePayments.abi.Pack("getPaymentFee", account, opCommand)
}

// UnpackGetPaymentFee is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x57abf78b.
//
// Solidity: function getPaymentFee((bytes32,string) _account, bytes32 _opCommand) view returns(uint256 _fee)
func (teePayments *TeePayments) UnpackGetPaymentFee(data []byte) (*big.Int, error) {
	out, err := teePayments.abi.Unpack("getPaymentFee", data)
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
func (teePayments *TeePayments) PackGetPaymentHash(account ITeePaymentsBasePMWMultisigAccount, paymentId uint64) []byte {
	enc, err := teePayments.abi.Pack("getPaymentHash", account, paymentId)
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
func (teePayments *TeePayments) TryPackGetPaymentHash(account ITeePaymentsBasePMWMultisigAccount, paymentId uint64) ([]byte, error) {
	return teePayments.abi.Pack("getPaymentHash", account, paymentId)
}

// UnpackGetPaymentHash is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x0079dfe8.
//
// Solidity: function getPaymentHash((bytes32,string) _account, uint64 _paymentId) view returns(bytes32 _paymentHash)
func (teePayments *TeePayments) UnpackGetPaymentHash(data []byte) ([32]byte, error) {
	out, err := teePayments.abi.Unpack("getPaymentHash", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGetWalletAccounts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a54c1b0.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (teePayments *TeePayments) PackGetWalletAccounts(walletId [32]byte) []byte {
	enc, err := teePayments.abi.Pack("getWalletAccounts", walletId)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletAccounts is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x3a54c1b0.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (teePayments *TeePayments) TryPackGetWalletAccounts(walletId [32]byte) ([]byte, error) {
	return teePayments.abi.Pack("getWalletAccounts", walletId)
}

// UnpackGetWalletAccounts is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x3a54c1b0.
//
// Solidity: function getWalletAccounts(bytes32 _walletId) view returns((bytes32,string)[])
func (teePayments *TeePayments) UnpackGetWalletAccounts(data []byte) ([]ITeePaymentsBasePMWMultisigAccount, error) {
	out, err := teePayments.abi.Unpack("getWalletAccounts", data)
	if err != nil {
		return *new([]ITeePaymentsBasePMWMultisigAccount), err
	}
	out0 := *abi.ConvertType(out[0], new([]ITeePaymentsBasePMWMultisigAccount)).(*[]ITeePaymentsBasePMWMultisigAccount)
	return out0, nil
}

// PackGetWalletId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5623b3f5.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (teePayments *TeePayments) PackGetWalletId(account ITeePaymentsBasePMWMultisigAccount) []byte {
	enc, err := teePayments.abi.Pack("getWalletId", account)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackGetWalletId is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5623b3f5.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (teePayments *TeePayments) TryPackGetWalletId(account ITeePaymentsBasePMWMultisigAccount) ([]byte, error) {
	return teePayments.abi.Pack("getWalletId", account)
}

// UnpackGetWalletId is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5623b3f5.
//
// Solidity: function getWalletId((bytes32,string) _account) view returns(bytes32)
func (teePayments *TeePayments) UnpackGetWalletId(data []byte) ([32]byte, error) {
	out, err := teePayments.abi.Unpack("getWalletId", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governance() view returns(address)
func (teePayments *TeePayments) PackGovernance() []byte {
	enc, err := teePayments.abi.Pack("governance")
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
func (teePayments *TeePayments) TryPackGovernance() ([]byte, error) {
	return teePayments.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (teePayments *TeePayments) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := teePayments.abi.Unpack("governance", data)
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
func (teePayments *TeePayments) PackGovernanceSettings() []byte {
	enc, err := teePayments.abi.Pack("governanceSettings")
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
func (teePayments *TeePayments) TryPackGovernanceSettings() ([]byte, error) {
	return teePayments.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (teePayments *TeePayments) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := teePayments.abi.Unpack("governanceSettings", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c60da1b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function implementation() view returns(address)
func (teePayments *TeePayments) PackImplementation() []byte {
	enc, err := teePayments.abi.Pack("implementation")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackImplementation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5c60da1b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function implementation() view returns(address)
func (teePayments *TeePayments) TryPackImplementation() ([]byte, error) {
	return teePayments.abi.Pack("implementation")
}

// UnpackImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (teePayments *TeePayments) UnpackImplementation(data []byte) (common.Address, error) {
	out, err := teePayments.abi.Unpack("implementation", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc0c53b8b.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (teePayments *TeePayments) PackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) []byte {
	enc, err := teePayments.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackInitialize is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xc0c53b8b.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function initialize(address _governanceSettings, address _initialGovernance, address _addressUpdater) returns()
func (teePayments *TeePayments) TryPackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) ([]byte, error) {
	return teePayments.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePayments *TeePayments) PackIsExecutor(address common.Address) []byte {
	enc, err := teePayments.abi.Pack("isExecutor", address)
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
func (teePayments *TeePayments) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return teePayments.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePayments *TeePayments) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := teePayments.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackPay is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x009ce938.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (teePayments *TeePayments) PackPay(account ITeePaymentsBasePMWMultisigAccount, paymentInstruction ITeePaymentsBasePaymentInstruction, claimBackAddress common.Address) []byte {
	enc, err := teePayments.abi.Pack("pay", account, paymentInstruction, claimBackAddress)
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
func (teePayments *TeePayments) TryPackPay(account ITeePaymentsBasePMWMultisigAccount, paymentInstruction ITeePaymentsBasePaymentInstruction, claimBackAddress common.Address) ([]byte, error) {
	return teePayments.abi.Pack("pay", account, paymentInstruction, claimBackAddress)
}

// UnpackPay is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x009ce938.
//
// Solidity: function pay((bytes32,string) _account, (string,bytes,uint256,uint256,bytes32) _paymentInstruction, address _claimBackAddress) payable returns(uint64 _paymentId)
func (teePayments *TeePayments) UnpackPay(data []byte) (uint64, error) {
	out, err := teePayments.abi.Unpack("pay", data)
	if err != nil {
		return *new(uint64), err
	}
	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	return out0, nil
}

// PackPaymentModel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb9d3dbb.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function paymentModel() pure returns(uint8)
func (teePayments *TeePayments) PackPaymentModel() []byte {
	enc, err := teePayments.abi.Pack("paymentModel")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackPaymentModel is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xbb9d3dbb.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function paymentModel() pure returns(uint8)
func (teePayments *TeePayments) TryPackPaymentModel() ([]byte, error) {
	return teePayments.abi.Pack("paymentModel")
}

// UnpackPaymentModel is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbb9d3dbb.
//
// Solidity: function paymentModel() pure returns(uint8)
func (teePayments *TeePayments) UnpackPaymentModel(data []byte) (uint8, error) {
	out, err := teePayments.abi.Unpack("paymentModel", data)
	if err != nil {
		return *new(uint8), err
	}
	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	return out0, nil
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (teePayments *TeePayments) PackProductionMode() []byte {
	enc, err := teePayments.abi.Pack("productionMode")
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
func (teePayments *TeePayments) TryPackProductionMode() ([]byte, error) {
	return teePayments.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (teePayments *TeePayments) UnpackProductionMode(data []byte) (bool, error) {
	out, err := teePayments.abi.Unpack("productionMode", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teePayments *TeePayments) PackProxiableUUID() []byte {
	enc, err := teePayments.abi.Pack("proxiableUUID")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackProxiableUUID is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x52d1902d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teePayments *TeePayments) TryPackProxiableUUID() ([]byte, error) {
	return teePayments.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teePayments *TeePayments) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := teePayments.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackReissue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1abcad2c.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _paymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, bool _startNew, address _claimBackAddress) payable returns(bool _finalized)
func (teePayments *TeePayments) PackReissue(account ITeePaymentsBasePMWMultisigAccount, paymentId uint64, paymentInstructions []ITeePaymentsBasePaymentInstruction, reissueFeeParams ITeePaymentsBaseReissueFeeParams, startNew bool, claimBackAddress common.Address) []byte {
	enc, err := teePayments.abi.Pack("reissue", account, paymentId, paymentInstructions, reissueFeeParams, startNew, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackReissue is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x1abcad2c.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _paymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, bool _startNew, address _claimBackAddress) payable returns(bool _finalized)
func (teePayments *TeePayments) TryPackReissue(account ITeePaymentsBasePMWMultisigAccount, paymentId uint64, paymentInstructions []ITeePaymentsBasePaymentInstruction, reissueFeeParams ITeePaymentsBaseReissueFeeParams, startNew bool, claimBackAddress common.Address) ([]byte, error) {
	return teePayments.abi.Pack("reissue", account, paymentId, paymentInstructions, reissueFeeParams, startNew, claimBackAddress)
}

// UnpackReissue is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x1abcad2c.
//
// Solidity: function reissue((bytes32,string) _account, uint64 _paymentId, (string,bytes,uint256,uint256,bytes32)[] _paymentInstructions, (uint256[],int16[][],uint16[]) _reissueFeeParams, bool _startNew, address _claimBackAddress) payable returns(bool _finalized)
func (teePayments *TeePayments) UnpackReissue(data []byte) (bool, error) {
	out, err := teePayments.abi.Unpack("reissue", data)
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
func (teePayments *TeePayments) PackSwitchToProductionMode() []byte {
	enc, err := teePayments.abi.Pack("switchToProductionMode")
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
func (teePayments *TeePayments) TryPackSwitchToProductionMode() ([]byte, error) {
	return teePayments.abi.Pack("switchToProductionMode")
}

// PackTeePaymentsConfigVerifier is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf71c6c75.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (teePayments *TeePayments) PackTeePaymentsConfigVerifier() []byte {
	enc, err := teePayments.abi.Pack("teePaymentsConfigVerifier")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTeePaymentsConfigVerifier is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf71c6c75.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (teePayments *TeePayments) TryPackTeePaymentsConfigVerifier() ([]byte, error) {
	return teePayments.abi.Pack("teePaymentsConfigVerifier")
}

// UnpackTeePaymentsConfigVerifier is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xf71c6c75.
//
// Solidity: function teePaymentsConfigVerifier() view returns(address)
func (teePayments *TeePayments) UnpackTeePaymentsConfigVerifier(data []byte) (common.Address, error) {
	out, err := teePayments.abi.Unpack("teePaymentsConfigVerifier", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackTeePaymentsFeeScheduleManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe9e33e8.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function teePaymentsFeeScheduleManager() view returns(address)
func (teePayments *TeePayments) PackTeePaymentsFeeScheduleManager() []byte {
	enc, err := teePayments.abi.Pack("teePaymentsFeeScheduleManager")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTeePaymentsFeeScheduleManager is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xfe9e33e8.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function teePaymentsFeeScheduleManager() view returns(address)
func (teePayments *TeePayments) TryPackTeePaymentsFeeScheduleManager() ([]byte, error) {
	return teePayments.abi.Pack("teePaymentsFeeScheduleManager")
}

// UnpackTeePaymentsFeeScheduleManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfe9e33e8.
//
// Solidity: function teePaymentsFeeScheduleManager() view returns(address)
func (teePayments *TeePayments) UnpackTeePaymentsFeeScheduleManager(data []byte) (common.Address, error) {
	out, err := teePayments.abi.Unpack("teePaymentsFeeScheduleManager", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackTeePaymentsRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaef828de.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (teePayments *TeePayments) PackTeePaymentsRegistry() []byte {
	enc, err := teePayments.abi.Pack("teePaymentsRegistry")
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackTeePaymentsRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaef828de.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (teePayments *TeePayments) TryPackTeePaymentsRegistry() ([]byte, error) {
	return teePayments.abi.Pack("teePaymentsRegistry")
}

// UnpackTeePaymentsRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (teePayments *TeePayments) UnpackTeePaymentsRegistry(data []byte) (common.Address, error) {
	out, err := teePayments.abi.Unpack("teePaymentsRegistry", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackUpdateContractAddresses is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xb00c0b76.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function updateContractAddresses(bytes32[] _contractNameHashes, address[] _contractAddresses) returns()
func (teePayments *TeePayments) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := teePayments.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (teePayments *TeePayments) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return teePayments.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (teePayments *TeePayments) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := teePayments.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (teePayments *TeePayments) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return teePayments.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// TeePaymentsGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePayments contract.
type TeePaymentsGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const TeePaymentsGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (TeePaymentsGovernanceCallTimelocked) ContractEventName() string {
	return TeePaymentsGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (teePayments *TeePayments) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*TeePaymentsGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != teePayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := teePayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePayments.abi.Events[event].Inputs {
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

// TeePaymentsGovernanceInitialised represents a GovernanceInitialised event raised by the TeePayments contract.
type TeePaymentsGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const TeePaymentsGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (TeePaymentsGovernanceInitialised) ContractEventName() string {
	return TeePaymentsGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (teePayments *TeePayments) UnpackGovernanceInitialisedEvent(log *types.Log) (*TeePaymentsGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != teePayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := teePayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePayments.abi.Events[event].Inputs {
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

// TeePaymentsGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePayments contract.
type TeePaymentsGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const TeePaymentsGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (TeePaymentsGovernedProductionModeEntered) ContractEventName() string {
	return TeePaymentsGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (teePayments *TeePayments) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*TeePaymentsGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != teePayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := teePayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePayments.abi.Events[event].Inputs {
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

// TeePaymentsInitialized represents a Initialized event raised by the TeePayments contract.
type TeePaymentsInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const TeePaymentsInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (TeePaymentsInitialized) ContractEventName() string {
	return TeePaymentsInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (teePayments *TeePayments) UnpackInitializedEvent(log *types.Log) (*TeePaymentsInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != teePayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsInitialized)
	if len(log.Data) > 0 {
		if err := teePayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePayments.abi.Events[event].Inputs {
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

// TeePaymentsPMWMultisigAccountAdded represents a PMWMultisigAccountAdded event raised by the TeePayments contract.
type TeePaymentsPMWMultisigAccountAdded struct {
	WalletId             [32]byte
	SourceId             [32]byte
	AccountAddress       string
	AuthorizationAddress common.Address
	InitialNonce         uint64
	Raw                  *types.Log // Blockchain specific contextual infos
}

const TeePaymentsPMWMultisigAccountAddedEventName = "PMWMultisigAccountAdded"

// ContractEventName returns the user-defined event name.
func (TeePaymentsPMWMultisigAccountAdded) ContractEventName() string {
	return TeePaymentsPMWMultisigAccountAddedEventName
}

// UnpackPMWMultisigAccountAddedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event PMWMultisigAccountAdded(bytes32 indexed walletId, bytes32 sourceId, string accountAddress, address authorizationAddress, uint64 initialNonce)
func (teePayments *TeePayments) UnpackPMWMultisigAccountAddedEvent(log *types.Log) (*TeePaymentsPMWMultisigAccountAdded, error) {
	event := "PMWMultisigAccountAdded"
	if len(log.Topics) == 0 || log.Topics[0] != teePayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsPMWMultisigAccountAdded)
	if len(log.Data) > 0 {
		if err := teePayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePayments.abi.Events[event].Inputs {
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

// TeePaymentsTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePayments contract.
type TeePaymentsTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeePaymentsTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (TeePaymentsTimelockedGovernanceCallCanceled) ContractEventName() string {
	return TeePaymentsTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (teePayments *TeePayments) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*TeePaymentsTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != teePayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := teePayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePayments.abi.Events[event].Inputs {
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

// TeePaymentsTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePayments contract.
type TeePaymentsTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeePaymentsTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (TeePaymentsTimelockedGovernanceCallExecuted) ContractEventName() string {
	return TeePaymentsTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (teePayments *TeePayments) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*TeePaymentsTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != teePayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := teePayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePayments.abi.Events[event].Inputs {
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

// TeePaymentsUpgraded represents a Upgraded event raised by the TeePayments contract.
type TeePaymentsUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const TeePaymentsUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (TeePaymentsUpgraded) ContractEventName() string {
	return TeePaymentsUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (teePayments *TeePayments) UnpackUpgradedEvent(log *types.Log) (*TeePaymentsUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 || log.Topics[0] != teePayments.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsUpgraded)
	if len(log.Data) > 0 {
		if err := teePayments.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePayments.abi.Events[event].Inputs {
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
func (teePayments *TeePayments) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], teePayments.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return teePayments.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["AlreadyInProductionMode"].ID.Bytes()[:4]) {
		return teePayments.UnpackAlreadyInProductionModeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["AuthorizationAddressZero"].ID.Bytes()[:4]) {
		return teePayments.UnpackAuthorizationAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return teePayments.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return teePayments.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return teePayments.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["GovernedAddressZero"].ID.Bytes()[:4]) {
		return teePayments.UnpackGovernedAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["GovernedAlreadyInitialized"].ID.Bytes()[:4]) {
		return teePayments.UnpackGovernedAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return teePayments.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["InvalidPaymentId"].ID.Bytes()[:4]) {
		return teePayments.UnpackInvalidPaymentIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["InvalidPaymentInstructionCount"].ID.Bytes()[:4]) {
		return teePayments.UnpackInvalidPaymentInstructionCountError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["InvalidRecipientAddress"].ID.Bytes()[:4]) {
		return teePayments.UnpackInvalidRecipientAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["LengthsMismatch"].ID.Bytes()[:4]) {
		return teePayments.UnpackLengthsMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return teePayments.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["OnlyAuthorizationAddress"].ID.Bytes()[:4]) {
		return teePayments.UnpackOnlyAuthorizationAddressError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["OnlyExecutor"].ID.Bytes()[:4]) {
		return teePayments.UnpackOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return teePayments.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return teePayments.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["OnlySystemExtensionId"].ID.Bytes()[:4]) {
		return teePayments.UnpackOnlySystemExtensionIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["OnlyWalletOwner"].ID.Bytes()[:4]) {
		return teePayments.UnpackOnlyWalletOwnerError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["PMWMultisigAccountAddressAlreadySet"].ID.Bytes()[:4]) {
		return teePayments.UnpackPMWMultisigAccountAddressAlreadySetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["PMWMultisigAccountNotRegistered"].ID.Bytes()[:4]) {
		return teePayments.UnpackPMWMultisigAccountNotRegisteredError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["PaymentAmountZero"].ID.Bytes()[:4]) {
		return teePayments.UnpackPaymentAmountZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["PaymentHashMismatch"].ID.Bytes()[:4]) {
		return teePayments.UnpackPaymentHashMismatchError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["StartNewRequired"].ID.Bytes()[:4]) {
		return teePayments.UnpackStartNewRequiredError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["TimelockCallNotFound"].ID.Bytes()[:4]) {
		return teePayments.UnpackTimelockCallNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["TimelockNotAllowedYet"].ID.Bytes()[:4]) {
		return teePayments.UnpackTimelockNotAllowedYetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return teePayments.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return teePayments.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["UnsupportedSourceId"].ID.Bytes()[:4]) {
		return teePayments.UnpackUnsupportedSourceIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["WalletNotInProduction"].ID.Bytes()[:4]) {
		return teePayments.UnpackWalletNotInProductionError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePayments.abi.Errors["WrongKeyType"].ID.Bytes()[:4]) {
		return teePayments.UnpackWrongKeyTypeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// TeePaymentsAddressEmptyCode represents a AddressEmptyCode error raised by the TeePayments contract.
type TeePaymentsAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func TeePaymentsAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (teePayments *TeePayments) UnpackAddressEmptyCodeError(raw []byte) (*TeePaymentsAddressEmptyCode, error) {
	out := new(TeePaymentsAddressEmptyCode)
	if err := teePayments.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsAlreadyInProductionMode represents a AlreadyInProductionMode error raised by the TeePayments contract.
type TeePaymentsAlreadyInProductionMode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInProductionMode()
func TeePaymentsAlreadyInProductionModeErrorID() common.Hash {
	return common.HexToHash("0x72a23ec2b1fdbc3de7d5f411b572fcf8ad8409803196930f63ddf40a2a2b2594")
}

// UnpackAlreadyInProductionModeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInProductionMode()
func (teePayments *TeePayments) UnpackAlreadyInProductionModeError(raw []byte) (*TeePaymentsAlreadyInProductionMode, error) {
	out := new(TeePaymentsAlreadyInProductionMode)
	if err := teePayments.abi.UnpackIntoInterface(out, "AlreadyInProductionMode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsAuthorizationAddressZero represents a AuthorizationAddressZero error raised by the TeePayments contract.
type TeePaymentsAuthorizationAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AuthorizationAddressZero()
func TeePaymentsAuthorizationAddressZeroErrorID() common.Hash {
	return common.HexToHash("0xe4472ca7994bab3d6f38c9bdc243c438c59c017951973a129c8674da85c299a2")
}

// UnpackAuthorizationAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AuthorizationAddressZero()
func (teePayments *TeePayments) UnpackAuthorizationAddressZeroError(raw []byte) (*TeePaymentsAuthorizationAddressZero, error) {
	out := new(TeePaymentsAuthorizationAddressZero)
	if err := teePayments.abi.UnpackIntoInterface(out, "AuthorizationAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the TeePayments contract.
type TeePaymentsERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func TeePaymentsERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (teePayments *TeePayments) UnpackERC1967InvalidImplementationError(raw []byte) (*TeePaymentsERC1967InvalidImplementation, error) {
	out := new(TeePaymentsERC1967InvalidImplementation)
	if err := teePayments.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsERC1967NonPayable represents a ERC1967NonPayable error raised by the TeePayments contract.
type TeePaymentsERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func TeePaymentsERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (teePayments *TeePayments) UnpackERC1967NonPayableError(raw []byte) (*TeePaymentsERC1967NonPayable, error) {
	out := new(TeePaymentsERC1967NonPayable)
	if err := teePayments.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsFailedCall represents a FailedCall error raised by the TeePayments contract.
type TeePaymentsFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func TeePaymentsFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (teePayments *TeePayments) UnpackFailedCallError(raw []byte) (*TeePaymentsFailedCall, error) {
	out := new(TeePaymentsFailedCall)
	if err := teePayments.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsGovernedAddressZero represents a GovernedAddressZero error raised by the TeePayments contract.
type TeePaymentsGovernedAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAddressZero()
func TeePaymentsGovernedAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x695725f6b429bac0cc5a2fa88214e31520e0c69f2fb0e2fb881a7130a564f49f")
}

// UnpackGovernedAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAddressZero()
func (teePayments *TeePayments) UnpackGovernedAddressZeroError(raw []byte) (*TeePaymentsGovernedAddressZero, error) {
	out := new(TeePaymentsGovernedAddressZero)
	if err := teePayments.abi.UnpackIntoInterface(out, "GovernedAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsGovernedAlreadyInitialized represents a GovernedAlreadyInitialized error raised by the TeePayments contract.
type TeePaymentsGovernedAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAlreadyInitialized()
func TeePaymentsGovernedAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0xb2f0162156d83828197776d56bbfba0a7a9d5936fad70f0a3014ae18cc7aa33c")
}

// UnpackGovernedAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAlreadyInitialized()
func (teePayments *TeePayments) UnpackGovernedAlreadyInitializedError(raw []byte) (*TeePaymentsGovernedAlreadyInitialized, error) {
	out := new(TeePaymentsGovernedAlreadyInitialized)
	if err := teePayments.abi.UnpackIntoInterface(out, "GovernedAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsInvalidInitialization represents a InvalidInitialization error raised by the TeePayments contract.
type TeePaymentsInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func TeePaymentsInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (teePayments *TeePayments) UnpackInvalidInitializationError(raw []byte) (*TeePaymentsInvalidInitialization, error) {
	out := new(TeePaymentsInvalidInitialization)
	if err := teePayments.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsInvalidPaymentId represents a InvalidPaymentId error raised by the TeePayments contract.
type TeePaymentsInvalidPaymentId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPaymentId()
func TeePaymentsInvalidPaymentIdErrorID() common.Hash {
	return common.HexToHash("0xb11be8c4d37d5cf7116085b4f06afe80345db997f69552f4e6a826a12eea0c0f")
}

// UnpackInvalidPaymentIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPaymentId()
func (teePayments *TeePayments) UnpackInvalidPaymentIdError(raw []byte) (*TeePaymentsInvalidPaymentId, error) {
	out := new(TeePaymentsInvalidPaymentId)
	if err := teePayments.abi.UnpackIntoInterface(out, "InvalidPaymentId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsInvalidPaymentInstructionCount represents a InvalidPaymentInstructionCount error raised by the TeePayments contract.
type TeePaymentsInvalidPaymentInstructionCount struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidPaymentInstructionCount()
func TeePaymentsInvalidPaymentInstructionCountErrorID() common.Hash {
	return common.HexToHash("0xb944aca0ae5b99f4666f05d50060098c29717bc0c0e1c515201a3cee72b4f5c3")
}

// UnpackInvalidPaymentInstructionCountError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidPaymentInstructionCount()
func (teePayments *TeePayments) UnpackInvalidPaymentInstructionCountError(raw []byte) (*TeePaymentsInvalidPaymentInstructionCount, error) {
	out := new(TeePaymentsInvalidPaymentInstructionCount)
	if err := teePayments.abi.UnpackIntoInterface(out, "InvalidPaymentInstructionCount", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsInvalidRecipientAddress represents a InvalidRecipientAddress error raised by the TeePayments contract.
type TeePaymentsInvalidRecipientAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRecipientAddress()
func TeePaymentsInvalidRecipientAddressErrorID() common.Hash {
	return common.HexToHash("0x44d99fea8d272f0b82a2cdf27d16958dace6a85ad6302ff443201a39bfc4820a")
}

// UnpackInvalidRecipientAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRecipientAddress()
func (teePayments *TeePayments) UnpackInvalidRecipientAddressError(raw []byte) (*TeePaymentsInvalidRecipientAddress, error) {
	out := new(TeePaymentsInvalidRecipientAddress)
	if err := teePayments.abi.UnpackIntoInterface(out, "InvalidRecipientAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsLengthsMismatch represents a LengthsMismatch error raised by the TeePayments contract.
type TeePaymentsLengthsMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error LengthsMismatch()
func TeePaymentsLengthsMismatchErrorID() common.Hash {
	return common.HexToHash("0x586cb9e1ae7e957b53bdff143b03d56db14d810b68acd40d5f39245aac29e94f")
}

// UnpackLengthsMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error LengthsMismatch()
func (teePayments *TeePayments) UnpackLengthsMismatchError(raw []byte) (*TeePaymentsLengthsMismatch, error) {
	out := new(TeePaymentsLengthsMismatch)
	if err := teePayments.abi.UnpackIntoInterface(out, "LengthsMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsNotInitializing represents a NotInitializing error raised by the TeePayments contract.
type TeePaymentsNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func TeePaymentsNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (teePayments *TeePayments) UnpackNotInitializingError(raw []byte) (*TeePaymentsNotInitializing, error) {
	out := new(TeePaymentsNotInitializing)
	if err := teePayments.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsOnlyAuthorizationAddress represents a OnlyAuthorizationAddress error raised by the TeePayments contract.
type TeePaymentsOnlyAuthorizationAddress struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyAuthorizationAddress()
func TeePaymentsOnlyAuthorizationAddressErrorID() common.Hash {
	return common.HexToHash("0xdc65e71c09e77be7539dadee1dc86849e01c3e995f0c8d49e625476c2a1b1ed5")
}

// UnpackOnlyAuthorizationAddressError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyAuthorizationAddress()
func (teePayments *TeePayments) UnpackOnlyAuthorizationAddressError(raw []byte) (*TeePaymentsOnlyAuthorizationAddress, error) {
	out := new(TeePaymentsOnlyAuthorizationAddress)
	if err := teePayments.abi.UnpackIntoInterface(out, "OnlyAuthorizationAddress", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsOnlyExecutor represents a OnlyExecutor error raised by the TeePayments contract.
type TeePaymentsOnlyExecutor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExecutor()
func TeePaymentsOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x7fb6be02aaf40e420b0e7eb67e43a3feee949e8932a3c64e33b591a436670841")
}

// UnpackOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExecutor()
func (teePayments *TeePayments) UnpackOnlyExecutorError(raw []byte) (*TeePaymentsOnlyExecutor, error) {
	out := new(TeePaymentsOnlyExecutor)
	if err := teePayments.abi.UnpackIntoInterface(out, "OnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsOnlyGovernance represents a OnlyGovernance error raised by the TeePayments contract.
type TeePaymentsOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func TeePaymentsOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (teePayments *TeePayments) UnpackOnlyGovernanceError(raw []byte) (*TeePaymentsOnlyGovernance, error) {
	out := new(TeePaymentsOnlyGovernance)
	if err := teePayments.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the TeePayments contract.
type TeePaymentsOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func TeePaymentsOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (teePayments *TeePayments) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*TeePaymentsOnlyProductionOrPausedStatus, error) {
	out := new(TeePaymentsOnlyProductionOrPausedStatus)
	if err := teePayments.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsOnlySystemExtensionId represents a OnlySystemExtensionId error raised by the TeePayments contract.
type TeePaymentsOnlySystemExtensionId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemExtensionId()
func TeePaymentsOnlySystemExtensionIdErrorID() common.Hash {
	return common.HexToHash("0xbd49094d5a86ee3f3c4b39274f5f98391dde2cd16618a917e121333ac3d43ee4")
}

// UnpackOnlySystemExtensionIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemExtensionId()
func (teePayments *TeePayments) UnpackOnlySystemExtensionIdError(raw []byte) (*TeePaymentsOnlySystemExtensionId, error) {
	out := new(TeePaymentsOnlySystemExtensionId)
	if err := teePayments.abi.UnpackIntoInterface(out, "OnlySystemExtensionId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsOnlyWalletOwner represents a OnlyWalletOwner error raised by the TeePayments contract.
type TeePaymentsOnlyWalletOwner struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyWalletOwner()
func TeePaymentsOnlyWalletOwnerErrorID() common.Hash {
	return common.HexToHash("0x3132d62dfddd1d8b6a9dd805001453bbb6a84b746ebe51fd650df1a6704a8143")
}

// UnpackOnlyWalletOwnerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyWalletOwner()
func (teePayments *TeePayments) UnpackOnlyWalletOwnerError(raw []byte) (*TeePaymentsOnlyWalletOwner, error) {
	out := new(TeePaymentsOnlyWalletOwner)
	if err := teePayments.abi.UnpackIntoInterface(out, "OnlyWalletOwner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsPMWMultisigAccountAddressAlreadySet represents a PMWMultisigAccountAddressAlreadySet error raised by the TeePayments contract.
type TeePaymentsPMWMultisigAccountAddressAlreadySet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PMWMultisigAccountAddressAlreadySet()
func TeePaymentsPMWMultisigAccountAddressAlreadySetErrorID() common.Hash {
	return common.HexToHash("0xe3e804ecd31f69f9fde81b0b1994f9d3057e672dfe1cb39f32b68f667aa752e6")
}

// UnpackPMWMultisigAccountAddressAlreadySetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PMWMultisigAccountAddressAlreadySet()
func (teePayments *TeePayments) UnpackPMWMultisigAccountAddressAlreadySetError(raw []byte) (*TeePaymentsPMWMultisigAccountAddressAlreadySet, error) {
	out := new(TeePaymentsPMWMultisigAccountAddressAlreadySet)
	if err := teePayments.abi.UnpackIntoInterface(out, "PMWMultisigAccountAddressAlreadySet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsPMWMultisigAccountNotRegistered represents a PMWMultisigAccountNotRegistered error raised by the TeePayments contract.
type TeePaymentsPMWMultisigAccountNotRegistered struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PMWMultisigAccountNotRegistered()
func TeePaymentsPMWMultisigAccountNotRegisteredErrorID() common.Hash {
	return common.HexToHash("0x9d78814a2f7e263f4d4cc3000f076b4c90004105130f1880c0dc94acd2464be9")
}

// UnpackPMWMultisigAccountNotRegisteredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PMWMultisigAccountNotRegistered()
func (teePayments *TeePayments) UnpackPMWMultisigAccountNotRegisteredError(raw []byte) (*TeePaymentsPMWMultisigAccountNotRegistered, error) {
	out := new(TeePaymentsPMWMultisigAccountNotRegistered)
	if err := teePayments.abi.UnpackIntoInterface(out, "PMWMultisigAccountNotRegistered", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsPaymentAmountZero represents a PaymentAmountZero error raised by the TeePayments contract.
type TeePaymentsPaymentAmountZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentAmountZero()
func TeePaymentsPaymentAmountZeroErrorID() common.Hash {
	return common.HexToHash("0xa74a1c28a64fc70e7d718ae1b3af31f436e1c55b3754d607ef7ec829ec74034c")
}

// UnpackPaymentAmountZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentAmountZero()
func (teePayments *TeePayments) UnpackPaymentAmountZeroError(raw []byte) (*TeePaymentsPaymentAmountZero, error) {
	out := new(TeePaymentsPaymentAmountZero)
	if err := teePayments.abi.UnpackIntoInterface(out, "PaymentAmountZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsPaymentHashMismatch represents a PaymentHashMismatch error raised by the TeePayments contract.
type TeePaymentsPaymentHashMismatch struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error PaymentHashMismatch()
func TeePaymentsPaymentHashMismatchErrorID() common.Hash {
	return common.HexToHash("0x85628ea5f0ea30f18d9205be9d6ae9bd4b189b3dc00aeba346ab8f4c0e25674e")
}

// UnpackPaymentHashMismatchError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error PaymentHashMismatch()
func (teePayments *TeePayments) UnpackPaymentHashMismatchError(raw []byte) (*TeePaymentsPaymentHashMismatch, error) {
	out := new(TeePaymentsPaymentHashMismatch)
	if err := teePayments.abi.UnpackIntoInterface(out, "PaymentHashMismatch", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsStartNewRequired represents a StartNewRequired error raised by the TeePayments contract.
type TeePaymentsStartNewRequired struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error StartNewRequired()
func TeePaymentsStartNewRequiredErrorID() common.Hash {
	return common.HexToHash("0x5bf0683c38d9f6b6ae0bb670cff384af933a27d8f571ab8767ae09d925a434d5")
}

// UnpackStartNewRequiredError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error StartNewRequired()
func (teePayments *TeePayments) UnpackStartNewRequiredError(raw []byte) (*TeePaymentsStartNewRequired, error) {
	out := new(TeePaymentsStartNewRequired)
	if err := teePayments.abi.UnpackIntoInterface(out, "StartNewRequired", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsTimelockCallNotFound represents a TimelockCallNotFound error raised by the TeePayments contract.
type TeePaymentsTimelockCallNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockCallNotFound()
func TeePaymentsTimelockCallNotFoundErrorID() common.Hash {
	return common.HexToHash("0x262ac987dd09120380d393a0c52cdef1ca58949d1107dc26d8da9f7b527eb33d")
}

// UnpackTimelockCallNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockCallNotFound()
func (teePayments *TeePayments) UnpackTimelockCallNotFoundError(raw []byte) (*TeePaymentsTimelockCallNotFound, error) {
	out := new(TeePaymentsTimelockCallNotFound)
	if err := teePayments.abi.UnpackIntoInterface(out, "TimelockCallNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsTimelockNotAllowedYet represents a TimelockNotAllowedYet error raised by the TeePayments contract.
type TeePaymentsTimelockNotAllowedYet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockNotAllowedYet()
func TeePaymentsTimelockNotAllowedYetErrorID() common.Hash {
	return common.HexToHash("0x6124e5c2c545bec1d02048c68500eebdabeae0cc69995a629ebccf8a62f69083")
}

// UnpackTimelockNotAllowedYetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockNotAllowedYet()
func (teePayments *TeePayments) UnpackTimelockNotAllowedYetError(raw []byte) (*TeePaymentsTimelockNotAllowedYet, error) {
	out := new(TeePaymentsTimelockNotAllowedYet)
	if err := teePayments.abi.UnpackIntoInterface(out, "TimelockNotAllowedYet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the TeePayments contract.
type TeePaymentsUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func TeePaymentsUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (teePayments *TeePayments) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*TeePaymentsUUPSUnauthorizedCallContext, error) {
	out := new(TeePaymentsUUPSUnauthorizedCallContext)
	if err := teePayments.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the TeePayments contract.
type TeePaymentsUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func TeePaymentsUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (teePayments *TeePayments) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*TeePaymentsUUPSUnsupportedProxiableUUID, error) {
	out := new(TeePaymentsUUPSUnsupportedProxiableUUID)
	if err := teePayments.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsUnsupportedSourceId represents a UnsupportedSourceId error raised by the TeePayments contract.
type TeePaymentsUnsupportedSourceId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedSourceId()
func TeePaymentsUnsupportedSourceIdErrorID() common.Hash {
	return common.HexToHash("0xb9c2749e785a86b27011ac62f2eeace91e285734582a3d21db2e093d45e3124a")
}

// UnpackUnsupportedSourceIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedSourceId()
func (teePayments *TeePayments) UnpackUnsupportedSourceIdError(raw []byte) (*TeePaymentsUnsupportedSourceId, error) {
	out := new(TeePaymentsUnsupportedSourceId)
	if err := teePayments.abi.UnpackIntoInterface(out, "UnsupportedSourceId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsWalletNotInProduction represents a WalletNotInProduction error raised by the TeePayments contract.
type TeePaymentsWalletNotInProduction struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WalletNotInProduction()
func TeePaymentsWalletNotInProductionErrorID() common.Hash {
	return common.HexToHash("0x2b7f97d6ba9d96a13708bf15f7c7ed50584fa90c4e3f979e82b455f65581f045")
}

// UnpackWalletNotInProductionError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WalletNotInProduction()
func (teePayments *TeePayments) UnpackWalletNotInProductionError(raw []byte) (*TeePaymentsWalletNotInProduction, error) {
	out := new(TeePaymentsWalletNotInProduction)
	if err := teePayments.abi.UnpackIntoInterface(out, "WalletNotInProduction", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsWrongKeyType represents a WrongKeyType error raised by the TeePayments contract.
type TeePaymentsWrongKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WrongKeyType()
func TeePaymentsWrongKeyTypeErrorID() common.Hash {
	return common.HexToHash("0xfb224bce70b4062685f138c7eed15a72b903495475f7cd7843835692a7c58bbd")
}

// UnpackWrongKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WrongKeyType()
func (teePayments *TeePayments) UnpackWrongKeyTypeError(raw []byte) (*TeePaymentsWrongKeyType, error) {
	out := new(TeePaymentsWrongKeyType)
	if err := teePayments.abi.UnpackIntoInterface(out, "WrongKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}
