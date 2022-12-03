package heap

// ComparatorFn compares the priority between two values.
//
// If `a` has higher priority than `b`, it returns a positve int.
// If `a` has lower priority than `b`, it returns a negative int.
// Otherwise it returns 0.
type ComparatorFn[T any] func(a, b T) int
