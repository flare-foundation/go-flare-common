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

// Validate checks that rawURL is a well-formed HTTP(S) URL whose host does
// not resolve to a private, loopback, or link-local address. This prevents
// SSRF attacks where an attacker-controlled URL could reach internal services.
//
// Use Validate for early rejection of untrusted input. For the authoritative
// check at connection time (eliminating TOCTOU), use NewTransport or NewClient.
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

// NewClient returns an http.Client with NewTransport and the given timeout.
func NewClient(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout:   timeout,
		Transport: NewTransport(),
	}
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
