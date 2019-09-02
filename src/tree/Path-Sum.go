package tree

import "go-leetcode/src/datastructure"

func hasPathSum(root *datastructure.TreeNode, sum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		if root.Val == sum {
			return true
		} else {
			return false
		}
	}

	sum -= root.Val
	return hasPathSum(root.Left, sum) || hasPathSum(root.Right, sum)

}
