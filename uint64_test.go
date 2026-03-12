package feistel_test

import (
	"errors"
	"hash/fnv"
	"slices"
	"testing"

	"github.com/s0rg/feistel"
)

func Test_Uint64(t *testing.T) {
	t.Parallel()

	const (
		nRounds = 8
		nValue  = uint64(1276)
	)

	tmp, _ := feistel.Uint64Hash(fnv.New32a(), nRounds, nValue)
	tmp, _ = feistel.Uint64Hash(fnv.New32a(), nRounds, tmp)

	if tmp != nValue {
		t.Logf("want: %d got: %d", nValue, tmp)
		t.Fail()
	}
}

func Test_Uint64_Errors(t *testing.T) {
	t.Parallel()

	_, err := feistel.Uint64Hash(fnv.New128a(), 0, 1)
	if !errors.Is(err, feistel.ErrBadRoundsCount) {
		t.FailNow()
	}

	_, err = feistel.Uint64Hash(fnv.New128a(), 1, 1)
	if !errors.Is(err, feistel.ErrWrongHasherSize) {
		t.FailNow()
	}
}

func Test_Int64(t *testing.T) {
	t.Parallel()

	const (
		nRounds = 8
		nValue  = int64(-1276)
	)

	// helper, simple static cast from negative constant to uint makes govet crazy.
	toUint := func(v int64) (rv uint64) {
		return uint64(v)
	}

	tmp, _ := feistel.Uint64Hash(fnv.New32a(), nRounds, toUint(nValue))
	tmp, _ = feistel.Uint64Hash(fnv.New32a(), nRounds, tmp)

	if int64(tmp) != nValue {
		t.Logf("want: %d got: %d", nValue, tmp)
		t.Fail()
	}
}

func Test_Uint64Keys(t *testing.T) {
	t.Parallel()

	const nValue = uint64(1276)

	keys := [][]byte{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{3, 3, 3, 3},
	}

	tmp, _ := feistel.Uint64HashKeys(fnv.New32a(), keys, nValue)
	slices.Reverse(keys)
	tmp, _ = feistel.Uint64HashKeys(fnv.New32a(), keys, tmp)

	if tmp != nValue {
		t.Logf("want: %d got: %d", nValue, tmp)
		t.Fail()
	}
}

func Test_Uint64Keys_Errors(t *testing.T) {
	t.Parallel()

	const nValue = uint64(1276)

	keys := [][]byte{
		{1, 1, 1, 1},
	}

	_, err := feistel.Uint64HashKeys(fnv.New32a(), [][]byte{}, nValue)
	if !errors.Is(err, feistel.ErrBadRoundsCount) {
		t.FailNow()
	}

	_, err = feistel.Uint64HashKeys(fnv.New128a(), keys, nValue)
	if !errors.Is(err, feistel.ErrWrongHasherSize) {
		t.FailNow()
	}
}

func Benchmark_Uint64(b *testing.B) {
	var hash = fnv.New32a()

	b.ResetTimer()

	b.Run("Uint64Hash-1", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.Uint64Hash(hash, 1, 1)
		}
	})

	b.Run("Uint64Hash-4", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.Uint64Hash(hash, 4, 1)
		}
	})

	b.Run("Uint64Hash-8", func(b *testing.B) {
		for range b.N {
			_, _ = feistel.Uint64Hash(hash, 8, 1)
		}
	})
}
