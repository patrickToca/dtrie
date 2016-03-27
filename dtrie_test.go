package dtrie

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPopCount(t *testing.T) {
	b := []uint32{
		uint32(0x55555555), // 0x55555555 = 01010101 01010101 01010101 01010101
		uint32(0x33333333), // 0x33333333 = 00110011 00110011 00110011 00110011
		uint32(0x0F0F0F0F), // 0x0F0F0F0F = 00001111 00001111 00001111 00001111
		uint32(0x00FF00FF), // 0x00FF00FF = 00000000 11111111 00000000 11111111
		uint32(0x0000FFFF), // 0x0000FFFF = 00000000 00000000 11111111 11111111
	}
	for _, x := range b {
		assert.Equal(t, 16, popCount(x))
	}
}

func TestDefaultHash(t *testing.T) {
	assert.Equal(t,
		hash(map[int]string{11234: "foo"}),
		hash(map[int]string{11234: "foo"}))
	assert.NotEqual(t, hash("foo"), hash("bar"))
}

type testEntry struct {
	hash  uint32
	key   int
	value int
}

func (e *testEntry) KeyHash() uint32 {
	return e.hash
}

func (e *testEntry) Key() interface{} {
	return e.key
}

func (e *testEntry) Value() interface{} {
	return e.value
}

func (e *testEntry) String() string {
	return fmt.Sprint(e.value)
}

func TestInsert(t *testing.T) {
	n := emptyNode(0, 32)
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		n = insert(n, &testEntry{hash(i), i, i})
	}
	t.Logf("10M insertions: %v\n", time.Since(start))
}

func BenchmarkInsert(b *testing.B) {
	n := emptyNode(0, 32)
	for i := b.N; i > 0; i-- {
		n = insert(n, &testEntry{hash(i), i, i})
	}
	b.ReportAllocs()
}

func TestGet(t *testing.T) {
	n := emptyNode(0, 32)
	for i := 0; i < 10000000; i++ {
		n = insert(n, &testEntry{hash(i), i, i})
	}
	getBenchNode = n
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		x := get(n, hash(i), i)
		if x == nil {
			t.Fatalf("%v not found", i)
		}
	}
	t.Logf("10M gets:\t  %v\n", time.Since(start))
}

var getBenchNode *node

func BenchmarkGet(b *testing.B) {
	for i := b.N; i > 0; i-- {
		get(getBenchNode, hash(i), i)
	}
	b.ReportAllocs()
}

func TestRemove(t *testing.T) {
	n := emptyNode(0, 32)
	for i := 0; i < 10000000; i++ {
		n = insert(n, &testEntry{hash(i), i, i})
	}
	deleteBenchNode = n
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		n = remove(n, hash(i), i)
	}
	for _, e := range n.entries {
		if e != nil {
			t.Fatal("final node is not empty")
		}
	}
	t.Logf("10M deletions: %v\n", time.Since(start))
}

var deleteBenchNode *node

func BenchmarkRemove(b *testing.B) {
	for i := b.N; i > 0; i-- {
		deleteBenchNode = remove(deleteBenchNode, hash(i), i)
	}
	b.ReportAllocs()
}

func TestUpdate(t *testing.T) {
	n := emptyNode(0, 32)
	for i := 0; i < 10000000; i++ {
		n = insert(n, &testEntry{hash(i), i, i})
	}
	updateBenchNode = n
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		n = insert(n, &testEntry{hash(i), i, -i})
	}
	t.Logf("10M updates:   %v\n", time.Since(start))
}

var updateBenchNode *node

func BenchmarkUpdate(b *testing.B) {
	for i := b.N; i > 0; i-- {
		updateBenchNode = insert(updateBenchNode, &testEntry{hash(i), i, -i})
	}
}

func TestIterate(t *testing.T) {
	n := emptyNode(0, 32)
	for i := 0; i < 10000000; i++ {
		n = insert(n, &testEntry{hash(i), i, i})
	}
	stop := make(chan struct{})
	echan := iterate(n, stop)
	c := 0
	start := time.Now()
	for range echan {
		c++
	}
	assert.Equal(t, 10000000, c)
	t.Logf("10M iterations: %v\n", time.Since(start))
}
