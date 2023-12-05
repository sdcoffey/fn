package enum

import "sync/atomic"

func selectEnum[T any](e *enumerator[T], pred func(item T) bool) *enumerator[T] {
	outChan := make(chan T)
	closed := &atomic.Bool{}
	go func() {
		for item := range e.items {
			if pred(item) {
				outChan <- item
			}
		}
		close(outChan)
		closed.Store(true)
	}()
	return &enumerator[T]{items: outChan, closed: closed}
}

func mapEnum[T any, U any](e *enumerator[T], fn func(item T) U) *enumerator[U] {
	outChan := make(chan U)
	closed := &atomic.Bool{}
	go func() {
		for item := range e.items {
			outChan <- fn(item)
		}
		close(outChan)
		closed.Store(true)
	}()
	return &enumerator[U]{items: outChan, closed: closed}
}
