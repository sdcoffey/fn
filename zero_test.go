package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZero_primitive(t *testing.T) {
	assert.True(t, Zero(""))
	assert.False(t, Zero(" "))
}

func TestZero_struct(t *testing.T) {
	type example struct {
		Name string
	}
	assert.True(t, Zero(example{}))
	assert.False(t, Zero(example{Name: "abcd"}))
}
