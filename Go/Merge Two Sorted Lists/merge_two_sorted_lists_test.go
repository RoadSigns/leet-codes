package merge_two_sorted_lists_test

import (
	"github.com/roadsigns/leet-codes/go/merge_two_sorted_lists"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	four := merge_two_sorted_lists.ListNode{
		Val:  4,
		Next: nil,
	}

	three := merge_two_sorted_lists.ListNode{
		Val:  3,
		Next: &four,
	}

	two := merge_two_sorted_lists.ListNode{
		Val:  2,
		Next: &four,
	}

	one1 := merge_two_sorted_lists.ListNode{
		Val:  1,
		Next: &two,
	}

	one2 := merge_two_sorted_lists.ListNode{
		Val:  1,
		Next: &three,
	}

	// Expecting [1,1,2,3,4,4]
	expected := merge_two_sorted_lists.MergeTwoLists(&one1, &one2)

	assert.Equal(t, expected.Val, 1)
	expected = expected.Next
	assert.Equal(t, expected.Val, 1)
	expected = expected.Next
	assert.Equal(t, expected.Val, 2)
	expected = expected.Next
	assert.Equal(t, expected.Val, 3)
	expected = expected.Next
	assert.Equal(t, expected.Val, 4)
	expected = expected.Next
	assert.Equal(t, expected.Val, 4)
}
