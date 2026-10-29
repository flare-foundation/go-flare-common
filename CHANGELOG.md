# Changelog

## Unreleased

### Added

- `contracts.MustABI`, `MustEvent`, `MustMethod` and `MustError` look up a binding's ABI and its events, methods and custom errors by name for package-level variables and init functions, panicking when one is missing.
  `MustABI` takes the metadata of an abigen v1 or v2 binding.
- `abicoder.Coder` encodes and decodes a Go type as one ABI argument: `abicoder.New` and `MustNew` check once that the type fits the argument, so `Encode` and `Decode` check only the data.
  Each struct field must be named for the tuple component at its position, as abigen's structs are, and `New` refuses an indexed (event) argument.
- `abicoder.ArgumentsDecoder` decodes an argument list as `DecodeArguments` does: `abicoder.NewArgumentsDecoder` and `MustNewArgumentsDecoder` check once that the type fits the arguments, so `Decode` and `DecodeTo` check only the data.
  `NewArgumentsDecoder` refuses an indexed (event) argument, and a type with no place for an argument with `ErrShapeMismatch`.
- `Coder.Decode` and `ArgumentsDecoder.Decode` don't re-encode the decoded values to compare them with the input, as `Decode` and `DecodeArguments` do: with indexed arguments refused, the layout check alone refuses the same inputs.

### Changed

- **Breaking:** requires Go 1.27.3.
- **Behavioral:** `abicoder`'s decoders refuse a destination holding a pointer to anything but a `big.Int`, such as a `*uint64` field, with `ErrShapeMismatch`. go-ethereum decodes through such a pointer without allocating it, so decoding panicked on every input, or for a slice element on every non-empty slice.
- **Behavioral:** `abicoder`'s decoders refuse a destination type go-ethereum can't assign the decoded value to with `ErrShapeMismatch`, such as a `string` or `uint64` for a `uint256`, a `type U uint64` for a `uint64`, a `big.Int` field, an unexported field, or an interface the value doesn't implement.
  Decoding into one failed on every input with a recovered panic, or for a slice element on every non-empty slice.
- `abicoder` refuses an `abi.Type` built by hand with go-ethereum's unused `HashTy` as unsupported.
  No ABI parses to it and go-ethereum can't encode it, so decoding it failed in the round trip after decoding.

### Deprecated

- `structs.Decode` and `structs.DecodeTo`: use `abicoder.Decode` and `abicoder.DecodeTo`, or an `abicoder.Coder` from `abicoder.New` to decode one type repeatedly.

## [v1.5.0](https://github.com/flare-foundation/go-flare-common/tree/v1.5.0) - 2026-10-16

### Added

- `abicoder.DecodeArguments` and `DecodeArgumentsTo` decode a top-level argument list, as Solidity's `abi.encode(a, b, ...)` writes it, checking the data as `Decode` does and the field or element each argument is copied into as `Decode` checks its destination.
- `abicoder.ErrTruncated`.
- `panics.Recover` turns a panic into an error when deferred as `defer panics.Recover(&err)`.

### Changed

- **Behavioral:** `abicoder.Decode` and `DecodeTo` check the encoding's layout before decoding: data whose offsets aren't where the encoder puts them fails with `ErrNonCanonical`, and data too short with `ErrTruncated` instead of go-ethereum's error.
- **Behavioral:** `abicoder` refuses integers out of range for widths go-ethereum doesn't range-check, such as `uint128`, and refuses zero-length fixed arrays, empty tuples and arrays too large to encode as unsupported, whatever the data.
- **Behavioral:** `abicoder.Encode` refuses the types the decoders refuse, and values that don't fit their type, such as an `int256` at or above 2^255, which go-ethereum wraps modulo 2^256, an `address` or `function` from more than 32 bytes, which go-ethereum writes whole, shifting what follows, or an `address` from 32 bytes with the high ones set, with `ErrNonCanonical`.
- **Behavioral:** `abicoder`'s decoders refuse a Go array for any ABI type but a fixed-size array, `bytesN`, `address` or `function` with `ErrShapeMismatch`, instead of truncating or zero-filling a dynamic array, `bytes` or `string` into it.
- **Behavioral:** `abicoder`'s decoders refuse a struct whose field name or `abi` tag matches a tuple component at another position with `ErrShapeMismatch`: tuples still decode by position while `Encode` matches by name, so such a struct got swapped values.
- `abicoder.Encode` refuses a value of a recursive pointer type, such as `type P *P`, with an error instead of crashing in go-ethereum with a stack overflow.
- `abicoder`'s round-trip error no longer prints the input.

### Security

