package signing

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

const (
	port    = 6739
	apiKey  = "krompirusa"
	prvSeed = "nekejCesarNihceNeUgane"
)

var url = fmt.Sprintf("http://localhost:%d/sign", port)

func prepareServer() *Signer {
	randomStr := []byte(prvSeed)
	pkStr := crypto.Keccak256(randomStr)
	prv, _ := crypto.ToECDSA(pkStr)

	cfg := Config{
		Addr:       fmt.Sprintf(":%d", port),
		APIKeyName: "X-API-KEY",
		APIKeys:    []string{apiKey},
	}

	return New(cfg, prv)
}

func prepareRequestBody(n uint64) RequestBody {
	hashes := make([]common.Hash, n)

	for j := range n {
		hashes[j] = crypto.Keccak256Hash([]byte(strconv.FormatUint(j, 10)))
	}

	return RequestBody{
		Hashes: hashes,
	}
}

func prepareRequest(t *testing.T, body any) *http.Request {
	encodedBody, err := json.Marshal(body)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(encodedBody))
	require.NoError(t, err)

	return req
}

func sendRequest(t *testing.T, r *http.Request) *http.Response {
	client := &http.Client{}
	resp, err := client.Do(r)
	require.NoError(t, err)
	return resp

}

func TestSigner(t *testing.T) {
	signer := prepareServer()

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		err := signer.Run(ctx)
		require.Error(t, err)
		wg.Done()
	}()

	// happy path
	const nuOfHashes uint64 = 10

	body := prepareRequestBody(nuOfHashes)

	request := prepareRequest(t, body)

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", apiKey)

	resp := sendRequest(t, request)
	require.Equal(t, http.StatusOK, resp.StatusCode)
	decoder := json.NewDecoder(resp.Body)

	responseBody := ResponseBody{}
	err := decoder.Decode(&responseBody)
	require.NoError(t, err)
	require.Len(t, responseBody.Signatures, int(nuOfHashes))

	// to big request
	bodyToBig := prepareRequestBody(2000)
	requestFail0 := prepareRequest(t, bodyToBig)

	requestFail0.Header.Set("Content-Type", "application/json")
	requestFail0.Header.Set("X-API-KEY", apiKey)

	respFail0 := sendRequest(t, requestFail0)
	require.Equal(t, http.StatusRequestEntityTooLarge, respFail0.StatusCode)

	// unparsable body
	unknownBody := struct{ X int }{10}

	requestUnknown := prepareRequest(t, unknownBody)

	requestUnknown.Header.Set("Content-Type", "application/json")
	requestUnknown.Header.Set("X-API-KEY", apiKey)

	respUnknown := sendRequest(t, requestUnknown)
	require.Equal(t, http.StatusBadRequest, respUnknown.StatusCode)

	// unauthenticated
	requestUnauth := prepareRequest(t, body)

	requestUnauth.Header.Set("Content-Type", "application/json")

	respFailUnauth := sendRequest(t, requestUnauth)
	require.Equal(t, http.StatusUnauthorized, respFailUnauth.StatusCode)

	// missing "Content-Type"
	requestNoContentType := prepareRequest(t, body)
	requestNoContentType.Header.Set("X-API-KEY", apiKey)

	respFailNoContentType := sendRequest(t, requestNoContentType)
	require.Equal(t, http.StatusNotAcceptable, respFailNoContentType.StatusCode)

	// header too large
	requestLargeHeader := prepareRequest(t, body)
	requestLargeHeader.Header.Set("X-API-KEY", apiKey)
	requestLargeHeader.Header.Set("Content-Type", "application/json")
	for j := range uint64(1000) {
		requestLargeHeader.Header.Set(strconv.FormatUint(j, 10), strconv.FormatUint(j, 10))
	}

	respFailLargeHeader := sendRequest(t, requestLargeHeader)

	require.Equal(t, http.StatusRequestHeaderFieldsTooLarge, respFailLargeHeader.StatusCode)

	err = signer.Shutdown(ctx)
	require.NoError(t, err)
	cancel()

	wg.Wait()
}
