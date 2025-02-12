package call

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Params struct {
	Timeout               time.Duration
	MaxResponseSize       int64
	DisallowUnknownFields bool
}

// Post
//
// T must be a pointer type, otherwise the function will panic.
func Post[T any](ctx context.Context, url string, apiKey string, body []byte, response T, params Params) error {
	client := &http.Client{}

	if params.Timeout > 0 {
		client.Timeout = params.Timeout
	}

	request, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", apiKey)

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request responded with code %d", resp.StatusCode)
	}

	var reader io.Reader

	if params.MaxResponseSize > 0 {
		reader = &io.LimitedReader{R: resp.Body, N: params.MaxResponseSize}
	} else {
		reader = resp.Body
	}

	decoder := json.NewDecoder(reader)
	if params.DisallowUnknownFields {
		decoder.DisallowUnknownFields()
	}

	err = decoder.Decode(response)
	if err != nil {
		return err
	}

	return nil
}
