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
func isLeafNode(node *datastructure.TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func dfs(node *datastructure.TreeNode) (ans int) {
	if node.Left != nil {
		if isLeafNode(node.Left) {
			ans += node.Left.Val
		} else {
			ans += dfs(node.Left)
		}
	}
	if node.Right != nil && !isLeafNode(node.Right) {
		ans += dfs(node.Right)
	}
	return
}

func sumOfLeftLeaves(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	}
	return dfs(root)
}
