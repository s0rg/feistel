package feistel

import (
	"encoding/binary"
	"fmt"
	"hash"
)

// Uint64Hash runs Feistel network for specified integer value.
func Uint64Hash(hasher hash.Hash, rounds int, value uint64) (rv uint64, err error) {
	var buf [8]byte

	binary.LittleEndian.PutUint64(buf[:], value)

	res, err := HashBlock(hasher, rounds, buf[:])
	if err != nil {
		return 0, fmt.Errorf("hash block: %w", err)
	}

	return binary.LittleEndian.Uint64(res), nil
}

// Uint64HashKeys runs Feistel network for specified integer value and set of keys.
func Uint64HashKeys(hasher hash.Hash, keys [][]byte, value uint64) (rv uint64, err error) {
	var buf [8]byte

	binary.LittleEndian.PutUint64(buf[:], value)

	res, err := HashKeys(hasher, keys, buf[:])
	if err != nil {
		return 0, fmt.Errorf("hash keys: %w", err)
	}

	return binary.LittleEndian.Uint64(res), nil
}
