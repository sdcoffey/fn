package enum

import "sync/atomic"

// seq returns an enumerator that yields integers from `[start, end)`, incremented by `inc`.
func seq(start, end, inc int) *enumerator[int] {
	outputChan := make(chan int)
	done := &atomic.Bool{}

	go func() {
		defer close(outputChan)
		defer done.Store(true)

		if (end < start && inc > 0) || (start < end && inc < 0) || inc == 0 {
			return
		}

		for hasNext(start, end, inc) {
			outputChan <- start
			start += inc
		}
	}()

	return &enumerator[int]{items: outputChan, closed: done}
}

func hasNext(cMin, cMax, inc int) bool {
	if inc > 0 {
		return cMin < cMax
	} else {
		return cMin > cMax
	}
}
