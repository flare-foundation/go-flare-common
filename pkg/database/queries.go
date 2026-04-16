package database

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/logger"

	"gorm.io/gorm"
)

const maxQueryDuration = 15 * time.Second

// SetErrorLogger sets logger used to log errors on queries.
//
// Default is without logging.
func SetErrorLogger(l Logger) {
	if l != nil {
		errorLogger = l
	}
}

var errorLogger Logger = logger.Nop{}

// FetchLatestBlock returns the latest block in the database.
func FetchLatestBlock(
	ctx context.Context, db *gorm.DB, _ any,
) (Block, error) {
	return RetryWrapper(fetchLatestBlock, "fetching latest block")(ctx, db, nil)
}

func fetchLatestBlock(
	ctx context.Context, db *gorm.DB, _ any,
) (Block, error) {
	var blocks []Block

	err := db.WithContext(ctx).Order("timestamp DESC").Limit(1).Find(&blocks).Error
	if err != nil {
		return Block{}, err
	}

	if len(blocks) == 0 {
		return Block{}, errors.New("no blocks in database")
	}

	return blocks[0], nil
}

// LatestLogsParams holds parameters for fetching the latest logs by address and topic.
type LatestLogsParams struct {
	Address common.Address
	Topic0  common.Hash
	Number  int
}

// FetchLatestLogsByAddressAndTopic0 returns the last <Number> logs with Topic0 emitted by Address.
func FetchLatestLogsByAddressAndTopic0(
	ctx context.Context, db *gorm.DB, params LatestLogsParams,
) ([]Log, error) {
	return RetryWrapper(fetchLatestLogsByAddressAndTopic0, "fetching logs")(ctx, db, params)
}

func fetchLatestLogsByAddressAndTopic0(
	ctx context.Context, db *gorm.DB, params LatestLogsParams,
) ([]Log, error) {
	var logs []Log

	err := db.WithContext(ctx).Where("address = ? AND topic0 = ?",
		hex.EncodeToString(params.Address[:]), // encodes without 0x prefix and without checksum
		hex.EncodeToString(params.Topic0[:]),
	).Order("timestamp DESC").Limit(params.Number).Find(&logs).Error

	return logs, err
}

// LogsFullParams holds parameters for fetching logs with optional address and topic filters.
type LogsFullParams struct {
	Address common.Address
	Topics  [4]common.Hash
	Number  int // -1 for unlimited

}

// FetchLogsFull returns the last <Number> logs with Topics emitted by Address.
//
// If a topic or address has zero value, it is not included in conditions.
// If the Number is -1, it attempts to return all logs matching the conditions.
func FetchLogsFull(
	ctx context.Context, db *gorm.DB, params LogsFullParams,
) ([]Log, error) {
	return RetryWrapper(fetchLogsFull, "fetching logs")(ctx, db, params)
}

func fetchLogsFull(
	ctx context.Context, db *gorm.DB, params LogsFullParams,
) ([]Log, error) {
	var logs []Log

	pMap := make(map[string]any)

	za := common.Address{}

	if params.Address != za {
		pMap["address"] = hex.EncodeToString(params.Address[:])
	}

	zh := common.Hash{}

	for j := range params.Topics {
		if params.Topics[j] != zh {
			key := fmt.Sprintf("topic%d", j)
			pMap[key] = hex.EncodeToString(params.Topics[j][:])
		}
	}

	err := db.WithContext(ctx).Where(pMap).Order("timestamp DESC").Limit(params.Number).Find(&logs).Error

	return logs, err
}

// LogsParams holds parameters for fetching logs within a range by address and topic.
type LogsParams struct {
	Address  common.Address
	Topic0   common.Hash
	From, To int64 // blockNumber or timestamp depending on the function
}

// FetchLogsByAddressAndTopic0Timestamp fetches all logs matching address and topic0 from timestamp range (from, to], order by timestamp.
func FetchLogsByAddressAndTopic0Timestamp(ctx context.Context, db *gorm.DB, params LogsParams) ([]Log, error) {
	return RetryWrapper(fetchLogsByAddressAndTopic0Timestamp, "fetching logs")(ctx, db, params)
}

