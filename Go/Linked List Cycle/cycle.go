package cycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	f := head
	s := head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
		if s == f {
			return true
		}
	}

	return false
}
