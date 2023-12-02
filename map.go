package fn

// Map transforms allows you to transform all values into a slice of other values. The original slice is not modified.
func Map[Input any, Result any](inputs []Input, mapper func(item Input, index int) Result) []Result {
	results := make([]Result, len(inputs))
	for index, item := range inputs {
		results[index] = mapper(item, index)
	}
	return results
}
