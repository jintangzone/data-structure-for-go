package LinkedList

import (
	"fmt"
)

type LinkedNode struct {
	Next *LinkedNode
	Value interface{}
}

func (ln *LinkedNode) LinkTo(node *LinkedNode)  {
	ln.Next = node
}

type LinkedList struct {
	Head *LinkedNode
	Tail *LinkedNode
	NodeTotal int
}

func (ll *LinkedList) Space() int {
	return ll.NodeTotal
}

func (ll *LinkedList) Append(value interface{}) {
	node := new(LinkedNode)
	node.Value = value

	if ll.Tail != nil {
		ll.Tail.Next = node
	}

	ll.Tail = node
	ll.NodeTotal++
	if ll.NodeTotal == 1 {
		ll.Head = node
	}
}

func (ll *LinkedList) Prepend(value interface{}) {
	node := new(LinkedNode)
	node.Value = value
	node.LinkTo(ll.Head)
	ll.Head = node
	ll.NodeTotal++
	if ll.NodeTotal == 1 {
		ll.Tail = node
	}
}

func (ll *LinkedList) Lookup(n int) *LinkedNode {
	if n == 0 {
		return ll.Head
	}

	if n >= ll.NodeTotal {
		return ll.Tail
	}

	pNode := ll.Head
	times := 0
	for pNode != nil {
		if times == n {
			break
		}
		times++
		pNode = pNode.Next
	}
	return pNode
}

func (ll *LinkedList) Insert(n int, value interface{}) {
	if n >= ll.NodeTotal {
		ll.Append(value)
		return
	}

	if n == 0 {
		ll.Prepend(value)
		return
	}

	node := ll.Lookup(n)
	newNode := new(LinkedNode)
	newNode.Value = value
	newNode.LinkTo(node.Next)
	node.LinkTo(newNode)
}

func (ll *LinkedList) Delete(n int) {

	if n <= 1 {
		head := ll.Head
		ll.Head = ll.Head.Next
		head.Next = nil
		return
	}

	node := ll.Lookup(n-1)
	node.LinkTo(node.Next.Next)
}

func (ll *LinkedList) Traversal() {

	fmt.Println("NodeTotal:", ll.Space())
	fmt.Println("Head =>", ll.Head.Value)
	fmt.Println("Tail =>", ll.Tail.Value)
	fmt.Println("---------------------------------")

	pNode := ll.Head
	for pNode != nil {
		fmt.Printf("%v -> ", pNode.Value)
		pNode = pNode.Next
	}
	fmt.Println()
	fmt.Println("---------------------------------")
}