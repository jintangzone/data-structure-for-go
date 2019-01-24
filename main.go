package main

import (
	"data_structure/advance"
	"data_structure/core"
	"data_structure/dict"
	"data_structure/heap"
	"data_structure/linear"
	"data_structure/set"
	"data_structure/tree"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	// Array
	// testArray()

	// ArrayStack
	// testArrayStack()

	// ArrayQueue
	//testArrayQueue()

	// ArrayLoopQueue
	//testArrayLoopQueue()

	// queue pk
	//opCount := 100000
	//queue := linear.NewArrayQueue(opCount)
	//fmt.Printf("ArrayQueue execute times: %.6fs\n", testQueue(queue, opCount))
	//loopQueue := linear.NewArrayLoopQueue(opCount)
	//fmt.Printf("ArrayLoopQueue execute times: %.6fs\n", testQueue(loopQueue, opCount))
	//linkedListQueue := linear.NewLinkedListQueue()
	//fmt.Printf("ArrayLinkedListQueue execute times: %.6fs\n", testQueue(linkedListQueue, opCount))

	// linkedList
	// testLinkedlist()

	// LinkedListStack
	// testLinkedListStack()

	// stack pk
	//opCount := 10000
	//stack := linear.NewArrayStack(opCount)
	//fmt.Printf("ArrayStack execute times: %.6fs\n", testStack(stack, opCount))
	//linkedStack := linear.NewLinkedListStack()
	//fmt.Printf("LinkedListStack execute times: %.6fs\n", testStack(linkedStack, opCount))

	// linkedListQueue
	// testLinkedListQueue()

	// binary search tree
	// testBSTree()

	// Set
	// testSet()

	// BST Set VS Linked Set
	//bstSet := set.NewBSTSet()
	//fmt.Printf("BST set, execute: %.4f \n", testSetTime(bstSet, 100000))
	//linkSet := set.NewLinkedSet()
	//fmt.Printf("Linked set, execute: %.4f \n", testSetTime(linkSet, 100000))

	// testDict
	// testDict()

	// BST Set VS Linked Set
	//bstDict := dict.NewBSTDict()
	//fmt.Printf("BST dict, execute: %.4f \n", testDictTime(bstDict, 100000))
	//linkDict := dict.NewLinkedDict()
	//fmt.Printf("Linked dict, execute: %.4f \n", testDictTime(linkDict, 100000))

	// heap
	// testMaxArrayHeap()

	//testHeapify()

	// test segment tree
	//testSegmentTree()

	// test trie
	// testTrie()

	// test vs trie
	//trieVsBSTSet()

	// test UnionFindArr

	// testUionFindArr()

	// test avl
	// testAVL()

	// test AVLMap
	testAVLMap()
}

func testAVLMap()  {

	m := advance.NewAVLMap()

	mp := map[myKey]string{
		1 : "php",
		2 : "java",
	}

	for k, v := range mp {
		m.Add(k, v)
	}

	fmt.Println(m.Contains(myKey(1)))

	fmt.Println(m.Get(myKey(2)))

	m.Remove(myKey(2))

	fmt.Println(m.Contains(myKey(2)))

	m.Set(myKey(1), "golang")
	fmt.Println(m.Get(myKey(1)))
}

type myKey int

func (n myKey) Then(k core.Key) int {

	if n < k.(myKey) {
		return -1
	} else if n > k.(myKey) {
		return 1
	}
	return 0
}

func (n myKey) HashCode() int {
	return int(n)
}

func testAVL() {

	avl := advance.NewAVLTree()

	ks := []myKey{5,3,6,8,4,2}
	vs := []string{"php", "java", "golang", "object-c", "cpp", "clang"}
	for idx, k := range ks {
		v := vs[idx]
		avl.Add(k, v)
	}

	//fmt.Println("PreOrderNR:")
	avl.InOrder(func(k core.Key, v interface{}) {
		fmt.Println(k, "=>", v)
	})

	fmt.Println("is BST:", avl.IsBST())
	fmt.Println("is Balanced:", avl.IsBalanced())
	//fmt.Println()
	//
	//fmt.Println("LevelOrderNR:")
	//avl.LevelOrder(func(e advance.E) {
	//	node := e.(*myNode)
	//	fmt.Println(node.k, "=>", node.v)
	//})
	//
	//fmt.Println()
	//
	//max := avl.Maximum().(*myNode)
	//fmt.Println("max:", max.k, "=>", max.v)
	//
	//min := avl.Minimum().(*myNode)
	//fmt.Println("min:", min.k, "=>", min.v)
	//
	//avl.RemoveMax()
	//avl.RemoveMin()
	//
	//avl.PreOrder(func(e advance.E) {
	//	node := e.(*myNode)
	//	fmt.Println(node.k, "=>", node.v)
	//})
	//
	//fmt.Println()
	//
	//avl.Remove(&myNode{k: 4})
	//
	//avl.PreOrder(func(e advance.E) {
	//	node := e.(*myNode)
	//	fmt.Println(node.k, "=>", node.v)
	//})
}

