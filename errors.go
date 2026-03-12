package feistel

import "errors"

var (
	// ErrBadRoundsCount - invalid rounds value.
	ErrBadRoundsCount = errors.New("rounds value must be > 0")

	// ErrUnevenDataSize - incoming data size in not even.
	ErrUnevenDataSize = errors.New("data size must be even")

	// ErrWrongHasherSize - invalid hasher selected for given data size.
	ErrWrongHasherSize = errors.New("wrong hasher size - must be half of data size")
)
