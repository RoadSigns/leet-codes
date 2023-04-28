package valid_parentheses_test

import (
	"github.com/roadsigns/leet-codes/go/valid_parentheses"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert.True(t, valid_parentheses.IsValid("{}"))
	assert.True(t, valid_parentheses.IsValid("()[]{}"))
	assert.False(t, valid_parentheses.IsValid("{"))
	assert.False(t, valid_parentheses.IsValid("(]"))
}
