// Package events provides utilities for converting database event logs to chain event logs.
package events

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/flare-foundation/go-flare-common/pkg/database"
)

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

// ConvertDatabaseLogToChainLog converts a database event log to a chain event log.
// Delegates to database.Log.ToEthLog so the two converters can't drift.
func ConvertDatabaseLogToChainLog(dbLog database.Log) (*types.Log, error) {
	l, err := dbLog.ToEthLog()
	if err != nil {
		return nil, err
	}
	return &l, nil
}
