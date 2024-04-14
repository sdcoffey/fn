package fn

import "reflect"

// All returns true if all items in the slice satisfy the predicate.
func All[T any](items []T, pred func(item T, index int) bool) bool {
	for idx, item := range items {
		if !pred(item, idx) {
			return false
		}
	}
	return true
}

// AllNonZero is a helper function that returns true if all items in the slice are not the zero value for their type.
func AllNonZero[T comparable](items []T) bool {
	for _, item := range items {
		if reflect.ValueOf(item).IsZero() {
			return false
		}
	}
	return true
}
