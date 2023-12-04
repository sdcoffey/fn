package fn

// Partition takes `[]T` and a predicate function and returns two slices of `T`.
// The first slice contains all the items for which `pred` returned true, and the second slice contains all the items for which `pred` returned false.
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
