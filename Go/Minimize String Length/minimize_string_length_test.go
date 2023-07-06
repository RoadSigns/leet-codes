package minimized_string_length_test

import (
	"github.com/roadsigns/leet-codes/go/minimized_string_length"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountingElements(t *testing.T) {
	expected := 3
	got := minimized_string_length.MinimizedStringLength("aaabc")
	assert.Equal(t, expected, got)

	expected = 3
	got = minimized_string_length.MinimizedStringLength("cbbd")
	assert.Equal(t, expected, got)

	expected = 2
	got = minimized_string_length.MinimizedStringLength("dddaaa")
	assert.Equal(t, expected, got)
}
