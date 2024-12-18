package voters_test

import (
	"math/big"
	"slices"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/flare-foundation/go-flare-common/pkg/voters"
)

var (
	testVoters = []common.Address{
		common.HexToAddress("0xc783df8a850f42e7f7e57013759c285caa701eb6"),
		common.HexToAddress("0xead9c93b79ae7c1591b1fb5323bd777e86e150d4"),
		common.HexToAddress("0xe5904695748fe4a84b40b3fc79de2277660bd1d3"),
		common.HexToAddress("0x92561f28ec438ee9831d00d1d59fbdc981b762b2"),
		common.HexToAddress("0x2ffd013aaa7b5a7da93336c2251075202b33fb2b"),
	}

	testWeights = []uint16{100, 200, 300, 400, 500}
)

func TestInitialHashSeed(t *testing.T) {
	seed := voters.InitialHashSeed(big.NewInt(1), 2, 3)
	if seed != common.HexToHash("0x6e0c627900b24bd432fe7b1f713f1b0744091a646a9fe4a65a18dfed21f2949c") {
		t.Errorf("initial hash seed is not correct")
	}
}

func TestVoterSetInitialization(t *testing.T) {
	vs := voters.NewSet(testVoters, testWeights, nil)
	if vs == nil {
		t.Errorf("voter set is nil")
	} else if vs.TotalWeight != 1500 {
		t.Errorf("total weight is not correct")
	}
}

func TestBinarySearch(t *testing.T) {
	testPairs := []uint16{0, 1, 99, 100, 101, 105, 299, 300, 301, 305, 599, 600, 601, 605, 999, 1000, 1001, 1005}

	t.Run("test1", func(t *testing.T) {
		vs := voters.NewSet([]common.Address{common.HexToAddress("0xc783df8a850f42e7f7e57013759c285caa701eb6")}, []uint16{100}, nil)
		testResults := make([]int, 4)
		for i := 0; i <= 3; i++ {
			testResults[i] = vs.BinarySearch(testPairs[i])
		}
		cupaloy.SnapshotT(t, testResults)
	})

	t.Run("test2", func(t *testing.T) {
		vs := voters.NewSet(testVoters, testWeights, nil)
		test2Results := make([]int, len(testPairs))
		for i := 0; i < len(testPairs); i++ {
			test2Results[i] = vs.BinarySearch(testPairs[i])
		}
		cupaloy.SnapshotT(t, test2Results)
	})
}

func TestSelectVoters(t *testing.T) {
	vs := voters.NewSet(testVoters, testWeights, nil)
	seed := voters.InitialHashSeed(big.NewInt(1), 1, 1)
	voterSet, err := vs.RandomSelectThresholdWeightVoters(seed, 3000)

	voterSetHex := make([]string, 0)
	for key := range voterSet {
		voterSetHex = append(voterSetHex, key.Hex())
	}

	slices.Sort(voterSetHex)
	require.NoError(t, err)
	cupaloy.SnapshotT(t, voterSetHex)
}
