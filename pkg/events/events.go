// Package events provides utilities for converting database event logs to chain event logs.
package events

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/flare-foundation/go-flare-common/pkg/database"
)

const null = "NULL"

// SelectorFromMetadata returns the event selector (topic0) for the named event from contract metadata.
func SelectorFromMetadata(metaData *bind.MetaData, eventName string) (common.Hash, error) {
	a, err := metaData.GetAbi()
	if err != nil {
		return common.Hash{}, err
	}
	return SelectorFromABI(a, eventName)
}

// SelectorFromABI returns the event selector (topic0) for the named event from a parsed ABI.
func SelectorFromABI(a *abi.ABI, eventName string) (common.Hash, error) {
	ev, ok := a.Events[eventName]
	if !ok {
		return common.Hash{}, fmt.Errorf("event %s not found in ABI", eventName)
	}
	return ev.ID, nil
}

// ConvertDatabaseLogToChainLog converts a database event log to a chain event log for use in the log decoder.
// It only converts the fields used by the abigen log decoder (Topics and Data).
func ConvertDatabaseLogToChainLog(dbLog database.Log) (*types.Log, error) {
	data, err := hex.DecodeString(dbLog.Data)
	if err != nil {
		return nil, fmt.Errorf("decoding log data: %w", err)
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
