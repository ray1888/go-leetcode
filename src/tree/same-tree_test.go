package tree

import (
	"go-leetcode/src/datastructure"
	"testing"
)

/*
   1
  / \
 2   2
/ \ / \
3  4 4  3
*/
var sameTree1 *datastructure.TreeNode = &datastructure.TreeNode{
	1,
	&datastructure.TreeNode{2, &datastructure.TreeNode{3, nil, nil}, &datastructure.TreeNode{4, nil, nil}},
	&datastructure.TreeNode{2, &datastructure.TreeNode{4, nil, nil}, &datastructure.TreeNode{3, nil, nil}},
}

/*

     1
   / \
  2   2
   \   \
   3    3
*/

var sameTree2 *datastructure.TreeNode = &datastructure.TreeNode{
	1,
	&datastructure.TreeNode{2, nil, &datastructure.TreeNode{3, nil, nil}},
	&datastructure.TreeNode{2, nil, &datastructure.TreeNode{3, nil, nil}},
}

func TestIsSameTree(t *testing.T) {
	var tests = []struct {
		in       []*datastructure.TreeNode
		expected bool
	}{
		// Add Testcase Here
		{[]*datastructure.TreeNode{sameTree1, sameTree1}, true},
		{[]*datastructure.TreeNode{sameTree1, sameTree2}, false},
	}
	for _, tt := range tests {
		actual := isSameTreeRecursive(tt.in[0], tt.in[1])
		if actual != tt.expected {
			t.Fatalf("the test failed , expected %v,  actual is %v", tt.expected, actual)
		}
	}
}
