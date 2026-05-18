package signer

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/stretchr/testify/require"
)

const (
	port    = 6739
	apiKey  = "krompirusa"
	prvSeed = "nekejCesarNihceNeUgane"

	signEP     = "/sign"
	identityEP = "/id"
	decryptEP  = "/decrypt"
)

var url = fmt.Sprintf("http://localhost:%d", port)

func prepareServer(t *testing.T) (*Signer, *ecdsa.PrivateKey) {
	t.Helper()

	randomStr := []byte(prvSeed)
	pkStr := crypto.Keccak256(randomStr)
	prv, err := crypto.ToECDSA(pkStr)
	require.NoError(t, err)

	cfg := Config{
		Addr:       fmt.Sprintf("127.0.0.1:%d", port),
		APIKeyName: "X-API-KEY",
		APIKeys:    []string{apiKey},
	}

	s, err := New(cfg, prv)
	require.NoError(t, err)
	return s, prv
}

func prepareSignBody(n uint64) SignBody {
	hashes := make([]common.Hash, n)

	for j := range n {
		hashes[j] = crypto.Keccak256Hash([]byte(strconv.FormatUint(j, 10)))
	}

	return SignBody{
		Hashes: hashes,
	}
}

func preparePOSTRequest(t *testing.T, body any, endpoint string) *http.Request {
	t.Helper()

	encodedBody, err := json.Marshal(body)
	require.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, url+endpoint, bytes.NewBuffer(encodedBody))
	require.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	return req
}

func sendRequest(t *testing.T, r *http.Request) *http.Response {
	t.Helper()

	client := &http.Client{}
	resp, err := client.Do(r)
	require.NoError(t, err)
	return resp
}

