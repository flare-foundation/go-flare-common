package googlecloud

import (
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/convert"
	"github.com/golang-jwt/jwt/v4"
)

// This code is slightly modified from Google documentation.
type GoogleTeeClaims struct {
	HWModel  string   `json:"hwmodel"`
	SWName   string   `json:"swname"`
	SecBoot  bool     `json:"secboot"`
	EatNonce []string `json:"eat_nonce"`
	SubMods  subMods  `json:"submods"`
	jwt.StandardClaims
}

// Valid satisfies the jwt.Claims interface without additional checks.
func (c GoogleTeeClaims) Valid() error {
	return nil
}

func (c *GoogleTeeClaims) Platform() (common.Hash, error) {
	p, err := convert.StringToCommonHash(c.HWModel)
	if err != nil {
		return common.Hash{}, fmt.Errorf("unparsable HWModel: %w", err)
	}

	return p, nil
}

func (c *GoogleTeeClaims) CodeHash() (common.Hash, error) {
	ch, err := convert.Hex32StringToCommonHash(strings.TrimPrefix(c.SubMods.Container.ImageDigest, "sha256:"))
	if err != nil {
		return common.Hash{}, fmt.Errorf("unparsable ImageDigest: %w", err)
	}

	return ch, nil
}

type subMods struct {
	Container container `json:"container"`
}

type container struct {
	ImageDigest string `json:"image_digest"`
	ImageID     string `json:"image_id"`
}

// This code is an example of how to validate a PKI token. This library is not an official library,
// nor is it endorsed by Google.

// ExtractAndValidatePKIToken validates the PKI token returned from the attestation service is valid.
// Returns a valid jwt.Token or returns an error if invalid.
func ParseAndValidatePKIToken(attestationToken string, storedRootCertificate x509.Certificate) (jwt.Token, error) {
	keyFunc := extractAndValidateKey(storedRootCertificate)

	verifiedJWT, err := jwt.ParseWithClaims(attestationToken, &GoogleTeeClaims{}, keyFunc)
	return *verifiedJWT, err
}

func ParsePKITokenUnverified(attestationToken string) (*jwt.Token, error) {
	token, _, err := jwt.NewParser().ParseUnverified(attestationToken, &GoogleTeeClaims{})
	if err != nil {
		return nil, fmt.Errorf("could not parse token, %w", err)
	}
	return token, nil
}

func extractAndValidateKey(expectedRoot x509.Certificate) func(*jwt.Token) (any, error) {
	return func(token *jwt.Token) (any, error) {
		if alg := token.Header["alg"]; alg != "RS256" {
			return nil, fmt.Errorf("unsupported alg in header: %v, expected %v", alg, "RS256")
		}

		x5cs, ok := token.Header["x5c"]
		if !ok {
			return jwt.Token{}, errors.New("x5c headers missing")
		}

		x5cHeaders, ok := x5cs.([]any)
		if !ok {
			return jwt.Token{}, errors.New("x5c headers not slice")
		}

		certificates, err := ExtractCertificatesFromX5CHeader(x5cHeaders)
		if err != nil {
			return nil, fmt.Errorf("extracting certificates from x5c headers: %v", err)
		}

		err = certificates.Verify(expectedRoot)
		if err != nil {
			return nil, fmt.Errorf("verifying certificates: %v", err)
		}

		return certificates.Leaf.PublicKey, nil
	}
}

// PKICertificates contains the certificates extracted from the x5c header.
type PKICertificates struct {
	Leaf         *x509.Certificate
	Intermediate *x509.Certificate
	Root         *x509.Certificate
}

