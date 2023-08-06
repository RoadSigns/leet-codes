package cycle_test

import (
	"github.com/roadsigns/learn-go-with-tests/cycle"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasCycle(t *testing.T) {
	t.Run("has cycle", func(t *testing.T) {
		one := &cycle.ListNode{Val: 1}
		two := &cycle.ListNode{Val: 2}
		one.Next = two
		two.Next = one

		assert.True(t, cycle.HasCycle(one))
	})

	t.Run("doesn't have cycle", func(t *testing.T) {
		four := &cycle.ListNode{Val: -4}
		zero := &cycle.ListNode{Val: 0, Next: four}
		two := &cycle.ListNode{Val: 2, Next: zero}
		three := &cycle.ListNode{Val: 3, Next: two}

		assert.False(t, cycle.HasCycle(three))
	})

	t.Run("single node", func(t *testing.T) {
		one := &cycle.ListNode{Val: 1}
		assert.False(t, cycle.HasCycle(one))
	})
}
