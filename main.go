package main

import (
	"data_structure/core"
	"data_structure/linear"
	"fmt"
	"math/rand"
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
	testArrayLoopQueue()

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