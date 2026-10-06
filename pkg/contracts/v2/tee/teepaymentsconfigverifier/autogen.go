// Code generated via abigen V2 - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package teepaymentsconfigverifier

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
	Status         uint8
	AccountAddress string
}

// Signature is an auto generated low-level Go binding around an user-defined struct.
type Signature struct {
	V uint8
	R [32]byte
	S [32]byte
}

// TeePaymentsConfigVerifierMetaData contains all meta data concerning the TeePaymentsConfigVerifier contract.
var TeePaymentsConfigVerifierMetaData = bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccountAddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyInProductionMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AnchorSetEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CosignersThresholdNotMet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ERC1967InvalidImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ERC1967NonPayable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernedAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cosigner\",\"type\":\"address\"}],\"name\":\"InvalidCosigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestBody\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSigningPolicy\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyExecutor\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyProductionOrPausedStatus\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySystemExtensionId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockCallNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TimelockNotAllowedYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UUPSUnauthorizedCallContext\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"UUPSUnsupportedProxiableUUID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnsupportedSourceId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongKeyType\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encodedCall\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"allowedAfterTimestamp\",\"type\":\"uint256\"}],\"name\":\"GovernanceCallTimelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initialGovernance\",\"type\":\"address\"}],\"name\":\"GovernanceInitialised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governanceSettings\",\"type\":\"address\"}],\"name\":\"GovernedProductionModeEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"encodedCallHash\",\"type\":\"bytes32\"}],\"name\":\"TimelockedGovernanceCallExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UPGRADE_INTERFACE_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"cancelGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_encodedCall\",\"type\":\"bytes\"}],\"name\":\"executeGovernanceCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2Hub\",\"outputs\":[{\"internalType\":\"contractIFdc2Hub\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fdc2Verification\",\"outputs\":[{\"internalType\":\"contractIFdc2Verification\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareSystemsManager\",\"outputs\":[{\"internalType\":\"contractIFlareSystemsManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flareTeeManager\",\"outputs\":[{\"internalType\":\"contractIIFlareTeeManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressUpdater\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governanceSettings\",\"outputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIGovernanceSettings\",\"name\":\"_governanceSettings\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_initialGovernance\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_addressUpdater\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isExecutor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"productionMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_accountAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proofOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestAccountConfiguredAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_accountIndex\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Anchor[]\",\"name\":\"_anchors\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"_testOnTeeId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_proofOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_claimBackAddress\",\"type\":\"address\"}],\"name\":\"requestUtxoConfiguredAttestation\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchToProductionMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"teePaymentsRegistry\",\"outputs\":[{\"internalType\":\"contractITeePaymentsRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_contractNameHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"_contractAddresses\",\"type\":\"address[]\"}],\"name\":\"updateContractAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigAccountConfigured.PMWMultisigAccountStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigAccountConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyAccountConfiguredProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_walletId\",\"type\":\"bytes32\"},{\"components\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"signingPolicySignatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"teeSignatures\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structSignature[]\",\"name\":\"cosignerSignatures\",\"type\":\"tuple[]\"}],\"internalType\":\"structIFdc2Verification.Fdc2Signatures\",\"name\":\"signatures\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"attestationType\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sourceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"thresholdBIPS\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"proofOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"cosigners\",\"type\":\"address[]\"},{\"internalType\":\"uint64\",\"name\":\"cosignersThreshold\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIFdc2Hub.Fdc2ResponseHeader\",\"name\":\"header\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"accountIndex\",\"type\":\"uint32\"},{\"internalType\":\"bytes[]\",\"name\":\"publicKeys\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"threshold\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"genesisAnchorTxid\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"genesisAnchorVout\",\"type\":\"uint32\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Anchor[]\",\"name\":\"anchors\",\"type\":\"tuple[]\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.RequestBody\",\"name\":\"requestBody\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumIPMWMultisigUtxoConfigured.PMWMultisigUtxoStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.ResponseBody\",\"name\":\"responseBody\",\"type\":\"tuple\"}],\"internalType\":\"structIPMWMultisigUtxoConfigured.Proof\",\"name\":\"_proof\",\"type\":\"tuple\"}],\"name\":\"verifyUtxoConfiguredProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	ID:  "TeePaymentsConfigVerifier",
}

// TeePaymentsConfigVerifier is an auto generated Go binding around an Ethereum contract.
type TeePaymentsConfigVerifier struct {
	abi abi.ABI
}

// NewTeePaymentsConfigVerifier creates a new instance of TeePaymentsConfigVerifier.
func NewTeePaymentsConfigVerifier() *TeePaymentsConfigVerifier {
	parsed, err := TeePaymentsConfigVerifierMetaData.ParseABI()
	if err != nil {
		panic(errors.New("invalid ABI: " + err.Error()))
	}
	return &TeePaymentsConfigVerifier{abi: *parsed}
}

// Instance creates a wrapper for a deployed contract instance at the given address.
// Use this to create the instance object passed to abigen v2 library functions Call, Transact, etc.
func (c *TeePaymentsConfigVerifier) Instance(backend bind.ContractBackend, addr common.Address) *bind.BoundContract {
	return bind.NewBoundContract(addr, c.abi, backend, backend, backend)
}

// PackUPGRADEINTERFACEVERSION is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xad3cb1cc.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackUPGRADEINTERFACEVERSION() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("UPGRADE_INTERFACE_VERSION")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackUPGRADEINTERFACEVERSION() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("UPGRADE_INTERFACE_VERSION")
}

