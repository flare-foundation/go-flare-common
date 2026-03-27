package merkle_test

import (
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/flare-foundation/go-flare-common/pkg/merkle"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEmptyTree(t *testing.T) {
	tree := merkle.Tree{}
	_, err := tree.Root()
	assert.Equal(t, merkle.ErrEmptyTree, err)

	treeSlice := tree
	assert.Empty(t, treeSlice)

	numLeaves := tree.LeavesCount()
	assert.Equal(t, 0, numLeaves)

	sortedHashes := tree.Leaves()
	assert.Empty(t, sortedHashes)

	_, err = tree.Leaf(0)
	assert.Equal(t, merkle.ErrInvalidIndex, err)

	_, err = tree.Proof(0)
	assert.Equal(t, merkle.ErrInvalidIndex, err)

	_, err = tree.ProofFromHash(common.HexToHash("0x01"))
	assert.Equal(t, merkle.ErrHashNotFound, err)
}

func TestBuildEmpty(t *testing.T) {
	hashes := []common.Hash{}
	tree := merkle.Build(hashes, false)
	require.Empty(t, tree)

	treeHashed := merkle.Build(hashes, true)
	require.Empty(t, treeHashed)
}

func TestBuildOne(t *testing.T) {
	hashes := []common.Hash{common.HexToHash("0x01")}
	tree := merkle.Build(hashes, false)
	require.Len(t, tree, 1)
}

func TestSingleLeafTree(t *testing.T) {
	val := common.HexToHash("0x01")
	tree := merkle.Tree{val}

	root, err := tree.Root()
	require.NoError(t, err)
	assert.Equal(t, val, root)

	treeSlice := tree
	assert.Len(t, treeSlice, 1)
	assert.Equal(t, val, treeSlice[0])

	numLeaves := tree.LeavesCount()
	assert.Equal(t, 1, numLeaves)

	sortedHashes := tree.Leaves()
	assert.Len(t, sortedHashes, 1)
	assert.Equal(t, val, sortedHashes[0])

	hash, err := tree.Leaf(0)
	require.NoError(t, err)
	assert.Equal(t, val, hash)

	proof, err := tree.Proof(0)
	require.NoError(t, err)
	require.Empty(t, proof)

	verified := merkle.VerifyProof(val, proof, root)
	assert.True(t, verified)

	proof, err = tree.ProofFromHash(val)
	require.NoError(t, err)
	require.Empty(t, proof)

	verified = merkle.VerifyProof(val, proof, root)
	assert.True(t, verified)
}

func TestMultiLeafTree(t *testing.T) {
	vals := []string{
		"0x01", "0x02", "0x03", "0x04", "0x05",
	}

	tree := merkle.BuildFromHex(vals, true)

	root, err := tree.Root()
	require.NoError(t, err)
	cupaloy.SnapshotT(t, root.Hex())

	t.Run("TreeSlice", func(t *testing.T) {
		treeSlice := tree
		assert.Len(t, treeSlice, 9)
		cupaloy.SnapshotT(t, treeSlice)
	})

	numLeaves := tree.LeavesCount()
	assert.Equal(t, 5, numLeaves)

	sortedHashes := tree.Leaves()
	t.Run("SortedHashes", func(t *testing.T) {
		assert.Len(t, sortedHashes, 5)
		cupaloy.SnapshotT(t, sortedHashes)
	})

	for i := range sortedHashes {
		hash, err := tree.Leaf(i)
		require.NoError(t, err)
		assert.Equal(t, sortedHashes[i], hash)
	}

	for i, hash := range sortedHashes {
		t.Run(fmt.Sprintf("Proof_%d", i), func(t *testing.T) {
			proof, err := tree.Proof(i)
			require.NoError(t, err)
			cupaloy.SnapshotT(t, proof)

			verified := merkle.VerifyProof(hash, proof, root)
			assert.True(t, verified)
		})

		t.Run(fmt.Sprintf("ProofFromHash_%d", i), func(t *testing.T) {
			proof, err := tree.ProofFromHash(hash)
			require.NoError(t, err)
			cupaloy.SnapshotT(t, proof)

			verified := merkle.VerifyProof(hash, proof, root)
			assert.True(t, verified)
		})
	}
}

func TestSorting(t *testing.T) {
	vals := [][]string{
		{
			"0x01", "0x02", "0x03", "0x04", "0x05", // 1-5
		},
		{
			"0x05", "0x04", "0x03", "0x02", "0x01", // 1-5 shuffled
		},
		{
			"0x02", "0x01", "0x05", "0x04", "0x03", "0x01", // 1-5 shuffled with duplicated 1
		},
	}

	prevRoot := common.Hash{}

	for i, val := range vals {
		tree := merkle.BuildFromHex(val, false)
		root, err := tree.Root()
		require.NoError(t, err)

		require.Equal(t, 5, tree.LeavesCount())

		if i > 0 {
			require.Equal(t, prevRoot, root)
		}

		prevRoot = root
	}
}
