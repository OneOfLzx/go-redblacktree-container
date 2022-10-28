package go_redblacktree_container

import (
	"testing"
)

func TestSet(t *testing.T) {
	dataArray := []int{-686, -623, -589, -562, -532, -327, -238, -182, -176, -156, -103, -82, -66, -37, 1,
		23, 56, 58, 78, 91, 158, 176, 358, 377, 379, 389, 399, 566, 578, 591, 686}

	set := NewSet(nil, nil)

	var it *SetIterator = nil
	for _, val := range dataArray {
		it = set.Insert(val)
		if it.Value().(int) != val {
			t.Error("Insert error")
			return
		}
	}

	for _, val := range dataArray {
		it = set.Find(val)
		if it.Value().(int) != val {
			t.Error("Find error")
			return
		}
	}

	idx := len(dataArray) - 1
	for it.Prev() {
		idx--
		if it.Value().(int) != dataArray[idx] {
			t.Error("Wrong order, iter.Value: ", it.Value(), ", dataArray[", idx, "]: ", dataArray[idx])
			return
		}
	}

	for _, val := range dataArray {
		set.Remove(val)
	}
	if set.Size() != 0 {
		t.Error("Remove error")
	}
}
