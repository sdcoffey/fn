package fn

// Must allows you to return one value from a function that would normally return a value and an error.
// If the error is present, Must will panic.
func Must[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}
	return result
}
