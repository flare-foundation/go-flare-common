package offers_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	offersv1 "github.com/flare-foundation/go-flare-common/pkg/contracts/offers"
	"github.com/flare-foundation/go-flare-common/pkg/contracts/v2/offers"
)

// The v2 ABI renames _offers, which must not change the encoding.
func TestPackOfferRewardsMatchesV1(t *testing.T) {
	v1, err := offersv1.OffersMetaData.GetAbi()
	require.NoError(t, err)

	offer := offers.IFtsoRewardOffersManagerOffer{
		Amount:                    big.NewInt(1000),
		FeedId:                    [21]byte{1, 2, 3},
		MinRewardedTurnoutBIPS:    50,
		PrimaryBandRewardSharePPM: big.NewInt(700000),
		SecondaryBandWidthPPM:     big.NewInt(2000),
		ClaimBackAddress:          common.HexToAddress("0x01"),
	}

	want, err := v1.Pack("offerRewards", big.NewInt(12), []offersv1.IFtsoRewardOffersManagerOffer{offersv1.IFtsoRewardOffersManagerOffer(offer)})
	require.NoError(t, err)
	require.Equal(t, want, offers.NewOffers().PackOfferRewards(big.NewInt(12), []offers.IFtsoRewardOffersManagerOffer{offer}))
}
