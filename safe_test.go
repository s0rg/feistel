package feistel_test

import (
	"bytes"
	"hash/fnv"
	"slices"
	"testing"

	"github.com/s0rg/feistel"
)

func Test_HashBlockSafe_8(t *testing.T) {
	t.Parallel()

	src := []byte{0, 1, 2, 3, 4, 5, 6, 7}

	const nRounds = 8

	res, _ := feistel.HashBlockSafe(fnv.New32a(), nRounds, src)
	res, _ = feistel.HashBlock(fnv.New32a(), nRounds, res)

	if !bytes.Equal(res, src) {
		t.Logf("want: %v got: %v", src, res)
		t.Fail()
	}
}

func Test_HashKeysSafe_8(t *testing.T) {
	t.Parallel()

	src := []byte{0, 1, 2, 3, 4, 5, 6, 7}
	keys := [][]byte{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{3, 3, 3, 3},
	}

	res, _ := feistel.HashKeysSafe(fnv.New32a(), keys, src)
	slices.Reverse(keys)
	res, _ = feistel.HashKeys(fnv.New32a(), keys, res)

	if !bytes.Equal(res, src) {
		t.Logf("want: %v got: %v", src, res)
		t.Fail()
	}
}
