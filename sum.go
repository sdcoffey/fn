package fn

// number is a union of all number-ish types.
type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~complex64 | ~complex128 | ~uintptr
}

// Sum returns the summed value of values in the slice.
func Sum[T number](numbers []T) (result T) {
	return Reduce(numbers, func(sum T, input T, i int) T {
		return sum + input
	}, result)
}
