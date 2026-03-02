package googlecloud

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/convert"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
)

const rootCert = `-----BEGIN CERTIFICATE-----
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
uyT5CnJulxSC5CT1
-----END CERTIFICATE-----`

func TestEATNonceMarshaling(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tests := []struct {
			data     any
			expected EATNonce
		}{
			{
				data:     "",
				expected: []string{""},
			},
			{
				data:     "nonzero",
				expected: []string{"nonzero"},
			},
			{
				data:     []any{""},
				expected: []string{""},
			},
			{
				data:     []any{"1", "2"},
				expected: []string{"1", "2"},
			},
		}

		for _, test := range tests {
			jData, err := json.Marshal(test.data)
			require.NoError(t, err)

			var en EATNonce

			err = json.Unmarshal(jData, &en)
			require.NoError(t, err)

			require.Equal(t, test.expected, en)
		}
	})

	t.Run("fail", func(t *testing.T) {
		tests := []any{
			1,
			[]any{0, []any{1}},
		}

		for _, test := range tests {
			jData, err := json.Marshal(test)
			require.NoError(t, err)

			var x EATNonce

			err = json.Unmarshal(jData, &x)
			require.Error(t, err, test)
		}
	})
}

func TestParsePKITokenUnverifiedHappy(t *testing.T) {
	_, err := ParsePEMCertificate(rootCert)

	require.NoError(t, err)

	tokenString, err := os.ReadFile("test_attestation.jwt")
	require.NoError(t, err)

	tokenStruct, claimsStruct, err := ParsePKITokenUnverified(string(tokenString))
	require.NoError(t, err)

	platform, err := claimsStruct.Platform()
	require.NoError(t, err)

	hrPlatform := convert.CommonHashToString(platform)
	require.Equal(t, "GCP_AMD_SEV", hrPlatform)

	ch, err := claimsStruct.CodeHash()
	require.NoError(t, err)

	require.Equal(t, "cfc0496335c6a4eedd408e87b64c08945d9058963b31fd64ac92754d395d291f", ch.Hex()[2:])

	claimsStructT, ok := tokenStruct.Claims.(*GoogleTeeClaims)
	require.True(t, ok)
	require.Equal(t, claimsStruct, claimsStructT)
	require.False(t, tokenStruct.Valid)

	claimsMap := jwt.MapClaims{}
	tokenMap, claims2, err := ParsePKITokenUnverifiedClaims(string(tokenString), claimsMap)
	require.NoError(t, err)

	require.Equal(t, claimsMap, claims2)
	require.False(t, tokenMap.Valid)

	require.Equal(t, claimsStruct.HWModel, claimsMap["hwmodel"])
}

func TestParseAndValidatePKITokenHappy(t *testing.T) {
	tokenClaims := GoogleTeeClaims{
		HWModel: "testmodel",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "issuer",
		},
	}

	now := time.Now()

	signedToken, root := generateTestPKIToken(t, now, now, now, tokenClaims)

	parsedToken, parsedClaims, err := ParseAndValidatePKIToken(signedToken, root, nil, nil)
	require.NoError(t, err)
	require.NotNil(t, parsedToken)
	require.Equal(t, "testmodel", parsedClaims.HWModel)
	require.Equal(t, "issuer", parsedClaims.Issuer)
}

func TestParseAndValidatePKITokenFail(t *testing.T) {
	defClaims := GoogleTeeClaims{
		HWModel: "testmodel",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "issuer",
		},
	}

	now := time.Now()
	past := now.Add(-300 * time.Hour)
	future := now.Add(300 * time.Hour)

	t.Run("random strings", func(t *testing.T) {
		t.Parallel()

		examples := []string{
			"",
			"11",
			"..1",
			"1.1.1",
			"00.00.00",
		}

		for _, example := range examples {
			_, _, err := ParsePKITokenUnverified(example)
			require.Error(t, err)
		}
	})

	t.Run("invalid lifetime certificates", func(t *testing.T) {
		t.Parallel()

		tests := []struct {
			rt, it, lt time.Time
		}{
			{now, now, future},
			{past, now, now},
			{future, now, now},
			{now, past, now},
			{now, future, now},
			{past, past, now},
			{past, future, now},
			{future, past, now},
			{future, future, now},
			{now, now, past},
			{past, past, past},
		}

		for _, test := range tests {
			signedToken, root := generateTestPKIToken(t, test.rt, test.it, test.lt, defClaims)

			_, _, err := ParseAndValidatePKIToken(signedToken, root, nil, nil)
			require.Error(t, err)
		}
	})

	t.Run("invalid lifetime token", func(t *testing.T) {
		t.Parallel()
		claims := []jwt.RegisteredClaims{
			{
				Issuer:    "issuer",
				ExpiresAt: jwt.NewNumericDate(past), // expired token
			},
			{
				Issuer:    "issuer",
				NotBefore: jwt.NewNumericDate(future), // not before future time
			},
		}

		for _, claims := range claims {
			now := time.Now()
			signedToken, root := generateTestPKIToken(t, now, now, now, claims)

			_, _, err := ParseAndValidatePKIToken(signedToken, root, nil, nil)
			require.Error(t, err)
		}
	})

	t.Run("invalid certificates", func(t *testing.T) {
		t.Parallel()

		root, rootKey := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), true, nil, nil)

		invalidRoot := base64.StdEncoding.EncodeToString([]byte("not a real root"))
		invalidInter := base64.StdEncoding.EncodeToString([]byte("not a real inter"))
		invalidLeaf := base64.StdEncoding.EncodeToString([]byte("not a real leaf"))

		signedToken := assembleSignedToken(t, defClaims, []string{invalidLeaf, invalidInter, invalidRoot}, rootKey)
		_, _, err := ParseAndValidatePKIToken(signedToken, root, nil, nil)
		require.Error(t, err)
	})

	t.Run("nil certificates", func(t *testing.T) {
		t.Parallel()

		root, rootKey := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), true, nil, nil)

		signedToken := assembleSignedToken(t, defClaims, nil, rootKey)
		_, _, err := ParseAndValidatePKIToken(signedToken, root, nil, nil)
		require.Error(t, err)
	})
}

