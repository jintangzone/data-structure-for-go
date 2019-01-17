package core

type Dict interface {
	Add(k interface{}, v interface{})
	Remove(k interface{}) interface{}
	Contains(k interface{}) bool
	Get(k interface{}) interface{}
	Set(k interface{}, e interface{})
	GetSize() int
	IsEmpty() bool
}
