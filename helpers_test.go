package fn

import "reflect"

func castSlice[T number, R any](slc []T) []R {
	result := make([]R, len(slc))
	for i, v := range slc {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8, reflect.Float32, reflect.Float64:
			var r R
			result[i] = reflect.ValueOf(v).Convert(reflect.TypeOf(r)).Interface().(R)
		case reflect.Complex128, reflect.Complex64:
			var r R
			result[i] = reflect.ValueOf(v).Convert(reflect.TypeOf(r)).Interface().(R)
		}
	}
	return result
}
