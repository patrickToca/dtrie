package dtrie

import (
	"fmt"
	"hash/fnv"
)

func hash(value interface{}) uint32 {
	switch value.(type) {
	case int:
		return uint32(value.(int))
		//return uint32(0xffffffff) // for testing collisions
	}
	hasher := fnv.New32a()
	hasher.Write([]byte(fmt.Sprintf("%#v", value)))
	return hasher.Sum32()
}

func mask(hash, level uint32) uint32 {
	return (hash >> (5 * level)) & 0x01f
}

func setBit(bitmap uint32, pos uint32) uint32 {
	bitmap |= (1 << pos)
	return bitmap
}

func clearBit(bitmap uint32, pos uint32) uint32 {
	var mask uint32 = ^(1 << pos)
	bitmap &= mask
	return bitmap
}

func hasBit(bitmap uint32, pos uint32) bool {
	val := bitmap & (1 << pos)
	return val > 0
}

func popCount(bitmap uint32) int {
	// bit population count, see
	// http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
	bitmap -= (bitmap >> 1) & 0x55555555
	bitmap = (bitmap>>2)&0x33333333 + bitmap&0x33333333
	bitmap += bitmap >> 4
	bitmap &= 0x0f0f0f0f
	bitmap *= 0x01010101
	return int(byte(bitmap >> 24))
}