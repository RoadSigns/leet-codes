package binary_search_test

import (
	"github.com/roadsigns/leet-codes/go/binary_search"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	expected := 4
	got := binary_search.Search([]int{-1, 0, 3, 5, 9, 12}, 9)
	assert.Equal(t, expected, got)

	expected = -1
	got = binary_search.Search([]int{-1, 0, 3, 5, 9, 12}, 2)
	assert.Equal(t, expected, got)
}
