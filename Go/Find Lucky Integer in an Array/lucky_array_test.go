package lucky_array_test

import (
	"github.com/roadsigns/learn-go-with-tests/lucky_array"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindLucky(t *testing.T) {
	t.Run("only one lucky number", func(t *testing.T) {
		want := 2
		got := lucky_array.FindLucky([]int{2, 2, 3, 4})
		assert.Equal(t, want, got)
	})

	t.Run("two lucky numbers", func(t *testing.T) {
		want := 3
		got := lucky_array.FindLucky([]int{1, 2, 2, 3, 3, 3})
		assert.Equal(t, want, got)
	})

	t.Run("no lucky number", func(t *testing.T) {
		want := -1
		got := lucky_array.FindLucky([]int{2, 2, 2, 3, 3})
		assert.Equal(t, want, got)
	})
}
