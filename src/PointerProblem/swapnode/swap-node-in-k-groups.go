package swapnode

import "go-leetcode/src/datastructure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *datastructure.ListNode, k int) *datastructure.ListNode {
	if head == nil {
		return head
	}

	dummy := new(datastructure.ListNode)
	dummy.Next = head
	pre := dummy
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		nex := tail.Next
		head, tail = myReverse(head, tail)
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}

	return dummy.Next
}

func myReverse(head, tail *datastructure.ListNode) (*datastructure.ListNode, *datastructure.ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}
