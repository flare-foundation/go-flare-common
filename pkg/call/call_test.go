package call

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
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

func TestCallNoKey(t *testing.T) {
	ctx := context.Background()
	url := "http://localhost:5010"
	apiKey := "a"

	server := MockServerForTest(t, 5010)

	go func() {
		err := server.ListenAndServe()
		require.Error(t, err)
	}()

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

	err = server.Shutdown(ctx)
	require.NoError(t, err)
}
