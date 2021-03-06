package trie_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/ledger"
	"github.com/onflow/flow-go/ledger/common/utils"
	"github.com/onflow/flow-go/ledger/complete/mtrie/trie"
)

const (
	// ReferenceImplPathByteSize is the path length in reference implementation: 2 bytes.
	// Please do NOT CHANGE.
	ReferenceImplPathByteSize = 2
)

// TestEmptyTrie tests whether the root hash of an empty trie matches the formal specification.
// The expected value is coming from a reference implementation in python and is hard-coded here.
func Test_EmptyTrie(t *testing.T) {
	// Make new Trie (independently of MForest):
	emptyTrie, err := trie.NewEmptyMTrie(ReferenceImplPathByteSize)
	require.NoError(t, err)

	expectedRootHashHex := "6e24e2397f130d9d17bef32b19a77b8f5bcf03fb7e9e75fd89b8a455675d574a"
	require.Equal(t, expectedRootHashHex, hex.EncodeToString(emptyTrie.RootHash()))
}

// Test_TrieWithLeftRegister tests whether the root hash of trie with only the left-most
// register populated matches the formal specification.
// The expected value is coming from a reference implementation in python and is hard-coded here.
func Test_TrieWithLeftRegister(t *testing.T) {
	// Make new Trie (independently of MForest):
	emptyTrie, err := trie.NewEmptyMTrie(ReferenceImplPathByteSize)
	require.NoError(t, err)

	path := utils.TwoBytesPath(0)
	payload := utils.LightPayload(11, 12345)
	leftPopulatedTrie, err := trie.NewTrieWithUpdatedRegisters(emptyTrie, []ledger.Path{path}, []ledger.Payload{*payload})
	require.NoError(t, err)
	expectedRootHashHex := "ff472d38a97b3b1786c4dfffa0005370aa3c16805d342ed7618876df7101f760"
	require.Equal(t, expectedRootHashHex, hex.EncodeToString(leftPopulatedTrie.RootHash()))
}

// Test_TrieWithRightRegister tests whether the root hash of trie with only the right-most
// register populated matches the formal specification.
// The expected value is coming from a reference implementation in python and is hard-coded here.
func Test_TrieWithRightRegister(t *testing.T) {
	// Make new Trie (independently of MForest):
	emptyTrie, err := trie.NewEmptyMTrie(ReferenceImplPathByteSize)
	require.NoError(t, err)

	path := utils.TwoBytesPath(65535)
	payload := utils.LightPayload(12346, 54321)
	rightPopulatedTrie, err := trie.NewTrieWithUpdatedRegisters(emptyTrie, []ledger.Path{path}, []ledger.Payload{*payload})
	require.NoError(t, err)
	expectedRootHashHex := "d1fb1c7c84bcd02205fbc7bdf73ee8e943b8bb4b7db6bcc26ae7af67e507fb8d"
	require.Equal(t, expectedRootHashHex, hex.EncodeToString(rightPopulatedTrie.RootHash()))
}

// // Test_TrieWithMiddleRegister tests the root hash of trie holding only a single
// // allocated register somewhere in the middle.
// // The expected value is coming from a reference implementation in python and is hard-coded here.
func Test_TrieWithMiddleRegister(t *testing.T) {
	// Make new Trie (independently of MForest):
	emptyTrie, err := trie.NewEmptyMTrie(ReferenceImplPathByteSize)
	require.NoError(t, err)

	path := utils.TwoBytesPath(56809)
	payload := utils.LightPayload(12346, 59656)
	leftPopulatedTrie, err := trie.NewTrieWithUpdatedRegisters(emptyTrie, []ledger.Path{path}, []ledger.Payload{*payload})
	require.NoError(t, err)
	expectedRootHashHex := "b44a9a00c182ba2203fca6886c4c99b854f9f8279a9978b180ad10e82362e412"
	require.Equal(t, expectedRootHashHex, hex.EncodeToString(leftPopulatedTrie.RootHash()))
}

// Test_TrieWithManyRegisters tests whether the root hash of a trie storing 12001 randomly selected registers
// matches the formal specification.
// The expected value is coming from a reference implementation in python and is hard-coded here.
func Test_TrieWithManyRegisters(t *testing.T) {
	// Make new Trie (independently of MForest):
	emptyTrie, err := trie.NewEmptyMTrie(ReferenceImplPathByteSize)
	require.NoError(t, err)

	// allocate single random register
	rng := &LinearCongruentialGenerator{seed: 0}
	paths, payloads := deduplicateWrites(sampleRandomRegisterWrites(rng, 12001))
	updatedTrie, err := trie.NewTrieWithUpdatedRegisters(emptyTrie, paths, payloads)
	require.NoError(t, err)
	expectedRootHashHex := "18a7c33a0ecf148274f860246f23dffdc6d15dc846e0ae34f6887a43ec67124c"
	require.Equal(t, expectedRootHashHex, hex.EncodeToString(updatedTrie.RootHash()))
}

