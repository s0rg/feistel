package feistel

import (
	"fmt"
	"hash"
)

const (
	minCount = 1
	half     = 2
)

// HashBlock runs balanced Feistel network for specified bytes block its size must be even and twice of hasher size.
// Please note - this method damages provided data slice, use this if speed matters but original data content is not.
func HashBlock(hasher hash.Hash, rounds int, data []byte) (rv []byte, err error) {
	if rounds < minCount {
		return nil, fmt.Errorf("%w: got %d", ErrBadRoundsCount, rounds)
	}

	if len(data) == 0 || len(data)%half == 1 {
		return nil, fmt.Errorf("%w: got %d bytes", ErrUnevenDataSize, len(data))
	}

	l := len(data) / half

	if l != hasher.Size() {
		return nil, fmt.Errorf("%w: need: %d got: %d", ErrWrongHasherSize, l, hasher.Size())
	}

	return HashBlockUnsafe(hasher, rounds, data), nil
}

// HashKeys runs balanced Feistel network for specified bytes block with given set of keys. This method damages
// provided data slice.
func HashKeys(hasher hash.Hash, keys [][]byte, data []byte) (rv []byte, err error) {
	if len(keys) < minCount {
		return nil, fmt.Errorf("%w: keys got only %d items", ErrBadRoundsCount, len(keys))
	}

	if len(data) == 0 || len(data)%half == 1 {
		return nil, fmt.Errorf("%w: got %d bytes", ErrUnevenDataSize, len(data))
	}

	l := len(data) / half

	if l != hasher.Size() {
		return nil, fmt.Errorf("%w: need: %d got: %d", ErrWrongHasherSize, l, hasher.Size())
	}

	return HashKeysUnsafe(hasher, keys, data), nil
}

// HashBlockUnsafe runs balanced Feistel network for specified bytes slice without any bounds checks.
func HashBlockUnsafe(hasher hash.Hash, rounds int, data []byte) (rv []byte) {
	l := len(data) / half
	a, b := data[:l], data[l:]

	for range rounds {
		hasher.Write(b)
		a, b = b, xor(a, hasher.Sum(nil))
		hasher.Reset()
	}

	return append(b, a...)
}

// HashKeysUnsafe runs balanced Feistel network for specified bytes slice with given set of keys
// without any bounds checks.
func HashKeysUnsafe(hasher hash.Hash, keys [][]byte, data []byte) (rv []byte) {
	l := len(data) / half
	a, b := data[:l], data[l:]

	for _, k := range keys {
		hasher.Write(k)
		hasher.Write(b)
		a, b = b, xor(a, hasher.Sum(nil))
		hasher.Reset()
	}

	return append(b, a...)
}

func xor(s, k []byte) (rv []byte) {
	for i := range s {
		s[i] ^= k[i]
	}

	return s
}
