package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartition(t *testing.T) {
	values := Seq(0, 10, 1)
	evens, odds := Partition(values, func(item, index int) bool {
		return item%2 == 0
	})
	assert.Equal(t, []int{0, 2, 4, 6, 8}, evens)
	assert.Equal(t, []int{1, 3, 5, 7, 9}, odds)
}