func TestInvalidChains(t *testing.T) {
	now := time.Now()

	root, rootKey := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), true, nil, nil)
	leaf, leafKey := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), false, root, rootKey)

	inter0, intermediateKey0 := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), true, root, rootKey)
	leaf0, leafKey0 := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), false, inter0, intermediateKey0)

	inter1, intermediateKey1 := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), true, inter0, intermediateKey0)
	leaf1, leafKey1 := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), false, inter1, intermediateKey1)

	defClaims := GoogleTeeClaims{
		HWModel: "testmodel",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "issuer",
		},
	}

	tests := []struct {
		certs  []*x509.Certificate
		signer crypto.Signer
	}{
		{
			certs:  []*x509.Certificate{},
			signer: leafKey,
		},
		{
			certs:  []*x509.Certificate{leaf, root},
			signer: leafKey,
		},
		{
			certs:  []*x509.Certificate{leaf1, inter1, inter0, root},
			signer: leafKey1,
		},
		// todo do we allow enforce KeyUsage
		// {
		// 	certs:  []*x509.Certificate{inter1, inter0, root},
		// 	signer: intermediateKey1,
		// },
		{
			certs:  []*x509.Certificate{leaf1, inter0, root},
			signer: leafKey1,
		},
		{
			certs:  []*x509.Certificate{leaf1, inter1, root},
			signer: leafKey1,
		},
		{
			certs:  []*x509.Certificate{leaf0, inter1, root},
			signer: leafKey0,
		},
		{
			certs:  []*x509.Certificate{leaf0, inter0, root},
			signer: leafKey1,
		},
		{
			certs:  []*x509.Certificate{leaf1, inter0, inter1},
			signer: leafKey1,
		},
	}

	for _, test := range tests {
		signedToken := assembleSignedToken(t, defClaims, convertToHeaderEntry(test.certs), test.signer)

		_, _, err := ParseAndValidatePKIToken(signedToken, root, nil, nil)
		require.Error(t, err)
	}
}

