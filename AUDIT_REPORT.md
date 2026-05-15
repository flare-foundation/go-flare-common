# go-flare-common Security Audit

**Branch:** `xrpl` &nbsp; **Date:** 2026-05-06 &nbsp; **Scope:** all of `pkg/` except `pkg/contracts/*` (autogen)

## Tooling baseline

| Tool                  | Result                                                                       |
| --------------------- | ---------------------------------------------------------------------------- |
| `go vet ./...`        | clean                                                                        |
| `golangci-lint run`   | 0 issues                                                                     |
| `go test -race ./...` | all packages pass, no races detected                                         |
| `govulncheck ./...`   | 5 findings, all in `go-ethereum v1.16.7` (p2p/RLPx — unused by this library) |

**`govulncheck` actions:** bump `go-ethereum` to `v1.17.0` to clear all five (`GO-2026-4314`, `4315`, `4507`, `4508`, `4511`).
None reachable in current code paths since this library doesn't run a P2P node, but the upgrade is cheap and reduces SBOM noise.

---

## Executive summary

The audit surfaced **6 Critical**, **~20 High**, and **~25 Medium** findings across 13 packages.
The structural risks cluster in three areas:

1. **TEE attestation gateway is open by default.** `pkg/tee/attestation/googlecloud` deserializes JWT claims but never asserts `iss`, `aud`, image identity, `dbgstat`, or `secboot`.
   Worse, the actual gateway in `tee-node` calls `ParsePKITokenUnverified*`, which performs zero signature verification.
   An attacker who runs _any_ Confidential Space VM (or who fabricates a JWT body with no signature) can pose as the Flare TEE.
2. **`tee/signer` is an unbound signing oracle.** A static API key in TOML is the only authorization.
   `/sign` will sign any 32-byte hash; `/decrypt` will ECIES-decrypt any ciphertext addressed to the TEE key.
   There is no nonce store, no link to attestation, no domain separator, and (depending on `cfg.Addr`) no TLS.
   Cross-domain attacks (XRPL hash signed as if EVM, or vice versa) are structurally possible.
3. **XRPL parser & multi-sig assembly have crash- and protocol-level bugs.** `MarshalDER` and the ed25519/secp256k1 prefix dispatch panic on attacker-shaped inputs reachable from `ValidateMultiSig`.
   The aggregator silently truncates collected signatures to `Quorum` and counts signer _cardinality_ instead of summed `SignerWeight`.
   `CheckAndEncodePayment` produces a signing-form blob without enforcing memo validity, since `ValidateMemo` is never called from the entrypoint.

`go-ethereum 1.16.7` p2p CVEs (5) are not exploitable here but should be cleared by upgrading to ≥1.17.0.

---

## Closure status as of 2026-05-08

Annotates the original findings below with what landed on branch `xrpl`, what was deferred and why, and what remains.
Section bodies further down this document are the as-found analysis and remain unchanged.

### Closed

