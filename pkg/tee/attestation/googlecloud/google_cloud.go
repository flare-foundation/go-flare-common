// Package googlecloud provides validation and parsing of Google Cloud TEE attestation tokens.
package googlecloud

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/convert"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
)

// EATNonce allows the EAT nonce to be serialized both from a string or an array of strings.
type EATNonce []string

// UnmarshalJSON unmarshals an JSON encoded string or an array of strings into an EATNonce value.
func (e *EATNonce) UnmarshalJSON(data []byte) error {
	var err1, err2 error

	var arr []string
	if err1 = json.Unmarshal(data, &arr); err1 == nil {
		*e = arr
		return nil
	}
	var s string
	if err2 = json.Unmarshal(data, &s); err2 == nil {
		*e = []string{s}
		return nil
	}

	return fmt.Errorf("%w and %w", err1, err2)
}

// GoogleTeeClaims represents the claims present in a Google Cloud TEE attestation JWT.
//
// Based on https://cloud.google.com/confidential-computing/confidential-space/docs/reference/token-claims.
type GoogleTeeClaims struct {
	SecBoot               bool     `json:"secboot"`
	OEMID                 int      `json:"oemid"`
	HWModel               string   `json:"hwmodel"` // Hardware Model
	SWName                string   `json:"swname"`
	GoogleServiceAccounts []string `json:"google_service_accounts"`
	DebugStatus           string   `json:"dbgstat"`
	EATNonce              EATNonce `json:"eat_nonce"` // Entity Attestation Token nonce
	SubMods               SubMods  `json:"submods"`
	jwt.RegisteredClaims
}

type SubMods struct {
	ConfidentialSpace ConfidentialSpaceInfo `json:"confidential_space"`
	Container         Container             `json:"container"`
}

type ConfidentialSpaceInfo struct {
	SupportAttributes []string `json:"support_attributes"`
}

type Container struct {
	ImageID string `json:"image_id"`
}

// Platform is the utf-8 encoded hwmodel (hardware module) on which the confidential computing workload is running.
func (c *GoogleTeeClaims) Platform() (common.Hash, error) {
	p, err := convert.StringToCommonHash(c.HWModel)
	if err != nil {
		return common.Hash{}, fmt.Errorf("unparsable HWModel: %w", err)
	}

	return p, nil
}

// CodeHash is the image ID of the workload container.
func (c *GoogleTeeClaims) CodeHash() (common.Hash, error) {
	ch, err := convert.Hex32StringToCommonHash(strings.TrimPrefix(c.SubMods.Container.ImageID, "sha256:"))
	if err != nil {
		return common.Hash{}, fmt.Errorf("unparsable ImageID: %w", err)
	}

	return ch, nil
}

// ParseAndValidatePKIToken validates the PKI token returned from the Google cloud confidential compute is valid.
// Returns a valid jwt.Token and GoogleTeeClaims or returns an error if invalid.
// leafCRL and intermediateCRL are optional pre-fetched CRLs for revocation checking.
func ParseAndValidatePKIToken(attestationToken string, storedRootCertificate *x509.Certificate, leafCRL, intermediateCRL *x509.RevocationList) (*jwt.Token, *GoogleTeeClaims, error) {
	claims := &GoogleTeeClaims{}
	return ParseAndValidatePKITokenClaims(attestationToken, storedRootCertificate, leafCRL, intermediateCRL, claims)
}

func ParseAndValidatePKITokenClaims[T jwt.Claims](attestationToken string, storedRootCertificate *x509.Certificate, leafCRL, intermediateCRL *x509.RevocationList, claims T) (*jwt.Token, T, error) {
	keyFunc := extractAndValidateKey(storedRootCertificate, leafCRL, intermediateCRL)

	verifiedJWT, err := jwt.ParseWithClaims(attestationToken, claims, keyFunc, jwt.WithValidMethods([]string{"RS256"}))
	if err != nil {
		return nil, claims, fmt.Errorf("parsing and verifying: %w", err)
	}
	return verifiedJWT, claims, err
}

