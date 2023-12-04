package fn

import "sync/atomic"

type enumerator[T any] struct {
	items  <-chan T
	closed *atomic.Bool
}

func (e enumerator[T]) Closed() bool {
	return e.closed.Load()
}

func (e enumerator[T]) Each(fn func(item T)) {
	for item := range e.items {
		fn(item)
	}
}

func enumeratorFromSlice[T any](items []T) *enumerator[T] {
	ch := make(chan T, len(items))
	closed := &atomic.Bool{}
	go func() {
		for _, item := range items {
			ch <- item
		}
		close(ch)
		closed.Store(true)
	}()
	return &enumerator[T]{items: ch, closed: closed}
}
