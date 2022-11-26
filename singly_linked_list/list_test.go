package singlylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Push(t *testing.T) {
	cases := []struct {
		Val      int
		Expected []int
	}{
		{2, []int{2}},
		{3, []int{2, 3}},
		{0, []int{2, 3, 0}},
	}

	list := NewList[int]()

	for _, test := range cases {
		list.Push(test.Val)
		assert.Equal(t, test.Expected, list.ToSlice())
		assert.Equal(t, len(test.Expected), list.Length())
	}
}

func Test_Unshift(t *testing.T) {
	cases := []struct {
		Val      int
		Expected []int
	}{
		{2, []int{2}},
		{3, []int{3, 2}},
		{0, []int{0, 3, 2}},
	}

	list := NewList[int]()

	for _, test := range cases {
		list.Unshift(test.Val)
		assert.Equal(t, test.Expected, list.ToSlice())
		assert.Equal(t, len(test.Expected), list.Length())

	}
}

func Test_InsertAtIndex(t *testing.T) {
	cases := []struct {
		Val      int
		Index    int
		Expected []int
	}{
		{2, 1, []int{2}},
		{3, 0, []int{3, 2}},
		{0, 2, []int{3, 2, 0}},
		{5, 1, []int{3, 5, 2, 0}},
		{6, 2, []int{3, 5, 6, 2, 0}},
	}

	list := NewList[int]()

	for _, test := range cases {
		list.InsertAtIndex(test.Index, test.Val)
		assert.Equal(t, test.Expected, list.ToSlice())
		assert.Equal(t, len(test.Expected), list.Length())

	}
}
