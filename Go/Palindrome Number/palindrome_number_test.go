package palindrome_number_test

import (
	"github.com/roadsigns/leet-codes/go/palindrome_number"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, palindrome_number.IsPalindrome(121))
	assert.False(t, palindrome_number.IsPalindrome(-121))
	assert.False(t, palindrome_number.IsPalindrome(10))
}

func TestIsPalindromeWithoutStringConversion(t *testing.T) {
	assert.True(t, palindrome_number.IsPalindromeWithoutStringConversion(121))
	assert.False(t, palindrome_number.IsPalindromeWithoutStringConversion(-121))
	assert.False(t, palindrome_number.IsPalindromeWithoutStringConversion(10))
}
