package make_the_string_great_test

import (
	"github.com/roadsigns/leet_codes/make_the_string_great"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMakeGood(t *testing.T) {
	t.Run("clear string", func(t *testing.T) {
		got := make_the_string_great.MakeGood("aAbBcC")
		want := ""

		assert.Equal(t, want, got)
	})

	t.Run("one letter", func(t *testing.T) {
		got := make_the_string_great.MakeGood("leEeetcode")
		want := "leetcode"
		assert.Equal(t, want, got)
	})

	t.Run("letter is the other way around", func(t *testing.T) {
		got := make_the_string_great.MakeGood("Pp")
		want := ""
		assert.Equal(t, want, got)
	})

	t.Run("double captial", func(t *testing.T) {
		got := make_the_string_great.MakeGood("kkdsFuqUfSDKK")
		want := "kkdsFuqUfSDKK"
		assert.Equal(t, want, got)
	})
}
