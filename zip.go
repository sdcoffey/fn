package fn

func Zip[Key comparable, Value any](keys []Key, values []Value) map[Key]Value {
	results := make(map[Key]Value)

	for i := 0; i < Min(len(keys), len(values)); i++ {
		results[keys[i]] = values[i]
	}
	return results
}
