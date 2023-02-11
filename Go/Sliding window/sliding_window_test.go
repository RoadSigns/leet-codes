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
