package hashcode

import (
	"hash/fnv"
)

type Comparer interface {
	GetHashCode() int
}

func String(value string) uint64 {
	hash := fnv.New64()
	hash.Write([]byte(value))
	return hash.Sum64()
}
