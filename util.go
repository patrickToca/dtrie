// Copyright (c) 2016, Theodore Butler
// Use of this source code is governed by a BSD 2-Caluse
// license that can be found in the LICENSE file.

// Package dtrie provides an implementation of the dtrie data structure, which
// is a persistent hash trie that dynamically expands or shrinks to provide
// efficient memory allocation. This data structure is based on the papers
// Ideal Hash Trees by Phil Bagwell and Optimizing Hash-Array Mapped Tries for
// Fast and Lean Immutable JVM Collections by Michael J. Steindorfer and
// Jurgen J. Vinju
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
	case int:
		return uint32(value.(int))
	}
	hasher := fnv.New32a()
	hasher.Write([]byte(fmt.Sprintf("%#v", value)))
	return hasher.Sum32()
}
