package pointersort

import (
	"go-leetcode/src/datastructure"
	"testing"
)

type Test struct {
	Name   string
	Input  *datastructure.ListNode
	Output *datastructure.ListNode
}

func TestInsertionSortList(t *testing.T) {
	tests := []Test{
		{Name: "testing normal", Input: datastructure.GenListNode([]int{1, 3, 2, 4}), Output: datastructure.GenListNode([]int{1, 2, 3, 4})},
	}
	for _, test := range tests {
		output := insertionSortList(test.Input)
		p := test.Output
		q := output.Next
		for p != nil && q != nil {
			if p.Val != q.Val {
				t.Fatalf("Plus one Link list error")
			}
			p = p.Next
			q = q.Next
		}
	}
}

func TestInsertionSortListFast(t *testing.T) {

	tests := []Test{
		{Name: "testing normal", Input: datastructure.GenListNode([]int{1, 3, 2, 4}), Output: datastructure.GenListNode([]int{1, 2, 3, 4})},
	}
	for _, test := range tests {
		output := insertionSortListFast(test.Input)
		p := test.Output
		q := output.Next
		for p != nil && q != nil {
			if p.Val != q.Val {
				t.Fatalf("Plus one Link list error")
			}
			p = p.Next
			q = q.Next
		}
	}
}
