package linkedlist

//// recursive
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	}
//	if list2 == nil {
//		return list1
//	}
//
//	var head *ListNode
//
//	if list1.Val < list2.Val {
//		head = list1
//		list1 = list1.Next
//	} else {
//		head = list2
//		list2 = list2.Next
//	}
//
//	head.Next = mergeTwoLists(list1, list2)
//	return head
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList = &ListNode{}
	var out = newList
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			newList.Next, l1 = l1, l1.Next
			newList = newList.Next
		} else {
			newList.Next, l2 = l2, l2.Next
			newList = newList.Next
		}
	}
	if l1 != nil {
		newList.Next = l1
	} else if l2 != nil {
		newList.Next = l2
	}

	return out.Next
}
