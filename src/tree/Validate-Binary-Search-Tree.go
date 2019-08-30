package tree

import "go-leetcode/src/datastructure"

func GetMax(root *datastructure.TreeNode) *datastructure.TreeNode {
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func GetMin(root *datastructure.TreeNode) *datastructure.TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func isValidBST(root *datastructure.TreeNode) bool {
	if root == nil {
		return true
	}
	leftValid := (root.Left == nil) || root.Val > GetMax(root.Left).Val
	rightValid := (root.Right == nil) || root.Val < GetMin(root.Right).Val
	return leftValid && rightValid && isValidBST(root.Left) && isValidBST(root.Right)
}
