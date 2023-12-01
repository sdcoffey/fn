package fn

func Reduce[Input any, Output any](inputs []Input, reducer func(Output, Input, int) Output, value Output) Output {
	for index, item := range inputs {
		value = reducer(value, item, index)
	}
	return value
}
