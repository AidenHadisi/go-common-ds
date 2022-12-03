package heap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var comparator = func(a, b int) int {
	if a < b {
		return 1
	}

	if b < a {
		return -1
	}

	return 0
}

func Test_Insert(t *testing.T) {
	cases := []struct {
		Val      int
		Expected []int
	}{
		{2, []int{2}},
		{3, []int{2, 3}},
		{1, []int{1, 3, 2}},
		{0, []int{0, 1, 2, 3}},
		{-3, []int{-3, 0, 2, 3, 1}},
	}

	heap := NewHeap(comparator)

	for _, test := range cases {
		t.Run(fmt.Sprintf("test insert %d", test.Val), func(t *testing.T) {
			heap.Insert(test.Val)
			assert.Equal(t, test.Expected, heap.values)
			assert.Equal(t, len(test.Expected), heap.Length())
		})

	}
}

func Test_Remove(t *testing.T) {
	cases := []int{-3, 0, 1, 2, 3, 0}

	heap := NewHeap(comparator)
	heap.values = []int{-3, 0, 2, 3, 1}

	for _, expected := range cases {
		t.Run(fmt.Sprintf("test pop %d", expected), func(t *testing.T) {
			val := heap.Remove()
			assert.Equal(t, expected, val)
		})

	}
}

func Test_Peek(t *testing.T) {
	expected := 5

	heap := NewHeap(comparator)
	heap.values = []int{-3, 0, 2, 3, 1}

	assert.Equal(t, expected, heap.Length())

}
