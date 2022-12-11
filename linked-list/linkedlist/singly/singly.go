package singly

import (
	"reflect"
)

type SinglyLinkedList[T any] struct {
	head *node[T]
	tail *node[T]
}

func New[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{
		head: nil,
		tail: nil,
	}
}

func (s *SinglyLinkedList[T]) Insert(v T) {
	newNode := NewNode(v, nil)

	if s.head == nil {
		s.head = newNode
		s.tail = newNode
		return
	}

	s.tail.next = newNode
	s.tail = newNode
}

func (s *SinglyLinkedList[T]) InsertAt(index int, v T) {
	newNode := NewNode(v, nil)

	if index == 0 {
		if s.head == nil {
			s.head = newNode
			s.tail = newNode
			return
		} else {
			newNode.next = s.head
			s.head = newNode
			return
		}
	}

	if s.head == nil {
		s.head = newNode
		s.tail = newNode
		return
	}

	var prev *node[T]
	cur := s.head
	for i := 0; i < index; i++ {
		if cur.next != nil {
			prev = cur
			cur = cur.next
		} else {
			cur.next = newNode
			s.tail = newNode
			return
		}
	}

	newNode.next = cur
	prev.next = newNode
}

func (s *SinglyLinkedList[T]) Find(v T) bool {
	var isExists bool

	cur := s.head
	for cur != nil {
		if reflect.DeepEqual(cur.val, v) {
			isExists = true
			break
		}
		cur = cur.next
	}

	return isExists
}

func (s *SinglyLinkedList[T]) Delete() {
	if s.head == nil {
		return
	}

	if s.tail == s.head {
		s.head, s.tail = nil, nil
		return
	}

	var prev *node[T]
	cur := s.head
	for cur.next != nil {
		prev = cur
		cur = cur.next
	}

	s.tail = prev
	s.tail.next = nil
}

func (s *SinglyLinkedList[T]) ToList() []T {
	results := make([]T, 0)

	cur := s.head
	for cur != nil {
		results = append(results, cur.val)
		cur = cur.next
	}

	return results
}
