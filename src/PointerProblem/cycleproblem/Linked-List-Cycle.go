package cycleproblem

import "go-leetcode/src/datastructure"

func hasCycle(head *datastructure.ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head
	// TO ensure make two step problem
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}
