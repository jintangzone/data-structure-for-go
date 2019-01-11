package linear

import (
	"fmt"
	"strings"
)

type ArrayLoopQueue struct {
	data []interface{}
	front int
	tail int
	size int
}

func (q *ArrayLoopQueue) EnQueue(e interface{}) {

	if q.IsFull() {
		q.Resize(q.GetCapacity() * 2)
	}
	q.data[q.tail] = e
	q.tail = (q.tail+1) % len(q.data)
	q.size++
}

func (q *ArrayLoopQueue) DeQueue() interface{} {
	if q.IsEmpty() {
		panic("Cannot dequeue from an empty queue.")
	}
	ret := q.data[q.front]
	q.data[q.front] = nil
	q.front = (q.front + 1) % len(q.data)
	q.size--
	if q.size == (q.GetCapacity() / 4) {
		q.Resize(q.GetCapacity()/2)
	}
	return ret
}

func (q *ArrayLoopQueue) GetFront() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.data[q.front]
}

func (q *ArrayLoopQueue) GetNear() interface{} {
	if q.IsEmpty() {
		return nil
	}
	if q.tail == 0 {
		return q.data[len(q.data)-1]
	} else {
		return q.data[q.tail-1]
	}
}

func (q *ArrayLoopQueue) GetSize() int {
	return q.size
}

func (q *ArrayLoopQueue) GetCapacity() int {
	return len(q.data)-1
}

func (q *ArrayLoopQueue) IsEmpty() bool {
	return q.front == q.tail
}

func (q *ArrayLoopQueue) IsFull() bool {
	return (q.tail+1)%len(q.data) == q.front
}

func (q *ArrayLoopQueue) Resize(capacity int) {
	newData := make([]interface{}, capacity+1)
	for i := 0; i < q.GetSize(); i++ {
		newData[i] = q.data[(q.front+i)%len(q.data)]
	}
	q.data = newData
	q.front = 0
	q.tail = q.size
}


func (q *ArrayLoopQueue) ToString() string {
	arrInStr := fmt.Sprintf("Queue: size = %d, capacity = %d front [", q.GetSize(), q.GetCapacity())
	if q.GetSize() > 0 {
		eInArr := make([]string, 0)
		for i := q.front; i != q.tail; i = (i + 1)%len(q.data) {
			eInArr = append(eInArr, fmt.Sprintf("%v", q.data[i]))
		}

		arrInStr = arrInStr + strings.Join(eInArr, ", ")
	}

	arrInStr = arrInStr + "] tail"
	return arrInStr
}

func NewArrayLoopQueue(capacity int) *ArrayLoopQueue {
	q := new(ArrayLoopQueue)
	q.data = make([]interface{}, capacity+1)
	return q
}
