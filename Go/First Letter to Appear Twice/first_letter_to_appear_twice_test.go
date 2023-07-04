package first_letter_to_appear_twice_test

import (
	"github.com/roadsigns/leet-codes/go/first_letter_to_appear_twice"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeatedCharacter(t *testing.T) {
	expected := []byte("d")[0]
	got := first_letter_to_appear_twice.RepeatedCharacter("abcdd")
	assert.Equal(t, expected, got, "Expected %v, got %v", expected, got)
}
