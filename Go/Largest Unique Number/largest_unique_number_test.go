package largest_unique_number_test

import (
	"github.com/roadsigns/leet-codes/go/largest_unqiue_number"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestUniqueNumber(t *testing.T) {
	expected := 8
	got := largest_unique_number.LargestUniqueNumber([]int{5, 7, 3, 9, 4, 9, 8, 3, 1})
	assert.Equal(t, expected, got)
}
