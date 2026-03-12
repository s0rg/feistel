package feistel

import "hash"

// HashBlockSafe is a safe version of BlockHash method, its slower and uses one more alloc but does not change
// content of provided data slice, uses this if you need those bytes elsewhere.
func HashBlockSafe(hasher hash.Hash, rounds int, data []byte) (rv []byte, err error) {
	tmp := make([]byte, len(data))
	copy(tmp, data)

	return HashBlock(hasher, rounds, tmp)
}

// HashKeysSafe is a safe version of BlockKeys method - it does not change incoming data slice.
func HashKeysSafe(hasher hash.Hash, keys [][]byte, data []byte) (rv []byte, err error) {
	tmp := make([]byte, len(data))
	copy(tmp, data)

	return HashKeys(hasher, keys, tmp)
}