func fetchLogsByAddressAndTopic0Timestamp(ctx context.Context, db *gorm.DB, params LogsParams) ([]Log, error) {
	var logs []Log
	err := db.WithContext(ctx).Where(
		"address = ? AND topic0 = ? AND timestamp > ? AND timestamp <= ?",
		hex.EncodeToString(params.Address[:]), // encodes without 0x prefix and without checksum
		hex.EncodeToString(params.Topic0[:]),
		params.From,
		params.To,
	).Order("timestamp").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}

// FetchLogsByAddressAndTopic0FromTimestampToBlockNumber fetches all logs matching Address and Topic0 From timestamp (included) To block number (included), order by timestamp.
func FetchLogsByAddressAndTopic0FromTimestampToBlockNumber(ctx context.Context, db *gorm.DB, params LogsParams) ([]Log, error) {
	return RetryWrapper(fetchLogsByAddressAndTopic0FromTimestampToBlockNumber, "fetching logs")(ctx, db, params)
}

func fetchLogsByAddressAndTopic0FromTimestampToBlockNumber(ctx context.Context, db *gorm.DB, params LogsParams) ([]Log, error) {
	var logs []Log

	err := db.WithContext(ctx).Where(
		"address = ? AND topic0 = ? AND timestamp >= ? AND block_number <= ?",
		hex.EncodeToString(params.Address[:]), // encodes without 0x prefix and without checksum
		hex.EncodeToString(params.Topic0[:]),
		params.From,
		params.To,
	).Order("timestamp").Find(&logs).Error
	if err != nil {
		return nil, err
	}

	return logs, nil
}

// FetchLogsByAddressAndTopic0BlockNumber fetches all logs matching Address and Topic0 from block range (From, To], order by timestamp.
func FetchLogsByAddressAndTopic0BlockNumber(ctx context.Context, db *gorm.DB, params LogsParams) ([]Log, error) {
	return RetryWrapper(fetchLogsByAddressAndTopic0BlockNumber, "fetching logs")(ctx, db, params)
}

func fetchLogsByAddressAndTopic0BlockNumber(ctx context.Context, db *gorm.DB, params LogsParams) ([]Log, error) {
	var logs []Log

	err := db.WithContext(ctx).Where(
		"address = ? AND topic0 = ? AND block_number > ? AND block_number <= ?",
		hex.EncodeToString(params.Address[:]), // encodes without 0x prefix and without checksum
		hex.EncodeToString(params.Topic0[:]),
		params.From,
		params.To,
	).Order("timestamp").Find(&logs).Error
	if err != nil {
		return nil, err
	}

	return logs, nil
}

// TxParams holds parameters for fetching transactions within a range by address and function selector.
type TxParams struct {
	ToAddress   common.Address
	FunctionSel [4]byte
	From, To    int64 // blockNumber or timestamp depending on the function
}

// FetchTransactionsByAddressAndSelectorTimestamp fetches all transactions matching ToAddress and FunctionSel from timestamp range (From, To], order by timestamp.
func FetchTransactionsByAddressAndSelectorTimestamp(ctx context.Context, db *gorm.DB, params TxParams) ([]Transaction, error) {
	return RetryWrapper(fetchTransactionsByAddressAndSelectorTimestamp, "fetching transactions")(ctx, db, params)
}

