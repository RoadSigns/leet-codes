package two_sums

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanCalculateWhenOnlySameNumberNeeded(t *testing.T) {
	expected := []int{0, 1}
	got := twoSum([]int{3, 3}, 6)

	assert.Equal(t, expected, got)
}
