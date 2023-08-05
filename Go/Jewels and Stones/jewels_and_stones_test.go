package jewels_test

import (
	jewels "github.com/roadsigns/learn-go-with-tests/jewel_and_stones"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumJewelsInStone(t *testing.T) {
	t.Run("stones contain jewels", func(t *testing.T) {
		want := 3
		got := jewels.NumJewelsInStones("aA", "aAAbbbb")
		assert.Equal(t, want, got)
	})

	t.Run("stone contains no jewels", func(t *testing.T) {
		want := 0
		got := jewels.NumJewelsInStones("z", "ZZ")
		assert.Equal(t, want, got)
	})
}
