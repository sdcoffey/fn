package fn

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestReduce(t *testing.T) {
	ints := []int{10, 20, 30}

	keyToValue := Reduce(ints, func(output map[string]int, input int, i int) map[string]int {
		output[strconv.Itoa(i)] = input
		return output
	}, make(map[string]int))

	assert.Equal(t, map[string]int{
		"0": 10,
		"1": 20,
		"2": 30,
	}, keyToValue)
}
