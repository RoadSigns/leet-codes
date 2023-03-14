package roman_to_integer_test

import (
	"github.com/roadsigns/leet-codes/go/romain_to_integer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	expected := 4
	got := roman_to_integer.RomanToInt("IV")
	assert.Equal(t, expected, got)

	expected = 88
	got = roman_to_integer.RomanToInt("LXXXVIII")
	assert.Equal(t, expected, got)

	expected = 1994
	got = roman_to_integer.RomanToInt("MCMXCIV")
	assert.Equal(t, expected, got)

	expected = 3
	got = roman_to_integer.RomanToInt("III")
	assert.Equal(t, expected, got)

	expected = 1
	got = roman_to_integer.RomanToInt("I")
	assert.Equal(t, expected, got)
}