| Finding          | Commit    | Notes                                                                                                                                                                                                                                                                                                                           |
| ---------------- | --------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| C1, C2, H10, H11 | `cc29f9e` | `Policy` struct + `iss`/`aud`/`exp`/leeway/`secboot`/`dbgstat`/`image_id`/`hwmodel`/`eat_nonce` enforced by the library; required policy fields fail-closed at parse time. 16 new policy-enforcement subtests.                                                                                                                  |
| C3               | `8d5e4eb` | `ParsePKITokenUnverified*` doc expanded to spell out the safety contract (only safe for self-fetched tokens from the local metadata server). Function name kept; correct usage at `tee-node` is a separate-repo concern.                                                                                                        |
| H1, H2, H3, H4   | `305de56` | Bound checks on `secp256k1.MarshalDER` `rLen`/`sLen` (`[1,33]` pre-strip, `[1,32]` post-strip); length guards on `signing.ValidateMultiSig` `SigningPubKey[0:2]` and `ed25519.Validate` `pubBytes[0]`. Each fix has a NotPanics+Error regression case using the original panic-trigger input.                                   |
| H5               | `bf07f8d` | `PrivKeyFromSecret` requires the canonical 23-byte xrpl.js shape (`3+16+4`). Non-canonical secrets that previously produced non-interoperable keys are now rejected.                                                                                                                                                            |
| H6, H7           | `83b9fbe` | H6: `CheckAndEncodePayment` now walks `decoded["Memos"]`, hex-decodes `MemoType`/`MemoFormat`, runs `ValidateMemo`, and enforces rippled's 1024-byte serialized cap (`STTx.cpp`). H7: function name kept; doc expanded to spell out the for-signing canonical form contract.                                                    |
| H9               | `087aacb` | `Account.SignerList` is now `map[string]uint16` carrying `SignerWeight`. Quorum check sums weights instead of counting signers, mirroring rippled `STTx::checkMultiSign`. Today every weight is 1 (enforced in `pkg/xrpl/check`), so this is a no-op for current deployments — but the cross-package safety dependency is gone. |
| H13              | `9dd9563` | Documented that the caller must source the Confidential Space Root CA from a trusted external location and verify it out-of-band by fingerprint. Not embedded in the repo by design — embedding would create a single supply-chain target.                                                                                      |
| H14              | `639d8d6` | Unsafe `DecodeTo` body removed; `DecodeTo2`'s `abi.ConvertType` strict-shape implementation now lives under the `DecodeTo` name. The 9 production callers across `tee-node`/`tee-proxy` automatically gain strict-shape verification at next bump — no migration required.                                                      |
| H16              | `7174393` | `JSONRPC` gained an `http.RoundTripper` field. Default behavior unchanged (works for VPC-internal rippled). Callers expecting public-only URLs can pass `safeurl.NewTransport()`; the audit's "default to safeurl" framing didn't fit the deployment model where rippled lives on the VPC.                                      |
| govulncheck      | `d7ab0dc` | `go-ethereum` bumped from v1.16.7 to v1.17.2; the five p2p/RLPx CVEs (`GO-2026-4314/4315/4507/4508/4511`) cleared. Two remaining advisories are stdlib (fixed in `go1.26.3`), unrelated to this dep.                                                                                                                            |

### Deferred by design

| Finding    | Reason                                                                                                                                                                                                     |
| ---------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| C4, C5, C6 | TEE signer is loopback-only inside the host trust boundary. The audit's threat model (compromised remote orchestrator hitting `/sign`/`/decrypt`) doesn't apply. Anyone who reaches the host already wins. |
| H8         | `Finalize`'s non-deterministic signer subset has the same on-chain effect for current deployments. Byte-determinism of the finalized blob isn't relied on.                                                 |
| H12        | Variable-length x5c chain — pure future-proofing; not exploitable today, would only matter if Google rotates to a different chain depth.                                                                   |
| H15        | `op.IsValid` default-allowing non-`F_`-prefixed types is the documented extension namespace contract. Validation of consumer-defined types is the consumer's job.                                          |

### Outstanding (priority order)

#### High — before next release

- **H17** — `safeurl.Validate` is a TOCTOU helper masquerading as the authoritative check.
  Doc/rename only; tee-relay already uses `NewTransport`.
- **H18** — default `http.Client` redirect policy follows up to 10 redirects without re-validating the target.
- **H19** — `voters.NewSet` accepts duplicate voters in raw bytes; slices retain duplicates while `VoterDataMap` dedupes.
  Every consumer iterating slices double-counts weight.
- **H20** — `database.DoInTransaction` `recover()` swallows panic without re-panicking and has no named return; caller treats panicked-and-rolled-back transaction as `nil` error / committed.
  Real correctness bug, broad blast radius.
- **H21** — Static API key compared via map lookup; non-constant-time.
  Cheap defense even with loopback signer.
- **H22** — Signer plaintext `ListenAndServe`, `cfg.Addr` accepts any address.
  Partly mooted by the loopback-only assumption; cheap to add a loopback default at `New()`.

#### Medium — backlog

M1–M3 (XRPL codec alloc bounds — DoS via huge length prefix), M14 (Instruction field validation), M22/M23 (payload framing).
Plus the rest of the Medium/Low tiers below — see per-finding sections.

---

## Closure status update as of 2026-05-15

Branch `fullAudit`.
The post-2026-05-08 work closed the remaining Highs (H17–H22) and the full Medium tier (M1–M28).
The 2026-05-08 "Outstanding" section above is superseded by the tables below.

### Closed (Highs after 2026-05-08)