// ExtractCertificatesFromX5CHeader extracts the certificates from the given x5c header.
func ExtractCertificatesFromX5CHeader(x5cHeaders []any) (PKICertificates, error) {
	if x5cHeaders == nil {
		return PKICertificates{}, errors.New("x5c headers are nil")
	}

	if len(x5cHeaders) != 3 {
		return PKICertificates{}, fmt.Errorf("incorrect number of certificates, expected 3 certificates, got %v", len(x5cHeaders))
	}

	h, ok := x5cHeaders[0].(string)
	if !ok {
		return PKICertificates{}, fmt.Errorf("leaf certificate is not a string:% v", x5cHeaders[0])
	}
	leaf, err := ParseDERCertificate(h)
	if err != nil {
		return PKICertificates{}, fmt.Errorf("cannot parse leaf certificate: %v", err)
	}

	h, ok = x5cHeaders[1].(string)
	if !ok {
		return PKICertificates{}, fmt.Errorf("intermediate certificate is not a string:% v", x5cHeaders[1])
	}
	intermediate, err := ParseDERCertificate(h)
	if err != nil {
		return PKICertificates{}, fmt.Errorf("cannot parse intermediate certificate: %v", err)
	}

	h, ok = x5cHeaders[1].(string)
	if !ok {
		return PKICertificates{}, fmt.Errorf("root certificate is not a string:% v", x5cHeaders[2])
	}
	root, err := ParseDERCertificate(h)
	if err != nil {
		return PKICertificates{}, fmt.Errorf("cannot parse root certificate: %v", err)
	}

	certificates := PKICertificates{
		Leaf:         leaf,
		Intermediate: intermediate,
		Root:         root,
	}
	return certificates, nil
}

// ParseDERCertificate decodes the given DER certificate string and parses it into an x509 certificate.
func ParseDERCertificate(certificate string) (*x509.Certificate, error) {
	bytes, _ := base64.StdEncoding.DecodeString(certificate)

	cert, err := x509.ParseCertificate(bytes)
	if err != nil {
		return nil, fmt.Errorf("cannot parse certificate: %v", err)
	}

	return cert, nil
}

func (c *PKICertificates) Verify(expectedRoot x509.Certificate) error {
	// Verify the leaf certificate signature algorithm is an RSA key
	if c.Leaf.SignatureAlgorithm != x509.SHA256WithRSA {
		return errors.New("leaf certificate signature algorithm is not SHA256WithRSA")
	}

	// Verify the leaf certificate public key algorithm is RSA
	if c.Leaf.PublicKeyAlgorithm != x509.RSA {
		return errors.New("leaf certificate public key algorithm is not RSA")
	}

	if !expectedRoot.Equal(c.Root) {
		return errors.New(("unexpected root certificate"))
	}

	err := c.verifyLifetime()
	if err != nil {
		return fmt.Errorf("certificate lifetime: %v", err)
	}

	err = c.verifyChain()
	if err != nil {
		return err
	}

	return nil
}

// verifyLifetime checks that all certificate lifetimes are valid.
func (c *PKICertificates) verifyLifetime() error {
	if isCertificateLifetimeValid(c.Leaf) {
		return fmt.Errorf("leaf is not valid")
	}

	if isCertificateLifetimeValid(c.Intermediate) {
		return fmt.Errorf("intermediate is not valid")
	}

	if isCertificateLifetimeValid(c.Root) {
		return fmt.Errorf("root is not valid")
	}

	return nil
}

// verifyChain verifies the certificate chain from leaf to root.
// It also checks that all certificate lifetimes are valid.
func (c *PKICertificates) verifyChain() error {
	interPool := x509.NewCertPool()
	interPool.AddCert(c.Intermediate)

	rootPool := x509.NewCertPool()
	rootPool.AddCert(c.Root)

	_, err := c.Leaf.Verify(x509.VerifyOptions{
		Intermediates: interPool,
		Roots:         rootPool,
		KeyUsages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
	})

	if err != nil {
		return fmt.Errorf("failed to verify certificate chain: %v", err)
	}

	return nil
}

func isCertificateLifetimeValid(certificate *x509.Certificate) bool {
	currentTime := time.Now()
	if currentTime.Before(certificate.NotBefore) {
		return false
	}
	if currentTime.After(certificate.NotAfter) {
		return false
	}
	return true
}
