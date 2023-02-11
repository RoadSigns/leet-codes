package sliding_window

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLength(t *testing.T) {
	nums := []int{3, 1, 2, 7, 4, 2, 1, 1, 5}
	k := 8
	expected := 4

	got := findLength(nums, k)

	assert.Equal(t, expected, got)
}

func TestNumSubarrayProductLessThanK(t *testing.T) {
	nums := []int{10, 5, 2, 6}
	k := 100
	expected := 8

	got := numSubarrayProductLessThanK(nums, k)

	assert.Equal(t, expected, got)
}

func TestFindBestSubArray(t *testing.T) {
	nums := []int{3, -1, 4, 12, -8, 5, 6}
	k := 4
	expected := 18

	got := findBestSubarray(nums, k)

	assert.Equal(t, expected, got)
}
