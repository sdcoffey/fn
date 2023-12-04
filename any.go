package fn

import "reflect"

// Any returns true if any item in the slice satisfies the predicate
func Any[T any](items []T, pred func(item T, index int) bool) bool {
	for idx, item := range items {
		if pred(item, idx) {
			return true
		}
	}
	return false
}

// AnyNonZero is a helper function that returns true if any item in the slice is not the zero value for its type
func AnyNonZero[T comparable](items []T) bool {
	for _, item := range items {
		if !reflect.ValueOf(item).IsZero() {
			return true
		}
	}
	return false
}
