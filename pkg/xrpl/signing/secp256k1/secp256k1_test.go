package secp256k1

import (
	"crypto/ecdsa"
	"encoding/hex"

	"fmt"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	btcecdsa "github.com/btcsuite/btcd/btcec/v2/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/xrpl/hash"
	"github.com/stretchr/testify/require"
)

func TestAddress(t *testing.T) {
	priv, err := crypto.HexToECDSA("76C4FB30C5E1C68142495F367F08F7970783897300A8D29044E89044D6E7F872")
	require.NoError(t, err)

	require.Equal(t, "ravbaTwRkNqecy9Zdw8zwrw4uK5awjqhFd", PrvToAddress(priv))
	require.Equal(t, "038940A036EE380369B9FCC5929A0D431ABE789C8A44E5F48F564E2F6EB608B543", PrvToPub(priv))
}

func SignT(message []byte, privKey *ecdsa.PrivateKey) []byte {
	hash := hash.Sha512Half(message)
	priv, _ := btcec.PrivKeyFromBytes(privKey.D.Bytes())

	sig2 := btcecdsa.Sign(priv, hash)

	return sig2.Serialize()
}

func TestSignatureMarshaling(t *testing.T) {
	prv, err := crypto.HexToECDSA("8F1430C6570B7AB19F4FBB0F1F9A3071192EFB133F6216103F8B4D380B55DE75")
	require.NoError(t, err)

	for range 500 {
		for j := range 5 {
			msg := fmt.Appendf(nil, "neki%d", j)

			sig, err := sign(hash.Sha512Half(msg), prv)
			require.NoError(t, err, j)

			sigXRPL, err := SignXRPL(msg, prv)
			require.NoError(t, err, j)

			isValid := sig.Verify(hash.Sha512Half(msg), &prv.PublicKey)
			require.Truef(t, isValid, "invalid signature %v for priv key %v", j, prv.D)

			sigDER := sig.DER()

			sig2DER := SignT(msg, prv)
			require.Equal(t, sig2DER, sigDER, j)

			sigMar, err := MarshalDER(sigDER)
			require.NoError(t, err, j)

			require.Equal(t, sig.r, sigMar.r, j)
			require.Equal(t, sig.s, sigMar.s, j)

			require.Equal(t, sig2DER, sigXRPL, j)

			ok, err := Validate(msg, sigDER, hex.EncodeToString(toBytesCompressed(&prv.PublicKey)))
			require.NoError(t, err, j)

			require.True(t, ok)

			pub, err := sig.Recover(hash.Sha512Half(msg))
			require.NoError(t, err, j)

			require.Equal(t, pub, &prv.PublicKey)
		}

		prv, err = crypto.GenerateKey() // change prv for the next run
		require.NoError(t, err)
	}
}

func TestMarshaling(t *testing.T) {
	sigs := []string{"30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
		"30440220680BBD745004E9CFB6B13A137F505FB92298AD309071D16C7B982825188FD1AE022004200B1F7E4A6A84BB0E4FC09E1E3BA2B66EBD32F0E6D121A34BA3B04AD99BC1"}

	for j, sig := range sigs {
		b, err := hex.DecodeString(sig)
		require.NoError(t, err, j)

		_, err = MarshalDER(b)
		require.NoError(t, err, j)
	}
}

// Vectors ported from rippled src/test/protocol/SecretKey_test.cpp secp256k1TestVectors.
// Go does not implement rippled's Generator-class seed derivation, so only the raw
// 32-byte secret keys are used; pub-key prefix (0x02 even / 0x03 odd) and
// calcAccountID = RIPEMD160(SHA256(pub)) are checked against rippled's expected values.
func TestKeyDerivationVectors(t *testing.T) {
	tests := []struct {
		secKey string
		pubKey string
		addr   string
	}{
		{
			secKey: "1ACAAEDECE405B2A958212629E16F2EB46B153EEE94CDD350FDEFF52795525B7",
			pubKey: "0330E7FC9D56BB25D6893BA3F317AE5BCF33B3291BD63DB32654A313222F7FD020",
			addr:   "rHb9CJAWyB4rj91VRWn96DkukG4bwdtyTh",
		},
		{
			secKey: "E36928EE1260088D4790B399254779F79BDD48A9ECADC13BF4597B69ADE16F3E",
			pubKey: "020965CE5C65510CC6D7AA6498FD0CE1D8E9ADCD720012A51D366494B197F6B99E",
			addr:   "rfZCLUjvKSNmg5xMufb6fgq9VfP5biBDfU",
		},
		{
			secKey: "B4B6AB76337598E2BF430792EF140436E25C43ADA06BED8CA1CC807FEA3BA526",
			pubKey: "0395E97A55CE552EC0E3C000F60553632A3EEAF94734640BABD1543AD0A69005CD",
			addr:   "rpyTz8db86bWHi8E43GGexhRsLDwwMRta3",
		},
	}

	for _, tc := range tests {
		t.Run(tc.addr, func(t *testing.T) {
			priv, err := crypto.HexToECDSA(tc.secKey)
			require.NoError(t, err)

			require.Equal(t, tc.pubKey, PrvToPub(priv))
			require.Equal(t, tc.addr, PrvToAddress(priv))
		})
	}
}

// Canonicality vector from rippled src/test/protocol/SecretKey_test.cpp testCanonicality.
// pubKey 02 5096EB12... verifies digest 34C19028... against the fully-canonical DER signature.
func TestCanonicalityVerify(t *testing.T) {
	pub := "025096EB12D3E924234E7162369C11D8BF877EDA238778E7A31FF0AAC5D0DBCF37"

	sig, err := hex.DecodeString("3045022100B49D07F0E934BA468C0EFC78117791408D1FB8B63A6492AD395AC2F360F2466002205087" +
		"39DB0A2EF81676E39F459C8BBB07A09C3E9F9BEB696294D524D479D62740")
	require.NoError(t, err)

	digest, err := hex.DecodeString("34C19028C80D21F3F48C9354895F8D5BF0D5EE7FF457647CF655F5530A3022A7")
	require.NoError(t, err)

	sigP, err := MarshalDER(sig)
	require.NoError(t, err)

	pubBytes, err := hex.DecodeString(pub)
	require.NoError(t, err)

	require.True(t, sigP.VerifyHex(digest, pubBytes))
}

// rippled SecretKey.cpp sign(KeyType::secp256k1, ...): the signature is over sha512half(message),
// produced by secp256k1_ecdsa_sign with RFC6979 deterministic nonce, then DER-serialized.
// Here we verify Go's deterministic output matches btcec's (which also uses RFC6979).
func TestDeterministicSigningRFC6979(t *testing.T) {
	priv, err := crypto.HexToECDSA("1ACAAEDECE405B2A958212629E16F2EB46B153EEE94CDD350FDEFF52795525B7")
	require.NoError(t, err)

	msg := []byte("http://www.xrpl.org")

	sigGo, err := SignXRPL(msg, priv)
	require.NoError(t, err)

	sigBtc := SignT(msg, priv)
	require.Equal(t, sigBtc, sigGo)

	for range 5 {
		sigAgain, err := SignXRPL(msg, priv)
		require.NoError(t, err)
		require.Equal(t, sigGo, sigAgain)
	}

	ok, err := Validate(msg, sigGo, PrvToPub(priv))
	require.NoError(t, err)
	require.True(t, ok)
}
