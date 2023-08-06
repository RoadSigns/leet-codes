package remove_duplicates_test

import (
	"github.com/roadsigns/learn-go-with-tests/remove_duplicates"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	t.Run("has duplicates", func(t *testing.T) {
		a := &remove_duplicates.ListNode{Val: 3}
		b := &remove_duplicates.ListNode{Val: 2, Next: a}
		c := &remove_duplicates.ListNode{Val: 1, Next: b}
		d := &remove_duplicates.ListNode{Val: 1, Next: c}

		node := remove_duplicates.DeleteDuplicates(d)

		values := make(map[int]bool)
		for node != nil {
			if values[node.Val] {
				t.Errorf("duplicate exists")
			}
			values[node.Val] = true
			node = node.Next
		}

		if len(values) != 3 {
			t.Errorf("incorrect amount of nodes")
		}
	})

	t.Run("has multiple duplicates", func(t *testing.T) {
		a := &remove_duplicates.ListNode{Val: 3}
		b := &remove_duplicates.ListNode{Val: 3, Next: a}
		c := &remove_duplicates.ListNode{Val: 2, Next: b}
		d := &remove_duplicates.ListNode{Val: 1, Next: c}
		e := &remove_duplicates.ListNode{Val: 1, Next: d}

		node := remove_duplicates.DeleteDuplicates(e)

		values := make(map[int]bool)
		got := make([]int, 0)
		for node != nil {
			if values[node.Val] {
				t.Errorf("duplicate exists")
			}
			values[node.Val] = true
			got = append(got, node.Val)
			node = node.Next
		}

		if len(values) != 3 {
			t.Errorf("incorrect amount of nodes")
		}

		want := []int{1, 2, 3}
		for i, _ := range got {
			if got[i] != want[i] {
				t.Errorf("not the expected value")
			}
		}
	})

	t.Run("nil value", func(t *testing.T) {
		got := remove_duplicates.DeleteDuplicates(nil)
		if got != nil {
			t.Errorf("was expecting nil")
		}
	})

	t.Run("just duplicates", func(t *testing.T) {
		d := &remove_duplicates.ListNode{Val: 1}
		e := &remove_duplicates.ListNode{Val: 1, Next: d}

		node := remove_duplicates.DeleteDuplicates(e)

		values := make(map[int]bool)
		got := make([]int, 0)
		for node != nil {
			if values[node.Val] {
				t.Errorf("duplicate exists")
			}
			values[node.Val] = true
			got = append(got, node.Val)
			node = node.Next
		}

		if len(values) != 1 {
			t.Errorf("incorrect amount of nodes")
		}
	})
}
