package fn

// Zero returns whether the comparable value passed in is the zero value for its type
func Zero[T comparable](value T) bool {
	var zeroValue T
	return value == zeroValue
}
