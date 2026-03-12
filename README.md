# feistel

This package implements [Feistel network (also known as Feistel cipher)](https://en.wikipedia.org/wiki/Feistel_cipher) for golang.

# example

```go
package main

import (
	"fmt"
	"hash/fnv"

	"github.com/google/uuid"
	"github.com/s0rg/feistel"
)

const nRounds = 4

func main() {
	hash := fnv.New32a()

	fmt.Println("uin64:")

	a := uint64(12345)
	b, _ := feistel.Uint64Hash(hash, nRounds, a)
	fmt.Printf("\t%d -> %d\n", a, b)

	hash.Reset()
	c, _ := feistel.Uint64Hash(hash, nRounds, b)
	fmt.Printf("\t%d -> %d\n", b, c)

	fmt.Println("bytes:")

	id, _ := uuid.NewV7()
	id1, _ := feistel.BytesHash(hash, nRounds, id[:])
	fmt.Printf("\t%v -> %v\n", id[:], id1)

	hash.Reset()
	id2, _ := feistel.BytesHash(hash, nRounds, id1)
	fmt.Printf("\t%v -> %v\n", id1, id2)
	fmt.Printf("\toriginal : %s\n", id.String())
	fmt.Printf("\tresult   : %s\n", uuid.UUID(id2).String())
}
```

Output:

```
uin64:
	12345 -> 4322171683138304448
	4322171683138304448 -> 12345
bytes:
	[1 156 227 224 159 1 122 207 173 161 6 106 236 78 38 38] -> [230 100 151 53 165 69 5 210 93 19 100 238 94 48 129 250]
	[230 100 151 53 165 69 5 210 93 19 100 238 94 48 129 250] -> [1 156 227 224 159 1 122 207 173 161 6 106 236 78 38 38]
	original : 019ce3e0-9f01-7acf-ada1-066aec4e2626
	result   : 019ce3e0-9f01-7acf-ada1-066aec4e2626
```

# license

MIT
