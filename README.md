[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/feistel)](https://pkg.go.dev/github.com/s0rg/feistel)
[![License](https://img.shields.io/github/license/s0rg/feistel)](https://github.com/s0rg/feistel/blob/main/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/feistel)](go.mod)
[![Tag](https://img.shields.io/github/v/tag/s0rg/feistel?sort=semver)](https://github.com/s0rg/feistel/tags)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/feistel)](https://goreportcard.com/report/github.com/s0rg/feistel)
![Issues](https://img.shields.io/github/issues/s0rg/feistel)

# feistel

This package implements balanced variant of [Feistel network (also known as Feistel cipher)](https://en.wikipedia.org/wiki/Feistel_cipher) for golang.

TLDR: It takes bytes and hashes them to reversable form

# features

- small api based on standart data types and interfaces
- additional `Block*` functions to use it as part of hashing or crypto lib
- 100% test coverage

# caveats

- all incoming data slices sizes **must** be rounded up to selected hash size

# example

```go
package main

import (
	"fmt"
	"hash"
	"hash/fnv"

	"github.com/google/uuid"
	"github.com/s0rg/feistel"
)

const nRounds = 4

func hashUint64(h hash.Hash, rounds int, value uint64) {
	a, _ := feistel.Uint64Hash(h, rounds, value)
	fmt.Printf("\t%d -> %d\n", value, a)

	h.Reset()

	b, _ := feistel.Uint64Hash(h, rounds, a)
	fmt.Printf("\t%d -> %d\n", a, b)
}

func hashInt64(h hash.Hash, rounds int, value int64) {
	a, _ := feistel.Uint64Hash(h, rounds, uint64(value))
	fmt.Printf("\t%d -> %d\n", value, a)

	h.Reset()

	b, _ := feistel.Uint64Hash(h, rounds, a)
	fmt.Printf("\t%d -> %d\n", a, int64(b))
}

func hashBytes(h hash.Hash, rounds int, value []byte) (rv []byte) {
	a, _ := feistel.BytesHash(h, rounds, value)
	fmt.Printf("\t%v -> %v\n", value, a)

	h.Reset()

	b, _ := feistel.BytesHash(h, rounds, a)
	fmt.Printf("\t%v -> %v\n", a, b)

	return b
}

func main() {
	hash := fnv.New32a()

	fmt.Println("uin64:")
	hashUint64(hash, nRounds, 123456)
	hash.Reset()

	fmt.Println("int64:")
	hashInt64(hash, nRounds, -123456)
	hash.Reset()

	fmt.Println("bytes:")
	id, _ := uuid.NewV7()
	res := hashBytes(hash, nRounds, id[:])
	fmt.Printf("\toriginal : %s\n", id.String())
	fmt.Printf("\tresult   : %s\n", uuid.UUID(res).String())
}
```

Output:

```
uin64:
	123456 -> 11152740970488264477
	11152740970488264477 -> 123456
int64:
	-123456 -> 6380474240854226338
	6380474240854226338 -> -123456
bytes:
	[1 156 228 69 159 236 118 49 137 105 96 245 156 183 76 171] -> [217 207 28 95 81 184 4 143 250 166 52 24 222 245 102 47]
	[217 207 28 95 81 184 4 143 250 166 52 24 222 245 102 47] -> [1 156 228 69 159 236 118 49 137 105 96 245 156 183 76 171]
	original : 019ce445-9fec-7631-8969-60f59cb74cab
	result   : 019ce445-9fec-7631-8969-60f59cb74cab
```

# license

MIT
