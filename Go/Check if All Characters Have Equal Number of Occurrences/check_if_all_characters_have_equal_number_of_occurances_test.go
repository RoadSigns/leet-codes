package check_if_all_characters_have_equal_number_of_occurances_test

import (
	"github.com/roadsigns/leet-codes/go/check_if_all_characters_have_equal_number_of_occurances"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckIfPangram(t *testing.T) {
	assert.True(t, check_if_all_characters_have_equal_number_of_occurances.AreOccurrencesEqual("abacbc"))
	assert.False(t, check_if_all_characters_have_equal_number_of_occurances.AreOccurrencesEqual("aaabb"))
}
