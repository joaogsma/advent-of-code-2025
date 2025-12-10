package main

import "slices"

type Stack[T any] struct {
	Data []T
}

func EmptyStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(elem T) {
	s.Data = append(s.Data, elem)
}

func (s *Stack[T]) Pop() T {
	elem := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return elem
}

func (s *Stack[T]) Len() int {
	return len(s.Data)
}

func (s *Stack[T]) Peek(i int) *T {
	return &s.Data[len(s.Data)-1-i]
}

func (s *Stack[T]) ToSlice() []T {
	return slices.Clone(s.Data)
}
