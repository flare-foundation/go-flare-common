package types

import (
	"bytes"
	"fmt"
)

// add MPT issue

var Issue = &issue{}

type issue struct{}

func (i *issue) ToBytes(value any, _ bool) ([]byte, error) {
	valueMap, ok := value.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("invalid issue %v", value)
	}

	currency, ok := valueMap["currency"]
	if !ok {
		return nil, fmt.Errorf("invalid issue %v, no currency", value)
	}
	currencyStr, ok := currency.(string)
	if !ok {
		return nil, fmt.Errorf("invalid issue %v, invalid currency", value)
	}

	// XRP issue
	if currencyStr == "XRP" {
		if len(valueMap) != 1 {
			return nil, fmt.Errorf("invalid issue %v", value)
		}
		return make([]byte, 20), nil
	}

	// Token issue
	issuer, ok := valueMap["issuer"]
	if !ok {
		return nil, fmt.Errorf("invalid issue %v, no currency", value)
	}

	issuerStr, ok := issuer.(string)
	if !ok {
		return nil, fmt.Errorf("invalid issuer %v", issuer)
	}

	issuerBytes, err := ID(issuerStr)
	if err != nil {
		return nil, fmt.Errorf("issuer address: %v", err)
	}

	code, err := serializeCurrency(currencyStr)
	if err != nil {
		return nil, fmt.Errorf("invalid issuer %v: %v", issuerStr, err)
	}

	out := append(code, issuerBytes...)

	return out, nil
}

func (i *issue) ToJson(b *bytes.Buffer, _ int) (any, error) {
	out := make(map[string]any)

	cCode := make([]byte, 20)
	_, err := b.Read(cCode)
	if err != nil {
		return nil, fmt.Errorf("reading currency code: %v", err)
	}

	if bytes.Equal(cCode, make([]byte, 20)) {
		out["currency"] = "XRP"

		return out, nil
	}

	c, err := deserializeCurrency(cCode)
	if err != nil {
		return nil, fmt.Errorf("deserializing currency code: %v", err)
	}

	out["currency"] = c

	issuer := make([]byte, 20)
	_, err = b.Read(issuer)
	if err != nil {
		return nil, fmt.Errorf("reading issuer: %v", err)
	}

	a, err := Address(issuer)
	if err != nil {
		return nil, fmt.Errorf("issuer to address: %v", err)
	}

	out["issuer"] = a

	return out, nil
}
