package tree

import (
	"bytes"
	"data_structure/linear"
	"fmt"
	"strings"
)

type VisitNode func(e interface{})

type TreeNode struct {
	e interface{}
	left *TreeNode
	right *TreeNode
}

func (tn *TreeNode) Then(e interface{}) int {

	ret := 0
	switch e.(type) {
	case int:
		if e == tn.e {
			ret = 0
		} else if e.(int) > tn.e.(int) {
			ret = 1
		} else {
			ret = -1
		}
	case string:
		if e == tn.e {
			ret = 0
		} else if len(e.(string)) > len(tn.e.(string)) {
			ret = 1
		} else {
			ret = -1
		}
	}
	return ret
}

func NewTreeNode(e interface{}) *TreeNode {
	node := new(TreeNode)
	node.e = e
	node.left = nil
	node.right = nil
	return node
}

type BST struct {
	root *TreeNode
	size int
}

func (b *BST) GetSize() int {
	return b.size
}

func (b *BST) IsEmpty() bool {
	return b.size == 0
}

func (b *BST) Add(e interface{}) {
	b.root = b.addNode(b.root, e)
}

func (b *BST) addNode(node *TreeNode, e interface{}) *TreeNode {
	if node == nil {
		b.size++
		return NewTreeNode(e)
	}
	if node.Then(e) < 0 {
		node.left = b.addNode(node.left, e)
	} else if node.Then(e) > 0 {
		node.right = b.addNode(node.right, e)
	}
	return node
}

func (b *BST) Remove(e interface{}) {
	b.root = b.remove(b.root, e)
}

func (b *BST) remove(node *TreeNode, e interface{}) *TreeNode {

	if node == nil {
		return nil
	}

	if node.Then(e) < 0 {
		node.left = b.remove(node.left, e)
		return node
	} else if node.Then(e) > 0 {
		node.right = b.remove(node.right, e)
		return node
	} else {

		if node.left == nil {
			rightNode := node.right
			node.right = nil
			b.size--
			return rightNode
		}

		if node.right == nil {
			leftNode := node.left
			node.left = nil
			b.size--
			return leftNode
		}

		successor := b.minimum(node.right)
		successor.right = b.removeMini(node.right)
		successor.left = node.left

		node.left = nil
		node.right = nil

		return successor
	}
}

func (b *BST) RemoveMini() interface{} {
	ret := b.Minimum()
	b.root = b.removeMini(b.root)
	return ret
}

func (b *BST) removeMini(node *TreeNode) *TreeNode {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		b.size--
		return rightNode
	}
	node.left = b.removeMini(node.left)
	return node
}

func (b *BST) RemoveMax() interface{}  {
	ret := b.Maxmum()
	b.root = b.removeMax(b.root)
	return ret
}

func (b *BST) removeMax(node *TreeNode) *TreeNode {
	if node.right == nil {
		leftNode := node.left
		node.left = nil
		b.size--
		return leftNode
	}
	node.right = b.removeMax(node.right)
	return node
}

func (b *BST) Contains(element interface{}) bool {
	has := false
	b.PreOrder(func(e interface{}) {
		if e == element {
			has = true
		}
	})
	return has
}

func (b *BST) Minimum() interface{} {
	if b.size == 0 {
		panic("BST is empty")
	}
	return b.minimum(b.root).e
}

func (b *BST) minimum(node *TreeNode) *TreeNode {
	if node.left == nil {
		return node
	}
	return b.minimum(node.left)
}

func (b *BST) Maxmum() interface{} {
	if b.size == 0 {
		panic("BST is empty")
	}
	return b.maxmum(b.root).e
}

func (b *BST) maxmum(node *TreeNode) *TreeNode {
	if node.right == nil {
		return node
	}
	return b.maxmum(node.right)
}

func (b *BST) PreOrder(call VisitNode) {
	b.preOrder(b.root, call)
}

func (b *BST) preOrder(node *TreeNode, call VisitNode) {
	if node == nil {
		return
	}
	call(node.e)
	b.preOrder(node.left, call)
	b.preOrder(node.right, call)
}

func (b *BST) PreOrderNR(call VisitNode) {
	stack := linear.NewLinkedListStack()
	stack.Push(b.root)
	for !stack.IsEmpty() {
		cur := stack.Pop()
		curNode := cur.(*TreeNode)
		call(curNode.e)
		if curNode.right != nil {
			stack.Push(curNode.right)
		}
		if curNode.left != nil {
			stack.Push(curNode.left)
		}
	}
}

func (b *BST) InOrder(call VisitNode) {
	b.inOrder(b.root, call)
}

func (b *BST) inOrder(node *TreeNode, call VisitNode) {
	if node == nil {
		return
	}
	b.inOrder(node.left, call)
	call(node.e)
	b.inOrder(node.right, call)
}

func (b *BST) BackOrder(call VisitNode) {
	b.backOrder(b.root, call)
}

func (b *BST) backOrder(node *TreeNode, call VisitNode) {
	if node == nil {
		return
	}
	b.inOrder(node.left, call)
	b.inOrder(node.right, call)
	call(node.e)
}

func (b *BST) LevelOrder(call VisitNode) {
	queue := linear.NewLinkedListQueue()
	queue.EnQueue(b.root)

	for !queue.IsEmpty() {
		cur := queue.DeQueue()
		curNode := cur.(*TreeNode)

		call(curNode.e)

		if curNode.left != nil {
			queue.EnQueue(curNode.left)
		}
		if curNode.right != nil {
			queue.EnQueue(curNode.right)
		}
	}
}

func (b *BST) ToString() string {
	buff := new(bytes.Buffer)
	b.toStr(b.root, 0, buff)
	return buff.String()
}

func (b *BST) toStr(node *TreeNode, depth int, buff *bytes.Buffer)  {
	buff.WriteString(strings.Repeat("-", depth))
	if node == nil {
		buff.WriteString("nil \n")
		return
	}
	buff.WriteString(fmt.Sprintf("%v", node.e) + "\n")
	b.toStr(node.left, depth+1, buff)
	b.toStr(node.right, depth+1, buff)
}

func NewBST() *BST {
	return &BST{root: nil, size: 0}
}
