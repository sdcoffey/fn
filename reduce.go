package fn

// Reduce condenses a collection into an accumulated value which is the result of running each element
// in the collection though `reducer`, where each successive iteration is supplied the return value of the previous.
func Reduce[Input any, Output any](inputs []Input, reducer func(Output, Input, int) Output, value Output) Output {
	for index, item := range inputs {
		value = reducer(value, item, index)
	}
	return value
}
