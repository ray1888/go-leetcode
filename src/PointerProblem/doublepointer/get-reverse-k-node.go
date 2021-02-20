package doublepointer

import "go-leetcode/src/datastructure"

func getKthFromEnd(head *datastructure.ListNode, k int) *datastructure.ListNode {
	if head == nil {
		return head
	}

	fast := head
	slow := head
	for ; k > 0; k-- {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
