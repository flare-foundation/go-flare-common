package types

import (
	"fmt"
	"strings"
)

const (
	accountFlag  uint8 = 0x01
	currencyFlag uint8 = 0x10
	issuerFlag   uint8 = 0x20

	anotherFlag uint8 = 0xff
	endFlag     uint8 = 0x00
)

type PathSet struct{}

func (p *PathSet) ToBytes(value any, _ bool) ([]byte, error) {
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

func stepToBytes(step map[string]any) ([]byte, error) {
	flag := uint8(0)
	fieldCount := 0

	account, ae := step["account"]
	currency, ce := step["currency"]
	issuer, ie := step["issuer"]

	if ae {
		flag = flag | accountFlag
		fieldCount++
	}
	if ce {
		flag = flag | currencyFlag
		fieldCount++
	}
	if ie {
		flag = flag | issuerFlag
		fieldCount++
	}

	if fieldCount == 0 || fieldCount != len(step) {
		return nil, fmt.Errorf("invalid step %v", step)
	}

	out := make([]byte, 0, 41)
	out = append(out, flag)

	//account
	if ae {
		if ce || ie {
			return nil, fmt.Errorf("invalid path %v, account can only be by itself", step)
		}

		accStr, ok := account.(string)
		if !ok {
			return nil, fmt.Errorf("invalid account %v", account)
		}

		accountBytes, err := ID(accStr)
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

		xrpCurrency := curr == "XRP" || strings.TrimPrefix(curr, "0x") == "0000000000000000000000000000000000000000"
		if xrpCurrency && !ie {
			out = append(out, make([]byte, 20)...)
		} else if xrpCurrency && ie {
			return nil, fmt.Errorf("invalid path %v, XRP has no issuer", step)
		} else {
			currencyBytes, err := serializeCurrency(curr)
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

		issuerBytes, err := ID(i)
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
