package dict

import "fmt"

type LinkedDictNode struct {
	k interface{}
	v interface{}
	next *LinkedDictNode
}

func NewLinkedDictNode(k interface{}, v interface{}, next *LinkedDictNode) *LinkedDictNode  {
	return &LinkedDictNode{ k: k, v: v, next: next }
}

type LinkedDict struct {
	dummyHead *LinkedDictNode
	size int
}

func (ld *LinkedDict) Add(k interface{}, v interface{}) {
	node := ld.getNode(k)
	if node == nil {
		ld.dummyHead.next = NewLinkedDictNode(k, v, ld.dummyHead.next)
		ld.size++
	} else {
		node.v = v
	}
}

func (ld *LinkedDict) Remove(k interface{}) interface{} {

	pre := ld.dummyHead
	for pre.next != nil {
		if pre.next.k == k {
			break
		}
		pre = pre.next
	}

	if pre.next != nil {
		delNode := pre.next
		pre.next = delNode.next
		delNode.next = nil
		ld.size--
		return delNode.v
	}

	return nil
}

func (ld *LinkedDict) Contains(k interface{}) bool {
	return ld.getNode(k) != nil
}

func (ld *LinkedDict) Get(k interface{}) interface{} {
	node := ld.getNode(k)
	if node != nil {
		return node.v
	}
	return nil
}

func (ld *LinkedDict) getNode(k interface{}) *LinkedDictNode {
	p := ld.dummyHead.next
	for p != nil {
		if p.k == k {
			return p
		}
		p = p.next
	}
	return nil
}

func (ld *LinkedDict) Set(k interface{}, v interface{}) {
	node := ld.getNode(k)
	if node == nil {
		panic(fmt.Sprint(k) + " doesn't exist!")
	}
	node.v = v
}

func (ld *LinkedDict) GetSize() int {
	return ld.size
}

func (ld *LinkedDict) IsEmpty() bool {
	return ld.size == 0
}

func NewLinkedDict() *LinkedDict {
	return &LinkedDict{
		dummyHead: NewLinkedDictNode("", nil, nil),
		size: 0,
	}
}

