package fn

import (
	"reflect"
)

// CompactZero returns a new slice with all zero items removed.
func CompactZero[T comparable](items []T) []T {
	result := make([]T, 0, len(items))
	for _, item := range items {
		if !reflect.ValueOf(item).IsZero() {
			result = append(result, item)
		}
	}
	return result
}

// CompactNil returns a new slice with all nil items removed.
func CompactNil[T any](items []T) []T {
	result := make([]T, 0, len(items))
	for _, item := range items {
		value := reflect.ValueOf(item)
		switch value.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
			if !value.IsNil() {
				result = append(result, item)
			}
		case reflect.Invalid:
		default:
			result = append(result, item)
		}
	}
	return result
}
