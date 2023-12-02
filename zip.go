package fn

// Zip takes two slices and aggregates them in to a map.
// If `keys` and `values` are not the same length, Zip will use the shorter of the two to determine the length of the
// resulting map
func Zip[Key comparable, Value any](keys []Key, values []Value) map[Key]Value {
	results := make(map[Key]Value)

	for i := 0; i < Min(len(keys), len(values)); i++ {
		results[keys[i]] = values[i]
	}
	return results
}
