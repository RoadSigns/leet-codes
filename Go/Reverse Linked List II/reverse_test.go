package reverse_test

import (
	reverse "github.com/roadsigns/leet_codes/reverse_linked_list_ii"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	five := &reverse.ListNode{
		Val:  5,
		Next: nil,
	}
	four := &reverse.ListNode{
		Val:  4,
		Next: five,
	}
	three := &reverse.ListNode{
		Val:  3,
		Next: four,
	}
	two := &reverse.ListNode{
		Val:  2,
		Next: three,
	}
	one := &reverse.ListNode{
		Val:  1,
		Next: two,
	}
	head := reverse.ReverseBetween(one, 2, 4)

	want := []int{1, 4, 3, 2, 5}
	got := make([]int, 0)
	for head != nil {
		got = append(got, head.Val)
		head = head.Next
	}

	assert.ElementsMatch(t, want, got)
	for i := 0; i < len(want); i++ {
		assert.Equal(t, want[i], got[i])
	}
}
