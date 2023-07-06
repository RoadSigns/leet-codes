package missing_number_test

import (
	"github.com/roadsigns/leet-codes/go/missing_number"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	expected := 2
	nums := []int{0, 1, 3}
	got := missing_number.MissingNumber(nums)
	assert.Equal(t, expected, got)

	expected = 2
	nums = []int{0, 1}
	got = missing_number.MissingNumber(nums)
	assert.Equal(t, expected, got)
}
