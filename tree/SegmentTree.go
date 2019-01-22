package tree

import (
	"strconv"
	"strings"
)

type Merger interface {
	merge(a, b interface{}) interface{}
}

type SegmentTree struct {
	data []int
	tree []int
}

func (st *SegmentTree) Get(index int) int {

	if index < 0 || index >= len(st.data) {
		panic("Index is illegal.")
	}
	return st.data[index]
}

func (st *SegmentTree) GetSize() int {
	return len(st.data)
}

func (st *SegmentTree) leftChild(index int) int {
	return 2*index+1
}

func (st *SegmentTree) rightChild(index int) int {
	return 2*index+2
}

func (st *SegmentTree) buildSegmentTree(treeIndex, l, r int) {
	if l == r {
		st.tree[treeIndex] = st.data[l]
		return
	}

	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)

	mid := l + (r-l)/2

	st.buildSegmentTree(leftTreeIndex, l, mid)
	st.buildSegmentTree(rightTreeIndex, mid+1, r)

	st.tree[treeIndex] = st.tree[leftTreeIndex] + st.tree[rightTreeIndex]
}

func (st *SegmentTree) Query(ql, qr int) int {
	if ql < 0 || ql >= len(st.data) || qr < 0 || qr >= len(st.data) || ql > qr {
		panic("index is illegal.")
	}
	return st.query(0, 0, len(st.data) - 1, ql, qr)
}

// treeIndex为根节点，意思为查询以treeIndex为根的
func (st *SegmentTree) query(treeIndex, l, r, ql, qr int) int {
	if l == ql && r == qr {
		return st.tree[treeIndex]
	}

	mid := l + (r-l)/2
	leftChildIndex := st.leftChild(treeIndex)
	rightChildIndex := st.rightChild(treeIndex)

	if ql >= mid+1 {
		return st.query(rightChildIndex, mid+1, r, ql, qr)
	} else if qr <= mid {
		return st.query(leftChildIndex, l, mid, ql, qr)
	} else {
		leftRes := st.query(leftChildIndex, l, mid, ql, mid)
		rightRes := st.query(rightChildIndex, mid+1, r, mid+1, qr)
		return leftRes + rightRes
	}
}

func (st *SegmentTree) Set(index, e int)  {
	if index < 0 || index >= len(st.data) {
		panic("index is illegal.")
	}
	st.set(0, 0, len(st.data)-1, index, e)
}

func (st *SegmentTree) set(treeIndex, l, r, index, e int)  {
	if l == r {
		st.tree[treeIndex] = e
		return
	}

	leftTreeIndex := st.leftChild(treeIndex)
	rightTreeIndex := st.rightChild(treeIndex)
	mid := l + (r - l)/2

	if index >= mid+1 {
		st.set(rightTreeIndex, mid+1, r, index, e)
	} else {
		st.set(leftTreeIndex, l, mid, index, e)
	}
	st.tree[treeIndex] = st.tree[leftTreeIndex] + st.tree[rightTreeIndex]
}

func (st *SegmentTree) ToString() string {

	str := make([]string, 0)
	str = append(str, "segment tree: [")

	for i := 0; i < len(st.tree); i++ {
		if st.tree[i] != 0 {
			str = append(str, strconv.Itoa(st.tree[i]))
		} else {
			str = append(str, "nil")
		}
	}

	str = append(str, "]")

	return strings.Join(str, ", ")
}

func NewSegmentTree(arr []int) *SegmentTree {
	st := &SegmentTree{ data: arr }
	st.tree = make([]int, 4*len(arr))
	st.buildSegmentTree(0, 0, len(st.data)-1)
	return st
}

