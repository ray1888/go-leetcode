package tree

import (
	"go-leetcode/src/datastructure"
)

func isSymmetric(left, right *datastructure.TreeNode) bool {
	if left != nil && right != nil {
		return left.Val == right.Val &&
			isSymmetric(left.Left, right.Right) &&
			isSymmetric(right.Left, left.Right)
	} else {
		return left == nil && right == nil
	}
}

func IsSymmetricRecursive(root *datastructure.TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetric(root.Left, root.Right)
}
