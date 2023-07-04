package first_unique_character_in_a_string_test

import (
	"github.com/roadsigns/leet-codes/go/first_unique_character_in_a_string"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	expected := 0
	got := first_unique_character_in_a_string.FirstUniqChar("leetcode")
	assert.Equal(t, expected, got)

	expected = 2
	got = first_unique_character_in_a_string.FirstUniqChar("loveleetcode")
	assert.Equal(t, expected, got)

	expected = -1
	got = first_unique_character_in_a_string.FirstUniqChar("aabb")
	assert.Equal(t, expected, got)
}

func TestFirstUniqCharWithRuneValue(t *testing.T) {
	expected := 0
	got := first_unique_character_in_a_string.FirstUniqCharWithRuneValue("leetcode")
	assert.Equal(t, expected, got)

	expected = 2
	got = first_unique_character_in_a_string.FirstUniqCharWithRuneValue("loveleetcode")
	assert.Equal(t, expected, got)

	expected = -1
	got = first_unique_character_in_a_string.FirstUniqCharWithRuneValue("aabb")
	assert.Equal(t, expected, got)
}
