package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("All", func(t *testing.T) {
		assert.True(t, All([]int{1, 2, 3, 4, 5}, func(item int, index int) bool {
			return item > 0
		}))
		assert.False(t, All([]int{1, 2, 3, 4, 5}, func(item int, index int) bool {
			return item > 2
		}))
	})

	t.Run("AllNonZero", func(t *testing.T) {
		assert.True(t, AllNonZero([]int{1, 2, 3, 4, 5}))
		assert.False(t, AllNonZero([]int{1, 2, 3, 0, 5}))
	})
}
