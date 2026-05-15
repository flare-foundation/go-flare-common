# Audit Report — Pass 2 (Independent Review)

**Date:** 2026-05-15 **Branch:** `fullAudit` **Scope:** All hand-written Go packages under `pkg/` (abigen-generated `autogen.go` files excluded).
Module meta (`go.mod`, `abi_update.sh`) also covered.
**Method:** 16 parallel review agents, each with an independent prompt and a fixed bug-class taxonomy.
Agents were explicitly instructed **not** to read the prior `AUDIT_REPORT.md`, to avoid anchoring.

## Severity rubric

- **Critical:** remote-exploitable; silent data corruption; private-key exposure; signature forgery.
- **High:** panic reachable from untrusted input; unbounded resource use; crypto misuse (malleability, weak nonce, missing canonicalisation); consensus drift.
- **Medium:** panic from misconfig; race in observable state; error swallowed in path that matters.
- **Low:** GOAI.md style violation with real consequence; minor correctness.
- **Info:** defensive note; no functional impact.

## Counts

| Severity | Count |
| -------- | ----- |
| Critical | 0     |
| High     | 10    |
| Medium   | 18    |
| Low      | 27    |
| Info     | 47    |

## High-severity headline

- **Merkle tree has no domain separation** between leaf and internal nodes — second-preimage proof forgery is possible without caller-side discipline (F-MERKLE-1).
- **XRPL encoder/decoder ambiguities** — AccountID VL-prefix ignored on decode (F-ENC-1) and reserved `0xA0` first-byte accepted with the stray `0x20` corrupting the IOU exponent (F-ENC-2).
  Both enable rippled-vs-Go consensus drift on the same signed bytes.
- **secp256k1 verifier accepts high-S signatures** — malleability not blocked at the verification boundary (F-SIGNCURVE-1).
- **`pkg/database`** has two paths (`FetchLogs`, `FetchTransactions`, `fetchLogsFull`) that materialise full tables into memory when called with zero-value filters (F-DB-1, F-DB-2).
- **`pkg/xrpl/aggregator`** has two distinct High bugs: `Finalize` treats `Quorum` (a weight threshold) as a signer count, blocking finalization for weighted multisig (F-AGG-1); and the aggregator has no synchronisation despite a multi-source signature-collection role (F-AGG-3).
- **`pkg/priority`** has a missed-wakeup race in `next()` that strands items (F-PRI-1) and an `Add`/`AddFast` post-cancel hang (F-PRI-2).

---

# Findings by package

## pkg/call & pkg/safeurl

### F-CALL-1: `MaxResponseSize` zero-value silently breaks all 200-OK responses

- **Location:** `pkg/call/post.go:86-94` (also `:74-80`)
- **Severity:** Medium
- **Class:** Logic bug / misconfig → silent failure
- **Trigger:** Caller constructs `Params{Timeout: ...}` and omits `MaxResponseSize`.
- **Impact:** `io.LimitedReader{N: 0}` returns EOF on first read; `json.Decoder.Decode` fails on every successful response.
  The error-path branch (l.75-80) likewise captures an empty body.
  `PostWithRetry` then retries POSTs to exhaustion.
- **Code:**

```go
respLimited := &io.LimitedReader{R: resp.Body, N: p.MaxResponseSize}
decoder := json.NewDecoder(respLimited)
response := new(T)
err = decoder.Decode(response)
```

### F-CALL-2: `PostRawWithRetry` reads the entire request body with no size cap

- **Location:** `pkg/call/post.go:101-105`
- **Severity:** Low
- **Class:** Unbounded read
- **Trigger:** Caller passes a `body io.Reader` from network/file/attacker source.
- **Impact:** `io.ReadAll(body)` is unbounded; a large/never-ending reader OOMs the process.
  `MaxResponseSize` is not consulted here.
- **Code:**

```go
bodyBytes, err := io.ReadAll(body)
```

### F-CALL-3: Brittle `Content-Type` exact-string match drops server error reason

- **Location:** `pkg/call/post.go:74`
- **Severity:** Low
- **Class:** Inverted-string match
- **Trigger:** Server emits `text/plain`, `text/plain;charset=utf-8` (no space), `text/plain; charset=UTF-8`, etc.
- **Impact:** Plain-text error body silently discarded; caller only sees status code.
  Use `mime.ParseMediaType` or case-insensitive prefix match.
- **Code:**

```go
if resp.Header.Get("Content-Type") == "text/plain; charset=utf-8" {
```

### F-CALL-4: `safeurl.Validate` uses uncancellable DNS lookup with no timeout

- **Location:** `pkg/safeurl/safeurl.go:42`
- **Severity:** Low
- **Class:** Missing ctx propagation
- **Trigger:** Hostile or slow authoritative server; loss of resolver connectivity.
- **Impact:** Blocks for the OS resolver timeout (tens of seconds), cannot be cancelled.
  The dialer at line 87 already uses ctx; `Validate` is the asymmetric outlier.
- **Code:**

```go
ips, err := net.LookupHost(host)
```

### F-CALL-5: `Params.Transport` defaults to `http.DefaultTransport` (no SSRF protection)

- **Location:** `pkg/call/post.go:52`
- **Severity:** Info
- **Class:** Defensive default / SSRF posture
- **Trigger:** Caller omits `Params.Transport`.
- **Impact:** `call` looks "safe" but dials any address unless caller wires `safeurl.NewTransport()` explicitly.
  Worth documenting in the field comment or defaulting to safeurl on nil.

### F-CALL-6: Dialer probes resolved IPs serially with full connect timeout each

- **Location:** `pkg/safeurl/safeurl.go:104-113`
- **Severity:** Info
- **Class:** Amplification
- **Trigger:** Hostile DNS reply with N unreachable IPs.
- **Impact:** N × `defaultConnectTimeout` (10s) tarpit per call.
  Use happy-eyeballs / parallel dial or cap against ctx deadline.

---

## pkg/database

### F-DB-1: `fetchLogsFull` allows unbounded full-table scan with empty filters

- **Location:** `pkg/database/queries.go:115`
- **Severity:** High
- **Class:** Unbounded reads / DoS
- **Trigger:** `LogsFullParams{Address: zero, Topics: zero, Number: -1}`.
  `pMap` empty, gorm `Where(emptyMap)` adds no WHERE, `Limit(-1)` is unlimited.
- **Impact:** Full table scan of `logs` materialised into a slice — OOM, DB load spike.
- **Code:**

```go
pMap := make(map[string]any)
err := db.WithContext(ctx).Where(pMap).Order("timestamp DESC").Limit(params.Number).Find(&logs).Error
```

### F-DB-2: `FetchLogs` / `FetchTransactions` allow unbounded reads on all-nil filters

- **Location:** `pkg/database/db.go:217` (FetchLogs), `:255` (FetchTransactions)
- **Severity:** High
- **Class:** Unbounded reads / DoS
- **Trigger:** All pointer fields in `LogsQuery` / `TxQuery` nil — no WHERE, no LIMIT.
- **Impact:** Loads the entire `logs`/`transactions` table (potentially billions of rows on a C-chain indexer) into memory.
- **Code:**

```go
func dbFetchLogs(ctx context.Context, db *gorm.DB, params LogsQuery) ([]Log, error) {
    q := db.WithContext(ctx)
    if params.From != nil { q = q.Where("timestamp > ?", *params.From) }
    // ... all optional ...
    err := q.Order("timestamp, id").Find(&logs).Error
```

### F-DB-3: `RetryWrapper` retries deterministic "not found" sentinels for 15 s

- **Location:** `pkg/database/queries.go:328`, `:62`, `:321`
- **Severity:** Medium
- **Class:** Error handling
- **Trigger:** Pristine indexer DB / missing state row.
  `fetchLatestBlock` and `fetchState` return plain `errors.New(...)`, not `backoff.Permanent`.
- **Impact:** Hot loops (`WaitCIndexerToSync`, `FetchFirstDBBlockTs`) hang for 15 s on every call against an empty DB.
  Compare `retryWithConfig` in db.go which honors `PermanentError`.
- **Code:**

```go
if len(blocks) == 0 {
    return Block{}, errors.New("no blocks in database")   // not Permanent
}
```

### F-DB-4: Per-range query helpers have no LIMIT

