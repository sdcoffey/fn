package fn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertSumAs[T number](t *testing.T, slc []T, expected T) {
	assert.Equal(t, expected, Sum[T](slc))
}

func TestSum_Int(t *testing.T) {
	assertSumAs(t, []int{10, 20, 30}, 60)
}

func TestSum_Int8(t *testing.T) {
	assertSumAs(t, []int8{10, 20, 30}, 60)
}

func TestSum_Int16(t *testing.T) {
	assertSumAs(t, []int16{10, 20, 30}, 60)
}

func TestSum_Int32(t *testing.T) {
	assertSumAs(t, []int32{10, 20, 30}, 60)
}

func TestSum_Int64(t *testing.T) {
	assertSumAs(t, []int64{10, 20, 30}, 60)
}

func TestSum_Uint(t *testing.T) {
	assertSumAs(t, []uint{10, 20, 30}, 60)
}

func TestSum_Uint8(t *testing.T) {
	assertSumAs(t, []uint8{10, 20, 30}, 60)
}

func TestSum_Uint16(t *testing.T) {
	assertSumAs(t, []uint16{10, 20, 30}, 60)
}

func TestSum_Uint32(t *testing.T) {
	assertSumAs(t, []uint32{10, 20, 30}, 60)
}

func TestSum_Uint64(t *testing.T) {
	assertSumAs(t, []uint64{10, 20, 30}, 60)
}

func TestSum_Float32(t *testing.T) {
	assertSumAs(t, []float32{10, 20, 30}, 60)
}

func TestSum_Float64(t *testing.T) {
	assertSumAs(t, []float64{10, 20, 30}, 60)
}

func TestSum_Complex64(t *testing.T) {
	assertSumAs(t, []complex64{10, 20, 30}, 60)
}

func TestSum_Complex128(t *testing.T) {
	assertSumAs(t, []complex128{10, 20, 30}, 60)
}

func TestSum_uintptr(t *testing.T) {
	assertSumAs(t, []uintptr{10, 20, 30}, 60)
}
