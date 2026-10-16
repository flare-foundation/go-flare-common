# Envelope refusal vectors

Every decoder of the CSP proposal envelope (`TeeStructs.CspEnvelope`) — in Go,
Python, Solidity or anything else — must **refuse** each `.hex` file here.
The envelope has no version field; a decoder accepts only the canonical
encoding (decode, encode again, require the same bytes), and these are the
shapes that rule exists to catch.

Format: lowercase hex, whitespace insignificant (one 32-byte word per line).
Each file is a variant of `../envelope_layout.hex`, which every decoder must
**accept** and decode to the fixture in `envelope_test.go` (`layoutFixture`).

| File | What it is | Why it must be refused |
|---|---|---|
| `trailing_byte.hex` | the layout vector plus one `0x00` byte | two byte strings would decode to one message |
| `trailing_bytes.hex` | the layout vector plus 31 `0x00` bytes (not a whole word) | as above |
| `appended_word.hex` | the layout vector plus one zero 32-byte word | as above; a lenient decoder ignores a whole trailing word |
| `appended_field.hex` | the fixture encoded as the tuple with one more `bytes32` (`0x…ff`) after `accountHash` (`cast abi-encode`) | a layout with an appended field: an old decoder reads the head and ignores the tail, and only the re-encoding (whose offsets are 32 lower) shows it |
| `with_version.hex` | the fixture encoded with a leading `uint16 version = 3` (the earlier layout; `cast abi-encode`) | the version word is gone; the old layout must not be read as the new one |

The layout vector and the two `cast`-encoded refusals were produced by
Foundry's `cast abi-encode`, independently of go-ethereum. `TestDecodeRefusesSharedVectors`
in `envelope_test.go` loads every `.hex` file here.
