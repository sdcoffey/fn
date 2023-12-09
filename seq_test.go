package fn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequence(t *testing.T) {
	assert.Equal(t, []int{30, 31, 32, 33, 34}, Seq(30, 35, 1))
	assert.Equal(t, []int{0, 2, 4, 6, 8, 10}, Seq(0, 11, 2))
	assert.Equal(t, []int{-10, -12, -14, -16, -18}, Seq(-10, -20, -2))
	assert.Equal(t, []int{}, Seq(10, 0, 1))
	assert.Equal(t, []int{}, Seq(0, 10, -1))
	assert.Equal(t, []int{}, Seq(0, 10, 0))
	assert.Equal(t, []int{}, Seq(10, 10, 1))
}
