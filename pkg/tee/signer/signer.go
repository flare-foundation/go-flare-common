// Package signer provides an HTTP server for cryptographic signing and decryption operations
// using ECDSA keys, with API key-based authorization for each request.
package signer

import (
	"context"
	"crypto/ecdsa"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

const (
	defaultMaxReqBodySize        = int64(10 << 10)  // 10 KiB for maximal request body size (enough for roughly 145 signatures)
	defaultMaxReqBodySizeDecrypt = int64(500 << 10) // 500 KiB for maximal request body for decrypt
	defaultMaxHeaderBytes        = 1 << 10          // 1 KiB for maximal request header (only expect)

	defaultWriteTimeout      = 3 * time.Second
	defaultReadTimeout       = 3 * time.Second
	defaultReadHeaderTimeout = 1 * time.Second
)

// Config holds configuration for the Signer server.
type Config struct {
	Addr       string   `toml:"addr"`
	APIKeyName string   `toml:"api_key_name"`
	APIKeys    []string `toml:"api_keys"`
	Limits     Limits   `toml:"limits"`

	// Accept every request with no API-key check. Tests or isolated environments only.
	// Cannot be combined with APIKeyName or APIKeys.
	AllowUnauthenticated bool `toml:"allow_unauthenticated"`
}

type Limits struct {
	MaxReqBodySize        int64 `toml:"max_req_body_size"`
	MaxReqBodySizeDecrypt int64 `toml:"max_req_body_size_decrypt"`
	MaxHeaderBytes        int   `toml:"max_header_bytes"`

	WriteTimeout      time.Duration `toml:"write_timeout"`
	ReadTimeout       time.Duration `toml:"read_timeout"`
	ReadHeaderTimeout time.Duration `toml:"read_header_timeout"`
}

// setDefaults sets default values for any zero-valued fields in Limits.
func (l *Limits) setDefaults() {
	if l.MaxReqBodySize == 0 {
		l.MaxReqBodySize = defaultMaxReqBodySize
	}
	if l.MaxReqBodySizeDecrypt == 0 {
		l.MaxReqBodySizeDecrypt = defaultMaxReqBodySizeDecrypt
	}
	if l.MaxHeaderBytes == 0 {
		l.MaxHeaderBytes = defaultMaxHeaderBytes
	}
	if l.WriteTimeout == 0 {
		l.WriteTimeout = defaultWriteTimeout
	}
	if l.ReadTimeout == 0 {
		l.ReadTimeout = defaultReadTimeout
	}
	if l.ReadHeaderTimeout == 0 {
		l.ReadHeaderTimeout = defaultReadHeaderTimeout
	}
}

// Signer wraps an HTTP server that provides signing and decryption endpoints.
type Signer struct {
	*http.Server
}

// assertLoopbackAddr returns nil iff addr is a host:port whose host portion
// is a literal loopback IP. The check is intentionally strict: hostnames are
// rejected because their resolution can change between New() and Listen(); the
// safety invariant must hold at construction time, not at dial time.
func assertLoopbackAddr(addr string) error {
	if addr == "" {
		return errors.New("address must be specified; an unset address binds to all interfaces, breaking the signer's loopback-only invariant")
	}
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return fmt.Errorf("parsing %q: %w", addr, err)
	}
	if host == "" {
		return fmt.Errorf("address %q has empty host; signer must bind a loopback address explicitly (127.0.0.1 or ::1)", addr)
	}
	ip := net.ParseIP(host)
	if ip == nil {
		return fmt.Errorf("address host %q is not a literal IP; signer must bind a loopback IP (127.0.0.1 or ::1), not a hostname whose resolution can change", host)
	}
	if !ip.IsLoopback() {
		return fmt.Errorf("address host %q is not a loopback IP; signer must bind 127.0.0.1 or ::1", host)
	}
	return nil
}