- **Location:** `pkg/database/queries.go:154,174,196,225,247,269`
- **Severity:** Medium
- **Class:** Unbounded reads
- **Trigger:** Wide `(From, To]` window passed to any `fetchLogsBy*` / `fetchTransactionsBy*` helper.
- **Impact:** Millions of matching rows accumulated into a slice; memory pressure and sustained DB load.
  Public-API surface of the shared library.

### F-DB-5: `WaitCIndexerToSync` performs `Retries+1` checks

- **Location:** `pkg/database/utils.go:107`
- **Severity:** Low
- **Class:** Off-by-one
- **Trigger:** Any call.
  Loop runs Retries iterations and then performs an additional "final time" check unconditionally.
- **Impact:** Surplus DB load, behavior diverges from doc-comment contract.
  `Retries == 0` still performs one fetch.

### F-DB-6: DSN with password may leak via `gorm.Open` error

- **Location:** `pkg/database/utils.go:57`
- **Severity:** Low
- **Class:** Credential exposure
- **Trigger:** `gorm.Open` (or underlying mysql driver) returns an error embedding the parsed DSN.
- **Impact:** If callers log the error, MySQL password is exposed.
  Redact or re-wrap before returning.

### F-DB-7: `uint64 → int64` conversion on `state.BlockTimestamp`

- **Location:** `pkg/database/utils.go:117,143,198`
- **Severity:** Info
- **Class:** Integer overflow on conversion
- **Trigger:** `BlockTimestamp > math.MaxInt64` (unreachable for valid Unix-seconds).
- **Impact:** None practically; GOAI pitfall list flags `int(uint64Var)` patterns.

### F-DB-8: gorm `Limit(0)` returns no rows — surprising for default-initialised struct

- **Location:** `pkg/database/queries.go:92,137`
- **Severity:** Info
- **Class:** API contract
- **Trigger:** `LatestLogsParams{Number: 0}` or `LogsFullParams{Number: 0}` (the zero value).
- **Impact:** Caller silently gets an empty result.
  Documented in code, but worth tightening to a typed sentinel.

---

## pkg/encoding

### F-ENC2-1: `TransformSignatureVRStoRSV` accepts V > 28, producing an out-of-range recid

- **Location:** `pkg/encoding/signature.go:59-73`
- **Severity:** Medium
- **Class:** Crypto misuse / missing range check
- **Trigger:** 65-byte `vrs` with first byte in `[29, 255]`.
  Recent fix added lower bound (`vrs[0] < 27` rejected) but no upper bound.
- **Impact:** EIP-155 V values (`chainID*2 + 35 + {0,1}`) or garbage become `rsv[64] = V-27`, e.g. V=37 → recid 10, V=0xFF → recid 228.
  Downstream `ecrecover` rejects; liveness loss with the caller's configuration bug masked.
- **Code:**

```go
if vrs[0] < 27 {
    return nil, fmt.Errorf("invalid V byte %d, expected >= 27 (input may already be normalised)", vrs[0])
}
// missing: if vrs[0] > 28 { return nil, ... }
rsv[64] = vrs[0] - 27
```

### F-ENC2-2: `TransformSignatureRSVtoVRS` accepts recid byte outside {0, 1}

- **Location:** `pkg/encoding/signature.go:77-88`
- **Severity:** Medium
- **Class:** Crypto misuse / missing range check (symmetric to F-ENC2-1)
- **Trigger:** 65-byte `rsv` whose last byte is not 0 or 1.
- **Impact:** `vrs[0] = rsv[64] + 27` is computed with no validation.
  Values 2-228 produce out-of-spec V bytes (29-255); ≥ 229 wrap (e.g. recid=229 → V=0).
  Malformed signature propagated as valid.

### F-ENC2-3: `EncodeSignatures` truncates `Index` to uint16, breaking sort/uniqueness

- **Location:** `pkg/encoding/signature.go:23-53`
- **Severity:** Medium
- **Class:** Logic / signed-unsigned narrowing
- **Trigger:** `IndexedSignature.Index` in `[math.MaxUint16+1, math.MaxInt]`.
  The monotonicity check compares full `int` values and passes for `[65535, 65536]`; then `convert.Uint16ToBytes(uint16(...))` truncates.
- **Impact:** A caller can pass sorted, distinct `int` indices `[65535, 65536]` and produce an encoded payload whose on-the-wire index sequence is `[65535, 0]` — duplicate and decreasing.
- **Code:**

```go
if signature.Index < 0 {
    return nil, errors.New("payload index not set")
}
// missing: if signature.Index > math.MaxUint16 { return nil, ... }
if prevIndex >= signature.Index { ... }
indexBytes := convert.Uint16ToBytes(uint16(signature.Index))
```

### F-ENC2-4: No low-S / canonical-signature enforcement in transforms

- **Location:** `pkg/encoding/signature.go:59-88`
- **Severity:** Info
- **Class:** Crypto / malleability (defensive)
- **Trigger:** Output consumed by a non-`ecrecover` verifier.
- **Impact:** Transforms are pure byte-layout; callers must enforce low-S (`S ≤ N/2`) independently.
  Doc note recommended.

### F-ENC2-5: Stale doc comment on `EncodeSignatures` still claims panic-on-wrong-length

- **Location:** `pkg/encoding/signature.go:20-22`
- **Severity:** Low
- **Class:** Doc accuracy
- **Impact:** Function now returns an error on a non-65-byte input; doc still says "will panic".
  GOAI requires doc accuracy.

---

## pkg/xrpl/signing (signing.go + ed25519 + secp256k1)

### F-SIGNCURVE-1: secp256k1 verifier accepts non-canonical (high-S) signatures

- **Location:** `pkg/xrpl/signing/secp256k1/signature.go:182-205`, `:34-117`
- **Severity:** High
- **Class:** Crypto misuse — malleability / missing canonicalisation
- **Trigger:** Caller submits a signature whose `s > N/2`.
  `MarshalDER` parses without low-S check; `crypto.VerifySignature` (go-ethereum) does not enforce low-S — that lives in `crypto.ValidateSignatureValues`.
  The library accepts a signature rippled's `STSignature::isValid` will reject.
- **Impact:** A signature from a legitimate signer can be malleated into `(r, N-s)` by any party in the network path.
  `JoinMultisig` accepts both; the assembled `STX_BLOB` fails on-chain submission.
  Silent local-vs-chain divergence in crypto-critical code.
- **Code:**

```go
copy(sigOut.s[32-sLen:], sig[sStart:sEnd])  // no s <= N/2 check
// ...
func (sig *Signature) VerifyHex(hash []byte, pub []byte) bool {
    return crypto.VerifySignature(pub, hash, sig.RS())
}
```

### F-SIGNCURVE-2: `Signature.DER()` panics on zero `r` or `s`

- **Location:** `pkg/xrpl/signing/secp256k1/signature.go:133-170`
- **Severity:** Medium
- **Class:** Panic — slice OOB (`big.Int.Bytes()` returns empty for zero)
- **Trigger:** `*Signature` built via `MarshalRS` with all-zero or zero-r / zero-s bytes, then `DER()` invoked.
- **Impact:** `rb[0]` indexes an empty slice — runtime panic.
  `MarshalRS` does not enforce r,s != 0.

### F-SIGNCURVE-3: Data race on `Signer.value` lazy cache

- **Location:** `pkg/xrpl/signing/signer/signer.go:25-43`
- **Severity:** Medium
- **Class:** Concurrency — data race
- **Trigger:** Two goroutines call `Value()` on the same `*Signer` whose `value` is nil.
  Read/write of `*big.Int` is racy.
- **Impact:** Race-detector failures; torn pointer reads theoretically possible.
  The API does not forbid sharing `*Signer`.

### F-SIGNCURVE-4: secp256k1 helpers nil-deref on malformed `*ecdsa.PrivateKey`

- **Location:** `pkg/xrpl/signing/secp256k1/secp256k1.go:94-112`, `signature.go:246-260`
- **Severity:** Low
- **Class:** Panic — nil deref
- **Trigger:** Caller passes `&ecdsa.PrivateKey{D: d}` without populating `PublicKey.X,Y` to `PrvToAddress`/`PrvToID`/`PrvToPub`.
- **Impact:** `toBytesCompressed` calls `pub.Y.Bit(0)` → nil deref.
  Re-derive X,Y from D or guard.

### F-SIGNCURVE-5: Public-key prefix routing is partially case-sensitive

