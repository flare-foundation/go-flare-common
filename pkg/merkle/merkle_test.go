package merkle_test

import (
	"fmt"
	"testing"

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

	expectedRoot := common.HexToHash("0xc9f523a58f30cdf128d0c9b9b5c47307fffd7d0edcdadf7c623c2b01b48cdf82")
	root, err := tree.Root()
	require.NoError(t, err)
	assert.Equal(t, expectedRoot, root)

	expectedTree := merkle.Tree{
		common.HexToHash("0xc9f523a58f30cdf128d0c9b9b5c47307fffd7d0edcdadf7c623c2b01b48cdf82"),
		common.HexToHash("0x6e106312c68ae7878ba9ba4885ad84fb8f8e153350600d368cedee43b5ce1903"),
		common.HexToHash("0xa95fdf6e1d2ef63e224adadbf70e466c4ac233aaa7fda9afb4862e884c8185f2"),
		common.HexToHash("0xa8d08734ea7322e06e0a776297d6f37272e6f5a616160a77c77841e0759bf0ca"),
		common.HexToHash("0x036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0"),
		common.HexToHash("0x405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace"),
		common.HexToHash("0x8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b"),
		common.HexToHash("0xb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6"),
		common.HexToHash("0xc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b"),
	}
	t.Run("TreeSlice", func(t *testing.T) {
		assert.Len(t, tree, 9)
		assert.Equal(t, expectedTree, tree)
	})

	numLeaves := tree.LeavesCount()
	assert.Equal(t, 5, numLeaves)

	expectedLeaves := []common.Hash{
		common.HexToHash("0x036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0"),
		common.HexToHash("0x405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace"),
		common.HexToHash("0x8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b"),
		common.HexToHash("0xb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6"),
		common.HexToHash("0xc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b"),
	}
	sortedHashes := tree.Leaves()
	t.Run("SortedHashes", func(t *testing.T) {
		assert.Len(t, sortedHashes, 5)
		assert.Equal(t, expectedLeaves, sortedHashes)
	})

	for i := range sortedHashes {
		hash, err := tree.Leaf(i)
		require.NoError(t, err)
		assert.Equal(t, sortedHashes[i], hash)
	}

	expectedProofs := [][]common.Hash{
		{
			common.HexToHash("0xa8d08734ea7322e06e0a776297d6f37272e6f5a616160a77c77841e0759bf0ca"),
			common.HexToHash("0xa95fdf6e1d2ef63e224adadbf70e466c4ac233aaa7fda9afb4862e884c8185f2"),
		},
		{
			common.HexToHash("0x8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b"),
			common.HexToHash("0x6e106312c68ae7878ba9ba4885ad84fb8f8e153350600d368cedee43b5ce1903"),
		},
		{
			common.HexToHash("0x405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace"),
			common.HexToHash("0x6e106312c68ae7878ba9ba4885ad84fb8f8e153350600d368cedee43b5ce1903"),
		},
		{
			common.HexToHash("0xc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b"),
			common.HexToHash("0x036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0"),
			common.HexToHash("0xa95fdf6e1d2ef63e224adadbf70e466c4ac233aaa7fda9afb4862e884c8185f2"),
		},
		{
			common.HexToHash("0xb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf6"),
			common.HexToHash("0x036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db0"),
			common.HexToHash("0xa95fdf6e1d2ef63e224adadbf70e466c4ac233aaa7fda9afb4862e884c8185f2"),
		},
	}

	for i, hash := range sortedHashes {
		t.Run(fmt.Sprintf("Proof_%d", i), func(t *testing.T) {
			proof, err := tree.Proof(i)
			require.NoError(t, err)
			assert.Equal(t, expectedProofs[i], proof)

			verified := merkle.VerifyProof(hash, proof, root)
			assert.True(t, verified)
		})

		t.Run(fmt.Sprintf("ProofFromHash_%d", i), func(t *testing.T) {
			proof, err := tree.ProofFromHash(hash)
			require.NoError(t, err)
			assert.Equal(t, expectedProofs[i], proof)

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
