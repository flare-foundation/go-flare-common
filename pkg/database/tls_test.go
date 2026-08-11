package database

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// writeTestCertPair writes a self-signed certificate and its key to dir and
// returns their paths.
func writeTestCertPair(t *testing.T) (certPath, keyPath string) {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)

	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "test-ca"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}
	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	require.NoError(t, err)

	keyDER, err := x509.MarshalECPrivateKey(key)
	require.NoError(t, err)

	dir := t.TempDir()
	certPath = filepath.Join(dir, "cert.pem")
	keyPath = filepath.Join(dir, "key.pem")
	writePEM(t, certPath, "CERTIFICATE", der)
	writePEM(t, keyPath, "EC PRIVATE KEY", keyDER)
	return certPath, keyPath
}

func writePEM(t *testing.T, path, blockType string, der []byte) {
	t.Helper()
	require.NoError(t, os.WriteFile(path, pem.EncodeToMemory(&pem.Block{Type: blockType, Bytes: der}), 0o600))
}

func TestResolveTLS(t *testing.T) {
	caPath, keyPath := writeTestCertPair(t)

	tests := []struct {
		name          string
		cfg           TLSConfig
		expected      string
		expectedError string
	}{
		{
			name:     "zero value is plaintext",
			cfg:      TLSConfig{},
			expected: "",
		},
		{
			name:     "off is plaintext",
			cfg:      TLSConfig{Mode: "off"},
			expected: "",
		},
		{
			name:          "off with ca_cert",
			cfg:           TLSConfig{Mode: "off", CACert: caPath},
			expectedError: `require mode "verify"`,
		},
		{
			name:          "empty mode with server_name",
			cfg:           TLSConfig{ServerName: "db.example.com"},
			expectedError: `require mode "verify"`,
		},
		{
			name:          "empty mode with client key",
			cfg:           TLSConfig{ClientKey: keyPath},
			expectedError: `require mode "verify"`,
		},
		{
			name:     "skip-verify",
			cfg:      TLSConfig{Mode: "skip-verify"},
			expected: "skip-verify",
		},
		{
			name:          "skip-verify with ca_cert",
			cfg:           TLSConfig{Mode: "skip-verify", CACert: caPath},
			expectedError: `require mode "verify"`,
		},
		{
			name:          "skip-verify with client pair",
			cfg:           TLSConfig{Mode: "skip-verify", ClientCert: caPath, ClientKey: keyPath},
			expectedError: `require mode "verify"`,
		},
		{
			name:     "verify with system roots",
			cfg:      TLSConfig{Mode: "verify"},
			expected: "true",
		},
		{
			name:          "unknown mode",
			cfg:           TLSConfig{Mode: "required"},
			expectedError: `unknown mode "required"`,
		},
		{
			name:          "uppercase mode is rejected",
			cfg:           TLSConfig{Mode: "VERIFY"},
			expectedError: `unknown mode "VERIFY"`,
		},
		{
			name:          "client cert without key",
			cfg:           TLSConfig{Mode: "verify", ClientCert: caPath},
			expectedError: "client_cert and client_key must be set together",
		},
		{
			name:          "client key without cert",
			cfg:           TLSConfig{Mode: "verify", ClientKey: keyPath},
			expectedError: "client_cert and client_key must be set together",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			name, err := resolveTLS(test.cfg)

			if test.expectedError != "" {
				require.ErrorContains(t, err, test.expectedError)
				return
			}
			require.NoError(t, err)
			require.Equal(t, test.expected, name)
		})
	}
}

func TestResolveTLSRegistersUniqueNames(t *testing.T) {
	caPath, _ := writeTestCertPair(t)
	cfg := TLSConfig{Mode: "verify", CACert: caPath}

	first, err := resolveTLS(cfg)
	require.NoError(t, err)
	second, err := resolveTLS(cfg)
	require.NoError(t, err)

	assert.True(t, strings.HasPrefix(first, tlsNamePrefix))
	assert.NotEqual(t, first, second)
}

// TestResolveTLSRegisteredConfigContent checks that what resolveTLS registers
// under the DSN name is the config buildTLSConfig produced: the driver clones
// the registered config while parsing the DSN.
func TestResolveTLSRegisteredConfigContent(t *testing.T) {
	caPath, keyPath := writeTestCertPair(t)
	name, err := resolveTLS(TLSConfig{
		Mode:       "verify",
		CACert:     caPath,
		ServerName: "db.example.com",
		ClientCert: caPath,
		ClientKey:  keyPath,
	})
	require.NoError(t, err)

	parsed, err := mysql.ParseDSN("u:p@tcp(h:3306)/db?tls=" + name)
	require.NoError(t, err)
	require.NotNil(t, parsed.TLS)

	pemBytes, err := os.ReadFile(caPath)
	require.NoError(t, err)
	expected := x509.NewCertPool()
	require.True(t, expected.AppendCertsFromPEM(pemBytes))

	assert.True(t, parsed.TLS.RootCAs.Equal(expected)) // ca_cert replaces system roots exclusively
	assert.Equal(t, "db.example.com", parsed.TLS.ServerName)
	assert.Equal(t, uint16(tls.VersionTLS12), parsed.TLS.MinVersion)
	assert.Len(t, parsed.TLS.Certificates, 1)
	assert.False(t, parsed.TLS.InsecureSkipVerify)
}

func TestDeregisterTLS(t *testing.T) {
	caPath, _ := writeTestCertPair(t)
	name, err := resolveTLS(TLSConfig{Mode: "verify", CACert: caPath})
	require.NoError(t, err)

	deregisterTLS(name)
	_, err = mysql.ParseDSN("u:p@tcp(h:3306)/db?tls=" + name)
	require.ErrorContains(t, err, "unknown config name")

	// Builtin DSN values are not ours to remove.
	deregisterTLS("true")
	_, err = mysql.ParseDSN("u:p@tcp(h:3306)/db?tls=true")
	require.NoError(t, err)
}

