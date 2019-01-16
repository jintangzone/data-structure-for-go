package set

import (
	"data_structure/tree"
	"fmt"
	"strings"
)

type BSTSet struct {
	tree *tree.BST
}

func (s *BSTSet) Add(e interface{}) {
	s.tree.Add(e)
}

func (s *BSTSet) Remove(e interface{}) {
	s.tree.Remove(e)
}

func (s *BSTSet) Contains(e interface{}) bool {
	return s.tree.Contains(e)
}

func (s *BSTSet) IsEmpty() bool {
	return s.tree.IsEmpty()
}

func (s *BSTSet) GetSize() int {
	return s.tree.GetSize()
}

func (s *BSTSet) ToString() string {
	strs := make([]string, 0, s.tree.GetSize())
	s.tree.PreOrder(func(e interface{}) {
		if e != nil {
			strs = append(strs, fmt.Sprintf("%v", e))
		}
	})
	return "BST Set: [" + strings.Join(strs, ",") + "]"
}

func NewBSTSet() *BSTSet {
	s := new(BSTSet)
	s.tree = tree.NewBST()
	return s
}