| Finding  | Commit    | Notes                                                                                                                                                   |
| -------- | --------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| H17, H18 | `21ed254` | safeurl `Validate` doc-marked non-authoritative; default client `CheckRedirect` caps chain length, rejects https→http downgrade, re-validates each hop. |
| H19      | `a351e13` | `voters.NewSet`/`FromRawBytes` rejects duplicate voter addresses; fixture added.                                                                        |
| H20      | `c2c6f03` | `DoInTransaction` named return + `err = fmt.Errorf("panic in transaction: %v", r)` after rollback; panic no longer treated as committed.                |
| H21      | `7430e74` | API-key compared via HMAC + `subtle.ConstantTimeCompare`.                                                                                               |
| H22      | `4e9da0a` | tee/signer requires explicit loopback bind address.                                                                                                     |

### Closed (Mediums)

| Finding                  | Commit    | Notes                                                                                                                                                               |
| ------------------------ | --------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| M1–M3                    | `a1a22c1` | xrpl/encoding/types: depth-bound STObject/STArray; length-prefix amplification guard before alloc; MPT reserved-bit strictness.                                     |
| M4–M5                    | `f16e049` | xrpl/transactions: Sequence/SigningPubKey required; self-payment reject doc-scoped to native payments.                                                              |
| M6, M7, M9 (M8 doc-only) | `7d0e3b1` | xrpl/signing: JoinMultisig validates signers and strict order; strict-DER in secp256k1; M8 documented as doc-only (Go GC + math/big make true zeroization a lie).   |
| M10                      | `03cf271` | aggregator.AddSignatures returns `common.Hash` (exported handle). SIG change.                                                                                       |
| M11                      | `49afd26` | xrpl/check rejects non-Validated ledger response.                                                                                                                   |
| M12, M13                 | `5a4ca21` | tee/attestation: `Policy.RequireCRL` fails closed; `Policy.AllowedLeafEKUs` pins chain EKUs; `Verify` kept as compat wrapper over `VerifyWithPolicy`.               |
| M14                      | `acfd151` | tee/instruction `Data.Validate` enforces MaxMessageSize=64KiB, MaxCosigners=32, threshold/uniqueness, op.IsValidPair.                                               |
| M15                      | `a6e72bd` | tee/xrpl native-only scope: nil Amount / sender==recipient / non-empty TokenId rejected; `CheckNativePayment` forced.                                               |
| M16, M17                 | `70481d5` | M16 doc-only note (real binding owed in pkg/tee/structs); M17 signer.Run goroutine leak fixed via buffered chan + fresh shutdown ctx.                               |
| M18, M21                 | `244074d` | safeurl: CGNAT/0.0.0.0/8/TEST-NET CIDRs blocked; explicit transport timeouts.                                                                                       |
| M19                      | `674deb7` | retry.Params Multiplier/MaxDelay/Jitter; call.Params NoRetryStatuses; MaxAttempts≤0 kept as unlimited with doc note; transport-error gating deferred intentionally. |
| M20                      | `6c025a9` | retry.Execute returns directly from the select on ctx-done; same error format preserved.                                                                            |
| M22, M23                 | `46dc10a` | BuildMessage returns `(string, error)` and bounds payload at MaxUint16 (SIG change); ExtractPayloads rejects duplicate protocolID.                                  |
| M24                      | `3099253` | policy.Storage replaces embedded `sync.Mutex` with private `mu` field (drops public Lock/Unlock).                                                                   |
| M25, M26                 | `98330f2` | database.Config gains `PoolConfig` (MaxOpen/Idle/Lifetime/IdleTime); `DoInTransaction` surfaces Begin and Rollback errors.                                          |
| M27                      | `eaaaaf4` | doc-only: leaf/internal domain-separation precondition formalized on package doc and Build/SortedHashPair/VerifyProof.                                              |
| M28                      | `89341f5` | priority queue channels allocated in `NewWithLogger`; `InitiateAndRun` only spawns goroutines. Add-before-Initiate is safe.                                         |

### Outstanding

None for Critical, High, or Medium tiers.
Low/Info backlog (~25 items in the section below) remains unscheduled.

---

## Critical (P0 — fix immediately)

### C1. TEE attestation: `iss`/`aud` never validated

**Location:** `pkg/tee/attestation/googlecloud/google_cloud.go:97-105` Validator passes only `WithValidMethods(["RS256"])`.
No `WithIssuer`, no `WithAudience`.
Any Google Confidential Space VM in the world produces a token signed by the same root chain — without `iss`/`aud` pinning, **anyone running their own TEE workload passes verification**.
**Fix:** Pin `iss="https://confidentialcomputing.googleapis.com"` and the project's expected `aud`.
Audit the only call site in `tee-node/internal/attestation/attestation.go:48` — it uses `ParsePKITokenUnverifiedClaims`, see C3.

