package trie

type Trie[T comparable] struct {
	root *Node[T]
}

func (t *Trie[T]) Inert(data []T) *Trie[T] {
	curNode := t.root
	for _, d := range data {
		if _, ok := curNode.children[d]; !ok {
			curNode.children[d] = NewNode(d)
			curNode.children[d].parent = curNode
		}

		curNode = curNode.children[d]
	}

	curNode.isEndOfWord = true

	return t
}

func (t *Trie[T]) Exists(data []T) bool {
	curNode := t.root

	for _, d := range data {
		if _, ok := curNode.children[d]; !ok {
			return false
		}

		curNode = curNode.children[d]
	}

	return curNode.isEndOfWord
}
