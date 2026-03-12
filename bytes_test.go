package feistel_test

import (
	"bytes"
	"errors"
	"hash/fnv"
	"slices"
	"testing"

	"github.com/s0rg/feistel"
)

func Test_BytesHash(t *testing.T) {
	t.Parallel()

	const nRounds = 8

	src := []byte("12345-test-string-to-hash!-67890")

	res, _ := feistel.BytesHash(fnv.New32a(), nRounds, src)
	res, _ = feistel.BytesHash(fnv.New32a(), nRounds, res)

	if !bytes.Equal(res, src) {
		t.Logf("want: %v got: %v", src, res)
		t.Fail()
	}
}

func Test_BytesHash_Errors(t *testing.T) {
	t.Parallel()

	_, err := feistel.BytesHash(fnv.New32a(), 0, []byte{})
	if !errors.Is(err, feistel.ErrBadRoundsCount) {
		t.FailNow()
	}

	_, err = feistel.BytesHash(fnv.New32a(), 1, []byte{})
	if !errors.Is(err, feistel.ErrUnevenDataSize) {
		t.FailNow()
	}
}

func Test_BytesHashKeys(t *testing.T) {
	t.Parallel()

	keys := [][]byte{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{3, 3, 3, 3},
	}

	src := []byte("12345-test-string-to-hash!-67890")

	res, _ := feistel.BytesHashKeys(fnv.New32a(), keys, src)
	slices.Reverse(keys)
	res, _ = feistel.BytesHashKeys(fnv.New32a(), keys, res)

	if !bytes.Equal(res, src) {
		t.Logf("want: %v got: %v", src, res)
		t.Fail()
	}
}

func Test_BytesHashKeys_Errors(t *testing.T) {
	t.Parallel()

	keys := [][]byte{
		{1, 1, 1, 1},
	}

	_, err := feistel.BytesHashKeys(fnv.New32a(), [][]byte{}, []byte{})
	if !errors.Is(err, feistel.ErrBadRoundsCount) {
		t.FailNow()
	}

	_, err = feistel.BytesHashKeys(fnv.New32a(), keys, []byte{})
	if !errors.Is(err, feistel.ErrUnevenDataSize) {
		t.FailNow()
	}
}

func Benchmark_BytesHash(b *testing.B) {
	var (
		src  = []byte("12345-test-string-to-hash!-67890")
		hash = fnv.New32a()
		keys = [][]byte{
			{1, 1, 1, 1},
			{2, 2, 2, 2},
			{3, 3, 3, 3},
			{4, 4, 4, 4},
		}
	)

	b.ResetTimer()

	b.Run("BytesHash-1", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.BytesHash(hash, 1, src)
		}
	})

	b.Run("BytesHash-4", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.BytesHash(hash, 4, src)
		}
	})

	b.Run("BytesHash-8", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.BytesHash(hash, 8, src)
		}
	})

	b.Run("BytesHashKeys-1", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.BytesHashKeys(hash, keys[:1], src)
		}
	})

	b.Run("BytesHashKeys-2", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.BytesHashKeys(hash, keys[:2], src)
		}
	})

	b.Run("BytesHashKeys-4", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.BytesHashKeys(hash, keys, src)
		}
	})
}
