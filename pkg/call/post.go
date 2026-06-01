// Package call provides HTTP POST request utilities with retry support.
package call

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"slices"
	"strings"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/retry"
)

// APIKey holds an HTTP header name-value pair for API authentication.
type APIKey struct {
	Name string
	Key  string
}

// Params configures HTTP request behavior.
type Params struct {
	Timeout         time.Duration // maximal time to wait for a response
	MaxResponseSize int64         // response body cap in bytes; 0 means unbounded (PostRawWithRetry requires > 0)
	// Transport sets the HTTP client transport. Nil uses http.DefaultTransport (no SSRF filtering); pass safeurl.NewTransport() to enable it.
	Transport http.RoundTripper

	// NoRetryStatuses lists HTTP response statuses that must NOT trigger a retry
	// in the *WithRetry helpers. The status and any error are still surfaced to
	// the caller; only the retry loop is suppressed. When nil/empty (default),
	// the helpers retry on any error from PostRaw (legacy behavior; dangerous for
	// non-idempotent POSTs). Independent of acceptableFail, which converts a status
	// to a successful result; NoRetryStatuses preserves the error.
	NoRetryStatuses []int
}

// NoAPIKey is an empty APIKey that skips adding an authentication header.
var NoAPIKey = APIKey{"", ""}

// Response holds the HTTP status code and the unmarshaled response body.
type Response[T any] struct {
	Status  int
	Message *T
}

// PostRaw sends a post request with body and apiKey in header to url and unmarshals the response to a struct of type T.
func PostRaw[T any](ctx context.Context, url string, apiKey APIKey, body io.Reader, p Params) (Response[T], error) {
	resOut := Response[T]{}

	client := &http.Client{Timeout: p.Timeout, Transport: p.Transport}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return resOut, err
	}

	request.Header.Set("Content-Type", "application/json")

	if len(apiKey.Name) > 0 {
		request.Header.Set(apiKey.Name, apiKey.Key)
	}

	resp, err := client.Do(request)
	if err != nil {
		return resOut, err
	}
	defer resp.Body.Close() //nolint:errcheck // closing response body, error not actionable

	resOut.Status = resp.StatusCode

	if resp.StatusCode != http.StatusOK {
		if isTextPlain(resp.Header.Get("Content-Type")) {
			buf := new(strings.Builder)
			_, err := io.Copy(buf, capReader(resp.Body, maxErrorReasonSize))
			if err == nil {
				return resOut, fmt.Errorf("request responded with code %d, reason: %s", resp.StatusCode, buf.String())
			}
		}

		return resOut, fmt.Errorf("request responded with code %d", resp.StatusCode)
	}

	decoder := json.NewDecoder(capReader(resp.Body, p.MaxResponseSize))

	response := new(T)
	err = decoder.Decode(response)
	if err != nil {
		return resOut, fmt.Errorf("decoding response: %w", err)
	}
	resOut.Message = response

	return resOut, nil
}

// PostRawWithRetry sends a post request and retries on unsuccessful attempts according to retry parameters.
// The request body is fully buffered up to p.MaxResponseSize to enable retries; oversize bodies are rejected.
func PostRawWithRetry[T any](ctx context.Context, url string, apiKey APIKey, body io.Reader, p Params, acceptableFail []int, rp retry.Params) (Response[T], error) {
	if p.MaxResponseSize <= 0 {
		return Response[T]{}, errors.New("MaxResponseSize must be set to buffer retry body")
	}
	limited := &io.LimitedReader{R: body, N: p.MaxResponseSize + 1}
	bodyBytes, err := io.ReadAll(limited)
	if err != nil {
		return Response[T]{}, fmt.Errorf("reading request body: %w", err)
	}
	if int64(len(bodyBytes)) > p.MaxResponseSize {
		return Response[T]{}, fmt.Errorf("request body exceeds MaxResponseSize (%d bytes)", p.MaxResponseSize)
	}

	fn := func() (Response[T], error) {
		return PostRaw[T](ctx, url, apiKey, bytes.NewReader(bodyBytes), p)
	}

	return executeWithRetry(ctx, fn, p.NoRetryStatuses, acceptableFail, rp)
}

// maxErrorReasonSize caps the text/plain reason read from a non-OK response.
const maxErrorReasonSize = 64 << 10 // 64 KiB

// capReader returns r capped to n bytes, or r unchanged if n <= 0 (unbounded).
func capReader(r io.Reader, n int64) io.Reader {
	if n <= 0 {
		return r
	}
	return &io.LimitedReader{R: r, N: n}
}

// isTextPlain reports whether contentType is a text/plain media type, ignoring charset and whitespace.
func isTextPlain(contentType string) bool {
	if contentType == "" {
		return false
	}
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}
	return mediaType == "text/plain"
}

// Post sends a post request with marshaled body and apiKey in header to url and unmarshals the response of type T.
func Post[S, T any](ctx context.Context, url string, apiKey APIKey, body S, p Params) (Response[T], error) {
	b, err := json.Marshal(body)
	if err != nil {
		return Response[T]{}, err
	}

	return PostRaw[T](ctx, url, apiKey, bytes.NewReader(b), p)
}

// PostWithRetry sends a post request with marshaled body and retries on unsuccessful attempts according to retry parameters.
func PostWithRetry[S, T any](ctx context.Context, url string, apiKey APIKey, body S, p Params, acceptableFail []int, rp retry.Params) (Response[T], error) {
	fn := func() (Response[T], error) {
		return Post[S, T](ctx, url, apiKey, body, p)
	}

	return executeWithRetry(ctx, fn, p.NoRetryStatuses, acceptableFail, rp)
}

// executeWithRetry runs fn through retry.Execute, applying the *WithRetry helpers'
// per-status gating. acceptableFail eats the error and stops retrying (legacy);
// noRetry preserves the error and stops retrying; otherwise the error triggers retry.
func executeWithRetry[T any](
	ctx context.Context,
	fn func() (Response[T], error),
	noRetry []int,
	acceptableFail []int,
	rp retry.Params,
) (Response[T], error) {
	type result struct {
		r   Response[T]
		err error
	}
	// retry.Execute keeps Value only from the successful attempt, so on
	// exhaustion the last Response (and its Status) would be lost; capture it.
	var last Response[T]
	wrapped := func() (result, error) {
		r, err := fn()
		last = r
		if err == nil {
			return result{r: r}, nil
		}
		if slices.Contains(acceptableFail, r.Status) {
			return result{r: r}, nil
		}
		if slices.Contains(noRetry, r.Status) {
			return result{r: r, err: err}, nil
		}
		return result{r: r}, err
	}
	res := retry.Execute(ctx, wrapped, rp)
	if res.Value.err != nil {
		return res.Value.r, res.Value.err
	}
	if res.Success {
		return res.Value.r, res.Err
	}
	return last, res.Err
}
