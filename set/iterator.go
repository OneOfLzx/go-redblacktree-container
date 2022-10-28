package go_redblacktree_container

import (
	rbTree "github.com/OneOfLzx/go-redblacktree"
)

// SetIterator the iterator of Set
type SetIterator struct {
	node *rbTree.RedBlackTreeNode
}

func (it *SetIterator) Value() interface{} {
	return it.node.Value().(setNodeValueEntry).value
}

func (it *SetIterator) Prev() bool {
	if prevNode := it.node.PrevNode(); nil != prevNode {
		it.node = prevNode
		return true
	}
	return false
}

func (it *SetIterator) Next() bool {
	if nextNode := it.node.NextNode(); nil != nextNode {
		it.node = nextNode
		return true
	}
	return false
}