func TestSigner(t *testing.T) {
	signer, prv := prepareServer(t)

	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	var wg2 sync.WaitGroup
	wg2.Add(1)

	wg.Go(func() {
		wg2.Done()
		err := signer.Run(ctx)
		require.Error(t, err)
	})

	wg2.Wait()
	time.Sleep(100 * time.Millisecond)

	t.Run("sign happy path", func(t *testing.T) {
		t.Parallel()

		const nuOfHashes uint64 = 10

		body := prepareSignBody(nuOfHashes)

		request := preparePOSTRequest(t, body, signEP)

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("X-API-KEY", apiKey)

		resp := sendRequest(t, request)

		defer resp.Body.Close() //nolint:errcheck

		require.Equal(t, http.StatusOK, resp.StatusCode)

		require.Equal(t, resp.Header.Get("Content-Type"), "application/json")

		decoder := json.NewDecoder(resp.Body)
		responseBody := SignedBody{}
		err := decoder.Decode(&responseBody)
		require.NoError(t, err)
		require.Len(t, responseBody.Signatures, int(nuOfHashes))

		recovered, err := crypto.SigToPub(accounts.TextHash(body.Hashes[0][:]), responseBody.Signatures[0])
		require.NoError(t, err)

		require.Equal(t, prv.PublicKey, *recovered)
	})

	t.Run("sign to big request", func(t *testing.T) {
		t.Parallel()

		bodyToBig := prepareSignBody(2000)
		requestFail0 := preparePOSTRequest(t, bodyToBig, signEP)

		requestFail0.Header.Set("Content-Type", "application/json")
		requestFail0.Header.Set("X-API-KEY", apiKey)

		respFail0 := sendRequest(t, requestFail0)
		defer respFail0.Body.Close() //nolint:errcheck

		require.Equal(t, http.StatusRequestEntityTooLarge, respFail0.StatusCode)
	})

	t.Run("sign unparsable", func(t *testing.T) {
		t.Parallel()

		unknownBody := struct{ X int }{10}

		requestUnknown := preparePOSTRequest(t, unknownBody, signEP)

		requestUnknown.Header.Set("Content-Type", "application/json")
		requestUnknown.Header.Set("X-API-KEY", apiKey)

		respUnknown := sendRequest(t, requestUnknown)

		require.Equal(t, http.StatusBadRequest, respUnknown.StatusCode)
		err := respUnknown.Body.Close()
		require.NoError(t, err)
	})

	t.Run("sign unauthorized", func(t *testing.T) {
		t.Parallel()

		body := prepareSignBody(1)

		requestUnauth := preparePOSTRequest(t, body, signEP)

		requestUnauth.Header.Set("Content-Type", "application/json")

		respFailUnauth := sendRequest(t, requestUnauth)
		require.Equal(t, http.StatusUnauthorized, respFailUnauth.StatusCode)

		err := respFailUnauth.Body.Close()
		require.NoError(t, err)
	})

	t.Run("sign missing Content-Type", func(t *testing.T) {
		t.Parallel()

		body := prepareSignBody(1)

		requestNoContentType := preparePOSTRequest(t, body, signEP)
		requestNoContentType.Header.Set("X-API-KEY", apiKey)
		requestNoContentType.Header.Del("Content-Type")

		respFailNoContentType := sendRequest(t, requestNoContentType)
		require.Equal(t, http.StatusNotAcceptable, respFailNoContentType.StatusCode)

		err := respFailNoContentType.Body.Close()
		require.NoError(t, err)
	})

	t.Run("sign header too large", func(t *testing.T) {
		t.Parallel()
		body := prepareSignBody(1)

		requestLargeHeader := preparePOSTRequest(t, body, signEP)
		requestLargeHeader.Header.Set("X-API-KEY", apiKey)
		requestLargeHeader.Header.Set("Content-Type", "application/json")
		for j := range uint64(1000) {
			requestLargeHeader.Header.Set(strconv.FormatUint(j, 10), strconv.FormatUint(j, 10))
		}

		respFailLargeHeader := sendRequest(t, requestLargeHeader)

		require.Equal(t, http.StatusRequestHeaderFieldsTooLarge, respFailLargeHeader.StatusCode)

		respFailLargeHeader.Body.Close() //nolint:errcheck
	})

	t.Run("identity happy path", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequest(http.MethodGet, url+identityEP, nil)
		require.NoError(t, err)

		req.Header.Set("X-API-KEY", apiKey)
		req.Header.Set("Content-Type", "application/json")

		resp := sendRequest(t, req)
		defer resp.Body.Close() //nolint:errcheck

		require.Equal(t, http.StatusOK, resp.StatusCode)
		decoder := json.NewDecoder(resp.Body)

		pk := PublicKey{}
		err = decoder.Decode(&pk)
		require.NoError(t, err)

		expectedPK := pubKeyToStruct(&prv.PublicKey)

		require.Equal(t, expectedPK, pk)
	})

	t.Run("identity unauthorized", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequest(http.MethodGet, url+identityEP, nil)
		require.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp := sendRequest(t, req)

		require.Equal(t, http.StatusUnauthorized, resp.StatusCode)

		err = resp.Body.Close()
		require.NoError(t, err)
	})

	t.Run("decrypt happy path", func(t *testing.T) {
		t.Parallel()

		plaintext := []byte("plaintextThatIsShort")

		pke := ecies.ImportECDSAPublic(&prv.PublicKey)

		cipher, err := ecies.Encrypt(rand.Reader, pke, plaintext, nil, nil)
		require.NoError(t, err)

		eb := EncryptedBody{
			Cipher: cipher,
		}

		req := preparePOSTRequest(t, eb, decryptEP)

		req.Header.Set("X-API-KEY", apiKey)

		resp := sendRequest(t, req)
		defer resp.Body.Close() //nolint:errcheck

		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.Equal(t, resp.Header.Get("Content-Type"), "application/json")

		responseBody := DecryptedBody{}

		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&responseBody)
		require.NoError(t, err)

		require.Equal(t, hexutil.Bytes(plaintext), responseBody.Plain)
	})

	t.Run("decrypt unparsable", func(t *testing.T) {
		t.Parallel()

		unknownBody := struct{ X int }{10}

		requestUnknown := preparePOSTRequest(t, unknownBody, decryptEP)

		requestUnknown.Header.Set("Content-Type", "application/json")
		requestUnknown.Header.Set("X-API-KEY", apiKey)

		respUnknown := sendRequest(t, requestUnknown)
		require.Equal(t, http.StatusBadRequest, respUnknown.StatusCode)

		err := respUnknown.Body.Close()
		require.NoError(t, err)
	})

	t.Run("invalid endpoint", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequest(http.MethodGet, url+identityEP+signEP, nil)
		require.NoError(t, err)

		req.Header.Set("X-API-KEY", apiKey)
		req.Header.Set("Content-Type", "application/json")

		resp := sendRequest(t, req)

		require.Equal(t, http.StatusNotFound, resp.StatusCode)

		err = resp.Body.Close()
		require.NoError(t, err)
	})

	t.Cleanup(func() {
		err := signer.Shutdown(ctx)
		require.NoError(t, err)
		cancel()

		wg.Wait()
	})
}

