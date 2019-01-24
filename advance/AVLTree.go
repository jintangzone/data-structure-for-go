package advance

import (
	"data_structure/core"
	"fmt"
	"math"
)

type AVLTreeNodeCall func(k core.Key, v interface{})

type AVLTreeNode struct {
	k core.Key
	v interface{}
	left *AVLTreeNode
	right *AVLTreeNode
	height int
}

func NewAVLTreeNode(k core.Key, v interface{}, left, right *AVLTreeNode) *AVLTreeNode {
	return &AVLTreeNode{k, v, left, right, 1}
}

type AVLTree struct {
	root *AVLTreeNode
	size int
}

func (avl *AVLTree) GetSize() int {
	return avl.size
}

func (avl *AVLTree) IsEmpty() bool {
	return avl.size == 0
}

func (avl *AVLTree) IsBST() bool {
	arr := make([]core.Key, 0)

	avl.InOrder(func(k core.Key, v interface{}) {
		arr = append(arr, k)
	})

	for i := 1; i < len(arr); i++ {
		if arr[i-1].Then(arr[i]) > 0 {
			return false
		}
	}
	return true
}

func (avl *AVLTree) IsBalanced() bool {
	return avl.isBalanced(avl.root)
}

func (avl *AVLTree) isBalanced(node *AVLTreeNode) bool {
	if node == nil {
		return true
	}
	balanceFactor := avl.getBalanceFactor(node)
	if math.Abs(float64(balanceFactor)) > 1 {
		return false
	}
	return avl.isBalanced(node.left) && avl.isBalanced(node.right)
}

func (avl *AVLTree) getHeight(node *AVLTreeNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (avl *AVLTree) getBalanceFactor(node *AVLTreeNode) int {
	if node == nil {
		return 0
	}
	return avl.getHeight(node.left) - avl.getHeight(node.right)
}

func (avl *AVLTree) leftRotate(y *AVLTreeNode) *AVLTreeNode  {
	x := y.right
	T2 := x.left
	x.left = y
	y.right = T2
	y.height = 1 + int(math.Max(float64(avl.getHeight(y.left)), float64(avl.getHeight(y.right))))
	x.height = 1 + int(math.Max(float64(avl.getHeight(x.left)), float64(avl.getHeight(x.right))))
	return x
}

func (avl *AVLTree) rightRotate(y *AVLTreeNode) *AVLTreeNode {
	x := y.left
	T3 := x.right
	x.right = y
	y.left = T3
	y.height = 1 + int(math.Max(float64(avl.getHeight(y.left)), float64(avl.getHeight(y.right))))
	x.height = 1 + int(math.Max(float64(avl.getHeight(x.left)), float64(avl.getHeight(x.right))))
	return x
}

func (avl *AVLTree) Add(k core.Key, v interface{})  {
	avl.root = avl.add(avl.root, k, v)
}

func (avl *AVLTree) add(node *AVLTreeNode, k core.Key, v interface{}) *AVLTreeNode {
	if node == nil {
		avl.size++
		return NewAVLTreeNode(k, v, nil, nil)
	}
	if k.Then(node.k) < 0 {
		node.left = avl.add(node.left, k, v)
	} else if k.Then(node.k) > 0 {
		node.right = avl.add(node.right, k, v)
	} else {
		node.v = v
	}

	node.height = 1 + int(math.Max(float64(avl.getHeight(node.left)), float64(avl.getHeight(node.right))))

	balanceFactor := avl.getBalanceFactor(node)

	//if math.Abs(float64(balanceFactor)) > 1 {
	//	fmt.Println("unbalanced", ":", balanceFactor)
	//}

	// LL
	if balanceFactor > 1 && avl.getBalanceFactor(node.left) >= 0 {
		return avl.rightRotate(node)
	}

	// RR
	if balanceFactor < -1 && avl.getBalanceFactor(node.right) <= 0 {
		return avl.leftRotate(node)
	}

	// LR
	if balanceFactor > 1 && avl.getBalanceFactor(node.left) < 0 {
		node.left = avl.leftRotate(node.left)
		return avl.rightRotate(node)
	}

	// RL
	if balanceFactor < -1 && avl.getBalanceFactor(node.right) > 0 {
		node.right = avl.rightRotate(node.right)
		return avl.leftRotate(node)
	}

	return node
}

func (avl *AVLTree) Remove(k core.Key) interface{} {
	ret := avl.Get(k)
	avl.root = avl.remove(avl.root, k)
	return ret
}

func (avl *AVLTree) remove(node *AVLTreeNode, k core.Key) *AVLTreeNode {
	if node == nil {
		return nil
	}

	retNode := new(AVLTreeNode)
	if k.Then(node.k) < 0 {
		node.left = avl.remove(node.left, k)
		retNode = node
	} else if k.Then(node.k) > 0 {
		node.right = avl.remove(node.right, k)
		retNode = node
	} else {

		if node.left == nil {
			rightNode := node.right
			node.right = nil
			avl.size--
			retNode = rightNode
		} else if node.right == nil {
			leftNode := node.left
			node.left = nil
			avl.size--
			retNode = leftNode
		} else {
			successor := avl.minimum(node.right)
			successor.right = avl.remove(node.right, successor.k)
			successor.left = node.left
			node.left = nil
			node.right = nil

			retNode = successor
		}

	}

	if retNode == nil {
		return nil
	}

	retNode.height = 1 + int(math.Max(float64(avl.getHeight(retNode.left)), float64(avl.getHeight(retNode.right))))

	balanceFactor := avl.getBalanceFactor(retNode)

	// LL
	if balanceFactor > 1 && avl.getBalanceFactor(retNode.left) >= 0 {
		return avl.rightRotate(retNode)
	}

	// RR
	if balanceFactor < -1 && avl.getBalanceFactor(retNode.right) <= 0 {
		return avl.leftRotate(retNode)
	}

	// LR
	if balanceFactor > 1 && avl.getBalanceFactor(retNode.left) < 0 {
		retNode.left = avl.leftRotate(retNode.left)
		return avl.rightRotate(retNode)
	}

	// RL
	if balanceFactor < -1 && avl.getBalanceFactor(retNode.right) > 0 {
		retNode.right = avl.rightRotate(retNode.right)
		return avl.leftRotate(retNode)
	}

	return retNode
}

func (avl *AVLTree) RemoveMin() interface{} {
	ret := avl.Minimum()
	avl.root = avl.removeMin(avl.root)
	return ret
}

func (avl *AVLTree) removeMin(node *AVLTreeNode) *AVLTreeNode {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		avl.size--
		return rightNode
	}
	node.left = avl.removeMin(node.left)
	return node
}

