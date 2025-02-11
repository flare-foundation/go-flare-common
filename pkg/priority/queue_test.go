package priority

import (
	"math/rand/v2"
	"testing"

	"github.com/flare-foundation/go-flare-common/pkg/heapt"
	"github.com/stretchr/testify/require"
)

type wInt int

func (x wInt) self() wInt {
	return x
}

func (x wInt) Less(y wInt) bool {
	return x < y
}

func TestPushPopRandom(t *testing.T) {
	size := 25

	for j := 0; j < 1; j++ {
		list := rand.Perm(size)
		queue := Queue[int, wInt]{}

		for i := range list {
			heapt.Push(&queue, &Item[int, wInt]{
				value:  list[i],
				weight: wInt(list[i]),
			})
		}

		for i := 0; i < size; i++ {
			item, _ := heapt.Pop(&queue)
			require.Equal(t, size-1-i, item.value)
			require.Equal(t, wInt(size-1-i), item.weight.self())
		}
	}
}

func TestAddValuePopRandom(t *testing.T) {
	size := 25

	for j := 0; j < 1; j++ {
		list := rand.Perm(size)
		queue := Queue[int, wInt]{}

		for i := range list {
			queue.AddValue(list[i], wInt(list[i]))
		}

		for i := 0; i < size; i++ {
			item, _ := heapt.Pop(&queue)
			require.Equal(t, size-1-i, item.value)
			require.Equal(t, wInt(size-1-i), item.weight.self())
		}
	}
}
