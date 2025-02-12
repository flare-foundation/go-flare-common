package call

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

type RequestData struct {
	A int    `json:"a"`
	B string `json:"b"`
	C []byte `json:"c"`
}

func MockServerForTest(t *testing.T, port int) *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		MockResponse(t, writer, request)
	})

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      r,
	}

	return server
}

func MockResponse(t *testing.T, writer http.ResponseWriter, request *http.Request) {
	body, err := io.ReadAll(request.Body)
	require.NoError(t, err)

	var rd RequestData
	err = json.Unmarshal(body, &rd)
	require.NoError(t, err)

	rd.A++

	responseBytes, err := json.Marshal(rd)
	require.NoError(t, err)

	_, err = writer.Write(responseBytes)
	require.NoError(t, err)
}

func TestCallPostNoKey(t *testing.T) {
	ctx := context.Background()
	url := "http://localhost:5010"
	apiKey := ""

	server := MockServerForTest(t, 5010)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		err := server.ListenAndServe()
		require.ErrorIs(t, err, http.ErrServerClosed)
		wg.Done()
	}()

	t.Run("happy path", func(t *testing.T) {
		request := RequestData{
			A: 0,
			B: "b",
			C: []byte{0, 1, 2, 3},
		}

		body, err := json.Marshal(request)

		require.NoError(t, err)

		result := new(RequestData)

		err = Post(ctx, url, apiKey, body, result, Params{
			Timeout:               0,
			MaxResponseSize:       0,
			DisallowUnknownFields: false,
		})
		require.NoError(t, err)
		require.Equal(t, 1, result.A)
		require.Equal(t, "b", result.B)
		require.Equal(t, []byte{0, 1, 2, 3}, result.C)
	})

	t.Run("unknown field", func(t *testing.T) {
		type OtherRequestData struct {
			X int
			Y string
		}

		request := OtherRequestData{
			X: 0,
			Y: "s",
		}

		body, err := json.Marshal(request)
		require.NoError(t, err)

		result := new(OtherRequestData)

		err = Post(ctx, url, apiKey, body, result, Params{
			Timeout:               0,
			MaxResponseSize:       0,
			DisallowUnknownFields: true,
		})
		require.Error(t, err)
	})

	t.Run("limit response", func(t *testing.T) {
		request := RequestData{
			A: 0,
			B: "b",
			C: []byte{0, 1, 2, 3},
		}

		body, err := json.Marshal(request)

		require.NoError(t, err)

		result := new(RequestData)

		err = Post(ctx, url, apiKey, body, result, Params{
			Timeout:               0,
			MaxResponseSize:       10,
			DisallowUnknownFields: false,
		})
		require.Error(t, err)
	})

	err := server.Shutdown(ctx)
	require.NoError(t, err)
	wg.Wait()
}
