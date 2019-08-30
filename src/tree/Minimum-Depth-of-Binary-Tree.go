package tree

import (
	"go-leetcode/src/datastructure"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minDepth(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left != nil && root.Right != nil {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	} else {
		if root.Left == nil {
			return minDepth(root.Right) + 1
		} else {
			return minDepth(root.Left) + 1
		}
	}
}

func MinDepthRecursive(root *datastructure.TreeNode) int {
	if root == nil {
		return 1
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}
