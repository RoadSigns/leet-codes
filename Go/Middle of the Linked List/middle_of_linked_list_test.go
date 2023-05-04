package middle_of_linked_list_test

import (
	"github.com/roadsigns/leet-codes/go/middle_of_linked_list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddleOfLinkedListOdd(t *testing.T) {
	expected := 3

	five := middle_of_linked_list.ListNode{
		Val:  5,
		Next: nil,
	}
	four := middle_of_linked_list.ListNode{
		Val:  4,
		Next: &five,
	}
	three := middle_of_linked_list.ListNode{
		Val:  3,
		Next: &four,
	}
	two := middle_of_linked_list.ListNode{
		Val:  2,
		Next: &three,
	}
	one := middle_of_linked_list.ListNode{
		Val:  1,
		Next: &two,
	}

	got := middle_of_linked_list.MiddleNode(&one)
	assert.Equal(t, expected, got.Val)
}

func TestMiddleOfLinkedListEven(t *testing.T) {
	expected := 4

	six := middle_of_linked_list.ListNode{
		Val:  6,
		Next: nil,
	}
	five := middle_of_linked_list.ListNode{
		Val:  5,
		Next: &six,
	}
	four := middle_of_linked_list.ListNode{
		Val:  4,
		Next: &five,
	}
	three := middle_of_linked_list.ListNode{
		Val:  3,
		Next: &four,
	}
	two := middle_of_linked_list.ListNode{
		Val:  2,
		Next: &three,
	}
	one := middle_of_linked_list.ListNode{
		Val:  1,
		Next: &two,
	}

	got := middle_of_linked_list.MiddleNode(&one)
	assert.Equal(t, expected, got.Val)
}