func TestConfigs(t *testing.T) {
	lts := Limits{
		maxReqBodySize:        200,
		maxReqBodySizeDecrypt: 500,
		maxHeaderBytes:        1 << 10,
		writeTimeout:          1 * time.Second,
		readTimeout:           1 * time.Second,
		readHeaderTimeout:     1 * time.Second,
	}

	t.Run("limits", func(t *testing.T) {
		// nothing changed
		lts.setDefaults()
		require.Equal(t, int64(200), lts.maxReqBodySize)
		require.Equal(t, int64(500), lts.maxReqBodySizeDecrypt)
		require.Equal(t, 1<<10, lts.maxHeaderBytes)
		require.Equal(t, 1*time.Second, lts.writeTimeout)
		require.Equal(t, 1*time.Second, lts.readTimeout)
		require.Equal(t, 1*time.Second, lts.readHeaderTimeout)

		ltsDef := Limits{}
		ltsDef.setDefaults()
		require.Equal(t, defaultMaxReqBodySize, ltsDef.maxReqBodySize)
		require.Equal(t, defaultMaxReqBodySizeDecrypt, ltsDef.maxReqBodySizeDecrypt)
		require.Equal(t, defaultMaxHeaderBytes, ltsDef.maxHeaderBytes)
		require.Equal(t, defaultWriteTimeout, ltsDef.writeTimeout)
		require.Equal(t, defaultReadTimeout, ltsDef.readTimeout)
		require.Equal(t, defaultReadHeaderTimeout, ltsDef.readHeaderTimeout)
	})

	cfg := Config{
		Addr:       fmt.Sprintf("127.0.0.1:%d", port),
		APIKeyName: "X-API-KEY",
		APIKeys:    []string{apiKey},
		Limits:     lts,
	}

	ctx, cancel := context.WithCancel(context.Background())

	prv, err := crypto.GenerateKey()
	require.NoError(t, err)

	s, err := New(cfg, prv)
	require.NoError(t, err)

	var wg sync.WaitGroup

	wg.Go(func() {
		err := s.Run(ctx)
		require.Error(t, err)
	})

	time.Sleep(1 * time.Second)

	t.Run("sign with custom limit", func(t *testing.T) {
		const nuOfHashesFail uint64 = 3

		body := prepareSignBody(nuOfHashesFail)
		request := preparePOSTRequest(t, body, signEP)

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("X-API-KEY", apiKey)

		resp := sendRequest(t, request)
		require.Equal(t, http.StatusRequestEntityTooLarge, resp.StatusCode)

		err := resp.Body.Close()
		require.NoError(t, err)

		const nuOfHashesSuccess uint64 = 1

		body = prepareSignBody(nuOfHashesSuccess)
		request = preparePOSTRequest(t, body, signEP)

		request.Header.Set("Content-Type", "application/json")
		request.Header.Set("X-API-KEY", apiKey)

		resp = sendRequest(t, request)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		err = resp.Body.Close()
		require.NoError(t, err)
	})
	t.Run("decrypt with custom limit ok", func(t *testing.T) {
		textLength := 10
		plaintext := bytes.Repeat([]byte("a"), textLength)

		pke := ecies.ImportECDSAPublic(&prv.PublicKey)

		cipher, err := ecies.Encrypt(rand.Reader, pke, plaintext, nil, nil)
		require.NoError(t, err)

		eb := EncryptedBody{
			Cipher: cipher,
		}

		req := preparePOSTRequest(t, eb, decryptEP)

		req.Header.Set("X-API-KEY", apiKey)

		resp := sendRequest(t, req)
		require.Equal(t, http.StatusOK, resp.StatusCode)

		err = resp.Body.Close()
		require.NoError(t, err)
	})

	t.Run("decrypt with custom limit fail", func(t *testing.T) {
		textLengthOk := 150
		plaintext := bytes.Repeat([]byte("a"), textLengthOk)

		pke := ecies.ImportECDSAPublic(&prv.PublicKey)

		cipher, err := ecies.Encrypt(rand.Reader, pke, plaintext, nil, nil)
		require.NoError(t, err)

		eb := EncryptedBody{
			Cipher: cipher,
		}

		req := preparePOSTRequest(t, eb, decryptEP)
		req.Header.Set("X-API-KEY", apiKey)

		resp := sendRequest(t, req)

		require.Equal(t, http.StatusRequestEntityTooLarge, resp.StatusCode)

		err = resp.Body.Close()
		require.NoError(t, err)
	})

	t.Cleanup(func() {
		err := s.Shutdown(ctx)
		require.NoError(t, err)
		cancel()
		wg.Wait()
	})
}

