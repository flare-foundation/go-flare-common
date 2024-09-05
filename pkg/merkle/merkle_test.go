package merkle_test

import (
	"fmt"
	"testing"

	"github.com/bradleyjkemp/cupaloy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/flarenetwork/libs/go-flare-common/pkg/merkle"
)

func TestEmptyTree(t *testing.T) {
	tree := merkle.Tree{}
	_, err := tree.Root()
	assert.Equal(t, err, merkle.ErrEmptyTree)

	treeSlice := tree
	assert.Len(t, treeSlice, 0)

	numLeaves := tree.LeavesCount()
	assert.Equal(t, numLeaves, 0)

	sortedHashes := tree.Leaves()
	assert.Len(t, sortedHashes, 0)

	_, err = tree.GetLeaf(0)
	assert.Equal(t, err, merkle.ErrInvalidIndex)

	_, err = tree.GetProof(0)
	assert.Equal(t, err, merkle.ErrInvalidIndex)

	_, err = tree.GetProofFromHash(common.HexToHash("0x01"))
	assert.Equal(t, err, merkle.ErrHashNotFound)
}

func TestBuildEmpty(t *testing.T) {
	hashes := []common.Hash{}
	tree := merkle.Build(hashes, false)

	if len(tree) > 0 {
		t.Error("not empty tree")
	}

	treeHashed := merkle.Build(hashes, true)

	if len(treeHashed) > 0 {
		t.Error("not empty tree")
	}

}

func TestBuildOne(t *testing.T) {
	hashes := []common.Hash{common.HexToHash("0x01")}
	tree := merkle.Build(hashes, false)

	if len(tree) != 1 {
		t.Error("tree too big")
	}
}

func TestSingleLeafTree(t *testing.T) {
	val := common.HexToHash("0x01")
	tree := merkle.Tree{val}

	root, err := tree.Root()
	require.NoError(t, err)
	assert.Equal(t, root, val)

	treeSlice := tree
	assert.Len(t, treeSlice, 1)
	assert.Equal(t, treeSlice[0], val)

	numLeaves := tree.LeavesCount()
	assert.Equal(t, numLeaves, 1)

	sortedHashes := tree.Leaves()
	assert.Len(t, sortedHashes, 1)
	assert.Equal(t, sortedHashes[0], val)

	hash, err := tree.GetLeaf(0)
	require.NoError(t, err)
	assert.Equal(t, hash, val)

	proof, err := tree.GetProof(0)
	require.NoError(t, err)
	require.Len(t, proof, 0)

	verified := merkle.VerifyProof(val, proof, root)
	assert.True(t, verified)

	proof, err = tree.GetProofFromHash(val)
	require.NoError(t, err)
	require.Len(t, proof, 0)

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
	assert.Equal(t, numLeaves, 5)

	sortedHashes := tree.Leaves()
	t.Run("SortedHashes", func(t *testing.T) {
		assert.Len(t, sortedHashes, 5)
		cupaloy.SnapshotT(t, sortedHashes)
	})

	for i := range sortedHashes {
		hash, err := tree.GetLeaf(i)
		require.NoError(t, err)
		assert.Equal(t, hash, sortedHashes[i])
	}

	for i, hash := range sortedHashes {
		t.Run(fmt.Sprintf("Proof_%d", i), func(t *testing.T) {
			proof, err := tree.GetProof(i)
			require.NoError(t, err)
			cupaloy.SnapshotT(t, proof)

			verified := merkle.VerifyProof(hash, proof, root)
			assert.True(t, verified)
		})

		t.Run(fmt.Sprintf("ProofFromHash_%d", i), func(t *testing.T) {
			proof, err := tree.GetProofFromHash(hash)
			require.NoError(t, err)
			cupaloy.SnapshotT(t, proof)

			verified := merkle.VerifyProof(hash, proof, root)
			assert.True(t, verified)
		})
	}
}
