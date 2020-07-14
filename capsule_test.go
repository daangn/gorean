package gorean

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intArrayToString(t *testing.T) {
	test1 := []int{0, 1, 2, 3}
	actual := intArrayToString(test1)

	assert.Equal(t, "0,1,2,3", actual, "It'll be equal")
}