// Test_FullTrie tests whether the root hash of a fully-populated trie storing
// matches the formal specification.
// The expected value is coming from a reference implementation in python and is hard-coded here.
func Test_FullTrie(t *testing.T) {
	// Make new Trie (independently of MForest):
	emptyTrie, err := trie.NewEmptyMTrie(ReferenceImplPathByteSize)
	require.NoError(t, err)

	// allocate single random register
	capacity := 65536
	rng := &LinearCongruentialGenerator{seed: 0}
	paths := make([]ledger.Path, 0, capacity)
	payloads := make([]ledger.Payload, 0, capacity)
	for i := 0; i < capacity; i++ {
		paths = append(paths, utils.TwoBytesPath(uint16(i)))
		temp := rng.next()
		payload := utils.LightPayload(temp, temp)
		payloads = append(payloads, *payload)
	}
	updatedTrie, err := trie.NewTrieWithUpdatedRegisters(emptyTrie, paths, payloads)
	require.NoError(t, err)
	expectedRootHashHex := "0a1e74e7a4dfcc916dcafbd3f1c826280f047cd5608295f01a32c9af5949898f"
	require.Equal(t, expectedRootHashHex, hex.EncodeToString(updatedTrie.RootHash()))
}

// TestUpdateTrie tests whether iteratively updating a Trie matches the formal specification.
// The expected root hashes are coming from a reference implementation in python and is hard-coded here.
func Test_UpdateTrie(t *testing.T) {
	expectedRootHashes := []string{
		"a8dc0574fdeeaab4b5d3b2a798c19bee5746337a9aea735ebc4dfd97311503c5",
		"6fb27c151f44ba50128c2a6b5ecec19343edf7b68b7b733b64cb5df3c0de4a8b",
		"1c3fccdf4a7e4234b9fb9c576e2a919bca259600056c4f14317bde7f22ad2c5d",
		"5ea61ef89f333a8695057ef3d650745b61c5ffeabc9663c5f1c288b755ff43da",
		"42bcd108195c12eb0122fac0389128a5b5073c1ab8717c225e1f6a9c8b8bc7b6",
		"194d139211362feb28ad1bd56f4c030228748c0045ad6d47665d450b66fb3da2",
		"f5f5cef0b91fdf0cfb10d535b122df7e4b5cb6df47fcf69f3cde80ef2dd23674",
		"28d7a59926dcd6025c744660b95cef52955e6413727a628b314da0e5b4c02ba6",
		"24869a02eecb3f56c37979eee9868170ed78d571f896245ca308dcb59eb8f09d",
		"99f3bbd9fbf19c3a3560c62d845ee6e4f8abc086dd891429c7f297470783a50e",
		"fa53339233bce843b6938f22556bdc9395a401dbabc163185386f750810ea993",
		"93828998941ce554a5c2e780d9951d179e83b28df1ef9c0d6479c176fe3b4a7f",
		"7dde1add8114622f8f01714c1dafae50718e20aad673a043b04b37f5e3ff57a0",
		"aba0dfcd53f8768a9b146b2f50b6ce43d47c45e1b961d9f64a65b2492906543b",
		"950a669dfc88bb8fff0497f677a095da75b506c5f759ebdd31ea0f7536eb81e7",
		"18a7c33a0ecf148274f860246f23dffdc6d15dc846e0ae34f6887a43ec67124c",
		"9574e25612daebf7dcd3e61c707a3fc6a2f23976776befc7671c17b3820db89b",
		"a490e00118ded37c89c358372c118b3b197a7693a294be438bb6557b65fb2265",
		"0f158d9b863a903f59b3e7b7fb35caf595789912b7dae41cb74f986d7b6f247f",
		"a5730e2e89daa48e01802bc83eb14c6ea52f5f38760ad2e844f8f038cbe87c8a",
	}

	// Make new Trie (independently of MForest):
	emptyTrie, err := trie.NewEmptyMTrie(ReferenceImplPathByteSize)
	require.NoError(t, err)

	// allocate single random register
	rng := &LinearCongruentialGenerator{seed: 0}
	path := utils.TwoBytesPath(rng.next())
	temp := rng.next()
	payload := utils.LightPayload(temp, temp)
	updatedTrie, err := trie.NewTrieWithUpdatedRegisters(emptyTrie, []ledger.Path{path}, []ledger.Payload{*payload})
	require.NoError(t, err)
	expectedRootHashHex := "a8dc0574fdeeaab4b5d3b2a798c19bee5746337a9aea735ebc4dfd97311503c5"
	require.Equal(t, expectedRootHashHex, hex.EncodeToString(updatedTrie.RootHash()))

	for r := 0; r < 20; r++ {
		paths, payloads := deduplicateWrites(sampleRandomRegisterWrites(rng, r*100))
		updatedTrie, err = trie.NewTrieWithUpdatedRegisters(updatedTrie, paths, payloads)
		require.NoError(t, err)
		require.Equal(t, expectedRootHashes[r], hex.EncodeToString(updatedTrie.RootHash()))
	}
}

