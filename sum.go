package fn

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Sum[T Number](numbers []T) (result T) {
	return Reduce(numbers, func(sum T, input T, i int) T {
		return sum + input
	}, result)
}
