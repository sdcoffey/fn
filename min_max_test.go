package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertMaxAs[T Orderable](t *testing.T, items []T, expected T) {
	assert.Equal(t, expected, Max(items...))
}

func assertMinAs[T Orderable](t *testing.T, items []T, expected T) {
	assert.Equal(t, expected, Min(items...))

}

func TestMax_Ints(t *testing.T) {
	assertMaxAs(t, []int{500, 100, 3000}, 3000)
}

func TestMax_Strings(t *testing.T) {
	assertMaxAs(t, []string{"a", "b", "c"}, "c")
}

func TestMin_Ints(t *testing.T) {
	assertMinAs(t, []int{500, 100, 3000}, 100)
}

func TestMin_Strings(t *testing.T) {
	assertMinAs(t, []string{"c", "a", "z"}, "a")
}
