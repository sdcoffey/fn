package fn

func castSlice[T Number, R Number](slc []T) []R {
	result := make([]R, len(slc))
	for i, v := range slc {
		result[i] = R(v)
	}
	return result
}
