# Audit Summary — `go-flare-common`

Two-pass security audit on branch `fullAudit`.
Range: `0410a58` (initial audit report) → `e3613ff` (HEAD). 77 commits across three rounds.

See `AUDIT_REPORT.md` (pass 1) and `AUDIT_REPORT_PASS2.md` (pass 2 + Low/Info appendix) for per-finding detail and closure tables.

## Pass 1 (~2026-05-06 → 2026-05-13)

Closed 1 Critical + all Highs + 28 Mediums + scoped Lows from `AUDIT_REPORT.md`.

### Crypto & signing

- ed25519/secp256k1 canonical signatures, strict DER, signing-domain prefix enforcement, sorted multisig order.
- XRPL aggregator now sums signer weights for quorum (was counting signers); validated multisig path.
- TEE signer: HMAC + constant-time API-key compare; explicit loopback bind; nil-prv guard.
- TEE attestation: PKI token claims enforced; chain EKUs pinned; CRL fails closed.

### Concurrency

- `priority` queue: closed `Add`-race and missed-wakeup race; ctx-aware lanes; channels allocated in constructor.
- `logger`: `Set` and getter via `atomic.Pointer`; `Nop` honors `Panic`/`Fatal`.
- `database`: `SetErrorLogger` via `atomic.Pointer`.

### Inputs / panics / DoS

- `xrpl/encoding/types`: bounded recursion + length-prefix alloc, MPT bit-mask reject.
- `tee/instruction`: shape + op-policy validation on parsed data.
- `payload`: size cap + duplicate-protocolID reject.
- `safeurl`: extra special-use CIDRs blocked; transport timeouts; redirect cap + downgrade guard.
- `policy.Storage.Add`: underflow guards; `voters.NewSet` rejects duplicates.

### API correctness

- `xrpl/check`: requires `Validated` ledger; `xrpl/transactions`: `Sequence`/`SigningPubKey` required, memo validated.
- `database.DoInTransaction`: panic → returned error; pool-config knobs added.
- `retry`/`call`: backoff/jitter knobs, ctx-aware mid-sleep cancel, deny-list for POST retry.

## Pass 2 — Criticals / Highs / Mediums (2026-05-15, 14 commits)

Independent re-review by 16 blind agents surfaced 10 Highs + 18 Mediums the first pass missed; all closed.

- `merkle`: reject malformed `Tree`; leaf/internal domain-separation docs (the `SortedHashPair` shape was kept — protocol-coordinated rollout, not a library fix).
- `xrpl/encoding`: closed decoder ambiguities vs rippled (AccountID VL, `0xA0` token byte, IOU exponent corruption).
- `xrpl/signing`: enforce low-S, zero-scalar, race-safe `Signer.Value` cache.
- `database`: unbounded-read warnings on public `Fetch*` (kept caller-discretion semantics).
- `xrpl/aggregator`: weight-quorum honored; map-iteration → deterministic; mutex-protected.
- `priority`: `Add`/`AddFast` take `ctx` and return `(*Item, error)` (breaking).
- `voters`: enforce `uint16` weight-sum cap in `NewSet`.
- `tee/signer`: nil-key reject; empty-API-key config warn.
- `xrpl/defs/main`: `nth` int16 range check; format-before-truncate for `autogen.go`.

## Pass 2 — Lows / Infos (2026-05-18, 17 commits)

Remaining 27 Lows + 47 Infos closed in per-package batches.

### Breaking-API changes

Out-of-tree callers will hit compile errors at next bump.

- `safeurl.Validate(rawURL)` → `Validate(ctx, rawURL)`.
- `merkle.BuildFromHex` now returns `(Tree, error)` and rejects non-32-byte hex.
- `priority.New` / `NewWithLogger` return `*PriorityQueue` instead of value.
- `priority` logger interface gained `Panicf` and was exported as `Logger`.
- `heapt.Remove(h, i)` returns `(T, bool)` and refuses OOB.
- `storage.Cyclic.Size()` returns `int` (was truncating `K`).
- `voters.InitialHashSeed(...)` returns `(Hash, error)`.
- `xrpl/check.JSONRPC.Info(addr)` → `Info(ctx, addr)` with a 30 s timeout.

### Other notable fixes

- `database`: DSN password redacted in errors.
- `xrpl/signing`: case-insensitive pubkey prefix routing; derive `PublicKey.X/Y` from `D` when missing.
- `policy.Hash`: handles short inputs without slice OOB.
- `tee/signer`: sanitised error responses; Content-Type enforcement on decrypt; low-S in `RecoverSignersPubKey`.
- `tee/attestation`: empty `EATNonce` warning.
- `xrpl/defs/main`: deterministic `autogen.go` via sorted map keys; bounds-checked FIELDS entries; closed input handle.
- `abi_update.sh`: `set -euo pipefail`; `cd "$(dirname "$0")"`; refuses to write literal `null`.

## Verification at HEAD `e3613ff`

- `go test ./...` — all packages pass.
- `go test -race ./...` — green (notable because pass 2 touched concurrency code in priority, voters, signer, aggregator, logger).
- `go vet ./...` — clean.
- `golangci-lint run ./...` — 0 issues.
