// Package googlecloud provides validation and parsing of Google Cloud TEE attestation tokens.
package googlecloud

import (
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/convert"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
)

// ConfidentialSpaceIssuer is the iss claim Google Confidential Space sets on attestation tokens.
const ConfidentialSpaceIssuer = "https://confidentialcomputing.googleapis.com"

// productionDebugStatus is the dbgstat value Confidential Space sets when the VM
// has had debugging disabled since boot. Any other value indicates a debug-mode VM
// with reduced isolation guarantees.
const productionDebugStatus = "disabled-since-boot"

const defaultLeeway = 30 * time.Second

// Policy carries the per-deployment configuration the verifier supplies. The library
// performs the comparisons; the consumer only supplies the values it alone knows.
//
// Each optional check runs only when the policy carries the information to perform it:
// an allowlist/value field is checked when non-empty, a Require* boolean when true.
// Every field therefore defaults to "skip", so the zero Policy verifies token
// AUTHENTICITY ONLY — signature, x5c chain to the pinned root, RS256, iss, and exp —
// but NOT workload IDENTITY. Almost every production caller should set at least
// AllowedImageIDs and RequireSecBoot. iss is always checked; supplying Issuer only
// overrides the expected value (intended for test environments).
//
// Fields group by where they are consumed:
//   - Always (JWT/chain level, both entry points): Audience, Issuer, Leeway,
//     RequireCRL, AllowedLeafEKUs.
//   - Only by Apply (the full ParseAndValidatePKIToken path; the generic-claims
//     variant does not run these): AllowedImageIDs, AllowedHWModels, EATNonce,
//     RequireSecBoot, AllowedDebugStatuses.
//
// Do not add a validate() that requires Apply-only fields: they have no effect on the
// ParseAndValidatePKITokenClaims path and requiring them there would reject callers
// that legitimately assert those claims themselves.
type Policy struct {
	// Audience is the expected aud claim. Empty skips the aud check.
	// The TEE workload requests an attestation token with this audience from the
	// Google metadata server; the verifier should require the same value back.
	Audience string

	// Issuer overrides the expected iss claim. Empty defaults to ConfidentialSpaceIssuer.
	// iss is always checked; this field only changes the expected value (intended for tests).
	Issuer string

	// Leeway is the clock skew tolerance for exp/iat/nbf. Zero means defaultLeeway.
	Leeway time.Duration

	// RequireCRL causes verification to fail closed when a certificate
	// declares CRLDistributionPoints but the caller did not supply the
	// corresponding x509.RevocationList. The default false keeps the
	// prior behavior (warn + skip), which is appropriate when CRLs are
	// fetched out-of-band or when the leaf has no DPs (Confidential Space
	// leaves typically don't).
	RequireCRL bool

	// AllowedLeafEKUs is the set of Extended Key Usage values the chain
	// must authorize for the leaf certificate. Nil or empty falls back to
	// the leaf's own declared ExtKeyUsage, which at least prevents the
	// chain validator from accepting any EKU via ExtKeyUsageAny.
	AllowedLeafEKUs []x509.ExtKeyUsage

	// AllowedImageIDs is the allowlist of acceptable workload container image IDs
	// (the sha256 hash from submods.container.image_id). Empty skips the image_id check.
	AllowedImageIDs map[common.Hash]struct{}

	// AllowedHWModels restricts to specific hardware models (hwmodel claim).
	// Empty skips the hwmodel check.
	AllowedHWModels map[string]struct{}

	// EATNonce is the per-request challenge to bind. When non-empty, the token's
	// eat_nonce list must contain this value. Empty skips replay binding.
	EATNonce string

	// RequireSecBoot, when true, rejects tokens reporting secboot=false.
	// Default false skips the check.
	RequireSecBoot bool

	// AllowedDebugStatuses is the allowlist of acceptable dbgstat claim values
	// (e.g. productionDebugStatus). Empty skips the dbgstat check.
	AllowedDebugStatuses []string
}

func (p Policy) issuer() string {
	if p.Issuer != "" {
		return p.Issuer
	}
	return ConfidentialSpaceIssuer
}

func (p Policy) leeway() time.Duration {
	if p.Leeway == 0 {
		return defaultLeeway
	}
	return p.Leeway
}

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
//
// The library enforces the full check chain against policy: x5c chain to root, RS256-only,
// iss/exp claims, leeway, and — when the corresponding policy field is set — aud, secboot,
// dbgstat, hwmodel, image_id allowlist, and eat_nonce. The consumer supplies the
// per-deployment values; comparison is done here. See Policy: unset fields skip their check,
// so an under-specified policy verifies authenticity but not workload identity.
//
// storedRootCertificate is compared against the chain's root via x509.Certificate.Equal,
// so the caller is responsible for sourcing it from a trusted external location (e.g.
// Google's published Confidential Space Root CA) and verifying it out-of-band — typically
// by SHA-256 fingerprint — before passing it in. The library does NOT embed the root,
// to avoid making this repo a supply-chain target for cert substitution.
func ParseAndValidatePKIToken(attestationToken string, storedRootCertificate *x509.Certificate, leafCRL, intermediateCRL *x509.RevocationList, policy Policy) (*jwt.Token, *GoogleTeeClaims, error) {
	claims := &GoogleTeeClaims{}
	parsed, claims, err := parseAndVerifyJWT(attestationToken, storedRootCertificate, leafCRL, intermediateCRL, policy, claims)
	if err != nil {
		return nil, claims, err
	}
	if err := claims.Apply(policy); err != nil {
		return nil, claims, fmt.Errorf("applying policy: %w", err)
	}
	return parsed, claims, nil
}

