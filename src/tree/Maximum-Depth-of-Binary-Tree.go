package tree

import (
	"go-leetcode/src/datastructure"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxDepth(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left != nil && root.Right != nil {
		return max(minDepth(root.Left), minDepth(root.Right)) + 1
	} else {
		if root.Left == nil {
			return maxDepth(root.Right) + 1
		} else {
			return maxDepth(root.Left) + 1
		}
	}
}

func MaxDepthRecursive(root *datastructure.TreeNode) int {
	if root == nil {
		return 1
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
