package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// PoolConfig tunes the underlying database/sql connection pool. Zero values
// keep the stdlib defaults (unlimited open/idle, no lifetime/idletime cap).
type PoolConfig struct {
	MaxOpenConns    int           `toml:"max_open_conns" envconfig:"DB_MAX_OPEN_CONNS"`
	MaxIdleConns    int           `toml:"max_idle_conns" envconfig:"DB_MAX_IDLE_CONNS"`
	ConnMaxLifetime time.Duration `toml:"conn_max_lifetime" envconfig:"DB_CONN_MAX_LIFETIME"`
	ConnMaxIdleTime time.Duration `toml:"conn_max_idle_time" envconfig:"DB_CONN_MAX_IDLE_TIME"`
}

// Config holds MySQL database connection parameters.
type Config struct {
	Host       string     `toml:"host" envconfig:"DB_HOST"`
	Port       int        `toml:"port" envconfig:"DB_PORT"`
	Database   string     `toml:"database" envconfig:"DB_DATABASE"`
	Username   string     `toml:"username" envconfig:"DB_USERNAME"`
	Password   string     `toml:"password" envconfig:"DB_PASSWORD"`
	LogQueries bool       `toml:"log_queries"`
	Pool       PoolConfig `toml:"pool"`
}

// Connect returns a gorm.DB specified in the config.
func Connect(cfg *Config) (*gorm.DB, error) {
	dbConfig := mysql.Config{
		User:                 cfg.Username,
		Passwd:               cfg.Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		DBName:               cfg.Database,
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	var gormLogLevel gormLogger.LogLevel
	if cfg.LogQueries {
		gormLogLevel = gormLogger.Info
	} else {
		gormLogLevel = gormLogger.Silent
	}
	gormConfig := gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogLevel),
	}
	db, err := gorm.Open(gormMysql.Open(dbConfig.FormatDSN()), &gormConfig)
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("getting underlying sql.DB: %w", err)
	}
	applyPoolConfig(sqlDB, cfg.Pool)

	return db, nil
}

// applyPoolConfig sets database/sql pool knobs from p. Zero values are left untouched.
func applyPoolConfig(sqlDB *sql.DB, p PoolConfig) {
	if p.MaxOpenConns > 0 {
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	}
	if p.MaxIdleConns > 0 {
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	}
	if p.ConnMaxLifetime > 0 {
		sqlDB.SetConnMaxLifetime(p.ConnMaxLifetime)
	}
	if p.ConnMaxIdleTime > 0 {
		sqlDB.SetConnMaxIdleTime(p.ConnMaxIdleTime)
	}
}

// SyncParams configures the retry behavior for waiting on C-chain indexer synchronization.
type SyncParams struct {
	Retries            int           // Maximal number of retries
	OutOfSyncTolerance time.Duration // Delay that is tolerable
	MaxSleepTime       time.Duration // Maximal duration of sleep between retries
	MinSleepTime       time.Duration // Minimal duration of sleep between retries
}

// WaitCIndexerToSync waits for C-chain indexer DB to sync.
//
// If db is not up to date, the check is performed again after 1/20 of the delay (bound by MaxSleepTime and MinSleepTime).
// Retries specifies at most how many times the check is done.
// If the check does not succeed by then, error is returned.
//
// Logger for logging can be provided. If it is nil, no logging is done.
func WaitCIndexerToSync(ctx context.Context, db *gorm.DB, params SyncParams, l syncLogger) error {
	if l == nil {
		l = logger.Nop{}
	}

	k := 0
	for k < params.Retries {
		if k > 0 {
			l.Debugf("Checking database for %v/%v time", k, params.Retries+1)
		}
		state, err := FetchState(ctx, db, nil)
		if err != nil {
			return fmt.Errorf("fetching state: %w", err)
		}

		dbTime := time.Unix(int64(state.BlockTimestamp), 0)

		outOfSync := time.Since(dbTime)
		if outOfSync < params.OutOfSyncTolerance {
			l.Debug("Database in sync")
			return nil
		}

		l.Warnf("Database out of sync. Delayed for %v", outOfSync)
		sleepTime := min(params.MaxSleepTime, outOfSync/20)
		sleepTime = max(sleepTime, params.MinSleepTime)
		l.Warnf("Sleeping for %v", sleepTime)
		k++
		select {
		case <-time.After(sleepTime):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	l.Warnf("Checking database for the final time")
	state, err := FetchState(ctx, db, nil)
	if err != nil {
		return fmt.Errorf("fetching state: %w", err)
	}

	dbTime := time.Unix(int64(state.BlockTimestamp), 0)
	outOfSync := time.Since(dbTime)
	if outOfSync > params.OutOfSyncTolerance {
		return fmt.Errorf("database out of sync after %v retries. Delayed for %v", params.Retries, outOfSync)
	}
	l.Debug("Database in sync")
	return nil
}

// syncLogger is an interface for logging used by WaitCIndexerToSync.
type syncLogger interface {
	Debug(...any)
	Debugf(string, ...any)
	Warnf(string, ...any)
}

// DoInTransaction executes operations within a single database transaction,
// rolling back on any returned error or panic. A panic from inside an
// operation is converted into a returned error so the caller does not treat
// a panicked-and-rolled-back transaction as committed. Begin and Rollback
// errors are surfaced rather than silently dropped.
func DoInTransaction(db *gorm.DB, operations ...func(db *gorm.DB) error) (err error) {
	tx := db.Begin()
	if tx.Error != nil {
		return fmt.Errorf("beginning transaction: %w", tx.Error)
	}
	defer func() {
		if r := recover(); r != nil {
			rerr := tx.Rollback().Error
			if rerr != nil {
				err = fmt.Errorf("panic in transaction: %v (rollback failed: %v)", r, rerr)
				return
			}
			err = fmt.Errorf("panic in transaction: %v", r)
		}
	}()

	for _, f := range operations {
		if opErr := f(tx); opErr != nil {
			if rerr := tx.Rollback().Error; rerr != nil {
				return fmt.Errorf("%w (rollback failed: %v)", opErr, rerr)
			}
			return opErr
		}
	}
	return tx.Commit().Error
}

// CheckDelay checks whether db is delayed by more than tolerance.
func CheckDelay(ctx context.Context, db *gorm.DB, tolerance time.Duration) error {
	state, err := FetchState(ctx, db, nil)
	if err != nil {
		return fmt.Errorf("fetching state: %w", err)
	}

	dbTime := time.Unix(int64(state.BlockTimestamp), 0)

	outOfSync := time.Since(dbTime)
	if outOfSync > tolerance {
		return fmt.Errorf("database out of sync for %v", outOfSync)
	}
	return nil
}
