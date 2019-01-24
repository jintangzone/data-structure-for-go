package advance

import "data_structure/core"

type AVLMap struct {
	avl *AVLTree
}

func (m *AVLMap) Add(k core.Key, v interface{}) {
	m.avl.Add(k, v)
}

func (m *AVLMap) Remove(k core.Key) interface{} {
	return m.avl.Remove(k)
}

func (m *AVLMap) Contains(k core.Key) bool {
	return m.avl.Contain(k)
}

func (m *AVLMap) Get(k core.Key) interface{} {
	return m.avl.Get(k)
}

func (m *AVLMap) Set(k core.Key, v interface{}) {
	m.avl.Set(k, v)
}

func (m *AVLMap) KeySet() []core.Key {
	keySets := make([]core.Key, 0)
	m.avl.LevelOrder(func(k core.Key, v interface{}) {
		keySets = append(keySets, k)
	})
	return keySets
}

func (m *AVLMap) GetSize() int {
	return m.GetSize()
}

func (m *AVLMap) IsEmpty() bool {
	return m.IsEmpty()
}

func NewAVLMap() *AVLMap {
	return &AVLMap{ avl: NewAVLTree() }
}

