package PointerProblem

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast := head
	slow := head
	// TO ensure make two step problem
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
