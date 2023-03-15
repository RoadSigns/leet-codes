package find_the_index_of_first_occurrence_in_a_string_test

import (
	"github.com/roadsigns/leet-codes/go/find_the_index_of_first_occurrence_in_a_string"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrStr(t *testing.T) {
	expected := 0
	got := find_the_index_of_first_occurrence_in_a_string.StrStr("sadbutsad", "sad")
	assert.Equal(t, expected, got)

	expected = 6
	got = find_the_index_of_first_occurrence_in_a_string.StrStr("badbutsad", "sad")
	assert.Equal(t, expected, got)

	expected = -1
	got = find_the_index_of_first_occurrence_in_a_string.StrStr("leetcode", "leeto")
	assert.Equal(t, expected, got)

	expected = 0
	got = find_the_index_of_first_occurrence_in_a_string.StrStr("z", "z")
	assert.Equal(t, expected, got)
}
