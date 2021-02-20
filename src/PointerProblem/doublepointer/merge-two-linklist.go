package doublepointer

import "go-leetcode/src/datastructure"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeInBetween(list1 *datastructure.ListNode, a int, b int, list2 *datastructure.ListNode) *datastructure.ListNode {

	var paPre *datastructure.ListNode
	var pb *datastructure.ListNode

	work1 := list1
	work2 := list1

	for i := 0; i <= a-1; i++ {
		if work1 == nil {
			break
		}
		if i == a-1 {
			paPre = work1
		}
		work1 = work1.Next
	}

	for i := 0; i <= b; i++ {
		if work2 == nil {
			break
		}
		if i == b {
			pb = work2
		}
		work2 = work2.Next
	}

	paPre.Next = list2
	k2 := list2
	for k2.Next != nil {
		k2 = k2.Next
	}
	k2.Next = pb.Next
	return list1
}
