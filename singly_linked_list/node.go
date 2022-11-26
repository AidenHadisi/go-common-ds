package singlylinkedlist

type Node[T any] struct {
	val  T
	next *Node[T]
}

func NewNode[T any](val T) *Node[T] {
	return &Node[T]{val: val}
}