### C2. TEE attestation: workload image identity not verified

**Location:** `pkg/tee/attestation/googlecloud/google_cloud.go:80-87` (and the _absence_ of any allowlist comparison) `CodeHash()` returns whatever the token says; `image_reference`, `dbgstat`, `secboot`, and `submods.container.image_digest` are deserialized but not asserted.
Recent commit `5b952b6` switched from `image_digest` → `image_id`; verify on-chain registry stores the matching value (else every attestation mismatches).
**Fix:** Assert `dbgstat == "disabled-since-boot"`, `secboot == true`, and pin `image_digest` (or `image_id`) against an allowlist.
Add `image_reference` and pin the registry path.

### C3. TEE attestation: `ParsePKITokenUnverified*` is the path actually used by `tee-node`

**Location:** `google_cloud.go:107-121`; caller `tee/tee-node/internal/attestation/attestation.go:48` The `tee-node` extracts `CodeHash`/`Platform` via the unverified parser — **zero signature checking, zero claims checking**.
An attacker hand-crafts a JWT body with arbitrary `image_id` and `hwmodel`, no key needed, the gateway opens.
**Fix:** Either rename/unexport `ParsePKITokenUnverified*` into an `unsafe` subpackage, or fix `tee-node` to call `ParseAndValidatePKIToken` with pinned root + claims.
Add a package-level doc warning.

### C4. TEE signer: unbound signing oracle

**Location:** `pkg/tee/signer/signer.go:81-106, 158-235` `POST /sign` accepts an arbitrary list of 32-byte hashes and signs them.
No link to `instruction`/`op`/`attestation`.
No nonce store.
No per-key allowlist.
A compromised orchestrator (which legitimately holds the API key to operate) becomes a full signing oracle that can extract signatures over any digest the TEE key controls — including cross-domain (an XRPL transaction blob is indistinguishable to `/sign` from an EVM message hash).
**Fix:** Make the signer aware of the instruction it serves: accept `Instruction` (data + signature recoverable to authorized originator), recompute on-chain-authenticated `HashForSigning`, derive the message-set, and refuse anything else. Add per-domain prefixes (`keccak("flare.tee.evm.v1" || h)` vs `"flare.tee.xrpl.v1"`). Replace static API key with attestation-bound session.

### C5. TEE signer: `/decrypt` is a generic ECIES oracle

**Location:** `pkg/tee/signer/signer.go:301-344, 366-377` `/decrypt` decrypts whatever ciphertext the API-keyed caller submits and returns plaintext.
ECIES authenticates the ciphertext but there's no caller-supplied AD, no policy, no rate limit, no link to attestation.
Anyone with the API key recovers any sealed key-share/backup addressed to this TEE key.
**Fix:** Refuse generic decryption.
Require an authenticated `Instruction` of type `op.Wallet`/`KeyDataProviderRestore` whose `OriginalMessage` is the only ciphertext accepted.
Bind ECIES `s1` to `instructionId`.

### C6. TEE signer: no replay/nonce protection on `/sign` or `/decrypt`

**Location:** `pkg/tee/signer/signer.go:169-213, 301-344` Same `SignBody` POSTed N times → N identical signatures.
Same ciphertext → unbounded decrypt repeats.
If a higher-layer protocol depends on the signer rate-limiting or refusing repeats (e.g., key-proof binding to a one-shot challenge), that property does not exist here.
**Fix:** Bounded LRU + persisted high-watermark per `instructionId`; reject repeats.

---

## High (P1)

### XRPL signing — parser panics reachable from `ValidateMultiSig`

| #   | Location                                        | Issue                                                                                                                                                                   |
| --- | ----------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| H1  | `pkg/xrpl/signing/secp256k1/signature.go:64,82` | `MarshalDER` `copy(sigOut.r[32-rLen:], …)` with attacker-controlled `rLen ≥ 33` → negative slice index panic. Reachable via `Validate → MarshalDER → ValidateMultiSig`. |
| H2  | `pkg/xrpl/signing/secp256k1/signature.go:77`    | `MarshalDER` indexes `sig[sStart]` without bound check when `sLen == 0`.                                                                                                |
| H3  | `pkg/xrpl/signing/signing.go:33`                | `s.SigningPubKey[0:2]` panics on signer with pubkey shorter than 2 chars.                                                                                               |
| H4  | `pkg/xrpl/signing/ed25519/ed25519.go:166`       | `pubBytes[0]` indexed before length check; empty hex-decoded pubkey panics.                                                                                             |

