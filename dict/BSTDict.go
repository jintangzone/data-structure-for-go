package dict

import (
	"fmt"
	"hash/crc32"
	"io"
)

type BSTDictNode struct {
	k interface{}
	v interface{}
	left *BSTDictNode
	right *BSTDictNode
}

func (bdn *BSTDictNode) CompareTo(k interface{}) int {
	ret := 0
	if bdn.k != k {
		switch k.(type) {
		case string:
			iEEE := crc32.NewIEEE()
			io.WriteString(iEEE, bdn.k.(string))
			bdnK := iEEE.Sum32()

			iEEE.Reset()
			io.WriteString(iEEE, k.(string))
			K := iEEE.Sum32()

			if K < bdnK {
				ret = -1
			} else {
				ret = 1
			}
		case int:
			if k.(int) < bdn.k.(int) {
				ret = -1
			} else {
				ret = 1
			}
		}
	}
	return ret
}

func NewBSTDictNode(k interface{}, v interface{}, left, right *BSTDictNode) *BSTDictNode  {
	return &BSTDictNode{ k: k, v: v, left: left, right: right }
}

type BSTDict struct {
	root *BSTDictNode
	size int
}

func (ld *BSTDict) Add(k interface{}, v interface{}) {
	ld.root = ld.add(ld.root, k, v)
}

func (ld *BSTDict) add(node *BSTDictNode, k interface{}, v interface{}) *BSTDictNode {
	if node == nil {
		ld.size++
		return NewBSTDictNode(k, v, nil, nil)
	}
	if node.CompareTo(k) < 0 {
		node.left = ld.add(node.left, k, v)
	} else if node.CompareTo(k) > 0 {
		node.right = ld.add(node.right, k, v)
	} else {
		node.v = v
	}
	return node
}

func (ld *BSTDict) Remove(k interface{}) interface{} {
	node := ld.getNode(ld.root, k)
	if node != nil {
		ld.root = ld.remove(ld.root, k)
		return node.v
	}
	return nil
}

func (ld *BSTDict) remove(node *BSTDictNode, k interface{}) *BSTDictNode {

	if node == nil {
		return nil
	}

	if node.CompareTo(k) < 0 {
		node.left = ld.remove(node.left, k)
		return node
	} else if node.CompareTo(k) > 0 {
		node.right = ld.remove(node.right, k)
		return node
	} else {

		if node.left == nil {
			right := node.right
			node.right = nil
			ld.size--
			return right
		}

		if node.right == nil {
			left := node.left
			node.left = nil
			ld.size--
			return left
		}

		s := ld.minimum(node.right)
		s.right = ld.removeMin(node.right)
		s.left = node.left

		node.left, node.right = nil, nil
		return s
	}
}

func (ld *BSTDict) minimum(node *BSTDictNode) *BSTDictNode {
	if node.left == nil {
		return node
	}
	return ld.minimum(node.left)
}

func (ld *BSTDict) removeMin(node *BSTDictNode) *BSTDictNode {
	if node.left == nil {
		right := node.right
		node.right = nil
		ld.size--
		return right
	}
	node.left = ld.removeMin(node.left)
	return node
}

func (ld *BSTDict) removeMax(node *BSTDictNode) *BSTDictNode {
	if node.right == nil {
		left := node.left
		node.left = nil
		ld.size--
		return left
	}
	node.right = ld.removeMax(node.right)
	return node
}

func (ld *BSTDict) Contains(k interface{}) bool {
	return ld.getNode(ld.root, k) != nil
}

func (ld *BSTDict) Get(k interface{}) interface{} {
	node := ld.getNode(ld.root, k)
	if node != nil {
		return node.v
	}
	return nil
}

func (ld *BSTDict) getNode(node *BSTDictNode, k interface{}) *BSTDictNode {
	if node == nil {
		return nil
	}
	if node.CompareTo(k) < 0 {
		return ld.getNode(node.left, k)
	} else if node.CompareTo(k) > 0 {
		return ld.getNode(node.right, k)
	} else {
		return node
	}
}

func (ld *BSTDict) Set(k interface{}, v interface{}) {
	node := ld.getNode(ld.root, k)
	if node == nil {
		panic(fmt.Sprint(k) + " doesn't exist!")
	}
	node.v = v
}

func (ld *BSTDict) GetSize() int {
	return ld.size
}

func (ld *BSTDict) IsEmpty() bool {
	return ld.size == 0
}

func NewBSTDict() *BSTDict {
	return &BSTDict{
		root: NewBSTDictNode("", nil, nil, nil),
		size: 0,
	}
}