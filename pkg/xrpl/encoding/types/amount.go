package types

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"
	"strconv"
	"strings"
)

// TODO amounts without issuer/amount.... MPT support

const e = byte(69) // E in utf8 encoding

const (
	xrpValuePrefix uint64 = 0x4000000000000000
)
const xrpValueByteLen int = 8
const maxXRPAmount uint64 = 1e17 // 100 billion drops aka 10^17

const (
	precision                    = 15
	exponentNormalization int64  = 97
	minNormalizedExponent int64  = 1                // -96 + 97
	maxNormalizedExponent int64  = 177              // 80 + 97
	minSignificant        uint64 = 1000000000000000 // 10^15
	maxSignificant        uint64 = 9999999999999999 // 10^16-1
)

const xrpBitMask uint8 = 0b10000000
const signBitMask uint8 = 0b01000000
const valueBitMask uint8 = 0b00111111

const StandardCodeRegex = `[0-9A-Za-z?!@#$%^&*<>(){}\[\]|]`

// Amount is for serialization of Amount fields. https://xrpl.org/docs/references/protocol/binary-format#amount-fields
//
// Amount implements Encoder interface.
type amount struct{}

var Amount = &amount{}

// ToBytes serializes Amount field.
func (a *amount) ToBytes(value any, _ bool) ([]byte, error) {
	switch value := value.(type) {
	case string:
		return xrpToBytes(value)
	case map[string]any:
		return tokenToBytes(value)
	default:
		return nil, fmt.Errorf("invalid amount: %v", value)
	}
}

// xrpToBytes serializes an XRP amount.
func xrpToBytes(amount string) ([]byte, error) {
	val, err := strconv.ParseUint(amount, 10, 62)
	if err != nil {
		return nil, fmt.Errorf("parsing xrp amount %v", err)
	}
	if val > maxXRPAmount {
		return nil, fmt.Errorf("amount to large %v", val)
	}

	valPrefix := val | xrpValuePrefix
	valBytes := make([]byte, xrpValueByteLen)

	binary.BigEndian.PutUint64(valBytes, valPrefix)

	return valBytes, nil
}

// tokenToBytes serializes an issued currency amount.
func tokenToBytes(amount map[string]any) ([]byte, error) {
	if len(amount) != 3 {
		return nil, fmt.Errorf("wrong number of fields")
	}

	value, err := extractString(amount, "value")
	if err != nil {
		return nil, fmt.Errorf("extracting value: %v", err)
	}
	currency, err := extractString(amount, "currency")
	if err != nil {
		return nil, fmt.Errorf("extracting currency: %v", err)
	}
	issuer, err := extractString(amount, "issuer")
	if err != nil {
		return nil, fmt.Errorf("extracting issuer: %v", err)
	}

	issuerBytes, err := ID(issuer)
	if err != nil {
		return nil, fmt.Errorf("issuer address: %v", err)
	}
	currencyBytes, err := serializeCurrency(currency)
	if err != nil {
		return nil, fmt.Errorf("currency: %v", err)
	}
	normalizedValue, err := serializeTokenValue(value)
	if err != nil {
		return nil, fmt.Errorf("value: %v", err)
	}

	out := make([]byte, 0, 8+20+20)

	out = append(out, normalizedValue...)
	out = append(out, currencyBytes...)
	out = append(out, issuerBytes...)

	return out, nil
}

func extractString(values map[string]any, key string) (string, error) {
	value, ok := values[key]
	if !ok {
		return "", fmt.Errorf("missing '%s'", key)
	}
	valueStr, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("value for key `%s` is not string: %v", key, value)
	}

	return valueStr, nil
}

var disallowedCodes = map[string]bool{
	"XRP": true,
	"0000000000000000000000005852500000000000": true,
	"0000000000000000000000000000000000000000": true,
}

