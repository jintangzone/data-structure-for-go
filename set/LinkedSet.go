package set

import (
	"data_structure/core"
	"fmt"
	"strings"
)

type LinkedSet struct {
	link *core.LinkedList
}

func (s *LinkedSet) Add(e interface{}) {
	if !s.link.Contains(e) {
		s.link.AddFirst(e)
	}
}

func (s *LinkedSet) Remove(e interface{}) {
	s.link.RemoveElement(e)
}

func (s *LinkedSet) Contains(e interface{}) bool {
	return s.link.Contains(e)
}

func (s *LinkedSet) IsEmpty() bool {
	return s.link.IsEmpty()
}

func (s *LinkedSet) GetSize() int {
	return s.link.GetSize()
}

func (s *LinkedSet) ToString() string {
	strs := make([]string, 0, s.link.GetSize())
	for i := 0; i < s.link.GetSize(); i++ {
		strs = append(strs, fmt.Sprintf("%v", s.link.Get(i)))
	}
	return "Linked Set: [" + strings.Join(strs, ",") + "]"
}

func NewLinkedSet() *LinkedSet {
	s := new(LinkedSet)
	s.link = core.NewLinkedList()
	return s
}
