package types

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
)

const (
	accountFlag  uint8 = 0x01
	currencyFlag uint8 = 0x10
	issuerFlag   uint8 = 0x20

	anotherFlag uint8 = 0xff
	endFlag     uint8 = 0x00
)

type PathSet struct{}

func (*PathSet) ToBytes(value any, _ bool) ([]byte, error) {
	paths, ok := value.([]any)
	if !ok {
		return nil, fmt.Errorf("invalid pathSet %v", value)
	}
	if len(paths) < 1 || len(paths) > 6 {
		return nil, fmt.Errorf("invalid pathSet length %v", value)
	}

	out := make([]byte, 0, 83*len(paths))

	for j, path := range paths {
		pathSlice, ok := path.([]any)
		if !ok {
			return nil, fmt.Errorf("invalid pathSet %v, path %v", value, path)
		}

		pathBytes, err := pathToBytes(pathSlice)
		if err != nil {
			return nil, fmt.Errorf("invalid path set: %v", err)
		}

		out = append(out, pathBytes...)

		if j+1 == len(paths) {
			out = append(out, endFlag)
		} else {
			out = append(out, anotherFlag)
		}
	}

	return out, nil
}

func (*PathSet) ToJSON(b *bytes.Buffer, _ int) (any, error) {
	out := make([]any, 0)
	l := 0

	for {
		flag, path, err := readPath(b)
		if err != nil {
			return nil, fmt.Errorf("reading next path: %v", err)
		}
		switch flag {
		case endFlag:
			out = append(out, path)
			l++
			if l > 6 {
				return nil, errors.New("pathSet too large")
			}
			return out, nil
		case anotherFlag:
			out = append(out, path)
			if l > 6 {
				return nil, errors.New("pathSet too large")
			}
			l++
		default:
			return nil, fmt.Errorf("unknown end flag %v", flag)
		}
	}
}

func stepToBytes(step map[string]any) ([]byte, error) {
	flag := uint8(0)
	fieldCount := 0

	account, ae := step["account"]
	currency, ce := step["currency"]
	issuer, ie := step["issuer"]

	if ae {
		flag |= accountFlag
		fieldCount++
	}
	if ce {
		flag |= currencyFlag
		fieldCount++
	}
	if ie {
		flag |= issuerFlag
		fieldCount++
	}

	if fieldCount == 0 || fieldCount != len(step) {
		return nil, fmt.Errorf("invalid step %v", step)
	}

	out := make([]byte, 0, 41)
	out = append(out, flag)

	// account
	if ae {
		if ce || ie {
			return nil, fmt.Errorf("invalid path %v, account can only be by itself", step)
		}

		accStr, ok := account.(string)
		if !ok {
			return nil, fmt.Errorf("invalid account %v", account)
		}

		accountBytes, err := address.ID(accStr)
		if err != nil {
			return nil, fmt.Errorf("invalid account %v", step)
		}

		out = append(out, accountBytes...)

		return out, nil
	}

	// currency
	if ce {
		curr, ok := currency.(string)
		if !ok {
			return nil, fmt.Errorf("invalid currency %v", currency)
		}

		xrpCurrency := curr == XRP || strings.TrimPrefix(curr, "0x") == "0000000000000000000000000000000000000000"

		switch xrpCurrency {
		case true:
			if ie {
				return nil, fmt.Errorf("invalid path %v, XRP has no issuer", step)
			}
			out = append(out, make([]byte, 20)...)

		case false:
			currencyBytes, err := serializeCurrency(curr, disallowedCodes)
			if err != nil {
				return nil, fmt.Errorf("invalid path %v, invalid issuer %v", step, err)
			}

			out = append(out, currencyBytes...)
		}
	}

	// issuer
	if ie {
		i, ok := issuer.(string)
		if !ok {
			return nil, fmt.Errorf("invalid issuer %v", issuer)
		}

		issuerBytes, err := address.ID(i)
		if err != nil {
			return nil, fmt.Errorf("invalid issuer %v", step)
		}

		out = append(out, issuerBytes...)
	}

	return out, nil
}

func pathToBytes(path []any) ([]byte, error) {
	if len(path) < 1 || len(path) > 8 {
		return nil, fmt.Errorf("invalid path length %v", path)
	}
	out := make([]byte, 0, len(path)*41)

	for _, step := range path {
		stepMap, ok := step.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("invalid step: %v", step)
		}

		stepBytes, err := stepToBytes(stepMap)
		if err != nil {
			return nil, fmt.Errorf("invalid path: %v", err)
		}
		out = append(out, stepBytes...)
	}

	return out, nil
}

func readPath(b *bytes.Buffer) (byte, any, error) {
	out := make([]any, 0)
	l := 0

	for {
		flag, step, err := readStep(b)
		if err != nil {
			return flag, nil, fmt.Errorf("reading next step: %v", err)
		}

		if flag == anotherFlag || flag == endFlag {
			if l < 1 {
				return flag, nil, errors.New("empty path")
			}
			return flag, out, nil
		}

		if l >= 6 {
			return flag, nil, errors.New("path too long")
		}

		out = append(out, step)
		l++
	}
}

func readStep(b *bytes.Buffer) (byte, any, error) {
	flag, err := b.ReadByte()
	if err != nil {
		return 0, nil, fmt.Errorf("reading flag: %v", err)
	}

	if flag == anotherFlag || flag == endFlag {
		return flag, nil, nil
	}

	out := make(map[string]any)

	switch flag {
	case 0x01:
		account, err := AccountID.ToJSON(b, 0)
		if err != nil {
			return flag, nil, fmt.Errorf("reading account: %v", err)
		}

		out["account"] = account
		return flag, out, nil

	case 0x10:
		currency, err := readCurrency(b)
		if err != nil {
			return flag, nil, fmt.Errorf("currency: %v", err)
		}

		out["currency"] = currency
		return flag, out, nil

	case 0x20:
		issuer, err := AccountID.ToJSON(b, 0)
		if err != nil {
			return flag, nil, fmt.Errorf("issuer: %v", err)
		}

		out["issuer"] = issuer
		return flag, out, nil

	case 0x30:
		currency, err := readCurrency(b)
		if err != nil {
			return flag, nil, fmt.Errorf("currency before issuer: %v", err)
		}

		out["currency"] = currency

		issuer, err := AccountID.ToJSON(b, 0)
		if err != nil {
			return flag, nil, fmt.Errorf("issuer after currency: %v", err)
		}

		out["issuer"] = issuer

		return flag, out, nil

	default:
		return flag, nil, fmt.Errorf("illegal path flag %x", flag)
	}
}

func readCurrency(b *bytes.Buffer) (string, error) {
	const l = 20
	c := make([]byte, l)
	n, err := b.Read(c)
	if err != nil {
		return "", fmt.Errorf("reading bytes: %v", err)
	}
	if n != l {
		return "nil", outOfBytes("currency", l, n)
	}

	if bytes.Equal(c, make([]byte, l)) {
		return XRP, nil
	}
	currency, err := deserializeCurrency(c)
	if err != nil {
		return "", fmt.Errorf("deserializing currency code: %v", err)
	}

	return currency, nil
}
