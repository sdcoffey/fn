package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestSequenceGenerator(t *testing.T) {
	assert.Equal(t, []int{30, 32, 34}, readAllChan(GenSeq(30, 35, 2)))
	assert.Equal(t, []int{30, 32, 34}, readAllChan(GenSeq(30, 35, 2)))
	assert.Equal(t, []int{}, readAllChan(GenSeq(10, 0, 1)))
	assert.Equal(t, []int{}, readAllChan(GenSeq(0, 10, -1)))
	assert.Equal(t, []int{}, readAllChan(GenSeq(0, 10, 0)))
	assert.Equal(t, []int{}, readAllChan(GenSeq(10, 10, 1)))
}
