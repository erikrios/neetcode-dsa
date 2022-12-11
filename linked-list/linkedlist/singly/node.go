package singly

type node[T any] struct {
	val  T
	next *node[T]
}

func NewNode[T any](val T, next *node[T]) *node[T] {
	return &node[T]{
		val:  val,
		next: next,
	}
}
