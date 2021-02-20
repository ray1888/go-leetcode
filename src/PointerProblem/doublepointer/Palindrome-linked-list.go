package doublepointer

import "go-leetcode/src/datastructure"

func isPalindrome(head *datastructure.ListNode) bool {
	if head == nil {
		return true
	}

	p := head
	count := 0
	for p != nil {
		p = p.Next
		count++
	}

	cur := head
	pre := &datastructure.ListNode{}

	for i := 0; i < (count / 2); i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	if count%2 == 1 {
		cur = cur.Next
	}
	for cur != nil && pre != nil {
		if pre.Val != cur.Val {
			return false
		}
		pre = pre.Next
		cur = cur.Next
	}
	return true
}
