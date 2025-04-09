package types

import "fmt"

// add MPT issue

type Issue struct{}

func (h *Issue) ToBytes(value any, _ bool) ([]byte, error) {
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

	issuerBytes, err := trimAddress(issuerStr)
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
