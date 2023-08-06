package duplicates_test

import (
	"github.com/roadsigns/learn-go-with-tests/duplicates"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	t.Run("contains duplicates", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		assert.True(t, duplicates.ContainsDuplicate(nums))
	})

	t.Run("does not contains duplicates", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		assert.False(t, duplicates.ContainsDuplicate(nums))
	})

	t.Run("contains multiple duplicates", func(t *testing.T) {
		nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
		assert.True(t, duplicates.ContainsDuplicate(nums))
	})
}
