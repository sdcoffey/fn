package fn

// First will return the first value that satisfies the predicate, and the index at which the value was found
// (or -1 if it was not found).
func First[T any](items []T, pred func(item T, index int) bool) (zero T, idx int) {
	for index, item := range items {
		if pred(item, index) {
			return item, index
		}
	}

	return zero, -1
}
