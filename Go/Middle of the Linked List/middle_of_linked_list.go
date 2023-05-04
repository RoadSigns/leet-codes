package middle_of_linked_list

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MiddleNode(head *ListNode) *ListNode {
	count := 0.0
	originalHead := head
	for head.Next != nil {
		count++
		head = head.Next
	}

	newHead := originalHead
	middle := int(math.Ceil(count / 2))
	for i := 0; i < middle; i++ {
		newHead = newHead.Next
	}

	return newHead
}
