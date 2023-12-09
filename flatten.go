package fn

// Flatten takes a 2D slice `[][]T` and flattens it to one dimension, preserving the original order.
func Flatten[T any](slice [][]T) []T {
	result := make([]T, 0, len(slice))

	for _, innerSlc := range slice {
		result = append(result, innerSlc...)
	}
	return result
}