// New creates a new Signer from the given config and ECDSA private key.
//
// The Signer listens to POST requests on /sign and /decrypt, and GET requests on /id.
// API key validation is performed for each request.
//
// cfg.Addr must be a host:port whose host portion resolves to a loopback
// address (127.0.0.0/8 or ::1). This enforces the loopback boundary: the
// signer is a local-only oracle, so a remote attacker cannot reach it.
// Empty or non-loopback Addr values are rejected at New() time, before any
// listener opens.
func New(cfg Config, prv *ecdsa.PrivateKey) (*Signer, error) {
	if prv == nil {
		// All handlers dereference prv (PublicKey for /id, crypto.Sign for
		// /sign, ECIES decrypt for /decrypt). A nil key passes New() only
		// to crash the listener on the first authenticated request — by
		// which point the network surface is already exposed.
		return nil, errors.New("signer private key is nil")
	}
	if err := assertLoopbackAddr(cfg.Addr); err != nil {
		return nil, fmt.Errorf("signer address: %w", err)
	}

	apiKeys, err := newAPIKeys(cfg)
	if err != nil {
		return nil, fmt.Errorf("initializing API keys: %w", err)
	}

	cfg.Limits.setDefaults()

	mux := http.NewServeMux()

	mux.Handle("POST /sign", signHandler(prv, cfg.Limits.MaxReqBodySize))
	mux.Handle("POST /decrypt", decryptHandler(prv, cfg.Limits.MaxReqBodySizeDecrypt))

	mux.Handle("GET /id", idHandler(&prv.PublicKey))

	handler := prepare(mux, apiKeys)

	server := http.Server{
		Addr:                         cfg.Addr,
		Handler:                      handler,
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  cfg.Limits.ReadTimeout,
		ReadHeaderTimeout:            cfg.Limits.ReadHeaderTimeout,
		WriteTimeout:                 cfg.Limits.WriteTimeout,
		MaxHeaderBytes:               cfg.Limits.MaxHeaderBytes,
	}

	return &Signer{Server: &server}, nil
}

// Run starts the HTTP server and listens for requests until the context is
// cancelled or ListenAndServe returns. The ListenAndServe result is propagated
// back to the caller (http.ErrServerClosed on graceful shutdown, or the real listener error).
// The result channel is buffered so the goroutine never blocks if Run already returned via
// ctx.Done; Shutdown uses a fresh bounded context so a cancelled parent does not no-op it.
func (s *Signer) Run(ctx context.Context) error {
	if s == nil {
		return errors.New("no signer")
	}

	c := make(chan error, 1)

	go func() {
		c <- s.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := s.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("server shut down: %w", err)
		}
		// Drain the ListenAndServe goroutine; it returns http.ErrServerClosed
		// after Shutdown completes. Propagate that as the Run result so the
		// caller sees the conventional sentinel.
		return <-c
	case err := <-c:
		return err
	}
}

// apiKeys hold API key digests for authorization of incoming requests.
//
// Each configured API key is stored as HMAC-SHA256(secret, key), where secret
// is a startup-random 32-byte value. On every request, the provided header
// value is HMAC'd under the same secret and compared with subtle.ConstantTimeCompare
// against each stored digest. This makes the comparison length- and content-
// independent in time and avoids storing raw keys in memory.
type apiKeys struct {
	name     string
	secret   []byte
	digests  [][]byte
	allowAll bool
}

// newAPIKeys builds an apiKeys struct from Config.
// AllowUnauthenticated yields an allow-all authorizer; otherwise APIKeyName and APIKeys must be non-empty.
func newAPIKeys(cfg Config) (apiKeys, error) {
	if cfg.AllowUnauthenticated {
		if cfg.APIKeyName != "" || len(cfg.APIKeys) > 0 {
			return apiKeys{}, errors.New("AllowUnauthenticated must not be combined with APIKeyName or APIKeys")
		}
		return apiKeys{allowAll: true}, nil
	}
	if cfg.APIKeyName == "" || len(cfg.APIKeys) == 0 {
		return apiKeys{}, errors.New("empty APIKeyName or APIKeys: set AllowUnauthenticated to run without authentication")
	}

	secret := make([]byte, 32)
	if _, err := rand.Read(secret); err != nil {
		return apiKeys{}, fmt.Errorf("generating HMAC secret: %w", err)
	}

	digests := make([][]byte, len(cfg.APIKeys))
	for j, k := range cfg.APIKeys {
		digests[j] = hmacKey(secret, k)
	}

	return apiKeys{
		name:    cfg.APIKeyName,
		secret:  secret,
		digests: digests,
	}, nil
}

// hmacKey returns HMAC-SHA256(secret, key).
func hmacKey(secret []byte, key string) []byte {
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(key))
	return h.Sum(nil)
}

// authorize checks whether the request header contains a valid API key.
// Comparison is constant-time and runs against every configured digest so
// match position does not leak through timing.
func (ak *apiKeys) authorize(h *http.Header) bool {
	if ak.allowAll {
		return true
	}
	provided := hmacKey(ak.secret, h.Get(ak.name))
	var ok byte
	for _, d := range ak.digests {
		ok |= byte(subtle.ConstantTimeCompare(provided, d))
	}
	return ok == 1
}

// SignBody is the request body format for the /sign endpoint.
type SignBody struct {
	Hashes []common.Hash `json:"hashes"`
}

// SignedBody is the response body format for the /sign endpoint.
type SignedBody struct {
	Signatures []hexutil.Bytes `json:"signatures"`
}

