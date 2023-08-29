package reverse

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	// Return nil if the list is empty
	if head == nil {
		return nil
	}

	// Initialize a dummy node that points to the head
	dummy := &ListNode{Next: head}
	prev := dummy

	// Move 'prev' to its place which is at position 'left - 1'
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	// Reverse the sublist from position 'left' to 'right'
	prevRev, currRev := reverse(prev.Next, right-left+1)

	// Connect the 'prev.Next' with the reversed sublist and
	// also connect the end of reversed sublist with the remaining list 'currRev'
	prev.Next.Next = currRev
	prev.Next = prevRev

	// Return the next node of dummy node
	return dummy.Next
}

func reverse(head *ListNode, len int) (*ListNode, *ListNode) {
	var prev *ListNode
	// Reverse 'len' nodes of the list and keep track of the next node after the reversed sublist
	for i := 0; i < len; i++ {
		nextRev := head.Next
		head.Next = prev
		prev = head
		head = nextRev
	}
	// Return the head of reversed sublist and the next node after the reversed sublist
	return prev, head
}