func testUionFindArr()  {

	uf := advance.NewUnionFindArr(10)

	fmt.Println(uf.ToString())
	fmt.Println(uf.IsConnected(1, 2))
	uf.Union(1, 2)
	fmt.Println(uf.ToString())
	fmt.Println(uf.IsConnected(1, 2))
	uf.Union(2, 3)
	fmt.Println(uf.ToString())

}

//生成随机字符串
func GetRandomString(l int64) string{
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var i int64
	for i = 0 ; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func trieVsBSTSet()  {

	//s := set.NewBSTSet()
	t := advance.NewTrie()

	opCount := 100000

	arr := make([]string, 0)
	for i := 0; i < opCount; i++ {
		arr = append(arr, GetRandomString(5))
	}

	startTime := time.Now().UnixNano()
	//for _, word := range arr  {
	//	s.Add(word)
	//}
	//for _, word := range arr  {
	//	s.Contains(word)
	//}
	endTime := time.Now().UnixNano()
	//fmt.Printf("BSTSet execute times: %.4fs \n", float32(endTime-startTime)/1000000000.0)

	startTime = time.Now().UnixNano()
	for _, word := range arr  {
		t.Add(word)
	}
	for _, word := range arr  {
		t.Contains(word)
	}
	endTime = time.Now().UnixNano()
	fmt.Printf("Trie execute times: %.4fs \n", float32(endTime-startTime)/1000000000.0)

}

func testTrie() {
	words := []string{
		"java",
		"php",
		"hello",
		"world",
	}

	t := advance.NewTrie()

	for _, str := range words {
		t.Add(str)
	}

	fmt.Println(t.Contains("php"))
	fmt.Println(t.Contains("java"))
	fmt.Println(t.Contains("cpp"))
}

func testSegmentTree()  {

	arr := []int{-2, 0, 3, -5, 2, -1}

	st := tree.NewSegmentTree(arr)

	fmt.Println(st.Query(0, 5))

}

func testHeapify()  {

	opCount := 100000000

	arr := make([]int, opCount)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < opCount; i++ {
		v := r.Intn(opCount)
		arr[i] = v
	}

	mah1 := heap.NewMaxArrayHeap(opCount)

	startTime := time.Now().UnixNano()
	for i := 0; i < opCount; i++ {
		mah1.Add(arr[i])
	}
	endTime := time.Now().UnixNano()
	fmt.Printf("add execute times: %.4fs \n", float32(endTime-startTime)/1000000000.0)

	startTime = time.Now().UnixNano()
	heap.NewMaxArrayHeapFromArr(arr)
	endTime = time.Now().UnixNano()
	fmt.Printf("heapify execute times: %.4fs \n", float32(endTime-startTime)/1000000000.0)

}

func testMaxArrayHeap()  {

	opCount := 100000

	mah := heap.NewMaxArrayHeap(opCount)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < opCount; i++ {
		v := r.Intn(opCount)
		mah.Add(v)
	}

	arr := make([]int, opCount)
	for i := 0; i < opCount; i++ {
		arr[i] = mah.ExtractMax()
	}

	for i := 1; i < opCount; i++ {
		if arr[i-1] < arr[i] {
			panic("Err")
		}
	}

	fmt.Println("Test MaxHeap Completed.")

}

func testDictTime(dict core.Dict, opCount int) float32 {

	startTime := time.Now().UnixNano()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < opCount; i++ {
		v := r.Intn(opCount)
		dict.Add(strconv.Itoa(v), v)
	}
	endTime := time.Now().UnixNano()
	return float32(endTime-startTime)/1000000000.0

}

func testDict()  {

	data := []string{"php", "java", "python", "ruby", "object-c", "cpp", "clang", "golang"}

	linkedDict := dict.NewLinkedDict()
	for _, word := range data {
		linkedDict.Add(word, len(word))
	}

	fmt.Println("size: ", linkedDict.GetSize())
	fmt.Println("len of 'php': ", linkedDict.Get("php"))
	fmt.Println("len of 'java': ", linkedDict.Get("java"))

	BSTDict := dict.NewBSTDict()
	for _, word := range data {
		BSTDict.Add(word, len(word))
	}

	fmt.Println("size: ", BSTDict.GetSize())
	fmt.Println("len of 'php': ", BSTDict.Get("php"))
	fmt.Println("len of 'java': ", BSTDict.Get("java"))
}

func testSetTime(s core.Set, opCount int) float32 {
	startTime := time.Now().UnixNano()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < opCount; i++ {
		s.Add(r.Intn(opCount))
	}
	endTime := time.Now().UnixNano()
	return float32(endTime-startTime)/1000000000.0
}

func testSet()  {

	bstSet := set.NewBSTSet()
	linkSet := set.NewLinkedSet()

	strs := []string{ "php", "java", "golang", "object-c", "clang", "cpp", "java" }
	for _, str := range strs {
		bstSet.Add(str)
		linkSet.Add(str)
	}

	fmt.Println(bstSet.ToString())
	fmt.Println(linkSet.ToString())

	bstSet.Remove("php")
	linkSet.Remove("php")

	fmt.Println(bstSet.ToString())
	fmt.Println(linkSet.ToString())
}

func testBSTree()  {
	bst := tree.NewBST()
	nums := []int{5,3,6,8,4,2}
	for _, n := range nums {
		bst.Add(n)
	}

	bst.LevelOrder(func(e interface{}) {

		fmt.Println(e)

	})

	//bst.PreOrderNR(func(e interface{}) {
	//	fmt.Println(e)
	//})
	//fmt.Println()
	//
	//bst.InOrder(func(e interface{}) {
	//	fmt.Println(e)
	//})
	//fmt.Println()
	//
	//bst.BackOrder(func(e interface{}) {
	//	fmt.Println(e)
	//})
	//
	//fmt.Println(bst.Contains(8))
	//
	//bst.RemoveMax()
	//
	//fmt.Println(bst.Contains(8))
	//fmt.Println(bst.ToString())
	//
	//bst.RemoveMini()
	//fmt.Println(bst.ToString())
	//
	//bst.RemoveMini()
	//fmt.Println(bst.ToString())
	//bst.RemoveMini()
	//fmt.Println(bst.ToString())
}

func sum(arr []int, index int) int {
	if index == len(arr) {
		return 0
	}
	return arr[index] + sum(arr, index+1)
}

func testLinkedListQueue()  {
	linkedlistQueue := linear.NewLinkedListQueue()
	for i := 0; i < 5; i++ {
		linkedlistQueue.EnQueue(i)
	}
	fmt.Println(linkedlistQueue.ToString())
	linkedlistQueue.DeQueue()
	fmt.Println(linkedlistQueue.ToString())
}

func testLinkedListStack()  {
	stack := linear.NewLinkedListStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
		fmt.Println(stack.ToString())
	}

	stack.Pop()
	fmt.Println(stack.ToString())
}

