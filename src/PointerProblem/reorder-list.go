package PointerProblem

func reorderListIterSpaceO1(head *ListNode) {
	if head == nil {
		return
	}
	work := head
	for work.Next != nil && work.Next.Next != nil {
		slow := work
		fast := work
		fast = fast.Next
		for fast.Next != nil {
			slow = slow.Next
			fast = fast.Next
		}
		tmp := work.Next
		work.Next = fast
		fast.Next = tmp
		slow.Next = nil
		work = work.Next.Next
	}
	return
}

func reorderListIterSpaceOn(head *ListNode) {
	if head == nil {
		return
	}
	nodes := make([]*ListNode, 0)
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	i := 0
	j := len(nodes) - 1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
