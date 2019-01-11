package linear

import (
	"data_structure/core"
	"fmt"
	"strings"
)

type ArrayStack struct {
	array *core.Array
}

func (as *ArrayStack) Push(e interface{}) {
	as.array.Append(e)
}

func (as *ArrayStack) Pop() interface{} {
	return as.array.Pop()
}

func (as *ArrayStack) Peek() interface{} {
	return as.array.GetLast()
}

func (as *ArrayStack) GetSize() int {
	return as.array.GetSize()
}

func (as *ArrayStack) IsEmpty() bool {
	return as.array.IsEmpty()
}

func (as *ArrayStack) ToString() string {
	arrInStr := fmt.Sprintf("Stack: size = %d, capacity = %d [", as.GetSize(), as.array.GetCapacity())
	if as.GetSize() > 0 {
		eInArr := make([]string, as.GetSize())
		for i := 0; i < as.GetSize(); i++ {
			eInArr[i] = fmt.Sprintf("%v", as.array.Get(i))
		}
		arrInStr = arrInStr + strings.Join(eInArr, ",")
	}
	arrInStr = arrInStr + "] top"
	return arrInStr
}

func NewArrayStack(capacity int) *ArrayStack {
	as := new(ArrayStack)
	as.array = core.NewArray(capacity)
	return as
}
