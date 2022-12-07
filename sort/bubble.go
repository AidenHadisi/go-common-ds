package sort

func BubbleSort[T any](arr []T, comparator ComparatorFn[T]) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swaps := false
		for j := 0; j < n-i-1; j++ {
			if comparator(arr[j], arr[j+1]) > 0 {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swaps = true
			}
		}

		if !swaps {
			break
		}
	}
}
