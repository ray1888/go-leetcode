package tree

import "go-leetcode/src/datastructure"

func preorder(root *datastructure.TreeNode, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	preorder(root.Left, result)
	preorder(root.Right, result)
}

func preorderTraversal(root *datastructure.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	preorder(root, &result)
	return result
}
