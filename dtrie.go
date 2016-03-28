package dtrie

// Dtrie is a persistent hash trie that dynamically expands or shrinks
// to provide efficient memory allocation
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

// Size returns the number of entries in the Dtrie
func (d *Dtrie) Size() (size int) {
	for range iterate(d.root, nil) {
		size++
	}
	return size
}
