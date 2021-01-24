package PointerProblem

/*
	此处处理的方法，先把指针复制一份连接下去，然后通过移动指针把多于的部分截断
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	n, p := 1, head
	for p.Next != nil {
		p = p.Next
		n++
	}
	p.Next = head
	k %= n
	for i := 1; i <= n-k; i++ {
		p = p.Next
	}
	head, p.Next = p.Next, nil
	return head
}
