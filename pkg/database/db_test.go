package database

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func ptr[T any](v T) *T { return &v }

func newTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, _ := InMemoryDB(t, t.Name())
	require.NoError(t, db.AutoMigrate(&Log{}, &Transaction{}, &State{}))

	return db
}

// ---------------------------------------------------------------------------
// FetchLogs
// ---------------------------------------------------------------------------

func TestDBFetchLogs(t *testing.T) {
	db := newTestDB(t)

	addr1 := common.HexToAddress("0x1111111111111111111111111111111111111111")
	addr2 := common.HexToAddress("0x2222222222222222222222222222222222222222")
	topic := common.HexToHash("0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	topic1 := common.HexToHash("0xcccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc")

	logs := []Log{
		{BaseEntity: BaseEntity{ID: 1}, Address: "1111111111111111111111111111111111111111", Topic0: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Topic1: "cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc", Timestamp: 100, BlockNumber: 10, TransactionHash: "tx1", LogIndex: 0},
		{BaseEntity: BaseEntity{ID: 2}, Address: "1111111111111111111111111111111111111111", Topic0: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", Timestamp: 200, BlockNumber: 20, TransactionHash: "tx2", LogIndex: 0},
		{BaseEntity: BaseEntity{ID: 3}, Address: "2222222222222222222222222222222222222222", Topic0: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", Topic1: "cccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccccc", Timestamp: 300, BlockNumber: 30, TransactionHash: "tx3", LogIndex: 0},
	}
	require.NoError(t, db.Create(&logs).Error)

	tests := []struct {
		name    string
		params  LogsQuery
		wantIDs []uint64
	}{
		{
			name:    "no filter returns all",
			params:  LogsQuery{},
			wantIDs: []uint64{1, 2, 3},
		},
		{
			name:    "filter by address",
			params:  LogsQuery{Address: &addr1},
			wantIDs: []uint64{1, 2},
		},
		{
			name:    "filter by topic0",
			params:  LogsQuery{Topic0: &topic},
			wantIDs: []uint64{1, 3},
		},
		{
			name:    "filter by topic1",
			params:  LogsQuery{Topic1: &topic1},
			wantIDs: []uint64{1, 3},
		},
		{
			name:    "filter by address and topic0",
			params:  LogsQuery{Address: &addr1, Topic0: &topic},
			wantIDs: []uint64{1},
		},
		{
			name:    "filter by timestamp From (exclusive)",
			params:  LogsQuery{From: ptr(uint64(100))},
			wantIDs: []uint64{2, 3},
		},
		{
			name:    "filter by timestamp To (inclusive)",
			params:  LogsQuery{To: ptr(uint64(200))},
			wantIDs: []uint64{1, 2},
		},
		{
			name:    "filter by timestamp range",
			params:  LogsQuery{From: ptr(uint64(100)), To: ptr(uint64(200))},
			wantIDs: []uint64{2},
		},
		{
			name:    "filter by block FromB (exclusive)",
			params:  LogsQuery{FromB: ptr(uint64(10))},
			wantIDs: []uint64{2, 3},
		},
		{
			name:    "filter by block ToB (inclusive)",
			params:  LogsQuery{ToB: ptr(uint64(20))},
			wantIDs: []uint64{1, 2},
		},
		{
			name:    "zero-value address pointer is applied",
			params:  LogsQuery{Address: ptr(common.Address{})},
			wantIDs: []uint64{},
		},
		{
			name:    "no match returns empty",
			params:  LogsQuery{Address: &addr2, Topic0: &topic, From: ptr(uint64(300))},
			wantIDs: []uint64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbFetchLogs(context.Background(), db, tt.params)
			require.NoError(t, err)

			gotIDs := make([]uint64, len(got))
			for i, l := range got {
				gotIDs[i] = l.ID
			}
			require.Equal(t, tt.wantIDs, gotIDs)
		})
	}
}

func TestDBFetchLogsOrderedByTimestamp(t *testing.T) {
	db := newTestDB(t)

	logs := []Log{
		{BaseEntity: BaseEntity{ID: 1}, Timestamp: 300, BlockNumber: 3, TransactionHash: "tx1", LogIndex: 0},
		{BaseEntity: BaseEntity{ID: 2}, Timestamp: 100, BlockNumber: 1, TransactionHash: "tx2", LogIndex: 0},
		{BaseEntity: BaseEntity{ID: 3}, Timestamp: 200, BlockNumber: 2, TransactionHash: "tx3", LogIndex: 0},
	}
	require.NoError(t, db.Create(&logs).Error)

	got, err := dbFetchLogs(context.Background(), db, LogsQuery{})
	require.NoError(t, err)
	require.Equal(t, []uint64{2, 3, 1}, []uint64{got[0].ID, got[1].ID, got[2].ID})
}

func TestDBFetchLogsMethod(t *testing.T) {
	db := newTestDB(t)

	logs := []Log{
		{BaseEntity: BaseEntity{ID: 1}, Address: "1111111111111111111111111111111111111111", Timestamp: 10, TransactionHash: "tx1", LogIndex: 0},
	}
	require.NoError(t, db.Create(&logs).Error)

	got, err := NewDB(db).FetchLogs(context.Background(), LogsQuery{})
	require.NoError(t, err)
	require.Len(t, got, 1)
	require.Equal(t, uint64(1), got[0].ID)
}

// ---------------------------------------------------------------------------
// FetchTransactions
// ---------------------------------------------------------------------------

func TestDBFetchTransactions(t *testing.T) {
	db := newTestDB(t)

	toAddr := common.HexToAddress("0xAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	fromAddr := common.HexToAddress("0xBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	sel := [4]byte{0xde, 0xad, 0xbe, 0xef}

	txs := []Transaction{
		{BaseEntity: BaseEntity{ID: 1}, ToAddress: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", FromAddress: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", FunctionSig: "deadbeef", Timestamp: 100, BlockNumber: 10, Hash: "h1"},
		{BaseEntity: BaseEntity{ID: 2}, ToAddress: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", FromAddress: "cccccccccccccccccccccccccccccccccccccccc", FunctionSig: "deadbeef", Timestamp: 200, BlockNumber: 20, Hash: "h2"},
		{BaseEntity: BaseEntity{ID: 3}, ToAddress: "dddddddddddddddddddddddddddddddddddddddd", FromAddress: "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", FunctionSig: "cafebabe", Timestamp: 300, BlockNumber: 30, Hash: "h3"},
	}
	require.NoError(t, db.Create(&txs).Error)

	tests := []struct {
		name    string
		params  TxQuery
		wantIDs []uint64
	}{
		{
			name:    "no filter returns all",
			params:  TxQuery{},
			wantIDs: []uint64{1, 2, 3},
		},
		{
			name:    "filter by ToAddr",
			params:  TxQuery{ToAddr: &toAddr},
			wantIDs: []uint64{1, 2},
		},
		{
			name:    "filter by FromAddr",
			params:  TxQuery{FromAddr: &fromAddr},
			wantIDs: []uint64{1, 3},
		},
		{
			name:    "filter by selector",
			params:  TxQuery{Selector: &sel},
			wantIDs: []uint64{1, 2},
		},
		{
			name:    "filter by timestamp range",
			params:  TxQuery{From: ptr(uint64(100)), To: ptr(uint64(200))},
			wantIDs: []uint64{2},
		},
		{
			name:    "filter by block range",
			params:  TxQuery{FromB: ptr(uint64(10)), ToB: ptr(uint64(20))},
			wantIDs: []uint64{2},
		},
		{
			name:    "combined filters",
			params:  TxQuery{ToAddr: &toAddr, Selector: &sel},
			wantIDs: []uint64{1, 2},
		},
		{
			name:    "no match returns empty",
			params:  TxQuery{ToAddr: &toAddr, From: ptr(uint64(300))},
			wantIDs: []uint64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbFetchTransactions(context.Background(), db, tt.params)
			require.NoError(t, err)

			gotIDs := make([]uint64, len(got))
			for i, tx := range got {
				gotIDs[i] = tx.ID
			}
			require.Equal(t, tt.wantIDs, gotIDs)
		})
	}
}

func TestDBFetchTransactionsOrderedByTimestamp(t *testing.T) {
	db := newTestDB(t)

	txs := []Transaction{
		{BaseEntity: BaseEntity{ID: 1}, Hash: "h1", Timestamp: 300},
		{BaseEntity: BaseEntity{ID: 2}, Hash: "h2", Timestamp: 100},
		{BaseEntity: BaseEntity{ID: 3}, Hash: "h3", Timestamp: 200},
	}
	require.NoError(t, db.Create(&txs).Error)

	got, err := dbFetchTransactions(context.Background(), db, TxQuery{})
	require.NoError(t, err)
	require.Equal(t, []uint64{2, 3, 1}, []uint64{got[0].ID, got[1].ID, got[2].ID})
}

func TestDBFetchTransactionsMethod(t *testing.T) {
	db := newTestDB(t)

	txs := []Transaction{
		{BaseEntity: BaseEntity{ID: 1}, Hash: "h1", Timestamp: 10},
	}
	require.NoError(t, db.Create(&txs).Error)

	got, err := NewDB(db).FetchTransactions(context.Background(), TxQuery{})
	require.NoError(t, err)
	require.Len(t, got, 1)
	require.Equal(t, uint64(1), got[0].ID)
}

// ---------------------------------------------------------------------------
// FetchState
// ---------------------------------------------------------------------------

func TestDBFetchState(t *testing.T) {
	db := newTestDB(t)

	states := []State{
		{BaseEntity: BaseEntity{ID: 1}, Name: string(StateFirstDBBlock), Index: 1, BlockTimestamp: 1000},
		{BaseEntity: BaseEntity{ID: 2}, Name: string(StateLastDBBlock), Index: 99, BlockTimestamp: 9900},
		{BaseEntity: BaseEntity{ID: 3}, Name: string(StateLastChain), Index: 105, BlockTimestamp: 10500},
	}
	require.NoError(t, db.Create(&states).Error)

	tests := []struct {
		name      string
		params    StateQuery
		wantIndex uint64
	}{
		{name: "first db block", params: StateQuery{Name: StateFirstDBBlock}, wantIndex: 1},
		{name: "last db block", params: StateQuery{Name: StateLastDBBlock}, wantIndex: 99},
		{name: "last chain block", params: StateQuery{Name: StateLastChain}, wantIndex: 105},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dbFetchState(context.Background(), db, tt.params)
			require.NoError(t, err)
			require.Equal(t, tt.wantIndex, got.Index)
		})
	}
}

func TestDBFetchStateNotFound(t *testing.T) {
	db := newTestDB(t)

	_, err := dbFetchState(context.Background(), db, StateQuery{Name: StateLastDBBlock})
	require.ErrorIs(t, err, ErrNotFound)
}

func TestDBFetchStateMethod(t *testing.T) {
	db := newTestDB(t)

	states := []State{
		{BaseEntity: BaseEntity{ID: 1}, Name: string(StateLastDBBlock), Index: 42, BlockTimestamp: 5000},
	}
	require.NoError(t, db.Create(&states).Error)

	got, err := NewDB(db).FetchState(context.Background(), StateQuery{Name: StateLastDBBlock})
	require.NoError(t, err)
	require.Equal(t, uint64(42), got.Index)
}

func TestDBFetchStateNotFoundMethod(t *testing.T) {
	db := newTestDB(t)

	_, err := NewDB(db).FetchState(context.Background(), StateQuery{Name: StateLastDBBlock})
	require.ErrorIs(t, err, ErrNotFound)
}

// ---------------------------------------------------------------------------
// retryWithConfig
// ---------------------------------------------------------------------------

func TestRetryWithConfigSuccessOnFirstAttempt(t *testing.T) {
	calls := 0
	query := func(_ context.Context, _ *gorm.DB, _ int) (string, error) {
		calls++
		return "ok", nil
	}

	wrapped := retryWithConfig(query, "test", DefaultRetryConfig(), nil)
	result, err := wrapped(context.Background(), nil, 0)

	require.NoError(t, err)
	require.Equal(t, "ok", result)
	require.Equal(t, 1, calls)
}

func TestRetryWithConfigSuccessAfterRetries(t *testing.T) {
	calls := 0
	query := func(_ context.Context, _ *gorm.DB, _ int) (string, error) {
		calls++
		if calls < 3 {
			return "", errors.New("transient")
		}
		return "ok", nil
	}

	cfg := RetryConfig{
		MaxDuration:     5 * time.Second,
		InitialInterval: time.Millisecond,
		Multiplier:      1.0,
		MaxInterval:     time.Millisecond,
	}

	wrapped := retryWithConfig(query, "test", cfg, nil)
	result, err := wrapped(context.Background(), nil, 0)

	require.NoError(t, err)
	require.Equal(t, "ok", result)
	require.Equal(t, 3, calls)
}

func TestRetryWithConfigMaxDuration(t *testing.T) {
	calls := 0
	query := func(_ context.Context, _ *gorm.DB, _ int) (string, error) {
		calls++
		return "", errors.New("always fails")
	}

	cfg := RetryConfig{
		MaxDuration:     50 * time.Millisecond,
		InitialInterval: 10 * time.Millisecond,
		Multiplier:      1.0,
		MaxInterval:     10 * time.Millisecond,
	}

	wrapped := retryWithConfig(query, "test", cfg, nil)
	_, err := wrapped(context.Background(), nil, 0)

	require.Error(t, err)
	require.Equal(t, "always fails", err.Error())
	require.Greater(t, calls, 1)
}

func TestRetryWithConfigContextCancellation(t *testing.T) {
	calls := 0
	query := func(_ context.Context, _ *gorm.DB, _ int) (string, error) {
		calls++
		return "", errors.New("fail")
	}

	cfg := RetryConfig{
		MaxDuration:     10 * time.Second,
		InitialInterval: time.Second,
		Multiplier:      1.0,
		MaxInterval:     time.Second,
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(20 * time.Millisecond)
		cancel()
	}()

	wrapped := retryWithConfig(query, "test", cfg, nil)
	_, err := wrapped(ctx, nil, 0)

	require.ErrorIs(t, err, context.Canceled)
	require.Equal(t, 1, calls)
}

func TestRetryWithConfigPermanentError(t *testing.T) {
	calls := 0
	query := func(_ context.Context, _ *gorm.DB, _ int) (string, error) {
		calls++
		return "", Permanent(ErrNotFound)
	}

	cfg := RetryConfig{
		MaxDuration:     5 * time.Second,
		InitialInterval: time.Millisecond,
		Multiplier:      1.0,
		MaxInterval:     time.Millisecond,
	}

	wrapped := retryWithConfig(query, "test", cfg, nil)
	_, err := wrapped(context.Background(), nil, 0)

	require.ErrorIs(t, err, ErrNotFound)
	require.Equal(t, 1, calls)
}

func TestRetryWithConfigZeroConfig(t *testing.T) {
	calls := 0
	query := func(_ context.Context, _ *gorm.DB, _ int) (string, error) {
		calls++
		return "", errors.New("fail")
	}

	wrapped := retryWithConfig(query, "test", RetryConfig{}, nil)
	_, err := wrapped(context.Background(), nil, 0)

	require.Error(t, err)
	require.LessOrEqual(t, calls, 2)
}

// ---------------------------------------------------------------------------
// DB struct
// ---------------------------------------------------------------------------

func TestDBWithConfig(t *testing.T) {
	db := newTestDB(t)

	fastCfg := RetryConfig{
		MaxDuration:     50 * time.Millisecond,
		InitialInterval: 5 * time.Millisecond,
		Multiplier:      1.0,
		MaxInterval:     5 * time.Millisecond,
	}

	original := NewDB(db)
	custom := original.WithConfig(fastCfg)

	require.Equal(t, DefaultRetryConfig(), original.cfg)
	require.Equal(t, fastCfg, custom.cfg)
}

func TestDBWithLogger(t *testing.T) {
	db := newTestDB(t)

	var logged []string
	l := &dbTestLogger{msgs: &logged}

	original := NewDB(db)
	custom := original.WithLogger(l)

	require.IsType(t, logger.Nop{}, original.logger)
	require.Equal(t, l, custom.logger)
}

type dbTestLogger struct {
	msgs *[]string
}

func (l *dbTestLogger) Errorf(format string, args ...any) {
	*l.msgs = append(*l.msgs, fmt.Sprintf(format, args...))
}

func TestDBChainedWithLoggerAndConfig(t *testing.T) {
	db := newTestDB(t)

	var logged []string
	l := &dbTestLogger{msgs: &logged}

	fastCfg := RetryConfig{
		MaxDuration:     30 * time.Millisecond,
		InitialInterval: 5 * time.Millisecond,
		Multiplier:      1.0,
		MaxInterval:     5 * time.Millisecond,
	}

	d := NewDB(db).WithLogger(l).WithConfig(fastCfg)

	require.Equal(t, l, d.logger)
	require.Equal(t, fastCfg, d.cfg)

	got, err := d.FetchLogs(context.Background(), LogsQuery{})
	require.NoError(t, err)
	require.Empty(t, got)
}

func TestDBWithLoggerNil(t *testing.T) {
	db := newTestDB(t)

	d := NewDB(db).WithLogger(nil)

	require.IsType(t, logger.Nop{}, d.logger)
}

// ---------------------------------------------------------------------------
// ToEthLog
// ---------------------------------------------------------------------------

func TestToEthLog(t *testing.T) {
	l := Log{
		Data:   "deadbeef",
		Topic0: "0000000000000000000000000000000000000000000000000000000000000001",
		Topic1: "0000000000000000000000000000000000000000000000000000000000000002",
		Topic2: "",
		Topic3: "NULL",
	}

	ethLog, err := l.ToEthLog()
	require.NoError(t, err)
	require.Equal(t, []byte{0xde, 0xad, 0xbe, 0xef}, ethLog.Data)
	require.Len(t, ethLog.Topics, 2)
	require.Equal(t, common.HexToHash("0x01"), ethLog.Topics[0])
	require.Equal(t, common.HexToHash("0x02"), ethLog.Topics[1])
}
