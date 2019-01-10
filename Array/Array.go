package Array

import (
	"fmt"
)

type Array struct {
	data []interface{}
	size int
}

func NewArray(capacity int) *Array {
	arr := new(Array)
	arr.size = 0
	arr.data = make([]interface{}, capacity)
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

func (arr *Array) Prepend(e interface{}) {
	arr.Insert(0, e)
}

func (arr *Array) Append(e interface{}) {
	arr.Insert(arr.size, e)
}

func (arr *Array) Insert(index int, e interface{})  {
	if arr.size == cap(arr.data) {
		arr.Resize(arr.size*2)
	}
	if index < 0 || index > arr.size {
		panic("Insert Failed, Out of range Index.")
	}

	for i := arr.size-1; i >= index; i-- {
		arr.data[i+1] = arr.data[i]
	}
	arr.data[index] = e
	arr.size++
}

func (arr *Array) ToString() string {
	arrInStr := fmt.Sprintf("Array: size = %d, capacity = %d ", arr.size, cap(arr.data))
	arrInStr = arrInStr + fmt.Sprintf("%v", arr.data[:arr.size])
	return arrInStr
}

func (arr *Array) Get(index int) interface{} {
	if index < 0 || index > arr.size {
		panic("Get Failed, Index is illegal.")
	}
	return arr.data[index]
}

func (arr *Array) Set(index int, e interface{}) {
	if index < 0 || index > arr.size {
		panic("Get Failed, Index is illegal.")
	}
	arr.data[index] = e
}

func (arr *Array) Contains(e interface{}) bool {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return true
		}
	}
	return false
}

func (arr *Array) Index(e interface{}) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return i
		}
	}
	return -1
}

func (arr *Array) Delete(index int) interface{} {
	if index < 0 || index >= arr.size {
		panic("Delete Failed, Index is illegal.")
	}
	ret := arr.data[index]
	for i := index + 1; i < arr.size; i++ {
		arr.data[i-1] = arr.data[i]
	}
	arr.size--
	arr.data[arr.size] = nil

	if arr.size == cap(arr.data)/2 {
		arr.Resize(cap(arr.data)/2)
	}
	return ret
}

func (arr *Array) Pop() interface{} {
	return arr.Delete(arr.size-1)
}

func (arr *Array) UnShift() interface{} {
	return arr.Delete(0)
}

func (arr *Array) DeleteElement(e interface{}) bool {
	index := arr.Index(e)
	if index > 0 {
		arr.Delete(index)
		return true
	}
	return false
}

func (arr *Array) Resize(capacity int) {
	newData := make([]interface{}, capacity)
	copy(newData, arr.data)
	arr.data = newData
}