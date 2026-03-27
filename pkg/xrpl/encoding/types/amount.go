package types

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
)

type amountType uint8

const (
	invalid amountType = iota
	xrpAmount
	tokenAmount
	mptAmount
)

const e = byte(69) // E in utf8 encoding

const XRP = "XRP"

const maxMPTValue = uint64(0x7FFFFFFFFFFFFFFF)

var maxMPTValueBig = new(big.Int).SetUint64(maxMPTValue)

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

const (
	typeBitMask uint8 = 0b10100000
	xrpType     uint8 = 0b00000000
	tokenType0  uint8 = 0b10000000
	tokenType1  uint8 = 0b10100000
	mptType     uint8 = 0b00100000

	signBitMask        uint8 = 0b01000000
	signedValueBitMask uint8 = 0b01111111
	valueBitMask       uint8 = 0b00111111
)

const StandardCodeRegex = `[0-9A-Za-z?!@#$%^&*<>(){}\[\]|]`

// Amount is for serialization of Amount fields. https://xrpl.org/docs/references/protocol/binary-format#amount-fields
//
// Amount implements Encoder interface.
type amount struct{}

var Amount = &amount{}

// ToBytes serializes Amount field.
func (*amount) ToBytes(value any, _ bool) ([]byte, error) {
	switch value := value.(type) {
	case string:
		return xrpToBytes(value)
	case map[string]any:
		switch decideAmountType(value) {
		case tokenAmount:
			return tokenToBytes(value)
		case mptAmount:
			return mptToBytes(value)
		default:
			return nil, fmt.Errorf("invalid amount struct: %v", value)
		}
	default:
		return nil, fmt.Errorf("invalid amount: %v", value)
	}
}

// ToJSON decodes amount field.
func (a *amount) ToJSON(b *bytes.Buffer, _ int) (any, error) {
	firstByte, err := b.ReadByte()
	if err != nil {
		return nil, fmt.Errorf("cannot read first byte: %w", err)
	}

	amountType := firstByte & typeBitMask

	switch amountType {
	case xrpType:
		return xrpToJSON(firstByte, b)
	case tokenType0, tokenType1:
		return tokenToJSON(firstByte, b)
	case mptType:
		return mptToJSON(firstByte, b)
	default:
		return nil, fmt.Errorf("impossible, unknown amount type: %v", amountType) // unreachable
	}
}

// xrpToBytes serializes an XRP amount.
func xrpToBytes(amount string) ([]byte, error) {
	val, err := strconv.ParseUint(amount, 10, 62)
	if err != nil {
		return nil, fmt.Errorf("xrp amount %w", err)
	}
	if val > maxXRPAmount {
		return nil, fmt.Errorf("amount too large %v", val)
	}

	valPrefix := val | xrpValuePrefix
	valBytes := make([]byte, xrpValueByteLen)

	binary.BigEndian.PutUint64(valBytes, valPrefix)

	return valBytes, nil
}

// tokenToBytes serializes an issued currency amount.
func tokenToBytes(amount map[string]any) ([]byte, error) {
	if len(amount) != 3 {
		return nil, errors.New("wrong number of fields")
	}

	value, err := extractString(amount, "value")
	if err != nil {
		return nil, fmt.Errorf("extracting value: %w", err)
	}
	currency, err := extractString(amount, "currency")
	if err != nil {
		return nil, fmt.Errorf("extracting currency: %w", err)
	}
	issuer, err := extractString(amount, "issuer")
	if err != nil {
		return nil, fmt.Errorf("extracting issuer: %w", err)
	}

	issuerBytes, err := address.ID(issuer)
	if err != nil {
		return nil, fmt.Errorf("issuer address: %w", err)
	}
	currencyBytes, err := serializeCurrency(currency, disallowedCodes)
	if err != nil {
		return nil, fmt.Errorf("currency: %w", err)
	}
	normalizedValue, err := serializeTokenValue(value)
	if err != nil {
		return nil, fmt.Errorf("value: %w", err)
	}

	out := make([]byte, 0, 8+20+20)

	out = append(out, normalizedValue...)
	out = append(out, currencyBytes...)
	out = append(out, issuerBytes...)

	return out, nil
}

