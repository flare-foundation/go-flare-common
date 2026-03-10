package aggregator

import (
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/signer"
	"github.com/stretchr/testify/require"
)

func buildTrustSetTx() map[string]any {
	tx := make(map[string]any)
	tx["TransactionType"] = "TrustSet"
	tx["Account"] = "rEuLyBCvcw4CFmzv8RepSiAoNgF8tTGJQC"
	tx["Fee"] = "30000"
	tx["Flags"] = 262144
	tx["LimitAmount"] = map[string]any{
		"currency": "USD",
		"issuer":   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
		"value":    "100",
	}
	tx["Sequence"] = 2
	tx["SigningPubKey"] = ""
	return tx
}

var (
	testSigner1 = &signer.Signer{
		Account:       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
		TxnSignature:  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
		SigningPubKey: "02B3EC4E5DD96029A647CFA20DA07FE1F85296505552CCAC114087E66B46BD77DF",
	}
	testSigner2 = &signer.Signer{
		Account:       "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
		TxnSignature:  "30440220680BBD745004E9CFB6B13A137F505FB92298AD309071D16C7B982825188FD1AE022004200B1F7E4A6A84BB0E4FC09E1E3BA2B66EBD32F0E6D121A34BA3B04AD99BC1",
		SigningPubKey: "028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B",
	}
)

