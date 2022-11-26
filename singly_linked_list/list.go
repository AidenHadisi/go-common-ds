package singlylinkedlist

type List[T any] struct {
	head, tail *Node[T]
	length     int
}

func NewList[T any]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Push(val T) *List[T] {
	n := NewNode(val)

	if l.tail != nil {
		l.tail.next = n
	} else {
		l.head = n
	}

	l.tail = n
	l.length++

	return l
}

func (l *List[T]) Unshift(val T) *List[T] {
	n := NewNode(val)

	n.next = l.head

	if l.tail == nil {
		l.tail = n
	}

	l.head = n
	l.length++

	return l
}

func (l *List[T]) InsertAtIndex(index int, val T) *List[T] {
	curNode := &Node[T]{next: l.head}
	dummy := curNode

	n := NewNode(val)

	for index > 0 && curNode.next != nil {
		curNode = curNode.next
		index--
	}

	n.next = curNode.next
	curNode.next = n

	l.head = dummy.next
	l.length++

	return l
}

func (l *List[T]) Length() int {
	return l.length
}

func (l *List[T]) ToSlice() []T {
	slice := make([]T, 0)
	curNode := l.head

	for curNode != nil {
		slice = append(slice, curNode.val)
		curNode = curNode.next
	}

	return slice
}

func (l *List[T]) Traverse(cb func(index int, val T)) {
	curNode := l.head
	index := 0

	for curNode != nil {
		cb(index, curNode.val)
		index++
		curNode = curNode.next
	}
}

func (l *List[T]) Clear() *List[T] {
	l.head = nil
	l.tail = nil
	l.length = 0

	return l
}
