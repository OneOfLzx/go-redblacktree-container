package go_redblacktree_container

import (
	"testing"
)

type testKey struct {
	idx  int
	name string
}

type testKeyValuePair struct {
	key   testKey
	value int
}

func testEqualFunc(v1 interface{}, v2 interface{}) bool {
	val1 := v1.(testKey)
	val2 := v2.(testKey)
	return val1.idx == val2.idx
}

func testSmallerFunc(v1 interface{}, v2 interface{}) bool {
	val1 := v1.(testKey)
	val2 := v2.(testKey)
	return val1.idx < val2.idx
}

func TestMap(t *testing.T) {
	dataArray := []testKeyValuePair{
		{testKey{0, "0"}, -686},
		{testKey{1, "1"}, -623},
		{testKey{2, "2"}, -589},
		{testKey{3, "3"}, -562},
		{testKey{4, "4"}, -532},
		{testKey{5, "5"}, -327},
		{testKey{6, "6"}, -238},
		{testKey{7, "7"}, -182},
		{testKey{8, "8"}, -176},
		{testKey{9, "9"}, -156},
		{testKey{10, "10"}, -103},
		{testKey{11, "11"}, -82},
		{testKey{12, "12"}, -66},
		{testKey{13, "13"}, -37},
		{testKey{14, "14"}, 1},
		{testKey{15, "15"}, 23},
		{testKey{16, "16"}, 56},
		{testKey{17, "17"}, 58},
		{testKey{18, "18"}, 78},
		{testKey{19, "19"}, 91},
		{testKey{20, "20"}, 158},
		{testKey{21, "21"}, 176},
		{testKey{22, "22"}, 358},
		{testKey{23, "23"}, 377},
		{testKey{24, "24"}, 379},
		{testKey{25, "25"}, 389},
		{testKey{26, "26"}, 399},
		{testKey{27, "27"}, 566},
		{testKey{28, "28"}, 578},
		{testKey{29, "29"}, 591},
		{testKey{30, "30"}, 686}}

	myMap := NewMap(testEqualFunc, testSmallerFunc)

	var it *MapIterator = nil
	for _, val := range dataArray {
		it = myMap.Insert(val.key, val.value)
		if it.Key().(testKey).idx != val.key.idx ||
			it.Key().(testKey).name != val.key.name ||
			it.Value().(int) != val.value {
			t.Error("Insert error")
			return
		}
	}

	for _, val := range dataArray {
		it = myMap.Find(val.key)
		if it.Key().(testKey).idx != val.key.idx ||
			it.Key().(testKey).name != val.key.name ||
			it.Value().(int) != val.value {
			t.Error("Find error")
			return
		}
	}

	idx := len(dataArray) - 1
	for it.Prev() {
		idx--
		if it.Key().(testKey).idx != dataArray[idx].key.idx ||
			it.Key().(testKey).name != dataArray[idx].key.name ||
			it.Value().(int) != dataArray[idx].value {
			t.Error("Wrong order")
			return
		}
	}

	for _, val := range dataArray {
		myMap.Remove(val.key)
	}
	if myMap.Size() != 0 {
		t.Error("Remove error")
	}
}
