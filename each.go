package fn

func Each[T any](items []T, fn func(item T, index int)) {
	for i, item := range items {
		fn(item, i)
	}
}
