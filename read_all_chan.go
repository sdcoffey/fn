package fn

func readAllChan[T any](ch <-chan T) []T {
	values := make([]T, 0)
	for chanItem := range ch {
		values = append(values, chanItem)
	}
	return values
}