// TestAPIKeyHMACCompare covers audit finding H21: API-key comparison runs
// against HMAC digests with subtle.ConstantTimeCompare. The test exercises
// the behavioural contract — correct keys accepted, wrong/empty keys
// rejected, multiple configured keys all accepted — independent of timing.
func TestAPIKeyHMACCompare(t *testing.T) {
	cfg := Config{
		APIKeyName: "X-API-KEY",
		APIKeys:    []string{"alpha", "beta", "gamma"},
	}
	ak, err := newAPIKeys(cfg)
	require.NoError(t, err)
	require.Len(t, ak.digests, 3)
	require.Len(t, ak.secret, 32)

	mk := func(v string) *http.Header {
		h := http.Header{}
		if v != "" {
			h.Set(cfg.APIKeyName, v)
		}
		return &h
	}

	for _, valid := range cfg.APIKeys {
		require.Truef(t, ak.authorize(mk(valid)), "valid key %q should authorize", valid)
	}
	require.False(t, ak.authorize(mk("")))
	require.False(t, ak.authorize(mk("delta")))
	require.False(t, ak.authorize(mk("alpha1")))
	require.False(t, ak.authorize(mk("alph")))
}

// TestAssertLoopbackAddr covers audit finding H22: New() must reject any
// configuration that would bind the signer outside loopback, because the
// deferred C4/C5/C6 unbound-oracle findings rely on the signer being
// reachable only from inside the host trust boundary.
func TestAssertLoopbackAddr(t *testing.T) {
	good := []string{
		"127.0.0.1:0",
		"127.0.0.1:8080",
		"127.1.2.3:8080",
		"[::1]:8080",
	}
	for _, a := range good {
		t.Run("accept "+a, func(t *testing.T) {
			require.NoError(t, assertLoopbackAddr(a))
		})
	}

	bad := []struct {
		addr string
		want string
	}{
		{"", "must be specified"},
		{":8080", "empty host"},
		{"0.0.0.0:8080", "not a loopback"},
		{"192.168.1.5:8080", "not a loopback"},
		{"10.0.0.1:8080", "not a loopback"},
		{"localhost:8080", "not a literal IP"},
		{"example.com:8080", "not a literal IP"},
		{"127.0.0.1", "parsing"},
		{"::1:8080", "parsing"},
	}
	for _, tc := range bad {
		t.Run("reject "+tc.addr, func(t *testing.T) {
			err := assertLoopbackAddr(tc.addr)
			require.Error(t, err)
			require.Contains(t, err.Error(), tc.want)
		})
	}
}

func TestNewRejectsNonLoopbackAddr(t *testing.T) {
	prv, err := crypto.GenerateKey()
	require.NoError(t, err)

	cfg := Config{
		Addr:       ":8080",
		APIKeyName: "X-API-KEY",
		APIKeys:    []string{"k"},
	}
	s, err := New(cfg, prv)
	require.Error(t, err)
	require.Nil(t, s)
	require.Contains(t, err.Error(), "signer address")
}

// TestNewRejectsNilPrivateKey covers audit finding F-TEE-2: a nil
// *ecdsa.PrivateKey used to pass New() and surface as a deref panic on
// the first /sign or /id request — exposing the listener before the
// crash. New must reject up front.
func TestNewRejectsNilPrivateKey(t *testing.T) {
	cfg := Config{
		Addr:       "127.0.0.1:0",
		APIKeyName: "X-API-KEY",
		APIKeys:    []string{"k"},
	}
	s, err := New(cfg, nil)
	require.Error(t, err)
	require.Nil(t, s)
	require.Contains(t, err.Error(), "private key is nil")
}

func TestAPIKeyHMACSecretsAreDistinct(t *testing.T) {
	cfg := Config{APIKeyName: "X-API-KEY", APIKeys: []string{"k"}}
	a, err := newAPIKeys(cfg)
	require.NoError(t, err)
	b, err := newAPIKeys(cfg)
	require.NoError(t, err)
	require.NotEqual(t, a.secret, b.secret, "each New() must use a fresh HMAC secret")
	require.NotEqual(t, a.digests[0], b.digests[0], "digest under different secret must differ")
}
