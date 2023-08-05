package frequency_test

import (
	frequency "github.com/roadsigns/learn-go-with-tests/sort_characters_by_frequency"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFrequencySort(t *testing.T) {
	t.Run("change order", func(t *testing.T) {
		want := "eetr"
		got := frequency.FrequencySort("tree")
		assert.Equal(t, want, got)
	})

	t.Run("no need to change the order", func(t *testing.T) {
		want := "cccaaa"
		got := frequency.FrequencySort("cccaaa")
		assert.Equal(t, want, got)
	})
	t.Run("capital letters are handled", func(t *testing.T) {
		want := "bbAa"
		got := frequency.FrequencySort("Aabb")
		assert.Equal(t, want, got)
	})
}
