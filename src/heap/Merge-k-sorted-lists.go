package heap

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKListsMinHeap(lists []*ListNode) *ListNode {

}

func mergeTwoList(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for left != nil && right != nil {
		if left.Val > right.Val {
			cur.Next = right
			right = right.Next
		} else {
			cur.Next = left
			left = left.Next
		}
		cur = cur.Next
	}
	if left != nil {
		cur.Next = left
		left = left.Next
		cur = cur.Next
	}
	if right != nil {
		cur.Next = right
		right = right.Next
		cur = cur.Next
	}
	return dummy.Next
}

func merge(nodelist []*ListNode, start, end int) *ListNode {
	if start == end {
		return nodelist[start]
	} else if start > end {
		return nil
	}
	mid := start + (end-start)/2
	left := merge(nodelist, start, mid)
	right := merge(nodelist, mid+1, end)
	return mergeTwoList(left, right)
}

func mergeKListsDivideConq(lists []*ListNode) *ListNode {
	if len(lists) == 0 || lists == nil {
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}
