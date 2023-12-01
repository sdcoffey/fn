package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZip(t *testing.T) {
	keys := []string{"one", "two", "three"}
	values := []int{1, 2, 3}

	assert.Equal(t, map[string]int{
		"one": 1, "two": 2, "three": 3,
	}, Zip(keys, values))
}

func TestZip_Unequal(t *testing.T) {
	assert.Equal(t, map[string]int{
		"one": 1, "two": 2,
	}, Zip([]string{"one", "two"}, []int{1, 2, 3}))

	assert.Equal(t, map[string]int{
		"one": 1, "two": 2,
	}, Zip([]string{"one", "two", "three"}, []int{1, 2}))

}
