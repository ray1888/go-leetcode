package swapnode

import "go-leetcode/src/datastructure"

func swapPairs(head *datastructure.ListNode) *datastructure.ListNode {
	if head == nil {
		return nil
	}
	dummy := new(datastructure.ListNode)
	dummy.Next = head
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := cur.Next.Next
		cur.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		cur = node1
	}
	return dummy.Next
}
