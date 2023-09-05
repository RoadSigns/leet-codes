package remove_adjacent_duplicates_test

import (
	"github.com/roadsigns/leet_codes/remove_adjacent_duplicates"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	t.Run("remove duplicates", func(t *testing.T) {
		s := "abbaca"
		got := remove_adjacent_duplicates.RemoveDuplicates(s)
		want := "ca"
		assert.Equal(t, want, got)
	})

	t.Run("no duplicates", func(t *testing.T) {
		s := "abc"
		got := remove_adjacent_duplicates.RemoveDuplicates(s)
		want := "abc"
		assert.Equal(t, want, got)
	})

	t.Run("only duplicates", func(t *testing.T) {
		s := "abccba"
		got := remove_adjacent_duplicates.RemoveDuplicates(s)
		want := ""
		assert.Equal(t, want, got)
	})
}
