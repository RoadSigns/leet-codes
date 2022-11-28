package fizz_buzz_test

import (
	"github.com/roadsigns/leet-codes/go/fizz_buzz"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzzWhenAllPossibleCombosAreAvailable(t *testing.T) {
	expected := []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	got := fizz_buzz.FizzBuzz(15)
	assert.Equal(t, expected, got)
}

func TestFizzBuzzWhenNoPossibleCombosAreAvailable(t *testing.T) {
	expected := []string{"1", "2"}
	got := fizz_buzz.FizzBuzz(2)
	assert.Equal(t, expected, got)
}
