package fn

func Any[T any](items []T, pred func(item T, index int) bool) bool {
	for idx, item := range items {
		if pred(item, idx) {
			return true
		}
	}
	return false
}

func AnyNonZero[T comparable](items []T) bool {
	for _, item := range items {
		if !Zero(item) {
			return true
		}
	}
	return false
}
