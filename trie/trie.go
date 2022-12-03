package trie

type Trie[T comparable] struct {
	root  *Node[T]
	words int
}

func NewTrie[T comparable]() *Trie[T] {
	return &Trie[T]{
		root: &Node[T]{},
	}
}

func (t *Trie[T]) Insert(word []T) *Trie[T] {
	cur := t.root
	for _, w := range word {
		if _, ok := cur.children[w]; !ok {
			cur.children[w] = NewNode(w)
			cur.children[w].parent = cur
		}

		cur = cur.children[w]
	}

	cur.isEndOfWord = true
	t.words++

	return t
}

func (t *Trie[T]) WordExists(word []T) bool {
	cur := t.root

	for _, w := range word {
		if _, ok := cur.children[w]; !ok {
			return false
		}

		cur = cur.children[w]
	}

	return cur.isEndOfWord
}

func (t *Trie[T]) PrefixExists(prefix []T) bool {
	cur := t.root

	for _, w := range prefix {
		if _, ok := cur.children[w]; !ok {
			return false
		}

		cur = cur.children[w]
	}
	return true

}

func (t *Trie[T]) WordCount() int {
	return t.words
}
