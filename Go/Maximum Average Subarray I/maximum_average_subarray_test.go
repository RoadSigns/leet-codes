package maximum_average_subarray_i

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMaxAverage(t *testing.T) {
	expected := 12.75
	got := findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)
	assert.Equal(t, expected, got)

	expected = 5.0
	got = findMaxAverage([]int{5}, 1)
	assert.Equal(t, expected, got)
}
