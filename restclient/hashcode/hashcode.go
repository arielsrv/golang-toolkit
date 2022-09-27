package hashcode

import (
	"hash/fnv"
)

type Comparer interface {
	GetHashCode() uint64
}

func GetValue(value string) uint64 {
	hash := fnv.New64()
	hash.Write([]byte(value))
	return hash.Sum64()
}
