package sort

func SelectionSort[T any](arr []T, comparator ComparatorFn[T]) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if comparator(arr[j], arr[min]) < 0 {
				min = j
			}
		}

		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
}
