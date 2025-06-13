package call

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/retry"
	"github.com/stretchr/testify/require"
)

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
		defer r.Body.Close() //nolint:errcheck

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
	time.Sleep(10 * time.Millisecond)
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
