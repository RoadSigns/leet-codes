package counting_elements_test

import (
	"github.com/roadsigns/leet-codes/go/counting_elements"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountingElements(t *testing.T) {
	expected := 2
	nums := []int{1, 2, 3}
	got := counting_elements.CountElements(nums)
	assert.Equal(t, expected, got)
}
