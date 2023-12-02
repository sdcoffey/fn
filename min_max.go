package fn

type Orderable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string | ~uintptr
}

// Max returns the maximum value in some orderable set, or the zero value of the set type if it has len 0
func Max[T Orderable](items ...T) (zero T) {
	if len(items) == 0 {
		return zero
	}

	maxItem := items[0]
	for _, item := range items[1:] {
		if item > maxItem {
			maxItem = item
		}
	}
	return maxItem
}

// Min returns the minimum value in some orderable set, or the zero value of the set type if it has len 0
func Min[T Orderable](items ...T) (zero T) {
	if len(items) == 0 {
		return zero
	}
	minItem := items[0]
	for _, item := range items[1:] {
		if item < minItem {
			minItem = item
		}
	}
	return minItem
}
