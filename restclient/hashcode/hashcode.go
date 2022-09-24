package hashcode

import (
	"hash/crc32"
)

type Comparer interface {
	GetHashCode() int
}

func String(value string) int {
	checksum := int(crc32.ChecksumIEEE([]byte(value)))
	if checksum >= 0 {
		return checksum
	}
	if -checksum >= 0 {
		return -checksum
	}
	return 0
}
