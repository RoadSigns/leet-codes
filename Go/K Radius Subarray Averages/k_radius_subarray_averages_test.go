package k_radius_subarray_averages_test

import (
	"github.com/roadsigns/leet-codes/go/k_radius_subarray_averages"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAverages(t *testing.T) {
	ints := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	expected := []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}

	actual := k_radius_subarray_averages.GetAverages(ints, 3)

	assert.Equal(t, expected, actual)
}

func TestGetAveragesUsingRadius(t *testing.T) {
	ints := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	expected := []int{-1, -1, -1, 5, 4, 4, -1, -1, -1}

	actual := k_radius_subarray_averages.GetAverageUsingRadius(ints, 3)

	assert.Equal(t, expected, actual)
}
