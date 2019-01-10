package core

type Queue interface {
	EnQueue(e interface{})
	DeQueue() interface{}
	GetFront() interface{}
	GetSize() int
	IsEmpty() bool
}