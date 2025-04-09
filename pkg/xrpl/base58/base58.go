package base58

import (
	"fmt"
	"math/big"
	"strings"
	"unicode/utf8"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
)

const AlphabetXRPL = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz" // alphabet used for base 58 encoding by XRPL: https://xrpl.org/docs/references/protocol/data-types/base58-encodings

// Purloined from https://github.com/btcsuite/btcd/blob/master/btcutil/base58/base58.go
var (
	bigRadix10 = big.NewInt(58 * 58 * 58 * 58 * 58 * 58 * 58 * 58 * 58 * 58) // 58^10
	bigRadix   = [...]*big.Int{
		big.NewInt(0),
		big.NewInt(58),
		big.NewInt(58 * 58),
		big.NewInt(58 * 58 * 58),
		big.NewInt(58 * 58 * 58 * 58),
		big.NewInt(58 * 58 * 58 * 58 * 58),
		big.NewInt(58 * 58 * 58 * 58 * 58 * 58),
		big.NewInt(58 * 58 * 58 * 58 * 58 * 58 * 58),
		big.NewInt(58 * 58 * 58 * 58 * 58 * 58 * 58 * 58),
		big.NewInt(58 * 58 * 58 * 58 * 58 * 58 * 58 * 58 * 58),
		bigRadix10,
	}
)

var XRPLCoder *Coder

func init() {
	var err error
	XRPLCoder, err = NewCoder(AlphabetXRPL)

	if err != nil {
		logger.Panicf("invalid xrpl alphabet: %v", err)
	}
}

// Coder holds a base58 alphabet.
type Coder struct {
	alphabet string
}

// Decode decodes base58 encoded string to byte slice.
func (c *Coder) Decode(s string) ([]byte, error) {
	return decode(s, c.alphabet)
}

// Encode encodes byte slice to base58 string.
func (c *Coder) Encode(b []byte) string {
	return encode(b, c.alphabet)
}

// NewCoder returns a coder for base58 alphabet. It checks that the alphabet is valid.
func NewCoder(alphabet string) (*Coder, error) {
	if len(alphabet) != 58 {
		return nil, fmt.Errorf("alphabet %s is not 58 characters long", alphabet)
	}
	if utf8.RuneCountInString(alphabet) != 58 {
		return nil, fmt.Errorf("alphabet %s has nonstandard runes", alphabet)
	}

	if err := uniqueCharacters(alphabet); err != nil {
		return nil, fmt.Errorf("invalid alphabet %s: %v", alphabet, err)
	}

	return &Coder{alphabet: alphabet}, nil
}

// uniqueCharacters errors if there any of the runes in string s appears more than once.
func uniqueCharacters(s string) error {
	check := make(map[rune]bool)
	for _, c := range s {
		_, exists := check[c]
		if exists {
			return fmt.Errorf("character %s appears multiple times", string(c))
		}
		check[c] = true
	}
	return nil
}

// decode decodes a base58 string in alphabet to a byte slice.
func decode(b, alphabet string) ([]byte, error) {
	answer := big.NewInt(0)
	scratch := new(big.Int)

	// Calculating with big.Int is slow for each iteration.
	//    x += b58[b[i]] * j
	//    j *= 58
	//
	// Instead we can try to do as much calculations on int64.
	// We can represent a 10 digit base58 number using an int64.
	//
	// Hence we'll try to convert 10, base58 digits at a time.
	// The rough idea is to calculate `t`, such that:
	//
	//   t := b58[b[i+9]] * 58^9 ... + b58[b[i+1]] * 58^1 + b58[b[i]] * 58^0
	//   x *= 58^10
	//   x += t
	//
	// Of course, in addition, we'll need to handle boundary condition when `b` is not multiple of 58^10.
	// In that case we'll use the bigRadix[n] lookup for the appropriate power.
	for t := b; len(t) > 0; {
		n := min(len(t), 10)

		total := uint64(0)
		for _, v := range t[:n] {
			tmp := strings.Index(alphabet, string(v))
			if tmp == -1 {
				return nil, fmt.Errorf("bad Base58 string: %s", b)
			}
			total = total*58 + uint64(tmp)
		}

		answer.Mul(answer, bigRadix[n])
		scratch.SetUint64(total)
		answer.Add(answer, scratch)

		t = t[n:]
	}

	tmpval := answer.Bytes()

	leadingZeros := 0
	for numZeros := range len(b) {
		if b[numZeros] != alphabet[0] {
			break
		}
		leadingZeros++
	}

	flen := leadingZeros + len(tmpval)
	val := make([]byte, flen)
	copy(val[leadingZeros:], tmpval)
	return val, nil
}

// encode encodes a byte slice to a base58 string in a given alphabet.
func encode(b []byte, alphabet string) string {
	x := new(big.Int)
	x.SetBytes(b)

	// maximum length of output is log58(2^(8*len(b))) == len(b) * 8 / log(58)
	maxlen := int(float64(len(b))*1.365658237309761) + 1
	answer := make([]byte, 0, maxlen)
	mod := new(big.Int)
	for x.Sign() > 0 {
		// Calculating with big.Int is slow for each iteration.
		//    x, mod = x / 58, x % 58
		//
		// Instead we can try to do as much calculations on int64.
		//    x, mod = x / 58^10, x % 58^10
		//
		// Which will give us mod, which is 10 digit base58 number.
		// We'll loop that 10 times to convert to the answer.

		x.DivMod(x, bigRadix10, mod)
		if x.Sign() == 0 {
			// When x = 0, we need to ensure we don't add any extra zeros.
			m := mod.Int64()
			for m > 0 {
				answer = append(answer, alphabet[m%58])
				m /= 58
			}
		} else {
			m := mod.Int64()
			for range 10 {
				answer = append(answer, alphabet[m%58])
				m /= 58
			}
		}
	}

	// leading zero bytes
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, alphabet[0])
	}

	// reverse
	alen := len(answer)
	for i := range alen / 2 {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}
	return string(answer)
}