// Test_UnallocateRegisters tests whether unallocating registers matches the formal specification.
// Unallocating here means, to set the stored register value to an empty byte slice
// The expected value is coming from a reference implementation in python and is hard-coded here.
func Test_UnallocateRegisters(t *testing.T) {
	rng := &LinearCongruentialGenerator{seed: 0}
	emptyTrie, err := trie.NewEmptyMTrie(ReferenceImplPathByteSize)
	require.NoError(t, err)

	// we first draw 99 random key-value pairs that will be first allocated and later unallocated:
	paths1, payloads1 := deduplicateWrites(sampleRandomRegisterWrites(rng, 99))
	updatedTrie, err := trie.NewTrieWithUpdatedRegisters(emptyTrie, paths1, payloads1)
	require.NoError(t, err)

	// we then write an additional 117 registers
	paths2, payloads2 := deduplicateWrites(sampleRandomRegisterWrites(rng, 117))
	updatedTrie, err = trie.NewTrieWithUpdatedRegisters(updatedTrie, paths2, payloads2)
	require.NoError(t, err)

	// and now we override the first 99 registers with default values, i.e. unallocate them
	payloads0 := make([]ledger.Payload, len(payloads1))
	updatedTrie, err = trie.NewTrieWithUpdatedRegisters(updatedTrie, paths1, payloads0)
	require.NoError(t, err)

	// this should be identical to the first 99 registers never been written
	expectedRootHashHex := "ce4883f826deaec46317901b7a274a2f9706bc1d1b2cf6869ca1447afb23b2d5"
	comparisionTrie, err := trie.NewTrieWithUpdatedRegisters(emptyTrie, paths2, payloads2)
	require.NoError(t, err)
	require.Equal(t, expectedRootHashHex, hex.EncodeToString(comparisionTrie.RootHash()))
	require.Equal(t, expectedRootHashHex, hex.EncodeToString(updatedTrie.RootHash()))
}

type LinearCongruentialGenerator struct {
	seed uint64
}

func (rng *LinearCongruentialGenerator) next() uint16 {
	rng.seed = (rng.seed*1140671485 + 12820163) % 65536
	return uint16(rng.seed)
}

// sampleRandomRegisterWrites generates path-payload tuples for `number` randomly selected registers;
// caution: registers might repeat
func sampleRandomRegisterWrites(rng *LinearCongruentialGenerator, number int) ([]ledger.Path, []ledger.Payload) {

	paths := make([]ledger.Path, 0, number)
	payloads := make([]ledger.Payload, 0, number)
	for i := 0; i < number; i++ {
		path := utils.TwoBytesPath(rng.next())
		paths = append(paths, path)
		t := rng.next()
		payload := utils.LightPayload(t, t)
		payloads = append(payloads, *payload)
	}
	return paths, payloads
}

// deduplicateWrites retains only the last register write
func deduplicateWrites(paths []ledger.Path, payloads []ledger.Payload) ([]ledger.Path, []ledger.Payload) {
	payloadMapping := make(map[string]int)
	if len(paths) != len(payloads) {
		panic("size mismatch (paths and payloads)")
	}
	for i, path := range paths {
		// we override the latest in the slice
		payloadMapping[string(path)] = i
	}
	dedupedPaths := make([]ledger.Path, 0, len(payloadMapping))
	dedupedPayloads := make([]ledger.Payload, 0, len(payloadMapping))
	for path := range payloadMapping {
		dedupedPaths = append(dedupedPaths, []byte(path))
		dedupedPayloads = append(dedupedPayloads, payloads[payloadMapping[path]])
	}
	return dedupedPaths, dedupedPayloads
}
