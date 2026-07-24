package safe

import (
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

const testChainID uint64 = 31337

var testSafe = common.HexToAddress("0x5afe5afe5afe5afe5afe5afe5afe5afe5afe5afe")

func testTx(nonce uint64) Tx {
	return Tx{
		To:        common.HexToAddress("0x1a9C4A0f9D76c0b1D91d22E24E573a9b377618aE"),
		Data:      []byte{0xde, 0xad, 0xbe, 0xef},
		Operation: OperationCall,
		Nonce:     new(big.Int).SetUint64(nonce),
	}
}

// signRaw signs the Safe tx hash directly (Safe v ∈ {27, 28}).
func signRaw(t *testing.T, txHash common.Hash, key *ecdsa.PrivateKey) []byte {
	t.Helper()
	sig, err := crypto.Sign(txHash[:], key)
	require.NoError(t, err)
	sig[64] += 27
	return sig
}

// signEthSign signs the EIP-191 prefixed Safe tx hash (Safe v ∈ {31, 32}).
func signEthSign(t *testing.T, txHash common.Hash, key *ecdsa.PrivateKey) []byte {
	t.Helper()
	sig, err := crypto.Sign(accounts.TextHash(txHash[:]), key)
	require.NoError(t, err)
	sig[64] += 31
	return sig
}

func generateOwners(t *testing.T, n int) ([]*ecdsa.PrivateKey, []common.Address) {
	t.Helper()
	keys := make([]*ecdsa.PrivateKey, n)
	addrs := make([]common.Address, n)
	for i := range n {
		key, err := crypto.GenerateKey()
		require.NoError(t, err)
		keys[i] = key
		addrs[i] = crypto.PubkeyToAddress(key.PublicKey)
	}
	return keys, addrs
}

func TestHashChangesWithEveryField(t *testing.T) {
	base := testTx(7)
	baseHash, err := base.Hash(testChainID, testSafe)
	require.NoError(t, err)

	variants := []Tx{}
	v := testTx(8) // different nonce
	variants = append(variants, v)
	v = testTx(7)
	v.Data = []byte{0x01} // different data
	variants = append(variants, v)
	v = testTx(7)
	v.To = common.HexToAddress("0x00000000000000000000000000000000deadbeef") // different target
	variants = append(variants, v)
	v = testTx(7)
	v.Operation = 1 // DELEGATECALL
	variants = append(variants, v)
	v = testTx(7)
	v.Value = big.NewInt(1)
	variants = append(variants, v)

	for i, variant := range variants {
		got, err := variant.Hash(testChainID, testSafe)
		require.NoError(t, err)
		require.NotEqual(t, baseHash, got, "variant %d must change the hash", i)
	}

	// Different domain: chain and safe address.
	got, err := base.Hash(testChainID+1, testSafe)
	require.NoError(t, err)
	require.NotEqual(t, baseHash, got)
	got, err = base.Hash(testChainID, common.HexToAddress("0x00000000000000000000000000000000deadbeef"))
	require.NoError(t, err)
	require.NotEqual(t, baseHash, got)
}

func TestRecoverSignersBothEncodings(t *testing.T) {
	keys, addrs := generateOwners(t, 2)
	tx := testTx(3)
	txHash, err := tx.Hash(testChainID, testSafe)
	require.NoError(t, err)

	blob := append(signRaw(t, txHash, keys[0]), signEthSign(t, txHash, keys[1])...)
	recovered := RecoverSigners(txHash, blob)
	require.ElementsMatch(t, addrs, recovered)
}

func TestRecoverSignersSkipsUnrecoverableEntries(t *testing.T) {
	keys, addrs := generateOwners(t, 1)
	tx := testTx(0)
	txHash, err := tx.Hash(testChainID, testSafe)
	require.NoError(t, err)

	// v == 1 approved-hash entry (owner in r) cannot be recovered offline.
	approvedHash := make([]byte, SignatureLength)
	copy(approvedHash[12:32], addrs[0][:])
	approvedHash[64] = 1

	blob := append(approvedHash, signRaw(t, txHash, keys[0])...)
	recovered := RecoverSigners(txHash, blob)
	require.Equal(t, []common.Address{addrs[0]}, recovered)
}

func TestApprovalVerify(t *testing.T) {
	keys, addrs := generateOwners(t, 3)
	const executedNonce = 5

	tx := testTx(executedNonce)
	txHash, err := tx.Hash(testChainID, testSafe)
	require.NoError(t, err)
	blob := append(signRaw(t, txHash, keys[0]), signRaw(t, txHash, keys[1])...)

	approval := Approval{
		ExecTransaction: ExecTransactionInputs{
			To:         tx.To,
			Data:       tx.Data,
			Operation:  tx.Operation,
			Signatures: blob,
		},
		Nonce: executedNonce,
	}

	// Correct nonce, threshold met.
	require.NoError(t, approval.Verify(testChainID, testSafe, addrs, 2))

	// Wrong nonce → signatures recover to non-signers → rejected, no scan.
	wrong := approval
	wrong.Nonce = executedNonce + 1
	require.ErrorIs(t, wrong.Verify(testChainID, testSafe, addrs, 2), ErrThresholdNotMet)

	// Correct nonce but threshold higher than signatures present.
	require.ErrorIs(t, approval.Verify(testChainID, testSafe, addrs, 3), ErrThresholdNotMet)

	// Non-owner signer set.
	_, strangers := generateOwners(t, 3)
	require.ErrorIs(t, approval.Verify(testChainID, testSafe, strangers, 2), ErrThresholdNotMet)
}

func TestExecTransactionCalldataRoundTrip(t *testing.T) {
	in := ExecTransactionInputs{
		To:             common.HexToAddress("0x1a9C4A0f9D76c0b1D91d22E24E573a9b377618aE"),
		Value:          big.NewInt(0),
		Data:           []byte{0xde, 0xad, 0xbe, 0xef},
		Operation:      OperationCall,
		SafeTxGas:      big.NewInt(100000),
		BaseGas:        big.NewInt(21000),
		GasPrice:       big.NewInt(0),
		GasToken:       common.Address{},
		RefundReceiver: common.Address{},
		Signatures:     []byte{0x01, 0x02, 0x03},
	}

	calldata, err := EncodeExecTransactionCalldata(in)
	require.NoError(t, err)
	require.Equal(t, ExecTransactionSelector[:], calldata[:4])

	out, err := DecodeExecTransactionCalldata(calldata)
	require.NoError(t, err)
	require.Equal(t, in.To, out.To)
	require.Zero(t, in.Value.Cmp(out.Value))
	require.Equal(t, in.Data, out.Data)
	require.Equal(t, in.Operation, out.Operation)
	require.Zero(t, in.SafeTxGas.Cmp(out.SafeTxGas))
	require.Zero(t, in.BaseGas.Cmp(out.BaseGas))
	require.Equal(t, in.Signatures, out.Signatures)
}

func TestDecodeExecTransactionCalldataRejectsGarbage(t *testing.T) {
	_, err := DecodeExecTransactionCalldata(nil)
	require.ErrorIs(t, err, ErrInvalidExecTransactionCalldata)
	_, err = DecodeExecTransactionCalldata([]byte{0x6a, 0x76})
	require.ErrorIs(t, err, ErrInvalidExecTransactionCalldata)
	// Right selector, truncated arguments.
	_, err = DecodeExecTransactionCalldata(append(ExecTransactionSelector[:], 0xde, 0xad))
	require.Error(t, err)
	// Wrong selector.
	_, err = DecodeExecTransactionCalldata([]byte{0x00, 0x01, 0x02, 0x03, 0x04})
	require.ErrorIs(t, err, ErrInvalidExecTransactionCalldata)
}

func TestExecTransactionSelector(t *testing.T) {
	// The manually-defined execTransaction signature must hash to the canonical
	// Safe v1.3.0 selector 0x6a761202 — guards against a typo in the signature.
	require.Equal(t, [4]byte{0x6a, 0x76, 0x12, 0x02}, ExecTransactionSelector)
}
