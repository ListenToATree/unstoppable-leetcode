package linkedlist

func removeElements(head *ListNode, val int) *ListNode {
	sentinel := &ListNode{Next: head}

	prev, curr := sentinel, head

	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	return sentinel.Next
}
