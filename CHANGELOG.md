# Changelog

## Unreleased

### Changed

- Logged durations are rounded.
  Console keeps milliseconds from a second up and microseconds below, `took=1.424s` instead of `took=1.423528535s`.
  `json` keeps whole microseconds, `"took":1423.529`.

## [v1.4.0](https://github.com/flare-foundation/go-flare-common/tree/v1.4.0) - 2026-10-13

### Structured logging on log/slog

`pkg/logger` is rebuilt on the standard library's `log/slog` and moves to structured logging.
zap and lumberjack are gone from the dependency tree.
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
2. Remove `SyncFileLogger` calls. The function is deprecated and does nothing.
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
