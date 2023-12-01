package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelect(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5}
	evens := Select(nums, func(item int, index int) bool {
		return item%2 == 0
	})
	assert.Equal(t, []int{0, 2, 4}, evens)
}

func TestReject(t *testing.T) {
	nums := []int{-2, -1, 0, 1, 2}
	negatives := Reject(nums, func(item int, index int) bool {
		return item >= 0
	})
	assert.Equal(t, []int{-2, -1}, negatives)
}