// ParsePKITokenUnverified parses a Google Cloud TEE attestation JWT without verifying the signature.
func ParsePKITokenUnverified(attestationToken string) (*jwt.Token, *GoogleTeeClaims, error) {
	claims := &GoogleTeeClaims{}

	return ParsePKITokenUnverifiedClaims(attestationToken, claims)
}

// ParsePKITokenUnverifiedClaims parses a jwt token with custom provided claims.
func ParsePKITokenUnverifiedClaims[T jwt.Claims](jwtToken string, claims T) (*jwt.Token, T, error) {
	token, _, err := jwt.NewParser().ParseUnverified(jwtToken, claims)
	if err != nil {
		return nil, claims, fmt.Errorf("parsing token: %w", err)
	}
	return token, claims, nil
}

// extractAndValidateKey returns a jwt.KeyFunc that extracts and validates the public key used for verification from a JWT token's x5c header chain.
// It verifies the certificate chain against the expected root certificate and returns the leaf certificate's public key.
func extractAndValidateKey(expectedRoot *x509.Certificate, leafCRL, intermediateCRL *x509.RevocationList) jwt.Keyfunc {
	return func(token *jwt.Token) (any, error) {
		x5cs, ok := token.Header["x5c"]
		if !ok {
			return nil, errors.New("x5c headers missing")
		}

		x5cHeaders, ok := x5cs.([]any)
		if !ok {
			return nil, errors.New("x5c headers not slice")
		}

		certificates, err := ExtractCertificatesFromX5CHeader(x5cHeaders)
		if err != nil {
			return nil, fmt.Errorf("extracting certificates from x5c headers: %w", err)
		}

		err = certificates.Verify(expectedRoot, leafCRL, intermediateCRL)
		if err != nil {
			return nil, fmt.Errorf("verifying certificates: %w", err)
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
		return PKICertificates{}, fmt.Errorf("incorrect number of certificates, expected 3, got %v", len(x5cHeaders))
	}

	parsedHeaders := make([]*x509.Certificate, 3)

	for j := range x5cHeaders {
		h, ok := x5cHeaders[j].(string)
		if !ok {
			return PKICertificates{}, fmt.Errorf("certificate at index %d is not a string:% v", j, x5cHeaders[j])
		}

		cert, err := ParseDERCertificate(h)
		if err != nil {
			return PKICertificates{}, fmt.Errorf("parsing certificate at index %d: %w", j, err)
		}

		parsedHeaders[j] = cert
	}

	certificates := PKICertificates{
		Leaf:         parsedHeaders[0],
		Intermediate: parsedHeaders[1],
		Root:         parsedHeaders[2],
	}
	return certificates, nil
}

// ParseDERCertificate decodes the given DER certificate string and parses it into an x509 certificate.
func ParseDERCertificate(certificate string) (*x509.Certificate, error) {
	bytes, err := base64.StdEncoding.DecodeString(certificate)
	if err != nil {
		return nil, fmt.Errorf("decoding certificate: %w", err)
	}

	cert, err := x509.ParseCertificate(bytes)
	if err != nil {
		return nil, fmt.Errorf("parsing certificate: %w", err)
	}

	return cert, nil
}

// ParsePEMCertificate decodes the given PEM certificate string and parses it into an x509 certificate.
func ParsePEMCertificate(certificate string) (*x509.Certificate, error) {
	block, _ := pem.Decode([]byte(certificate))
	if block == nil {
		return nil, errors.New("cannot decode certificate")
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parsing certificate: %w", err)
	}

	return cert, nil
}

// Verify checks the following:
//   - Leaf signature algorithm is SHA256 With RSA,
//   - Leaf public key algorithm is RSA,
//   - Root certificate matches the expected root certificate,
//   - All certificates have valid lifetimes,
//   - Certificate chain is valid from leaf to root,
//   - Leaf and intermediate certificates are not revoked (if CRLs are provided).
func (c *PKICertificates) Verify(expectedRoot *x509.Certificate, leafCRL, intermediateCRL *x509.RevocationList) error {
	// Verify the leaf certificate signature algorithm is an RSA key
	if c.Leaf.SignatureAlgorithm != x509.SHA256WithRSA {
		return errors.New("leaf certificate signature algorithm is not SHA256WithRSA")
	}

	// Verify the leaf certificate public key algorithm is RSA
	if c.Leaf.PublicKeyAlgorithm != x509.RSA {
		return errors.New("leaf certificate public key algorithm is not RSA")
	}

	if !expectedRoot.Equal(c.Root) {
		return errors.New("unexpected root certificate")
	}

	err := c.verifyLifetime()
	if err != nil {
		return fmt.Errorf("lifetime: %w", err)
	}

	err = c.verifyChain()
	if err != nil {
		return err
	}

	err = c.verifyCRL(leafCRL, intermediateCRL)
	if err != nil {
		return err
	}

	return nil
}

// verifyLifetime checks that all certificate lifetimes are valid.
func (c *PKICertificates) verifyLifetime() error {
	if !isCertificateLifetimeValid(c.Leaf) {
		return errors.New("leaf invalid")
	}

	if !isCertificateLifetimeValid(c.Intermediate) {
		return errors.New("intermediate invalid")
	}

	if !isCertificateLifetimeValid(c.Root) {
		return errors.New("root invalid")
	}

	return nil
}

// verifyChain verifies the certificate chain from leaf to root.
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
		return fmt.Errorf("verifying certificate chain: %w", err)
	}

	return nil
}

