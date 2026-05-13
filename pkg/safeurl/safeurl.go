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

// NewTransport returns an http.Transport whose dialer resolves DNS and validates
// every resolved IP before opening a connection. This eliminates the TOCTOU
// window that exists when validation and dialing are separate steps.
func NewTransport() *http.Transport {
	d := &net.Dialer{}
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

	return true
}
