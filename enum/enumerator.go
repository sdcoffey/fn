package enum

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

func (e enumerator[T]) selectEnum(fn func(item T) bool) *enumerator[T] {
	outChan := make(chan T)
	closed := &atomic.Bool{}
	go func() {
		defer close(outChan)
		defer closed.Store(true)
		for item := range e.items {
			if fn(item) {
				outChan <- item
			}
		}
	}()
	return &enumerator[T]{items: outChan, closed: closed}
}

func (e enumerator[T]) ReadAll() (items []T) {
	for item := range e.items {
		items = append(items, item)
	}
	return items
}

func fromFunc[T any](initialValue T, next func(lastValue T) (T, bool)) *enumerator[T] {
	ch := make(chan T)
	closed := &atomic.Bool{}
	go func() {
		defer close(ch)
		defer closed.Store(true)
		value := initialValue
		cont := true
		for cont {
			ch <- value
			value, cont = next(value)
		}
	}()
	return &enumerator[T]{items: ch, closed: closed}
}

func fromSlice[T any](items []T) *enumerator[T] {
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

func fromChan[T any](ch <-chan T) *enumerator[T] {
	closed := &atomic.Bool{}
	outChan := make(chan T)
	go func() {
		for item := range ch {
			outChan <- item
		}
		close(outChan)
		closed.Store(true)
	}()
	return &enumerator[T]{items: outChan, closed: closed}
}
