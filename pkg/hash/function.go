package hash

import (
	"fmt"
	"hash/maphash"
)

var stringHash maphash.Hash

func String(s string) int {
	_, _ = stringHash.WriteString(s)
	defer stringHash.Reset()
	return int(stringHash.Sum64())
}

func Universal[T any](a T) int {
	return String(fmt.Sprint(a))
}
