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

// split and reverse version
func isPalindromeReverse(head *datastructure.ListNode) bool {
	if head == nil {
		return true
	}
	slow := head
	fast := head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reverseHead := reverse(slow.Next)
	p1 := head
	for reverseHead != nil && p1 != nil {
		if reverseHead.Val != p1.Val {
			return false
		}
		p1 = p1.Next
		reverseHead = reverseHead.Next
	}
	// slow.Next = reverse(reverseHead)
	return true
}

func reverse(head *datastructure.ListNode) *datastructure.ListNode {
	var pre *datastructure.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
