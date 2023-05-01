package valid_anagram_test

import (
	"github.com/roadsigns/leet-codes/go/valid_anagram"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert.True(t, valid_anagram.IsAnagram("anagram", "nagaram"))
	assert.False(t, valid_anagram.IsAnagram("rat", "car"))
}
