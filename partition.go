package fn

// Partition takes a slice of T and a predicate function and returns two slices of T.
// The first slice contains all the items that passed the predicate function, and the second slice contains all the items that failed the predicate function.
func Partition[T any](items []T, pred func(item T, index int) bool) ([]T, []T) {
	passed, failed := make([]T, 0, len(items)), make([]T, 0, len(items))
	for idx, item := range items {
		if pred(item, idx) {
			passed = append(passed, item)
		} else {
			failed = append(failed, item)
		}
	}
	return passed, failed
}
