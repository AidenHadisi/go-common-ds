package trie

type Node[T comparable] struct {
	val         T
	isEndOfWord bool
	children    map[T]*Node[T]
	parent      *Node[T]
}

func NewNode[T comparable](val T) *Node[T] {
	return &Node[T]{
		val: val,
	}
}
