// Package merkle provides a Merkle tree implementation with proof generation and verification.
//
// SECURITY: leaf/internal domain separation is the caller's responsibility.
// Internal nodes are keccak256(min(x,y) || max(x,y)) with no built-in domain
// tag, and proofs bind only the multiset of sibling hashes (not left/right
// position). A 32-byte value that equals any internal-node hash of the tree
// can therefore be presented as a "leaf" and VerifyProof will accept a forged
// inclusion proof — a second-preimage attack on the tree.
//
// To prevent this, callers MUST hash leaves of a fixed, non-collidable
// structure (e.g., keccak256(abi.encode(...)) over a struct that cannot
// collide with the 64-byte two-hash concatenation used at internal nodes).
// All in-tree callers (rewards, FDC, XRP) already do this. Building a tree
// directly from raw 32-byte values is unsafe.
package merkle

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	ErrEmptyTree    = errors.New("empty tree")
	ErrInvalidTree  = errors.New("invalid tree layout")
	ErrInvalidIndex = errors.New("invalid index")
	ErrHashNotFound = errors.New("hash not found")
)

// Tree is Merkle tree implementation with helper functions.
type Tree []common.Hash

// Build builds the Merkle tree from a slice of leaf hashes.
//
// If initialHash is true, each leaf hash is hashed again before building the
// tree. NOTE: this does NOT provide leaf/internal domain separation — both
// the rehashed leaves and the internal nodes are plain keccak256 outputs.
// Caller-side domain separation (see package doc) is still required.
//
// Build sorts the input hashes and silently drops duplicates, so the leaf
// position in the resulting tree has no relation to the caller's input order.
// Use ProofFromHash to look proofs up by leaf value.
func Build(hashes []common.Hash, initialHash bool) Tree {
	if initialHash {
		hashes = mapSingleHash(hashes)
	}

	n := len(hashes)

	if n == 0 {
		return Tree{}
	}

	// Hashes must be sorted to enable binary search.
	sortedHashes := make([]common.Hash, n)
	copy(sortedHashes, hashes)
	sort.Slice(sortedHashes, func(i, j int) bool {
		return bytes.Compare(sortedHashes[i][:], sortedHashes[j][:]) < 0
	})

	sortedHashes = removeDuplicates(sortedHashes)
	n = len(sortedHashes)

	// 2*n-1 overflows int on 32-bit platforms when n > 2^30. The make below
	// would then panic with a misleading "len out of range"; reject early.
	if n > (int(^uint(0)>>1)-1)/2+1 {
		return Tree{}
	}

	tree := make(Tree, 2*n-1)
	copy(tree[n-1:], sortedHashes)

	for i := n - 2; i >= 0; i-- {
		tree[i] = SortedHashPair(tree[2*i+1], tree[2*i+2])
	}

	return tree
}

// removeDuplicates removes duplicated hashes from an ordered slice.
// The slice must be sorted; the function does not work correctly otherwise.
func removeDuplicates(hashes []common.Hash) []common.Hash {
	output := make([]common.Hash, 0, len(hashes))
	for i := range hashes {
		if i == 0 || hashes[i] != hashes[i-1] {
			output = append(output, hashes[i])
		}
	}

	return output
}

// BuildFromHex builds the Merkle tree from a slice of hex-encoded leaf hashes.
// If initialHash is true, each leaf hash is hashed again before building the tree.
//
// Each hex value must decode to exactly 32 bytes (with or without a 0x prefix);
// short / long inputs that common.HexToHash would silently pad or truncate are
// rejected with an error. The same leaf-domain-separation precondition applies;
// see the package doc.
func BuildFromHex(hexValues []string, initialHash bool) (Tree, error) {
	hashes := make([]common.Hash, 0, len(hexValues))
	for i, v := range hexValues {
		s := strings.TrimPrefix(strings.TrimPrefix(v, "0x"), "0X")
		if len(s) != 2*common.HashLength {
			return nil, fmt.Errorf("hex leaf %d: expected %d hex chars, got %d", i, 2*common.HashLength, len(s))
		}
		b, err := hex.DecodeString(s)
		if err != nil {
			return nil, fmt.Errorf("hex leaf %d: %w", i, err)
		}
		hashes = append(hashes, common.BytesToHash(b))
	}

	return Build(hashes, initialHash), nil
}

