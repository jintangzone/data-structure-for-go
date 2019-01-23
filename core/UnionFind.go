package core

type UnionFind interface {
	IsConnected(p, q int) bool
	Union(p, q int)
	GetSize() int
}