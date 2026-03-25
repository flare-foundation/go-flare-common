// Package events provides utilities for converting database event logs to chain event logs.
package events

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/flare-foundation/go-flare-common/pkg/database"
)

const null = "NULL"

// ConvertDatabaseLogToChainLog converts a database event log to a chain event log for use in the log decoder.
// It only converts the fields used by the abigen log decoder (Topics and Data).
func ConvertDatabaseLogToChainLog(dbLog database.Log) (*types.Log, error) {
	data, err := hex.DecodeString(dbLog.Data)
	if err != nil {
		return nil, err
	}

	var topics []common.Hash

	if dbLog.Topic0 != null {
		topics = append(topics, common.HexToHash(dbLog.Topic0))
	}
	if dbLog.Topic1 != null {
		topics = append(topics, common.HexToHash(dbLog.Topic1))
	}
	if dbLog.Topic2 != null {
		topics = append(topics, common.HexToHash(dbLog.Topic2))
	}
	if dbLog.Topic3 != null {
		topics = append(topics, common.HexToHash(dbLog.Topic3))
	}
	return &types.Log{
		Topics:         topics,
		Data:           data,
		BlockNumber:    dbLog.BlockNumber,
		BlockTimestamp: dbLog.Timestamp,
		Address:        common.HexToAddress(dbLog.Address),
		Index:          uint(dbLog.LogIndex),
		TxHash:         common.HexToHash(dbLog.TransactionHash),
	}, nil
}
