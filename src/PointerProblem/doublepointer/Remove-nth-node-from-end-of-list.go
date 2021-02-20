package doublepointer

import "go-leetcode/src/datastructure"

func removeNthFromEnd(head *datastructure.ListNode, n int) *datastructure.ListNode {
	dummy := &datastructure.ListNode{}
	dummy.Next = head

	fast := dummy
	slow := dummy

	for n > 0 && fast.Next != nil {
		fast = fast.Next
		n--
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
