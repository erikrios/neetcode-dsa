package queue

import (
	"errors"
	"linked-list/linkedlist"
	"linked-list/linkedlist/doubly"
)

var ErrEmptyQueue = errors.New("The queue is empty.")

type Queue[T any] struct {
	ll linkedlist.LinkedList[T]
}

func New[T any]() *Queue[T] {
	ll := doubly.New[T]()
	return &Queue[T]{ll: ll}
}

func (q *Queue[T]) Enqueue(v T) {
	q.ll.InsertAt(0, v)
}

func (q *Queue[T]) Dequeue() (result T, err error) {
	if q.IsEmpty() {
		err = ErrEmptyQueue
		return
	}

	result = q.ll.Delete()
	return
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.ToList()) == 0
}

func (q *Queue[T]) ToList() []T {
	listRev := q.ll.ToList()

	if len(listRev) == 0 {
		return []T{}
	}

	results := make([]T, len(listRev))

	for i := 0; i <= len(listRev)/2; i++ {
		results[i], results[len(listRev)-i-1] = listRev[len(listRev)-i-1], listRev[i]
	}

	return results
}
