package running_sum_of_1d_array_test

import (
	"github.com/roadsigns/leet-codes/go/running_sum_of_1d_array"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunningSum(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	got := running_sum_of_1d_array.RunningSum([]int{1, 1, 1, 1, 1})
	assert.Equal(t, expected, got)

	expected = []int{3, 4, 6, 16, 17}
	got = running_sum_of_1d_array.RunningSum([]int{3, 1, 2, 10, 1})
	assert.Equal(t, expected, got)
}
