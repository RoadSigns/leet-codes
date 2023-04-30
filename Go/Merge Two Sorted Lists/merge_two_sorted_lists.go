package merge_two_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	h := &ListNode{Val: 0, Next: nil}
	c := h

	for {
		if list1 == nil || list2 == nil {
			break
		}

		if list1.Val < list2.Val {
			c.Next = list1
			list1 = list1.Next
		} else {
			c.Next = list2
			list2 = list2.Next
		}
		c = c.Next
	}

	if list1 != nil {
		c.Next = list1
	} else {
		c.Next = list2
	}

	return h.Next
}
