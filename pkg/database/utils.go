package database

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	var gormLogLevel logger.LogLevel
	if cfg.LogQueries {
		gormLogLevel = logger.Info
	} else {
		gormLogLevel = logger.Silent
	}
	gormConfig := gorm.Config{
		Logger: logger.Default.LogMode(gormLogLevel),
	}
	return gorm.Open(gormMysql.Open(dbConfig.FormatDSN()), &gormConfig)
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