**Combined fix:** input validation at every entrypoint.
Replace hand-rolled DER with `dcrd/secp256k1/v4/ecdsa.ParseDERSignature`.
Reject `rLen/sLen ∉ [1,32]` after the leading-zero strip; check `len(SigningPubKey) >= 2` before indexing; check `len(pubBytes) == 33` before reading prefix.

### XRPL ed25519 secret length too permissive

**H5.** `pkg/xrpl/signing/ed25519/ed25519.go:65-96` — `PrivKeyFromSecret` accepts `len(secretBytes) >= 7` but canonical XRPL ed25519 secret is **exactly 23 bytes**.
Different lengths still produce a valid (but non-canonical) key, breaking interoperability with rippled/xrpl.js.
**Fix:** `if len(secretBytes) != 3+16+4 { ... }` like `seed.DecodeFamilySeed`.

### XRPL transactions — entrypoint contract gaps

| #   | Location                                 | Issue                                                                                                                                                                           |
| --- | ---------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| H6  | `pkg/xrpl/transactions/payment.go:12-82` | Memo validation unreachable: `CheckAndEncodePayment` never calls `ValidateMemo`, so the recent rippled-`isMemoOkay` alignment (`a11b779`) is bypassed by the actual entrypoint. |
| H7  | `pkg/xrpl/transactions/payment.go:13`    | `types.Encode(tx, true)` returns the **signing-form** blob, but the function name and contract give no hint. Callers can submit unsigned blobs or assume the wrong wrapping.    |

**Fix:** in `CheckAndEncodePayment`, walk `decoded["Memos"]` and run `ValidateMemo` per entry.
Rename to `EncodeForSigning` or split signing/wire variants explicitly.

### XRPL multi-sig — quorum/weight semantics

| #   | Location                                                  | Issue                                                                                                                                                                                                                       |
| --- | --------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| H8  | `pkg/xrpl/aggregator/aggregator.go:158-189`               | `Finalize` truncates collected signatures to `Quorum` via map iteration. Redundant honest sigs dropped non-deterministically; two `Finalize` calls for same id produce different blobs.                                     |
| H9  | `pkg/xrpl/aggregator/aggregator.go:139` + `Account` model | Counts signer **cardinality** not summed `SignerWeight`. Safe today only because `check.go:173` enforces `weight==1`; if invariant ever relaxes (or `Account` built from raw on-ledger SignerList), quorum semantics break. |

**Fix:** include all collected signers in the multi-sig blob (rippled allows up to 32).
Propagate `SignerWeight` end-to-end and sum properly; mirror rippled `STTx::checkMultiSign`.

### TEE attestation — additional gaps

| #   | Location                       | Issue                                                                                                                                                                                                                                                                                    |
| --- | ------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| H10 | `google_cloud.go:97-105`       | No `WithExpirationRequired()`, no leeway, no nonce binding. Tokens with no `exp` verify forever; `eat_nonce` parsed but never bound to a per-request challenge.                                                                                                                          |
| H11 | `google_cloud.go:44-54, 70-77` | `dbgstat`, `secboot`, `hwmodel` deserialized but never asserted. Debug-mode VM (reduced isolation) verifies.                                                                                                                                                                             |
| H12 | `google_cloud.go:164-166`      | `len(x5cHeaders) != 3` hard-rejects any other depth. Brittle to Google chain rotation.                                                                                                                                                                                                   |
| H13 | `google_cloud.go:240-242`      | Trust anchor matched by `(*x509.Certificate).Equal` rather than fingerprint. Mis-loaded `expectedRoot` (writable file, network-fetched) silently substitutes. **Fix:** embed the Confidential Space Root CA via `//go:embed`, or compare SHA-256 fingerprint against a hard-coded value. |

### TEE structs / op routing

**H14.** `pkg/tee/structs/decode.go:34-40` — `DecodeTo` wraps `dest` as `struct{ X any }`, defeating ABI strict struct-shape verification.
Wrong destination type silently misinterprets a TEE message.
**Fix:** delete `DecodeTo` or rewrite like `DecodeTo2[T]` (`abi.ConvertType`); audit callers.

