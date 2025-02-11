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

type SyncParams struct {
	Retries            int           // Maximal number of retires
	OutOfSyncTolerance time.Duration // Delay that is tolerable
	MaxSleepTime       time.Duration // Maximal duration of sleep between retries
	MinSleepTime       time.Duration // Minimal duration of sleep between retries
}

// WaitToSync waits for C-chain indexer DB to sync.
//
// If db is not up to date, the check is performed again after 1/20 of the delay (bound by MaxSleepTime and MinSleepTime).
// Reties specifies at most how many times is the check done.
// If the check does not succeed by then, panic is raised.
func WaitCIndexerToSync(ctx context.Context, db *gorm.DB, params SyncParams) {
	k := 0
	for k < params.Retries {
		if k > 0 {
			logger.Debugf("Checking database for %v/%v time", k, params.Retries)
		}
		state, err := FetchState(ctx, db, nil)
		if err != nil {
			logger.Panic("database error:", err)
		}

		dbTime := time.Unix(int64(state.BlockTimestamp), 0)

		outOfSync := time.Since(dbTime)
		if outOfSync < params.OutOfSyncTolerance {
			logger.Debug("Database in sync")
			return
		}

		logger.Warnf("Database out of sync. Delayed for %v", outOfSync)
		sleepTime := min(params.MaxSleepTime, outOfSync/20)
		sleepTime = max(sleepTime, params.MinSleepTime)
		logger.Warnf("Sleeping for %v", sleepTime)
		k++
		time.Sleep(sleepTime)
	}

	logger.Warnf("Checking database for the final time")
	state, err := FetchState(ctx, db, nil)
	if err != nil {
		logger.Panic("database error:", err)
	}

	dbTime := time.Unix(int64(state.BlockTimestamp), 0)
	outOfSync := time.Since(dbTime)
	if outOfSync > params.OutOfSyncTolerance {
		logger.Panic("Database out of sync after %v retries. Delayed for %v", params.Retries, outOfSync)
	}
	logger.Debug("Database in sync")
}

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
