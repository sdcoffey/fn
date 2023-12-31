package fn

// Select reduces a slice to only values for which `selector` returns true.
func Select[T any](items []T, selector func(item T, index int) bool) []T {
	result := make([]T, 0, len(items))
	for index, item := range items {
		if selector(item, index) {
			result = append(result, item)
		}
	}
	return result
}

// Reject reduces a slice to only values for which `rejector` returns false.
func Reject[T any](items []T, rejector func(item T, index int) bool) []T {
	return Select(items, func(item T, index int) bool {
		return !rejector(item, index)
	})
}
