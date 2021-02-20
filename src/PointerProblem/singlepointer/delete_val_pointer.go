package singlepointer

import "go-leetcode/src/datastructure"

func deleteNode(head *datastructure.ListNode, val int) *datastructure.ListNode {
	if head == nil {
		return head
	}
	dummy := &datastructure.ListNode{}
	last := dummy
	dummy.Next = head
	work := head

	for work != nil {
		if work.Val == val {
			if work.Next != nil {
				last.Next = work.Next
			} else {
				last.Next = nil
			}
			return dummy.Next
		}
		last = work
		work = work.Next
	}
	return dummy.Next
}
