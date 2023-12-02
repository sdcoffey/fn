package fn

// Flatten takes a two-dimensional array of T and flattens it down to one dimension, preserving the original order
func Flatten[T any](slice [][]T) []T {
	result := make([]T, 0, len(slice))

	for _, innerSlc := range slice {
		for _, item := range innerSlc {
			result = append(result, item)
		}
	}
	return result
}
