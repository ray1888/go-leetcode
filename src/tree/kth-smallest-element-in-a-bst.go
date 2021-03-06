package tree

import "go-leetcode/src/datastructure"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *datastructure.TreeNode, k int) int {

	res := 0
	stack := []*datastructure.TreeNode{}
	var cur *datastructure.TreeNode = root

	for len(stack) != 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}

		cur = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		k--
		if k == 0 {
			return cur.Val
		}

		cur = cur.Right
	}
	return res

}
