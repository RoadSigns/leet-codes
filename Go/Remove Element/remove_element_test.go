package remove_element_test

import (
	"github.com/roadsigns/leet_codes/remove_element"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	t.Run("Slice has two int", func(t *testing.T) {
		nums := []int{3, 2, 2, 3}
		got := remove_element.RemoveElement(nums, 3)
		want := 2

		assert.Equal(t, want, got)
	})
}