func TestExtractCertificatesFromX5CHeader(t *testing.T) {
	now := time.Now()

	root, rootKey := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), true, nil, nil)
	leaf, _ := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), false, root, rootKey)

	inter0, intermediateKey0 := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), true, root, rootKey)
	leaf0, _ := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), false, inter0, intermediateKey0)

	inter1, intermediateKey1 := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), true, inter0, intermediateKey0)
	leaf1, _ := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), false, inter1, intermediateKey1)

	t.Run("success", func(t *testing.T) {
		x5c := []any{pemEncode(leaf0), pemEncode(inter0), pemEncode(root)}
		certs, err := ExtractCertificatesFromX5CHeader(x5c)
		require.NoError(t, err)
		require.NotNil(t, certs.Leaf)
		require.NotNil(t, certs.Intermediate)
		require.NotNil(t, certs.Root)

		require.NoError(t, certs.Verify(root, nil, nil))
	})
	t.Run("wrong type", func(t *testing.T) {
		certs, err := ExtractCertificatesFromX5CHeader([]any{123, 456, 789})
		require.Equal(t, PKICertificates{}, certs)
		require.Error(t, err)
	})
	t.Run("wrong number of certs", func(t *testing.T) {
		certs, err := ExtractCertificatesFromX5CHeader([]any{pemEncode(leaf)})
		require.Equal(t, PKICertificates{}, certs)
		require.Error(t, err)

		certs, err = ExtractCertificatesFromX5CHeader([]any{pemEncode(leaf), pemEncode(root)})
		require.Equal(t, PKICertificates{}, certs)
		require.Error(t, err)

		certs, err = ExtractCertificatesFromX5CHeader([]any{pemEncode(leaf1), pemEncode(inter1), pemEncode(inter0), pemEncode(root)})
		require.Equal(t, PKICertificates{}, certs)
		require.Error(t, err)

		certs, err = ExtractCertificatesFromX5CHeader([]any{})
		require.Equal(t, PKICertificates{}, certs)
		require.Error(t, err)
	})
	t.Run("invalid leaf base64", func(t *testing.T) {
		certs, err := ExtractCertificatesFromX5CHeader([]any{"!!!", "!!!", "!!!"})
		require.Equal(t, PKICertificates{}, certs)
		require.Error(t, err)
	})
	t.Run("invalid intermediate base64", func(t *testing.T) {
		certs, err := ExtractCertificatesFromX5CHeader([]any{pemEncode(leaf0), "!!!", "!!!"})
		require.Equal(t, PKICertificates{}, certs)
		require.Error(t, err)
	})
	t.Run("invalid root base64", func(t *testing.T) {
		invalidDER := base64.StdEncoding.EncodeToString([]byte("not a real certificate"))
		certs, err := ExtractCertificatesFromX5CHeader([]any{pemEncode(leaf0), pemEncode(inter0), invalidDER})
		require.Equal(t, PKICertificates{}, certs)
		require.Error(t, err)
	})
	t.Run("nil x5cHeaders", func(t *testing.T) {
		certs, err := ExtractCertificatesFromX5CHeader(nil)
		require.Equal(t, PKICertificates{}, certs)
		require.Error(t, err)
	})
}

func TestVerifyCRL(t *testing.T) {
	now := time.Now()

	root, rootKey := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), true, nil, nil)
	inter, interKey := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), true, root, rootKey)
	leaf, leafKey := generateTestCertificate(t, now.Add(-time.Hour), now.Add(time.Hour), false, inter, interKey)

	claims := GoogleTeeClaims{
		HWModel: "testmodel",
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "issuer",
		},
	}

	signedToken := assembleSignedToken(t, claims, convertToHeaderEntry([]*x509.Certificate{leaf, inter, root}), leafKey)

	t.Run("nil CRLs succeed with warning", func(t *testing.T) {
		_, _, err := ParseAndValidatePKIToken(signedToken, root, nil, nil)
		require.NoError(t, err)
	})

	t.Run("valid CRLs with no revocations", func(t *testing.T) {
		leafCRL := generateTestCRL(t, inter, interKey, nil)
		interCRL := generateTestCRL(t, root, rootKey, nil)

		_, _, err := ParseAndValidatePKIToken(signedToken, root, leafCRL, interCRL)
		require.NoError(t, err)
	})

	t.Run("leaf CRL not yet valid", func(t *testing.T) {
		now := time.Now()
		leafCRL := generateTestCRLWithTimes(t, inter, interKey, nil, now.Add(time.Hour), now.Add(2*time.Hour))

		_, _, err := ParseAndValidatePKIToken(signedToken, root, leafCRL, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "not yet valid")
	})

	t.Run("intermediate CRL expired", func(t *testing.T) {
		now := time.Now()
		interCRL := generateTestCRLWithTimes(t, root, rootKey, nil, now.Add(-2*time.Hour), now.Add(-time.Hour))

		_, _, err := ParseAndValidatePKIToken(signedToken, root, nil, interCRL)
		require.Error(t, err)
		require.Contains(t, err.Error(), "expired")
	})

	t.Run("revoked leaf certificate", func(t *testing.T) {
		leafCRL := generateTestCRL(t, inter, interKey, []*big.Int{leaf.SerialNumber})

		_, _, err := ParseAndValidatePKIToken(signedToken, root, leafCRL, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "revoked")
	})

	t.Run("revoked intermediate certificate", func(t *testing.T) {
		interCRL := generateTestCRL(t, root, rootKey, []*big.Int{inter.SerialNumber})

		_, _, err := ParseAndValidatePKIToken(signedToken, root, nil, interCRL)
		require.Error(t, err)
		require.Contains(t, err.Error(), "revoked")
	})

	t.Run("CRL signed by wrong issuer", func(t *testing.T) {
		// Leaf CRL signed by root instead of intermediate
		wrongLeafCRL := generateTestCRL(t, root, rootKey, nil)

		_, _, err := ParseAndValidatePKIToken(signedToken, root, wrongLeafCRL, nil)
		require.Error(t, err)
		require.Contains(t, err.Error(), "CRL signature verification failed")
	})

	t.Run("intermediate CRL signed by wrong issuer", func(t *testing.T) {
		// Intermediate CRL signed by intermediate instead of root
		wrongInterCRL := generateTestCRL(t, inter, interKey, nil)

		_, _, err := ParseAndValidatePKIToken(signedToken, root, nil, wrongInterCRL)
		require.Error(t, err)
		require.Contains(t, err.Error(), "CRL signature verification failed")
	})

	t.Run("CRL with other serial does not revoke", func(t *testing.T) {
		otherSerial := big.NewInt(999999)
		leafCRL := generateTestCRL(t, inter, interKey, []*big.Int{otherSerial})
		interCRL := generateTestCRL(t, root, rootKey, []*big.Int{otherSerial})

		_, _, err := ParseAndValidatePKIToken(signedToken, root, leafCRL, interCRL)
		require.NoError(t, err)
	})
}