**H15.** `pkg/tee/op/op.go:124-131` — `IsValid` default-allows any non-`F_`-prefixed type/command pair (textual prefix check).
Tests confirm `{"x","x"}` validates.
**Fix:** default-deny — require `validSystemPairs` membership for system types; require enumerated extensions for non-system types.

### Network safety

| #   | Location                                              | Issue                                                                                                                                                                                                                                                                                                 |
| --- | ----------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| H16 | `pkg/call/post.go:28,44`                              | `Params.Transport` plumbed straight into `http.Client` with no constraint. The only in-tree caller (`pkg/xrpl/check/check.go:94`) leaves it nil — **SSRF protection added in `73521b6` is not actually wired up by the only consumer**.                                                               |
| H17 | `pkg/safeurl/safeurl.go:21-53`                        | `Validate` does `LookupHost` and returns; nothing prevents a caller from then using `http.Get` (DNS-rebinding TOCTOU).                                                                                                                                                                                |
| H18 | `pkg/safeurl/safeurl.go:58-96`, `pkg/call/post.go:44` | Default `http.Client` redirect policy follows up to 10 redirects without scheme/host re-validation. With `safeurl.NewTransport` each redirect IP is re-checked at dial — but `https→http` downgrade and proxy-abuse still possible. With `Validate`-only flow, redirect target is never re-validated. |

**Fix:** make safeurl the default in `call.PostRaw`.
Either remove `Validate` (require connections through `NewTransport`) or rename to `ValidateOptimistic` with a doc warning.
Provide `NewClient` with a `CheckRedirect` that caps chain length, rejects downgrade, and re-validates the target.

### Policy / payload

