package Array

type Array struct {
	data []int
	size int
}

func NewArray(capacity int) *Array {
	arr := new(Array)
	arr.size = 0
	arr.data = make([]int, 0, capacity)
	return arr
}

func (arr *Array) GetSize() int {
	return arr.size
}

func (arr *Array) GetCapacity() int {
	return cap(arr.data)
}

func (arr *Array) IsEmpty() bool {
	return arr.size == 0
}

