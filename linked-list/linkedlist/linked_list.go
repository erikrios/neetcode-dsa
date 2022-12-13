package linkedlist

type LinkedList[T any] interface {
	Insert(v T)
	InsertAt(index int, v T)
	Find(v T)
	Delete()
	ToList() []T
}
