package fn

import (
	"reflect"
	"strings"
	"unicode"
)

// Present returns true if the passed value is not blank (see Blank)
func present(value any) bool {
	return !blank(value)
}

type Blankable interface {
	Blank() bool
}

// Blank returns true if the passed value is its type's zero value, a nil pointer, or an all whitespace string
// if value is a boolean, Blank will return true if the value is false
// If value is a slice or array, Blank will return true if the length is 0
// if value is a map, Blank will return true if the length is 0
// if value is a struct, Blank will return true if all fields are blank and it has no methods
func blank(value any) bool {
	// Special handling for error types
	if err, isErr := value.(error); isErr {
		return err == nil
	} else if blankable, isBlankable := value.(Blankable); isBlankable {
		return blankable.Blank()
	}

	rType := reflect.TypeOf(value)
	if rType == nil {
		return true
	}

	rValue := reflect.ValueOf(value)
	switch rValue.Kind() {
	case reflect.String:
		return allWhiteSpace(rValue.String())
	case reflect.Bool:
		return !rValue.Interface().(bool)
	case reflect.Slice, reflect.Array:
		return rValue.Len() == 0
	case reflect.Map:
		if rValue.Len() == 0 {
			return true
		}
	case reflect.Ptr:
		if rValue.IsNil() {
			return true
		} else {
			return blank(rValue.Elem().Interface())
		}
	case reflect.Interface, reflect.Chan, reflect.Func, reflect.UnsafePointer:
		rValue.IsZero()
		return rValue.IsNil()
	default:
		return reflect.DeepEqual(value, reflect.Zero(rType).Interface())
	}

	return false
}

func allWhiteSpace(value string) bool {
	for _, char := range strings.TrimSpace(value) {
		if !unicode.IsSpace(char) {
			return false
		}
	}
	return true
}
