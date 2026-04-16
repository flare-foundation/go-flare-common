package database

import (
	"context"
	"encoding/hex"
	"errors"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/logger"

	"gorm.io/gorm"
)

// ErrNotFound is returned when a query finds no matching rows.
var ErrNotFound = errors.New("not found")

// StateName identifies a named state row in the indexer State table.
type StateName string

const (
	StateFirstDBBlock StateName = "first_database_block"
	StateLastDBBlock  StateName = "last_database_block"
	StateLastChain    StateName = "last_chain_block"
)

// StateQuery holds parameters for [DB.FetchState].
type StateQuery struct {
	Name StateName
}

// LogsQuery holds optional filter parameters for [DB.FetchLogs].
// A nil field is not included in the query.
// A pointer to the zero value is a valid defined value and is included.
type LogsQuery struct {
	From    *uint64 // timestamp > From
	To      *uint64 // timestamp <= To
	FromB   *uint64 // block_number > FromB
	ToB     *uint64 // block_number <= ToB
	Address *common.Address
	Topic0  *common.Hash
	Topic1  *common.Hash
	Topic2  *common.Hash
	Topic3  *common.Hash
}

// TxQuery holds optional filter parameters for [DB.FetchTransactions].
// A nil field is not included in the query.
// A pointer to the zero value is a valid defined value and is included.
type TxQuery struct {
	From     *uint64 // timestamp > From
	To       *uint64 // timestamp <= To
	FromB    *uint64 // block_number > FromB
	ToB      *uint64 // block_number <= ToB
	ToAddr   *common.Address
	FromAddr *common.Address
	Selector *[4]byte // function selector
}

// Logger is the interface used to log errors in DB operations.
type Logger interface {
	Errorf(string, ...any)
}

// RetryConfig holds timing parameters for retry operations.
type RetryConfig struct {
	// MaxDuration is the total time budget for all attempts.
	MaxDuration time.Duration
	// InitialInterval is the wait before the first retry.
	InitialInterval time.Duration
	// Multiplier scales the interval after each failed attempt.
	Multiplier float64
	// MaxInterval caps the per-attempt wait.
	MaxInterval time.Duration
}

// DefaultRetryConfig returns the default retry timing configuration.
func DefaultRetryConfig() RetryConfig {
	return RetryConfig{
		MaxDuration:     15 * time.Second,
		InitialInterval: 500 * time.Millisecond,
		Multiplier:      1.5,
		MaxInterval:     60 * time.Second,
	}
}

// PermanentError wraps an error that should not be retried.
// Use [Permanent] to create one.
type PermanentError struct {
	Err error
}

func (e *PermanentError) Error() string { return e.Err.Error() }
func (e *PermanentError) Unwrap() error { return e.Err }

// Permanent marks an error as non-retryable.
// The retry wrapper unwraps it before returning so callers see the original error.
func Permanent(err error) error {
	return &PermanentError{Err: err}
}

const minRetryInterval = time.Millisecond

// retryWithConfig wraps a query function with exponential backoff retries governed by cfg.
// On each failed attempt it logs the error via retryLogger.
// Errors wrapped with [Permanent] are returned immediately without retrying.
// Context cancellation aborts retries immediately.
func retryWithConfig[F any, P any](
	query func(context.Context, *gorm.DB, P) (F, error),
	errorMsg string,
	cfg RetryConfig,
	retryLogger Logger,
) func(context.Context, *gorm.DB, P) (F, error) {
	if retryLogger == nil {
		retryLogger = logger.Nop{}
	}

	return func(ctx context.Context, db *gorm.DB, params P) (F, error) {
		interval := cfg.InitialInterval
		start := time.Now()

		for {
			result, err := query(ctx, db, params)
			if err == nil {
				return result, nil
			}

			var perm *PermanentError
			if errors.As(err, &perm) {
				return result, perm.Err
			}

			if time.Since(start) >= cfg.MaxDuration {
				return result, err
			}

			if interval < minRetryInterval {
				interval = minRetryInterval
			}

			retryLogger.Errorf("error %s: %v, retrying after %v", errorMsg, err, interval)

			timer := time.NewTimer(interval)
			select {
			case <-timer.C:
			case <-ctx.Done():
				timer.Stop()
				return result, ctx.Err()
			}

			interval = min(time.Duration(float64(interval)*cfg.Multiplier), cfg.MaxInterval)
		}
	}
}