func mapSingleHash(hashes []common.Hash) []common.Hash {
	output := make([]common.Hash, len(hashes))

	for i := range hashes {
		output[i] = crypto.Keccak256Hash(hashes[i].Bytes())
	}

	return output
}

// SortedHashPair returns keccak256(min(x,y) || max(x,y)). This is the
// internal-node hash; it has no domain tag distinguishing it from a leaf,
// so the caller-side leaf-domain precondition (see package doc) is what
// blocks second-preimage attacks.
func SortedHashPair(x, y common.Hash) common.Hash {
	if bytes.Compare(x[:], y[:]) <= 0 {
		return crypto.Keccak256Hash(x.Bytes(), y.Bytes())
	}

	return crypto.Keccak256Hash(y.Bytes(), x.Bytes())
}

// validate returns ErrInvalidTree if the tree's heap-layout invariant is
// broken (non-empty tree with even length). A valid non-empty tree has
// length 2n-1 (always odd). Callers handle the empty case separately.
func (t Tree) validate() error {
	if len(t) > 0 && len(t)%2 == 0 {
		return ErrInvalidTree
	}
	return nil
}

// Root returns the Merkle root of the tree.
func (t Tree) Root() (common.Hash, error) {
	if len(t) == 0 {
		return common.Hash{}, ErrEmptyTree
	}
	if err := t.validate(); err != nil {
		return common.Hash{}, err
	}

	return t[0], nil
}

// LeavesCount returns the number of leaves in the tree.
func (t Tree) LeavesCount() int {
	if len(t) == 0 {
		return 0
	}

	return (len(t) + 1) / 2
}

// Leaves returns all leaves in a slice. The returned slice is a fresh copy;
// mutating it does not affect the underlying tree.
func (t Tree) Leaves() []common.Hash {
	numLeaves := t.LeavesCount()
	if numLeaves == 0 {
		return nil
	}

	out := make([]common.Hash, numLeaves)
	copy(out, t[numLeaves-1:])
	return out
}

// Leaf returns the i-th leaf.
func (t Tree) Leaf(i int) (common.Hash, error) {
	if err := t.validate(); err != nil {
		return common.Hash{}, err
	}
	numLeaves := t.LeavesCount()
	if numLeaves == 0 || i < 0 || i >= numLeaves {
		return common.Hash{}, ErrInvalidIndex
	}

	pos := len(t) - numLeaves + i
	return t[pos], nil
}

// Proof returns the Merkle proof for the i-th leaf.
func (t Tree) Proof(i int) ([]common.Hash, error) {
	if err := t.validate(); err != nil {
		return nil, err
	}
	numLeaves := t.LeavesCount()
	if numLeaves == 0 || i < 0 || i >= numLeaves {
		return nil, ErrInvalidIndex
	}

	var proof []common.Hash

	for pos := len(t) - numLeaves + i; pos > 0; pos = parent(pos) {
		sibling := pos + ((2 * (pos % 2)) - 1)
		proof = append(proof, t[sibling])
	}

	return proof, nil
}

// parent returns the index of the parent node.
func parent(i int) int {
	return (i - 1) / 2
}

// ProofFromHash returns the proof that hash is stored in a leaf of the tree.
// Returns ErrHashNotFound if hash is not stored in the tree.
func (t Tree) ProofFromHash(hash common.Hash) ([]common.Hash, error) {
	i, err := t.binarySearch(hash)
	if err != nil {
		return nil, err
	}

	return t.Proof(i)
}

func (t Tree) binarySearch(hash common.Hash) (int, error) {
	if err := t.validate(); err != nil {
		return 0, err
	}
	leaves := t.Leaves()
	i := sort.Search(len(leaves), func(i int) bool {
		return bytes.Compare(leaves[i][:], hash[:]) >= 0
	})

	if i < len(leaves) && leaves[i] == hash {
		return i, nil
	}

	return 0, ErrHashNotFound
}

// VerifyProof verifies a Merkle proof for a given leaf against the Merkle root.
//
// leaf must come from the caller's leaf domain (see package doc). VerifyProof
// binds only the multiset of sibling hashes along the path — not the position
// of the leaf — so any 32-byte value that happens to equal an internal-node
// hash will verify with the right proof. The caller's leaf encoding is the
// only thing preventing this.
func VerifyProof(leaf common.Hash, proof []common.Hash, root common.Hash) bool {
	hash := leaf
	for _, pair := range proof {
		hash = SortedHashPair(pair, hash)
	}

	return hash == root
}