// signHandler returns an HTTP handler that signs hashes provided in the request body.
func signHandler(prv *ecdsa.PrivateKey, maxReqBodySize int64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusNotAcceptable)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, maxReqBodySize)

		defer r.Body.Close() //nolint:errcheck // closing request body, error not actionable

		rb := SignBody{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&rb)
		if err != nil {
			var maxErr *http.MaxBytesError
			if errors.As(err, &maxErr) {
				http.Error(w, "request too large", http.StatusRequestEntityTooLarge)
				return
			}

			msg := "unparsable request body"

			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		signatures, err := signHashes(rb.Hashes, prv)
		if err != nil {
			http.Error(w, "signing failed", http.StatusInternalServerError)
			return
		}
		response := SignedBody{Signatures: signatures}

		encoder := json.NewEncoder(w)

		if err := encoder.Encode(response); err != nil {
			http.Error(w, "encoding response", http.StatusInternalServerError)
			return
		}
	}
}

// SignEthMessage computes the Ethereum Signed Message hash and signs it with the given ECDSA private key.
func SignEthMessage(hash common.Hash, prv *ecdsa.PrivateKey) (hexutil.Bytes, error) {
	toSign := accounts.TextHash(hash[:])
	return crypto.Sign(toSign, prv)
}

// signHashes creates an Ethereum Signed Message signature for each hash and returns them in a slice.
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

// idHandler returns an HTTP handler that responds with the coordinates of the signer's public key.
func idHandler(pubKey *ecdsa.PublicKey) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pk := pubKeyToStruct(pubKey)
		encoder := json.NewEncoder(w)

		if err := encoder.Encode(pk); err != nil {
			http.Error(w, "encoding response", http.StatusInternalServerError)
			return
		}
	}
}

// prepare wraps the handler with API key authorization and sets Content-Type to "application/json".
func prepare(h http.Handler, ak apiKeys) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !ak.authorize(&r.Header) {
			msg := fmt.Sprintf("unauthorized: unrecognized %s", ak.name)
			http.Error(w, msg, http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}

// PublicKey holds the X and Y coordinates of an ECDSA public key.
type PublicKey struct {
	X common.Hash `json:"x"`
	Y common.Hash `json:"y"`
}

// pubKeyToStruct converts an ECDSA public key to a PublicKey struct.
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

// EncryptedBody is the request body format for the /decrypt endpoint.
type EncryptedBody struct {
	Cipher hexutil.Bytes `json:"cipher"`
}

// DecryptedBody is the response body format for the /decrypt endpoint.
type DecryptedBody struct {
	Plain hexutil.Bytes `json:"plain"`
}

// decryptHandler returns an HTTP handler that decrypts the cipher from the request body using the provided ECDSA private key.
func decryptHandler(prv *ecdsa.PrivateKey, maxReqBodySize int64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusNotAcceptable)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, maxReqBodySize)

		defer r.Body.Close() //nolint:errcheck // closing request body, error not actionable

		eb := EncryptedBody{}
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		err := decoder.Decode(&eb)
		if err != nil {
			var maxErr *http.MaxBytesError
			if errors.As(err, &maxErr) {
				http.Error(w, "request too large", http.StatusRequestEntityTooLarge)
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
			return
		}

		response := DecryptedBody{
			Plain: plain,
		}

		encoder := json.NewEncoder(w)

		err = encoder.Encode(response)
		if err != nil {
			http.Error(w, "encoding response", http.StatusInternalServerError)
			return
		}
	}
}

// ECDSAPubKeyToECIES converts an ECDSA public key on secp256k1 to an ECIES public key.
func ECDSAPubKeyToECIES(pubKey *ecdsa.PublicKey) (*ecies.PublicKey, error) {
	if pubKey.Curve != secp256k1.S256() && pubKey.Curve != crypto.S256() {
		return nil, errors.New("curve not S256")
	}

	return &ecies.PublicKey{X: pubKey.X, Y: pubKey.Y, Curve: ecies.DefaultCurve, Params: ecies.ECIES_AES128_SHA256}, nil
}

// ECDSAPrivKeyToECIES converts an ECDSA private key on secp256k1 to an ECIES private key.
func ECDSAPrivKeyToECIES(privKey *ecdsa.PrivateKey) (*ecies.PrivateKey, error) {
	pubKey, err := ECDSAPubKeyToECIES(&privKey.PublicKey)
	if err != nil {
		return nil, err
	}

	return &ecies.PrivateKey{PublicKey: *pubKey, D: privKey.D}, nil
}

// decrypt decrypts the given cipher text using the provided ECDSA private key and returns the plain text.
func decrypt(cipher []byte, privKeyECDSA *ecdsa.PrivateKey) (hexutil.Bytes, error) {
	privKeyDecryption, err := ECDSAPrivKeyToECIES(privKeyECDSA)
	if err != nil {
		return nil, err
	}
	plainText, err := privKeyDecryption.Decrypt(cipher, nil, nil)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}
