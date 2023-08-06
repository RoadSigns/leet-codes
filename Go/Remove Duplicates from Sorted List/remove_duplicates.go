package remove_duplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	s := head

	for s != nil {
		if s.Next != nil {
			if s.Val == s.Next.Val {
				s.Next = s.Next.Next
				continue
			}
		}
		s = s.Next
	}

	return head
}
