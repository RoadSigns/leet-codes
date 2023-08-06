package node_swap_test

import (
	"github.com/roadsigns/learn-go-with-tests/node_swap"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	t.Run("swap pairs", func(t *testing.T) {
		want := []int{2, 1, 4, 3}
		a := &node_swap.ListNode{Val: 4}
		b := &node_swap.ListNode{Val: 3, Next: a}
		c := &node_swap.ListNode{Val: 2, Next: b}
		d := &node_swap.ListNode{Val: 1, Next: c}

		result := node_swap.SwapPairs(d)

		got := make([]int, 0)
		for result != nil {
			got = append(got, result.Val)
			result = result.Next
		}

		for i, _ := range want {
			if got[i] != want[i] {
				t.Errorf("expected %d, got %d", want[i], got[i])
			}
		}
	})
}
