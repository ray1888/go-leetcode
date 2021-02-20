package PointerProblem

func reverseList(head *ListNode) *ListNode {
	work := head
	var pre *ListNode
	for work != nil {
		next := work.Next
		work.Next = pre
		pre = work
		work = next
	}
	head = pre
	return head
}
