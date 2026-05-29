// Package xrpl provides XRPL payment transaction construction and fee scheduling for TEE instructions.
package xrpl

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/payments"
)

// PaymentTxFromInstruction prepares a transaction from the Payment Instruction Message.
// Scope: native XRP payments only — TokenId must be empty, signaling "use the chain's native asset".
// Hardening:
//
//   - i.Amount must be non-nil (the on-chain decoder allows nil through);
//   - SenderAddress and RecipientAddress must differ (self-payment is only meaningful via
//     Paths/SendMax cross-currency conversion, which this entrypoint does not support);
//   - TokenId must be empty so an issued-token payment cannot accidentally produce a native-XRP tx;
//   - the resulting tx is run through CheckNativePayment (CheckNullify for the nullify
//     path) so the caller never gets back a structurally invalid blob.
func PaymentTxFromInstruction(i payments.ITeePaymentsPaymentInstructionMessage, try int) (map[string]any, error) {
	if i.Amount == nil {
		return nil, errors.New("nil Amount")
	}
	if i.SenderAddress == i.RecipientAddress {
		return nil, errors.New("sender equals recipient")
	}
	if len(i.TokenId) != 0 {
		return nil, fmt.Errorf("non-empty TokenId %s; this entrypoint only builds native XRP payments", hex.EncodeToString(i.TokenId))
	}
	if i.Nonce > math.MaxUint32 {
		return nil, fmt.Errorf("nonce %d exceeds XRPL Sequence (UInt32) range", i.Nonce)
	}

	tx := make(map[string]any)

	tx["TransactionType"] = "Payment"
	tx["Account"] = i.SenderAddress
	tx["Destination"] = i.RecipientAddress
	tx["Amount"] = i.Amount.String()
	tx["SigningPubKey"] = ""
	tx["Sequence"] = uint32(i.Nonce) //nolint:gosec // XRPL Sequence is UInt32; on-chain Nonce bounded to fit

	// [
	//   {
	//     "Memo": {
	//       "MemoData": paymentReference
	//     }
	//   }
	// ]
	memos := []any{
		map[string]any{
			"Memo": map[string]any{
				"MemoData": hex.EncodeToString(i.PaymentReference[:]),
			}},
	}
	tx["Memos"] = memos

	feeFixture, err := ParseFeeEntry(i.FeeSchedule, try)
	if err != nil {
		return nil, err
	}

	fee := feeFixture.Fee(i.MaxFee)

	if feeFixture.Nullify {
		return Nullify(i, fee)
	}

	tx["Fee"] = feeFixture.Fee(i.MaxFee)

	if err := CheckNativePayment(tx); err != nil {
		return nil, fmt.Errorf("native payment validation: %w", err)
	}

	return tx, nil
}

// Nullify prepares a transaction that nullifies the transaction prepared for the Payment Instruction Message.
// Cannot be used as a replacement as the fee is not raised. The result is validated by CheckNullify.
func Nullify(i payments.ITeePaymentsPaymentInstructionMessage, fee string) (map[string]any, error) {
	if i.Nonce > math.MaxUint32 {
		return nil, fmt.Errorf("nonce %d exceeds XRPL Sequence (UInt32) range", i.Nonce)
	}

	tx := make(map[string]any)

	tx["TransactionType"] = "AccountSet"
	tx["Account"] = i.SenderAddress
	tx["SigningPubKey"] = ""
	tx["Sequence"] = uint32(i.Nonce) //nolint:gosec // bounded by the guard above

	// [
	//   {
	//     "Memo": {
	//       "MemoData": paymentReference
	//     }
	//   }
	// ]
	memos := []any{
		map[string]any{
			"Memo": map[string]any{
				"MemoData": hex.EncodeToString(i.PaymentReference[:]),
			}},
	}
	tx["Memos"] = memos

	tx["Fee"] = fee

	if err := CheckNullify(tx); err != nil {
		return nil, fmt.Errorf("nullify validation: %w", err)
	}

	return tx, nil
}

// CheckNativePayment checks the following for a native xrpl payment for multi-signing:
//   - correct transaction type
//   - Account address and Destination address are set and are not equal (validity of the addresses is checked during serializing)
//   - Fee is set and nonzero
//   - Amount is for native token and is nonzero
//   - Sequence is set (formatting of sequence is checked during serialization)
//   - SigningPubKey is set to empty string
func CheckNativePayment(tx map[string]any) error {
	// Transaction type
	if tx["TransactionType"] != "Payment" {
		return fmt.Errorf("invalid transaction type %v", tx["TransactionType"])
	}

	// Addresses
	account, ok := tx["Account"]
	if !ok {
		return errors.New("missing Account")
	}

	destination, ok := tx["Destination"]
	if !ok {
		return errors.New("missing Destination")
	}

	if account == destination {
		return errors.New("account is equal to the destination")
	}

	// Fee
	if err := checkNonzeroFee(tx); err != nil {
		return err
	}

	// Amount
	amount, ok := tx["Amount"]
	if !ok {
		return errors.New("missing Amount")
	}
	amountStr, ok := amount.(string)
	if !ok {
		return fmt.Errorf("invalid Amount type %v: %v", amount, reflect.TypeOf(amount))
	}
	amountBig, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		return fmt.Errorf("unparsable integer for Amount: %v", amountStr)
	}
	if amountBig.Cmp(big.NewInt(0)) != 1 {
		return fmt.Errorf("non positive integer for Amount: %v", amountStr)
	}

	// Sequence
	_, ok = tx["Sequence"]
	if !ok {
		return errors.New("missing Sequence")
	}

	// SigningPubKey
	if err := checkEmptySigningPubKey(tx); err != nil {
		return err
	}

	return nil
}

