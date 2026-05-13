// Package safeurl provides HTTP client utilities that guard against SSRF attacks
// by validating URLs and resolved IP addresses before making connections.
package safeurl

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

// Validate is a non-authoritative, optimistic check that rawURL is a
// well-formed HTTP(S) URL whose host currently resolves to public addresses
// only. It is NOT a defense against SSRF on its own: there is a TOCTOU gap
// between the resolution Validate performs and the resolution the eventual
// http.Client performs at dial time. A DNS-rebinding attacker can satisfy
// Validate and still steer the connection to an internal address.
//
// Use Validate only for early rejection of obviously bad input. For the
// authoritative check (resolve-and-validate-at-dial, eliminating TOCTOU),
// build the client with NewTransport or NewClient — those validate every
// resolved IP again inside the dialer, so the validated address is the
// connected address.
func Validate(rawURL string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("parsing URL: %w", err)
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return fmt.Errorf("unsupported scheme %q, only http and https are allowed", u.Scheme)
	}

	host := u.Hostname()
	if host == "" {
		return errors.New("empty host")
	}

	ips, err := net.LookupHost(host)
	if err != nil {
		return fmt.Errorf("resolving host %q: %w", host, err)
	}

	for _, ipStr := range ips {
		ip := net.ParseIP(ipStr)
		if ip == nil {
			return fmt.Errorf("invalid IP %q for host %q", ipStr, host)
		}

		if !isPublicIP(ip) {
			return fmt.Errorf("host %q resolves to non-public address %s", host, ipStr)
		}
	}

	return nil
}

// Default transport timeouts (audit M21). The defaults bound slowloris-style
// attacks where a malicious endpoint trickles bytes to keep a goroutine and
// connection pinned indefinitely.
const (
	defaultConnectTimeout      = 10 * time.Second
	defaultTLSHandshakeTimeout = 10 * time.Second
	defaultResponseHeaderTime  = 30 * time.Second
	defaultExpectContinueTime  = 1 * time.Second
	defaultIdleConnTimeout     = 90 * time.Second
)

// NewTransport returns an http.Transport whose dialer resolves DNS and validates
// every resolved IP before opening a connection. This eliminates the TOCTOU
// window that exists when validation and dialing are separate steps. The
// transport carries connect, TLS-handshake, response-header, expect-continue,
// and idle-connection timeouts so a slow or malicious remote cannot pin
// resources indefinitely (audit M21).
func NewTransport() *http.Transport {
	d := &net.Dialer{Timeout: defaultConnectTimeout}
	return &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			host, port, err := net.SplitHostPort(addr)
			if err != nil {
				return nil, fmt.Errorf("parsing dial address: %w", err)
			}

			ips, err := net.DefaultResolver.LookupHost(ctx, host)
			if err != nil {
				return nil, fmt.Errorf("resolving host %q: %w", host, err)
			}

			var validIPs []string
			for _, ipStr := range ips {
				ip := net.ParseIP(ipStr)
				if ip == nil {
					return nil, fmt.Errorf("invalid resolved IP %q for host %q", ipStr, host)
				}
				if !isPublicIP(ip) {
					return nil, fmt.Errorf("host %q resolves to non-public address %s", host, ipStr)
				}
				validIPs = append(validIPs, ipStr)
			}

			// Try each resolved IP to preserve fallback behavior for multi-homed hosts.
			var lastErr error
			for _, ipStr := range validIPs {
				conn, err := d.DialContext(ctx, network, net.JoinHostPort(ipStr, port))
				if err == nil {
					return conn, nil
				}
				lastErr = err
			}
			return nil, fmt.Errorf("connecting to host %q: %w", host, lastErr)
		},
		TLSHandshakeTimeout:   defaultTLSHandshakeTimeout,
		ResponseHeaderTimeout: defaultResponseHeaderTime,
		ExpectContinueTimeout: defaultExpectContinueTime,
		IdleConnTimeout:       defaultIdleConnTimeout,
	}
}

// maxRedirects caps the redirect chain on clients built by NewClient. The
// default http.Client follows up to 10 redirects; that's enough room for a
// malicious origin to chain hops, downgrade scheme, or shuffle hosts. Five
// is generous for legitimate traffic.
const maxRedirects = 5

// NewClient returns an http.Client with NewTransport, the given timeout, and
// a CheckRedirect that caps chain length and rejects https→http downgrade.
// Per-hop IP re-validation is already performed by NewTransport's dialer, so
// each redirected target is dial-checked against public-IP policy regardless
// of CheckRedirect.
func NewClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout:       timeout,
		Transport:     NewTransport(),
		CheckRedirect: checkRedirect,
	}
}

// checkRedirect enforces redirect-chain length cap and rejects https→http
// downgrades. It does not need to re-resolve the next-hop host: NewTransport's
// dialer validates every resolved IP on every connection, so downgrading via
// DNS rebinding or chain-hopping to an internal host still fails at dial time.
func checkRedirect(req *http.Request, via []*http.Request) error {
	if len(via) >= maxRedirects {
		return fmt.Errorf("stopped after %d redirects", maxRedirects)
	}
	if len(via) > 0 && via[0].URL.Scheme == "https" && req.URL.Scheme == "http" {
		return errors.New("redirect downgrades https to http")
	}
	return nil
}

// nonPublicCIDRs holds extra address ranges that net.IP's built-in predicates
// do not cover (audit M18):
//   - 100.64.0.0/10: CGNAT (RFC 6598). Used by ISPs for shared address space;
//     resolvable to a carrier's internal infrastructure.
//   - 0.0.0.0/8: "this network" (RFC 1122). The all-zeros address is
//     IsUnspecified, but any other 0.x.x.x is also reserved and not routable.
//   - 192.0.0.0/24, 192.0.2.0/24, 198.18.0.0/15, 198.51.100.0/24,
//     203.0.113.0/24, 240.0.0.0/4: IANA special-use ranges.
//   - 64:ff9b::/96 (IPv4-IPv6 well-known prefix), 64:ff9b:1::/48: IPv4
//     translation.
var nonPublicCIDRs = func() []*net.IPNet {
	specs := []string{
		"100.64.0.0/10",   // CGNAT
		"0.0.0.0/8",       // "this network"
		"192.0.0.0/24",    // IETF protocol assignments
		"192.0.2.0/24",    // TEST-NET-1
		"198.18.0.0/15",   // benchmarking
		"198.51.100.0/24", // TEST-NET-2
		"203.0.113.0/24",  // TEST-NET-3
		"240.0.0.0/4",     // Reserved (former Class E)
		// Note: do not add "::ffff:0:0/96" — Go's net.IP represents all
		// IPv4 addresses in that range internally, so it would match every
		// public IPv4 host.
		"64:ff9b::/96",   // IPv4-IPv6 well-known prefix
		"64:ff9b:1::/48", // IPv4-IPv6 local-use prefix
		"100::/64",       // Discard prefix
	}
	out := make([]*net.IPNet, 0, len(specs))
	for _, s := range specs {
		_, n, err := net.ParseCIDR(s)
		if err == nil {
			out = append(out, n)
		}
	}
	return out
}()

// isPublicIP returns true if the IP is a globally routable unicast address.
func isPublicIP(ip net.IP) bool {
	if ip.IsLoopback() ||
		ip.IsPrivate() ||
		ip.IsLinkLocalUnicast() ||
		ip.IsLinkLocalMulticast() ||
		ip.IsUnspecified() ||
		ip.IsMulticast() {
		return false
	}
	for _, cidr := range nonPublicCIDRs {
		if cidr.Contains(ip) {
			return false
		}
	}
	return true
}
