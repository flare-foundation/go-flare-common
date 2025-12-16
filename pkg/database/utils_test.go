package database

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCheckDelay(t *testing.T) {
	tolerance := 1 * time.Minute

	tests := []struct {
		name          string
		setupDB       func(t *testing.T, db *gorm.DB)
		expectedError string
	}{
		{
			name: "DB in sync",
			setupDB: func(t *testing.T, db *gorm.DB) {
				t.Helper()

				require.NoError(t, db.AutoMigrate(&State{}))
				// Block timestamp is recent, so delay is less than tolerance
				inSyncTimestamp := time.Now().Add(-30 * time.Second).Unix()
				state := State{
					Name:           "last_database_block",
					BlockTimestamp: uint64(inSyncTimestamp),
				}
				require.NoError(t, db.Create(&state).Error)
			},
			expectedError: "",
		},
		{
			name: "DB out of sync",
			setupDB: func(t *testing.T, db *gorm.DB) {
				t.Helper()

				require.NoError(t, db.AutoMigrate(&State{}))
				// Block timestamp is old, so delay is more than tolerance
				outOfSyncTimestamp := time.Now().Add(-2 * time.Minute).Unix()
				state := State{
					Name:           "last_database_block",
					BlockTimestamp: uint64(outOfSyncTimestamp),
				}
				require.NoError(t, db.Create(&state).Error)
			},
			expectedError: "database out of sync for",
		},
		{
			name: "DB error on fetch",
			setupDB: func(t *testing.T, db *gorm.DB) {
				t.Helper()
			},
			expectedError: "database error: no such table: states",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			db, _ := InMemoryDB(t, test.name)
			if test.setupDB != nil {
				test.setupDB(t, db)
			}

			// Close the db connection after the test
			sqlDB, err := db.DB()
			require.NoError(t, err)
			defer sqlDB.Close() //nolint:errcheck

			err = CheckDelay(t.Context(), db, tolerance)

			if test.expectedError == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func InMemoryDB(t *testing.T, name string) (*gorm.DB, string) {
	t.Helper()

	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", name)
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	return db, dsn
}
