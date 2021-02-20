package pointersort

import "go-leetcode/src/datastructure"

func insertionSortList(head *datastructure.ListNode) *datastructure.ListNode {
	var p *datastructure.ListNode
	var next *datastructure.ListNode
	dummy := &datastructure.ListNode{Val: 0, Next: nil}
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

func insertionSortListFast(head *datastructure.ListNode) *datastructure.ListNode {
	var next *datastructure.ListNode
	dummy := &datastructure.ListNode{Val: 0, Next: nil}
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
