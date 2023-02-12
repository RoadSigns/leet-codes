package max_consecutive_ones_iii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestOnes(t *testing.T) {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2

	expected := 6
	got := longestOnes(nums, k)
	assert.Equal(t, expected, got)

	nums = []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k = 3
	expected = 10
	got = longestOnes(nums, k)
	assert.Equal(t, expected, got)
}
