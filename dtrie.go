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

// Dtrie is a persistent hash trie that dynamically expands or shrinks
// to provide efficient memory allocation.
type Dtrie struct {
	root   *node
	hasher func(v interface{}) uint32
}

// New creates an empty DTrie with the given hashing function.
// If nil is passed in, the default hashing function will be used.
func New(hasher func(v interface{}) uint32) *Dtrie {
	if hasher == nil {
		hasher = defaultHasher
	}
	return &Dtrie{
		root:   emptyNode(0, 32),
		hasher: hasher,
	}
}

// Size returns the number of entries in the Dtrie.
func (d *Dtrie) Size() (size int) {
	for range iterate(d.root, nil) {
		size++
	}
	return size
}

// Get returns the Entry for the associated key or returns nil if the
// key does not exist.
func (d *Dtrie) Get(key interface{}) Entry {
	return get(d.root, d.hasher(key), key)
}

// Insert adds an entry to the Dtrie, replacing the existing value if
// the key already exists and returns the resulting Dtrie.
func (d *Dtrie) Insert(entry Entry) *Dtrie {
	root := insert(d.root, entry)
	return &Dtrie{root, d.hasher}
}

// Remove deletes the value for the associated key if it exists and returns
// the resulting Dtrie.
func (d *Dtrie) Remove(key interface{}) *Dtrie {
	root := remove(d.root, d.hasher(key), key)
	return &Dtrie{root, d.hasher}
}

// Iterator returns a read-only channel of Entries from the Dtrie. If a stop
// channel is provided, closing it will terminate and close the iterator
// channel. Note that if a cancel channel is not used and not every entry is
// read from the iterator, a goroutine will leak.
func (d *Dtrie) Iterator(stop <-chan struct{}) <-chan Entry {
	return iterate(d.root, stop)
}
