package safeurl

import (
	"fmt"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr string
	}{
		{
			name:    "empty string",
			url:     "",
			wantErr: "unsupported scheme",
		},
		{
			name:    "file scheme",
			url:     "file:///etc/passwd",
			wantErr: "unsupported scheme",
		},
		{
			name:    "ftp scheme",
			url:     "ftp://example.com/file",
			wantErr: "unsupported scheme",
		},
		{
			name:    "no scheme",
			url:     "example.com/path",
			wantErr: "unsupported scheme",
		},
		{
			name:    "loopback IPv4",
			url:     "http://127.0.0.1/path",
			wantErr: "non-public address",
		},
		{
			name:    "loopback IPv6",
			url:     "http://[::1]/path",
			wantErr: "non-public address",
		},
		{
			name:    "private 10.x",
			url:     "https://10.0.0.1/api",
			wantErr: "non-public address",
		},
		{
			name:    "private 172.16.x",
			url:     "https://172.16.0.1/api",
			wantErr: "non-public address",
		},
		{
			name:    "private 192.168.x",
			url:     "https://192.168.1.1/api",
			wantErr: "non-public address",
		},
		{
			name:    "link-local",
			url:     "http://169.254.169.254/latest/meta-data",
			wantErr: "non-public address",
		},
		{
			name:    "empty host",
			url:     "http:///path",
			wantErr: "empty host",
		},
		{
			name:    "unresolvable host",
			url:     "https://this-host-does-not-exist-abc123xyz.invalid",
			wantErr: "resolving host",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.url)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestValidatePublicURL(t *testing.T) {
	// This test requires network access; skip in isolated environments.
	err := Validate("https://example.com")
	assert.NoError(t, err)
}

func TestIsPublicIP(t *testing.T) {
	tests := []struct {
		name   string
		ip     string
		public bool
	}{
		{"loopback v4", "127.0.0.1", false},
		{"loopback v6", "::1", false},
		{"private 10", "10.0.0.1", false},
		{"private 172", "172.16.0.1", false},
		{"private 192", "192.168.0.1", false},
		{"link-local", "169.254.1.1", false},
		{"unspecified v4", "0.0.0.0", false},
		{"unspecified v6", "::", false},
		{"multicast", "224.0.0.1", false},
		{"public v4", "93.184.216.34", true},
		{"public v6", "2606:2800:220:1:248:1893:25c8:1946", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := parseIP(t, tt.ip)
			assert.Equal(t, tt.public, isPublicIP(ip))
		})
	}
}

func TestNewClientBlocksLoopback(t *testing.T) {
	// Start a server on localhost.
	handler := http.NewServeMux()
	handler.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	defer listener.Close() //nolint:errcheck // test cleanup

	server := &http.Server{Handler: handler}
	go server.Serve(listener) //nolint:errcheck // test server
	defer server.Close()      //nolint:errcheck // test cleanup

	url := fmt.Sprintf("http://%s/", listener.Addr().String())

	// A safe client must refuse to connect to loopback.
	client := NewClient(5 * time.Second)
	_, err = client.Get(url) //nolint:bodyclose // response is nil on error
	require.Error(t, err)
	assert.Contains(t, err.Error(), "non-public address")

	// Verify the server is actually reachable with a normal client.
	resp, err := http.Get(url) //nolint:bodyclose,gosec,noctx // minimal test call
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func parseIP(t *testing.T, s string) net.IP {
	t.Helper()

	ip := net.ParseIP(s)
	require.NotNil(t, ip, "failed to parse IP %q", s)

	return ip
}
