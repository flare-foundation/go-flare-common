# Changelog

## Unreleased

### Added

- Structured logging on the standard library's `log/slog`: `Debugw`, `Infow`, `Warnw`, `Errorw`, `Panicw`, `Fatalw` and `With`, on the package and on `*logger.Log`.
  `[logger] format` selects `console` (default) or `json`, and `level` takes DEBUG, INFO, WARN, ERROR or FATAL.
  `NO_COLOR` turns off console colour.

### Deprecated

- File logging. `file`, `max_file_size`, `max_backups`, `max_age_days` and `console` are accepted and ignored, and `SyncFileLogger` does nothing.
  Services write to standard output, and the container runtime handles retention and rotation.

### Changed

- **Breaking:** `Logger()` returns `*logger.Log` instead of `*zap.SugaredLogger`, and the zap and lumberjack dependencies are gone.
  Callers using zap methods must update, and `WithOptions(zap.AddCallerSkip(...))` is no longer needed for correct source locations.
- `Nop.Fatal`, `Nop.Fatalf` and `Nop.Fatalw` exit with 1 instead of panicking, as the real logger does.

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
