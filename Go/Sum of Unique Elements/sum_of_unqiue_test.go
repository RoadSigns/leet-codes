package sum_test

import (
	sum "github.com/roadsigns/learn-go-with-tests/sum_of_unquie"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfUnique(t *testing.T) {
	t.Run("mixture of uniques", func(t *testing.T) {
		want := 4
		got := sum.SumOfUnique([]int{1, 2, 3, 2})
		assert.Equal(t, want, got)
	})

	t.Run("no uniques", func(t *testing.T) {
		want := 0
		got := sum.SumOfUnique([]int{1, 1, 1, 1, 1})
		assert.Equal(t, want, got)
	})

	t.Run("all unique", func(t *testing.T) {
		want := 15
		got := sum.SumOfUnique([]int{1, 2, 3, 4, 5})
		assert.Equal(t, want, got)
	})
}
