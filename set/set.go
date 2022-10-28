package go_redblacktree_container

import rbTree "github.com/OneOfLzx/go-redblacktree"

// Set underlying red-black tree
type Set struct {
	tree rbTree.RedBlackTree // underlying red-black tree
	size uint64              // size of underlying red-black tree node
}

// NewSet construct a Set instance.
// Rules for passing parameters (equalFunc, smallerFunc):
// When the Set save built-in type, caller can ignore this function
// and set this function to nil. Otherwise, setNodeValueEntryComparator will not know
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
func NewSet(equalFunc SetExtComparatorFunc, smallerFunc SetExtComparatorFunc) *Set {
	comparator := setNodeValueEntryComparator{
		extEqualFunc:   equalFunc,
		extSmallerFunc: smallerFunc,
	}
	return &Set{
		tree: rbTree.RedBlackTree{
			Comparator: comparator,
		},
		size: uint64(0),
	}
}

// Find return iterator of the result node. Return nil if error
func (s *Set) Find(val interface{}) *SetIterator {
	if node, ok := s.tree.FindNode(setNodeValueEntry{value: val}); ok {
		return &SetIterator{node: node}
	}
	return nil
}

// Insert inserts an entry and return iterator of the result node. Return nil if error
func (s *Set) Insert(val interface{}) *SetIterator {
	if node, ok := s.tree.AddNode(setNodeValueEntry{value: val}); ok {
		s.size++
		return &SetIterator{node: node}
	}
	return nil
}

// Remove an entry from Set
func (s *Set) Remove(val interface{}) {
	if _, ok := s.tree.FindNode(setNodeValueEntry{value: val}); !ok {
		return
	}
	s.tree.RemoveNodeByVal(setNodeValueEntry{value: val})
	s.size--
}

// RemoveByIter an entry from Set by iterator
func (s *Set) RemoveByIter(it *SetIterator) {
	s.Remove(it.Value())
}

// Size the number of Set node
func (s *Set) Size() uint64 {
	return s.size
}