- **Location:** `pkg/xrpl/signing/signing.go:43-50`
- **Severity:** Low
- **Class:** Logic
- **Trigger:** Mixed-case prefixes like `"Ed"` are rejected, though `hex.DecodeString` is case-insensitive.
- **Impact:** Fails closed (no forgery) but rejects valid rippled-encoded keys.

### F-SIGNCURVE-6: `MarshalDER` does not range-check `r`, `s` against curve order N

- **Location:** `pkg/xrpl/signing/secp256k1/signature.go:34-117`
- **Severity:** Info
- **Class:** Defence-in-depth
- **Trigger:** DER signature with r or s ≥ N (32-byte big-endian above curve order).
- **Impact:** `crypto.VerifySignature` rejects today; recommended to enforce `0 < r,s < N` at the parser boundary.

### F-SIGNCURVE-7: `Recover` does not validate `hash` length

- **Location:** `pkg/xrpl/signing/secp256k1/signature.go:241-243`
- **Severity:** Info
- **Class:** Input validation
- **Impact:** `crypto.SigToPub` returns an error rather than panicking; opaque diagnostic.

### F-SIGNCURVE-8: `Compare` documents but does not enforce `value != nil`

- **Location:** `pkg/xrpl/signing/signer/signer.go:145-147`
- **Severity:** Info
- **Class:** API contract
- **Impact:** Caller missing `Value()` first nil-derefs in `Cmp`.
  All in-tree callers call `Value()` first.

---

## pkg/xrpl/signing (seed + signer + utils)

### F-SIGN-1: Seed-decode error wrap propagates raw input string from base58 layer

- **Location:** `pkg/xrpl/signing/seed/seed.go:35`
- **Severity:** Medium
- **Class:** Error handling / secret leakage
- **Trigger:** Malformed family-seed passed to `DecodeFamilySeed`.
  Underlying `base58.Coder.Decode` returns `fmt.Errorf("bad Base58 string: %s", b)`; `DecodeFamilySeed` wraps with `%w`, preserving the seed string in the chain.
- **Impact:** Family-seed string is private-key-equivalent.
  If callers log the error (common), the seed lands in logs.
  Return a non-wrapping sentinel here.
- **Code:**

```go
raw, err := base58.XRPLCoder.Decode(s)
if err != nil {
    return nil, fmt.Errorf("decoding: %w", err)
}
```

### F-SIGN-2: `Compare` dereferences `s.value` without nil check

- **Location:** `pkg/xrpl/signing/signer/signer.go:146`
- **Severity:** Low
- **Class:** Panic / nil deref
- **Trigger:** External caller invokes `Compare(a, b)` without calling `Value()` first.
- **Impact:** Nil deref panic; documented precondition but exported API.

### F-SIGN-3: `Signer.Value()` lazy cache is not concurrency-safe

- **Location:** `pkg/xrpl/signing/signer/signer.go:25-43`
- **Severity:** Low
- **Class:** Concurrency — data race
- **Impact:** Same as F-SIGNCURVE-3 (duplicate find from a different agent).

### F-SIGN-4: `Signer.Value()` accepts empty `Account` and produces zero `big.Int`

- **Location:** `pkg/xrpl/signing/signer/signer.go:35-42`
- **Severity:** Info
- **Class:** Input boundary
- **Impact:** `base58.Decode("") = []byte{}`; `new(big.Int).SetBytes(nil) = 0`.
  No exploit path (downstream `AccountID.ToBytes` rejects empty), but `Value()` could enforce 25-byte structure.

---

## pkg/merkle

### F-MERKLE-1: No leaf/internal-node domain separation enables second-preimage proof forgery

- **Location:** `pkg/merkle/merkle.go:99-105`, layout at `:53-58`
- **Severity:** High
- **Class:** Crypto — second preimage / weak domain separation
- **Trigger:** Caller builds a tree from raw 32-byte hashes (or `Build(rawHashes, false)`); leaves are not in a separate hash domain from internal nodes. `SortedHashPair(x, y) = keccak256(min || max)` is used uniformly.
- **Impact:** An attacker who chooses a leaf, or who finds an internal-node hash already present, can present that 32-byte value as a "leaf" and `VerifyProof` accepts a forged inclusion proof.
  For non-power-of-2 leaf counts, internal nodes sit at the same indexing level as real leaves (e.g., n=5, `tree[3]` is internal but pairs with leaf `tree[4]`); a prover claiming `tree[3]` is a leaf produces a proof `[tree[4], tree[2]]` that verifies.
  The doc acknowledges this; the API does not enforce.
- **Code:**

```go
func SortedHashPair(x, y common.Hash) common.Hash {
    if bytes.Compare(x[:], y[:]) <= 0 {
        return crypto.Keccak256Hash(x.Bytes(), y.Bytes())
    }
    return crypto.Keccak256Hash(y.Bytes(), x.Bytes())
}
```

### F-MERKLE-2: Proof verification binds only the multiset of sibling hashes, not position

- **Location:** `pkg/merkle/merkle.go:195-202`
- **Severity:** Info (tied to F-MERKLE-1)
- **Class:** Crypto — proof malleability / missing direction bits
- **Impact:** OpenZeppelin `MerkleProof` semantic — safe only if leaves are in a separate hash domain.

### F-MERKLE-3: Panic on out-of-bounds sibling for manually-constructed Tree of invalid length

- **Location:** `pkg/merkle/merkle.go:155-158`
- **Severity:** Medium
- **Class:** Panic — slice OOB
- **Trigger:** `Tree` is an exported type with no constructor enforcement.
  Anyone can build a Tree of arbitrary length, then call `Proof(i)`.
  Traced: `Tree{a,b,c,d}.Proof(1)` → sibling=4 on len-4 slice → panic.
- **Impact:** Process panic from deserialising a `Tree` blob from an untrusted source or constructing from precomputed nodes.
  `LeavesCount` returns nonzero on these inputs.

### F-MERKLE-4: `Leaves()` returns a subslice aliased to the underlying tree

- **Location:** `pkg/merkle/merkle.go:126-133`
- **Severity:** Low
- **Class:** Slice aliasing
- **Impact:** Mutating the returned slice corrupts the tree's leaf level, silently invalidating subsequent `Proof`/`Root` results.
  Concurrent reads race.

### F-MERKLE-5: Silent deduplication of input leaves breaks input-position indexing

- **Location:** `pkg/merkle/merkle.go:50-51`
- **Severity:** Low
- **Class:** Logic — silent input mutation
- **Trigger:** Caller passes a leaf list with duplicates and later uses `Proof(i)` expecting `i` to index into the original input.
- **Impact:** Build silently collapses duplicates then sorts; caller's index has no relation to post-build leaf index.
  Docstring does not mention dedup.

### F-MERKLE-6: `BuildFromHex` silently truncates or pads non-32-byte hex inputs

- **Location:** `pkg/merkle/merkle.go:78-85`
- **Severity:** Low
- **Class:** Input boundary
- **Impact:** `common.HexToHash` silently left-pads short / right-truncates long inputs with no error.
  Two distinct hex inputs sharing the last 32 bytes collide into the same leaf; dedup then collapses them.

### F-MERKLE-7: No proof-length cap in `VerifyProof`

- **Location:** `pkg/merkle/merkle.go:195-202`
- **Severity:** Info
- **Class:** Resource — unbounded CPU
- **Impact:** One keccak256 per proof element with no cap.
  Linear CPU work on untrusted proof; no OOM.

### F-MERKLE-8: `2*n - 1` allocation size has no integer-overflow guard

- **Location:** `pkg/merkle/merkle.go:53`
- **Severity:** Info
- **Class:** Integer overflow (theoretical on 32-bit)
- **Impact:** Trusted-caller boundary, but exported function.

### F-MERKLE-9: `initialHash` parameter is misleading and provides no security property

- **Location:** `pkg/merkle/merkle.go:32-35`, `:87-95`
- **Severity:** Info
- **Class:** API design / crypto misuse risk
- **Impact:** Parameter name invites the misreading that `initialHash=true` is the "safe option" providing domain separation between leaves and internal nodes — it does not.

### F-MERKLE-10: Tree value-receiver methods on exported slice type permit concurrent-mutation data races

- **Location:** `pkg/merkle/merkle.go:26,108,117,126,136,147,170`
- **Severity:** Info
- **Class:** Concurrency
- **Impact:** Standard Go slice race; the intended use is immutable-after-Build, so this is defensive — but `type Tree []common.Hash` is exported.

