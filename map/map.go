package go_redblacktree_container

import rbTree "github.com/OneOfLzx/go-redblacktree"

// Map underlying red-black tree
type Map struct {
	tree rbTree.RedBlackTree // underlying red-black tree
	size uint64              // size of underlying red-black tree node
}

// NewMap construct a Map instance.
// Rules for passing parameters (equalFunc, smallerFunc):
// When the Map save built-in type, caller can ignore this function
// and set this function to nil. Otherwise, mapNodeValueEntryComparator will not know
// how to compare value, so that caller need to implement this function and custom the
// comparison operation. If not nil, using this function first for comparison operation.
//
// type with default EQUAL comparison operation:
// - int8
// - uint8
// - int16
// - uint16
// - int32
// - uint32
// - int64
// - uint64
// - int
// - uint
// - uintptr
// - float32
// - float64
// - complex64
// - complex128
// - string
// type with default SMALLER comparison operation:
// - int8
// - uint8
// - int16
// - uint16
// - int32
// - uint32
// - int64
// - uint64
// - int
// - uint
// - uintptr
// - float32
// - float64
// - string
func NewMap(equalFunc MapExtComparatorFunc, smallerFunc MapExtComparatorFunc) *Map {
	comparator := mapNodeValueEntryComparator{
		extEqualFunc:   equalFunc,
		extSmallerFunc: smallerFunc,
	}
	return &Map{
		tree: rbTree.RedBlackTree{
			Comparator: comparator,
		},
		size: uint64(0),
	}
}

// Find return iterator of the result node. Return nil if error
func (m *Map) Find(key interface{}) *MapIterator {
	if node, ok := m.tree.FindNode(mapNodeValueEntry{key: key, value: nil}); ok {
		return &MapIterator{node: node}
	}
	return nil
}

// Insert inserts a key-value pair entry and return iterator of the result node. Return nil if error.
// If the key is already exist, Insert will modify the value of this key-value pair
func (m *Map) Insert(key interface{}, val interface{}) *MapIterator {
	if node, ok := m.tree.FindNode(mapNodeValueEntry{key: key, value: nil}); ok {
		m.tree.RemoveNode(node)
		m.size--
	}
	if node, ok := m.tree.AddNode(mapNodeValueEntry{key: key, value: val}); ok {
		m.size++
		return &MapIterator{node: node}
	}
	return nil
}

// Remove a  key-value pair entry from Map
func (m *Map) Remove(key interface{}) {
	if _, ok := m.tree.FindNode(mapNodeValueEntry{key: key, value: nil}); !ok {
		return
	}
	m.tree.RemoveNodeByVal(mapNodeValueEntry{key: key, value: nil})
	m.size--
}

// RemoveByIter a  key-value pair entry from Map by iterator
func (m *Map) RemoveByIter(it *MapIterator) {
	m.Remove(it.Key())
}

// Size the number of Map node
func (m *Map) Size() uint64 {
	return m.size
}
