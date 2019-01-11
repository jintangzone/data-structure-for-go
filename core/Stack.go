package core

type Stack interface {
	Push(e interface{})
	Pop() interface{}
	Peek() interface{}
	GetSize() int
	IsEmpty() bool
}