package call

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/retry"
)

type APIKey struct {
	Name string
	Key  string
}

type CallParams struct {
	Timeout         time.Duration // maximal time to wait for a response
	MaxResponseSize int64         // maximal response size in bytes
}

var NoAPIKey = APIKey{"", ""}

// PostRaw sends a post request with body and apiKey in header to url and unmarshals the response to a struct of type T.
func PostRaw[T any](ctx context.Context, url string, apiKey APIKey, body io.Reader, p CallParams) (*T, error) {
	client := &http.Client{Timeout: p.Timeout}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")

	if len(apiKey.Name) > 0 {
		request.Header.Set(apiKey.Name, apiKey.Key)
	}

	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request responded with code %d", resp.StatusCode)
	}

	respLimited := &io.LimitedReader{R: resp.Body, N: p.MaxResponseSize}
	defer resp.Body.Close()

	decoder := json.NewDecoder(respLimited)
	decoder.DisallowUnknownFields()

	response := new(T)

	err = decoder.Decode(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// PostRawWithRetry sends a post request and retries on unsuccessful attempts according to retry parameters.
func PostRawWithRetry[T any](ctx context.Context, url string, apiKey APIKey, body io.Reader, p CallParams, rp retry.Params) (*T, error) {
	fn := func() (*T, error) {
		return PostRaw[T](ctx, url, apiKey, body, p)
	}

	res := retry.Execute(ctx, fn, rp)

	return res.Value, res.Err
}

// Post sends a post request with marshaled body and apiKey in header to url and unmarshals the response of type T.
func Post[S, T any](ctx context.Context, url string, apiKey APIKey, body S, p CallParams) (*T, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return PostRaw[T](ctx, url, apiKey, bytes.NewReader(b), p)
}

// PostWithRetry sends a post request with marshaled body and retries on unsuccessful attempts according to retry parameters.
func PostWithRetry[S, T any](ctx context.Context, url string, apiKey APIKey, body S, p CallParams, rp retry.Params) (*T, error) {
	fn := func() (*T, error) {
		return Post[S, T](ctx, url, apiKey, body, p)
	}

	res := retry.Execute(ctx, fn, rp)

	return res.Value, res.Err
}
