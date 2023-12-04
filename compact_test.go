package fn

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompact(t *testing.T) {
	ints := []int{0, 1, 2, 3, 0, 5}
	assert.Equal(t, []int{1, 2, 3, 5}, CompactZero(ints))

	bools := []bool{true, false, true, false}
	assert.Equal(t, []bool{true, true}, CompactZero(bools))

	strings := []string{"", "a", "b", "", "c"}
	assert.Equal(t, []string{"a", "b", "c"}, CompactZero(strings))
}

func TestCompactNil(t *testing.T) {
	reqs := []*struct{}{nil, {}, nil, {}}
	assert.Equal(t, []*struct{}{{}, {}}, CompactNil(reqs))

	bools := []bool{true, false, true, false}
	assert.Equal(t, []bool{true, false, true, false}, CompactNil(bools))

	errs := []error{nil, nil, nil, errors.New("example-error")}
	assert.Equal(t, []error{errors.New("example-error")}, CompactNil(errs))
}