// ParseAndValidatePKITokenClaims is the generic-claims variant of ParseAndValidatePKIToken.
// JWT-level checks (signature, iss, exp, leeway, and aud when Policy.Audience is set) are
// performed by this function. The Apply-only claim checks (secboot, dbgstat, image_id,
// hwmodel, eat_nonce) are NOT applied — the caller is responsible for asserting them on the
// custom claims type, and the corresponding Policy fields are ignored here.
// Use ParseAndValidatePKIToken when consuming GoogleTeeClaims to get the full check chain.
func ParseAndValidatePKITokenClaims[T jwt.Claims](attestationToken string, storedRootCertificate *x509.Certificate, leafCRL, intermediateCRL *x509.RevocationList, policy Policy, claims T) (*jwt.Token, T, error) {
	return parseAndVerifyJWT(attestationToken, storedRootCertificate, leafCRL, intermediateCRL, policy, claims)
}

func parseAndVerifyJWT[T jwt.Claims](attestationToken string, storedRootCertificate *x509.Certificate, leafCRL, intermediateCRL *x509.RevocationList, policy Policy, claims T) (*jwt.Token, T, error) {
	keyFunc := extractAndValidateKey(storedRootCertificate, leafCRL, intermediateCRL, policy)

	opts := []jwt.ParserOption{
		jwt.WithValidMethods([]string{"RS256"}),
		jwt.WithIssuer(policy.issuer()),
		jwt.WithExpirationRequired(),
		jwt.WithLeeway(policy.leeway()),
	}
	// WithAudience("") would require an empty aud claim and reject every real token,
	// so the aud check is opt-in: add it only when an expected audience is set.
	if policy.Audience != "" {
		opts = append(opts, jwt.WithAudience(policy.Audience))
	}

	verifiedJWT, err := jwt.ParseWithClaims(attestationToken, claims, keyFunc, opts...)
	if err != nil {
		return nil, claims, fmt.Errorf("parsing and verifying: %w", err)
	}
	return verifiedJWT, claims, nil
}

// Apply asserts the Confidential-Space-specific claim values on c against the policy.
// Each check runs only when its policy field is set (RequireSecBoot true, or the relevant
// allowlist/nonce non-empty); an unset field skips its check. Returns nil if all enabled
// checks pass.
//
// Called automatically by ParseAndValidatePKIToken; callers using
// ParseAndValidatePKITokenClaims with a custom claim type are responsible for their own
// equivalent checks.
func (c *GoogleTeeClaims) Apply(p Policy) error {
	if p.RequireSecBoot && !c.SecBoot {
		return errors.New("secboot not enabled")
	}
	if len(p.AllowedDebugStatuses) > 0 && !slices.Contains(p.AllowedDebugStatuses, c.DebugStatus) {
		return fmt.Errorf("dbgstat %q not in allowlist", c.DebugStatus)
	}
	if len(p.AllowedHWModels) > 0 {
		if _, ok := p.AllowedHWModels[c.HWModel]; !ok {
			return fmt.Errorf("hwmodel %q not in allowlist", c.HWModel)
		}
	}
	if len(p.AllowedImageIDs) > 0 {
		ch, err := c.CodeHash()
		if err != nil {
			return fmt.Errorf("decoding image_id: %w", err)
		}
		if _, ok := p.AllowedImageIDs[ch]; !ok {
			return fmt.Errorf("image_id %s not in allowlist", ch.Hex())
		}
	}
	// WARNING: empty p.EATNonce silently skips replay binding.
	if p.EATNonce != "" {
		if !slices.Contains([]string(c.EATNonce), p.EATNonce) {
			return errors.New("eat_nonce does not contain expected challenge")
		}
	}
	return nil
}

// ParsePKITokenUnverified parses a Google Cloud TEE attestation JWT WITHOUT
// verifying the signature, certificate chain, or any claim. An attacker can
// produce a JWT with arbitrary header and claims, and a missing or junk
// signature; this function will return those claims without complaint.
//
// Safe only when the token comes from a trusted local source — typically the
// VM's own attestation token fetched from the Google Confidential Space
// metadata server (metadata.google.internal), where the VM boundary already
// authenticates the source. Reading self-claims (e.g. the VM's own image_id
// for self-reporting) is the intended use.
//
// For any token received from another party or over the wire, use
// ParseAndValidatePKIToken with a Policy instead.
func ParsePKITokenUnverified(attestationToken string) (*jwt.Token, *GoogleTeeClaims, error) {
	claims := &GoogleTeeClaims{}

	return ParsePKITokenUnverifiedClaims(attestationToken, claims)
}

