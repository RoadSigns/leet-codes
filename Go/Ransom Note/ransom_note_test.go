package ransom_note_test

import (
	"github.com/roadsigns/leet-codes/go/ransom_note"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	assert.False(t, ransom_note.CanConstruct("a", "b"))
	assert.True(t, ransom_note.CanConstruct("abc", "cab"))
}
