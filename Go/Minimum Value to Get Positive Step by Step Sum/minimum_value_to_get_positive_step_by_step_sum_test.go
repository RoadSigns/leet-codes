package minimum_value_to_get_positive_step_by_step_sum_test

import (
	"github.com/roadsigns/leet-codes/go/minimum_value_to_get_positive_step_by_step_sum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStartValue(t *testing.T) {
	expected := 5
	got := minimum_value_to_get_positive_step_by_step_sum.MinStartValue([]int{1, -2, -3})
	assert.Equal(t, expected, got)

	expected = 1
	got = minimum_value_to_get_positive_step_by_step_sum.MinStartValue([]int{1, 2})
	assert.Equal(t, expected, got)
}
