package stack

import "errors"

var (
	StackEmptyErr = errors.New("stack is empty")
)

type Stack[T any] struct {
	items []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{items: make([]T, 0)}
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (v T, err error) {
	if s.isEmpty() {
		err = StackEmptyErr
		return
	}

	v = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return
}

func (s *Stack[T]) Peek() (v T, err error) {
	if s.isEmpty() {
		err = StackEmptyErr
		return
	}

	v = s.items[len(s.items)-1]
	return
}

func (s *Stack[T]) isEmpty() bool {
	return len(s.items) == 0
}
