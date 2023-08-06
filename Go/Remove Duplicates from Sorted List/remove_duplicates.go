package remove_duplicates

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	values := make(map[int]bool)
	s := &ListNode{Val: head.Val}
	result := s
	values[s.Val] = true
	head = head.Next

	for head != nil {
		if values[head.Val] {
			head = head.Next
			continue
		}

		values[head.Val] = true
		s.Next = &ListNode{Val: head.Val}
		head = head.Next
		s = s.Next
	}

	return result
}
