package doublepointer

import "go-leetcode/src/datastructure"

func kthToLast(head *datastructure.ListNode, k int) int {
	if head == nil || k == 0 {
		return 0
	}
	work := head
	slow := head
	for ; k > 0; k-- {
		if work == nil {
			return 0
		}
		work = work.Next
	}

	for work != nil {
		slow = slow.Next
		work = work.Next
	}
	return slow.Val
}
