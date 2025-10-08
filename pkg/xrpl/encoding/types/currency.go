package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

var Currency = &currency{}

type currency struct{}

func (c *currency) ToBytes(value any, _ bool) ([]byte, error) {
	code, ok := value.(string)
	if !ok {
		return nil, fmt.Errorf("value %v is not string", value)
	}

	return serializeCurrency(code, map[string]bool{})
}

func (c *currency) ToJSON(b *bytes.Buffer, _ int) (any, error) {
	l := 20
	code := make([]byte, l)
	n, err := b.Read(code)
	if err != nil {
		return nil, fmt.Errorf("cannot read currency value: %v", err)
	}
	if n != l {
		return nil, outOfBytes("currency value", l, n)
	}

	return deserializeCurrency(code)
}

func serializeCurrency(currency string, disallowed map[string]bool) ([]byte, error) {
	if len(currency) > 5 {
		// allow 0x prefixed nonstandard codes
		currency = strings.TrimPrefix(currency, "0x")
	}

	if disallowed[currency] {
		return nil, fmt.Errorf("currency %s is not allowed", currency)
	}

	switch len(currency) {
	case 3:
		return serializeStandardCode(currency)
	case 40:
		return serializeNonstandardCode(currency)
	}

	return nil, fmt.Errorf("invalid length. Allowed 3 or 40 got %d", len(currency))
}

func serializeStandardCode(code string) ([]byte, error) {
	r := regexp.MustCompile(StandardCodeRegex)

	m := r.FindAllString(code, -1)
	if len(code) != 3 || len(m) != 3 {
		return nil, fmt.Errorf("invalid currency code %s", code)
	}

	codeBytes := make([]byte, 20)
	copy(codeBytes[12:], []byte(code)) // 12 bytes zero | 3 bytes currency | 5 bytes bytes zero
	return codeBytes, nil
}

func serializeNonstandardCode(code string) ([]byte, error) {
	out, err := hex.DecodeString(code)
	if err != nil {
		return nil, fmt.Errorf("invalid nonstandard currency code: %v", err)
	}

	return out, nil
}

func deserializeCurrency(c []byte) (string, error) {
	if len(c) != 20 {
		return "", fmt.Errorf("invalid currency length %v should be 20", len(c))
	}

	if c[0] == 0 {
		isIso := true
		for j := 1; j < 12; j++ {
			if c[j] != 0 {
				isIso = false
			}
		}

		for j := 15; j < 20; j++ {
			if c[j] != 0 {
				isIso = false
			}
		}

		if isIso {
			out := string(c[12:15])
			if out == XRP {
				return "", fmt.Errorf("invalid currency %v", out)
			}
			return out, nil
		}
	}

	return strings.ToUpper(hex.EncodeToString(c)), nil
}
