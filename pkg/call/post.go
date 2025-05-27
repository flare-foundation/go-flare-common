package call

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"slices"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/retry"
)

type APIKey struct {
	Name string
	Key  string
}

type Params struct {
	Timeout         time.Duration // maximal time to wait for a response
	MaxResponseSize int64         // maximal response size in bytes
}

var NoAPIKey = APIKey{"", ""}

type Response[T any] struct {
	Status  int
	Message *T
}

// PostRaw sends a post request with body and apiKey in header to url and unmarshals the response to a struct of type T.
func PostRaw[T any](ctx context.Context, url string, apiKey APIKey, body io.Reader, p Params) (Response[T], error) {
	resOut := Response[T]{}

	client := &http.Client{Timeout: p.Timeout}

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

	resOut.Status = resp.StatusCode

	if resp.StatusCode != http.StatusOK {
		return resOut, fmt.Errorf("request responded with code %d", resp.StatusCode)
	}

	respLimited := &io.LimitedReader{R: resp.Body, N: p.MaxResponseSize}
	defer resp.Body.Close() //nolint:errcheck

	decoder := json.NewDecoder(respLimited)
	// decoder.DisallowUnknownFields() // todo make this optional maybe

	response := new(T)
	err = decoder.Decode(response)
	if err != nil {
		return resOut, fmt.Errorf("decoding response: %v", err)
	}
	resOut.Message = response

	return resOut, nil
}

// PostRawWithRetry sends a post request and retries on unsuccessful attempts according to retry parameters.
func PostRawWithRetry[T any](ctx context.Context, url string, apiKey APIKey, body io.Reader, p Params, acceptableFail []int, rp retry.Params) (Response[T], error) {
	fn := func() (Response[T], error) {
		r, err := PostRaw[T](ctx, url, apiKey, body, p)
		if err != nil && !slices.Contains(acceptableFail, r.Status) {
			return r, err
		}
		return r, nil
	}

	res := retry.Execute(ctx, fn, rp)

	return res.Value, res.Err
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
		r, err := Post[S, T](ctx, url, apiKey, body, p)
		if err != nil && !slices.Contains(acceptableFail, r.Status) {
			return r, err
		}
		return r, nil
	}

	res := retry.Execute(ctx, fn, rp)

	return res.Value, res.Err
}