func generateTestCRL(
	t *testing.T,
	issuer *x509.Certificate,
	issuerKey crypto.Signer,
	revokedSerials []*big.Int,
) *x509.RevocationList {
	now := time.Now()
	return generateTestCRLWithTimes(t, issuer, issuerKey, revokedSerials, now.Add(-time.Hour), now.Add(time.Hour))
}

func generateTestCRLWithTimes(
	t *testing.T,
	issuer *x509.Certificate,
	issuerKey crypto.Signer,
	revokedSerials []*big.Int,
	thisUpdate time.Time,
	nextUpdate time.Time,
) *x509.RevocationList {
	t.Helper()

	revokedEntries := make([]x509.RevocationListEntry, 0, len(revokedSerials))
	for _, serial := range revokedSerials {
		revokedEntries = append(revokedEntries, x509.RevocationListEntry{
			SerialNumber:   serial,
			RevocationTime: thisUpdate.Add(-time.Minute),
		})
	}

	template := &x509.RevocationList{
		RevokedCertificateEntries: revokedEntries,
		Number:                    big.NewInt(1),
		ThisUpdate:                thisUpdate,
		NextUpdate:                nextUpdate,
	}

	crlBytes, err := x509.CreateRevocationList(rand.Reader, template, issuer, issuerKey)
	require.NoError(t, err)

	crl, err := x509.ParseRevocationList(crlBytes)
	require.NoError(t, err)

	return crl
}

func generateTestCertificate(
	t *testing.T,
	notBefore, notAfter time.Time,
	isCA bool,
	parent *x509.Certificate,
	parentKey crypto.Signer,
) (*x509.Certificate, *rsa.PrivateKey) {
	t.Helper()
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	serial := big.NewInt(time.Now().UnixNano())

	template := &x509.Certificate{
		SerialNumber:          serial,
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		SignatureAlgorithm:    x509.SHA256WithRSA,
		PublicKeyAlgorithm:    x509.RSA,
		IsCA:                  isCA,
		BasicConstraintsValid: true,
	}

	if isCA {
		template.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageCRLSign
	} else {
		template.KeyUsage = x509.KeyUsageDigitalSignature
	}

	if parent == nil {
		parent = template
		parentKey = priv
	}

	certDER, err := x509.CreateCertificate(rand.Reader, template, parent, &priv.PublicKey, parentKey)
	require.NoError(t, err)
	cert, err := x509.ParseCertificate(certDER)
	require.NoError(t, err)
	return cert, priv
}

func generateTestPKIToken(
	t *testing.T,
	tr, ti, tl time.Time,
	claims jwt.Claims,
) (string, *x509.Certificate) {
	t.Helper()

	root, rootKey := generateTestCertificate(t, tr.Add(-time.Hour), tr.Add(time.Hour), true, nil, nil)
	inter, intermediateKey := generateTestCertificate(t, ti.Add(-time.Hour), ti.Add(time.Hour), true, root, rootKey)
	leaf, leafKey := generateTestCertificate(t, tl.Add(-time.Hour), tl.Add(time.Hour), false, inter, intermediateKey)

	x5c := []*x509.Certificate{leaf, inter, root}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["x5c"] = x5c

	signedToken := assembleSignedToken(t, claims, convertToHeaderEntry(x5c), leafKey)

	return signedToken, root
}

func assembleSignedToken(t *testing.T, claims jwt.Claims, x5c []string, key crypto.Signer) string {
	t.Helper()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	token.Header["x5c"] = x5c

	signedToken, err := token.SignedString(key)
	require.NoError(t, err)

	return signedToken
}

func convertToHeaderEntry(certs []*x509.Certificate) []string {
	cs := make([]string, 0, len(certs))
	for _, cert := range certs {
		cs = append(cs, pemEncode(cert))
	}

	return cs
}

func pemEncode(cert *x509.Certificate) string {
	return base64.StdEncoding.EncodeToString(cert.Raw)
}
