// Package xrpl provides XRPL payment transaction construction and fee scheduling for TEE instructions.
package xrpl

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/payments"
)

// PaymentTxFromInstruction prepares a transaction from the Payment Instruction Message.
func PaymentTxFromInstruction(i payments.ITeePaymentsPaymentInstructionMessage, try int) (map[string]any, error) {
	tx := make(map[string]any)

	tx["TransactionType"] = "Payment"
	tx["Account"] = i.SenderAddress
	tx["Destination"] = i.RecipientAddress
	tx["Amount"] = i.Amount.String()
	tx["SigningPubKey"] = ""
	tx["Sequence"] = i.Nonce

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
		return Nullify(i, fee), nil
	}

	tx["Fee"] = feeFixture.Fee(i.MaxFee)

	return tx, nil
}

// Nullify prepares a transaction that nullifies the transaction prepared for the the Payment Instruction Message.
// Cannot be used as a replacement as the fee is not raised.
func Nullify(i payments.ITeePaymentsPaymentInstructionMessage, fee string) map[string]any {
	tx := make(map[string]any)

	tx["TransactionType"] = "AccountSet"
	tx["Account"] = i.SenderAddress
	tx["SigningPubKey"] = ""
	tx["Sequence"] = i.Nonce

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

	return tx
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

// Fee returns the FeeBIPS of max.
func (s *ScheduledFee) Fee(max *big.Int) string {
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
	if len(schedule) < (try+1)*4 {
		return ScheduledFee{}, errors.New("try beyond schedule length")
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
