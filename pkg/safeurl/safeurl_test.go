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
			name:    "site-local IPv6",
			url:     "http://[fec0::1]/path",
			wantErr: "non-public address",
		},
		{
			name:    "6to4 IPv6",
			url:     "http://[2002:7f00:1::1]/path",
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
			err := Validate(t.Context(), tt.url)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestValidatePublicURL(t *testing.T) {
	// This test requires network access; skip in isolated environments.
	err := Validate(t.Context(), "https://example.com")
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
		// Additional cases.
		{"CGNAT 100.64", "100.64.0.1", false},
		{"CGNAT 100.127", "100.127.255.254", false},
		{"this-network 0.1.2.3", "0.1.2.3", false},
		{"TEST-NET-1", "192.0.2.42", false},
		{"TEST-NET-2", "198.51.100.42", false},
		{"TEST-NET-3", "203.0.113.42", false},
		{"benchmarking 198.18", "198.18.0.1", false},
		{"reserved 240", "240.0.0.1", false},
		{"public v4", "93.184.216.34", true},
		{"public v6", "2606:2800:220:1:248:1893:25c8:1946", true},
		// IPv4 special-use and mapped forms.
		{"6to4 relay anycast", "192.88.99.1", false},
		{"v4-mapped loopback", "::ffff:127.0.0.1", false},
		{"v4-mapped private", "::ffff:10.0.0.1", false},
		{"v4-mapped public", "::ffff:93.184.216.34", true},
		// IPv6 outside 2000::/3.
		{"v4-compatible loopback", "::127.0.0.1", false},
		{"NAT64 well-known prefix", "64:ff9b::7f00:1", false},
		{"NAT64 local-use prefix", "64:ff9b:1::1", false},
		{"discard prefix", "100::1", false},
		{"dummy prefix", "100:0:0:1::1", false},
		{"below 2000::/3", "1fff::1", false},
		{"above 2000::/3", "4000::1", false},
		{"SRv6 SIDs", "5f00::1", false},
		{"ULA", "fd00::1", false},
		{"link-local v6", "fe80::1", false},
		{"site-local v6", "fec0::1", false},
		{"site-local v6 upper", "feff:ffff::1", false},
		{"multicast v6", "ff02::1", false},
		// IPv6 special-use inside 2000::/3.
		{"Teredo", "2001::1", false},
		{"benchmarking v6", "2001:2::1", false},
		{"ORCHID", "2001:10::1", false},
		{"ORCHIDv2 is globally reachable", "2001:20::1", true},
		{"documentation 2001:db8", "2001:db8::1", false},
		{"6to4 embedding loopback", "2002:7f00:1::1", false},
		{"documentation 3fff", "3fff::1", false},
		// 2000::/3 boundaries and another public host.
		{"2000::/3 lower bound", "2000::1", true},
		{"2000::/3 upper bound", "3fff:ffff::1", true},
		{"public v6 2a00", "2a00:1450:4001:80b::200e", true},
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

// TestCheckRedirect verifies that NewClient's redirect policy caps chain
// length and rejects https→http downgrade.
func TestCheckRedirect(t *testing.T) {
	mk := func(scheme string) *http.Request {
		r, err := http.NewRequest(http.MethodGet, scheme+"://example.com", nil)
		require.NoError(t, err)
		return r
	}

	t.Run("first hop allowed", func(t *testing.T) {
		require.NoError(t, checkRedirect(mk("https"), nil))
	})

	t.Run("chain under limit allowed", func(t *testing.T) {
		via := []*http.Request{mk("https"), mk("https"), mk("https")}
		require.NoError(t, checkRedirect(mk("https"), via))
	})

	t.Run("chain at limit rejected", func(t *testing.T) {
		via := make([]*http.Request, maxRedirects)
		for i := range via {
			via[i] = mk("https")
		}
		err := checkRedirect(mk("https"), via)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "stopped after")
	})

	t.Run("https-to-http downgrade rejected", func(t *testing.T) {
		via := []*http.Request{mk("https")}
		err := checkRedirect(mk("http"), via)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "downgrades")
	})

	t.Run("http-to-https upgrade allowed", func(t *testing.T) {
		via := []*http.Request{mk("http")}
		require.NoError(t, checkRedirect(mk("https"), via))
	})
}
