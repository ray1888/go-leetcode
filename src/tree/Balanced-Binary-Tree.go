package tree

import "go-leetcode/src/datastructure"

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func getHeight(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func isBalanced(root *datastructure.TreeNode) bool {
	if root == nil {
		return true
	}
	return (isBalanced(root.Left) && isBalanced(root.Right) && (abs(getHeight(root.Left)-getHeight(root.Right))) <= 1)
}
