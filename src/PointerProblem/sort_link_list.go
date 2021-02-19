package PointerProblem

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLinkList(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	p := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return dummy.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	right := sortList(slow.Next)
	slow.Next = nil
	left := sortList(head)
	return mergeTwoLinkList(left, right)
}

func swapNode(a, b *ListNode) {
	tmp := a.Val
	a.Val = b.Val
	b.Val = tmp
}

func sortListQuickSortRec(head, end *ListNode) {
	if head == end || head.Next == end {
		return
	}
	slow := head
	fast := head.Next

	pivot := head.Val
	for fast != end {
		if fast.Val < pivot {
			slow = slow.Next
			swapNode(slow, fast)
		}
		fast = fast.Next
	}
	swapNode(head, slow)
	sortListQuickSortRec(head, slow)
	sortListQuickSortRec(slow.Next, end)
}

func sortListQuickSort(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	sortListQuickSortRec(head, nil)
	return head
}
