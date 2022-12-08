package sort

func InsertionSort[T any](arr []T, comparator ComparatorFn[T]) {
	n := len(arr)

	for i := 1; i < n; i++ {
		for j := i; j > 0 && comparator(arr[j-1], arr[j]) > 0; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
	}
}
