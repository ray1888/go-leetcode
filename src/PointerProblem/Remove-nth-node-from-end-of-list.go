package PointerProblem

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	fast := dummy
	slow := dummy

	for n > 0 && fast.Next != nil {
		fast = fast.Next
		n -= 1
	}

	if n != 0 {
		return dummy.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
