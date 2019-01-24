package core

type Map interface {
	Add(k Key, v interface{})
	Remove(k Key) interface{}
	Contains(k Key) bool
	KeySet() []Key
	Get(k Key) interface{}
	Set(k Key, v interface{})
	GetSize() int
	IsEmpty() bool
}