// verifyCRL checks whether the leaf or intermediate certificates have been revoked
// according to the provided CRLs. If a CRL is nil, a warning is logged and the check is skipped.
func (c *PKICertificates) verifyCRL(leafCRL, intermediateCRL *x509.RevocationList) error {
	if err := checkCRL("leaf", c.Leaf, leafCRL, c.Intermediate); err != nil {
		return err
	}

	if err := checkCRL("intermediate", c.Intermediate, intermediateCRL, c.Root); err != nil {
		return err
	}

	return nil
}

// checkCRL verifies a single certificate against its CRL. The issuer is the certificate
// that should have signed the CRL.
func checkCRL(name string, cert *x509.Certificate, crl *x509.RevocationList, issuer *x509.Certificate) error {
	if crl == nil {
		if len(cert.CRLDistributionPoints) == 0 {
			// No CRL distribution points — expected for some certs (e.g. Google Confidential Space leaf certs).
			return nil
		}
		logger.Warnf("%s certificate CRL was not provided despite CRL distribution points being present, skipping CRL check", name)
		return nil
	}

	now := time.Now()
	if now.Before(crl.ThisUpdate) {
		return fmt.Errorf("%s CRL is not yet valid (thisUpdate: %s)", name, crl.ThisUpdate.Format(time.RFC3339))
	}
	if !crl.NextUpdate.IsZero() && now.After(crl.NextUpdate) {
		return fmt.Errorf("%s CRL is expired (nextUpdate: %s)", name, crl.NextUpdate.Format(time.RFC3339))
	}

	if err := crl.CheckSignatureFrom(issuer); err != nil {
		return fmt.Errorf("verifying %s CRL signature: %w", name, err)
	}

	for _, entry := range crl.RevokedCertificateEntries {
		if entry.SerialNumber.Cmp(cert.SerialNumber) == 0 {
			return fmt.Errorf("%s certificate has been revoked (serial: %s)", name, cert.SerialNumber.String())
		}
	}

	return nil
}

// isCertificateLifetimeValid checks that the current time is within the certificate's validity period.
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
