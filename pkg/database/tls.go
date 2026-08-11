package database

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync/atomic"

	"github.com/go-sql-driver/mysql"
)

// TLSMode selects how the MySQL connection is encrypted. See TLSConfig.
type TLSMode string

// Modes accepted in TLSConfig.Mode.
const (
	// TLSModeOff (or empty) keeps the connection plaintext, matching the
	// historical behavior; every other TLSConfig field must be empty.
	TLSModeOff TLSMode = "off"
	// TLSModeSkipVerify encrypts but does not authenticate the server, so it
	// does not protect against an active man-in-the-middle.
	TLSModeSkipVerify TLSMode = "skip-verify"
	// TLSModeVerify verifies the server certificate and hostname — against the
	// system roots by default, or exclusively against TLSConfig.CACert when set.
	TLSModeVerify TLSMode = "verify"
)

// TLSConfig configures TLS for the MySQL connection. The zero value keeps the
// connection unencrypted.
//
// CACert, ServerName, ClientCert and ClientKey are only valid with
// TLSModeVerify. ServerName overrides the hostname expected in the server
// certificate (defaults to Config.Host). ClientCert and ClientKey must be set
// together and enable mutual TLS.
type TLSConfig struct {
	Mode       TLSMode `toml:"mode" envconfig:"DB_TLS_MODE"`
	CACert     string  `toml:"ca_cert" envconfig:"DB_TLS_CA_CERT"`
	ServerName string  `toml:"server_name" envconfig:"DB_TLS_SERVER_NAME"`
	ClientCert string  `toml:"client_cert" envconfig:"DB_TLS_CLIENT_CERT"`
	ClientKey  string  `toml:"client_key" envconfig:"DB_TLS_CLIENT_KEY"`
}

// custom reports whether t sets any field that needs a custom driver TLS config.
func (t TLSConfig) custom() bool {
	return t.CACert != "" || t.ServerName != "" || t.ClientCert != "" || t.ClientKey != ""
}

// tlsNamePrefix marks driver TLS config names owned by this package.
const tlsNamePrefix = "flare"

// tlsNameCount makes registered driver TLS config names unique per Connect.
var tlsNameCount atomic.Uint64

// resolveTLS validates t and returns the value for the DSN tls parameter,
// registering a custom driver config when one is needed. Empty means plaintext.
// The caller drops the registration with deregisterTLS once the DSN is parsed.
func resolveTLS(t TLSConfig) (string, error) {
	switch t.Mode {
	case "", TLSModeOff:
		if t.custom() {
			return "", fmt.Errorf("ca_cert, server_name, client_cert and client_key require mode %q (mode is %q)", TLSModeVerify, t.Mode)
		}
		return "", nil
	case TLSModeSkipVerify:
		if t.custom() {
			return "", fmt.Errorf("ca_cert, server_name, client_cert and client_key require mode %q (mode is %q)", TLSModeVerify, t.Mode)
		}
		return string(TLSModeSkipVerify), nil
	case TLSModeVerify:
		if !t.custom() {
			return "true", nil // driver builtin: system roots + hostname check
		}
		tlsCfg, err := buildTLSConfig(t)
		if err != nil {
			return "", err
		}
		name := fmt.Sprintf("%s%d", tlsNamePrefix, tlsNameCount.Add(1))
		if err := mysql.RegisterTLSConfig(name, tlsCfg); err != nil {
			return "", fmt.Errorf("registering driver tls config: %w", err)
		}
		return name, nil
	default:
		return "", fmt.Errorf("unknown mode %q (valid: %q, %q, %q)", t.Mode, TLSModeOff, TLSModeSkipVerify, TLSModeVerify)
	}
}

// deregisterTLS drops a driver TLS config registered by resolveTLS; other
// values (builtins like "true" or "skip-verify") are left alone. Safe once the
// DSN has been parsed — the driver clones the config during parsing and never
// consults the registry again.
func deregisterTLS(name string) {
	if strings.HasPrefix(name, tlsNamePrefix) {
		mysql.DeregisterTLSConfig(name)
	}
}

// buildTLSConfig creates the tls.Config for TLSModeVerify with custom fields.
// The driver fills ServerName from Config.Host when left empty.
func buildTLSConfig(t TLSConfig) (*tls.Config, error) {
	if (t.ClientCert == "") != (t.ClientKey == "") {
		return nil, errors.New("client_cert and client_key must be set together")
	}

	cfg := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: t.ServerName,
	}

	if t.CACert != "" {
		pem, err := os.ReadFile(t.CACert)
		if err != nil {
			return nil, fmt.Errorf("reading ca_cert: %w", err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("no certificates found in ca_cert %s", t.CACert)
		}
		cfg.RootCAs = pool
	}

	if t.ClientCert != "" {
		cert, err := tls.LoadX509KeyPair(t.ClientCert, t.ClientKey)
		if err != nil {
			return nil, fmt.Errorf("loading client cert/key: %w", err)
		}
		cfg.Certificates = []tls.Certificate{cert}
	}

	return cfg, nil
}
