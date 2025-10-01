package types

import (
	"bytes"
	"fmt"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/address"
)

// add MPT issue

var Issue = &issue{}

type issue struct{}

func (*issue) ToBytes(value any, _ bool) ([]byte, error) {
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
	if currencyStr == XRP {
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

	issuerBytes, err := address.ID(issuerStr)
	if err != nil {
		return nil, fmt.Errorf("issuer address: %v", err)
	}

	code, err := serializeCurrency(currencyStr)
	if err != nil {
		return nil, fmt.Errorf("invalid issuer %v: %v", issuerStr, err)
	}

	code = append(code, issuerBytes...)

	return code, nil
}

func (*issue) ToJSON(b *bytes.Buffer, _ int) (any, error) {
	out := make(map[string]any)

	l := 20
	cCode := make([]byte, l)
	n, err := b.Read(cCode)
	if err != nil {
		return nil, fmt.Errorf("reading currency code: %v", err)
	}
	if n != l {
		return nil, outOfBytes("currency code", l, n)
	}

	if bytes.Equal(cCode, make([]byte, 20)) {
		out["currency"] = XRP

		return out, nil
	}

	c, err := deserializeCurrency(cCode)
	if err != nil {
		return nil, fmt.Errorf("deserializing currency code: %v", err)
	}

	out["currency"] = c

	l = 20
	issuer := make([]byte, l)
	n, err = b.Read(issuer)
	if err != nil {
		return nil, fmt.Errorf("reading issuer: %v", err)
	}
	if n != l {
		return nil, outOfBytes("issuer", l, n)
	}

	a, err := address.Address(issuer)
	if err != nil {
		return nil, fmt.Errorf("issuer to address: %v", err)
	}

	out["issuer"] = a

	return out, nil
}
