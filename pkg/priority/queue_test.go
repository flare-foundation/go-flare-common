package priority

import (
	"math/rand/v2"
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/heapt"
	"github.com/stretchr/testify/require"
)

func TestPushPopRandom(t *testing.T) {
	size := 25

	for j := 0; j < 1; j++ {
		list := rand.Perm(size)
		queue := Queue[int]{}

		for i := range list {
			heapt.Push(&queue, &Item[int]{
				value:  list[i],
				weight: list[i],
			})
		}

		for i := 0; i < size; i++ {
			item, _ := heapt.Pop(&queue)
			require.Equal(t, size-1-i, item.value)
			require.Equal(t, size-1-i, item.weight)
		}
	}
}

func TestAddValuePopRandom(t *testing.T) {
	size := 25

	for j := 0; j < 1; j++ {
		list := rand.Perm(size)
		queue := Queue[int]{}

		for i := range list {
			queue.AddValue(list[i], list[i])
		}

		for i := 0; i < size; i++ {
			item, _ := heapt.Pop(&queue)
			require.Equal(t, size-1-i, item.value)
			require.Equal(t, size-1-i, item.weight)
		}
	}
}
