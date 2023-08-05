package substring_test

import (
	substring "github.com/roadsigns/learn-go-with-tests/longest_substring_without_repeat_characters"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("longest substring is first letters", func(t *testing.T) {
		want := 3
		got := substring.LengthOfLongestSubstring("abcabcbb")
		assert.Equal(t, want, got)
	})

	t.Run("longest substring is one letter", func(t *testing.T) {
		want := 1
		got := substring.LengthOfLongestSubstring("bbbbb")
		assert.Equal(t, want, got)
	})

	t.Run("longest substring is in the middle", func(t *testing.T) {
		want := 3
		got := substring.LengthOfLongestSubstring("pwwkew")
		assert.Equal(t, want, got)
	})

	t.Run("letter in substring is straight after the original", func(t *testing.T) {
		want := 2
		got := substring.LengthOfLongestSubstring("aab")
		assert.Equal(t, want, got)
	})

	t.Run("letter repeats half way into a substring", func(t *testing.T) {
		want := 3
		got := substring.LengthOfLongestSubstring("dvdf")
		assert.Equal(t, want, got)
	})
}
