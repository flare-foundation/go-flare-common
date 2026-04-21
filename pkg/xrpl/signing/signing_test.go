package signing

import (
	"crypto/ed25519"
	"encoding/hex"
	"maps"
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
	xed25519 "github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/ed25519"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/signing/signer"
	"github.com/stretchr/testify/require"
)

func TestJoinMultisig(t *testing.T) {
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

	s1 := &signer.Signer{
		Account:       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
		TxnSignature:  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
		SigningPubKey: "02B3EC4E5DD96029A647CFA20DA07FE1F85296505552CCAC114087E66B46BD77DF",
	}

	s2 := &signer.Signer{
		Account:       "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
		TxnSignature:  "30440220680BBD745004E9CFB6B13A137F505FB92298AD309071D16C7B982825188FD1AE022004200B1F7E4A6A84BB0E4FC09E1E3BA2B66EBD32F0E6D121A34BA3B04AD99BC1",
		SigningPubKey: "028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B",
	}

	signers := []*signer.Signer{s1, s2}

	blob, err := JoinMultisig(tx, signers)
	require.NoError(t, err)

	const expectedBlob = "1200142200040000240000000263D5038D7EA4C680000000000000000000000000005553440000000000B5F762798A53D543A014CAF8B297CFF8F2F937E868400000000000753073008114A3780F5CB5A44D366520FC44055E8ED44D9A2270F3E010732102B3EC4E5DD96029A647CFA20DA07FE1F85296505552CCAC114087E66B46BD77DF744730450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E58114204288D2E47F8EF6C99BCC457966320D12409711E1E0107321028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B744630440220680BBD745004E9CFB6B13A137F505FB92298AD309071D16C7B982825188FD1AE022004200B1F7E4A6A84BB0E4FC09E1E3BA2B66EBD32F0E6D121A34BA3B04AD99BC181147908A7F0EDD48EA896C3580A399F0EE78611C8E3E1F1"
	expectedBlobByte, err := hex.DecodeString(expectedBlob)
	require.NoError(t, err)

	require.Equal(t, expectedBlobByte, blob)
}

// End-to-end: sign with an Ed25519 key derived from rippled's ed25519TestVectors[0]
// (src/test/protocol/SecretKey_test.cpp: seed AF41..1C97 -> addr rVAEQBhWT6nZ4woEifdN3TMMdUZaxeXnR)
// via SignTxMultisig, then confirm ValidateMultiSig accepts the produced signer.
// Covers: utils.Prepare (multi) -> ed25519.Sign -> ed25519.Validate chain using a
// rippled-provenance key whose pubkey prefix is 0xED and whose accountID path is
// RIPEMD160(SHA256(pub)) per rippled src/libxrpl/protocol/AccountID.cpp calcAccountID.
func TestValidateMultiSigRoundTripED(t *testing.T) {
	seed16, err := hex.DecodeString("AF41FF66F75EBD3A6B18FB7A1DF61C97")
	require.NoError(t, err)

	rawPriv := hash.Sha512Half(seed16)
	priv := ed25519.NewKeyFromSeed(rawPriv)

	tx := map[string]any{
		"TransactionType": "TrustSet",
		"Account":         "rEuLyBCvcw4CFmzv8RepSiAoNgF8tTGJQC",
		"Fee":             "30000",
		"Flags":           262144,
		"LimitAmount": map[string]any{
			"currency": "USD",
			"issuer":   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
			"value":    "100",
		},
		"Sequence":      2,
		"SigningPubKey": "",
	}

	s, err := xed25519.SignTxMultisig(tx, priv)
	require.NoError(t, err)
	require.Equal(t, "rVAEQBhWT6nZ4woEifdN3TMMdUZaxeXnR", s.Account)

	ok, err := ValidateMultiSig(tx, s)
	require.NoError(t, err)
	require.True(t, ok)

	tampered := &signer.Signer{
		Account:       s.Account,
		TxnSignature:  s.TxnSignature,
		SigningPubKey: s.SigningPubKey,
	}
	txTampered := maps.Clone(tx)
	txTampered["Sequence"] = 3
	ok, err = ValidateMultiSig(txTampered, tampered)
	require.NoError(t, err)
	require.False(t, ok)
}
