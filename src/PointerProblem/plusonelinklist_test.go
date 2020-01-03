package PointerProblem

import "testing"

func TestPlusOneLinkList(t *testing.T) {
	tests := []Test{
		{Name: "testing normal", Input: genListNode([]int{1, 2, 3}), Output: genListNode([]int{1, 2, 4})},
		{Name: "testing double nine", Input: genListNode([]int{1, 2, 9, 9}), Output: genListNode([]int{1, 3, 0, 0})},
		{Name: "testing single nine", Input: genListNode([]int{1, 2, 9}), Output: genListNode([]int{1, 3, 0})},
		{Name: "testing all nine", Input: genListNode([]int{9, 9, 9}), Output: genListNode([]int{1, 0, 0, 0})},
	}
	for _, test := range tests {
		output := PlusOneLinkList(test.Input)
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