---

## pkg/voters + pkg/policy

### F-POLICY-1: `NewSet` silently allows `TotalWeight` overflow → wrong voter selected

- **Location:** `pkg/voters/voters.go:50-53`
- **Severity:** Medium
- **Class:** Integer overflow / silent data corruption
- **Trigger:** Comment "sum does not exceed uint16, guaranteed by the smart contract" is enforced in `FromRawBytes` but not in `NewSet`.
  A non-`FromRawBytes` caller (tests, tooling, future internal use) hits unchecked overflow; `vs.TotalWeight += w` wraps; `vs.thresholds` becomes non-monotonic.
- **Impact:** `BinarySearch` returns the wrong index for the same `randomWeight`; `RandomSelectThresholdWeightVoters` picks a voter different from the on-chain protocol's choice.
  Silent disagreement with chain consensus.
- **Code:**

```go
// sum does not exceed uint16 limit, guaranteed by the smart contract
for i, w := range weights {
    vs.thresholds[i] = vs.TotalWeight
    vs.TotalWeight += w
}
```

### F-POLICY-2: `BinarySearch` panics on empty voter set

- **Location:** `pkg/voters/voters.go:131`
- **Severity:** Low
- **Class:** Panic / slice OOB
- **Trigger:** Exported `BinarySearch` invoked on a `Set` whose `voters/weights` are empty; `right := -1` then `vs.thresholds[right]` panics.

### F-POLICY-3: `Hash` panics for inputs of 32 bytes or fewer

- **Location:** `pkg/policy/policy.go:166`
- **Severity:** Low
- **Class:** Panic / slice OOB
- **Trigger:** `len(b) <= 32` → padded buffer of length 32; next access `b[32:64]` is out of range.
  Production callers always pass ≥43 bytes; unreachable today but exported.

### F-POLICY-4: `Storage.Add` notStrict mode lets `rewardEpochID` go backwards

- **Location:** `pkg/policy/storage.go:62-79`
- **Severity:** Low
- **Class:** Broken invariant
- **Trigger:** `notStrict == true` only checks `sp.StartVotingRoundID < prev.StartVotingRoundID`; does not check `sp.RewardEpochID >= prev.RewardEpochID`.
  Caller can insert `{RewardEpochID:3, StartVotingRoundID:200}` after `{RewardEpochID:5, StartVotingRoundID:100}`.
- **Impact:** The "sorted by rewardEpochID (and by startVotingRoundID)" invariant is violated.
  `ForVotingRound`'s `isLast` (last RewardEpochID) misbehaves; downstream ordering consumers break.

### F-POLICY-5: `Voters()` returns the internal slice — caller can mutate Set state

- **Location:** `pkg/voters/voters.go:192-194`
- **Severity:** Low
- **Class:** Slice aliasing
- **Impact:** Mutating reorders / overwrites voters; `VoterDataMap` still describes the original ordering.
  Document immutability or return a clone.

### F-POLICY-6: `selectVoterIndex` would panic via `Mod(_, 0)` if reached with `TotalWeight == 0`

- **Location:** `pkg/voters/voters.go:122-126`
- **Severity:** Info
- **Class:** big.Int division by zero (theoretical)

### F-POLICY-7: `VoterAddress` / `VoterWeight` panic on out-of-range index

- **Location:** `pkg/voters/voters.go:157-164`
- **Severity:** Info
- **Class:** Slice OOB

### F-POLICY-8: `Set` exports `SubmitToSigningAddress`/`VoterDataMap`/`TotalWeight` without sync guarantees

- **Location:** `pkg/voters/voters.go:21-29,66`
- **Severity:** Info
- **Class:** Concurrency
- **Impact:** No clone on adoption; document read-only after construction.

### F-POLICY-9: `InitialHashSeed` panics on oversized `rewardEpochSeed`

- **Location:** `pkg/voters/voters.go:74-86`
- **Severity:** Info
- **Class:** big.Int bounds
- **Impact:** `FillBytes` panics if `BitLen() > 256`.
  Production seed bounded; latent.

### F-POLICY-10: `Storage.RemoveBefore` keeps a policy one round longer than the docstring claims

- **Location:** `pkg/policy/storage.go:94-107`
- **Severity:** Info
- **Class:** Off-by-one (doc vs code)
- **Impact:** Predicate uses `<` but docstring says "removes all signingPolicies that ended strictly before votingRoundID".
  Slight over-retention.

### F-POLICY-11: `Equals` uses non-constant-time `bytes.Equal` on raw policy bytes

- **Location:** `pkg/policy/policy.go:68-73`
- **Severity:** Info
- **Class:** Defensive note (non-issue — public material, no secret)

### F-POLICY-12: `parse.go` `init()` panics on filterer construction

- **Location:** `pkg/policy/parse.go:20-31`
- **Severity:** Info
- **Class:** Panic at package init (intentional fail-fast)

---

## pkg/logger + pkg/events + pkg/payload

### F-LOG-1: Level parse error swallowed; invalid config silently downgrades to InfoLevel

- **Location:** `pkg/logger/logger.go:77`
- **Severity:** Medium
- **Class:** Error handling / misconfiguration
- **Trigger:** `Config.Level` contains an unrecognised string (typo, case, empty).
- **Impact:** `zapcore.ParseLevel` returns the zero value `InfoLevel`.
  Error logged via the just-built INFO logger but execution proceeds; `atom.SetLevel(level)` pins to Info.
  A node operator who intended DEBUG silently runs at INFO.
  Persistent misconfig source because the same code path runs on every `Set(...)` reload.
- **Code:**

```go
level, err := zapcore.ParseLevel(config.Level)
if err != nil {
    sugared.Errorf("Wrong level %s", config.Level)
}
atom.SetLevel(level)
```

### F-LOG-2: `Nop.Panic` and `Nop.Fatal` do not panic / exit — silent control-flow flip

- **Location:** `pkg/logger/nop.go:19,22`
- **Severity:** Medium
- **Class:** Logic / interface contract
- **Trigger:** Caller substitutes `logger.Nop{}` for a real logger and invokes `Panic`/`Fatal` (used in `pkg/database/db.go:115,169,178`, `pkg/database/queries.go:34,41`, `pkg/database/utils.go:104`).
- **Impact:** Real logger's `Panic`/`Fatal` terminate (panic / `os.Exit`); Nop versions return normally.
  Code that relies on unreachable-after semantics (`logger.Fatal("invariant"); return ok`) continues executing in unexpected states.
- **Code:**

```go
func (Nop) Panic(args ...any)         {}
func (Nop) Fatal(args ...any)         {}
```

### F-LOG-3: `SyncFileLogger` reads `current()` three times — torn logger view under `Set`

- **Location:** `pkg/logger/logger.go:88-94`
- **Severity:** Low
- **Class:** Concurrency / atomicity
- **Impact:** Three loads from `loggerPtr` can observe up to three different `*zap.SugaredLogger` instances on the fatal path.
  Defeats the recent `atomic.Pointer` fix on the fatal path.

### F-LOG-4: `events.go` accepts empty-string topic as a real topic

- **Location:** `pkg/events/events.go:45-56`
- **Severity:** Low
- **Class:** Logic / input validation
- **Trigger:** `database.Log.TopicN == ""` (schema has no NOT NULL).
- **Impact:** `common.HexToHash("")` appended as zero-hash topic; sibling converter at `pkg/database/entities.go:73-77` correctly checks both `""` and `"NULL"`.
  Divergence between two converters in the same module.

### F-LOG-5: `common.HexToHash` / `HexToAddress` silently accept malformed input

- **Location:** `pkg/events/events.go:46,49,52,55,62,64`; `pkg/payload/payload.go:67`
- **Severity:** Info
- **Class:** Defensive note
- **Impact:** go-ethereum's hex helpers never error; bad input padded/truncated.
  Trust-boundary note for callers.

### F-LOG-6: `uint(dbLog.LogIndex)` narrows uint64 → uint on 32-bit targets

- **Location:** `pkg/events/events.go:63`
- **Severity:** Info
- **Class:** Integer overflow (32-bit only)

### F-LOG-7: `Panicf`/`Fatalf` call `SyncFileLogger` even when no file logger is configured

- **Location:** `pkg/logger/logger.go:165-176`
- **Severity:** Info
- **Class:** Cosmetic / log noise

### F-LOG-8: `BuildMessageForSigning` ignores oversize cap (defensive only)

