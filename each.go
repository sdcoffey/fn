package fn

// Each calls the given function for each item in the slice.
func Each[T any](items []T, fn func(item T, index int)) {
	for i, item := range items {
		fn(item, i)
	}
}
