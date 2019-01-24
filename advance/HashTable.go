package advance

import (
	"data_structure/core"
	"fmt"
)

var UpperTol = 10
var LowerTol = 2
var InitCapacity = 7

type HashTable struct {
	hashTable []*AVLMap
	m int
	size int
}

func (ht *HashTable) hash(key core.Key) int {
	return (key.HashCode() & 0x7fffffff) % ht.m
}

func (ht *HashTable) Add(key core.Key, v interface{}) {
	avlMap := ht.hashTable[ht.hash(key)]
	if avlMap.Contains(key) {
		avlMap.Add(key, v)
	} else {
		avlMap.Add(key, v)
		ht.size++
		if ht.size >= UpperTol*ht.m {
			ht.resize(2*ht.m)
		}
	}
}

func (ht *HashTable) Remove(key core.Key) interface{} {
	avlMap := ht.hashTable[ht.hash(key)]
	var ret interface{}
	if avlMap.Contains(key) {
		ret = avlMap.Remove(key)
		ht.size--
		if ht.size < UpperTol*ht.m && ht.m / 2 >= InitCapacity {
			ht.resize(ht.m/2)
		}
	}
	return ret
}

func (ht *HashTable) resize(newM int) {
	newHashTable := make([]*AVLMap, newM)
	for i := 0; i < newM; i++ {
		newHashTable[i] = NewAVLMap()
	}
	ht.m = newM
	for i := 0; i < ht.m; i++ {
		avlMap := newHashTable[i]
		for _, key := range avlMap.KeySet() {
			newHashTable[ht.hash(key)].Add(key, avlMap.Get(key))
		}
	}
	ht.hashTable = newHashTable
}

func (ht *HashTable) Set(key core.Key, v interface{})  {
	avlMap := ht.hashTable[ht.hash(key)]
	if !avlMap.Contains(key) {
		panic(fmt.Sprint(key) + "doesn't exists!")
	}
	avlMap.Set(key, v)
}

func (ht *HashTable) Get(key core.Key) interface{} {
	return ht.hashTable[ht.hash(key)].Get(key)
}

func NewHasTable(m int) *HashTable {
	return &HashTable{
		hashTable: make([]*AVLMap, m),
		m: m,
		size: 0,
	}
}