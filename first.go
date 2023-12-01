package fn

func First[T any](items []T, pred func(item T, index int) bool) (zero T, idx int) {
	for index, item := range items {
		if pred(item, index) {
			return item, index
		}
	}

	return zero, -1
}
