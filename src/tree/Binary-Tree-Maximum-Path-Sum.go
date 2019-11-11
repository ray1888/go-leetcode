package tree

import (
	"go-leetcode/src/datastructure"
	"math"
)

func maxMax(a, b, c int) int {
	if a > b {
		if c > a {
			return c
		} else {
			return a
		}
	} else {
		if c > b {
			return c
		} else {
			return b
		}
	}
}

func maxPathFrom(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	}
	leftmax := max(maxPathFrom(root.Left), 0)
	rightmax := max(maxPathFrom(root.Right), 0)
	return root.Val + max(leftmax, rightmax)
}

// O(n^2)   space: O(N)
func maxPathSumBrutleForce(root *datastructure.TreeNode) int {
	if root == nil {
		return -1 * int(math.Pow(2, 32))
	}
	leftMax := max(maxPathFrom(root.Left), 0)
	rightMax := max(maxPathFrom(root.Right), 0)
	return maxMax(root.Val+leftMax+rightMax,
		maxPathSumBrutleForce(root.Left),
		maxPathSumBrutleForce(root.Right))
}

func maxPathSumFromRootandCompute(root *datastructure.TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	leftMax := max(maxPathSumFromRootandCompute(root.Left, result), 0)
	rightMax := max(maxPathSumFromRootandCompute(root.Right, result), 0)
	*result = max(*result, root.Val+rightMax+leftMax)
	return root.Val + max(leftMax, rightMax)
}

func maxPathSum(root *datastructure.TreeNode) int {
	if root == nil {
		return -1 * int(math.Pow(2, 32))
	}
	result := -1 * int(math.Pow(2, 32))
	maxPathSumFromRootandCompute(root, &result)
	return result
}
