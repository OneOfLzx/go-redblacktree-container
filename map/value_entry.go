package go_redblacktree_container

import (
	rbTree "github.com/OneOfLzx/go-redblacktree"
)

// mapNodeValueEntry the value entry of underlying red-black tree node.
// Implement RedBlackTreeNodeValEntry interface
type mapNodeValueEntry struct {
	rbTree.RedBlackTreeNodeValEntry
	key   interface{}
	value interface{}
}

// MapExtComparatorFunc when the Map save built-in type, caller can ignore this function
// and set this function to nil. Otherwise, mapNodeValueEntryComparator will not know
// how to compare value, so that caller need to implement this function and custom the
// comparison operation
type MapExtComparatorFunc func(interface{}, interface{}) bool

// mapNodeValueEntryComparator the comparator of Map.
// Implement NodeValueEntryComparator interface
type mapNodeValueEntryComparator struct {
	rbTree.NodeValueEntryComparator
	extEqualFunc, extSmallerFunc MapExtComparatorFunc
}

func (c mapNodeValueEntryComparator) Equal(val1 rbTree.RedBlackTreeNodeValEntry, val2 rbTree.RedBlackTreeNodeValEntry) bool {
	ret := false
	v1 := val1.(mapNodeValueEntry).key
	v2 := val2.(mapNodeValueEntry).key
	switch v1.(type) {
	case int8:
		if v1.(int8) == v2.(int8) {
			ret = true
		}
	case uint8:
		if v1.(uint8) == v2.(uint8) {
			ret = true
		}
	case int16:
		if v1.(int16) == v2.(int16) {
			ret = true
		}
	case uint16:
		if v1.(uint16) == v2.(uint16) {
			ret = true
		}
	case int32:
		if v1.(int32) == v2.(int32) {
			ret = true
		}
	case uint32:
		if v1.(uint32) == v2.(uint32) {
			ret = true
		}
	case int64:
		if v1.(int64) == v2.(int64) {
			ret = true
		}
	case uint64:
		if v1.(uint64) == v2.(uint64) {
			ret = true
		}
	case int:
		if v1.(int) == v2.(int) {
			ret = true
		}
	case uint:
		if v1.(uint) == v2.(uint) {
			ret = true
		}
	case uintptr:
		if v1.(uintptr) == v2.(uintptr) {
			ret = true
		}
	case float32:
		if v1.(float32) == v2.(float32) {
			ret = true
		}
	case float64:
		if v1.(float64) == v2.(float64) {
			ret = true
		}
	case complex64:
		if v1.(complex64) == v2.(complex64) {
			ret = true
		}
	case complex128:
		if v1.(complex128) == v2.(complex128) {
			ret = true
		}
	case string:
		if v1.(string) == v2.(string) {
			ret = true
		}
	default:
		if nil == c.extEqualFunc {
			panic("Do not know how to compare(equal)!")
		}
		ret = c.extEqualFunc(v1, v2)
	}
	return ret
}

func (c mapNodeValueEntryComparator) Smaller(val1 rbTree.RedBlackTreeNodeValEntry, val2 rbTree.RedBlackTreeNodeValEntry) bool {
	ret := false
	v1 := val1.(mapNodeValueEntry).key
	v2 := val2.(mapNodeValueEntry).key
	switch v1.(type) {
	case int8:
		if v1.(int8) < v2.(int8) {
			ret = true
		}
	case uint8:
		if v1.(uint8) < v2.(uint8) {
			ret = true
		}
	case int16:
		if v1.(int16) < v2.(int16) {
			ret = true
		}
	case uint16:
		if v1.(uint16) < v2.(uint16) {
			ret = true
		}
	case int32:
		if v1.(int32) < v2.(int32) {
			ret = true
		}
	case uint32:
		if v1.(uint32) < v2.(uint32) {
			ret = true
		}
	case int64:
		if v1.(int64) < v2.(int64) {
			ret = true
		}
	case uint64:
		if v1.(uint64) < v2.(uint64) {
			ret = true
		}
	case int:
		if v1.(int) < v2.(int) {
			ret = true
		}
	case uint:
		if v1.(uint) < v2.(uint) {
			ret = true
		}
	case uintptr:
		if v1.(uintptr) < v2.(uintptr) {
			ret = true
		}
	case float32:
		if v1.(float32) < v2.(float32) {
			ret = true
		}
	case float64:
		if v1.(float64) < v2.(float64) {
			ret = true
		}
	case string:
		if v1.(string) < v2.(string) {
			ret = true
		}
	default:
		if nil == c.extSmallerFunc {
			panic("Do not know how to compare(smaller)!")
		}
		ret = c.extSmallerFunc(v1, v2)
	}
	return ret
}
