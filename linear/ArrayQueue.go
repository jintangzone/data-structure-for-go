package linear

import (
	"data_structure/linear/core"
	"fmt"
	"strings"
)

type ArrayQueue struct {
	array *core.Array
}

func (q *ArrayQueue) EnQueue(e interface{}) {
	q.array.Append(e)
}

func (q *ArrayQueue) DeQueue() interface{} {
	return q.array.UnShift()
}

func (q *ArrayQueue) GetFront() interface{} {
	return q.array.GetFirst()
}

func (q *ArrayQueue) GetSize() int {
	return q.array.GetSize()
}

func (q *ArrayQueue) GetCapacity() int {
	return q.array.GetCapacity()
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.array.IsEmpty()
}

func (q *ArrayQueue) ToString() string {
	arrInStr := fmt.Sprintf("Queue: size = %d, capacity = %d front [", q.GetSize(), q.array.GetCapacity())
	if q.GetSize() > 0 {
		eInArr := make([]string, q.GetSize())
		for i := 0; i < q.GetSize(); i++ {
			eInArr[i] = fmt.Sprintf("%v", q.array.Get(i))
		}
		arrInStr = arrInStr + strings.Join(eInArr, ",")
	}
	arrInStr = arrInStr + "] tail"
	return arrInStr
}

func NewArrayQueue(capacity int) *ArrayQueue {
	q := new(ArrayQueue)
	q.array = core.NewArray(capacity)
	return q
}