func serializeCurrency(currency string) ([]byte, error) {
	if len(currency) > 5 {
		// allow 0x prefixed nonstandard codes
		currency = strings.TrimPrefix(currency, "0x")
	}

	if disallowedCodes[currency] {
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

func deserializeCurrency(c []byte) (string, error) {
	if len(c) != 20 {
		return "", fmt.Errorf("invalid currency length %v should be 20", len(c))
	}

	if c[0] == 0 {
		isIso := true
		for j := 1; j < 13; j++ {
			if c[j] != 0 {
				isIso = false
			}
		}

		for j := 15; j < 21; j++ {
			if c[j] != 0 {
				isIso = false
			}
		}

		if isIso {
			out := string(c[13:16])
			if out == "XRP" {
				return "", fmt.Errorf("invalid currency %v", out)
			}
			return out, nil
		}
	}

	return hex.EncodeToString(c), nil
}

func serializeStandardCode(code string) ([]byte, error) {
	r := regexp.MustCompile(StandardCodeRegex)

	m := r.FindAllString(code, -1)
	if len(code) != 3 || len(m) != 3 {
		return nil, fmt.Errorf("invalid currency code %s", code)
	}

	codeBytes := make([]byte, 20)
	copy(codeBytes[12:], []byte(code)) // 12 bytes zero | 3 bytes currency | 5 bytes bytes zero
	return codeBytes[:], nil
}

// todo consider disallowing 0x00 prefixed nonstandard codes
func serializeNonstandardCode(code string) ([]byte, error) {
	out, err := hex.DecodeString(code)

	if err != nil {
		return nil, fmt.Errorf("invalid nonstandard currency code: %v", err)
	}

	return out, nil
}

func serializeTokenValue(value string) ([]byte, error) {
	fl, _, err := new(big.Float).Parse(value, 10)
	if err != nil {
		return nil, fmt.Errorf("unparsable value %v", value)
	}

	sign := uint64(1) // positive

	out := make([]byte, 8)

	switch fl.Cmp(big.NewFloat(0)) {
	case 0:
		//return  "0x80000000"
		out[0] = 8 << 4
		return out, nil
	case -1:
		// set sign byte to 0 and change fl to positive
		sign = 0
		fl = fl.Neg(fl)
	}

	significant, exponent, err := format(fl, precision, 56)
	if err != nil {
		return nil, fmt.Errorf("formatting float %v: %v", fl, err)
	}

	significantCheck, _, err := format(fl, precision+1, 64)
	if err != nil {
		return nil, fmt.Errorf("formatting float %v: %v", fl, err)
	}

	if significant*10 != significantCheck {
		return nil, fmt.Errorf("precision overflow %v, at most %v nonzero leading digits", value, precision)
	}

	normalizedExponent := exponent + exponentNormalization - precision
	if normalizedExponent < minNormalizedExponent || normalizedExponent > maxNormalizedExponent {
		return nil, fmt.Errorf("exponent %d out of range [%d,%d]", exponent, int64(minNormalizedExponent)-exponentNormalization, int64(maxNormalizedExponent)-exponentNormalization)
	}

	if significant < minSignificant || significant > maxSignificant {
		return nil, fmt.Errorf("significant part %d out of range [%d,%d]", significant, minSignificant, maxSignificant)
	}

	exponentN := uint64(normalizedExponent)

	outUint := 1<<63 | sign<<62 | exponentN<<54 | significant // not xrp (1bit) | sign (1bit) | exponent (8bits) | significant (5bits)
	binary.BigEndian.PutUint64(out, outUint)

	return out, nil
}

func format(f *big.Float, p int, bitSize int) (uint64, int64, error) {
	formatted := f.Text(e, p) // "d.ddddddddd...dddddE±dd"

	separated := strings.Split(formatted, ".")

	leadingDigit := separated[0]
	rest := separated[1]

	second := strings.Split(rest, "E")

	significantDec := leadingDigit + second[0]
	exponentStr := second[1]

	significant, err := strconv.ParseUint(significantDec, 10, bitSize)
	if err != nil {
		return 0, 0, fmt.Errorf("parsing significant part: %v", err)
	}

	exponent, err := strconv.ParseInt(exponentStr, 10, 8)
	if err != nil {
		return 0, 0, fmt.Errorf("parsing exponent: %v", err)
	}

	return significant, exponent, err
}

const exponentMask = 0xff << 54
const significantMask = (1 << 54) - 1

func deserializeTokenAmount(a []byte) (string, error) {
	if len(a) != 8 {
		return "", fmt.Errorf("invalid token amount %v", a)
	}

	firstByte := a[0]

	number := binary.BigEndian.Uint64(a)

	exponentNormalized := (number & exponentMask) >> 54
	exponent := int64(exponentNormalized) - exponentNormalization

	sign := ""

	val := number & significantMask

	if firstByte&signBitMask == 0 && val != 0 {
		sign = "-"
	}

	floatStr := sign + strconv.FormatUint(val, 10) + "e" + strconv.FormatInt(exponent, 10)

	fl, ok := new(big.Float).SetString(floatStr)
	if !ok {
		return "", fmt.Errorf("not parsable float %v", floatStr)
	}

	return fl.Text(0x67, -1), nil
}

func (a *amount) ToJson(b *bytes.Buffer, _ int) (any, error) {
	firstByte, err := b.ReadByte()
	if err != nil {
		return nil, fmt.Errorf("cannot read first byte: %v", err)
	}

	first := firstByte & xrpBitMask

	switch first {
	case xrpBitMask:
		return xrpToJson(firstByte, b)
	case 0:
		return tokenToJson(firstByte, b)
	default:
		return nil, fmt.Errorf("impossible, first bit neither 1 nor 0: %v", first)
	}
}

func xrpToJson(firstByte byte, b *bytes.Buffer) (string, error) {
	v := make([]byte, 8)
	v[0] = firstByte & valueBitMask // remove sign and "not XRP bit"

	_, err := b.Read(v[1:])
	if err != nil {
		return "", fmt.Errorf("cannot read xrp amount: %v", err)
	}

	var out string

	sign := firstByte & signBitMask
	if sign == 0 {
		out = "-"
	}
	value := binary.BigEndian.Uint64(v)
	out += strconv.FormatUint(value, 10)

	return out, nil
}

func tokenToJson(firstByte byte, b *bytes.Buffer) (map[string]any, error) {
	out := make(map[string]any)

	// amount
	a := make([]byte, 8)
	a[0] = firstByte & valueBitMask // remove sign and "not XRP bit"

	_, err := b.Read(a[1:])
	if err != nil {
		return nil, fmt.Errorf("cannot read token amount: %v", err)
	}

	amount, err := deserializeTokenAmount(a)
	if err != nil {
		return nil, fmt.Errorf("deserializing amount: %v", err)
	}
	out["amount"] = amount

	// currency
	c := make([]byte, 20)
	_, err = b.Read(c)
	if err != nil {
		return nil, fmt.Errorf("cannot read token currency: %v", err)
	}

	cCode, err := deserializeCurrency(c)
	if err != nil {
		return nil, fmt.Errorf("deserializing currency: %v", err)
	}
	out["currency"] = cCode

	// issuer
	i := make([]byte, 20)
	_, err = b.Read(i)
	if err != nil {
		return nil, fmt.Errorf("cannot read token issuer: %v", err)
	}

	iAddress, err := Address(i)
	if err != nil {
		return nil, fmt.Errorf("deserializing issuer: %v", err)
	}

	out["issuer"] = iAddress

	return out, nil
}
