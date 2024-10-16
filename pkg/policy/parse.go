package policy

import (
	"log"

	"github.com/ethereum/go-ethereum/common"

	"github.com/flare-fundation/go-flare-common/pkg/contracts/registry"
	"github.com/flare-fundation/go-flare-common/pkg/contracts/relay"
	"github.com/flare-fundation/go-flare-common/pkg/database"
	"github.com/flare-fundation/go-flare-common/pkg/events"
	"github.com/flare-fundation/go-flare-common/pkg/logger"
)

var RelayFilterer *relay.RelayFilterer

var RegistryFilterer *registry.RegistryFilterer

func init() {
	var err error
	RelayFilterer, err = relay.NewRelayFilterer(common.Address{}, nil)
	if err != nil {
		logger.Panic(err)
	}

	RegistryFilterer, err = registry.NewRegistryFilterer(common.Address{}, nil)

	if err != nil {
		logger.Panic(err)
	}

	RegistryFilterer, err = registry.NewRegistryFilterer(common.Address{}, nil)

	if err != nil {
		log.Panic(err)
	}
}

func ParseSigningPolicyInitializedEvent(dbLog database.Log) (*relay.RelaySigningPolicyInitialized, error) {
	contractLog, err := events.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}

	return RelayFilterer.ParseSigningPolicyInitialized(*contractLog)
}

func ParseVoterRegisteredEvent(dbLog database.Log) (*registry.RegistryVoterRegistered, error) {
	contractLog, err := events.ConvertDatabaseLogToChainLog(dbLog)
	if err != nil {
		return nil, err
	}

	return RegistryFilterer.ParseVoterRegistered(*contractLog)
}
