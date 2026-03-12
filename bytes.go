package feistel

import (
	"fmt"
	"hash"
)

// BytesHash runs balanced Feistel network for specified bytes slice.
func BytesHash(hasher hash.Hash, rounds int, data []byte) (rv []byte, err error) {
	if rounds < minCount {
		return nil, fmt.Errorf("%w: got %d", ErrBadRoundsCount, rounds)
	}

	bsize := hasher.Size() * half
	nblocks := len(data) / bsize

	if nblocks < 1 || len(data) > nblocks*bsize {
		return nil, fmt.Errorf("%w: need at least %d bytes", ErrUnevenDataSize, (nblocks+1)*bsize)
	}

	rv = make([]byte, len(data))
	buf := make([]byte, bsize)

	for i := range nblocks {
		j := i * bsize

		copy(buf, data[j:j+bsize])
		copy(rv[j:], HashBlockUnsafe(hasher, rounds, buf))
	}

	return rv, nil
}

// BytesHashKeys runs balanced Feistel network for specified bytes slice with given set of keys.
func BytesHashKeys(hasher hash.Hash, keys [][]byte, data []byte) (rv []byte, err error) {
	if len(keys) < minCount {
		return nil, fmt.Errorf("%w: got %d keys", ErrBadRoundsCount, len(keys))
	}

	bsize := hasher.Size() * half
	nblocks := len(data) / bsize

	if nblocks < 1 || len(data) > nblocks*bsize {
		return nil, fmt.Errorf("%w: need at least %d bytes", ErrUnevenDataSize, (nblocks+1)*bsize)
	}

	rv = make([]byte, len(data))
	buf := make([]byte, bsize)

	for i := range nblocks {
		j := i * bsize

		copy(buf, data[j:j+bsize])
		copy(rv[j:], HashKeysUnsafe(hasher, keys, buf))
	}

	return rv, nil
}
