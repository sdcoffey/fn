package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirst(t *testing.T) {
	ints := []int{-2, -1, 0, 1, 2}
	first, index := First(ints, func(item, index int) bool {
		return item > 0
	})
	assert.Equal(t, 1, first)
	assert.Equal(t, 3, index)
}

func TestFirst_zero(t *testing.T) {
	ints := []int{-2, -1, 0, 1, 2}
	first, index := First(ints, func(item, index int) bool {
		return item > 10
	})
	assert.Equal(t, 0, first)
	assert.Equal(t, -1, index)
}
