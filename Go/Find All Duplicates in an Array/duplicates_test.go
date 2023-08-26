package duplicates_test

import (
	duplicates "github.com/roadsigns/leet_codes/find_all_duplicates_in_an_array"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	t.Run("find duplicates", func(t *testing.T) {
		nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
		got := duplicates.FindDuplicates(nums)
		want := []int{2, 3}
		assert.Equal(t, want, got)
	})

	t.Run("find single duplicate", func(t *testing.T) {
		nums := []int{1, 1, 2}
		got := duplicates.FindDuplicates(nums)
		want := []int{1}
		assert.Equal(t, want, got)
	})

	t.Run("find no duplicates", func(t *testing.T) {
		nums := []int{1}
		got := duplicates.FindDuplicates(nums)
		want := []int{}
		assert.Equal(t, want, got)
	})
}
