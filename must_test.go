package fn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMust_ValidResult(t *testing.T) {
	value := "some-value"
	result := Must(value, nil)
	assert.Equal(t, "some-value", result)
}

func TestMust_Error(t *testing.T) {
	value := "some-value"

	assert.Panics(t, func() {
		Must(value, fmt.Errorf("some-error"))
	})
}
