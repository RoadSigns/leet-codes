package node_swap

type ListNode struct {
	Val  int
	Next *ListNode
}

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	r := head
	for head != nil && head.Next != nil {
		p := head.Val
		c := head.Next.Val
		head.Val = c
		head.Next.Val = p

		head = head.Next.Next
	}

	return r
}
