package core

type Map interface {
	Add(k Key, v interface{})
	Remove(k Key)
	Contains(k Key) bool
	Get(k Key) interface{}
	Set(k Key, v interface{})
	GetSize() int
	IsEmpty() bool
}