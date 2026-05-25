package transactions

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"

	"github.com/flare-foundation/go-flare-common/pkg/xrpl/encoding/types"
)

// CheckAndEncodePayment validates and serializes a payment transaction.
//
// Scope: plain native or single-currency-IOU payments. Self-payments
// (Account == Destination) are rejected unconditionally because the only
// rippled-meaningful self-payment is a cross-currency conversion via
// Paths/SendMax, which this function does not support. Callers that need
// conversion paths must build and sign their transaction by another route.
//
// Amount validation: native amounts are explicitly rejected when zero.
// Non-native (IOU) amount validation — including zero-value rejection — is
// delegated to types.Encode (STAmount canonical form per rippled); the
// entrypoint does not add an extra zero check on the IOU branch.
//
// The returned bytes are the for-signing canonical form produced by
// types.Encode(tx, true): non-signing fields (TxnSignature, Signers, …) are
// omitted, and the STX\0 prefix is NOT included. The caller is expected to
// hash this with SHA-512-Half (after prepending STX\0) to obtain the digest
// to sign, then assemble the wire-form blob separately by re-encoding the tx
// with the signature attached. Do NOT submit the returned bytes to rippled
// directly — they are not a valid wire transaction.
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

	// Sequence and SigningPubKey are soeREQUIRED in rippled TxFormats.cpp.
	// SigningPubKey is empty in the multi-sig flow, so only check presence.
	if _, ok := decoded["Sequence"]; !ok {
		return nil, errors.New("missing Sequence")
	}
	if _, ok := decoded["SigningPubKey"]; !ok {
		return nil, errors.New("missing SigningPubKey")
	}

	if err := validatePaymentMemos(decoded); err != nil {
		return nil, err
	}

	feeAny, ok := decoded["Fee"] // format is checked during encoding
	if !ok {
		return nil, errors.New("missing Fee")
	}
	feeStr, ok := feeAny.(string)
	if !ok {
		return nil, fmt.Errorf("invalid fee %v", feeAny) // check that fee is xrp amount
	}
	fee, err := strconv.ParseUint(feeStr, 10, 64)
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
		amount, err := strconv.ParseUint(amountStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid amount %s: %w", amountStr, err) // should never happen
		}
		if amount == 0 {
			return nil, errors.New("zero amount")
		}
	}

	return encoded, nil
}

// rippledMemoSerializedLimit mirrors the cap in rippled STTx::isMemoOkay
// (libxrpl/protocol/STTx.cpp): the serialized Memos array (each Memo object
// with its field tags plus the array-end marker) must not exceed 1024 bytes.
const rippledMemoSerializedLimit = 1024

// validatePaymentMemos enforces rippled's isMemoOkay (STTx.cpp) on every memo
// in the decoded payment:
//   - the serialized Memos array fits within rippledMemoSerializedLimit;
//   - MemoType and MemoFormat hex-decode and contain only RFC 3986 URL
//     characters. MemoData is unconstrained beyond being valid hex (already
//     enforced by the codec on the Encode/Decode round-trip).
func validatePaymentMemos(decoded map[string]any) error {
	memosAny, ok := decoded["Memos"]
	if !ok {
		return nil
	}
	memos, ok := memosAny.([]any)
	if !ok {
		return fmt.Errorf("invalid Memos shape: %T", memosAny)
	}

	serialized, err := types.STArray.ToBytes(memos, true)
	if err != nil {
		return fmt.Errorf("re-serializing Memos: %w", err)
	}
	if len(serialized) > rippledMemoSerializedLimit {
		return fmt.Errorf("memos exceed rippled %d-byte serialized limit (got %d)", rippledMemoSerializedLimit, len(serialized))
	}

	for i, entry := range memos {
		wrapper, ok := entry.(map[string]any)
		if !ok {
			return fmt.Errorf("memo %d: not an object", i)
		}
		memoAny, ok := wrapper["Memo"]
		if !ok {
			return fmt.Errorf("memo %d: missing Memo wrapper", i)
		}
		memo, ok := memoAny.(map[string]any)
		if !ok {
			return fmt.Errorf("memo %d: invalid Memo shape: %T", i, memoAny)
		}
		for _, field := range [...]string{"MemoType", "MemoFormat"} {
			v, present := memo[field]
			if !present {
				continue
			}
			s, ok := v.(string)
			if !ok {
				return fmt.Errorf("memo %d: %s not a string", i, field)
			}
			b, err := hex.DecodeString(s)
			if err != nil {
				return fmt.Errorf("memo %d: %s invalid hex: %w", i, field, err)
			}
			if !ValidateMemo(b) {
				return fmt.Errorf("memo %d: %s contains disallowed bytes per rippled isMemoOkay", i, field)
			}
		}
	}
	return nil
}