- `abicoder`: offsets aliasing one region no longer make decoding allocate per alias. v1.3.0 and v1.4.0 guarded `string` only, so `bytes[]` and nested arrays could still turn about 150 KB into gigabytes.
- `abicoder`: decoding an integer go-ethereum holds in a `*big.Int`, such as a `uint256`, into a Go array no longer crashes the process with a stack overflow that `recover` can't catch.

## [v1.4.0](https://github.com/flare-foundation/go-flare-common/tree/v1.4.0) - 2026-10-13

### Structured logging on log/slog

`pkg/logger` is rebuilt on the standard library's `log/slog` and moves to structured logging. zap and lumberjack are gone from the dependency tree.
**Every service that uses `pkg/logger` needs a code or config change to upgrade.**

What changes:

- Structured functions `Debugw`, `Infow`, `Warnw`, `Errorw`, `Panicw` and `Fatalw` take a fixed message and key-value fields, and `With` attaches fields to every line of a logger.
  The printf (`Infof`) and print (`Info`) functions are kept for compatibility.
- Logs go to standard output only.
  `[logger] format` selects `console` (default) or `json`.
- The console timestamp changes from `[01-02|15:04:05.000]` in local time to RFC 3339 in UTC with milliseconds, and fields follow the message as `key=value`.
  `NO_COLOR` turns off colour.
  Anything that parses the old console output needs updating.
- `json` writes one object per line with `time`, a lowercase `level`, `source`, `msg` and the fields as top-level keys.
- ERROR lines no longer carry a stack trace.
- `level` takes DEBUG, INFO, WARN, ERROR or FATAL, and the old DPANIC and PANIC are read as FATAL.
  `Panic`, `Panicf` and `Panicw` log at FATAL and then panic.
- An unknown level or format keeps the previous logger and logs an ERROR, instead of falling back to DEBUG.

To upgrade a service:

1. `Logger()` returns `*logger.Log` instead of `*zap.SugaredLogger`.
   Replace `*zap.SugaredLogger` in types and signatures with `*logger.Log` or a small interface, and drop `WithOptions(zap.AddCallerSkip(...))`, which source locations no longer need.
2. Remove `SyncFileLogger` calls.
   The function is deprecated and does nothing.
3. Remove `file`, `max_file_size`, `max_backups`, `max_age_days` and `console` from `[logger]`.
   They are ignored and named in one INFO line at startup.
   Set retention and rotation on the container runtime's logging driver instead.
4. Set `format = "json"` where a log shipper reads the logs.
5. Tests that rely on `Nop.Fatal`, `Nop.Fatalf` or `Nop.Fatalw` panicking must change, because they now exit with 1 as the real logger does.

New and changed log lines should use the structured functions, with identifiers in fields rather than in the message.

### Added

- Contract bindings generated with `abigen --v2` in `contracts/v2/...`, mirroring the v1 bindings in `contracts/...`.

## [v1.3.0](https://github.com/flare-foundation/go-flare-common/tree/v1.3.0) - 2026-09-30

### Added

- Packages `abicoder`, `call`, `convert`, `encoding`, `random`, `retry`, `safe`, `safeurl`, `signing`, `toml`, `tee/...` and `xrpl/...`.
- Contract bindings for ContractRegistry, EntityManager, FDC2 and TEE contracts.
- `database`: opt-in MySQL TLS.
- `policy`: source-bound signing-policy hash.

### Changed

- **Breaking:** requires Go 1.25.14 and go-ethereum v1.17.
- **Breaking:** `voters.NewSet`, `voters.InitialHashSeed`, `payload.BuildMessage`, `merkle.BuildFromHex`, `policy.NewSigningPolicy` and `database.WaitCIndexerToSync` return an error.
  `WaitCIndexerToSync` also takes a logger.
- **Breaking:** `policy.SigningPolicy.RewardEpochID` is `uint32`, `RawBytes` is a method, and `policy.Storage` no longer exposes its mutex.
- **Breaking:** `priority.New` returns a pointer; `Add` and `AddFast` take a context and return an error.
- **Breaking:** `heapt.Remove` returns `(T, bool)`.
- **Breaking:** `merkle.Tree` methods `GetLeaf`, `GetProof` and `GetProofFromHash` are renamed `Leaf`, `Proof` and `ProofFromHash`.
- **Breaking:** `contracts/registry` and `contracts/preregistry` are regenerated from the current ABIs, and most of their methods are gone.
- **Behavioral:** `database` queries ordered by timestamp break ties by `id`, so rows with equal timestamps come back in a deterministic order.

### Removed

- `restserver` package.
- `logger.GetLogger`; use `logger.Logger`.
- `merkle.NewFromHex` and `queue.NotRatedDequeue`.

### Security

- Raised minimum `golang.org/x/crypto`, `golang.org/x/text`, `filippo.io/edwards25519` and `go.opentelemetry.io/otel` to releases fixing the reachable advisories.
