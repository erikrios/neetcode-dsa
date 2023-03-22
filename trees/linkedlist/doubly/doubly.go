package doubly

import "reflect"

type DoublyLinkedList[T any] struct {
	head *node[T]
	tail *node[T]
}

func New[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		head: nil,
		tail: nil,
	}
}

func (d *DoublyLinkedList[T]) Insert(v T) {
	newNode := NewNode(v, nil, nil)

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	d.tail.next = newNode
	newNode.prev = d.tail
	d.tail = newNode
}

func (d *DoublyLinkedList[T]) InsertAt(index int, v T) {
	newNode := NewNode(v, nil, nil)

	if index == 0 {
		if d.head == nil {
			d.head = newNode
			d.tail = newNode
			return
		} else {
			newNode.next = d.head
			d.head.prev = newNode
			d.head = newNode
			return
		}
	}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	cur := d.head
	for i := 0; i < index; i++ {
		if cur.next != nil {
			cur = cur.next
		} else {
			cur.next = newNode
			newNode.prev = cur
			d.tail = newNode
			return
		}
	}

	newNode.next = cur
	newNode.prev = cur.prev
	cur.prev.next = newNode
	cur.prev = newNode
}

var isExists bool

func (d *DoublyLinkedList[T]) Find(v T) bool {

	cur := d.head
	for cur != nil {
		if reflect.DeepEqual(cur.val, v) {
			isExists = true
			break
		}
		cur = cur.next
	}

	return isExists
}

func (d *DoublyLinkedList[T]) Delete() (result T) {
	if d.head == nil {
		return
	}

	if d.head == d.tail {
		result = d.head.val
		d.head, d.tail = nil, nil
		return
	}

	result = d.tail.val
	oldTail := d.tail
	d.tail.prev.next = nil
	d.tail = d.tail.prev
	oldTail.prev = nil
	return
}

func (d *DoublyLinkedList[T]) ToList() []T {
	results := make([]T, 0)

	cur := d.head
	for cur != nil {
		results = append(results, cur.val)
		cur = cur.next
	}

	return results
}
