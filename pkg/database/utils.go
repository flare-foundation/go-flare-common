package database

import (
	"context"
	"fmt"
	"time"

	"github.com/flare-foundation/go-flare-common/pkg/logger"
	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// Config holds MySQL database connection parameters.
type Config struct {
	Host       string `toml:"host" envconfig:"DB_HOST"`
	Port       int    `toml:"port" envconfig:"DB_PORT"`
	Database   string `toml:"database" envconfig:"DB_DATABASE"`
	Username   string `toml:"username" envconfig:"DB_USERNAME"`
	Password   string `toml:"password" envconfig:"DB_PASSWORD"`
	LogQueries bool   `toml:"log_queries"`
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
	return gorm.Open(gormMysql.Open(dbConfig.FormatDSN()), &gormConfig)
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

// DoInTransaction executes operations within a single database transaction, rolling back on any error.
func DoInTransaction(db *gorm.DB, operations ...func(db *gorm.DB) error) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, f := range operations {
		if err := f(tx); err != nil {
			tx.Rollback()
			return err
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
