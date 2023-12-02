package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax_Ints(t *testing.T) {
	values := []int{10, -1, 40}
	expected := 40

	assert.Equal(t, expected, Max(values...))
	assert.Equal(t, int8(expected), Max(castSlice[int, int8](values)...))
	assert.Equal(t, int16(expected), Max(castSlice[int, int16](values)...))
	assert.Equal(t, int32(expected), Max(castSlice[int, int32](values)...))
	assert.Equal(t, int64(expected), Max(castSlice[int, int64](values)...))
}

func TestMax_Uints(t *testing.T) {
	values := []uint{10, 0, 40}
	var expected uint = 40

	assert.Equal(t, expected, Max(values...))
	assert.Equal(t, uint8(expected), Max(castSlice[uint, uint8](values)...))
	assert.Equal(t, uint16(expected), Max(castSlice[uint, uint16](values)...))
	assert.Equal(t, uint32(expected), Max(castSlice[uint, uint32](values)...))
	assert.Equal(t, uint64(expected), Max(castSlice[uint, uint64](values)...))
}

func TestMax_floats(t *testing.T) {
	values := []float64{-30.4, 103.3, 40.2}

	assert.Equal(t, 103.3, Max(values...))
	assert.Equal(t, float32(103.3), Max(castSlice[float64, float32](values)...))
}

func TestMax_Strings(t *testing.T) {
	values := []string{"c", "d", "a"}
	assert.Equal(t, "d", Max(values...))
}

func TestMax_Uintptr(t *testing.T) {
	values := []uintptr{392, 3929, 128}
	assert.Equal(t, uintptr(3929), Max(values...))
}

type intish = int
type int8ish = int8
type int16ish = int16
type int32ish = int32
type int64ish = int64

func TestMax_intIsh(t *testing.T) {
	values := []intish{10, -1, 40}
	expected := 40

	assert.Equal(t, expected, Max(values...))
	assert.Equal(t, int8ish(expected), Max(castSlice[intish, int8ish](values)...))
	assert.Equal(t, int16ish(expected), Max(castSlice[intish, int16ish](values)...))
	assert.Equal(t, int32ish(expected), Max(castSlice[intish, int32ish](values)...))
	assert.Equal(t, int64ish(expected), Max(castSlice[intish, int64ish](values)...))
}

type uintish = uint
type uint8ish = uint8
type uint16ish = uint16
type uint32ish = uint32
type uint64ish = uint64

func TestMax_uintIsh(t *testing.T) {
	values := []uintish{10, 1, 40}
	expected := uintish(40)

	assert.Equal(t, expected, Max(values...))
	assert.Equal(t, uint8ish(expected), Max(castSlice[uintish, uint8ish](values)...))
	assert.Equal(t, uint16ish(expected), Max(castSlice[uintish, uint16ish](values)...))
	assert.Equal(t, uint32ish(expected), Max(castSlice[uintish, uint32ish](values)...))
	assert.Equal(t, uint64ish(expected), Max(castSlice[uintish, uint64ish](values)...))
}
