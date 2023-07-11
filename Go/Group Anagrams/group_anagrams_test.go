package group_anagrams_test

import (
	"github.com/roadsigns/leet-codes/go/group_anagrams"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	expected := [][]string{
		{"eat", "tea", "ate"},
		{"tan", "nat"},
		{"bat"},
	}
	got := group_anagrams.GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	assert.Equal(t, expected, got)
}
