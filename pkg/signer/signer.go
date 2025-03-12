package signing

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

const (
	maxReqBodySize = 10 * (1 << 10) // 10 kB for maximal request body size (enough for roughly 145 signatures)
	maxHeaderBytes = (1 << 10)      // 1 kB for maximal request header (only expect )

	writeTimeout      = 3 * time.Second
	readTimeout       = 3 * time.Second
	readHeaderTimeout = 1 * time.Second
)

type Signer struct {
	*http.Server
}

// Run starts to listen and serve to requests.
func (s *Signer) Run(ctx context.Context) error {
	if s == nil {
		logger.Debug("no signer")
		return nil
	}

	c := make(chan error)

	go func() {
		err := s.ListenAndServe()
		c <- err
	}()

	var err error

	select {
	case <-ctx.Done():
		err = s.Shutdown(ctx)
		err = fmt.Errorf("server shut down:%v", err)
	case err = <-c:
	}

	return err
}

type Config struct {
	Addr       string   `toml:"addr"`
	APIKeyName string   `toml:"api_key_name"`
	APIKeys    []string `toml:"api_keys"`
}

type apiKeys struct {
	name string
	keys map[string]bool
}

// NewAPIKey builds apiKey from the config.
func newAPIKeys(cfg Config) apiKeys {
	keys := make(map[string]bool)
	for j := range cfg.APIKeys {
		keys[cfg.APIKeys[j]] = true
	}

	return apiKeys{
		name: cfg.APIKeyName,
		keys: keys,
	}
}

// Authorize checks whether the header has the correct authorization.
func (ak *apiKeys) authorize(h *http.Header) bool {
	providedKey := h.Get(ak.name)
	return ak.keys[providedKey]
}

// New creates new signer from config and private key.
//
// Signer listens to POST requests on cfg.Addr/sign.
// The body of the request should be json marshaled RequestBody, the header should include api key.
func New(cfg Config, prv *ecdsa.PrivateKey) *Signer {
	apiKey := newAPIKeys(cfg)

	handler := http.NewServeMux()

	server := http.Server{
		Addr:                         cfg.Addr,
		Handler:                      handler,
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  readTimeout,
		ReadHeaderTimeout:            readHeaderTimeout,
		WriteTimeout:                 writeTimeout,
		MaxHeaderBytes:               maxHeaderBytes,
	}

	handler.Handle("POST /sign", signHandler(prv, apiKey))

	return &Signer{&server}
}

// Request body format
type RequestBody struct {
	Hashes []common.Hash `json:"hashes"`
}

// Response body format
type ResponseBody struct {
	Signatures [][]byte `json:"signatures"`
}

// signHandler returns a HandlerFunction that authorizes requests with encoded RequestBody and responses with the signatures of the provided hashes.
func signHandler(prv *ecdsa.PrivateKey, ak apiKeys) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !ak.authorize(&r.Header) {
			errorString := "Unauthorized, provide valid " + ak.name + " api key"
			http.Error(w, errorString, http.StatusUnauthorized)
			return
		}

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusNotAcceptable)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		r.Body = http.MaxBytesReader(w, r.Body, maxReqBodySize)

		defer r.Body.Close()

		rb := RequestBody{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&rb)
		if err != nil {
			if err.Error() == "http: request body too large" {
				msg := "request to large"
				http.Error(w, msg, http.StatusRequestEntityTooLarge)
				return
			}

			msg := "unparsable request body"

			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		signatures, err := signHashes(rb.Hashes, prv)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response := ResponseBody{signatures}

		encoder := json.NewEncoder(w)

		err = encoder.Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// signEthMessage computes keccak256("\x19Ethereum Signed Message:\n"32{message}) and (ECDSA) signs it with private key.
func SignEthMessage(hash common.Hash, prv *ecdsa.PrivateKey) ([]byte, error) {
	toSign := accounts.TextHash(hash[:])
	return crypto.Sign(toSign, prv)
}

// signHashes creates Eth Message sign of each hash and returns it in a slice.
func signHashes(hashes []common.Hash, prv *ecdsa.PrivateKey) ([][]byte, error) {
	signatures := make([][]byte, len(hashes))

	for j := range hashes {
		signature, err := SignEthMessage(hashes[j], prv)
		if err != nil {
			return nil, err
		}

		signatures[j] = signature
	}

	return signatures, nil
}
