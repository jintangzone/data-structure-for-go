package linear

import (
	"data_structure/core"
	"fmt"
)

type LinkedListQueue struct {
	head *core.Node
	tail *core.Node
	size int
}

func (llq *LinkedListQueue) EnQueue(e interface{}) {
	if llq.tail == nil {
		llq.tail = core.NewNode(e, nil)
		llq.head = llq.tail
	} else {
		llq.tail.Next = core.NewNode(e, nil)
		llq.tail = llq.tail.Next
	}
	llq.size++
}

func (llq *LinkedListQueue) DeQueue() interface{} {
	if llq.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	ret := llq.head
	llq.head = llq.head.Next
	ret.Next = nil
	if llq.head == nil {
		llq.tail = nil
	}
	llq.size--
	return ret.E
}

func (llq *LinkedListQueue) GetFront() interface{} {
	if llq.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	return llq.head.E
}

func (llq *LinkedListQueue) GetSize() int {
	return llq.size
}

func (llq *LinkedListQueue) IsEmpty() bool {
	return llq.size == 0
}

func (llq *LinkedListQueue) ToString() string {
	str := "LinkedListQueue: front "
	cur := llq.head
	for cur != nil {
		str = str + fmt.Sprintf("%v", cur.E) + " -> "
		cur = cur.Next
	}
	str = str + "NULL tail"
	return str
}

func NewLinkedListQueue() *LinkedListQueue {
	queue := new(LinkedListQueue)
	queue.head = nil
	queue.tail = nil
	queue.size = 0
	return queue
}