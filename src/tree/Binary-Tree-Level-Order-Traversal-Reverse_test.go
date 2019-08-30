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

var symtreeLevelR *datastructure.TreeNode = &datastructure.TreeNode{
	1,
	&datastructure.TreeNode{2, &datastructure.TreeNode{3, nil, nil}, &datastructure.TreeNode{4, nil, nil}},
	&datastructure.TreeNode{2, &datastructure.TreeNode{4, nil, nil}, &datastructure.TreeNode{3, nil, nil}},
}

var nonsymtreeLevelR *datastructure.TreeNode = &datastructure.TreeNode{
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

var leveltreeR *datastructure.TreeNode = &datastructure.TreeNode{
	3,
	&datastructure.TreeNode{9,
		nil,
		nil},
	&datastructure.TreeNode{20,
		&datastructure.TreeNode{15, nil, nil},
		&datastructure.TreeNode{7, nil, nil}},
}

func TestLevelRTravel(t *testing.T) {
	var tests = []struct {
		in       *datastructure.TreeNode
		expected [][]int
	}{
		// Add Testcase Here
		{symtreeLevelR, [][]int{[]int{3, 4, 4, 3}, []int{2, 2}, []int{1}}},
		{nonsymtreeLevelR, [][]int{[]int{1}, []int{3, 4, 4, 3}, []int{2, 2}, []int{1}}},
		{leveltreeR, [][]int{[]int{15, 7}, []int{9, 20}, []int{3}}},
	}
	for _, tt := range tests {
		actual := levelOrderBottomEasy(tt.in)
		if len(actual) != len(tt.expected) {
			t.Fatalf("the test failed , level transfer failed, level num is not match, expected %d, actual %d",
				len(tt.expected), len(actual))
		}
		for i := 0; i < len(actual); i++ {
			if len(actual[i]) != len(tt.expected[i]) {
				t.Fatalf("the test failed , level %d , item number not match, expected %v, actual %v",
					i, tt.expected[i], actual[i])
			}
			for j := 0; j < len(actual[i]); j++ {
				if actual[i][j] != tt.expected[i][j] {
					t.Fatalf("the test failed , expected %v,  actual is %v", tt.expected[i], actual[i])
				}
			}
		}
	}
}
