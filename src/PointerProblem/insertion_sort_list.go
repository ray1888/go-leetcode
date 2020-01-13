package PointerProblem

func insertionSortList(head *ListNode) *ListNode {
	var p *ListNode
	var next *ListNode
	dummy := &ListNode{0, nil}
	cur := head
	for cur != nil {
		next = cur.Next
		p = dummy
		for p.Next != nil && p.Next.Val <= cur.Val {
			p = p.Next
		}
		cur.Next = p.Next
		p.Next = cur
		cur = next
	}
	return dummy.Next
}

func insertionSortListFast(head *ListNode) *ListNode {
	var next *ListNode
	dummy := &ListNode{0, nil}
	p := dummy
	cur := head
	for cur != nil {
		next = cur.Next
		if p.Next != nil && cur.Val < p.Next.Val {
			p = dummy
		}
		for p.Next != nil && p.Next.Val <= cur.Val {
			p = p.Next
		}
		cur.Next = p.Next
		p.Next = cur
		cur = next
	}
	return dummy.Next
}