// DB holds a database connection, logger, and retry configuration.
// Use [NewDB] to create one. Methods return copies so WithLogger and WithConfig
// can be chained without modifying the original.
type DB struct {
	db     *gorm.DB
	logger Logger
	cfg    RetryConfig
}

// NewDB returns a DB with the default retry configuration and a no-op logger.
func NewDB(db *gorm.DB) *DB {
	return &DB{
		db:     db,
		logger: logger.Nop{},
		cfg:    DefaultRetryConfig(),
	}
}

// WithLogger returns a copy of d with logger set.
func (d *DB) WithLogger(l Logger) *DB {
	c := *d
	if l == nil {
		l = logger.Nop{}
	}
	c.logger = l
	return &c
}

// WithConfig returns a copy of d with cfg set.
func (d *DB) WithConfig(cfg RetryConfig) *DB {
	c := *d
	c.cfg = cfg
	return &c
}

// FetchState returns the State row matching params.Name.
func (d *DB) FetchState(ctx context.Context, params StateQuery) (State, error) {
	return retryWithConfig(dbFetchState, "fetching state", d.cfg, d.logger)(ctx, d.db, params)
}

// FetchLogs returns all logs matching the non-nil fields in params, ordered by timestamp ascending.
func (d *DB) FetchLogs(ctx context.Context, params LogsQuery) ([]Log, error) {
	return retryWithConfig(dbFetchLogs, "fetching logs", d.cfg, d.logger)(ctx, d.db, params)
}

// FetchTransactions returns all transactions matching the non-nil fields in params, ordered by timestamp ascending.
func (d *DB) FetchTransactions(ctx context.Context, params TxQuery) ([]Transaction, error) {
	return retryWithConfig(dbFetchTransactions, "fetching transactions", d.cfg, d.logger)(ctx, d.db, params)
}

func dbFetchState(ctx context.Context, db *gorm.DB, params StateQuery) (State, error) {
	var state State

	err := db.WithContext(ctx).Where("name = ?", params.Name).Take(&state).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return State{}, Permanent(ErrNotFound)
	}

	return state, err
}

func dbFetchLogs(ctx context.Context, db *gorm.DB, params LogsQuery) ([]Log, error) {
	q := db.WithContext(ctx)

	if params.From != nil {
		q = q.Where("timestamp > ?", *params.From)
	}
	if params.To != nil {
		q = q.Where("timestamp <= ?", *params.To)
	}
	if params.FromB != nil {
		q = q.Where("block_number > ?", *params.FromB)
	}
	if params.ToB != nil {
		q = q.Where("block_number <= ?", *params.ToB)
	}
	if params.Address != nil {
		q = q.Where("address = ?", hex.EncodeToString(params.Address[:]))
	}
	if params.Topic0 != nil {
		q = q.Where("topic0 = ?", hex.EncodeToString(params.Topic0[:]))
	}
	if params.Topic1 != nil {
		q = q.Where("topic1 = ?", hex.EncodeToString(params.Topic1[:]))
	}
	if params.Topic2 != nil {
		q = q.Where("topic2 = ?", hex.EncodeToString(params.Topic2[:]))
	}
	if params.Topic3 != nil {
		q = q.Where("topic3 = ?", hex.EncodeToString(params.Topic3[:]))
	}

	var logs []Log

	err := q.Order("timestamp, id").Find(&logs).Error

	return logs, err
}

func dbFetchTransactions(ctx context.Context, db *gorm.DB, params TxQuery) ([]Transaction, error) {
	q := db.WithContext(ctx)

	if params.From != nil {
		q = q.Where("timestamp > ?", *params.From)
	}
	if params.To != nil {
		q = q.Where("timestamp <= ?", *params.To)
	}
	if params.FromB != nil {
		q = q.Where("block_number > ?", *params.FromB)
	}
	if params.ToB != nil {
		q = q.Where("block_number <= ?", *params.ToB)
	}
	if params.ToAddr != nil {
		q = q.Where("to_address = ?", hex.EncodeToString(params.ToAddr[:]))
	}
	if params.FromAddr != nil {
		q = q.Where("from_address = ?", hex.EncodeToString(params.FromAddr[:]))
	}
	if params.Selector != nil {
		q = q.Where("function_sig = ?", hex.EncodeToString(params.Selector[:]))
	}

	var txs []Transaction

	err := q.Order("timestamp, id").Find(&txs).Error

	return txs, err
}
