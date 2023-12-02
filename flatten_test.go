package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlatten(t *testing.T) {
	items := [][]string{{"one"}, {"two", "three"}, {"four"}}
	assert.Equal(t, []string{"one", "two", "three", "four"}, Flatten(items))
}
