package hash

import "hash/maphash"

var stringHash maphash.Hash

func String(s string) int {
	_, _ = stringHash.WriteString(s)
	defer stringHash.Reset()
	return int(stringHash.Sum64())
}