// UnpackUPGRADEINTERFACEVERSION is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackUPGRADEINTERFACEVERSION(data []byte) (string, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("UPGRADE_INTERFACE_VERSION", data)
	if err != nil {
		return *new(string), err
	}
	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	return out0, nil
}

// PackCancelGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x16fc2f6d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function cancelGovernanceCall(bytes _encodedCall) returns()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackCancelGovernanceCall(encodedCall []byte) []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("cancelGovernanceCall", encodedCall)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackCancelGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("cancelGovernanceCall", encodedCall)
}

// PackExecuteGovernanceCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x20c5f99d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function executeGovernanceCall(bytes _encodedCall) returns()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackExecuteGovernanceCall(encodedCall []byte) []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("executeGovernanceCall", encodedCall)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackExecuteGovernanceCall(encodedCall []byte) ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("executeGovernanceCall", encodedCall)
}

// PackFdc2Hub is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xa7566ff3.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function fdc2Hub() view returns(address)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackFdc2Hub() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("fdc2Hub")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackFdc2Hub() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("fdc2Hub")
}

// UnpackFdc2Hub is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xa7566ff3.
//
// Solidity: function fdc2Hub() view returns(address)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackFdc2Hub(data []byte) (common.Address, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("fdc2Hub", data)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackFdc2Verification() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("fdc2Verification")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackFdc2Verification() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("fdc2Verification")
}

// UnpackFdc2Verification is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xbf2e9839.
//
// Solidity: function fdc2Verification() view returns(address)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackFdc2Verification(data []byte) (common.Address, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("fdc2Verification", data)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackFlareSystemsManager() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("flareSystemsManager")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackFlareSystemsManager() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("flareSystemsManager")
}

// UnpackFlareSystemsManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xfaae7fc9.
//
// Solidity: function flareSystemsManager() view returns(address)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackFlareSystemsManager(data []byte) (common.Address, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("flareSystemsManager", data)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackFlareTeeManager() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("flareTeeManager")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackFlareTeeManager() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("flareTeeManager")
}

// UnpackFlareTeeManager is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x453f7ab4.
//
// Solidity: function flareTeeManager() view returns(address)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackFlareTeeManager(data []byte) (common.Address, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("flareTeeManager", data)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackGetAddressUpdater() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("getAddressUpdater")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackGetAddressUpdater() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("getAddressUpdater")
}

// UnpackGetAddressUpdater is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5267a15d.
//
// Solidity: function getAddressUpdater() view returns(address _addressUpdater)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackGetAddressUpdater(data []byte) (common.Address, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("getAddressUpdater", data)
	if err != nil {
		return *new(common.Address), err
	}
	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	return out0, nil
}

// PackGovernance is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x5aa6e675.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function governance() view returns(address)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackGovernance() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("governance")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackGovernance() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("governance")
}

// UnpackGovernance is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackGovernance(data []byte) (common.Address, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("governance", data)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackGovernanceSettings() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("governanceSettings")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackGovernanceSettings() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("governanceSettings")
}

// UnpackGovernanceSettings is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x62354e03.
//
// Solidity: function governanceSettings() view returns(address)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackGovernanceSettings(data []byte) (common.Address, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("governanceSettings", data)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackImplementation() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("implementation")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackImplementation() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("implementation")
}

// UnpackImplementation is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackImplementation(data []byte) (common.Address, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("implementation", data)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackInitialize(governanceSettings common.Address, initialGovernance common.Address, addressUpdater common.Address) ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("initialize", governanceSettings, initialGovernance, addressUpdater)
}

// PackIsExecutor is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xdebfda30.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackIsExecutor(address common.Address) []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("isExecutor", address)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackIsExecutor(address common.Address) ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("isExecutor", address)
}

// UnpackIsExecutor is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xdebfda30.
//
// Solidity: function isExecutor(address _address) view returns(bool)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackIsExecutor(data []byte) (bool, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("isExecutor", data)
	if err != nil {
		return *new(bool), err
	}
	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	return out0, nil
}

// PackProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xe17f212e.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function productionMode() view returns(bool)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackProductionMode() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("productionMode")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackProductionMode() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("productionMode")
}

// UnpackProductionMode is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xe17f212e.
//
// Solidity: function productionMode() view returns(bool)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackProductionMode(data []byte) (bool, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("productionMode", data)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackProxiableUUID() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("proxiableUUID")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackProxiableUUID() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("proxiableUUID")
}

// UnpackProxiableUUID is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackProxiableUUID(data []byte) ([32]byte, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("proxiableUUID", data)
	if err != nil {
		return *new([32]byte), err
	}
	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	return out0, nil
}

// PackRequestAccountConfiguredAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x17005854.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function requestAccountConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackRequestAccountConfiguredAttestation(walletId [32]byte, sourceId [32]byte, accountAddress string, testOnTeeId common.Address, proofOwner common.Address, claimBackAddress common.Address) []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("requestAccountConfiguredAttestation", walletId, sourceId, accountAddress, testOnTeeId, proofOwner, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRequestAccountConfiguredAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x17005854.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function requestAccountConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, string _accountAddress, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackRequestAccountConfiguredAttestation(walletId [32]byte, sourceId [32]byte, accountAddress string, testOnTeeId common.Address, proofOwner common.Address, claimBackAddress common.Address) ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("requestAccountConfiguredAttestation", walletId, sourceId, accountAddress, testOnTeeId, proofOwner, claimBackAddress)
}

// PackRequestUtxoConfiguredAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f12700d.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function requestUtxoConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, uint32 _accountIndex, (bytes32,uint32)[] _anchors, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackRequestUtxoConfiguredAttestation(walletId [32]byte, sourceId [32]byte, accountIndex uint32, anchors []IPMWMultisigUtxoConfiguredAnchor, testOnTeeId common.Address, proofOwner common.Address, claimBackAddress common.Address) []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("requestUtxoConfiguredAttestation", walletId, sourceId, accountIndex, anchors, testOnTeeId, proofOwner, claimBackAddress)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackRequestUtxoConfiguredAttestation is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x0f12700d.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function requestUtxoConfiguredAttestation(bytes32 _walletId, bytes32 _sourceId, uint32 _accountIndex, (bytes32,uint32)[] _anchors, address _testOnTeeId, address _proofOwner, address _claimBackAddress) payable returns()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackRequestUtxoConfiguredAttestation(walletId [32]byte, sourceId [32]byte, accountIndex uint32, anchors []IPMWMultisigUtxoConfiguredAnchor, testOnTeeId common.Address, proofOwner common.Address, claimBackAddress common.Address) ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("requestUtxoConfiguredAttestation", walletId, sourceId, accountIndex, anchors, testOnTeeId, proofOwner, claimBackAddress)
}

// PackSwitchToProductionMode is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xf5a98383.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function switchToProductionMode() returns()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackSwitchToProductionMode() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("switchToProductionMode")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackSwitchToProductionMode() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("switchToProductionMode")
}

// PackTeePaymentsRegistry is the Go binding used to pack the parameters required for calling
// the contract method with ID 0xaef828de.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackTeePaymentsRegistry() []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("teePaymentsRegistry")
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackTeePaymentsRegistry() ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("teePaymentsRegistry")
}

// UnpackTeePaymentsRegistry is the Go binding that unpacks the parameters returned
// from invoking the contract method with ID 0xaef828de.
//
// Solidity: function teePaymentsRegistry() view returns(address)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackTeePaymentsRegistry(data []byte) (common.Address, error) {
	out, err := teePaymentsConfigVerifier.abi.Unpack("teePaymentsRegistry", data)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackUpdateContractAddresses(contractNameHashes [][32]byte, contractAddresses []common.Address) ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("updateContractAddresses", contractNameHashes, contractAddresses)
}

// PackUpgradeToAndCall is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x4f1ef286.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function upgradeToAndCall(address _newImplementation, bytes _data) payable returns()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackUpgradeToAndCall(newImplementation common.Address, data []byte) []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("upgradeToAndCall", newImplementation, data)
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackUpgradeToAndCall(newImplementation common.Address, data []byte) ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("upgradeToAndCall", newImplementation, data)
}

// PackVerifyAccountConfiguredProof is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8c876193.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function verifyAccountConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackVerifyAccountConfiguredProof(walletId [32]byte, proof IPMWMultisigAccountConfiguredProof) []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("verifyAccountConfiguredProof", walletId, proof)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVerifyAccountConfiguredProof is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x8c876193.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function verifyAccountConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(string,bytes[],uint64),(uint8,uint64)) _proof) returns()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackVerifyAccountConfiguredProof(walletId [32]byte, proof IPMWMultisigAccountConfiguredProof) ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("verifyAccountConfiguredProof", walletId, proof)
}

// PackVerifyUtxoConfiguredProof is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46969efe.  This method will panic if any
// invalid/nil inputs are passed.
//
// Solidity: function verifyUtxoConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof) returns()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) PackVerifyUtxoConfiguredProof(walletId [32]byte, proof IPMWMultisigUtxoConfiguredProof) []byte {
	enc, err := teePaymentsConfigVerifier.abi.Pack("verifyUtxoConfiguredProof", walletId, proof)
	if err != nil {
		panic(err)
	}
	return enc
}

// TryPackVerifyUtxoConfiguredProof is the Go binding used to pack the parameters required for calling
// the contract method with ID 0x46969efe.  This method will return an error
// if any inputs are invalid/nil.
//
// Solidity: function verifyUtxoConfiguredProof(bytes32 _walletId, ((bytes,(uint8,bytes32,bytes32)[],(uint8,bytes32,bytes32)[]),(bytes32,bytes32,uint16,address,address[],uint64,uint64),(uint32,bytes[],uint64,(bytes32,uint32)[]),(uint8,string)) _proof) returns()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) TryPackVerifyUtxoConfiguredProof(walletId [32]byte, proof IPMWMultisigUtxoConfiguredProof) ([]byte, error) {
	return teePaymentsConfigVerifier.abi.Pack("verifyUtxoConfiguredProof", walletId, proof)
}

// TeePaymentsConfigVerifierGovernanceCallTimelocked represents a GovernanceCallTimelocked event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierGovernanceCallTimelocked struct {
	EncodedCall           []byte
	EncodedCallHash       [32]byte
	AllowedAfterTimestamp *big.Int
	Raw                   *types.Log // Blockchain specific contextual infos
}

const TeePaymentsConfigVerifierGovernanceCallTimelockedEventName = "GovernanceCallTimelocked"

// ContractEventName returns the user-defined event name.
func (TeePaymentsConfigVerifierGovernanceCallTimelocked) ContractEventName() string {
	return TeePaymentsConfigVerifierGovernanceCallTimelockedEventName
}

// UnpackGovernanceCallTimelockedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceCallTimelocked(bytes encodedCall, bytes32 encodedCallHash, uint256 allowedAfterTimestamp)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackGovernanceCallTimelockedEvent(log *types.Log) (*TeePaymentsConfigVerifierGovernanceCallTimelocked, error) {
	event := "GovernanceCallTimelocked"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsConfigVerifier.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsConfigVerifierGovernanceCallTimelocked)
	if len(log.Data) > 0 {
		if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsConfigVerifier.abi.Events[event].Inputs {
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

// TeePaymentsConfigVerifierGovernanceInitialised represents a GovernanceInitialised event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierGovernanceInitialised struct {
	InitialGovernance common.Address
	Raw               *types.Log // Blockchain specific contextual infos
}

const TeePaymentsConfigVerifierGovernanceInitialisedEventName = "GovernanceInitialised"

// ContractEventName returns the user-defined event name.
func (TeePaymentsConfigVerifierGovernanceInitialised) ContractEventName() string {
	return TeePaymentsConfigVerifierGovernanceInitialisedEventName
}

// UnpackGovernanceInitialisedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernanceInitialised(address initialGovernance)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackGovernanceInitialisedEvent(log *types.Log) (*TeePaymentsConfigVerifierGovernanceInitialised, error) {
	event := "GovernanceInitialised"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsConfigVerifier.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsConfigVerifierGovernanceInitialised)
	if len(log.Data) > 0 {
		if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsConfigVerifier.abi.Events[event].Inputs {
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

// TeePaymentsConfigVerifierGovernedProductionModeEntered represents a GovernedProductionModeEntered event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierGovernedProductionModeEntered struct {
	GovernanceSettings common.Address
	Raw                *types.Log // Blockchain specific contextual infos
}

const TeePaymentsConfigVerifierGovernedProductionModeEnteredEventName = "GovernedProductionModeEntered"

// ContractEventName returns the user-defined event name.
func (TeePaymentsConfigVerifierGovernedProductionModeEntered) ContractEventName() string {
	return TeePaymentsConfigVerifierGovernedProductionModeEnteredEventName
}

// UnpackGovernedProductionModeEnteredEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event GovernedProductionModeEntered(address governanceSettings)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackGovernedProductionModeEnteredEvent(log *types.Log) (*TeePaymentsConfigVerifierGovernedProductionModeEntered, error) {
	event := "GovernedProductionModeEntered"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsConfigVerifier.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsConfigVerifierGovernedProductionModeEntered)
	if len(log.Data) > 0 {
		if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsConfigVerifier.abi.Events[event].Inputs {
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

// TeePaymentsConfigVerifierInitialized represents a Initialized event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierInitialized struct {
	Version uint64
	Raw     *types.Log // Blockchain specific contextual infos
}

const TeePaymentsConfigVerifierInitializedEventName = "Initialized"

// ContractEventName returns the user-defined event name.
func (TeePaymentsConfigVerifierInitialized) ContractEventName() string {
	return TeePaymentsConfigVerifierInitializedEventName
}

// UnpackInitializedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Initialized(uint64 version)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackInitializedEvent(log *types.Log) (*TeePaymentsConfigVerifierInitialized, error) {
	event := "Initialized"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsConfigVerifier.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsConfigVerifierInitialized)
	if len(log.Data) > 0 {
		if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsConfigVerifier.abi.Events[event].Inputs {
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

// TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled represents a TimelockedGovernanceCallCanceled event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeePaymentsConfigVerifierTimelockedGovernanceCallCanceledEventName = "TimelockedGovernanceCallCanceled"

// ContractEventName returns the user-defined event name.
func (TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled) ContractEventName() string {
	return TeePaymentsConfigVerifierTimelockedGovernanceCallCanceledEventName
}

// UnpackTimelockedGovernanceCallCanceledEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallCanceled(bytes32 encodedCallHash)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackTimelockedGovernanceCallCanceledEvent(log *types.Log) (*TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled, error) {
	event := "TimelockedGovernanceCallCanceled"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsConfigVerifier.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsConfigVerifierTimelockedGovernanceCallCanceled)
	if len(log.Data) > 0 {
		if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsConfigVerifier.abi.Events[event].Inputs {
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

// TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted represents a TimelockedGovernanceCallExecuted event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted struct {
	EncodedCallHash [32]byte
	Raw             *types.Log // Blockchain specific contextual infos
}

const TeePaymentsConfigVerifierTimelockedGovernanceCallExecutedEventName = "TimelockedGovernanceCallExecuted"

// ContractEventName returns the user-defined event name.
func (TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted) ContractEventName() string {
	return TeePaymentsConfigVerifierTimelockedGovernanceCallExecutedEventName
}

// UnpackTimelockedGovernanceCallExecutedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event TimelockedGovernanceCallExecuted(bytes32 encodedCallHash)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackTimelockedGovernanceCallExecutedEvent(log *types.Log) (*TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted, error) {
	event := "TimelockedGovernanceCallExecuted"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsConfigVerifier.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsConfigVerifierTimelockedGovernanceCallExecuted)
	if len(log.Data) > 0 {
		if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsConfigVerifier.abi.Events[event].Inputs {
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

// TeePaymentsConfigVerifierUpgraded represents a Upgraded event raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierUpgraded struct {
	Implementation common.Address
	Raw            *types.Log // Blockchain specific contextual infos
}

const TeePaymentsConfigVerifierUpgradedEventName = "Upgraded"

// ContractEventName returns the user-defined event name.
func (TeePaymentsConfigVerifierUpgraded) ContractEventName() string {
	return TeePaymentsConfigVerifierUpgradedEventName
}

// UnpackUpgradedEvent is the Go binding that unpacks the event data emitted
// by contract.
//
// Solidity: event Upgraded(address indexed implementation)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackUpgradedEvent(log *types.Log) (*TeePaymentsConfigVerifierUpgraded, error) {
	event := "Upgraded"
	if len(log.Topics) == 0 || log.Topics[0] != teePaymentsConfigVerifier.abi.Events[event].ID {
		return nil, errors.New("event signature mismatch")
	}
	out := new(TeePaymentsConfigVerifierUpgraded)
	if len(log.Data) > 0 {
		if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, event, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range teePaymentsConfigVerifier.abi.Events[event].Inputs {
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
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackError(raw []byte) (any, error) {
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["AccountAddressZero"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackAccountAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["AddressEmptyCode"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackAddressEmptyCodeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["AlreadyInProductionMode"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackAlreadyInProductionModeError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["AnchorLimitExceeded"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackAnchorLimitExceededError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["AnchorSetEmpty"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackAnchorSetEmptyError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["CosignersThresholdNotMet"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackCosignersThresholdNotMetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["ERC1967InvalidImplementation"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackERC1967InvalidImplementationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["ERC1967NonPayable"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackERC1967NonPayableError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["FailedCall"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackFailedCallError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["GovernedAddressZero"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackGovernedAddressZeroError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["GovernedAlreadyInitialized"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackGovernedAlreadyInitializedError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["InvalidAttestation"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackInvalidAttestationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["InvalidCosigner"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackInvalidCosignerError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["InvalidInitialization"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackInvalidInitializationError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["InvalidProof"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackInvalidProofError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["InvalidRequestBody"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackInvalidRequestBodyError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["InvalidSigningPolicy"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackInvalidSigningPolicyError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["NotInitializing"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackNotInitializingError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["OnlyExecutor"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackOnlyExecutorError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["OnlyGovernance"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackOnlyGovernanceError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["OnlyProductionOrPausedStatus"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackOnlyProductionOrPausedStatusError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["OnlySystemExtensionId"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackOnlySystemExtensionIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["TimelockCallNotFound"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackTimelockCallNotFoundError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["TimelockNotAllowedYet"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackTimelockNotAllowedYetError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["UUPSUnauthorizedCallContext"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackUUPSUnauthorizedCallContextError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["UUPSUnsupportedProxiableUUID"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackUUPSUnsupportedProxiableUUIDError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["UnsupportedSourceId"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackUnsupportedSourceIdError(raw[4:])
	}
	if bytes.Equal(raw[:4], teePaymentsConfigVerifier.abi.Errors["WrongKeyType"].ID.Bytes()[:4]) {
		return teePaymentsConfigVerifier.UnpackWrongKeyTypeError(raw[4:])
	}
	return nil, errors.New("Unknown error")
}

// TeePaymentsConfigVerifierAccountAddressZero represents a AccountAddressZero error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierAccountAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AccountAddressZero()
func TeePaymentsConfigVerifierAccountAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x7d01fca993146ff8e429fe0a909628970d27f0058569d84d8ec792b0847e5bd7")
}

// UnpackAccountAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AccountAddressZero()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackAccountAddressZeroError(raw []byte) (*TeePaymentsConfigVerifierAccountAddressZero, error) {
	out := new(TeePaymentsConfigVerifierAccountAddressZero)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "AccountAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierAddressEmptyCode represents a AddressEmptyCode error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierAddressEmptyCode struct {
	Target common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AddressEmptyCode(address target)
func TeePaymentsConfigVerifierAddressEmptyCodeErrorID() common.Hash {
	return common.HexToHash("0x9996b315c842ff135b8fc4a08ad5df1c344efbc03d2687aecc0678050d2aac89")
}

// UnpackAddressEmptyCodeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AddressEmptyCode(address target)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackAddressEmptyCodeError(raw []byte) (*TeePaymentsConfigVerifierAddressEmptyCode, error) {
	out := new(TeePaymentsConfigVerifierAddressEmptyCode)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "AddressEmptyCode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierAlreadyInProductionMode represents a AlreadyInProductionMode error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierAlreadyInProductionMode struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AlreadyInProductionMode()
func TeePaymentsConfigVerifierAlreadyInProductionModeErrorID() common.Hash {
	return common.HexToHash("0x72a23ec2b1fdbc3de7d5f411b572fcf8ad8409803196930f63ddf40a2a2b2594")
}

// UnpackAlreadyInProductionModeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AlreadyInProductionMode()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackAlreadyInProductionModeError(raw []byte) (*TeePaymentsConfigVerifierAlreadyInProductionMode, error) {
	out := new(TeePaymentsConfigVerifierAlreadyInProductionMode)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "AlreadyInProductionMode", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierAnchorLimitExceeded represents a AnchorLimitExceeded error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierAnchorLimitExceeded struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AnchorLimitExceeded()
func TeePaymentsConfigVerifierAnchorLimitExceededErrorID() common.Hash {
	return common.HexToHash("0x53d9693388bb3de1aa9c38a0e4b38f8bd50c274a2615f12bf28ed76d2566a5ee")
}

// UnpackAnchorLimitExceededError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AnchorLimitExceeded()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackAnchorLimitExceededError(raw []byte) (*TeePaymentsConfigVerifierAnchorLimitExceeded, error) {
	out := new(TeePaymentsConfigVerifierAnchorLimitExceeded)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "AnchorLimitExceeded", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierAnchorSetEmpty represents a AnchorSetEmpty error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierAnchorSetEmpty struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error AnchorSetEmpty()
func TeePaymentsConfigVerifierAnchorSetEmptyErrorID() common.Hash {
	return common.HexToHash("0xb06cd5887ea5632bef5615859f9d82914081bba13f3044570559e85b2214c80e")
}

// UnpackAnchorSetEmptyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error AnchorSetEmpty()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackAnchorSetEmptyError(raw []byte) (*TeePaymentsConfigVerifierAnchorSetEmpty, error) {
	out := new(TeePaymentsConfigVerifierAnchorSetEmpty)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "AnchorSetEmpty", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierCosignersThresholdNotMet represents a CosignersThresholdNotMet error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierCosignersThresholdNotMet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error CosignersThresholdNotMet()
func TeePaymentsConfigVerifierCosignersThresholdNotMetErrorID() common.Hash {
	return common.HexToHash("0xff41656670fff33a7dbbd13446e34a04caed47f2d1c8607b1828a6aa9d9051b3")
}

// UnpackCosignersThresholdNotMetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error CosignersThresholdNotMet()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackCosignersThresholdNotMetError(raw []byte) (*TeePaymentsConfigVerifierCosignersThresholdNotMet, error) {
	out := new(TeePaymentsConfigVerifierCosignersThresholdNotMet)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "CosignersThresholdNotMet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierERC1967InvalidImplementation represents a ERC1967InvalidImplementation error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierERC1967InvalidImplementation struct {
	Implementation common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func TeePaymentsConfigVerifierERC1967InvalidImplementationErrorID() common.Hash {
	return common.HexToHash("0x4c9c8ce3ceb3130f17f7cdba48d89b5b0129f266a8bac114e6e315a41879b617")
}

// UnpackERC1967InvalidImplementationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967InvalidImplementation(address implementation)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackERC1967InvalidImplementationError(raw []byte) (*TeePaymentsConfigVerifierERC1967InvalidImplementation, error) {
	out := new(TeePaymentsConfigVerifierERC1967InvalidImplementation)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "ERC1967InvalidImplementation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierERC1967NonPayable represents a ERC1967NonPayable error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierERC1967NonPayable struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error ERC1967NonPayable()
func TeePaymentsConfigVerifierERC1967NonPayableErrorID() common.Hash {
	return common.HexToHash("0xb398979fa84f543c8e222f17890372c487baf85e062276c127fef521eea7224b")
}

// UnpackERC1967NonPayableError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error ERC1967NonPayable()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackERC1967NonPayableError(raw []byte) (*TeePaymentsConfigVerifierERC1967NonPayable, error) {
	out := new(TeePaymentsConfigVerifierERC1967NonPayable)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "ERC1967NonPayable", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierFailedCall represents a FailedCall error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierFailedCall struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error FailedCall()
func TeePaymentsConfigVerifierFailedCallErrorID() common.Hash {
	return common.HexToHash("0xd6bda27508c0fb6d8a39b4b122878dab26f731a7d4e4abe711dd3731899052a4")
}

// UnpackFailedCallError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error FailedCall()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackFailedCallError(raw []byte) (*TeePaymentsConfigVerifierFailedCall, error) {
	out := new(TeePaymentsConfigVerifierFailedCall)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "FailedCall", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierGovernedAddressZero represents a GovernedAddressZero error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierGovernedAddressZero struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAddressZero()
func TeePaymentsConfigVerifierGovernedAddressZeroErrorID() common.Hash {
	return common.HexToHash("0x695725f6b429bac0cc5a2fa88214e31520e0c69f2fb0e2fb881a7130a564f49f")
}

// UnpackGovernedAddressZeroError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAddressZero()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackGovernedAddressZeroError(raw []byte) (*TeePaymentsConfigVerifierGovernedAddressZero, error) {
	out := new(TeePaymentsConfigVerifierGovernedAddressZero)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "GovernedAddressZero", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierGovernedAlreadyInitialized represents a GovernedAlreadyInitialized error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierGovernedAlreadyInitialized struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error GovernedAlreadyInitialized()
func TeePaymentsConfigVerifierGovernedAlreadyInitializedErrorID() common.Hash {
	return common.HexToHash("0xb2f0162156d83828197776d56bbfba0a7a9d5936fad70f0a3014ae18cc7aa33c")
}

// UnpackGovernedAlreadyInitializedError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error GovernedAlreadyInitialized()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackGovernedAlreadyInitializedError(raw []byte) (*TeePaymentsConfigVerifierGovernedAlreadyInitialized, error) {
	out := new(TeePaymentsConfigVerifierGovernedAlreadyInitialized)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "GovernedAlreadyInitialized", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierInvalidAttestation represents a InvalidAttestation error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierInvalidAttestation struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidAttestation()
func TeePaymentsConfigVerifierInvalidAttestationErrorID() common.Hash {
	return common.HexToHash("0xbd8ba84d4af301e1b110d1f93f8971934a3ba3afdf4200e7f884f76f2dd27077")
}

// UnpackInvalidAttestationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidAttestation()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackInvalidAttestationError(raw []byte) (*TeePaymentsConfigVerifierInvalidAttestation, error) {
	out := new(TeePaymentsConfigVerifierInvalidAttestation)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "InvalidAttestation", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierInvalidCosigner represents a InvalidCosigner error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierInvalidCosigner struct {
	Cosigner common.Address
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidCosigner(address cosigner)
func TeePaymentsConfigVerifierInvalidCosignerErrorID() common.Hash {
	return common.HexToHash("0xe1e5f014cf4519a719ed6c051af3c17c84765602be52368f3bd57e413daeffcc")
}

// UnpackInvalidCosignerError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidCosigner(address cosigner)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackInvalidCosignerError(raw []byte) (*TeePaymentsConfigVerifierInvalidCosigner, error) {
	out := new(TeePaymentsConfigVerifierInvalidCosigner)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "InvalidCosigner", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierInvalidInitialization represents a InvalidInitialization error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierInvalidInitialization struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidInitialization()
func TeePaymentsConfigVerifierInvalidInitializationErrorID() common.Hash {
	return common.HexToHash("0xf92ee8a957075833165f68c320933b1a1294aafc84ee6e0dd3fb178008f9aaf5")
}

// UnpackInvalidInitializationError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidInitialization()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackInvalidInitializationError(raw []byte) (*TeePaymentsConfigVerifierInvalidInitialization, error) {
	out := new(TeePaymentsConfigVerifierInvalidInitialization)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "InvalidInitialization", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierInvalidProof represents a InvalidProof error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierInvalidProof struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidProof()
func TeePaymentsConfigVerifierInvalidProofErrorID() common.Hash {
	return common.HexToHash("0x09bde339c6b182be216ee7ef8ccff6338c6ef7993445216112ae575c5438fd27")
}

// UnpackInvalidProofError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidProof()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackInvalidProofError(raw []byte) (*TeePaymentsConfigVerifierInvalidProof, error) {
	out := new(TeePaymentsConfigVerifierInvalidProof)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "InvalidProof", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierInvalidRequestBody represents a InvalidRequestBody error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierInvalidRequestBody struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidRequestBody()
func TeePaymentsConfigVerifierInvalidRequestBodyErrorID() common.Hash {
	return common.HexToHash("0xfe524bcc2b70aad90ebe31fd595732afd3bed1387a6e1141569409a40554173c")
}

// UnpackInvalidRequestBodyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidRequestBody()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackInvalidRequestBodyError(raw []byte) (*TeePaymentsConfigVerifierInvalidRequestBody, error) {
	out := new(TeePaymentsConfigVerifierInvalidRequestBody)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "InvalidRequestBody", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierInvalidSigningPolicy represents a InvalidSigningPolicy error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierInvalidSigningPolicy struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error InvalidSigningPolicy()
func TeePaymentsConfigVerifierInvalidSigningPolicyErrorID() common.Hash {
	return common.HexToHash("0x3d7cc620812b32543e78b18734710415246f44f2e8d99c9198a7cb7573499dda")
}

// UnpackInvalidSigningPolicyError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error InvalidSigningPolicy()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackInvalidSigningPolicyError(raw []byte) (*TeePaymentsConfigVerifierInvalidSigningPolicy, error) {
	out := new(TeePaymentsConfigVerifierInvalidSigningPolicy)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "InvalidSigningPolicy", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierNotInitializing represents a NotInitializing error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierNotInitializing struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error NotInitializing()
func TeePaymentsConfigVerifierNotInitializingErrorID() common.Hash {
	return common.HexToHash("0xd7e6bcf8597daa127dc9f0048d2f08d5ef140a2cb659feabd700beff1f7a8302")
}

// UnpackNotInitializingError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error NotInitializing()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackNotInitializingError(raw []byte) (*TeePaymentsConfigVerifierNotInitializing, error) {
	out := new(TeePaymentsConfigVerifierNotInitializing)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "NotInitializing", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierOnlyExecutor represents a OnlyExecutor error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierOnlyExecutor struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyExecutor()
func TeePaymentsConfigVerifierOnlyExecutorErrorID() common.Hash {
	return common.HexToHash("0x7fb6be02aaf40e420b0e7eb67e43a3feee949e8932a3c64e33b591a436670841")
}

// UnpackOnlyExecutorError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyExecutor()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackOnlyExecutorError(raw []byte) (*TeePaymentsConfigVerifierOnlyExecutor, error) {
	out := new(TeePaymentsConfigVerifierOnlyExecutor)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "OnlyExecutor", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierOnlyGovernance represents a OnlyGovernance error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierOnlyGovernance struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyGovernance()
func TeePaymentsConfigVerifierOnlyGovernanceErrorID() common.Hash {
	return common.HexToHash("0x54348f03bacab21a8a36da53ca12d0f27a67eb02beb85de54c1375a35f4cce45")
}

// UnpackOnlyGovernanceError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyGovernance()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackOnlyGovernanceError(raw []byte) (*TeePaymentsConfigVerifierOnlyGovernance, error) {
	out := new(TeePaymentsConfigVerifierOnlyGovernance)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "OnlyGovernance", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierOnlyProductionOrPausedStatus represents a OnlyProductionOrPausedStatus error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierOnlyProductionOrPausedStatus struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlyProductionOrPausedStatus()
func TeePaymentsConfigVerifierOnlyProductionOrPausedStatusErrorID() common.Hash {
	return common.HexToHash("0x7e84272cf09561f32201945a4eb3a196392a46c2d5961caad2f686f27039fd59")
}

// UnpackOnlyProductionOrPausedStatusError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlyProductionOrPausedStatus()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackOnlyProductionOrPausedStatusError(raw []byte) (*TeePaymentsConfigVerifierOnlyProductionOrPausedStatus, error) {
	out := new(TeePaymentsConfigVerifierOnlyProductionOrPausedStatus)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "OnlyProductionOrPausedStatus", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierOnlySystemExtensionId represents a OnlySystemExtensionId error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierOnlySystemExtensionId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error OnlySystemExtensionId()
func TeePaymentsConfigVerifierOnlySystemExtensionIdErrorID() common.Hash {
	return common.HexToHash("0xbd49094d5a86ee3f3c4b39274f5f98391dde2cd16618a917e121333ac3d43ee4")
}

// UnpackOnlySystemExtensionIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error OnlySystemExtensionId()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackOnlySystemExtensionIdError(raw []byte) (*TeePaymentsConfigVerifierOnlySystemExtensionId, error) {
	out := new(TeePaymentsConfigVerifierOnlySystemExtensionId)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "OnlySystemExtensionId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierTimelockCallNotFound represents a TimelockCallNotFound error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierTimelockCallNotFound struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockCallNotFound()
func TeePaymentsConfigVerifierTimelockCallNotFoundErrorID() common.Hash {
	return common.HexToHash("0x262ac987dd09120380d393a0c52cdef1ca58949d1107dc26d8da9f7b527eb33d")
}

// UnpackTimelockCallNotFoundError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockCallNotFound()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackTimelockCallNotFoundError(raw []byte) (*TeePaymentsConfigVerifierTimelockCallNotFound, error) {
	out := new(TeePaymentsConfigVerifierTimelockCallNotFound)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "TimelockCallNotFound", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierTimelockNotAllowedYet represents a TimelockNotAllowedYet error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierTimelockNotAllowedYet struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error TimelockNotAllowedYet()
func TeePaymentsConfigVerifierTimelockNotAllowedYetErrorID() common.Hash {
	return common.HexToHash("0x6124e5c2c545bec1d02048c68500eebdabeae0cc69995a629ebccf8a62f69083")
}

// UnpackTimelockNotAllowedYetError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error TimelockNotAllowedYet()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackTimelockNotAllowedYetError(raw []byte) (*TeePaymentsConfigVerifierTimelockNotAllowedYet, error) {
	out := new(TeePaymentsConfigVerifierTimelockNotAllowedYet)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "TimelockNotAllowedYet", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierUUPSUnauthorizedCallContext represents a UUPSUnauthorizedCallContext error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierUUPSUnauthorizedCallContext struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnauthorizedCallContext()
func TeePaymentsConfigVerifierUUPSUnauthorizedCallContextErrorID() common.Hash {
	return common.HexToHash("0xe07c8dba242a06571ac65fe4bbe20522c9fb111cb33599b799ff8039c1ed18f4")
}

// UnpackUUPSUnauthorizedCallContextError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnauthorizedCallContext()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackUUPSUnauthorizedCallContextError(raw []byte) (*TeePaymentsConfigVerifierUUPSUnauthorizedCallContext, error) {
	out := new(TeePaymentsConfigVerifierUUPSUnauthorizedCallContext)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "UUPSUnauthorizedCallContext", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierUUPSUnsupportedProxiableUUID represents a UUPSUnsupportedProxiableUUID error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierUUPSUnsupportedProxiableUUID struct {
	Slot [32]byte
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func TeePaymentsConfigVerifierUUPSUnsupportedProxiableUUIDErrorID() common.Hash {
	return common.HexToHash("0xaa1d49a4c084bfa9aeeee2a0be65267a7f19ba7e1476b114dac513d2c14cb563")
}

// UnpackUUPSUnsupportedProxiableUUIDError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UUPSUnsupportedProxiableUUID(bytes32 slot)
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackUUPSUnsupportedProxiableUUIDError(raw []byte) (*TeePaymentsConfigVerifierUUPSUnsupportedProxiableUUID, error) {
	out := new(TeePaymentsConfigVerifierUUPSUnsupportedProxiableUUID)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "UUPSUnsupportedProxiableUUID", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierUnsupportedSourceId represents a UnsupportedSourceId error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierUnsupportedSourceId struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error UnsupportedSourceId()
func TeePaymentsConfigVerifierUnsupportedSourceIdErrorID() common.Hash {
	return common.HexToHash("0xb9c2749e785a86b27011ac62f2eeace91e285734582a3d21db2e093d45e3124a")
}

// UnpackUnsupportedSourceIdError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error UnsupportedSourceId()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackUnsupportedSourceIdError(raw []byte) (*TeePaymentsConfigVerifierUnsupportedSourceId, error) {
	out := new(TeePaymentsConfigVerifierUnsupportedSourceId)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "UnsupportedSourceId", raw); err != nil {
		return nil, err
	}
	return out, nil
}

// TeePaymentsConfigVerifierWrongKeyType represents a WrongKeyType error raised by the TeePaymentsConfigVerifier contract.
type TeePaymentsConfigVerifierWrongKeyType struct {
}

// ErrorID returns the hash of canonical representation of the error's signature.
//
// Solidity: error WrongKeyType()
func TeePaymentsConfigVerifierWrongKeyTypeErrorID() common.Hash {
	return common.HexToHash("0xfb224bce70b4062685f138c7eed15a72b903495475f7cd7843835692a7c58bbd")
}

// UnpackWrongKeyTypeError is the Go binding used to decode the provided
// error data into the corresponding Go error struct.
//
// Solidity: error WrongKeyType()
func (teePaymentsConfigVerifier *TeePaymentsConfigVerifier) UnpackWrongKeyTypeError(raw []byte) (*TeePaymentsConfigVerifierWrongKeyType, error) {
	out := new(TeePaymentsConfigVerifierWrongKeyType)
	if err := teePaymentsConfigVerifier.abi.UnpackIntoInterface(out, "WrongKeyType", raw); err != nil {
		return nil, err
	}
	return out, nil
}
