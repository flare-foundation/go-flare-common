package xrpl

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"

	"github.com/flare-foundation/go-flare-common/pkg/tee/structs/payment"
)

func PaymentTransactionMultisig(i payment.ITeePaymentsPaymentInstructionMessage) map[string]any {
	tx := make(map[string]any)

	tx["TransactionType"] = "Payment"
	tx["Account"] = i.SenderAddress
	tx["Destination"] = i.RecipientAddress
	tx["Amount"] = i.Amount.String()
	tx["SigningPubKey"] = ""
	tx["Sequence"] = i.Nonce // it is uint32 on blockchain

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
	tx["Fee"] = i.Fee.String()

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
		return fmt.Errorf("unparsable unsigned integer for Fee: %v: %v", fee, err)
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
		return fmt.Errorf("invalid fee type %v: %v", fee, reflect.TypeOf(fee))
	}
	amountBig, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		return fmt.Errorf("unparsable integer for Amount: %v: %v", fee, err)
	}
	if amountBig.Cmp(big.NewInt(0)) != 1 {
		return fmt.Errorf("non positive integer for Amount: %v: %v", fee, err)
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

	SigningPubKeyStr, ok := signingPubKey.(string)
	if !ok {
		return fmt.Errorf("invalid type of signingPubKey %v: %v", signingPubKey, reflect.TypeOf(signingPubKey))
	}

	if SigningPubKeyStr != "" {
		return errors.New("signingPubKey must be empty string")
	}

	return nil
}
