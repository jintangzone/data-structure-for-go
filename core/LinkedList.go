package core

import "fmt"

type Node struct {
	Next *Node
	E interface{}
}

func NewNode(e interface{}, next *Node) *Node {
	n := new(Node)
	n.E = e
	n.Next = next
	return n
}

type LinkedList struct {
	dummyHead *Node
	size int
}

func (link *LinkedList) GetSize() int {
	return link.size
}

func (link *LinkedList) IsEmpty() bool {
	return link.size == 0
}

func (link *LinkedList) AddFirst(e interface{}) {
	link.Insert(0, e)
}

func (link *LinkedList) AddLast(e interface{}) {
	link.Insert(link.size, e)
}

func (link *LinkedList) Insert(index int, e interface{}) {
	if index < 0 || index > link.size {
		panic("Insert failed, index illegal.")
	}

	prev := link.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	prev.Next = NewNode(e, prev.Next)
	link.size++
}

func (link *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= link.size {
		panic("Get failed, index illegal.")
	}

	cur := link.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	return cur.E
}

func (link *LinkedList) Set(index int, e interface{}) {
	if index < 0 || index >= link.size {
		panic("Set failed, index illegal.")
	}

	cur := link.dummyHead.Next
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.E = e
}

func (link *LinkedList) GetFirst() interface{} {
	return link.Get(0)
}

func (link *LinkedList) GetLast() interface{} {
	return link.Get(link.size-1)
}

func (link *LinkedList) Contains(e interface{}) bool {
	cur := link.dummyHead
	for cur != nil {
		if cur.E == e {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (link *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= link.size {
		panic("Remove failed, index illegal.")
	}

	prev := link.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.Next
	}

	removeNode := prev.Next
	prev.Next = removeNode.Next
	removeNode.Next = nil
	link.size--
	return removeNode.E
}

func (link *LinkedList) RemoveFirst() interface{} {
	return link.Remove(0)
}

func (link *LinkedList) RemoveLast() interface{} {
	return link.Remove(link.size-1)
}

func (link *LinkedList) ToString() string {
	str := "link: "
	cur := link.dummyHead
	for cur != nil {
		str = str + fmt.Sprintf("%v", cur.E) + " -> "
		cur = cur.Next
	}
	str = str + "NULL"
	return str
}

func NewLinkedList() *LinkedList {
	link := new(LinkedList)
	link.dummyHead = NewNode(nil, nil)
	link.size = 0
	return link
}