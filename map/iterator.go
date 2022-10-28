package go_redblacktree_container

import (
	rbTree "github.com/OneOfLzx/go-redblacktree"
)

// MapIterator the iterator of Map
type MapIterator struct {
	node *rbTree.RedBlackTreeNode
}

func (it *MapIterator) Key() interface{} {
	return it.node.Value().(mapNodeValueEntry).key
}

func (it *MapIterator) Value() interface{} {
	return it.node.Value().(mapNodeValueEntry).value
}

func (it *MapIterator) Prev() bool {
	if prevNode := it.node.PrevNode(); nil != prevNode {
		it.node = prevNode
		return true
	}
	return false
}

func (it *MapIterator) Next() bool {
	if nextNode := it.node.NextNode(); nil != nextNode {
		it.node = nextNode
		return true
	}
	return false
}
