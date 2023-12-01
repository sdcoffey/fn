package fn

func Map[Input any, Result any](inputs []Input, mapper func(item Input, index int) Result) []Result {
	results := make([]Result, len(inputs))
	for index, item := range inputs {
		results[index] = mapper(item, index)
	}
	return results
}
