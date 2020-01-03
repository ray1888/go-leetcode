package PointerProblem

import "testing"

type Test struct {
	Name   string
	Input  *ListNode
	Output *ListNode
}

func genListNode(data []int) *ListNode {
	head := new(ListNode)
	p := head
	for _, value := range data {
		p.Val = value
		p.Next = new(ListNode)
		p = head.Next
	}
	return head
}

func TestOddEvenList(t *testing.T) {
	tests := []Test{
		{Name: "testing normal", Input: genListNode([]int{1, 2, 3, 4}), Output: genListNode([]int{1, 3, 2, 4})},
		{Name: "testing double nine", Input: genListNode([]int{3, 6, 2, 8, 7}), Output: genListNode([]int{3, 2, 7, 6, 8})},
	}
	for _, test := range tests {
		output := oddEvenList(test.Input)
		p := test.Input
		q := output
		for p != nil && q != nil {
			if p.Val != q.Val {
				t.Fatalf("Plus one Link list error")
			}
			p = p.Next
			q = q.Next
		}
	}
}
