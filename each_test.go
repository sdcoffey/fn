package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEach(t *testing.T) {
	items, indices := []int{}, []int{}
	Each([]int{1, 2, 3}, func(item int, index int) {
		items = append(items, item)
		indices = append(indices, index)
	})
	assert.Equal(t, []int{1, 2, 3}, items)
	assert.Equal(t, []int{0, 1, 2}, indices)
}
