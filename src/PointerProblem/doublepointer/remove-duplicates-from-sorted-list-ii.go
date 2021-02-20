package doublepointer

import "go-leetcode/src/datastructure"

func deleteDuplicates(head *datastructure.ListNode) *datastructure.ListNode {
	if head == nil {
		return head
	}

	dummy := new(datastructure.ListNode)
	prev := dummy
	prev.Next = head
	cur := prev.Next
	for cur != nil {
		// 先把重复的给去除只剩余一个
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		// 把最后一个也给跳掉
		if prev.Next != cur {
			prev.Next = cur.Next
		} else {
			prev = prev.Next
		}
		cur = prev.Next
	}
	return dummy.Next
}
