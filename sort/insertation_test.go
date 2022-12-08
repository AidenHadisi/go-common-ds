package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertion(t *testing.T) {
	var comparator ComparatorFn[int] = func(a, b int) int {
		return a - b
	}

	var cases = []struct {
		name     string
		input    []int
		expected []int
	}{
		{"empty array", []int{}, []int{}},
		{"1 element", []int{1}, []int{1}},
		{"2 elements", []int{4, 1}, []int{1, 4}},
		{"many elements", []int{7, 3, 3, 8, 1, 9, 2}, []int{1, 2, 3, 3, 7, 8, 9}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			InsertionSort(c.input, comparator)
			assert.Equal(t, c.expected, c.input)
		})
	}
}