// TestConnectDeregistersTLSName checks that Connect drops its registry entry
// even when the connection attempt fails.
func TestConnectDeregistersTLSName(t *testing.T) {
	caPath, _ := writeTestCertPair(t)
	before := tlsNameCount.Load()

	// Port 1 has no listener, so Connect fails after building the DSN.
	_, err := Connect(&Config{
		Host: "127.0.0.1", Port: 1, Database: "d", Username: "u",
		TLS: TLSConfig{Mode: "verify", CACert: caPath},
	})
	require.Error(t, err)

	name := fmt.Sprintf("%s%d", tlsNamePrefix, before+1)
	_, err = mysql.ParseDSN("u:p@tcp(h:3306)/db?tls=" + name)
	require.ErrorContains(t, err, "unknown config name")
}

// TestConnectSurfacesTLSError pins the exported contract: an invalid TLS
// config must fail Connect with the validation error, before any dial.
func TestConnectSurfacesTLSError(t *testing.T) {
	_, err := Connect(&Config{Host: "localhost", Port: 3306, TLS: TLSConfig{Mode: "bogus"}})
	require.ErrorContains(t, err, "configuring mysql connection")
	require.ErrorContains(t, err, `unknown mode "bogus"`)
}

func TestBuildTLSConfig(t *testing.T) {
	caPath, keyPath := writeTestCertPair(t)

	t.Run("ca cert", func(t *testing.T) {
		cfg, err := buildTLSConfig(TLSConfig{Mode: "verify", CACert: caPath})
		require.NoError(t, err)
		assert.NotNil(t, cfg.RootCAs)
		assert.Equal(t, uint16(tls.VersionTLS12), cfg.MinVersion)
		assert.Empty(t, cfg.ServerName)
		assert.False(t, cfg.InsecureSkipVerify)
	})

	t.Run("server name only keeps system roots", func(t *testing.T) {
		cfg, err := buildTLSConfig(TLSConfig{Mode: "verify", ServerName: "db.example.com"})
		require.NoError(t, err)
		assert.Nil(t, cfg.RootCAs)
		assert.Equal(t, "db.example.com", cfg.ServerName)
	})

	t.Run("client pair", func(t *testing.T) {
		cfg, err := buildTLSConfig(TLSConfig{Mode: "verify", ClientCert: caPath, ClientKey: keyPath})
		require.NoError(t, err)
		require.Len(t, cfg.Certificates, 1)
	})

	t.Run("missing ca file", func(t *testing.T) {
		_, err := buildTLSConfig(TLSConfig{Mode: "verify", CACert: filepath.Join(t.TempDir(), "absent.pem")})
		require.ErrorContains(t, err, "reading ca_cert")
	})

	t.Run("garbage ca file", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "garbage.pem")
		require.NoError(t, os.WriteFile(path, []byte("not a pem"), 0o600))
		_, err := buildTLSConfig(TLSConfig{Mode: "verify", CACert: path})
		require.ErrorContains(t, err, "no certificates found")
	})

	t.Run("mismatched client pair", func(t *testing.T) {
		otherCert, _ := writeTestCertPair(t)
		_, err := buildTLSConfig(TLSConfig{Mode: "verify", ClientCert: otherCert, ClientKey: keyPath})
		require.ErrorContains(t, err, "loading client cert/key")
	})
}

func TestMysqlConfigDSN(t *testing.T) {
	caPath, _ := writeTestCertPair(t)
	base := Config{Host: "localhost", Port: 3306, Database: "flare", Username: "u", Password: "p"}

	tests := []struct {
		name          string
		tls           TLSConfig
		expected      string // substring of the DSN
		notExpected   string
		expectedError string
	}{
		{
			name:        "zero tls keeps plaintext dsn",
			tls:         TLSConfig{},
			notExpected: "tls=",
		},
		{
			name:        "off keeps plaintext dsn",
			tls:         TLSConfig{Mode: "off"},
			notExpected: "tls=",
		},
		{
			name:     "verify",
			tls:      TLSConfig{Mode: "verify"},
			expected: "tls=true",
		},
		{
			name:     "verify with ca",
			tls:      TLSConfig{Mode: "verify", CACert: caPath},
			expected: "tls=" + tlsNamePrefix,
		},
		{
			name:          "invalid tls surfaces error",
			tls:           TLSConfig{Mode: "bogus"},
			expectedError: "tls: unknown mode",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := base
			c.TLS = test.tls
			cfg, err := mysqlConfig(&c)

			if test.expectedError != "" {
				require.ErrorContains(t, err, test.expectedError)
				return
			}
			require.NoError(t, err)
			if test.expected != "" {
				assert.Contains(t, cfg.FormatDSN(), test.expected)
			}
			if test.notExpected != "" {
				assert.NotContains(t, cfg.FormatDSN(), test.notExpected)
			}
		})
	}
}

// TestMysqlConfigPlaintextDSNUnchanged pins the DSN without TLS to the exact
// driver config Connect built before TLS support was added.
func TestMysqlConfigPlaintextDSNUnchanged(t *testing.T) {
	legacy := mysql.Config{
		User:                 "u",
		Passwd:               "p",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "flare",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	cfg := Config{Host: "localhost", Port: 3306, Database: "flare", Username: "u", Password: "p"}

	for _, tlsCfg := range []TLSConfig{{}, {Mode: "off"}} {
		cfg.TLS = tlsCfg
		built, err := mysqlConfig(&cfg)
		require.NoError(t, err)
		require.Equal(t, legacy.FormatDSN(), built.FormatDSN())
	}
}
