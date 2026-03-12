package feistel_test

import (
	"bytes"
	"errors"
	"hash/fnv"
	"slices"
	"testing"

	"github.com/s0rg/feistel"
)

func Test_HashBlock_8(t *testing.T) {
	t.Parallel()

	src := []byte{0, 1, 2, 3, 4, 5, 6, 7}
	buf := make([]byte, 8)

	copy(buf, src)

	const nRounds = 8

	res, _ := feistel.HashBlock(fnv.New32a(), nRounds, buf)
	res, _ = feistel.HashBlock(fnv.New32a(), nRounds, res)

	if !bytes.Equal(res, src) {
		t.Logf("want: %v got: %v", src, res)
		t.Fail()
	}
}

func Test_HashBlock_16(t *testing.T) {
	t.Parallel()

	src := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	buf := make([]byte, 16)

	copy(buf, src)

	const nRounds = 8

	res, _ := feistel.HashBlock(fnv.New64a(), nRounds, buf)
	res, _ = feistel.HashBlock(fnv.New64a(), nRounds, res)

	if !bytes.Equal(res, src) {
		t.Logf("want: %v got: %v", src, res)
		t.Fail()
	}
}

func Test_HashBlock_String(t *testing.T) {
	t.Parallel()

	src := []byte("my-test-data-str")
	buf := make([]byte, len(src))

	copy(buf, src)

	const nRounds = 8

	res, _ := feistel.HashBlock(fnv.New64a(), nRounds, buf)
	res, _ = feistel.HashBlock(fnv.New64a(), nRounds, res)

	if !bytes.Equal(res, src) {
		t.Logf("want: %v got: %v", src, res)
		t.Fail()
	}
}

func Test_HashKeys(t *testing.T) {
	t.Parallel()

	src := []byte{0, 1, 2, 3, 4, 5, 6, 7}
	keys := [][]byte{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{3, 3, 3, 3},
	}

	buf := make([]byte, 8)

	copy(buf, src)

	res, _ := feistel.HashKeys(fnv.New32a(), keys, buf)
	slices.Reverse(keys)
	res, _ = feistel.HashKeys(fnv.New32a(), keys, res)

	if !bytes.Equal(res, src) {
		t.Logf("want: %v got: %v", src, res)
		t.Fail()
	}
}

func Test_HashKeys_Errors(t *testing.T) {
	t.Parallel()

	_, err := feistel.HashKeys(fnv.New32a(), [][]byte{}, []byte{})
	if !errors.Is(err, feistel.ErrBadRoundsCount) {
		t.FailNow()
	}

	_, err = feistel.HashKeys(fnv.New32a(), [][]byte{{1}}, []byte{1})
	if !errors.Is(err, feistel.ErrUnevenDataSize) {
		t.FailNow()
	}

	_, err = feistel.HashKeys(fnv.New32a(), [][]byte{{1}}, []byte{1, 2})
	if !errors.Is(err, feistel.ErrWrongHasherSize) {
		t.FailNow()
	}
}

func Test_HashBlock_Errors(t *testing.T) {
	t.Parallel()

	_, err := feistel.HashBlock(fnv.New32a(), 0, []byte{})
	if !errors.Is(err, feistel.ErrBadRoundsCount) {
		t.FailNow()
	}

	const nRounds = 8

	_, err = feistel.HashBlock(fnv.New32a(), nRounds, []byte{1})
	if !errors.Is(err, feistel.ErrUnevenDataSize) {
		t.FailNow()
	}

	_, err = feistel.HashBlock(fnv.New32a(), nRounds, []byte{1, 2})
	if !errors.Is(err, feistel.ErrWrongHasherSize) {
		t.FailNow()
	}
}

func Benchmark_HashBlock(b *testing.B) {
	var (
		buf  = []byte{1, 2, 3, 4, 5, 6, 7, 8}
		hash = fnv.New32a()
		keys = [][]byte{
			{1, 1, 1, 1},
			{2, 2, 2, 2},
			{3, 3, 3, 3},
			{4, 4, 4, 4},
		}
	)

	b.ResetTimer()

	b.Run("HashBlock-1", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.HashBlock(hash, 1, buf)
		}
	})

	b.Run("HashBlock-4", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.HashBlock(hash, 4, buf)
		}
	})

	b.Run("HashBlock-8", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.HashBlock(hash, 8, buf)
		}
	})

	b.Run("HashBlockSafe-1", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.HashBlockSafe(hash, 1, buf)
		}
	})

	b.Run("HashBlockSafe-4", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.HashBlockSafe(hash, 4, buf)
		}
	})

	b.Run("HashBlockSafe-8", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.HashBlockSafe(hash, 8, buf)
		}
	})

	b.Run("HashKeys-1", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.HashKeys(hash, keys[:1], buf)
		}
	})

	b.Run("HashKeys-2", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.HashKeys(hash, keys[:2], buf)
		}
	})

	b.Run("HashKeys-4", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.HashKeys(hash, keys, buf)
		}
	})
}
