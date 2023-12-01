package fn

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	ints := []int{0, 2, 4, 6}
	labeled := Map(ints, func(item int, index int) string {
		return fmt.Sprintf("%d:%d", index, item)
	})

	assert.Equal(t, []string{"0:0", "1:2", "2:4", "3:6"}, labeled)
}
