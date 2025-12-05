package main

import (
	"fmt"
	"strconv"

	"golang.org/x/exp/constraints"
)

type Range[T constraints.Integer | constraints.Float] struct {
	Begin T
	End   T
}

func (r Range[T]) Contains(value T) bool {
	return r.Begin <= value && value < r.End
}

func (r Range[T]) ContainsRange(other Range[T]) bool {
	return r.Begin <= other.Begin && other.End <= r.End
}

func (r Range[T]) IsEmpty() bool {
	return r.Begin >= r.End
}

func (r Range[T]) Intersects(other Range[T]) bool {
	if r.IsEmpty() || other.IsEmpty() {
		return false
	}
	if r.ContainsRange(other) || other.ContainsRange(r) {
		return true
	}
	return r.Contains(other.Begin) || r.Contains(other.End-1)
}

func (r Range[T]) Union(other Range[T]) Range[T] {
	if !r.Intersects(other) {
		panic("Disjoint ranges")
	}
	return Range[T]{min(r.Begin, other.Begin), max(r.End, other.End)}
}

func (r Range[T]) Intersect(other Range[T]) Range[T] {
	if !r.Intersects(other) {
		var zero T
		return Range[T]{zero, zero}
	}
	begin := max(r.Begin, other.Begin)
	end := min(r.End, other.End)
	return Range[T]{begin, end}
}

func (r Range[T]) String() string {
	return fmt.Sprintf("%s-%s", strconv.FormatUint(uint64(r.Begin), 10), strconv.FormatUint(uint64(r.End), 10))
}

func (r Range[T]) IsLeftOf(partition T) bool {
	return r.Begin < partition
}

func (r Range[T]) IsRightOf(partition T) bool {
	return r.End > partition
}