// ParsePKITokenUnverifiedClaims is the generic-claims variant of
// ParsePKITokenUnverified; the same safety contract applies — see that
// function's documentation.
func ParsePKITokenUnverifiedClaims[T jwt.Claims](jwtToken string, claims T) (*jwt.Token, T, error) {
	token, _, err := jwt.NewParser().ParseUnverified(jwtToken, claims)
	if err != nil {
		return nil, claims, fmt.Errorf("parsing token: %w", err)
	}
	return token, claims, nil
}

// extractAndValidateKey returns a jwt.KeyFunc that extracts and validates the public key used for verification from a JWT token's x5c header chain.
// It verifies the certificate chain against the expected root certificate and returns the leaf certificate's public key.
func extractAndValidateKey(expectedRoot *x509.Certificate, leafCRL, intermediateCRL *x509.RevocationList, policy Policy) jwt.Keyfunc {
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

		err = certificates.VerifyWithPolicy(expectedRoot, leafCRL, intermediateCRL, policy)
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
//
// Equivalent to VerifyWithPolicy(..., Policy{}). New code should call
// VerifyWithPolicy and supply the per-deployment hardening knobs.
func (c *PKICertificates) Verify(expectedRoot *x509.Certificate, leafCRL, intermediateCRL *x509.RevocationList) error {
	return c.VerifyWithPolicy(expectedRoot, leafCRL, intermediateCRL, Policy{})
}

// VerifyWithPolicy is the policy-aware variant of Verify. Policy.RequireCRL fails closed
// when a certificate declares DPs but no CRL was supplied; Policy.AllowedLeafEKUs constrains
// chain validation to specific Extended Key Usage values rather than ExtKeyUsageAny.
func (c *PKICertificates) VerifyWithPolicy(expectedRoot *x509.Certificate, leafCRL, intermediateCRL *x509.RevocationList, policy Policy) error {
	if c.Leaf == nil || c.Intermediate == nil || c.Root == nil {
		return errors.New("nil certificate in chain")
	}
	if expectedRoot == nil {
		return errors.New("nil expected root certificate")
	}

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

	err = c.verifyChain(policy)
	if err != nil {
		return err
	}

	err = c.verifyCRL(leafCRL, intermediateCRL, policy.RequireCRL)
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
//
// KeyUsages is pinned to policy.AllowedLeafEKUs when set; otherwise it falls back to the
// leaf's own declared ExtKeyUsage so the chain validator does not accept any EKU via
// ExtKeyUsageAny. If the leaf declares no EKU we keep ExtKeyUsageAny — there is nothing
// stricter to pin to without a deployment-specific Policy.AllowedLeafEKUs value.
func (c *PKICertificates) verifyChain(policy Policy) error {
	interPool := x509.NewCertPool()
	interPool.AddCert(c.Intermediate)

	rootPool := x509.NewCertPool()
	rootPool.AddCert(c.Root)

	keyUsages := policy.AllowedLeafEKUs
	if len(keyUsages) == 0 {
		keyUsages = c.Leaf.ExtKeyUsage
	}
	if len(keyUsages) == 0 {
		keyUsages = []x509.ExtKeyUsage{x509.ExtKeyUsageAny}
	}

	_, err := c.Leaf.Verify(x509.VerifyOptions{
		Intermediates: interPool,
		Roots:         rootPool,
		KeyUsages:     keyUsages,
	})
	if err != nil {
		return fmt.Errorf("verifying certificate chain: %w", err)
	}

	return nil
}

// verifyCRL checks whether the leaf or intermediate certificates have been revoked
// according to the provided CRLs. If a CRL is nil and requireCRL is false,
// a warning is logged and the check is skipped; if requireCRL is true, the
// function fails closed.
func (c *PKICertificates) verifyCRL(leafCRL, intermediateCRL *x509.RevocationList, requireCRL bool) error {
	if err := checkCRL("leaf", c.Leaf, leafCRL, c.Intermediate, requireCRL); err != nil {
		return err
	}

	if err := checkCRL("intermediate", c.Intermediate, intermediateCRL, c.Root, requireCRL); err != nil {
		return err
	}

	return nil
}

// checkCRL verifies a single certificate against its CRL. The issuer is the certificate
// that should have signed the CRL.
func checkCRL(name string, cert *x509.Certificate, crl *x509.RevocationList, issuer *x509.Certificate, requireCRL bool) error {
	if crl == nil {
		if len(cert.CRLDistributionPoints) == 0 {
			// No CRL distribution points — expected for some certs (e.g. Google Confidential Space leaf certs).
			return nil
		}
		if requireCRL {
			return fmt.Errorf("%s certificate declares CRL distribution points but no CRL was provided (Policy.RequireCRL is set)", name)
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
