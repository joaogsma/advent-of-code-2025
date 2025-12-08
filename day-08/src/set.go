// AI generated because Go should have this ò_ó
package main

import "iter"

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable](values ...T) *Set[T] {
	s := &Set[T]{m: make(map[T]struct{}, len(values))}
	for _, v := range values {
		s.m[v] = struct{}{}
	}
	return s
}

func (s *Set[T]) Add(v T) {
	s.m[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.m, v)
}

func (s *Set[T]) Contains(v T) bool {
	_, ok := s.m[v]
	return ok
}

func (s *Set[T]) Len() int {
	return len(s.m)
}

func (s *Set[T]) Clear() {
	clear(s.m)
}

func (s *Set[T]) IsSubsetOf(other *Set[T]) bool {
	for v := range s.m {
		if !other.Contains(v) {
			return false
		}
	}
	return true
}

func (s *Set[T]) Values() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		var i int = 0
		for v := range s.m {
			if !yield(i, v) {
				return
			}
			i++
		}
	}
}

func (s *Set[T]) ToSlice() []T {
	out := make([]T, 0, len(s.m))
	for v := range s.m {
		out = append(out, v)
	}
	return out
}

func (s *Set[T]) Equals(other *Set[T]) bool {
	if s.Len() != other.Len() {
		return false
	}
	return s.IsSubsetOf(other)
}

func (a *Set[T]) Union(b *Set[T]) *Set[T] {
	out := NewSet[T]()
	for v := range a.m {
		out.m[v] = struct{}{}
	}
	for v := range b.m {
		out.m[v] = struct{}{}
	}
	return out
}

func (a *Set[T]) Intersect(b *Set[T]) *Set[T] {
	out := NewSet[T]()

	// iterate smaller set for performance
	if a.Len() > b.Len() {
		a, b = b, a
	}

	for v := range a.m {
		if b.Contains(v) {
			out.m[v] = struct{}{}
		}
	}
	return out
}

func (a *Set[T]) Difference(b *Set[T]) *Set[T] {
	out := NewSet[T]()
	for v := range a.m {
		if !b.Contains(v) {
			out.m[v] = struct{}{}
		}
	}
	return out
}
