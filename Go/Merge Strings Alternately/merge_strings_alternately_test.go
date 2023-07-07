package merge_strings_alternately_test

import (
	"github.com/roadsigns/leet-codes/go/merge_strings_alternately"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	expected := "apbqcd"
	got := merge_strings_alternately.MergeAlternately("abcd", "pq")
	assert.Equal(t, expected, got)
}
