package fn

// Seq Returns a slice of integers from `[start, end)`, incremented by `inc`.
func Seq(start, end, inc int) []int {
	if (end < start && inc > 0) || (start < end && inc < 0) || inc == 0 {
		return []int{}
	}

	i := 0
	values := make([]int, 0, (end-start)/inc)

	for hasNext(start, end, inc) {
		values = append(values, start)
		start += inc
		i++
	}

	return values
}

func hasNext(cMin, cMax, inc int) bool {
	if inc > 0 {
		return cMin < cMax
	} else {
		return cMin > cMax
	}
}