- **Location:** `pkg/payload/message.go:61-76`
- **Severity:** Info
- **Class:** Defensive note (all inputs fixed-width)

---

## pkg/priority + pkg/queue + pkg/heapt

### F-PRI-1: `next()` missed-wakeup race — items stranded in heap

- **Location:** `pkg/priority/priority.go:198-239`
- **Severity:** High
- **Class:** Concurrency / data race on event signal (lossy signal channel)
- **Trigger:** Two `Add`/`AddFast` calls arrive while `next()` has drained and unlocked but has not reached the final `select`.
  Each `processIn(Fast)` non-blocking send to the unbuffered `empty` channel falls through `default` and is dropped.
  Items pushed into the heap with no wakeup reaching `next()`.
- **Impact:** Consumer goroutine blocks indefinitely with items present in the heap.
  Throughput drops; in a quiescent system the items stay stranded until another `Add` happens.
  Recent fix `b478145` addresses the both-lanes-empty + ctx case, but the drain-then-block race remains.
- **Code:**

```go
select {
case <-p.regular.empty:
default:
}
p.regular.Unlock()
select {
case <-p.fast.empty:
    ...
case <-p.regular.empty:
    ...
case <-ctx.Done():
```

Fix: hold the lock while observing `Len()==0` and arming the wait, or use `sync.Cond` / buffered (capacity-1) signal channel.

### F-PRI-2: `Add`/`AddFast` block forever after `ctx` cancel

- **Location:** `pkg/priority/priority.go:175,191`
- **Severity:** High
- **Class:** Goroutine leak / no ctx on channel send
- **Trigger:** After the ctx passed to `InitiateAndRun` is cancelled, both `processIn` and `processInFast` return.
  Any subsequent `Add`/`AddFast` sends on `p.in` / `p.inFast` (unbuffered) with no receiver and no ctx-select.
- **Impact:** Permanent goroutine leak.
  The pre-init case was fixed (`TestAddBeforeInitiateAndRun`, M28); the symmetric post-cancel case is unfixed.
  Caller has no way to learn the queue is shut down.
- **Code:**

```go
func (p *PriorityQueue[T, W]) Add(value T, weight W) *Item[Wrapped[T], W] {
    item := &Item[Wrapped[T], W]{ ... }
    p.in <- item  // no select on ctx
    return item
}
```

### F-PRI-3: Data race on `PriorityQueue.bo`

- **Location:** `pkg/priority/priority.go:112-114` (writer), `:321,:351-356` (reader)
- **Severity:** Medium
- **Class:** Data race
- **Trigger:** `SetBackOff` writes `p.bo` without sync; reader runs inside a retry goroutine.
- **Impact:** Race-detector hit; potential torn read on architectures without 8-byte atomic word loads.

### F-PRI-4: Unbounded heap growth — memory exhaustion

- **Location:** `pkg/priority/priority.go:51-67,131,150,328`
- **Severity:** Medium
- **Class:** Input boundary / unbounded resource use
- **Trigger:** `Params` has no `Size` field; producers outpace the consumer.
- **Impact:** OOM.
  Contrast `pkg/queue` which exposes `Size`.

### F-PRI-5: Items silently dropped on ctx cancel between pop and dispatch

- **Location:** `pkg/priority/priority.go:258-268`, `pkg/queue/queue.go:188-197`
- **Severity:** Medium
- **Class:** Logic / item loss
- **Trigger:** After `next()`/`dequeue` returns an item, if `limiter.Wait`, `incrementWorkers`, or `discard` returns ctx error, the item is dropped — no retry, no dead-letter.
- **Impact:** Silent message loss on graceful shutdown.

### F-PRI-6: `next()` discards `heapt.Pop` ok-flag — nil deref if invariant ever breaks

- **Location:** `pkg/priority/priority.go:201,214,228,233`
- **Severity:** Low
- **Class:** Defensive / type-assertion-like silent failure
- **Impact:** Each `Pop` sits behind `Len() > 0` check today; future refactor dropping the lock would panic on `wItem.value.item` nil deref.

### F-PRI-7: `decrementWorkers` panic on imbalance

- **Location:** `pkg/priority/priority.go:303-310`, `pkg/queue/queue.go:310-320`
- **Severity:** Low
- **Class:** Panic in deferred / `logger.Panic`
- **Impact:** Brittle invariant; convert to defensive log + return.

### F-PRI-8: `PriorityQueue` returned by value from constructor — copies mutex

- **Location:** `pkg/priority/priority.go:70,75`, `pkg/queue/queue.go:53`
- **Severity:** Low
- **Class:** Go pitfall — sync types copied by value
- **Impact:** `go vet` `copylocks` would warn if assigned to a value; return `*PriorityQueue` instead.

### F-PRI-9: `DequeueAsync` closure writes outer-scope `err` after caller returned

- **Location:** `pkg/queue/queue.go:203`
- **Severity:** Info
- **Class:** Closure variable capture / shadowing trap

### F-PRI-10: `Length()` reads two channel lengths non-atomically

- **Location:** `pkg/queue/queue.go:327-329`
- **Severity:** Info
- **Class:** Observable race
- **Impact:** Sum can be off by a small amount under concurrent enqueue/dequeue.

### F-PRI-11: `MaxAttempts` semantics doc mismatch

- **Location:** `pkg/priority/priority.go:33`
- **Severity:** Low
- **Class:** Documentation
- **Impact:** Doc says "retries attempted after the first try"; code does "total attempts including first" (test confirms).

### F-PRI-12: `AddValue` on `priority.Queue` is not lock-aware

- **Location:** `pkg/priority/queue.go:73-78`
- **Severity:** Info
- **Class:** API hazard

### F-PRI-13: `handleRetry` re-enqueues directly into the heap, skipping the channel

- **Location:** `pkg/priority/priority.go:320-334`
- **Severity:** Info
- **Class:** Consistency (intentional but worth flagging)

### F-PRI-14: heapt `Remove` lacks the bounds guard `Fix` has

- **Location:** `pkg/heapt/heapt.go:64-71`
- **Severity:** Info
- **Class:** Defensive asymmetry

---

## pkg/convert + pkg/random + pkg/retry + pkg/toml + pkg/storage

### F-UTIL-1: `K(len(s.values))` truncates when K is narrower than size; signed-K negative keys cause OOB panic

- **Location:** `pkg/storage/storage.go:28,45`
- **Severity:** Medium
- **Class:** Integer overflow / slice OOB
- **Trigger:** `Cyclic[K, T]` instantiated with K narrower than `size` (e.g., `int8` + size=200, `uint8` + size=300).
  Traced: K=int8, size=200, key=-200 → `K(200)` overflows to -56.
  `-200 % -56 = -32`.
  `+= K(200) = -56` → -88, still negative; `s.values[-88]` panics.
- **Impact:** Runtime panic from misconfig. For unsigned K, no panic but only `|K(size)|` slots used (silent capacity loss).
- **Code:**

```go
keyMod := key % K(len(s.values))      // K(len) overflows when K narrower than size
if keyMod < 0 {
    keyMod += K(len(s.values))         // still negative if K(len) is negative
}
```

### F-UTIL-2: `Size()` truncates length to K

- **Location:** `pkg/storage/storage.go:22-24`
- **Severity:** Low
- **Class:** Integer overflow
- **Impact:** Callers using `Size()` for indexing arithmetic compute wrong offsets.

### F-UTIL-3: Deprecated `NewCyclic(0)` returns a usable struct that panics on first access

- **Location:** `pkg/storage/storage.go:70-72`
- **Severity:** Medium
- **Class:** Division by zero
- **Impact:** `New` was added with a guard but `NewCyclic` left unsafe.
  First `Store`/`Get` panics on `key % K(0)`.

### F-UTIL-4: `*Cyclic` methods have no nil receiver guard; `New` is documented to return nil

- **Location:** `pkg/storage/storage.go:22,27,43,76-81`
- **Severity:** Low
- **Class:** Nil deref
- **Impact:** Documented contract; defensive note.

### F-UTIL-5: Jitter applied to saturated delay can wrap to a negative `time.Duration`

- **Location:** `pkg/retry/retry.go:110-117`
- **Severity:** Low
- **Class:** Integer overflow on conversion
- **Trigger:** Multiplier-driven saturation makes `d == math.MaxInt64`; jitter then multiplies by up to `1+Jitter`.
  `float64(math.MaxInt64) * 1.5` exceeds int64; conversion saturates implementation-defined; `max(..., 0)` clamps to 0 → immediate retry at extreme delays.

