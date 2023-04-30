package valid_palindrome_test

import (
	"github.com/roadsigns/leet-codes/go/valid_palindrome"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert.True(t, valid_palindrome.IsPalindrome("racecar"))
	assert.True(t, valid_palindrome.IsPalindrome("A man, a plan, a canal: Panama"))
}
