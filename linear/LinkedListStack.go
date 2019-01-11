package linear

import "data_structure/core"

type LinkedListStack struct {
	link *core.LinkedList
}

func (lls *LinkedListStack) Push(e interface{})  {
	lls.link.AddFirst(e)
}

func (lls *LinkedListStack) Pop() interface{}  {
	return lls.link.RemoveFirst()
}

func (lls *LinkedListStack) Peek() interface{} {
	return lls.link.GetFirst()
}

func (lls *LinkedListStack) GetSize() int {
	return lls.link.GetSize()
}

func (lls *LinkedListStack) IsEmpty() bool  {
	return lls.link.IsEmpty()
}

func (lls *LinkedListStack) ToString() string {
	str := "LinkedListStack: top "
	str = str + lls.link.ToString()
	return str
}

func NewLinkedListStack() *LinkedListStack  {
	link := new(LinkedListStack)
	link.link = core.NewLinkedList()
	return link
}