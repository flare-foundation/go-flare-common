package call

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/retry"
	"github.com/stretchr/testify/require"
)

// waitForServerReady polls addr by TCP dial until the listener is up.
func waitForServerReady(t *testing.T, addr string) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		conn, err := net.DialTimeout("tcp", addr, 100*time.Millisecond)
		if err == nil {
			_ = conn.Close()
			return
		}
		time.Sleep(5 * time.Millisecond)
	}
	t.Fatalf("server at %s not ready within deadline", addr)
}

func TestPOST(t *testing.T) {
	port := 4912

	handler := http.NewServeMux()

	type req struct {
		A int `json:"a"`
	}

	type res struct {
		B uint `json:"b"`
	}

	handler.HandleFunc("/abs", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)

		in := req{}
		err := decoder.Decode(&in)
		require.NoError(t, err)
		defer r.Body.Close() //nolint:errcheck // closing request body in test handler

		var b uint
		if in.A < 0 {
			b = uint(-in.A)
		} else {
			b = uint(in.A)
		}

		out := res{B: b}
		responseBytes, err := json.Marshal(out)
		if err != nil {
			w.WriteHeader(http.StatusInsufficientStorage)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		_, err = w.Write(responseBytes)
		if err != nil {
			w.WriteHeader(http.StatusInsufficientStorage)
		}
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	var wg sync.WaitGroup
	wg.Add(1)

	var wg2 sync.WaitGroup
	wg2.Add(1)

	go func() {
		wg2.Done()
		err := server.ListenAndServe()
		require.Error(t, err)

		wg.Done()
	}()

	request := req{
		A: -10,
	}

	reqMarshaled, err := json.Marshal(request)
	require.NoError(t, err)

	url := fmt.Sprintf("http://localhost:%d/abs", port)

	wg2.Wait()
	waitForServerReady(t, fmt.Sprintf("localhost:%d", port))
	responseRaw, err := PostRawWithRetry[res](context.Background(), url, NoAPIKey, bytes.NewReader(reqMarshaled), Params{
		Timeout:         3 * time.Second,
		MaxResponseSize: 300,
	}, []int{}, retry.Params{
		MaxAttempts: 2,
		Delay:       100 * time.Millisecond,
		Timeout:     3 * time.Second,
	})
	require.NoError(t, err)

	response, err := Post[req, res](context.Background(), url, NoAPIKey, request, Params{
		Timeout:         3 * time.Second,
		MaxResponseSize: 300,
	})
	require.NoError(t, err)

	expected := res{10}

	require.Equal(t, expected, *responseRaw.Message)
	require.Equal(t, expected, *response.Message)

	err = server.Shutdown(context.Background())
	require.NoError(t, err)

	wg.Wait()
}

type echoBody struct{}

func newAlwaysStatus(t *testing.T, code int) (*httptest.Server, *atomic.Int32) {
	t.Helper()
	var hits atomic.Int32
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.WriteHeader(code)
	}))
	t.Cleanup(s.Close)
	return s, &hits
}

func TestPostWithRetry_NoRetryStatuses_SuppressesRetry(t *testing.T) {
	srv, hits := newAlwaysStatus(t, http.StatusBadRequest)

	resp, err := PostWithRetry[echoBody, echoBody](
		context.Background(),
		srv.URL,
		NoAPIKey,
		echoBody{},
		Params{
			Timeout:         time.Second,
			MaxResponseSize: 256,
			NoRetryStatuses: []int{http.StatusBadRequest},
		},
		nil,
		retry.Params{MaxAttempts: 5, Delay: 0, Timeout: time.Second},
	)
	require.Error(t, err, "expected the HTTP error to be surfaced")
	require.Equal(t, http.StatusBadRequest, resp.Status)
	require.Equal(t, int32(1), hits.Load(), "must not retry when status is in NoRetryStatuses")
}

func TestPostWithRetry_LegacyDefaultRetries(t *testing.T) {
	srv, hits := newAlwaysStatus(t, http.StatusInternalServerError)

	_, err := PostWithRetry[echoBody, echoBody](
		context.Background(),
		srv.URL,
		NoAPIKey,
		echoBody{},
		Params{Timeout: time.Second, MaxResponseSize: 256},
		nil,
		retry.Params{MaxAttempts: 3, Delay: 0, Timeout: time.Second},
	)
	require.Error(t, err)
	require.Equal(t, int32(3), hits.Load(), "default (nil NoRetryStatuses) must retry to exhaustion")
}

func TestPostWithRetry_AcceptableFailWinsOverNoRetry(t *testing.T) {
	srv, hits := newAlwaysStatus(t, http.StatusTeapot)

	resp, err := PostWithRetry[echoBody, echoBody](
		context.Background(),
		srv.URL,
		NoAPIKey,
		echoBody{},
		Params{
			Timeout:         time.Second,
			MaxResponseSize: 256,
			NoRetryStatuses: []int{http.StatusTeapot},
		},
		[]int{http.StatusTeapot}, // acceptableFail
		retry.Params{MaxAttempts: 3, Delay: 0, Timeout: time.Second},
	)
	require.NoError(t, err, "acceptableFail must eat the error even when status is also in NoRetryStatuses")
	require.Equal(t, http.StatusTeapot, resp.Status)
	require.Equal(t, int32(1), hits.Load())
}
