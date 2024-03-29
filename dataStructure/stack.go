package main

import (
	"container/list"
)

type Stack struct {
	v *list.List
}

func (s *Stack) Push(val interface{}) {
	s.v.PushBack(val)
}

func (s *Stack) Pop() interface{} {
	back := s.v.Back()
	if back != nil {
		return s.v.Remove(back)
	}
	return nil
}

func NewStack() *Stack {
	return &Stack{list.New()}
}