func mptToBytes(amount map[string]any) ([]byte, error) {
	if len(amount) != 2 {
		return nil, errors.New("invalid amount struct")
	}

	value, err := extractString(amount, "value")
	if err != nil {
		return nil, fmt.Errorf("extracting value: %w", err)
	}

	absValue, isNegative := mptSign(value)

	valueBytes, err := serializeMPTValue(absValue)
	if err != nil {
		return nil, fmt.Errorf("serializing value: %w", err)
	}

	indicator := byte(0x60) // cMPToken | cPositive
	if isNegative && valueBytes != [8]byte{} {
		indicator = 0x20 // cMPToken only (no cPositive bit)
	}

	issuanceID, err := extractString(amount, "mpt_issuance_id")
	if err != nil {
		return nil, fmt.Errorf("extracting mpt_issuance_id: %w", err)
	}

	isID, err := hex.DecodeString(issuanceID)
	if err != nil {
		return nil, fmt.Errorf("mpt_issuance_id: %w", err)
	}

	if len(isID) != 24 {
		return nil, errors.New("mpt_issuance_id invalid length")
	}

	out := make([]byte, 0, 33)
	out = append(out, indicator)
	out = append(out, valueBytes[:]...)
	out = append(out, isID...)

	return out, nil
}

// mptSign returns the absolute value string and whether the original was negative.
// Hex values (0x...) are never treated as negative.
func mptSign(value string) (string, bool) {
	if strings.HasPrefix(value, "-") && !strings.HasPrefix(strings.ToLower(value), "-0x") {
		return value[1:], true
	}
	return value, false
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
	XRP: true,
	"0000000000000000000000005852500000000000": true,
	"0000000000000000000000000000000000000000": true,
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
		// return  "0x8000000000000000"
		out[0] = 0x80
		return out, nil
	case -1:
		// set sign byte to 0 and change fl to positive
		sign = 0
		fl = fl.Neg(fl)
	}

	significant, exponent, err := format(fl, precision, 56)
	if err != nil {
		return nil, fmt.Errorf("formatting float %v: %w", fl, err)
	}

	significantCheck, _, err := format(fl, precision+1, 64)
	if err != nil {
		return nil, fmt.Errorf("formatting float %v: %w", fl, err)
	}

	if significant*10 != significantCheck {
		return nil, fmt.Errorf("precision overflow %v, at most %v nonzero leading digits", value, precision)
	}

	normalizedExponent := exponent + exponentNormalization - precision
	if normalizedExponent < minNormalizedExponent || normalizedExponent > maxNormalizedExponent {
		return nil, fmt.Errorf("exponent %d out of range [%d,%d]", exponent, minNormalizedExponent-exponentNormalization, maxNormalizedExponent-exponentNormalization)
	}

	if significant < minSignificant || significant > maxSignificant {
		return nil, fmt.Errorf("significant part %d out of range [%d,%d]", significant, minSignificant, maxSignificant)
	}

	exponentN := uint64(normalizedExponent)

	outUint := 1<<63 | sign<<62 | exponentN<<54 | significant // not xrp (1bit) | sign (1bit) | exponent (8bits) | significant (54bits)
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
		return 0, 0, fmt.Errorf("parsing significant part: %w", err)
	}

	exponent, err := strconv.ParseInt(exponentStr, 10, 8)
	if err != nil {
		return 0, 0, fmt.Errorf("parsing exponent: %w", err)
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

	val := number & significantMask

	sign := ""
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

func xrpToJSON(firstByte byte, b *bytes.Buffer) (string, error) {
	const l = 8
	v := make([]byte, l)
	v[0] = firstByte & valueBitMask // remove "not XRP bit" and sign

	n, err := b.Read(v[1:])
	if err != nil {
		return "", fmt.Errorf("cannot read xrp amount: %w", err)
	}
	if n != l-1 {
		return "", outOfBytes("xrp amount", l, n+1)
	}

	sign := firstByte & signBitMask
	if sign == 0 {
		return "", errors.New("negative xrp amount")
	}

	value := binary.BigEndian.Uint64(v)
	if value > maxXRPAmount {
		return "", errors.New("xrp amount too large")
	}

	return strconv.FormatUint(value, 10), nil
}

func tokenToJSON(firstByte byte, b *bytes.Buffer) (map[string]any, error) {
	out := make(map[string]any)

	// amount

	l := 8
	a := make([]byte, l)
	a[0] = firstByte & signedValueBitMask // remove "not XRP bit"

	n, err := b.Read(a[1:])
	if err != nil {
		return nil, fmt.Errorf("cannot read token amount: %w", err)
	}
	if n != l-1 {
		return nil, outOfBytes("token amount", l, n+1)
	}

	amount, err := deserializeTokenAmount(a)
	if err != nil {
		return nil, fmt.Errorf("deserializing amount: %w", err)
	}
	out["value"] = amount

	// currency
	l = 20
	c := make([]byte, l)
	n, err = b.Read(c)
	if err != nil {
		return nil, fmt.Errorf("cannot read token currency: %w", err)
	}
	if n != l {
		return nil, outOfBytes("token currency", l, n)
	}

	cCode, err := deserializeCurrency(c)
	if err != nil {
		return nil, fmt.Errorf("deserializing currency: %w", err)
	}
	out["currency"] = cCode

	// issuer
	l = 20
	i := make([]byte, l)
	n, err = b.Read(i)
	if err != nil {
		return nil, fmt.Errorf("cannot read token issuer: %w", err)
	}
	if n != l {
		return nil, outOfBytes("token issuer", l, n)
	}

	iAddress, err := address.Address(i)
	if err != nil {
		return nil, fmt.Errorf("deserializing issuer: %w", err)
	}

	out["issuer"] = iAddress

	return out, nil
}

func mptToJSON(firstByte byte, b *bytes.Buffer) (map[string]any, error) {
	out := make(map[string]any, 2)

	l := 8
	v := make([]byte, l)
	n, err := b.Read(v)
	if err != nil {
		return nil, fmt.Errorf("cannot read mpt value: %w", err)
	}
	if n != l {
		return nil, outOfBytes("mpt value", l, n)
	}

	bv := new(big.Int).SetBytes(v)

	if bv.Cmp(maxMPTValueBig) > 0 {
		return nil, errors.New("mpt value too large")
	}

	valueStr := bv.String()
	if firstByte&signBitMask == 0 && bv.Sign() != 0 {
		valueStr = "-" + valueStr
	}
	out["value"] = valueStr

	l = 24
	mii := make([]byte, l)
	n, err = b.Read(mii)
	if err != nil {
		return nil, fmt.Errorf("cannot read mpt issuance id: %w", err)
	}
	if n != l {
		return nil, outOfBytes("mpt issuance id", l, n)
	}

	out["mpt_issuance_id"] = strings.ToUpper(hex.EncodeToString(mii))

	return out, nil
}

func serializeMPTValue(value string) ([8]byte, error) {
	out := [8]byte{}

	zeroB := big.NewInt(0)

	value = strings.ToLower(value)

	valueBig := new(big.Int)
	var ok bool

	if strings.HasPrefix(value, "0x") {
		if strings.HasPrefix(value, "0x-") {
			return out, errors.New("invalid hex syntax")
		}

		valueBig, ok = valueBig.SetString(value[2:], 16)
		if !ok {
			return out, errors.New("invalid hexadecimal value format")
		}
	} else {
		values := strings.Split(value, "e")

		if len(values) == 0 || len(values) > 2 {
			return out, errors.New("invalid value format")
		}

		valueBig, ok = new(big.Int).SetString(values[0], 10)
		if !ok {
			return out, errors.New("invalid value format main part")
		}

		if len(values) == 2 {
			exponent, ok := new(big.Int).SetString(values[1], 10)
			if !ok {
				return out, errors.New("invalid value format exponent")
			}

			if exponent.Cmp(zeroB) == -1 {
				return out, errors.New("negative exponent not allowed")
			}

			multiplier := new(big.Int).Exp(big.NewInt(10), exponent, nil)

			valueBig.Mul(valueBig, multiplier)
		}
	}

	if valueBig.Cmp(maxMPTValueBig) == 1 || valueBig.Cmp(zeroB) == -1 {
		return out, errors.New("invalid value")
	}

	binary.BigEndian.PutUint64(out[:], valueBig.Uint64())

	return out, nil
}

func decideAmountType(amount map[string]any) amountType {
	if _, hasCurrency := amount["currency"]; hasCurrency {
		return tokenAmount
	}

	if _, hasMPTID := amount["mpt_issuance_id"]; hasMPTID {
		return mptAmount
	}

	return 0 // invalid
}
