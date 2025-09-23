package signer

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
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

// New creates new Signer from config and private key.
//
// Signer listens to POST requests on /id, /sign, and /decrypt endpoints.
// API key validation is done for each request.
func New(cfg Config, prv *ecdsa.PrivateKey) *Signer {
	apiKeys := newAPIKeys(cfg)

	mux := http.NewServeMux()

	mux.Handle("POST /sign", signHandler(prv))
	mux.Handle("POST /decrypt", decryptHandler(prv))

	mux.Handle("GET /id", idHandler(&prv.PublicKey))

	handler := prepare(mux, apiKeys)

	server := http.Server{
		Addr:                         cfg.Addr,
		Handler:                      handler,
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  readTimeout,
		ReadHeaderTimeout:            readHeaderTimeout,
		WriteTimeout:                 writeTimeout,
		MaxHeaderBytes:               maxHeaderBytes,
	}

	return &Signer{&server}
}

// SignBody is request body format for /sign.
type SignBody struct {
	Hashes []common.Hash `json:"hashes"`
}

// SignedBody is response body for /sign.
type SignedBody struct {
	Signatures []hexutil.Bytes `json:"signatures"`
}

// signHandler returns a handler function that signs hashes in the request SignBody.
func signHandler(prv *ecdsa.PrivateKey) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusNotAcceptable)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, maxReqBodySize)

		defer r.Body.Close() //nolint:errcheck

		rb := SignBody{}
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
		response := SignedBody{Signatures: signatures}

		encoder := json.NewEncoder(w)

		err = encoder.Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// signEthMessage computes keccak256("\x19Ethereum Signed Message:\n"32{message}) and (ECDSA) signs it with private key.
func SignEthMessage(hash common.Hash, prv *ecdsa.PrivateKey) (hexutil.Bytes, error) {
	toSign := accounts.TextHash(hash[:])
	return crypto.Sign(toSign, prv)
}

// signHashes creates Eth Message signature of each hash and returns it in a slice.
func signHashes(hashes []common.Hash, prv *ecdsa.PrivateKey) ([]hexutil.Bytes, error) {
	signatures := make([]hexutil.Bytes, len(hashes))

	for j := range hashes {
		signature, err := SignEthMessage(hashes[j], prv)
		if err != nil {
			return nil, err
		}

		signatures[j] = signature
	}

	return signatures, nil
}

// idHandler returns a handler function that responds with coordinates of the signer's public key.
func idHandler(pubKey *ecdsa.PublicKey) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pk := pubKeyToStruct(pubKey)
		encoder := json.NewEncoder(w)

		err := encoder.Encode(pk)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// prepare adds api key check and sets Content-Type to "application/json".
func prepare(h http.Handler, ak apiKeys) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !ak.authorize(&r.Header) {
			msg := fmt.Sprintf("unauthorized: unrecognized %s", ak.name)
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

type PublicKey struct {
	X common.Hash `json:"x"`
	Y common.Hash `json:"y"`
}

func pubKeyToStruct(key *ecdsa.PublicKey) PublicKey {
	var newKey PublicKey
	xBytes := key.X.Bytes()
	yBytes := key.Y.Bytes()

	if len(xBytes) < 32 {
		xBytes = append(make([]byte, 32-len(xBytes)), xBytes...)
	}
	if len(yBytes) < 32 {
		yBytes = append(make([]byte, 32-len(yBytes)), yBytes...)
	}

	copy(newKey.X[:], xBytes)
	copy(newKey.Y[:], yBytes)

	return newKey
}

// EncryptedBody is request body format for /decrypt.
type EncryptedBody struct {
	Cipher hexutil.Bytes `json:"cipher"`
}

// DecryptedBody is response body for /decrypt.
type DecryptedBody struct {
	Plain hexutil.Bytes `json:"plain"`
}

// decryptHandler returns a handler function that decrypts cipher from EncryptedBody.
func decryptHandler(prv *ecdsa.PrivateKey) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, maxReqBodySize)

		defer r.Body.Close() //nolint:errcheck

		eb := EncryptedBody{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&eb)
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

		plain, err := decrypt(eb.Cipher, prv)
		if err != nil {
			msg := "cannot decrypt"
			http.Error(w, msg, http.StatusBadRequest)
		}

		response := DecryptedBody{
			Plain: plain,
		}

		encoder := json.NewEncoder(w)

		err = encoder.Encode(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func decrypt(cipher []byte, privKeyECDSA *ecdsa.PrivateKey) (hexutil.Bytes, error) {
	privKeyDecryption := ecies.ImportECDSA(privKeyECDSA)
	plainText, err := privKeyDecryption.Decrypt(cipher, nil, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