**H19.** `pkg/voters/voters.go:39-58` (called from `pkg/policy/policy.go:149,56`) — Duplicate voters in raw bytes are accepted; `vs.voters[]/weights[]/thresholds[]` retain duplicate slots, so a duplicated address gets two random-selection windows and double-counted weight in any consumer iterating the slices (vs the dedup'd `VoterDataMap`).
**Fix:** error on duplicate during `FromRawBytes` / `NewSet`.
Add fixture.

### Database

**H20.** `pkg/database/utils.go:122-137` — `DoInTransaction` `recover()` neutralises panic without re-panicking, has no named return, returns nil error.
Caller treats panic-and-rolled-back transaction as committed.
**Fix:** named return + `err = fmt.Errorf("panic in transaction: %v", r)` after rollback.

### TEE signer — additional

| #   | Location                           | Issue                                                                                                                                                     |
| --- | ---------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------- |
| H21 | `pkg/tee/signer/signer.go:152-156` | API-key auth is `map[string]bool` lookup — non-constant-time. Hash with HMAC of a startup-random key, compare with `subtle.ConstantTimeCompare`.          |
| H22 | `pkg/tee/signer/signer.go:95-103`  | No TLS — `http.Server` constructed without `TLSConfig`, `Run` calls `ListenAndServe` plaintext. `cfg.Addr` accepts any address (no loopback enforcement). |

---

## Medium (P2)

| #   | Pkg                 | Location                                                  | Issue                                                                                                                                                                                               |
| --- | ------------------- | --------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| M1  | xrpl/encoding/types | `coder.go:288-330`, `object.go:32-67`, `array.go:60-89`   | No recursion-depth bound in `STObject`/`STArray`; nested-marker bombs amplify CPU/alloc churn.                                                                                                      |
| M2  | xrpl/encoding/types | `coder.go:159-191`, `hash.go:49-69`, `vector256.go:31-49` | Length prefix → `make([]byte, l)` of up to 918 744 with no comparison to remaining `b.Len()`. ~230 000× amplification per field. **Fix:** `if l > b.Len() { return outOfBytes(...) }` before alloc. |
| M3  | xrpl/encoding/types | `amount.go:96-106, 437-475`                               | MPT amount accepts unspecified indicator bit patterns; rippled would reject. Compatibility/consensus drift.                                                                                         |
| M4  | xrpl/transactions   | `payment.go:24-66`                                        | Required common fields `Sequence` / `SigningPubKey` not enforced (own test documents this).                                                                                                         |
| M5  | xrpl/transactions   | `payment.go:43-45`                                        | Self-payment unconditionally rejected — breaks legit XRP-to-IOU conversion (`Paths`/`SendMax`).                                                                                                     |
| M6  | xrpl/signing        | `signing.go:53-91`                                        | `JoinMultisig`/`JoinMultisigJSON` skip signature validation — commented-out checks.                                                                                                                 |
| M7  | xrpl/signing        | `secp256k1/signature.go:150-173`                          | `Validate` enforces low-S but not full XRPL "fully canonical" rule (r>0, s>0, strict DER). Wire-format gap with rippled.                                                                            |
| M8  | xrpl/signing        | `secp256k1/derivation.go`, `ed25519/ed25519.go`           | No zeroization of secret material; `math/big.Int` ops are non-constant-time.                                                                                                                        |
| M9  | xrpl/signing        | `signing.go:50-70`                                        | Doc says signers must be sorted, never enforces it. Call `signer.Sort` or assert.                                                                                                                   |
| M10 | xrpl/aggregator     | `aggregator.go:35`                                        | `AddSignatures` returns unexported `*transaction` whose only useful field (`id`) is unexported. External callers cannot finalize.                                                                   |
| M11 | xrpl/check          | `check.go:110-128`                                        | `AccountInfoResponse.Validated` not checked; SignerList from non-final ledger could be rolled back.                                                                                                 |
| M12 | tee/attestation     | `google_cloud.go:315-323`                                 | `checkCRL` returns success when `crl == nil` even if cert has DPs. Defense depends on caller passing CRL. **Fix:** `RequireCRL bool` option, fail closed when DPs present.                          |
| M13 | tee/attestation     | `google_cloud.go:287-291`                                 | `KeyUsages: ExtKeyUsageAny` weakens chain validation. Pin to leaf's actual EKU.                                                                                                                     |
| M14 | tee/instruction     | `instruction.go:15-56`                                    | `Data` parsed from on-chain events without size/field validation: variable bytes lengths, `CosignersThreshold ≤ len(Cosigners)`, distinct cosigners, `op.IsValidPair` — none enforced.              |
| M15 | tee/xrpl            | `payment.go:18-26`                                        | `i.Amount.String()` panics on nil `*big.Int`; no `Sender != Recipient` check; no `tokenId` discrimination; `CheckNativePayment` is optional.                                                        |
| M16 | tee/op              | `KeyProof` op                                             | Added as constant only; no struct/binding to attestation/instructionId/freshness nonce. Replay across attestation contexts possible.                                                                |
| M17 | tee/signer          | `signer.go:108-131`                                       | Goroutine leak on shutdown error path: unbuffered `c` channel, writer blocks forever after `select` takes `ctx.Done()` branch; cancelled ctx passed into `Shutdown`.                                |
| M18 | safeurl             | `safeurl.go:107-118`                                      | CGNAT (100.64.0.0/10) and "this network" (0.0.0.0/8 non-zero) **not** blocked. AWS metadata 169.254.169.254 _is_ (link-local). Add explicit CIDRs.                                                  |
| M19 | call+retry          | `call/post.go:93-110`, `retry/retry.go:19-23`             | `MaxAttempts <= 0` means **unlimited**; no exponential backoff/jitter; POST retried unconditionally → duplicate side effects.                                                                       |
| M20 | retry               | `retry/retry.go:54-73`                                    | Loop continues for one extra cycle on context-derived errors and wraps `context.Canceled`.                                                                                                          |
| M21 | safeurl             | `safeurl.go:58-96`                                        | No connect/TLS/header timeouts on transport. Slowloris vector.                                                                                                                                      |
| M22 | payload             | `message.go:39`                                           | `BuildMessage` truncates `len(payload)` to `uint16` without bounds check; receiver mis-frames.                                                                                                      |
| M23 | payload             | `payload.go:73`                                           | `ExtractPayloads` silently overwrites duplicate `protocolID` entries. Last-write-wins → spoofing surface.                                                                                           |
| M24 | policy              | `storage.go:19`                                           | `Storage` embeds `sync.Mutex` publicly. External `Lock()`/`Unlock()` exposed; `govet copylocks` triggers on accidental copy.                                                                        |
| M25 | database            | `utils.go:26-47`                                          | `Connect` doesn't tune connection pool — unlimited `MaxOpenConns`/`MaxIdleConns`/`ConnMaxLifetime`.                                                                                                 |
| M26 | database            | `utils.go:122-137`                                        | `DoInTransaction` doesn't check `Begin` error or `Rollback` errors.                                                                                                                                 |
| M27 | merkle              | `merkle.go:24-53, 90-96`                                  | No leaf/internal-node domain separation. Currently safe because all leaves are domain-tagged in callers (rewards, FDC, XRP). Document precondition or add tagged hashing.                           |
| M28 | priority            | `priority.go:180,196,116-128`                             | `Add`/`AddFast` write to `p.in`/`p.inFast` which are nil until `InitiateAndRun` is called and assigned without sync. Race / deadlock on misuse.                                                     |

---

## Low / Info

Selected items (~25 total — see per-package agent reports for full text):

- **xrpl/signing**: low-entropy seed validation absent; `MarshalDER` not strict-DER; `SignXRPL` exported without enforcing STX prefix; family-seed `maxAttempts=128` cap citation unverified.
- **xrpl/encoding/types**: IOU exponent/significand ranges not enforced on decode; ISO currency decode doesn't validate charset; PathSet path-count check fires _after_ over-limit append; XChainBridge encode-side dead check; `pathset.readCurrency` returns literal `"nil"` on truncated input.
- **xrpl/transactions**: zero-amount IOU not rejected; `strconv.ParseUint(_, 10, 0)` platform-dependent.
- **xrpl multisig**: `checkSigners` doesn't dedupe within `SignerEntries`; `lsfDisallowXRP` blanket-rejects.
- **xrpl/encoding**: `encoding/transaction.go` is empty stub.
- **xrpl base58/address/hash**: `Sha256RipeMD160` doc comment reversed; base58 length cap duplicated across 4 call sites instead of centralized; `bytes.Equal` for checksum compare (informational — no secret material).
- **policy/payload**: `0X` (uppercase) prefix not stripped; `Selector`/`From` not validated; `Hash` length-extension precondition undocumented; `findByVotingRoundID` inverted comment; `RewardEpochID` underflow when first inserted policy has ID 0.
- **database/storage**: `WaitCIndexerToSync` boundary inconsistency (`<` vs `>`); `SetErrorLogger` package-global without sync; deprecated `NewCyclic` doesn't guard `size <= 0`; `Limit(-1)` semantics undocumented for `LatestLogsParams.Number`.
- **encoding/signature**: `TransformSignatureVRStoRSV` panics on len ≠ 65, byte-subtraction underflow on already-normalised `v`, no low-S enforcement.
- **safeurl/call**: `MaxResponseSize=0` silently truncates; `InsecureSkipVerify` not used anywhere (good); HTTP/2 connection coalescing not re-validated (speculative).
- **lower-risk skim**: `voters.NewSet` returns bare nil on length mismatch; `queue` retry goroutines leak under stalled consumer; `priority.next()` blocks ignoring ctx; `events.HexToHash` silently zero-pads malformed topics; `logger.Set` mutates global without sync; `toml` strictness is caller-controlled.

---

## Suggested fix order

**This week:**

1. C3 — fix `tee-node` to call `ParseAndValidatePKIToken`; rename or unsafe-namespace `ParsePKITokenUnverified*` (gateway-of-record bug).
2. C1, C2, H10–H13 — pin issuer/audience/image identity and tighten `dbgstat`/`secboot`/CRL/EKU in attestation.
3. H1–H4 — input-validate the panic surfaces in `xrpl/signing` (one PR, all four).
4. C4–C6 — re-architect signer authorization (instruction-aware /sign + /decrypt, nonce store, attestation-bound session).
5. H16 — wire safeurl into `call.PostRaw` by default; revisit `xrpl/check` call site.
6. govulncheck — bump `go-ethereum` to `≥1.17.0`.

**Before next release:**

- H5–H9, H14–H15, H17–H22 plus M1–M3 (codec alloc bounds), M14 (instruction validation), M22–M23 (payload framing).

**Backlog:**

- All remaining Medium/Low.
  Most of the Mediums in `xrpl/signing` (M6–M9) and `xrpl/encoding/types` (M1–M3) are appropriate for one focused hardening sprint.

---

## Coverage note

13 parallel agents reviewed every non-test, non-autogen `.go` file in scope (~138 files, ~22k LOC) plus relevant tests for context, plus recent commit history.
Tooling baseline ran cleanly; no `-race` failures.
Out of scope: `pkg/contracts/*` (autogen abigen), CHANGELOG/docs.
