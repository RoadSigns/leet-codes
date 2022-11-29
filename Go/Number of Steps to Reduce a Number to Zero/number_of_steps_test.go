package number_of_steps_test

import (
	"github.com/roadsigns/leet-codes/go/number_of_steps"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanWorkOutEvenNumber(t *testing.T) {
	expected := 4
	got := number_of_steps.NumberOfSteps(8)

	assert.Equal(t, expected, got)
}

func TestCanWorkOutOddNumber(t *testing.T) {
	expected := 12
	got := number_of_steps.NumberOfSteps(123)

	assert.Equal(t, expected, got)
}

func TestCanWorkOutZero(t *testing.T) {
	expected := 0
	got := number_of_steps.NumberOfSteps(0)

	assert.Equal(t, expected, got)
}
