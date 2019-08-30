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
var successInput *datastructure.TreeNode = &datastructure.TreeNode{
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

var failedInput *datastructure.TreeNode = &datastructure.TreeNode{
	1,
	&datastructure.TreeNode{2, nil, &datastructure.TreeNode{3, nil, nil}},
	&datastructure.TreeNode{2, nil, &datastructure.TreeNode{3, nil, nil}},
}

func TestIsSymetric(t *testing.T) {
	var tests = []struct {
		in       *datastructure.TreeNode
		expected bool
	}{
		// Add Testcase Here
		{successInput, true},
		{failedInput, false},
	}
	for _, tt := range tests {
		actual := IsSymmetricRecursive(tt.in)
		if actual != tt.expected {
			t.Fatalf("the test failed , expected %v,  actual is %v", tt.expected, actual)
		}
	}
}
