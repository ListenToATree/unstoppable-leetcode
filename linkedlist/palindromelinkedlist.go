package linkedlist

//// Time: O(2n) Space: O(n)
//func isPalindrome(head *ListNode) bool {
//	arr := []int{}
//	for head != nil {
//		arr = append(arr, head.Val)
//		head = head.Next
//	}
//
//	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
//		if arr[i] != arr[j] {
//			return false
//		}
//	}
//	return true
//}

// Time: O(4n) Space: O(1)
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverse(firstHalfEnd.Next)

	p1, p2 := head, secondHalfStart
	res := true
	for res && p2 != nil {
		if p1.Val != p2.Val {
			res = false
		}
		p1, p2 = p1.Next, p2.Next
	}

	firstHalfEnd.Next = reverse(secondHalfStart)
	return res
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head
	for curr != nil {
		curr.Next, prev, curr = prev, curr, curr.Next
	}
	return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast.Next != nil || fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
