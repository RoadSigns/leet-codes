package intersection_of_mulitple_arrays_test

import (
	"github.com/roadsigns/leet-codes/go/intersection_of_mulitple_arrays"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersection(t *testing.T) {
	// [[3,1,2,4,5],[1,2,3,4],[3,4,5,6]]
	nums := [][]int{
		{3, 1, 2, 4, 5},
		{1, 2, 3, 4},
		{3, 4, 5, 6},
	}

	expected := []int{3, 4}
	got := intersection_of_mulitple_arrays.Intersection(nums)
	assert.Equal(t, expected, got)
}

func TestIntersectionReturnsEmptySlice(t *testing.T) {
	// [[1,2,3],[4,5,6]]
	nums := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	expected := []int{}
	got := intersection_of_mulitple_arrays.Intersection(nums)
	assert.Equal(t, expected, got)
}
