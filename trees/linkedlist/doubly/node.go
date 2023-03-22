package doubly

type node[T any] struct {
	val  T
	prev *node[T]
	next *node[T]
}

func NewNode[T any](val T, next, prev *node[T]) *node[T] {
	return &node[T]{
		val:  val,
		next: next,
		prev: prev,
	}
}
