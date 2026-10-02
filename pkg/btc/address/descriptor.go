package address

import (
	"errors"
	"fmt"
	"strings"
)

// BIP-380 output script descriptor checksum.
//
// A descriptor on the wire is `<expression>#<8-char-checksum>`, where the
// checksum is computed over the expression bytes using the algorithm in
// [BIP-380 § "Checksum"](https://github.com/bitcoin/bips/blob/master/bip-0380.mediawiki#checksum).
//
// We use the checksum to validate the `scriptDescriptor` field carried by
// the BtcAccountConfigured attestation. The format is what `bitcoind`'s
// `getdescriptorinfo` emits, so descriptors round-trip without modification.

// descriptorInputCharset is the 90-character alphabet from which a
// descriptor expression may be built. The index of each character within
// this string is its symbol (mod 32 for the checksum polynomial; the high
// bits encode a character class).
const descriptorInputCharset = "0123456789()[],'/*abcdefgh@:$%{}" +
	"IJKLMNOPQRSTUVWXYZ&+-.;<=>?!^_|~" +
	"ijklmnopqrstuvwxyzABCDEFGH`#\"\\ "

// descriptorChecksumCharset is the 32-character bech32 alphabet used to
// encode the 40-bit checksum as 8 chars.
const descriptorChecksumCharset = "qpzry9x8gf2tvdw0s3jn54khce6mua7l"

// descriptorPolymod is the BIP-380 polynomial step. It mirrors the bech32
// polymod but with a different generator tuned for the 40-bit checksum.
func descriptorPolymod(c uint64, val uint64) uint64 {
	c0 := c >> 35
	c = ((c & 0x7ffffffff) << 5) ^ val
	if c0&1 != 0 {
		c ^= 0xf5dee51989
	}
	if c0&2 != 0 {
		c ^= 0xa9fdca3312
	}
	if c0&4 != 0 {
		c ^= 0x1bab10e32d
	}
	if c0&8 != 0 {
		c ^= 0x3706b1677a
	}
	if c0&16 != 0 {
		c ^= 0x644d626ffd
	}
	return c
}

// ComputeDescriptorChecksum returns the 8-character BIP-380 checksum for
// the given descriptor expression (the part before `#`).
//
// Returns an error if expr contains a character outside the input
// alphabet.
func ComputeDescriptorChecksum(expr string) (string, error) {
	var c uint64 = 1
	var cls uint64 = 0
	var clscount uint64 = 0
	for i := 0; i < len(expr); i++ {
		ch := expr[i]
		pos := strings.IndexByte(descriptorInputCharset, ch)
		if pos < 0 {
			return "", fmt.Errorf("descriptor: invalid character %q at byte %d", ch, i)
		}
		c = descriptorPolymod(c, uint64(pos)&31)
		cls = cls*3 + uint64(pos)>>5
		clscount++
		if clscount == 3 {
			c = descriptorPolymod(c, cls)
			cls = 0
			clscount = 0
		}
	}
	if clscount > 0 {
		c = descriptorPolymod(c, cls)
	}
	// Append 8 zero symbols to ensure the polynomial absorbs the checksum
	// position; then XOR with 1 to make the empty input checksum nonzero.
	for range 8 {
		c = descriptorPolymod(c, 0)
	}
	c ^= 1

	out := make([]byte, 8)
	for j := range 8 {
		out[j] = descriptorChecksumCharset[(c>>(5*(7-j)))&31]
	}
	return string(out), nil
}

// VerifyDescriptorChecksum splits a full descriptor of the form
// `<expression>#<checksum>` and verifies the checksum. It returns the
// expression on success.
//
// Rejects:
//   - missing `#` separator,
//   - checksum length != 8,
//   - characters outside the BIP-380 input alphabet (expression) or
//     bech32 checksum alphabet (suffix),
//   - mismatched recomputed checksum.
func VerifyDescriptorChecksum(descriptor string) (string, error) {
	expr, got, found := strings.Cut(descriptor, "#")
	if !found {
		return "", errors.New("descriptor: missing '#' checksum separator")
	}
	if len(got) != 8 {
		return "", fmt.Errorf("descriptor: checksum length %d, want 8", len(got))
	}
	for i := 0; i < len(got); i++ {
		if strings.IndexByte(descriptorChecksumCharset, got[i]) < 0 {
			return "", fmt.Errorf("descriptor: checksum character %q at position %d is outside bech32 alphabet", got[i], i)
		}
	}
	want, err := ComputeDescriptorChecksum(expr)
	if err != nil {
		return "", err
	}
	if want != got {
		return "", fmt.Errorf("descriptor: checksum mismatch, want %q got %q", want, got)
	}
	return expr, nil
}