// CheckNullify validates a nullify (AccountSet) transaction for multi-signing:
//   - correct transaction type
//   - Account is set
//   - Fee is set and nonzero
//   - Sequence is set
//   - SigningPubKey is set to empty string
func CheckNullify(tx map[string]any) error {
	if tx["TransactionType"] != "AccountSet" {
		return fmt.Errorf("invalid transaction type %v", tx["TransactionType"])
	}

	if _, ok := tx["Account"]; !ok {
		return errors.New("missing Account")
	}

	if err := checkNonzeroFee(tx); err != nil {
		return err
	}

	if _, ok := tx["Sequence"]; !ok {
		return errors.New("missing Sequence")
	}

	return checkEmptySigningPubKey(tx)
}

// checkNonzeroFee reports whether tx carries a string Fee that parses to a nonzero UInt32.
func checkNonzeroFee(tx map[string]any) error {
	fee, ok := tx["Fee"]
	if !ok {
		return errors.New("missing Fee")
	}
	feeStr, ok := fee.(string)
	if !ok {
		return fmt.Errorf("invalid Fee type %v: %v", fee, reflect.TypeOf(fee))
	}
	feeUint, err := strconv.ParseUint(feeStr, 10, 32)
	if err != nil {
		return fmt.Errorf("unparsable unsigned integer for Fee: %v: %w", fee, err)
	}
	if feeUint == 0 {
		return errors.New("zero Fee set")
	}
	return nil
}

// checkEmptySigningPubKey reports whether tx carries an empty-string SigningPubKey.
func checkEmptySigningPubKey(tx map[string]any) error {
	signingPubKey, ok := tx["SigningPubKey"]
	if !ok {
		return errors.New("missing 'empty' SigningPubKey")
	}
	signingPubKeyStr, ok := signingPubKey.(string)
	if !ok {
		return fmt.Errorf("invalid type of signingPubKey %v: %v", signingPubKey, reflect.TypeOf(signingPubKey))
	}
	if signingPubKeyStr != "" {
		return errors.New("signingPubKey must be empty string")
	}
	return nil
}

type ScheduledFee struct {
	Nullify bool
	FeeBIPS uint16
	Delay   time.Duration
}

// Fee returns the FeeBIPS of max. A nil max is treated as zero.
func (s *ScheduledFee) Fee(max *big.Int) string {
	if max == nil {
		return "0"
	}
	bips := big.NewInt(int64(s.FeeBIPS))
	fee := new(big.Int).Mul(bips, max)
	fee.Div(fee, big.NewInt(10000))
	return fee.String()
}

func ParseFeeEntries(schedule []byte) ([]ScheduledFee, error) {
	if len(schedule)%4 != 0 {
		return nil, errors.New("invalid schedule length")
	}
	num := len(schedule) / 4

	out := make([]ScheduledFee, 0, num)

	for i := range num {
		entry := schedule[i*4 : (i+1)*4]
		fee, err := parseScheduledFee(entry)
		if err != nil {
			return nil, err
		}
		out = append(out, fee)
	}

	return out, nil
}

func ParseFeeEntry(schedule []byte, try int) (ScheduledFee, error) {
	if len(schedule)%4 != 0 {
		return ScheduledFee{}, errors.New("invalid schedule length")
	}
	// Check try against the entry count directly. The original guard
	// computed (try+1)*4 and compared against len(schedule): for negative
	// try the RHS becomes non-positive and the check passed, then the
	// slice with a negative bound panicked; for try near math.MaxInt the
	// (try+1)*4 multiplication wrapped to a large negative and the slice
	// likewise panicked. Range-checking try directly avoids both edges.
	entries := len(schedule) / 4
	if try < 0 || try >= entries {
		return ScheduledFee{}, fmt.Errorf("try %d out of range [0, %d)", try, entries)
	}

	entry := schedule[try*4 : (try+1)*4]
	fee, err := parseScheduledFee(entry)
	if err != nil {
		return ScheduledFee{}, err
	}

	return fee, nil
}

func parseScheduledFee(fixture []byte) (ScheduledFee, error) {
	if len(fixture) != 4 {
		return ScheduledFee{}, errors.New("invalid fixture length")
	}
	bips := binary.BigEndian.Uint16(fixture[0:2])

	bipsInt16 := int16(bips)
	if bipsInt16 == 0 || bipsInt16 < -10000 || bipsInt16 > 10000 {
		return ScheduledFee{}, errors.New("invalidBIPS")
	}

	delay := time.Duration(binary.BigEndian.Uint16(fixture[2:4])) * time.Second

	nullify := bipsInt16 < 0

	if nullify {
		bipsInt16 = -bipsInt16
	}

	return ScheduledFee{
		Nullify: nullify,
		FeeBIPS: uint16(bipsInt16),
		Delay:   delay,
	}, nil
}
