package occurances_test

import (
	occurances "github.com/roadsigns/learn-go-with-tests/number_of_occurrances"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueOccurrences(t *testing.T) {
	t.Run("has unique occurrences", func(t *testing.T) {
		nums := []int{1, 2, 2, 1, 1, 3}
		assert.True(t, occurances.UniqueOccurrences(nums))
	})

	t.Run("does not have unique occurrences", func(t *testing.T) {
		nums := []int{1, 2}
		assert.False(t, occurances.UniqueOccurrences(nums))
	})

	t.Run("has unique occurrences with negative values", func(t *testing.T) {
		nums := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
		assert.True(t, occurances.UniqueOccurrences(nums))
	})

	t.Run("has no unique with negative", func(t *testing.T) {
		nums := []int{3, 5, -2, -3, -6, -6}
		assert.False(t, occurances.UniqueOccurrences(nums))
	})
}