### F-UTIL-6: Multiplier-NaN handled benignly today (`Multiplier > 1` is false for NaN)

- **Location:** `pkg/retry/retry.go:99-106`
- **Severity:** Info
- **Class:** Defensive note for future refactors

### F-UTIL-7: `toml.Read`/`ReadTo` reads with no size limit and follows symlinks

- **Location:** `pkg/toml/toml.go:22-34`
- **Severity:** Info
- **Class:** Input boundaries
- **Impact:** Untrusted `filePath` → memory exhaustion or path traversal.
  Trust-boundary note.

### F-UTIL-8: Strict-mode TOML error reports only the first unknown field

- **Location:** `pkg/toml/toml.go:29-31`
- **Severity:** Info
- **Class:** Error ergonomics

### F-UTIL-9: `pkg/convert` clean

- **Location:** `pkg/convert/*.go`
- **Severity:** Info
- **Impact:** No truncation, no panic vectors visible.
  `BigToUint32Safe`/`BigToUint64Safe` bound-check before narrowing; `BytesToUintN` rejects oversize.

### F-UTIL-10: `pkg/random` clean

- **Location:** `pkg/random/random.go`
- **Severity:** Info
- **Impact:** Uses `crypto/rand` exclusively; bounds-checks negative length.

---

## pkg/xrpl/base58 + hash + address + defs

### F-XLOW-1: `Coder.Decode` has no built-in input-length cap

- **Location:** `pkg/xrpl/base58/base58.go:51`
- **Severity:** Info
- **Class:** Unbounded resource use (exported API)
- **Impact:** big.Int multiplication grows ~quadratically with input length.
  Every in-tree caller caps at 40 chars, but the exported method does not.

### F-XLOW-2: `Field.ID` declares an `err` that is never assigned

- **Location:** `pkg/xrpl/defs/definitions.go:62`
- **Severity:** Low
- **Class:** Style (dead variable, always-nil return)

### F-XLOW-3: `Sha256RipeMD160` doc comment mislabels canonical XRPL derivation as "legacy"

- **Location:** `pkg/xrpl/hash/hash.go:28`
- **Severity:** Info
- **Class:** Misleading doc

### F-XLOW-4: `PubToAddress` error message says "private key" when validating a public key

- **Location:** `pkg/xrpl/address/address.go:85`
- **Severity:** Low
- **Class:** Error message correctness

### F-XLOW-5: `IDToAddress` is an unchecked variant of `Address`

- **Location:** `pkg/xrpl/address/address.go:64`
- **Severity:** Info
- **Class:** API footgun (documented bypass)
- **Impact:** Produces a structurally valid base58 string from a non-20-byte id.

---

## pkg/xrpl/encoding

### F-ENC-1: AccountID decoder ignores VL length prefix → decoder/encoder ambiguity

- **Location:** `pkg/xrpl/encoding/types/accountid.go:31`
- **Severity:** High
- **Class:** Logic / silent data corruption / consensus mismatch
- **Trigger:** Wire AccountID field with VL prefix ≠ 20.
  `coder.decodeNext` calls `lengthDecode` and passes the result as `length` to `AccountID.ToJSON`, which discards it (`_`) and always reads exactly 20 bytes.
- **Impact:** With `VL < 20` over-reads into adjacent fields; with `VL > 20` surplus bytes are reinterpreted as the next field ID. rippled rejects `VL != 20` for AccountID; same payload encodes two ways accepted here, smuggling an alternate decoding under signed bytes.
- **Code:**

```go
func (*accountID) ToJSON(b *bytes.Buffer, _ int) (any, error) {
    const l = 20
    value := make([]byte, l)
    n, err := b.Read(value)
```

### F-ENC-2: tokenType1 (firstByte `0xA0`) accepted; stray `0x20` bit corrupts IOU exponent

- **Location:** `pkg/xrpl/encoding/types/amount.go:100,380-401,318-352`
- **Severity:** High
- **Class:** Logic / silent data corruption / encoding ambiguity
- **Trigger:** Wire amount with first byte `0xA0` (bits `0x80` and `0x20`).
  `Amount.ToJSON` routes both `tokenType0 (0x80)` and `tokenType1 (0xA0)` to `tokenToJSON`.
  `tokenToJSON` masks only `0x80` (`a[0] = firstByte & signedValueBitMask = 0x20`).
  `deserializeTokenAmount` extracts the exponent via `exponentMask = 0xff << 54`; the stray `0x20` lands inside that mask → normalised exponent acquires +128 (real exponent +31).
- **Impact:** Non-canonical IOU encoding accepted with corrupted value while rippled rejects it (reserved bit).
  Same signed bytes → different field values Flare-Go vs rippled — transaction-value malleability.
- **Code:**

```go
switch amountType {
case xrpType:
    return xrpToJSON(firstByte, b)
case tokenType0, tokenType1:
    return tokenToJSON(firstByte, b)
...
a[0] = firstByte & signedValueBitMask // remove "not XRP bit"  -- 0x20 stays in
```

### F-ENC-3: IOU exponent/significand canonical ranges not enforced on decode

- **Location:** `pkg/xrpl/encoding/types/amount.go:325-352`
- **Severity:** Medium
- **Class:** Logic / consensus drift / encoding ambiguity
- **Trigger:** IOU amount with exponent outside `[-96, 80]` or significand outside `[10^15, 10^16-1]` (and not zero).
  Comment in source acknowledges the missing enforcement.
- **Impact:** Non-canonical IOU encodings parse here but rippled rejects them; multiple wire forms map to the same numeric value with no canonicalisation.

### F-ENC-4: MPT zero with sign bit clear (firstByte `0x20`, mantissa 0) accepted as canonical

- **Location:** `pkg/xrpl/encoding/types/amount.go:441-472`
- **Severity:** Low
- **Class:** Consensus drift
- **Trigger:** MPT amount with `firstByte == 0x20` and zero mantissa.
- **Impact:** Canonical MPT zero is `firstByte == 0x60`; two byte patterns decode to the same value with no rejection.

### F-ENC-5: `deserializeCurrency` accepts non-ASCII ISO codes

- **Location:** `pkg/xrpl/encoding/types/currency.go:84-113`
- **Severity:** Low
- **Class:** Encoding ambiguity
- **Trigger:** Wire currency `c[12:15]` containing non-printable bytes.
- **Impact:** Round-trips as a binary-garbage 3-byte string here but the encoder side (`serializeStandardCode`) is strict — breaks `Decode∘Encode == identity`.

### F-ENC-6: PathSet `0x00` step accepted as delimiter at step level

- **Location:** `pkg/xrpl/encoding/types/pathset.go:227-235`
- **Severity:** Info
- **Class:** Step-flag/delimiter space conflation
- **Impact:** Bounded by downstream `l<1` empty-path check; no exploit.

### F-ENC-7: `convertInt64` uint64/float boundary `> 1<<63` lets `1<<63` slip through

- **Location:** `pkg/xrpl/encoding/types/uint.go:216,224,232`
- **Severity:** Low
- **Class:** Off-by-one
- **Trigger:** `uint64(1<<63)` exactly.
  Check is strictly greater; equality passes.
  `int64(value)` wraps to MinInt64.
- **Impact:** Bounded today by downstream `< 0` and `>= 1<<N` checks.
  Future callers without those checks would silently wrap.
  Fix to `>= 1<<63`.

### F-ENC-8: XChainBridge encode error messages reference wrong variable and wrap nil error

- **Location:** `pkg/xrpl/encoding/types/xchainbridge.go:46-47,71-73`
- **Severity:** Info
- **Class:** Error handling
- **Impact:** Defensive paths only — `address.ID` always returns either error or 20 bytes; `>= 256` branch unreachable.

### F-ENC-9: Top-level `Decode` accepts empty blob as empty object

- **Location:** `pkg/xrpl/encoding/types/coder.go:355-369`
- **Severity:** Info
- **Class:** API contract
- **Impact:** `Encode` refuses empty objects (l.42); top-level `Decode` accepts.
  Breaks `Decode∘Encode == identity`.

### F-ENC-10: `serializeTokenValue` error message misleading for huge exponents

- **Location:** `pkg/xrpl/encoding/types/amount.go:261-273`
- **Severity:** Info
- **Class:** Error message ergonomics

