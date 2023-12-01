package fn

type Orderable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64 | ~string | ~uintptr
}

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
