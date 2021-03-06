// Copyright (c) 2016, Theodore Butler
// Use of this source code is governed by a BSD 2-Caluse
// license that can be found in the LICENSE file.

package dtrie

import (
	"fmt"
	"hash/fnv"
)

func mask(hash, level uint32) uint32 {
	return (hash >> (5 * level)) & 0x01f
}

func setBit(bitmap uint32, pos uint32) uint32 {
	return bitmap | (1 << pos)
}

func clearBit(bitmap uint32, pos uint32) uint32 {
	return bitmap & ^(1 << pos)
}

func hasBit(bitmap uint32, pos uint32) bool {
	return (bitmap & (1 << pos)) != 0
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

func defaultHasher(value interface{}) uint32 {
	switch value.(type) {
	case uint8:
		return uint32(value.(uint8))
	case uint16:
		return uint32(value.(uint16))
	case uint32:
		return value.(uint32)
	case uint64:
		return uint32(value.(uint64))
	case int8:
		return uint32(value.(int8))
	case int16:
		return uint32(value.(int16))
	case int32:
		return uint32(value.(int32))
	case int64:
		return uint32(value.(int64))
	case uint:
		return uint32(value.(uint))
	case int:
		return uint32(value.(int))
	case uintptr:
		return uint32(value.(uintptr))
	case float32:
		return uint32(value.(float32))
	case float64:
		return uint32(value.(float64))
	}
	hasher := fnv.New32a()
	hasher.Write([]byte(fmt.Sprintf("%#v", value)))
	return hasher.Sum32()
}
