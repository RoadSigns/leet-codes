package backspace_string_test

import (
	"github.com/roadsigns/leet_codes/backspace_string"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBackspaceCompare(t *testing.T) {
	t.Run("has backspace to remove and are the same", func(t *testing.T) {
		a := "a#c"
		b := "b#c"
		assert.True(t, backspace_string.BackspaceCompare(a, b))
	})

	t.Run("has backspace to remove and aren't the same", func(t *testing.T) {
		a := "a#c"
		b := "b#b"
		assert.False(t, backspace_string.BackspaceCompare(a, b))
	})

	t.Run("has no backspace to remove and are the same", func(t *testing.T) {
		a := "ac"
		b := "ac"
		assert.True(t, backspace_string.BackspaceCompare(a, b))
	})

	t.Run("has no backspace to remove and are the same", func(t *testing.T) {
		a := "ac"
		b := "bc"
		assert.False(t, backspace_string.BackspaceCompare(a, b))
	})

	t.Run("has no backspace to remove and are the same", func(t *testing.T) {
		a := "y#fo##f"
		b := "y#f#o##f"
		assert.True(t, backspace_string.BackspaceCompare(a, b))
	})

	t.Run("has no backspace to remove and no characters", func(t *testing.T) {
		a := ""
		b := "bc"
		assert.False(t, backspace_string.BackspaceCompare(a, b))
	})
}