func fetchTransactionsByAddressAndSelectorTimestamp(ctx context.Context, db *gorm.DB, params TxParams) ([]Transaction, error) {
	var transactions []Transaction

	err := db.WithContext(ctx).Where(
		"to_address = ? AND function_sig = ? AND timestamp > ? AND timestamp <= ?",
		hex.EncodeToString(params.ToAddress[:]), // encodes without 0x prefix and without checksum
		hex.EncodeToString(params.FunctionSel[:]),
		params.From,
		params.To,
	).Order("timestamp").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// FetchTransactionsByAddressAndSelectorBlockNumber fetches all transactions matching ToAddress and FunctionSel from block number range (From, To], order by timestamp.
func FetchTransactionsByAddressAndSelectorBlockNumber(ctx context.Context, db *gorm.DB, params TxParams) ([]Transaction, error) {
	return RetryWrapper(fetchTransactionsByAddressAndSelectorBlockNumber, "fetching transactions")(ctx, db, params)
}

func fetchTransactionsByAddressAndSelectorBlockNumber(ctx context.Context, db *gorm.DB, params TxParams) ([]Transaction, error) {
	var transactions []Transaction

	err := db.WithContext(ctx).Where(
		"to_address = ? AND function_sig = ? AND block_number > ? AND block_number <= ?",
		hex.EncodeToString(params.ToAddress[:]), // encodes without 0x prefix and without checksum
		hex.EncodeToString(params.FunctionSel[:]),
		params.From,
		params.To,
	).Order("timestamp").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// FetchTransactionsByAddressAndSelectorFromBlockNumber fetches all transactions matching ToAddress and FunctionSel with blockNumber higher than From, order by timestamp.
func FetchTransactionsByAddressAndSelectorFromBlockNumber(ctx context.Context, db *gorm.DB, params TxParams) ([]Transaction, error) {
	return RetryWrapper(fetchTransactionsByAddressAndSelectorFromBlockNumber, "fetching transactions")(ctx, db, params)
}

func fetchTransactionsByAddressAndSelectorFromBlockNumber(ctx context.Context, db *gorm.DB, params TxParams) ([]Transaction, error) {
	var transactions []Transaction

	err := db.WithContext(ctx).Where(
		"to_address = ? AND function_sig = ? AND block_number > ?",
		hex.EncodeToString(params.ToAddress[:]), // encodes without 0x prefix and without checksum
		hex.EncodeToString(params.FunctionSel[:]),
		params.From,
	).Order("timestamp").Find(&transactions).Error
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

// FetchFirstDBBlockTs retrieves the first indexed block timestamp from the database.
func FetchFirstDBBlockTs(ctx context.Context, db *gorm.DB) (uint64, error) {
	state, err := RetryWrapper(fetchState, "fetching first db block state")(ctx, db, "first_database_block")
	if err != nil {
		return 0, err
	}
	return state.BlockTimestamp, nil
}

// FetchLastDBBlockTs retrieves the last indexed block timestamp from the database.
func FetchLastDBBlockTs(ctx context.Context, db *gorm.DB) (uint64, error) {
	state, err := RetryWrapper(fetchState, "fetching last db block state")(ctx, db, "last_database_block")
	if err != nil {
		return 0, err
	}
	return state.BlockTimestamp, nil
}

// FetchState returns the "last_database_block" state from the indexer database.
func FetchState(ctx context.Context, db *gorm.DB, _ any) (State, error) {
	return RetryWrapper(fetchState, "fetching state")(ctx, db, "last_database_block")
}

func fetchState(ctx context.Context, db *gorm.DB, stateName string) (State, error) {
	var states []State

	err := db.WithContext(ctx).Where(
		"name = ?",
		stateName,
	).Find(&states).Error

	if err != nil {
		return State{}, err
	}

	if len(states) == 0 {
		return State{}, errors.New("no states in database")
	}

	return states[0], nil
}

// RetryWrapper wraps a query function to retry with exponential backoff up to 15 seconds.
func RetryWrapper[F any, P any](query func(context.Context, *gorm.DB, P) (F, error), errorMsg string) func(context.Context, *gorm.DB, P) (F, error) {
	wrappedFunc := func(ctx context.Context, db *gorm.DB, params P) (F, error) {
		var returnValue F

		err := backoff.RetryNotify(
			func() error {
				var err error
				returnValue, err = query(ctx, db, params)
				return err
			},
			backoff.WithContext(backoff.NewExponentialBackOff(backoff.WithMaxElapsedTime(maxQueryDuration)), ctx),
			func(err error, duration time.Duration) {
				errorLogger.Errorf("error %s: %v, retrying after %v", errorMsg, err, duration)
			},
		)

		return returnValue, err
	}

	return wrappedFunc
}
