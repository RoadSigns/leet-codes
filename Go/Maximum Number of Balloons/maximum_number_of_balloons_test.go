package maximum_number_of_balloons_test

import (
	"github.com/roadsigns/leet-codes/go/maximum_number_of_balloons"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxNumberOfBalloons(t *testing.T) {
	expected := 1
	got := maximum_number_of_balloons.MaxNumberOfBalloons("nlaebolko")
	assert.Equal(t, expected, got)

	expected = 2
	got = maximum_number_of_balloons.MaxNumberOfBalloons("loonbalxballpoon")
	assert.Equal(t, expected, got)

	expected = 0
	got = maximum_number_of_balloons.MaxNumberOfBalloons("leetcode")
	assert.Equal(t, expected, got)
}

func TestMaxNumberOfBalloonsOneLoop(t *testing.T) {
	expected := 1
	got := maximum_number_of_balloons.MaxNumberOfBalloonsOneLoop("nlaebolko")
	assert.Equal(t, expected, got)

	expected = 2
	got = maximum_number_of_balloons.MaxNumberOfBalloonsOneLoop("loonbalxballpoon")
	assert.Equal(t, expected, got)

	expected = 0
	got = maximum_number_of_balloons.MaxNumberOfBalloonsOneLoop("leetcode")
	assert.Equal(t, expected, got)
}
