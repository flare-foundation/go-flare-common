package policy

import (
	"gitlab.com/ryancollingham/flare-common/pkg/contracts/relay"
	"gitlab.com/ryancollingham/flare-common/pkg/database"
	"gitlab.com/ryancollingham/flare-common/pkg/events"
	"gitlab.com/ryancollingham/flare-common/pkg/logger"

	"github.com/ethereum/go-ethereum/common"
)

var RelayFilterer *relay.RelayFilterer

var log = logger.GetLogger()

func init() {
	var err error
	RelayFilterer, err = relay.NewRelayFilterer(common.Address{}, nil)
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
