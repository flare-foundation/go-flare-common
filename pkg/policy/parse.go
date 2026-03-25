package policy

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-foundation/go-flare-common/pkg/contracts/registry"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-foundation/go-flare-common/pkg/database"
	"github.com/flare-foundation/go-flare-common/pkg/events"
)

// RelayFilterer is the shared filterer for parsing Relay contract events.
var RelayFilterer *relay.RelayFilterer

// RegistryFilterer is the shared filterer for parsing Registry contract events.
var RegistryFilterer *registry.RegistryFilterer

func init() {
	var err error
	RelayFilterer, err = relay.NewRelayFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}

	RegistryFilterer, err = registry.NewRegistryFilterer(common.Address{}, nil)
	if err != nil {
		panic(err)
	}
}

// ParseSigningPolicyInitializedEvent parses a SigningPolicyInitialized event from a database log.
func ParseSigningPolicyInitializedEvent(dbLog database.Log) (*relay.RelaySigningPolicyInitialized, error) {
	contractLog, err := events.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, fmt.Errorf("converting database log: %w", err)
	}

	return RelayFilterer.ParseSigningPolicyInitialized(*contractLog)
}

// ParseVoterRegisteredEvent parses a VoterRegistered event from a database log.
func ParseVoterRegisteredEvent(dbLog database.Log) (*registry.RegistryVoterRegistered, error) {
	contractLog, err := events.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, fmt.Errorf("converting database log: %w", err)
	}

	return RegistryFilterer.ParseVoterRegistered(*contractLog)
}
