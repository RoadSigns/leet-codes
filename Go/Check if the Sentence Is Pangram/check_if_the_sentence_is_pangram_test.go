package check_if_the_sentence_is_pangram_test

import (
	"github.com/roadsigns/leet-codes/go/check_if_the_sentence_is_pangram"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckIfPangram(t *testing.T) {
	assert.True(t, check_if_the_sentence_is_pangram.CheckIfPangram("thequickbrownfoxjumpsoverthelazydog"))
	assert.False(t, check_if_the_sentence_is_pangram.CheckIfPangram("leetcode"))
}
