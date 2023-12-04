package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunk(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, Chunk([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3))
	assert.Equal(t, [][]int{{1, 2, 3}, {4, 5, 6}, {7}}, Chunk([]int{1, 2, 3, 4, 5, 6, 7}, 3))
	assert.Equal(t, [][]int{{1, 2, 3}}, Chunk([]int{1, 2, 3}, 3))
}

func TestChunkWhile(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}, ChunkWhile([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, func(eltBefore, eltAfter int) bool {
		return eltBefore+1 == eltAfter
	}))
	assert.Equal(t, [][]int{{1, 2}, {4, 5}, {7}}, ChunkWhile([]int{1, 2, 4, 5, 7}, func(eltBefore, eltAfter int) bool {
		return eltBefore+1 == eltAfter
	}))
}