### F-ENC-11: `tokenToJSON` uses `b.Read` for fixed-size fields without `io.ReadFull`

- **Location:** `pkg/xrpl/encoding/types/amount.go:389,406,423` (and many other decoders)
- **Severity:** Info
- **Class:** Short-read handling
- **Impact:** Safe today (source is `*bytes.Buffer`); use `io.ReadFull` for robustness against future stream wrappers.

---

## pkg/xrpl/transactions + check + aggregator

### F-AGG-1: `Finalize` fails when weighted quorum is reached by fewer signers than the quorum threshold

- **Location:** `pkg/xrpl/aggregator/aggregator.go:180,194-211`
- **Severity:** High
- **Class:** Logic / aggregator threshold off-by-one / weight handling
- **Trigger:** `Account.Quorum` is a **summed-weight threshold** (per the doc, l.18-29), but `Finalize` passes it to `sort` as the **count** of required signers.
  When signers reach quorum _by weight_ but their _count_ is less than `a.Quorum` (e.g. `Quorum=2`, single signer with `SignerWeight=2`), `sort` allocates `make([]*Signer, 2)`, fills only one, and bails with `"quorum not reached"`.
- **Impact:** Finalize is unreachable for any weighted multisig where any single signer's weight is ≥ quorum, or any subset reaches quorum by weight without count.
  Funds-flow stuck.
- **Code:**

```go
s, err := sort(tx.signers, int(a.Quorum))
// ...
out := make([]*Signer, quorum)
i := 0
for j := range sig { out[i] = sig[j]; i++; if i == quorum { break } }
if i != quorum { return nil, errors.New("quorum not reached") }
```

### F-AGG-2: Non-deterministic signer subset selection in `Finalize`

- **Location:** `pkg/xrpl/aggregator/aggregator.go:194-211`
- **Severity:** Medium
- **Class:** Map iteration order
- **Trigger:** `len(tx.signers) > a.Quorum`.
  `sort` iterates `tx.signers` (a Go map) and stops at `quorum` entries.
- **Impact:** Same transaction can produce different on-wire blobs / different rippled tx IDs across re-finalizations; complicates idempotency.

### F-AGG-3: Data race on `Account.txs`, `tx.signers`, `tx.quorumReached`

- **Location:** `pkg/xrpl/aggregator/aggregator.go:47-157,170-191`
- **Severity:** High
- **Class:** Concurrency — data race
- **Trigger:** `Account` is designed to collect signatures arriving from multiple sources.
  `AddSignatures` mutates `a.txs[id]`, `tx.signers[s.Account]`, and `tx.quorumReached`; `Finalize` reads all three.
  No mutex.
- **Impact:** Concurrent calls produce undefined behaviour: map-concurrent-write panic, missed signatures, double-dispatch of the quorum-reached transition, torn `quorumReached` reads.

### F-AGG-4: Nil-map panic on first `AddSignatures` call when `Account` is zero-value-constructed

- **Location:** `pkg/xrpl/aggregator/aggregator.go:24-29,149`
- **Severity:** Medium
- **Class:** Panic — nil map write
- **Trigger:** No exported constructor; `txs` cannot be set externally; first `AddSignatures` panics on `a.txs[identifier] = tx`.
- **Impact:** Forward-trap; no in-tree callers today, but reachable for any external library consumer.

### F-AGG-5: `Info` uses `context.TODO()` and disables HTTP timeout

- **Location:** `pkg/xrpl/check/check.go:91-112`
- **Severity:** Low
- **Class:** Resource / deadline
- **Impact:** A hung rippled blocks the calling goroutine forever; caller cannot cancel.

### F-AGG-6: `transaction.id` field written but never read

- **Location:** `pkg/xrpl/aggregator/aggregator.go:31-36,87`
- **Severity:** Info
- **Class:** Dead state

### F-AGG-7: `int(a.Quorum)` cast bounded only by uint→int range

- **Location:** `pkg/xrpl/aggregator/aggregator.go:180`
- **Severity:** Info
- **Class:** Integer overflow (defensive — bounded by XRPL signer max)

### F-AGG-8: `Memo` pointer-receiver methods nil-deref on a nil `*Memo`

- **Location:** `pkg/xrpl/transactions/basic.go:51-72`
- **Severity:** Info
- **Class:** Nil deref (no in-tree caller; exported)

---

## pkg/tee/signer + op + instruction

### F-TEE-1: Empty `APIKeyName` + empty header allows unauthenticated requests

- **Location:** `pkg/tee/signer/signer.go:230-237`
- **Severity:** Medium
- **Class:** Crypto / authn bypass via misconfig
- **Trigger:** Operator deploys with `Config.APIKeyName == ""` or includes an empty string in `Config.APIKeys`.
  `newAPIKeys` does not reject either; `r.Header.Get("")` returns empty for every request; HMAC-SHA256(secret,"") equals the digest stored for the empty key.
- **Impact:** Anyone with network access to the (loopback-only) signer can sign arbitrary hashes / decrypt arbitrary ciphertexts.
- **Code:**

```go
func newAPIKeys(cfg Config) (apiKeys, error) {
    secret := make([]byte, 32)
    if _, err := rand.Read(secret); err != nil { ... }
    digests := make([][]byte, len(cfg.APIKeys))
    for j, k := range cfg.APIKeys { digests[j] = hmacKey(secret, k) }
    return apiKeys{name: cfg.APIKeyName, secret: secret, digests: digests}, nil
}
```

### F-TEE-2: `signer.New` does not nil-check the private key

- **Location:** `pkg/tee/signer/signer.go:118-149`
- **Severity:** Medium
- **Class:** Nil deref / panic from misconfig
- **Trigger:** Caller passes `prv == nil` (e.g. uninitialised wallet, failed key load whose error was not propagated).
  `New` succeeds; first POST to `/id` panics at `&prv.PublicKey`, first `/sign` panics inside `crypto.Sign`.

### F-TEE-3: `signHandler` reflects internal error strings to client

- **Location:** `pkg/tee/signer/signer.go:279-291`, `:421`, `:326`
- **Severity:** Low
- **Class:** Error message leaking internal detail
- **Impact:** Leaks library internals to the (authenticated, loopback) client.
  Violates GOAI.md §Errors "never expose internal error details through a server API".

### F-TEE-4: `decryptHandler` does not enforce `Content-Type`

- **Location:** `pkg/tee/signer/signer.go:382-404`
- **Severity:** Info
- **Class:** API surface inconsistency
- **Impact:** `signHandler` rejects non-JSON; `decryptHandler` accepts any content type.
  Behavioural inconsistency between two endpoints sharing the same trust model.

### F-TEE-5: `Content-Type` appended rather than set

- **Location:** `pkg/tee/signer/signer.go:341`
- **Severity:** Info
- **Class:** HTTP correctness / defensive

### F-TEE-6: Private key and API-key HMAC secret are never zeroized

- **Location:** `pkg/tee/signer/signer.go` (whole file)
- **Severity:** Info
- **Class:** Key zeroization (by design for TEE-resident signer)

### F-TEE-7: `signHandler` is an unbounded EIP-191 signing oracle

- **Location:** `pkg/tee/signer/signer.go:250-294,297-316`
- **Severity:** Info (deferred; documented loopback + API-key mitigation)
- **Class:** Crypto / oracle design
- **Impact:** No per-request domain separation.
  Mitigated by host trust boundary.

### F-TEE-8: `Data.HashFixed` and `HashForSigning` do not require `Validate`

- **Location:** `pkg/tee/instruction/instruction.go:92-99,158-164`
- **Severity:** Info
- **Class:** Input boundaries
- **Impact:** DoS hardening; recommend `HashFixed` call `Validate` first.

### F-TEE-9: `RecoverSignersPubKey` accepts non-canonical (high-S) signatures

- **Location:** `pkg/tee/instruction/instruction.go:180-188`
- **Severity:** Low
- **Class:** Crypto / signature malleability
- **Trigger:** `crypto.SigToPub` runs ECDSA recovery without enforcing low-S; `(r,s,v)` and `(r,n-s,v^1)` recover the same pubkey.
- **Impact:** If downstream uses the signature as a uniqueness/replay key, replay protection breaks.

### F-TEE-10: `op.IsValid` blanket-accepts unknown non-system, non-F-prefixed types

