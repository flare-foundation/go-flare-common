package policy

import (
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/contracts/registry"
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/contracts/relay"
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/database"
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/events"
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/logger"

	"github.com/ethereum/go-ethereum/common"
)

var RelayFilterer *relay.RelayFilterer

var RegistryFilterer *registry.RegistryFilterer

var log = logger.GetLogger()

func init() {
	var err error
	RelayFilterer, err = relay.NewRelayFilterer(common.Address{}, nil)
	if err != nil {
		log.Panic(err)
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
