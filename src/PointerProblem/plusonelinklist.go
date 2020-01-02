package PointerProblem

func PlusOneLinkList(head *ListNode) *ListNode {
	maybe := new(ListNode)
	notNine := maybe
	maybe.Next = head
	p := head
	for p != nil {
		if p.Val != 9 {
			notNine = p
		}
		p = p.Next
	}
	notNine.Val += 1
	p = notNine.Next
	for p != nil {
		p.Val = 0
		p = p.Next
	}
	if notNine == maybe {
		return maybe
	} else {
		return head
	}
}
