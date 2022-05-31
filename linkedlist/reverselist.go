package linkedlist

//// Time: O(n) Space: O(1)
//// iterative
//func reverseList(head *ListNode) *ListNode {
//	var prev *ListNode = nil
//	cur := head
//
//	for cur != nil {
//		// save the next node before break
//		tmp := cur.Next
//		// point the cur to prev
//		cur.Next = prev
//		// move one step further
//		prev = cur
//		cur = tmp
//
//		// in one line
//		//cur.Next, prev, cur = prev, cur, cur.Next
//	}
//
//	return prev
//}

// Time: O(n) Space: O(1)
// recursive
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversed := reverseList(head.Next)
	head.Next.Next, head.Next = head, nil

	return reversed
}