func testLinkedlist()  {
	link := core.NewLinkedList()
	for i := 0; i < 5; i++ {
		link.AddFirst(i)
		fmt.Println(link.ToString())
	}
	link.Insert(3, 999)
	fmt.Println(link.ToString())

	link.Remove(3)
	fmt.Println(link.ToString())

	link.RemoveFirst()
	fmt.Println(link.ToString())

	link.RemoveLast()
	fmt.Println(link.ToString())

}

func testStack(q core.Stack, opCount int) float32 {
	startTime := time.Now().UnixNano()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < opCount; i++ {
		q.Push(r.Intn(opCount))
	}
	for i := 0; i < opCount; i++ {
		q.Pop()
	}
	endTime := time.Now().UnixNano()
	return float32(endTime-startTime)/1000000000.0
}

func testQueue(q core.Queue, opCount int) float32 {
	startTime := time.Now().UnixNano()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < opCount; i++ {
		q.EnQueue(r.Intn(opCount))
	}
	for i := 0; i < opCount; i++ {
		q.DeQueue()
	}
	endTime := time.Now().UnixNano()
	return float32(endTime-startTime)/1000000000.0
}

func testArrayLoopQueue()  {
	loopQueue := linear.NewArrayLoopQueue(10)

	for i := 0; i < 5; i++ {
		loopQueue.EnQueue(i)
	}
	fmt.Println(loopQueue.ToString())
	loopQueue.DeQueue()
	fmt.Println(loopQueue.ToString())
}

func testArrayQueue()  {

	queue := linear.NewArrayQueue(20)

	for i := 0; i < 10; i++ {
		queue.EnQueue(i)
	}
	fmt.Println(queue.ToString())

	queue.DeQueue()
	fmt.Println(queue.ToString())
}

func testArrayStack()  {

	arrStack := linear.NewArrayStack(20)

	for i := 0; i < 10; i++ {
		arrStack.Push(i)
	}

	fmt.Println(arrStack.ToString())

	arrStack.Pop()

	fmt.Println(arrStack.ToString())
	arrStack.Pop()
	arrStack.Pop()
	arrStack.Pop()
	fmt.Println(arrStack.ToString())
}

