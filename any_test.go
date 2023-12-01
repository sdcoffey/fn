package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	ints := []int{2, 2, 2, 2, 2, 3}
	assert.True(t, Any(ints, func(item, index int) bool {
		return item%2 == 1
	}))

	ints = []int{2, 2, 2, 2, 2, 4}
	assert.False(t, Any(ints, func(item, index int) bool {
		return item%2 == 1
	}))
}

func TestAnyNonZero(t *testing.T) {
	truthy := []int{0, 0, 1, 0}
	assert.True(t, AnyNonZero(truthy))

	falsy := []string{"", "", ""}
	assert.False(t, AnyNonZero(falsy))
}

func TestAnyNonZero_struct(t *testing.T) {

	type Example struct {
		Name string
	}

	assert.True(t, AnyNonZero([]Example{{Name: "abcd"}, {}})) // true
}
