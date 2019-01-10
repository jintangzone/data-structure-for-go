package core

import (
	"fmt"
)

/*

Time Complexity analysis

insert，delete: O(n)

PS: append operation is O(1)
but if need to resize, copy array must traversal element of arr, so time complexity is O(n)

set, get: O(1), unknow index: O(n)

总结：

添加操作：O(n)

*/

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

// Time Complexity: O(n)
func (arr *Array) Prepend(e interface{}) {
	arr.Insert(0, e)
}

// Time Complexity: O(1)
func (arr *Array) Append(e interface{}) {
	arr.Insert(arr.size, e)
}

// Time Complexity: O(n)
// 最高复杂度是O(n)，则以最高复杂度为准
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

// TIme Complexity: O(1)
func (arr *Array) Get(index int) interface{} {
	if index < 0 || index > arr.size {
		panic("Get Failed, Index is illegal.")
	}
	return arr.data[index]
}

// TIme Complexity: O(1)
func (arr *Array) Set(index int, e interface{}) {
	if index < 0 || index > arr.size {
		panic("Get Failed, Index is illegal.")
	}
	arr.data[index] = e
}

// TIme Complexity: O(n)
func (arr *Array) Contains(e interface{}) bool {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return true
		}
	}
	return false
}

// TIme Complexity: O(n)
func (arr *Array) Index(e interface{}) int {
	for i := 0; i < arr.size; i++ {
		if arr.data[i] == e {
			return i
		}
	}
	return -1
}

// TIme Complexity: O(n)
// 最高复杂度是O(n)，则以最高复杂度为准
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

	if arr.size == cap(arr.data)/4 && cap(arr.data)/2 != 0 {
		arr.Resize(cap(arr.data)/2)
	}
	return ret
}

// TIme Complexity: O(1)
func (arr *Array) Pop() interface{} {
	return arr.Delete(arr.size-1)
}

// TIme Complexity: O(n)
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

// TIme Complexity: O(n)
func (arr *Array) Resize(capacity int) {
	newData := make([]interface{}, capacity)
	copy(newData, arr.data)
	arr.data = newData
}

func (arr *Array) GetLast() interface{}  {
	return arr.Get(arr.size-1)
}

func (arr *Array) GetFirst() interface{}  {
	return arr.Get(0)
}