func TestAddSignatures(t *testing.T) {
	tx := buildTrustSetTx()

	blob, err := signing.JoinMultisig(tx, []*signer.Signer{testSigner1, testSigner2})
	require.NoError(t, err)

	acc := Account{
		Address:    "rEuLyBCvcw4CFmzv8RepSiAoNgF8tTGJQC",
		SignerList: map[string]bool{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW": true},
		Quorum:     2,
		txs:        make(map[common.Hash]*transaction),
	}

	txResult, dispatch, err := acc.AddSignatures(blob)
	require.NoError(t, err)
	require.False(t, dispatch)

	require.Len(t, txResult.signers, 1, "signatures")
	require.Len(t, acc.txs, 1, "transactions")
}

func TestAddSignaturesFull(t *testing.T) {
	tx := buildTrustSetTx()

	blob, err := signing.JoinMultisig(tx, []*signer.Signer{testSigner1, testSigner2})
	require.NoError(t, err)

	acc := Account{
		Address: "rEuLyBCvcw4CFmzv8RepSiAoNgF8tTGJQC",
		SignerList: map[string]bool{
			"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW": true,
			"rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v": true,
		},
		Quorum: 2,
		txs:    make(map[common.Hash]*transaction),
	}

	txResult, dispatch, err := acc.AddSignatures(blob)
	require.NoError(t, err)
	require.True(t, dispatch)

	require.Len(t, txResult.signers, 2, "signatures")
	require.Len(t, acc.txs, 1, "transactions")

	txBlob, err := acc.Finalize(txResult.id)
	require.NoError(t, err)

	const expectedBlob = "1200142200040000240000000263D5038D7EA4C680000000000000000000000000005553440000000000B5F762798A53D543A014CAF8B297CFF8F2F937E868400000000000753073008114A3780F5CB5A44D366520FC44055E8ED44D9A2270F3E010732102B3EC4E5DD96029A647CFA20DA07FE1F85296505552CCAC114087E66B46BD77DF744730450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E58114204288D2E47F8EF6C99BCC457966320D12409711E1E0107321028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B744630440220680BBD745004E9CFB6B13A137F505FB92298AD309071D16C7B982825188FD1AE022004200B1F7E4A6A84BB0E4FC09E1E3BA2B66EBD32F0E6D121A34BA3B04AD99BC181147908A7F0EDD48EA896C3580A399F0EE78611C8E3E1F1"
	expectedBlobByte, err := hex.DecodeString(expectedBlob)
	require.NoError(t, err)

	require.Equal(t, expectedBlobByte, txBlob)
}

// TestAddSignaturesInvalidSignature verifies that a signer with a tampered signature is rejected.
func TestAddSignaturesInvalidSignature(t *testing.T) {
	tx := buildTrustSetTx()

	// Swap testSigner1's account/pubkey with testSigner2's signature — the
	// signature is valid DER but covers a different message (wrong account ID).
	s1WithWrongSig := &signer.Signer{
		Account:       testSigner1.Account,
		TxnSignature:  testSigner2.TxnSignature,
		SigningPubKey: testSigner1.SigningPubKey,
	}

	blob, err := signing.JoinMultisig(tx, []*signer.Signer{s1WithWrongSig})
	require.NoError(t, err)

	acc := Account{
		Address:    "rEuLyBCvcw4CFmzv8RepSiAoNgF8tTGJQC",
		SignerList: map[string]bool{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW": true},
		Quorum:     1,
		txs:        make(map[common.Hash]*transaction),
	}

	_, _, err = acc.AddSignatures(blob)
	require.Error(t, err)
	require.Contains(t, err.Error(), "no new valid signatures added")
}

// TestAddSignaturesWrongAccount verifies that a blob for a different account is rejected.
func TestAddSignaturesWrongAccount(t *testing.T) {
	tx := buildTrustSetTx()

	blob, err := signing.JoinMultisig(tx, []*signer.Signer{testSigner1, testSigner2})
	require.NoError(t, err)

	// Aggregator expects a different address than what is in the blob.
	acc := Account{
		Address:    "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
		SignerList: map[string]bool{"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW": true},
		Quorum:     1,
		txs:        make(map[common.Hash]*transaction),
	}

	_, _, err = acc.AddSignatures(blob)
	require.Error(t, err)
	require.Contains(t, err.Error(), "expected address")
}

// TestFinalizeWithoutQuorum verifies that Finalize returns an error when quorum has not been reached.
func TestFinalizeWithoutQuorum(t *testing.T) {
	tx := buildTrustSetTx()

	// Only s1 signs, but quorum = 2.
	blob, err := signing.JoinMultisig(tx, []*signer.Signer{testSigner1})
	require.NoError(t, err)

	acc := Account{
		Address: "rEuLyBCvcw4CFmzv8RepSiAoNgF8tTGJQC",
		SignerList: map[string]bool{
			"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW": true,
			"rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v": true,
		},
		Quorum: 2,
		txs:    make(map[common.Hash]*transaction),
	}

	txResult, dispatch, err := acc.AddSignatures(blob)
	require.NoError(t, err)
	require.False(t, dispatch)

	_, err = acc.Finalize(txResult.id)
	require.Error(t, err)
	require.Contains(t, err.Error(), "quorum not yet reached")
}

// TestFinalizeNonexistent verifies that Finalize returns an error for an unknown transaction ID.
func TestFinalizeNonexistent(t *testing.T) {
	acc := Account{
		Address:    "rEuLyBCvcw4CFmzv8RepSiAoNgF8tTGJQC",
		SignerList: map[string]bool{},
		Quorum:     1,
		txs:        make(map[common.Hash]*transaction),
	}

	_, err := acc.Finalize(common.Hash{})
	require.Error(t, err)
}

// TestAddSignaturesDuplicateSigner verifies that submitting the same signer twice does not duplicate it.
func TestAddSignaturesDuplicateSigner(t *testing.T) {
	tx := buildTrustSetTx()

	blob, err := signing.JoinMultisig(tx, []*signer.Signer{testSigner1, testSigner2})
	require.NoError(t, err)

	acc := Account{
		Address: "rEuLyBCvcw4CFmzv8RepSiAoNgF8tTGJQC",
		SignerList: map[string]bool{
			"rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW": true,
			"rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v": true,
		},
		Quorum: 2,
		txs:    make(map[common.Hash]*transaction),
	}

	// First submission: both valid signers are accepted.
	txResult, dispatch, err := acc.AddSignatures(blob)
	require.NoError(t, err)
	require.True(t, dispatch)
	require.Len(t, txResult.signers, 2)

	// Second submission of the same blob: no new valid signatures.
	_, _, err = acc.AddSignatures(blob)
	require.Error(t, err)
	require.Contains(t, err.Error(), "no new valid signatures added")
}
