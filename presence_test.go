package fn

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Testblank_numericValues(t *testing.T) {
	assert.True(t, blank(0))
	assert.True(t, blank(0.0))
	assert.True(t, blank(int8(0)))
	assert.True(t, blank(int16(0)))
	assert.True(t, blank(int32(0)))
	assert.True(t, blank(int64(0)))
	assert.True(t, blank(uint(0)))
	assert.True(t, blank(uint8(0)))
	assert.True(t, blank(uint16(0)))
	assert.True(t, blank(uint32(0)))
	assert.True(t, blank(uint64(0)))
	assert.True(t, blank(float32(0)))
	assert.True(t, blank(float64(0)))
	assert.True(t, blank(complex64(0)))
	assert.True(t, blank(complex128(0)))

	assert.False(t, blank(1))
	assert.False(t, blank(1.0))
	assert.False(t, blank(int8(1)))
	assert.False(t, blank(int16(1)))
	assert.False(t, blank(int32(1)))
	assert.False(t, blank(int64(1)))
	assert.False(t, blank(uint(1)))
	assert.False(t, blank(uint8(1)))
	assert.False(t, blank(uint16(1)))
	assert.False(t, blank(uint32(1)))
	assert.False(t, blank(uint64(1)))
	assert.False(t, blank(float32(1)))
	assert.False(t, blank(float64(1)))
	assert.False(t, blank(complex64(1)))
	assert.False(t, blank(complex128(1)))
}

func Testblank_bool(t *testing.T) {
	assert.True(t, blank(false))
	assert.False(t, blank(true))
}

func Testblank_stringValues(t *testing.T) {
	assert.True(t, blank(""))
	assert.True(t, blank(" "))
	assert.True(t, blank(" \t\n"))
	assert.True(t, blank(" \t\n\r"))
	assert.True(t, blank(" \t\n\r\v\f"))

	assert.False(t, blank("a"))
}

func Testblank_sliceValues(t *testing.T) {
	emptySlice := []int{}
	assert.True(t, blank(emptySlice))

	sliceWithValues := []int{0, 0}
	assert.False(t, blank(sliceWithValues))
}

func Testblank_mapValues(t *testing.T) {
	emptyMap := map[string]int{}
	assert.True(t, blank(emptyMap))

	mapWithValues := map[string]int{"a": 0, "b": 0}
	assert.False(t, blank(mapWithValues))
}

func Testblank_struct(t *testing.T) {
	type example struct {
		Name string
	}
	assert.True(t, blank(example{}))
	assert.False(t, blank(example{Name: "abcd"}))

	// test nil pointer
	var request *http.Request
	assert.True(t, blank(request))

	// test error
	var err error = nil
	assert.True(t, blank(err))

	// test non-blank error
	assert.False(t, blank(errors.New("example error")))

	// wrapped error
	wrappedError := fmt.Errorf("wrapped error: %w", errors.New("example error"))
	assert.False(t, blank(wrappedError))
}
