package PointerProblem

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	evenHead := head.Next
	evenwork := evenHead
	for evenwork != nil && evenwork.Next != nil {
		odd.Next = evenwork.Next
		odd = odd.Next
		evenwork.Next = odd.Next
		evenwork = evenwork.Next
	}
	odd.Next = evenHead
	return head
}
