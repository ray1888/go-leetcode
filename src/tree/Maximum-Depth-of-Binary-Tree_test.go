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

var symtreeMaxtree *datastructure.TreeNode = &datastructure.TreeNode{
	1,
	&datastructure.TreeNode{2, &datastructure.TreeNode{3, nil, nil}, &datastructure.TreeNode{4, nil, nil}},
	&datastructure.TreeNode{2, &datastructure.TreeNode{4, nil, nil}, &datastructure.TreeNode{3, nil, nil}},
}

var nonsymtreeMaxTree *datastructure.TreeNode = &datastructure.TreeNode{
	1,
	&datastructure.TreeNode{2,
		&datastructure.TreeNode{3, &datastructure.TreeNode{1, nil, nil}, nil},
		&datastructure.TreeNode{4, nil, nil}},
	&datastructure.TreeNode{2,
		&datastructure.TreeNode{4, nil, nil},
		&datastructure.TreeNode{3, nil, nil}},
}

/*
     3
   / \
  9  20
    /  \
   15   7
*/

var maxtree *datastructure.TreeNode = &datastructure.TreeNode{
	3,
	&datastructure.TreeNode{9,
		nil,
		nil},
	&datastructure.TreeNode{20,
		&datastructure.TreeNode{15, nil, nil},
		&datastructure.TreeNode{7, nil, nil}},
}

func TestMaxTree(t *testing.T) {
	var tests = []struct {
		in       *datastructure.TreeNode
		expected int
	}{
		// Add Testcase Here
		{symtreeMaxtree, 3},
		{nonsymtreeMaxTree, 4},
		{maxtree, 3},
	}
	for _, tt := range tests {
		actual := MaxDepthRecursive(tt.in)
		if actual != tt.expected {
			t.Fatalf("the test failed , expected %v,  actual is %v", tt.expected, actual)
		}
	}
}