- **Location:** `pkg/tee/op/op.go:132-138`
- **Severity:** Info
- **Class:** Input boundary / opcode dispatch
- **Impact:** Documented behaviour; downstream dispatcher must be aware.

### F-TEE-11: `Type.Hash`/`Command.Hash` collision for >32-byte names

- **Location:** `pkg/tee/op/op.go:104-106,127-129`
- **Severity:** Info
- **Class:** Domain separation (latent; current constants short)

### F-TEE-12: `Validate` allocates map for cosigner dup check

- **Location:** `pkg/tee/instruction/instruction.go:65-74`
- **Severity:** Info
- **Class:** Minor; not exploitable.

### F-TEE-13: No replay-protection / freshness binding on signed instruction hashes

- **Location:** `pkg/tee/instruction/instruction.go:117-164`
- **Severity:** Info (cross-cutting)
- **Class:** Crypto / nonce / replay
- **Impact:** No monotonic counter, attestation hash, or chain ID covered in `HashForSigning`.
  Replay across networks/TEE instances depends on `TeeID`+`InstructionID` uniqueness upstream.

---

## pkg/tee/xrpl + attestation + structs

### F-TEEX-1: `ParseFeeEntry` panics on negative `try`

- **Location:** `pkg/tee/xrpl/payment.go:235-249`
- **Severity:** Medium
- **Class:** Slice OOB / integer overflow on length fields
- **Trigger:** Caller passes negative `try`; `len(schedule) < (try+1)*4` evaluates with non-positive RHS so guard passes; `schedule[try*4:(try+1)*4]` panics on negative-bound slicing.
  `try` near `MaxInt` triggers the same via `(try+1)*4` overflow.
- **Impact:** `PaymentTxFromInstruction` forwards caller-supplied `try` directly; retry loops / RPC params / deserialised counters that drive this negative crash the TEE process.
- **Code:**

```go
if len(schedule) < (try+1)*4 {
    return ScheduledFee{}, errors.New("try beyond schedule length")
}
entry := schedule[try*4 : (try+1)*4]
```

### F-TEEX-2: `ScheduledFee.Fee` panics on nil `max`

- **Location:** `pkg/tee/xrpl/payment.go:208-213`
- **Severity:** Low
- **Class:** Nil deref / big.Int
- **Impact:** `PaymentTxFromInstruction` validates `i.Amount != nil` but not `i.MaxFee != nil`.
  ABI-decoded uint256 won't be nil, but direct callers will.

### F-TEEX-3: `EATNonce` policy default silently disables replay binding

- **Location:** `pkg/tee/attestation/googlecloud/google_cloud.go:257-261`
- **Severity:** Info
- **Class:** Logic / replay (nonce binding)
- **Trigger:** Deployment with empty `Policy.EATNonce` skips the `slices.Contains` check entirely.
- **Impact:** Cross-deployment replay protection is opt-in despite the rest of the policy being mandatory.
  Misconfig is hard to spot in tests.

### F-TEEX-4: Other production attestation claims not exposed by `Apply`

- **Location:** `pkg/tee/attestation/googlecloud/google_cloud.go:236-263`
- **Severity:** Info
- **Class:** Missing claim treated as valid
- **Impact:** `SupportAttributes` and `GoogleServiceAccounts` are parsed but never validated by `Apply`.
  Callers must read by hand.

### F-TEEX-5: Typo in x5c header error formatting (`% v` instead of `%v`)

- **Location:** `pkg/tee/attestation/googlecloud/google_cloud.go:345`
- **Severity:** Info
- **Class:** Style

---

## pkg/contracts/contractregistry + module meta

### F-META-1: Unchecked float64 → int16 narrowing on field `nth`

- **Location:** `pkg/xrpl/defs/main/main.go:44`
- **Severity:** Medium
- **Class:** Integer overflow on conversions
- **Trigger:** Any `nth` value in `definitions.json` outside `[-32768, 32767]`.
  The XRPL definitions file is an external input; future protocol changes / hostile copy could exceed.
- **Impact:** Silent wraparound on the generated `defs/autogen.go` table; wrong `Nth` produces wrong field tags on every signed XRPL transaction.
  No runtime error, no compile error — just bad signatures.
- **Code:**

```go
Nth: int16(math.Round(values["nth"].(float64))),
```

### F-META-2: Generator truncates output before formatting can fail

- **Location:** `pkg/xrpl/defs/main/main.go:114-123`
- **Severity:** Medium
- **Class:** Failure-recovery
- **Trigger:** `format.Source` returns an error after `O_TRUNC` has zeroed `autogen.go`.
- **Impact:** Zero-byte `autogen.go` on disk; package will not compile until rerun succeeds.
  Write to temp + rename on success.

### F-META-3: Non-deterministic generator output (map iteration order)

- **Location:** `pkg/xrpl/defs/main/main.go:90-100`
- **Severity:** Low
- **Class:** Reproducibility
- **Impact:** `autogen.go` diff churns between runs even when input unchanged.
  Defeats CI "in-sync?" check.

### F-META-4: Slice OOB on malformed Fields entry

- **Location:** `pkg/xrpl/defs/main/main.go:27-35`
- **Severity:** Low
- **Class:** Slice OOB on parsing
- **Trigger:** `FIELDS` row with fewer than two elements (upstream schema change).
- **Impact:** Index-out-of-range panic before the type-assertion guard.

### F-META-5: Input file handle leaked

- **Location:** `pkg/xrpl/defs/main/main.go:68-78`
- **Severity:** Info
- **Class:** Resource leak (trivial; short-lived `main`)

### F-META-6: `abi_update.sh` missing strict mode

- **Location:** `abi_update.sh:1`
- **Severity:** Low
- **Class:** Error handling / shell hardening
- **Trigger:** Any failure in `jq`, `mkdir`, `tr` inside loops.
- **Impact:** `jq '.abi' "$input_file"` on a JSON file lacking `.abi` writes literal `null` and exits 0; script reports success but leaves corrupted output.
  No shell-injection vector (no positional args read).

### F-META-7: `abi_update.sh` relies on undeclared PWD

- **Location:** `abi_update.sh:45,62,72`
- **Severity:** Info
- **Class:** Shell hardening
- **Impact:** Relative paths; running from a subdirectory silently fails.

### F-META-8: Standard abigen `out[0]` index without bounds check

- **Location:** `pkg/contracts/contractregistry/contractregistry.go:194,225-226,257,288,319,350`
- **Severity:** Info
- **Class:** Slice OOB (inherent to abigen output; file is generated boilerplate)

### F-META-9: Scope description vs. file header mismatch

- **Location:** `pkg/contracts/contractregistry/contractregistry.go:1-2`
- **Severity:** Info
- **Class:** Audit hygiene
- **Impact:** File is abigen-generated (`// Code generated - DO NOT EDIT.`); no handwritten contract code lives in `pkg/contracts/contractregistry/`.

### F-META-10: go.mod / linter Go-version skew

- **Location:** `go.mod:3`, `infolint.yml:3`
- **Severity:** Info
- **Class:** Configuration drift
- **Impact:** `go.mod` declares 1.25.1; `infolint.yml` pins 1.24.
  Likely intentional cross-compat lint; confirm.

---

# Appendix: methodology details

- 16 parallel `general-purpose` review agents, each given an enumerated bug-class list and a structured finding format.
- Each agent was explicitly instructed to **avoid reading `AUDIT_REPORT.md`** to prevent anchoring on the prior pass.
- All `autogen.go` files were skipped (`pkg/contracts/**/autogen.go`, `pkg/xrpl/defs/autogen.go` — the latter is a data table, audited only at its consumer boundary).
- Test files (`*_test.go`) were read only for context; no findings filed against them.

Headline gaps after Pass 2 — recommend prioritising:

1. **Merkle domain separation** (F-MERKLE-1) — protocol-level safety relies on caller discipline.
2. **XRPL encoder/decoder consensus drift** (F-ENC-1, F-ENC-2) — same signed bytes decode differently here vs rippled.
3. **secp256k1 low-S enforcement** (F-SIGNCURVE-1) — verify-side malleability gap.
4. **Database unbounded reads** (F-DB-1, F-DB-2) — OOM via public API of the shared lib.
5. **Priority queue concurrency** (F-PRI-1, F-PRI-2) — missed-wakeup + post-cancel hang.
6. **Aggregator weight vs count + data race** (F-AGG-1, F-AGG-3) — weighted multisig unreachable + concurrent-collection bugs.
