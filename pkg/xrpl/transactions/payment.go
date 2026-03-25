package transactions

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
)

// CheckAndEncodePayment validates and serializes a payment transaction.
func CheckAndEncodePayment(tx map[string]any, native bool) ([]byte, error) {
	encoded, err := types.Encode(tx, true)
	if err != nil {
		return nil, fmt.Errorf("encoding tx: %w", err)
	}

	// to get a canonical values
	decoded, err := types.Decode(encoded)
	if err != nil {
		return nil, fmt.Errorf("decoding tx: %w", err)
	}

	// TxType
	txType, ok := decoded["TransactionType"]
	if !ok {
		return nil, errors.New("missing TransactionType")
	}
	if txType != "Payment" {
		return nil, fmt.Errorf("transaction type should be 'Payment' is %s", txType)
	}

	account, ok := decoded["Account"] // format is checked during encoding
	if !ok {
		return nil, errors.New("missing Account")
	}

	destination, ok := decoded["Destination"] // format is checked during encoding
	if !ok {
		return nil, errors.New("missing Destination")
	}

	if account == destination {
		return nil, errors.New("destination should not be equal to the account")
	}

	feeAny, ok := decoded["Fee"] // format is checked during encoding
	if !ok {
		return nil, errors.New("missing Fee")
	}
	feeStr, ok := feeAny.(string)
	if !ok {
		return nil, fmt.Errorf("invalid fee %v", feeAny) // check that fee is xrp amount
	}
	fee, err := strconv.ParseUint(feeStr, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("invalid fee %s: %w", feeStr, err) // should never happen
	}
	if fee == 0 {
		return nil, errors.New("zero fee")
	}

	amountAny, ok := decoded["Amount"] // format is checked during encoding
	if !ok {
		return nil, errors.New("missing Amount")
	}
	if native {
		amountStr, ok := amountAny.(string)
		if !ok {
			return nil, fmt.Errorf("invalid native amount %v", amountAny) // check that fee is xrp amount
		}
		amount, err := strconv.ParseUint(amountStr, 10, 0)
		if err != nil {
			return nil, fmt.Errorf("invalid amount %s: %w", feeStr, err) // should never happen
		}
		if amount == 0 {
			return nil, errors.New("zero amount")
		}
	}

	return encoded, nil
}
