package main

type Queue[T any] struct {
	data []T
	head int
}

func EmptyQueue[T any]() Queue[T] {
	return Queue[T]{[]T{}, 0}
}

func FilledQueue[T any](data []T) Queue[T] {
	return Queue[T]{data, 0}
}

func (q *Queue[T]) Push(v T) {
	q.data = append(q.data, v)
}

func (q *Queue[T]) Pop() (T, bool) {
	if q.head >= len(q.data) {
		var zero T
		return zero, false
	}

	v := q.data[q.head]
	q.head++

	// Periodic compaction to avoid memory growth
	if q.head > len(q.data)/2 {
		q.data = append([]T(nil), q.data[q.head:]...)
		q.head = 0
	}

	return v, true
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head >= len(q.data)
}
