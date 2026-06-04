package googlecloud

import (
	"crypto/x509"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestVerifyWithPolicyRejectsNilCerts verifies VerifyWithPolicy errors instead
// of panicking when the certificate chain contains a nil certificate.
func TestVerifyWithPolicyRejectsNilCerts(t *testing.T) {
	var c PKICertificates
	err := c.VerifyWithPolicy(&x509.Certificate{}, nil, nil, Policy{})
	require.Error(t, err)
}

func TestVerifyWithPolicyRejectsNilExpectedRoot(t *testing.T) {
	c := PKICertificates{
		Leaf:         &x509.Certificate{},
		Intermediate: &x509.Certificate{},
		Root:         &x509.Certificate{},
	}
	err := c.VerifyWithPolicy(nil, nil, nil, Policy{})
	require.Error(t, err)
}
