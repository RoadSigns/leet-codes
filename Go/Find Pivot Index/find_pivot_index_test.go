package find_pivot_index_test

import (
	"github.com/roadsigns/leet-codes/go/find_pivot_index"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	assert.Equal(t, 3, find_pivot_index.PivotIndex([]int{1, 7, 3, 6, 5, 6}))
	assert.Equal(t, -1, find_pivot_index.PivotIndex([]int{1, 2, 3}))
	assert.Equal(t, 0, find_pivot_index.PivotIndex([]int{2, 1, -1}))
	assert.Equal(t, 2, find_pivot_index.PivotIndex([]int{-1, -1, -1, -1, -1, 0}))
	assert.Equal(t, 2, find_pivot_index.PivotIndex([]int{-1, -1, 0, 0, -1, -1}))
}