func (avl *AVLTree) RemoveMax() interface{} {
	ret := avl.Maximum()
	avl.root = avl.removeMax(avl.root)
	return ret
}

func (avl *AVLTree) removeMax(node *AVLTreeNode) *AVLTreeNode {
	if node.right == nil {
		leftNode := node.left
		node.left = nil
		avl.size--
		return leftNode
	}
	node.right = avl.removeMax(node.right)
	return node
}

func (avl *AVLTree) Minimum() interface{} {
	if avl.IsEmpty() {
		panic("avl tree is empty.")
	}
	return avl.minimum(avl.root).v
}

func (avl *AVLTree) minimum(node *AVLTreeNode) *AVLTreeNode {
	if node.left == nil {
		return node
	}
	return avl.minimum(node.left)
}

func (avl *AVLTree) Maximum() interface{} {
	if avl.IsEmpty() {
		panic("avl tree is empty.")
	}
	return avl.maximum(avl.root).v
}

func (avl *AVLTree) maximum(node *AVLTreeNode) *AVLTreeNode {
	if node.right == nil {
		return node
	}
	return avl.maximum(node.right)
}

func (avl *AVLTree) Contain(key core.Key) bool {
	has := false
	avl.PreOrder(func(k core.Key, v interface{}) {
		if k == key {
			has = true
		}
	})
	return has
}

func (avl *AVLTree) Get(k core.Key) interface{} {
	node := avl.get(avl.root, k)
	if node != nil {
		return node.v
	}
	return nil
}

func (avl *AVLTree) get(node *AVLTreeNode, k core.Key) *AVLTreeNode {
	if node == nil {
		return nil
	}
	if k.Then(node.k) < 0 {
		return avl.get(node.left, k)
	} else if k.Then(node.k) > 0 {
		return avl.get(node.right, k)
	} else {
		return node
	}
}

func (avl *AVLTree) Set(k core.Key, v interface{}) {
	node := avl.get(avl.root, k)
	if node == nil {
		panic(fmt.Sprint(k) + " doesn't exist!")
	}
	node.v = v
}

func (avl *AVLTree) PreOrder(call AVLTreeNodeCall) {
	avl.preOrder(avl.root, call)
}

func (avl *AVLTree) preOrder(node *AVLTreeNode, call AVLTreeNodeCall)  {
	if node == nil {
		return
	}
	call(node.k, node.v)
	avl.preOrder(node.left, call)
	avl.preOrder(node.right, call)
}

func (avl *AVLTree) PreOrderNR(call AVLTreeNodeCall) {
	stack := make([]*AVLTreeNode, 0)
	stack = append(stack, avl.root)
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		call(cur.k, cur.v)
		stack = stack[:len(stack)-1]
		if cur.right != nil {
			stack = append(stack, cur.right)
		}
		if cur.left != nil {
			stack = append(stack, cur.left)
		}
	}
}

func (avl *AVLTree) InOrder(call AVLTreeNodeCall) {
	avl.inOrder(avl.root, call)
}

func (avl *AVLTree) inOrder(node *AVLTreeNode, call AVLTreeNodeCall)  {
	if node == nil {
		return
	}
	avl.inOrder(node.left, call)
	call(node.k, node.v)
	avl.inOrder(node.right, call)
}

func (avl *AVLTree) BackOrder(call AVLTreeNodeCall) {
	avl.backOrder(avl.root, call)
}

func (avl *AVLTree) backOrder(node *AVLTreeNode, call AVLTreeNodeCall)  {
	if node == nil {
		return
	}
	avl.backOrder(node.left, call)
	avl.backOrder(node.right, call)
	call(node.k, node.v)
}

func (avl *AVLTree) LevelOrder(call AVLTreeNodeCall) {
	queue := make([]*AVLTreeNode, 0)
	queue = append(queue, avl.root)
	for len(queue) > 0 {
		top := queue[0]
		call(top.k, top.v)
		queue = queue[1:]
		if top.left != nil {
			queue = append(queue, top.left)
		}
		if top.right != nil {
			queue = append(queue, top.right)
		}
	}
}

func NewAVLTree() *AVLTree {
	return &AVLTree{
		root: nil,
		size: 0,
	}
}
