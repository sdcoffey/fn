package fn

func Zero[T comparable](value T) bool {
	var zeroValue T
	return value == zeroValue
}
