package heap

// Heap is a struct that offers min or max heap data strucutre implementation using given comparator function.
type Heap[T any] struct {
	values     []T
	comparator ComparatorFn[T]
}

// NewHeap creates a new heap using given comparator function.
func NewHeap[T any](comparator ComparatorFn[T]) *Heap[T] {
	return &Heap[T]{
		comparator: comparator,
	}
}

func (h *Heap[T]) Insert(val T) *Heap[T] {
	h.values = append(h.values, val)

	h.heapifyUp()

	return h
}

func (h *Heap[T]) Remove() T {
	if h.Length() == 0 {
		var val T
		return val
	}

	val := h.values[0]
	h.swap(0, h.Length()-1)
	h.values = h.values[:h.Length()-1]

	h.heapifyDown()

	return val

}

func (h *Heap[T]) heapifyUp() {
	var cur, parent int
	setCur := func(c int) {
		cur = c
		parent = h.getParentIndex(c)
	}

	setCur(len(h.values) - 1)

	for cur > 0 && h.comparator(h.values[cur], h.values[parent]) >= 1 {
		h.swap(parent, cur)
		setCur(parent)
	}

}

func (h *Heap[T]) heapifyDown() {
	cur, left, right, replace := 0, 0, 0, 0
	setCur := func(i int) {
		cur = i
		replace = i
		left = h.getLeftChild(i)
		right = h.getRightChild(i)
	}

	setCur(0)

	for right < h.Length() || left < h.Length() {

		if right < h.Length() && h.comparator(h.values[right], h.values[cur]) >= 1 {
			replace = right
		}

		if left < h.Length() && h.comparator(h.values[left], h.values[replace]) >= 1 {
			replace = left
		}

		if cur == replace {
			break
		}

		h.swap(cur, replace)
		setCur(replace)
	}

}

func (h *Heap[T]) Length() int {
	return len(h.values)
}

func (h *Heap[T]) getParentIndex(index int) int {
	return (index - 1) / 2
}

func (h *Heap[T]) getLeftChild(index int) int {
	return 2*index + 1
}

func (h *Heap[T]) getRightChild(index int) int {
	return 2*index + 2
}

func (h *Heap[T]) swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *Heap[T]) Peek() T {
	if len(h.values) == 0 {
		var noop T
		return noop
	}

	return h.values[0]
}