func testArray() {

	arr := core.NewArray(20)
	for i := 0; i < 10; i++ {
		arr.Append(i)
	}
	fmt.Println(arr.ToString())

	arr.Insert(1, 100)
	fmt.Println(arr.ToString())

	arr.Prepend(99)
	fmt.Println(arr.ToString())

	arr.Append(200)
	fmt.Println(arr.ToString())

	arr.Set(4, 101)
	fmt.Println(arr.ToString())

	fmt.Println("Get Index = 8:", arr.Get(8))

	fmt.Println("Array contains 101 ?", arr.Contains(101))
	fmt.Println("Array contains 201 ?", arr.Contains(201))
	fmt.Println("index of 9 =", arr.Index(9))

	fmt.Println("Delete index 5 =", arr.Delete(5))
	fmt.Println(arr.ToString())

	fmt.Println("Pop: ", arr.Pop())
	fmt.Println(arr.ToString())

	fmt.Println("UnShift: ", arr.UnShift())
	fmt.Println(arr.ToString())

	fmt.Println("Delete Element 1:", arr.DeleteElement(1))
	fmt.Println(arr.ToString())

	fmt.Println()
	fmt.Println("New String Array:")
	arrStr := core.NewArray(10)

	arrStr.Append("php")
	arrStr.Append("java")
	arrStr.Append("golang")
	arrStr.Append("objectc")
	arrStr.Append("python")
	arrStr.Append("ruby")
	arrStr.Append("cpp")
	arrStr.Append("clang")
	arrStr.Append("erlang")
	fmt.Println(arrStr.ToString())

	arrStr.Set(8, "ios")
	fmt.Println(arrStr.ToString())

	fmt.Println("Get Index = 8:", arrStr.Get(8))

	fmt.Println("Array contains php ?", arrStr.Contains("php"))
	fmt.Println("Array contains javascript ?", arrStr.Contains("javascript"))
	fmt.Println("index of python =", arrStr.Index("python"))

	fmt.Println("Delete index 5 =", arrStr.Delete(5))
	fmt.Println(arrStr.ToString())

	fmt.Println("Pop: ", arrStr.Pop())
	fmt.Println(arrStr.ToString())

	fmt.Println("UnShift: ", arrStr.UnShift())
	fmt.Println(arrStr.ToString())

	fmt.Println("Delete Element objectc:", arrStr.DeleteElement("objectc"))
	fmt.Println(arrStr.ToString())

	type Boy struct {
		name string
		age byte
	}

	boyArr := core.NewArray(10)

	kenny := Boy{name: "kenny", age:20}
	kern := Boy{name: "kern", age:22}
	jiekeng := Boy{name: "jiekeng", age:30}
	You := Boy{name: "You", age:29}
	Joe := Boy{name: "Joe", age:30}
	Tim := Boy{name: "Tim", age:34}

	boyArr.Append(kenny)
	boyArr.Append(kern)
	boyArr.Append(jiekeng)
	boyArr.Append(You)
	boyArr.Append(Joe)
	boyArr.Append(Tim)

	fmt.Println(boyArr.ToString())

	boyArr.Set(2, Boy{name: "halo", age: 33})
	fmt.Println(boyArr.ToString())

	fmt.Println("Get Index = 2:", boyArr.Get(2))

	fmt.Println("Array contains keeny ?", boyArr.Contains(kenny))
	fmt.Println("Array contains javascript ?", boyArr.Contains(Boy{name:"fan", age:10}))
	fmt.Println("index of joe =", boyArr.Index(Joe))

	fmt.Println("Delete index 2 =", boyArr.Delete(2))
	fmt.Println(boyArr.ToString())

	fmt.Println("Pop: ", boyArr.Pop())
	fmt.Println(boyArr.ToString())

	fmt.Println("UnShift: ", boyArr.UnShift())
	fmt.Println(boyArr.ToString())

	fmt.Println("Delete Element You:", boyArr.DeleteElement(You))
	fmt.Println(boyArr.ToString())

	arrAuto := core.NewArray(5)

	arrAuto.Append(1)
	arrAuto.Append(2)
	arrAuto.Append(3)
	arrAuto.Append(4)
	arrAuto.Append(5)

	fmt.Println(arrAuto.ToString())
	arrAuto.Append(6)
	arrAuto.Append(7)
	arrAuto.Append(8)
	fmt.Println(arrAuto.ToString())
	arrAuto.Pop()
	arrAuto.Pop()
	arrAuto.Pop()
	arrAuto.Pop()
	arrAuto.Pop()
	fmt.Println(arrAuto.ToString())

	arrAuto.Pop()
	arrAuto.Pop()
	fmt.Println(arrAuto.ToString())
}