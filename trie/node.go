package trie

type Node[T comparable] struct {
	val         T
	isEndOfWord bool
	parent      *Node[T]
	children    map[T]*Node[T]
}

func NewNode[T comparable](val T) *Node[T] {
	return &Node[T]{val: val}
}
