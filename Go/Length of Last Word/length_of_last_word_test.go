package length_of_last_word_test

import (
	"github.com/roadsigns/leet-codes/go/length_of_last_word"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	expect := 5
	got := length_of_last_word.LengthOfLastWord("Hello World")
	assert.Equal(t, expect, got)

	expect = 4
	got = length_of_last_word.LengthOfLastWord("   fly me   to   the moon  ")
	assert.Equal(t, expect, got)
}
