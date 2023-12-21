package datastructures

import "errors"

type LinkedStack struct {
	top    *Node
	length int
}

func (s *LinkedStack) GetLength() int {
	return s.length
}

func (s *LinkedStack) Stack(value any) error {
	s.top = &Node{value: value, next: s.top}
	s.length++
	return nil
}

func (s *LinkedStack) Unstack() (any, error) {
	if s.length == 0 {
		return nil, errors.New("Can't unstack empty stack!")
	}

	node := s.top

	s.top = node.next
	s.length--

	return node.value, nil
}

func (s *LinkedStack) Read() (any, error) {
	if s.length == 0 {
		return nil, errors.New("Can't read from empty stack!")
	}
	return s.top.value, nil
}
