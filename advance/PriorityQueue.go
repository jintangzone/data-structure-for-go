package advance

import "data_structure/heap"

type PriorityQueue struct {
	mah *heap.MaxArrayHeap
}


func (pq *PriorityQueue) EnQueue(e interface{}) {
	pq.mah.Add(e.(int))
}

func (pq *PriorityQueue) DeQueue() interface{} {
	return pq.mah.ExtractMax()
}

func (pq *PriorityQueue) GetFront() interface{} {
	return pq.mah.FindMax()
}

func (pq *PriorityQueue) GetSize() int {
	return pq.mah.GetSize()
}

func (pq *PriorityQueue) IsEmpty() bool {
	return pq.mah.IsEmpty()
}
