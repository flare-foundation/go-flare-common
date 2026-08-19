package csp

import "testing"

// Decode is handed bytes fetched from the open DAL by txid, with nothing binding
// them to a hash. Whatever it does with malformed input, it must not panic: a
// remote-triggerable crash in an attestation verifier stops that data provider
// voting, which is a liveness attack anyone can run.
func FuzzDecodeDoesNotPanic(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{0x00})
	f.Add(make([]byte, 32))
	f.Add(make([]byte, 1024))
	f.Fuzz(func(t *testing.T, data []byte) {
		_, _ = Decode(data)
	})
}
